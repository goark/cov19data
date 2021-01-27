package entity

import (
	"encoding/json"

	"github.com/spiegel-im-spiegel/cov19data/filter"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

//JapanData is entity class for japan COVID-19 forecast data by Google
type JapanData struct {
	Date                 values.Date       //発生日
	ForecastDate         values.Date       //計測日
	PrefCode             values.PrefJpCode //都道府県コード
	PrefName             string            //都道府県（ローマ字）
	PrefNamekanji        string            //都道府県（漢字）
	ForecastFlag         bool              //true なら予測値, false なら 実測値
	NewCases             json.Number       //追加感染者数（日別）
	CumulativeCases      json.Number       //感染者のべ数（実測値のみ）
	NewDeaths            json.Number       //追加死者数（日別）
	CumulativeDeaths     json.Number       //死者総数（実測値のみ）
	HospitalizedPatients json.Number       //入院・療養等患者数（実測値のみ・日別）
	Recovered            json.Number       //回復者数（実測値のみ・日別）
}

func New(prefCode, prefname, date, cumulativeCases, cumulativeDeaths, hospitalizedPatients, recovered, forecastDate, newDeathsForecast, newCasesForecast, newDeathsTrue, newCasesTrue, prefNamekanji string) (*JapanData, error) {
	dt, err := values.NewDateString(date)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("date", date))
	}
	fdt, err := values.NewDateString(forecastDate)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("forecastDate", forecastDate))
	}
	fflag := false
	newDeaths := newDeathsTrue
	newCases := newCasesTrue
	if dt.After(fdt) {
		fflag = true
		newDeaths = newDeathsForecast
		newCases = newCasesForecast
	}
	return &JapanData{
		Date:                 dt,                                //発生日
		ForecastDate:         fdt,                               //計測日
		PrefCode:             values.GetPrefJpCode(prefCode),    //都道府県コード
		PrefName:             prefname,                          //都道府県（ローマ字）
		PrefNamekanji:        prefNamekanji,                     //都道府県（漢字）
		ForecastFlag:         fflag,                             //true なら予測値, false なら 実測値
		NewCases:             json.Number(newCases),             //追加感染者数（日別）
		CumulativeCases:      json.Number(cumulativeCases),      //感染者のべ数（実測値のみ）
		NewDeaths:            json.Number(newDeaths),            //追加死者数（日別）
		CumulativeDeaths:     json.Number(cumulativeDeaths),     //死者総数（実測値のみ）
		HospitalizedPatients: json.Number(hospitalizedPatients), //入院・療養等患者数（実測値のみ・日別）
		Recovered:            json.Number(recovered),            //回復者数（実測値のみ・日別）
	}, nil
}

//CheckFilter method returns true if cheking filter is OK.
func (d *JapanData) CheckFilter(filter *filter.Filters) bool {
	return filter.Period(d.Date) && filter.PrefJpCode(d.PrefCode) && filter.RegionCode(values.WPRO)
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
