package tokyo

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"github.com/goark/cov19data/ecode"
	"github.com/goark/cov19data/filter"
	"github.com/goark/cov19data/histogram"
	"github.com/goark/cov19data/tokyo/entity"
	"github.com/goark/cov19data/values"
	"github.com/goark/csvdata"
	"github.com/goark/errs"
	"github.com/goark/fetch"
)

var urls = []string{
	"https://stopcovid19.metro.tokyo.lg.jp/data/130001_tokyo_covid19_patients_2020.csv",
	"https://stopcovid19.metro.tokyo.lg.jp/data/130001_tokyo_covid19_patients_2021.csv",
	"https://stopcovid19.metro.tokyo.lg.jp/data/130001_tokyo_covid19_patients_2022.csv",
	"https://stopcovid19.metro.tokyo.lg.jp/data/130001_tokyo_covid19_patients_2022-1.csv",
}

// Import class
type Import struct {
	reader io.Reader
	data   *csvdata.Rows
}

// New returns new Import instance
func New(r io.Reader) *Import {
	return &Import{reader: r, data: csvdata.NewRows(csvdata.New(r).WithTrimSpace(true).WithFieldsPerRecord(17), true)}
}

// NewWeb returns new Import instance
func NewWeb(ctx context.Context, cli fetch.Client) (*Import, error) {
	cc := newConcat()
	for i, u := range urls {
		if err := importFromWeb(ctx, cli, u, cc, i == 0); err != nil {
			return nil, errs.Wrap(err)
		}
	}
	return New(cc.reader()), nil
}

func importFromWeb(ctx context.Context, cli fetch.Client, ustr string, cc *concat, withHeader bool) error {
	u, err := fetch.URL(ustr)
	if err != nil {
		return errs.Wrap(err)
	}
	resp, err := cli.Get(u, fetch.WithContext(ctx))
	if err != nil {
		return errs.Wrap(err)
	}
	defer resp.Close()
	return cc.cat(resp.Body(), withHeader)
}

// Close method close reader if it has io.Closer interface.
func (i *Import) Close() {
	if i == nil {
		return
	}
	i.data.Close()
}

// RawReader method returns raw data stream
func (i *Import) RawReader() io.Reader {
	if i == nil {
		return nil
	}
	return i.reader
}

// Data method returns entity.GlobalData list
func (i *Import) Data(opts ...filter.FiltersOptFunc) ([]*entity.TokyoData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	filter := filter.New(opts...)
	records := []*entity.TokyoData{}
	for {
		record, err := i.next()
		if err != nil {
			if errs.Is(err, io.EOF) {
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

// Data method returns entity.GlobalData list
func (i *Import) Histogram(period values.Period, step int, opts ...filter.FiltersOptFunc) ([]*histogram.HistData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	if step < 1 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("period", period.String()), errs.WithContext("step", step))
	}
	histList, _ := histogram.NewList(period, step)
	if len(histList) == 0 {
		return nil, errs.Wrap(os.ErrInvalid, errs.WithContext("period", period.String()), errs.WithContext("step", step))
	}
	filter := filter.New(opts...)
	for {
		record, err := i.next()
		if err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return nil, errs.Wrap(err)
		}
		if record.CheckFilter(filter) {
			histogram.AddData(histList, record.Date, json.Number("1"), json.Number("0"))
		}
	}

	return histList, nil
}

func (i *Import) next() (*entity.TokyoData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	if err := i.data.Next(); err != nil {
		return nil, errs.Wrap(err)
	}
	return entity.New(
		i.data.Get(4),
		i.data.Get(1),
		i.data.Get(7),
		i.data.Get(8),
		i.data.Get(9),
		i.data.Get(15),
	)
}

/* Copyright 2020-2022 Spiegel
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
