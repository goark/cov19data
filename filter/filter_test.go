package filter

import (
	"testing"
	"time"

	"github.com/spiegel-im-spiegel/cov19data/values"
)

func TestPeriod(t *testing.T) {
	testCases := []struct {
		period values.Period
		date   values.Date
		res    bool
	}{
		{
			period: values.NewPeriod(
				values.NewDate(2020, time.Month(8), 1),
				values.NewDate(2020, time.Month(8), 31),
			),
			date: values.NewDate(2020, time.Month(7), 31),
			res:  false,
		},
		{
			period: values.NewPeriod(
				values.NewDate(2020, time.Month(8), 1),
				values.NewDate(2020, time.Month(8), 31),
			),
			date: values.NewDate(2020, time.Month(8), 1),
			res:  true,
		},
		{
			period: values.NewPeriod(
				values.NewDate(2020, time.Month(8), 1),
				values.NewDate(2020, time.Month(8), 31),
			),
			date: values.NewDate(2020, time.Month(8), 31),
			res:  true,
		},
		{
			period: values.NewPeriod(
				values.NewDate(2020, time.Month(8), 1),
				values.NewDate(2020, time.Month(8), 31),
			),
			date: values.NewDate(2020, time.Month(9), 1),
			res:  false,
		},
	}

	for _, tc := range testCases {
		res := New(WithPeriod(tc.period)).Period(tc.date)
		if res != tc.res {
			t.Errorf("Filters.Period(\"%v\") != \"%v\", want \"%v\".", tc.date, res, tc.res)
		}
	}
}

func TestCountryCode(t *testing.T) {
	testCases := []struct {
		fc  values.CountryCode
		c   values.CountryCode
		res bool
	}{
		{fc: values.CC_JP, c: values.CC_JP, res: true},
		{fc: values.CC_JP, c: values.CC_US, res: false},
	}

	for _, tc := range testCases {
		res := New(WithCountryCode(tc.fc)).CountryCode(tc.c)
		if res != tc.res {
			t.Errorf("Filters.CountryCode(\"%v\") != \"%v\", want \"%v\".", tc.c, res, tc.res)
		}
	}
}

func TestRegionCode(t *testing.T) {
	testCases := []struct {
		fc  values.RegionCode
		c   values.RegionCode
		res bool
	}{
		{fc: values.WPRO, c: values.WPRO, res: true},
		//{fc: values.WPRO, c: values.AFRO, res: false},
		{fc: values.WPRO, c: values.AFRO, res: true},
	}

	for _, tc := range testCases {
		res := New(WithRegionCode(tc.fc)).RegionCode(tc.c)
		if res != tc.res {
			t.Errorf("Filters.CountryCode(\"%v\") != \"%v\", want \"%v\".", tc.c, res, tc.res)
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
