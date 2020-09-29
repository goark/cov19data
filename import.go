package cov19data

import (
	"io"
	"os"

	"github.com/spiegel-im-spiegel/cov19data/client"
	"github.com/spiegel-im-spiegel/cov19data/csvdata"
	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/cov19data/entity"
	"github.com/spiegel-im-spiegel/cov19data/filter"
	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
)

//Import class
type Import struct {
	reader  io.Reader
	csvdata *csvdata.Reader
}

//New returns new Import instance
func New(r io.Reader) *Import {
	return &Import{reader: r, csvdata: csvdata.New(r, 8, true)}
}

//NewWeb returns new Import instance
func NewWeb(c *client.Client) (*Import, error) {
	r, err := c.Get("https://covid19.who.int/WHO-COVID-19-global-data.csv")
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return New(r), nil
}

//Close method close reader if it has io.Closer interface.
func (i *Import) Close() {
	if i == nil {
		return
	}
	if c, ok := i.reader.(io.Closer); ok {
		c.Close()
	}
}

//RawReader method returns raw data stream
func (i *Import) RawReader() io.Reader {
	if i == nil {
		return nil
	}
	return i.reader
}

//Data method returns entity.GlobalData list
func (i *Import) Data(opts ...filter.FiltersOptFunc) ([]*entity.GlobalData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	filter := filter.New(opts...)
	records := []*entity.GlobalData{}
	for {
		record, err := i.next()
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

//Data method returns entity.GlobalData list
func (i *Import) Histogram(period values.Period, step int, opts ...filter.FiltersOptFunc) ([]*histogram.HistData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	if step < 1 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("period", period.String()), errs.WithContext("step", step))
	}
	histList, p := histogram.NewList(period, step)
	if len(histList) == 0 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("period", period.String()), errs.WithContext("step", step))
	}
	period = p
	filter := filter.New(opts...)
	for {
		record, err := i.next()
		if err != nil {
			if errs.Is(err, ecode.ErrNoData) {
				break
			}
			return nil, errs.Wrap(err)
		}
		if record.CheckFilter(filter) {
			histogram.AddData(histList, record.Date, record.NewCases, record.NewDeaths)
		}
	}

	return histList, nil
}

func (i *Import) next() (*entity.GlobalData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	elms, err := i.csvdata.Next()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return entity.New(elms[0], elms[1], elms[3], elms[4], elms[5], elms[6], elms[7])
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
