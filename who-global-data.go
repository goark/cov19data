package cov19data

import (
	"github.com/spiegel-im-spiegel/cov19data/client"
	"github.com/spiegel-im-spiegel/cov19data/entity"
	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

//ImportWHOCSV function returns WHO COVID-19 global data from WHO site.
func ImportWHOCSV(c *client.Client, opts ...entity.FiltersOptFunc) ([]entity.WHOGlobalData, error) {
	r, err := c.WHOCasesData()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer r.Close()
	return entity.ImportWHOCSV(r, opts...)
}

//MakeHistogramWHO function returns WHO COVID-19 global data from WHO site.
func MakeHistogramWHO(c *client.Client, period values.Period, span int, opts ...entity.FiltersOptFunc) ([]*histogram.HistData, error) {
	r, err := c.WHOCasesData()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	defer r.Close()
	return histogram.MakeHistgramWHOFromCSV(r, period, span, opts...)
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