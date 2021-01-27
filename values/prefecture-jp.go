package values

import (
	"fmt"
	"strconv"
	"strings"
)

//PrefJpCode is prefecture code in Japan
type PrefJpCode uint

func GetPrefJpCode(s string) PrefJpCode {
	s = strings.TrimPrefix(strings.ToUpper(strings.TrimSpace(s)), "JP-")
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return PrefJpCode(0)
	}
	if n < 1 || n > 47 {
		return PrefJpCode(0)
	}
	return PrefJpCode(uint(n))
}

func (pc PrefJpCode) String() string {
	return fmt.Sprintf("%02d", uint(pc))
}

//UnmarshalJSON method returns result of Unmarshal for json.Unmarshal().
func (pc *PrefJpCode) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*pc = GetPrefJpCode(s)
	return nil
}

//MarshalJSON method returns string for json.Marshal().
func (pc PrefJpCode) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(pc.String())), nil
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
