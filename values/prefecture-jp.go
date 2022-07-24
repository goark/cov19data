package values

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

//PrefJpCode is prefecture code in Japan
type PrefJpCode uint

const (
	PCJP_UNKNOWN PrefJpCode = iota
	PCJP_HOKKAIDO
	PCJP_AOMORI
	PCJP_IWATE
	PCJP_MIYAGI
	PCJP_AKITA
	PCJP_YAMAGATA
	PCJP_FUKUSHIMA
	PCJP_IBARAKI
	PCJP_TOCHIGI
	PCJP_GUNMA
	PCJP_SAITAMA
	PCJP_CHIBA
	PCJP_TOKYO
	PCJP_KANAGAWA
	PCJP_NIIGATA
	PCJP_TOYAMA
	PCJP_ISHIKAWA
	PCJP_FUKUI
	PCJP_YAMANASHI
	PCJP_NAGANO
	PCJP_GIFU
	PCJP_SHIZUOKA
	PCJP_AICHI
	PCJP_MIE
	PCJP_SHIGA
	PCJP_KYOTO
	PCJP_OSAKA
	PCJP_HYOGO
	PCJP_NARA
	PCJP_WAKAYAMA
	PCJP_TOTTORI
	PCJP_SHIMANE
	PCJP_OKAYAMA
	PCJP_HIROSHIMA
	PCJP_YAMAGUCHI
	PCJP_TOKUSHIMA
	PCJP_KAGAWA
	PCJP_EHIME
	PCJP_KOCHI
	PCJP_FUKUOKA
	PCJP_SAGA
	PCJP_NAGASAKI
	PCJP_KUMAMOTO
	PCJP_OITA
	PCJP_MIYAZAKI
	PCJP_KAGOSHIMA
	PCJP_OKINAWA
)

var (
	PCJP_MIN = PCJP_HOKKAIDO
	PCJP_MAX = PCJP_OKINAWA
	pcNames  = map[PrefJpCode]string{
		PCJP_HOKKAIDO:  "HOKKAIDO",
		PCJP_AOMORI:    "AOMORI",
		PCJP_IWATE:     "IWATE",
		PCJP_MIYAGI:    "MIYAGI",
		PCJP_AKITA:     "AKITA",
		PCJP_YAMAGATA:  "YAMAGATA",
		PCJP_FUKUSHIMA: "FUKUSHIMA",
		PCJP_IBARAKI:   "IBARAKI",
		PCJP_TOCHIGI:   "TOCHIGI",
		PCJP_GUNMA:     "GUNMA",
		PCJP_SAITAMA:   "SAITAMA",
		PCJP_CHIBA:     "CHIBA",
		PCJP_TOKYO:     "TOKYO",
		PCJP_KANAGAWA:  "KANAGAWA",
		PCJP_NIIGATA:   "NIIGATA",
		PCJP_TOYAMA:    "TOYAMA",
		PCJP_ISHIKAWA:  "ISHIKAWA",
		PCJP_FUKUI:     "FUKUI",
		PCJP_YAMANASHI: "YAMANASHI",
		PCJP_NAGANO:    "NAGANO",
		PCJP_GIFU:      "GIFU",
		PCJP_SHIZUOKA:  "SHIZUOKA",
		PCJP_AICHI:     "AICHI",
		PCJP_MIE:       "MIE",
		PCJP_SHIGA:     "SHIGA",
		PCJP_KYOTO:     "KYOTO",
		PCJP_OSAKA:     "OSAKA",
		PCJP_HYOGO:     "HYOGO",
		PCJP_NARA:      "NARA",
		PCJP_WAKAYAMA:  "WAKAYAMA",
		PCJP_TOTTORI:   "TOTTORI",
		PCJP_SHIMANE:   "SHIMANE",
		PCJP_OKAYAMA:   "OKAYAMA",
		PCJP_HIROSHIMA: "HIROSHIMA",
		PCJP_YAMAGUCHI: "YAMAGUCHI",
		PCJP_TOKUSHIMA: "TOKUSHIMA",
		PCJP_KAGAWA:    "KAGAWA",
		PCJP_EHIME:     "EHIME",
		PCJP_KOCHI:     "KOCHI",
		PCJP_FUKUOKA:   "FUKUOKA",
		PCJP_SAGA:      "SAGA",
		PCJP_NAGASAKI:  "NAGASAKI",
		PCJP_KUMAMOTO:  "KUMAMOTO",
		PCJP_OITA:      "OITA",
		PCJP_MIYAZAKI:  "MIYAZAKI",
		PCJP_KAGOSHIMA: "KAGOSHIMA",
		PCJP_OKINAWA:   "OKINAWA",
	}
)

func prefJpCode(cd uint) PrefJpCode {
	if cd < uint(PCJP_MIN) || cd > uint(PCJP_MAX) {
		return PCJP_UNKNOWN
	}
	return PrefJpCode(cd)
}

func GetFromPrefCode(s string) PrefJpCode {
	s = strings.TrimPrefix(strings.ToUpper(strings.TrimSpace(s)), "JP-")
	n, err := strconv.ParseUint(s, 10, 8)
	if err != nil {
		return PCJP_UNKNOWN
	}
	return prefJpCode(uint(n))
}

func GetFromPrefName(s string) PrefJpCode {
	for k, v := range pcNames {
		if strings.EqualFold(v, s) {
			return k
		}
	}
	return PCJP_UNKNOWN
}

func (pc PrefJpCode) String() string {
	return fmt.Sprintf("%02d", uint(pc))
}

func (pc PrefJpCode) Name() string {
	for k, v := range pcNames {
		if pc == k {
			return v
		}
	}
	return ""
}

func (pc PrefJpCode) NameLower() string {
	return strings.ToLower(pc.Name())
}

func (pc PrefJpCode) Title() string {
	return cases.Title(language.Und, cases.NoLower).String(pc.NameLower())
}

//UnmarshalJSON method returns result of Unmarshal for json.Unmarshal().
func (pc *PrefJpCode) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*pc = GetFromPrefCode(s)
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
