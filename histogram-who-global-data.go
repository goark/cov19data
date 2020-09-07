package cov19data

import (
	"os"
	"sort"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/cov19data/entity"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

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
	histList := []*HistData{}
	start := data[0].Date
	end := data[len(data)-1].Date
	next := end
	flag := false
	for {
		to := next
		next = to.AddDay(-span)
		from := next.AddDay(1)
		if values.NewPeriod(from, to).Contains(start) {
			from = start
			flag = true
		}
		histList = append(histList, NewHistData(values.NewPeriod(from, to), 0, 0))
		if flag {
			break
		}

	}
	sort.Slice(histList, func(i, j int) bool {
		return histList[i].Period.StringEnd() < histList[j].Period.StringEnd()
	})
	for _, d := range data {
		for _, h := range histList {
			if h.Period.Contains(d.Date) {
				if n, err := d.NewCases.Int64(); err == nil {
					h.AddCases(n)
				}
				if n, err := d.NewDeaths.Int64(); err == nil {
					h.AddDeaths(n)
				}
				break
			}
		}
	}
	return histList, nil
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
