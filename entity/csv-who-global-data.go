package entity

import (
	"io"

	"github.com/spiegel-im-spiegel/cov19data/csvdata"
	"github.com/spiegel-im-spiegel/errs"
)

//CsvReaderWHO is a class of reader for WHO global CSV data
type CsvReaderWHO struct {
	*csvdata.Reader
}

//NewCsvReaderWHO creates a new CsvReaderWHO instance.
func NewCsvReaderWHO(r io.Reader) *CsvReaderWHO {
	return &CsvReaderWHO{csvdata.New(r, 8, true)}
}

//Next method returns next record of WHOGlobalData
func (cr *CsvReaderWHO) Next() (WHOGlobalData, error) {
	elms, err := cr.Reader.Next()
	if err != nil {
		return WHOGlobalData{}, errs.Wrap(err)
	}
	return newWHOGlobalData(elms[0], elms[1], elms[3], elms[4], elms[5], elms[6], elms[7])
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
