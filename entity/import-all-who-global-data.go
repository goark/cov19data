package entity

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/errs"
)

//ImportWHOCSV function creates list of WHOGlobalData from CSV.
func ImportWHOCSV(r io.Reader, opts ...FiltersOptFunc) ([]WHOGlobalData, error) {
	filter := NewFilters(opts...)
	records := []WHOGlobalData{}
	cr := NewCsvReaderWHO(r)
	for {
		record, err := cr.Next()
		if err != nil {
			if errs.Is(err, ecode.ErrNoData) {
				break
			}
			return nil, errs.Wrap(err)
		}
		if record.CheckFilter(filter) {
			records = append(records, record)
		}
	}
	return records, nil
}

//ExportWHOCSV function returns CSV string from list of WHOGlobalData.
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

//ExportWHOJSON function returns JSON string from list of WHOGlobalData.
func ExportWHOJSON(data []WHOGlobalData) ([]byte, error) {
	if len(data) == 0 {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	return json.Marshal(data)
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
