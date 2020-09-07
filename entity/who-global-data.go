package entity

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

//WHOGlobalData is entity class for WHO COVID-19 global data
type WHOGlobalData struct {
	Date             values.Date
	CountryCode      values.CountryCode
	WHORegion        values.RegionCode
	NewCases         json.Number
	CumulativeCases  json.Number
	NewDeaths        json.Number
	CumulativeDeaths json.Number
}

func newWHOGlobalData(date, countryCode, regionCode, newCases, cumulativeCases, newDeaths, cumulativeDeaths string) (WHOGlobalData, error) {
	dt, err := values.NewDateString(date)
	if err != nil {
		return WHOGlobalData{}, errs.Wrap(err, errs.WithContext("date", date))
	}
	return WHOGlobalData{
		Date:             dt,
		CountryCode:      values.GetCountryCode(countryCode),
		WHORegion:        values.GetRegionCode(regionCode),
		NewCases:         json.Number(newCases),
		CumulativeCases:  json.Number(cumulativeCases),
		NewDeaths:        json.Number(newDeaths),
		CumulativeDeaths: json.Number(cumulativeDeaths),
	}, nil

}

func ExportWHOJSON(data []WHOGlobalData) ([]byte, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	return json.Marshal(data)
}

func ExportWHOCSV(data []WHOGlobalData) ([]byte, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	buf := &bytes.Buffer{}
	cw := csv.NewWriter(buf)
	cw.Comma = ','
	if err := cw.Write([]string{
		"Date_reported",
		"Country_code",
		"Country",
		"WHO_region",
		"New_cases",
		"Cumulative_cases",
		"New_deaths",
		"Cumulative_deaths",
	}); err != nil {
		return nil, errs.Wrap(err)
	}
	for _, d := range data {
		if err := cw.Write([]string{
			d.Date.String(),
			d.CountryCode.String(),
			d.CountryCode.Name(),
			d.WHORegion.String(),
			string(d.NewCases),
			string(d.CumulativeCases),
			string(d.NewDeaths),
			string(d.CumulativeDeaths),
		}); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	cw.Flush()
	return buf.Bytes(), nil
}

func ImportWHOCSV(r io.Reader, opts ...FiltersOptFunc) ([]WHOGlobalData, error) {
	filter := NewFilters(opts...)
	records := []WHOGlobalData{}
	cr := csv.NewReader(r)
	cr.Comma = ','
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading
	header := true
	for {
		elms, err := cr.Read()
		if err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return nil, errs.Wrap(err)
		}
		if len(elms) < 8 {
			return nil, errs.Wrap(ecode.ErrInvalidRecord, errs.WithContext("record", elms))
		}
		if !header {
			record, err := newWHOGlobalData(elms[0], elms[1], elms[3], elms[4], elms[5], elms[6], elms[7])
			if err != nil {
				return nil, errs.Wrap(err)
			}
			if filter.Period(record.Date) && filter.CountryCode(record.CountryCode) && filter.RegionCode(record.WHORegion) {
				records = append(records, record)
			}
		} else {
			header = false
		}
	}
	return records, nil
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
