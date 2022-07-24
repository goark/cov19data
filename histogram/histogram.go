package histogram

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"sort"
	"strconv"

	"github.com/goark/cov19data/ecode"
	"github.com/goark/cov19data/values"
	"github.com/goark/errs"
)

//HistData is class of cases data record for histgram.
type HistData struct {
	Period values.Period
	Cases  float64
	Deaths float64
}

//New function creates a new HistData instance.
func New(period values.Period, cases, deaths float64) *HistData {
	return &HistData{Period: period, Cases: cases, Deaths: deaths}
}

//Contains method returns true if scape of this contains date of parameter.
func (h *HistData) Contains(dt values.Date) bool {
	if h == nil {
		return false
	}
	return h.Period.Contains(dt)
}

//AddCases method adds cases count in HistData
func (h *HistData) AddCases(cases float64) *HistData {
	if h == nil {
		return nil
	}
	h.Cases += cases
	return h
}

//AddDeaths method adds deaths count in HistData
func (h *HistData) AddDeaths(deaths float64) *HistData {
	if h == nil {
		return nil
	}
	h.Deaths += deaths
	return h
}

//NewList creates list of HistData.
func NewList(p values.Period, step int) ([]*HistData, values.Period) {
	histList := []*HistData{}
	max := values.Period{}
	if p.IsZero() {
		return histList, max
	}
	if step < 1 {
		return histList, max
	}
	start := p.Start
	end := p.End
	next := end
	for {
		to := next
		next = to.AddDay(-step)
		from := next.AddDay(1)
		histList = append(histList, New(values.NewPeriod(from, to), 0, 0))
		if values.NewPeriod(from, to).Contains(p.Start) {
			break
		}
		start = from
	}
	sort.Slice(histList, func(i, j int) bool {
		return histList[i].Period.End.Before(histList[j].Period.End)
	})
	return histList, values.NewPeriod(start, end)
}

//AddData adds data into HistData list.
func AddData(histList []*HistData, dt values.Date, cases, deaths json.Number) {
	for _, h := range histList {
		if h.Period.Contains(dt) {
			if n, err := cases.Float64(); err == nil {
				h.AddCases(n)
			}
			if n, err := deaths.Float64(); err == nil {
				h.AddDeaths(n)
			}
			return
		}
	}
}

//ExportHistCSV exports CSV string from list of HistData.
func ExportCSV(data []*HistData) ([]byte, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	buf := &bytes.Buffer{}
	cw := csv.NewWriter(buf)
	cw.Comma = ','
	if err := cw.Write([]string{
		"Date_from",
		"Date_to",
		"Cases",
		"Deaths",
	}); err != nil {
		return nil, errs.Wrap(err)
	}
	for _, d := range data {
		if err := cw.Write([]string{
			d.Period.StringStart(),
			d.Period.StringEnd(),
			strconv.FormatFloat(d.Cases, 'f', -1, 64),
			strconv.FormatFloat(d.Deaths, 'f', -1, 64),
		}); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	cw.Flush()
	return buf.Bytes(), nil
}

//ExportJSON function returns JSON string from list of HistData.
func ExportJSON(data []*HistData) ([]byte, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	return json.Marshal(data)
}

/* Copyright 2020-2021 Spiegel
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
