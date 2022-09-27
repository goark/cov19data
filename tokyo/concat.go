package tokyo

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/goark/csvdata"
	"github.com/goark/errs"
)

type concat struct {
	buf *bytes.Buffer
}

func newConcat() *concat {
	return &concat{&bytes.Buffer{}}
}

func (cc *concat) cat(r io.Reader, withHeader bool) error {
	data := csvdata.NewRows(csvdata.New(r).WithTrimSpace(true), true)
	if withHeader {
		hdr, err := data.Header()
		if err != nil {
			return errs.Wrap(err)
		}
		fmt.Fprintln(cc.buf, strings.Join(hdr, ","))
	}
	for {
		if err := data.Next(); err != nil {
			if errs.Is(err, io.EOF) {
				return nil
			}
			return errs.Wrap(err)
		}
		fmt.Fprintln(cc.buf, strings.Join(data.Row(), ","))
	}
}

func (cc *concat) reader() io.Reader {
	return cc.buf
}

/* Copyright 2022 Spiegel
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
