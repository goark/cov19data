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

//MakeHistgramTokyo function create list of HistData from entity.TokyoData list
func MakeHistgramTokyo(data []entity.TokyoData, step int) ([]*HistData, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	if step < 1 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("step", step))
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].Date.Before(data[j].Date)
	})
	histList := NewHistList(values.NewPeriod(data[0].Date, data[len(data)-1].Date), step)
	for _, d := range data {
		setHistTokyoData(histList, d)
	}
	return histList, nil
}

//MakeHistgramTokyoFromCSV function create list of HistData from CSV
func MakeHistgramTokyoFromCSV(r io.Reader, period values.Period, step int, opts ...entity.FiltersOptFunc) ([]*HistData, error) {
	if step < 1 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("step", step))
	}
	histList := NewHistList(period, step)
	if len(histList) > 0 {
		period = values.NewPeriod(histList[0].Period.Start, period.End)
	}
	filter := entity.NewFilters(append(opts, entity.WithFilterPeriod(period))...)
	cr := entity.NewCsvReaderTokyo(r)
	for {
		record, err := cr.Next()
		if err != nil {
			if errs.Is(err, ecode.ErrNoData) {
				break
			}
			return nil, errs.Wrap(err)
		}
		if record.CheckFilter(filter) {
			setHistTokyoData(histList, record)
		}
	}
	return histList, nil
}

func setHistTokyoData(histList []*HistData, data entity.TokyoData) {
	for _, h := range histList {
		if h.Period.Contains(data.Date) {
			h.AddCases(1)
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
