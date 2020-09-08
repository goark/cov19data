package client_test

import (
	"encoding/csv"
	"fmt"
	"strings"

	"github.com/spiegel-im-spiegel/cov19data/client"
)

var bom = []byte{0xEF, 0xBB, 0xBF}

func trimBOM(s string) string {
	return strings.Trim(s, string(bom))
}

func ExampleWHOCasesData() {
	resp, err := client.Default().WHOCasesData()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Close()

	cr := csv.NewReader(resp)
	cr.Comma = ','
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading

	elms, err := cr.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(trimBOM(strings.Join(elms, ",")))
	//Output:
	//Date_reported,Country_code,Country,WHO_region,New_cases,Cumulative_cases,New_deaths,Cumulative_deaths
}

func ExampleTokyoPatientsData() {
	resp, err := client.Default().TokyoPatientsData()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Close()

	cr := csv.NewReader(resp)
	cr.Comma = ','
	cr.LazyQuotes = true       // a quote may appear in an unquoted field and a non-doubled quote may appear in a quoted field.
	cr.TrimLeadingSpace = true // leading

	elms, err := cr.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(trimBOM(strings.Join(elms, ",")))
	//Output:
	//No,全国地方公共団体コード,都道府県名,市区町村名,公表_年月日,曜日,発症_年月日,患者_居住地,患者_年代,患者_性別,患者_属性,患者_状態,患者_症状,患者_渡航歴の有無フラグ,備考,退院済フラグ
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
