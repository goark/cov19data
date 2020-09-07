package cov19data

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

//HistData is class of cases data record for histgram.
type HistData struct {
	Period values.Period
	Cases  int64
	Deaths int64
}

//NewHistData function creates a new HistData instance.
func NewHistData(period values.Period, cases, deaths int64) *HistData {
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
func (h *HistData) AddCases(cases int64) *HistData {
	if h == nil {
		return nil
	}
	h.Cases += cases
	return h
}

//AddDeaths method adds deaths count in HistData
func (h *HistData) AddDeaths(deaths int64) *HistData {
	if h == nil {
		return nil
	}
	h.Deaths += deaths
	return h
}

func ExportHistCSV(data []*HistData) ([]byte, error) {
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
			strconv.FormatInt(d.Cases, 10),
			strconv.FormatInt(d.Deaths, 10),
		}); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	cw.Flush()
	return buf.Bytes(), nil
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
