package google

import (
	"context"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/cov19data/csvdata"
	"github.com/spiegel-im-spiegel/cov19data/ecode"
	"github.com/spiegel-im-spiegel/cov19data/filter"
	"github.com/spiegel-im-spiegel/cov19data/google/entity"
	"github.com/spiegel-im-spiegel/cov19data/histogram"
	"github.com/spiegel-im-spiegel/cov19data/values"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/fetch"
)

//Import class
type Import struct {
	reader  io.Reader
	csvdata *csvdata.Reader
}

//New returns new Import instance
func New(r io.Reader) *Import {
	return &Import{reader: r, csvdata: csvdata.New(r, 25, true)}
}

//NewWeb returns new Import instance
func NewWeb(ctx context.Context, cli fetch.Client) (*Import, error) {
	u, err := fetch.URL("https://storage.googleapis.com/covid-external/forecast_JAPAN_PREFECTURE_28.csv")
	if err != nil {
		return nil, errs.Wrap(err)
	}
	resp, err := cli.Get(u, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return New(resp.Body()), nil
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
func (i *Import) Data(opts ...filter.FiltersOptFunc) ([]*entity.JapanData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	filter := filter.New(opts...)
	records := []*entity.JapanData{}
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
	entity.Sort(records)
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

func (i *Import) next() (*entity.JapanData, error) {
	if i == nil {
		return nil, errs.Wrap(ecode.ErrNoData)
	}
	elms, err := i.csvdata.Next()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return entity.New(elms[0], elms[1], elms[2], elms[15], elms[16], elms[17], elms[18], elms[19], elms[20], elms[21], elms[22], elms[23], elms[24])
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
