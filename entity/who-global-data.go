package entity

import (
	"encoding/json"

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

//CheckFilter method returns true if cheking filter is OK.
func (d WHOGlobalData) CheckFilter(filter *Filters) bool {
	return filter.Period(d.Date) && filter.CountryCode(d.CountryCode) && filter.RegionCode(d.WHORegion)
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
