package entity

import (
	"io"

	"github.com/spiegel-im-spiegel/cov19data/csvdata"
	"github.com/spiegel-im-spiegel/errs"
)

//CsvReaderTokyo is a class of reader for Tokyo COVID-19 patients CSV data
type CsvReaderTokyo struct {
	*csvdata.Reader
}

//NewCsvReaderTokyo creates a new CsvReaderTokyo instance.
func NewCsvReaderTokyo(r io.Reader) *CsvReaderTokyo {
	return &CsvReaderTokyo{csvdata.New(r, 16, true)}
}

//Next method returns next record of TokyoData
func (cr *CsvReaderTokyo) Next() (TokyoData, error) {
	elms, err := cr.Reader.Next()
	if err != nil {
		return TokyoData{}, errs.Wrap(err)
	}
	return newTokyoData(elms[4], elms[1], elms[7], elms[8], elms[9], elms[15])
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
