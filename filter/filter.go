package filter

import "github.com/spiegel-im-spiegel/cov19data/values"

//Filters is a filter class for entity classes
type Filters struct {
	periods      []values.Period
	countryCodes []values.CountryCode
	regionCodes  []values.RegionCode
}

//FiltersOptFunc type is self-referential function type for New functions. (functional options pattern)
type FiltersOptFunc func(*Filters)

//New function returns a  new Filters instance with options.
func New(opts ...FiltersOptFunc) *Filters {
	f := &Filters{
		periods:      []values.Period{},
		countryCodes: []values.CountryCode{},
		regionCodes:  []values.RegionCode{},
	}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

//WithPeriod function returns FiltersOptFunc function value.
//This function is used in New functions that represents Marketplace data.
func WithPeriod(period values.Period) FiltersOptFunc {
	return func(f *Filters) {
		if f != nil {
			existFlag := false
			for _, p := range f.periods {
				if p.Equal(period) {
					existFlag = true
					break
				}
			}
			if !existFlag {
				f.periods = append(f.periods, period)
			}
		}
	}
}

//WithCountryCode function returns FiltersOptFunc function value.
//This function is used in New functions that represents Marketplace data.
func WithCountryCode(code values.CountryCode) FiltersOptFunc {
	return func(f *Filters) {
		if f != nil {
			existFlag := false
			for _, c := range f.countryCodes {
				if c == code {
					existFlag = true
					break
				}
			}
			if !existFlag {
				f.countryCodes = append(f.countryCodes, code)
			}
		}
	}
}

//WithRegionCode function returns FiltersOptFunc function value.
//This function is used in New functions that represents Marketplace data.
func WithRegionCode(code values.RegionCode) FiltersOptFunc {
	return func(f *Filters) {
		if f != nil {
			existFlag := false
			for _, c := range f.regionCodes {
				if c == code {
					existFlag = true
					break
				}
			}
			if !existFlag {
				f.regionCodes = append(f.regionCodes, code)
			}
		}
	}
}

//Period method returns true if Periods filter contains date of parameter.
func (f *Filters) Period(dt values.Date) bool {
	if f == nil {
		return true
	}
	if len(f.periods) == 0 {
		return true
	}
	for _, p := range f.periods {
		if p.Contains(dt) {
			return true
		}
	}
	return false
}

//CountryCode method returns true if CountryCodes filter contains date of parameter.
func (f *Filters) CountryCode(code values.CountryCode) bool {
	if f == nil {
		return true
	}
	if len(f.countryCodes) == 0 {
		return true
	}
	for _, c := range f.countryCodes {
		if c == code {
			return true
		}
	}
	return false
}

//CountryCode method returns true if CountryCodes filter contains date of parameter.
func (f *Filters) RegionCode(code values.RegionCode) bool {
	if f == nil {
		return true
	}
	if len(f.regionCodes) == 0 {
		return true
	}
	for _, c := range f.regionCodes {
		if c == code {
			return true
		}
	}
	return false
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
