package values

import (
	"strconv"
	"strings"
)

//RegionCode is WHO region code
type RegionCode int

const (
	UnknownRO RegionCode = iota // Unknown region
	AFRO                        // Africa
	AMRO                        // Americas
	EMRO                        // Eastern Mediterranean
	EURO                        // Europe
	SEARO                       // South-East Asia
	WPRO                        // Western Pacific
	OtherRO                     // Other
)

var regionCodeMap = map[RegionCode]string{
	AFRO:    "AFRO",  // Africa
	AMRO:    "AMRO",  // Americas
	EMRO:    "EMRO",  // Eastern Mediterranean
	EURO:    "EURO",  // Europe
	SEARO:   "SEARO", // South-East Asia
	WPRO:    "WPRO",  // Western Pacific
	OtherRO: "Other", // Other
}

var regionNameMap = map[RegionCode]string{
	AFRO:    "Africa",
	AMRO:    "Americas",
	EMRO:    "Eastern Mediterranean",
	EURO:    "Europe",
	SEARO:   "South-East Asia",
	WPRO:    "Western Pacific",
	OtherRO: "Other",
}

//GetRegionCode function returns RegionCode instrance from string.
func GetRegionCode(s string) RegionCode {
	s = strings.TrimSpace(s)
	for k, v := range regionCodeMap {
		if strings.EqualFold(v, s) {
			return k
		}
	}
	return UnknownRO
}

func (rc RegionCode) String() string {
	if s, ok := regionCodeMap[rc]; ok {
		return s
	}
	return ""
}

//Name method returns country name.
func (rc RegionCode) Name() string {
	if s, ok := regionNameMap[rc]; ok {
		return s
	}
	return ""
}

//UnmarshalJSON method returns result of Unmarshal for json.Unmarshal().
func (rc *RegionCode) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*rc = GetRegionCode(s)
	return nil
}

//MarshalJSON method returns string for json.Marshal().
func (rc *RegionCode) MarshalJSON() ([]byte, error) {
	if rc == nil {
		return []byte(`""`), nil
	}
	return []byte(strconv.Quote(rc.String())), nil
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
