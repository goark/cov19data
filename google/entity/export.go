package entity

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/cov19data/filter"
	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

func Sort(data []*JapanData) {
	sort.Slice(data, func(i, j int) bool {
		if data[i].PrefCode < data[j].PrefCode {
			return true
		}
		if data[i].PrefCode > data[j].PrefCode {
			return false
		}
		return data[i].Date.Before(data[j].Date)
	})
}

//ExportCSV function returns CSV string from list of WHOGlobalData.
func ExportCSV(data []*JapanData) ([]byte, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	buf := &bytes.Buffer{}
	cw := csv.NewWriter(buf)
	cw.Comma = ','
	if err := cw.Write([]string{
		"target_prediction_date",
		"forecast_date",
		"japan_prefecture_code",
		"prefecture_name",
		"prefecture_name_kanji",
		"forecast_flag",
		"cumulative_confirmed",
		"cumulative_deaths",
		"new_confirmed",
		"new_deaths",
		"hospitalized_patients",
		"recovered",
	}); err != nil {
		return nil, errs.Wrap(err)
	}
	for _, d := range data {
		if err := cw.Write([]string{
			d.Date.String(),
			d.ForecastDate.String(),
			"JP-" + d.PrefCode.String(),
			d.PrefName,
			d.PrefNamekanji,
			fmt.Sprintf("%v", d.ForecastFlag),
			string(d.CumulativeCases),
			string(d.CumulativeDeaths),
			string(d.NewCases),
			string(d.NewDeaths),
			string(d.HospitalizedPatients),
			string(d.Recovered),
		}); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	cw.Flush()
	return buf.Bytes(), nil
}

//ExportJSON function returns JSON string from list of WHOGlobalData.
func ExportJSON(data []*JapanData) ([]byte, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	return json.Marshal(data)
}

//ExportJSON function returns JSON string from list of WHOGlobalData.
func ExportHistgram(data []*JapanData, period values.Period, step int, opts ...filter.FiltersOptFunc) ([]*histogram.HistData, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	if step < 1 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("period", period.String()), errs.WithContext("step", step))
	}
	histList, p := histogram.NewList(period, step)
	if len(histList) == 0 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("period", period.String()), errs.WithContext("step", step))
	}
	period = p
	filter := filter.New(opts...)
	for _, record := range data {
		if record != nil && record.CheckFilter(filter) {
			histogram.AddData(histList, record.Date, record.NewCases, record.NewDeaths)
		}

	}
	return histList, nil
}

/* Copyright 2021 Spiegel
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
