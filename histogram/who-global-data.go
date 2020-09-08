package histogram

import (
	"io"
	"os"
	"sort"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/cov19data/entity"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

//MakeHistgramWHO function create list of HistData from entity.WHOGlobalData list
func MakeHistgramWHO(data []entity.WHOGlobalData, span int) ([]*HistData, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	if span < 1 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("span", span))
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.Before(data[j].Date)
	})
	histList := NewHistList(values.NewPeriod(data[0].Date, data[len(data)-1].Date), span)
	for _, d := range data {
		setHistWHOData(histList, d)
	}
	return histList, nil
}

//MakeHistgramWHOFromCSV function create list of HistData from CSV
func MakeHistgramWHOFromCSV(r io.Reader, period values.Period, span int, opts ...entity.FiltersOptFunc) ([]*HistData, error) {
	if span < 1 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("span", span))
	}
	filter := entity.NewFilters(append(opts, entity.WithFilterPeriod(period))...)
	histList := NewHistList(period, span)
	cr := entity.NewCsvReaderWHO(r)
	for {
		record, err := cr.Next()
		if err != nil {
			if errs.Is(err, ecode.ErrNoData) {
				break
			}
			return nil, errs.Wrap(err)
		}
		if record.CheckFilter(filter) {
			setHistWHOData(histList, record)
		}
	}
	return histList, nil
}

func setHistWHOData(histList []*HistData, data entity.WHOGlobalData) {
	for _, h := range histList {
		if h.Period.Contains(data.Date) {
			if n, err := data.NewCases.Int64(); err == nil {
				h.AddCases(n)
			}
			if n, err := data.NewDeaths.Int64(); err == nil {
				h.AddDeaths(n)
			}
			return
		}
	}
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
