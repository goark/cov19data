package entity

import (
	"github.com/goark/cov19data/filter"
	"github.com/goark/cov19data/values"
	"github.com/goark/errs"
)

//WHOGlobalData is entity class for Tokyo COVID-19 patients data
type TokyoData struct {
	Date         values.Date //発生日付
	LocalGovCode string      //地方公共団体コード
	Address      string      //対象者の居住地
	Age          string      //対象者の年代
	Gender       string      //対象者の性別
	LeaveFlag    string      //退院フラグ（1:退院または死亡）
}

func New(date, localGovCode, address, age, gender, leaveFlag string) (*TokyoData, error) {
	dt, err := values.NewDateString(date)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("date", date))
	}
	return &TokyoData{
		Date:         dt,
		LocalGovCode: localGovCode,
		Address:      address,
		Age:          age,
		Gender:       gender,
		LeaveFlag:    leaveFlag,
	}, nil
}

//CheckFilter method returns true if cheking filter is OK.
func (d *TokyoData) CheckFilter(filter *filter.Filters) bool {
	return filter.Period(d.Date) && filter.CountryCode(values.CC_JP) && filter.RegionCode(values.WPRO)
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
