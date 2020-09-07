package client

import (
	"errors"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/cov19data/ecode"
)

func TestIllegalURL(t *testing.T) {
	testCases := []struct {
		rawurl string
		err    error
	}{
		{rawurl: "", err: ecode.ErrInvalidRequest},
		{rawurl: ":::", err: ecode.ErrInvalidRequest},
		{rawurl: "foo", err: ecode.ErrInvalidRequest},
		{rawurl: "https://baldanders.info/not-exist.txt", err: ecode.ErrHTTPStatus},
	}

	for _, tc := range testCases {
		if _, err := ((*Client)(nil)).Get(tc.rawurl); !errors.Is(err, tc.err) {
			t.Errorf("\"%v\" != \"%v\"", err, tc.err)
		} else {
			fmt.Printf("Info: %+v\n", err)
		}
	}
}

func TestZeroValue(t *testing.T) {
	testCases := []struct {
		err error
	}{
		{err: ecode.ErrInvalidRequest},
	}

	for _, tc := range testCases {
		if _, err := (&Client{}).WHOCasesData(); !errors.Is(err, tc.err) {
			t.Errorf("\"%v\" != \"%v\"", err, tc.err)
		} else {
			fmt.Printf("Info: %+v\n", err)
		}
	}
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
