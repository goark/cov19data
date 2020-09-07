package values

import (
	"strconv"
	"strings"
)

//CountryCode is country code by WHO.
type CountryCode int

const (
	CC_Other CountryCode = iota // Other
	CC_AD                       // Andorra
	CC_AE                       // United Arab Emirates
	CC_AF                       // Afghanistan
	CC_AG                       // Antigua and Barbuda
	CC_AI                       // Anguilla
	CC_AL                       // Albania
	CC_AM                       // Armenia
	CC_AO                       // Angola
	CC_AR                       // Argentina
	CC_AT                       // Austria
	CC_AU                       // Australia
	CC_AW                       // Aruba
	CC_AZ                       // Azerbaijan
	CC_BA                       // Bosnia and Herzegovina
	CC_BB                       // Barbados
	CC_BD                       // Bangladesh
	CC_BE                       // Belgium
	CC_BF                       // Burkina Faso
	CC_BG                       // Bulgaria
	CC_BH                       // Bahrain
	CC_BI                       // Burundi
	CC_BJ                       // Benin
	CC_BL                       // Saint Barthélemy
	CC_BM                       // Bermuda
	CC_BN                       // Brunei Darussalam
	CC_BO                       // Bolivia (Plurinational State of)
	CC_BQ                       // Bonaire, Sint Eustatius and Saba
	CC_BR                       // Brazil
	CC_BS                       // Bahamas
	CC_BT                       // Bhutan
	CC_BW                       // Botswana
	CC_BY                       // Belarus
	CC_BZ                       // Belize
	CC_CA                       // Canada
	CC_CD                       // Democratic Republic of the Congo
	CC_CF                       // Central African Republic
	CC_CG                       // Congo
	CC_CH                       // Switzerland
	CC_CI                       // Côte d’Ivoire
	CC_CL                       // Chile
	CC_CM                       // Cameroon
	CC_CN                       // China
	CC_CO                       // Colombia
	CC_CR                       // Costa Rica
	CC_CU                       // Cuba
	CC_CV                       // Cabo Verde
	CC_CW                       // Curaçao
	CC_CY                       // Cyprus
	CC_CZ                       // Czechia
	CC_DE                       // Germany
	CC_DJ                       // Djibouti
	CC_DK                       // Denmark
	CC_DM                       // Dominica
	CC_DO                       // Dominican Republic
	CC_DZ                       // Algeria
	CC_EC                       // Ecuador
	CC_EE                       // Estonia
	CC_EG                       // Egypt
	CC_ER                       // Eritrea
	CC_ES                       // Spain
	CC_ET                       // Ethiopia
	CC_FI                       // Finland
	CC_FJ                       // Fiji
	CC_FK                       // Falkland Islands (Malvinas)
	CC_FO                       // Faroe Islands
	CC_FR                       // France
	CC_GA                       // Gabon
	CC_GB                       // The United Kingdom
	CC_GD                       // Grenada
	CC_GE                       // Georgia
	CC_GF                       // French Guiana
	CC_GG                       // Guernsey
	CC_GH                       // Ghana
	CC_GI                       // Gibraltar
	CC_GL                       // Greenland
	CC_GM                       // Gambia
	CC_GN                       // Guinea
	CC_GP                       // Guadeloupe
	CC_GQ                       // Equatorial Guinea
	CC_GR                       // Greece
	CC_GT                       // Guatemala
	CC_GU                       // Guam
	CC_GW                       // Guinea-Bissau
	CC_GY                       // Guyana
	CC_HN                       // Honduras
	CC_HR                       // Croatia
	CC_HT                       // Haiti
	CC_HU                       // Hungary
	CC_ID                       // Indonesia
	CC_IE                       // Ireland
	CC_IL                       // Israel
	CC_IM                       // Isle of Man
	CC_IN                       // India
	CC_IQ                       // Iraq
	CC_IR                       // Iran (Islamic Republic of)
	CC_IS                       // Iceland
	CC_IT                       // Italy
	CC_JE                       // Jersey
	CC_JM                       // Jamaica
	CC_JO                       // Jordan
	CC_JP                       // Japan
	CC_KE                       // Kenya
	CC_KG                       // Kyrgyzstan
	CC_KH                       // Cambodia
	CC_KM                       // Comoros
	CC_KN                       // Saint Kitts and Nevis
	CC_KR                       // Republic of Korea
	CC_KW                       // Kuwait
	CC_KY                       // Cayman Islands
	CC_KZ                       // Kazakhstan
	CC_LA                       // Lao People's Democratic Republic
	CC_LB                       // Lebanon
	CC_LC                       // Saint Lucia
	CC_LI                       // Liechtenstein
	CC_LK                       // Sri Lanka
	CC_LR                       // Liberia
	CC_LS                       // Lesotho
	CC_LT                       // Lithuania
	CC_LU                       // Luxembourg
	CC_LV                       // Latvia
	CC_LY                       // Libya
	CC_MA                       // Morocco
	CC_MC                       // Monaco
	CC_MD                       // Republic of Moldova
	CC_ME                       // Montenegro
	CC_MF                       // Saint Martin
	CC_MG                       // Madagascar
	CC_MK                       // North Macedonia
	CC_ML                       // Mali
	CC_MM                       // Myanmar
	CC_MN                       // Mongolia
	CC_MP                       // Northern Mariana Islands (Commonwealth of the)
	CC_MQ                       // Martinique
	CC_MR                       // Mauritania
	CC_MS                       // Montserrat
	CC_MT                       // Malta
	CC_MU                       // Mauritius
	CC_MV                       // Maldives
	CC_MW                       // Malawi
	CC_MX                       // Mexico
	CC_MY                       // Malaysia
	CC_MZ                       // Mozambique
	CC_NA                       // Namibia
	CC_NC                       // New Caledonia
	CC_NE                       // Niger
	CC_NG                       // Nigeria
	CC_NI                       // Nicaragua
	CC_NL                       // Netherlands
	CC_NO                       // Norway
	CC_NP                       // Nepal
	CC_NZ                       // New Zealand
	CC_OM                       // Oman
	CC_PA                       // Panama
	CC_PE                       // Peru
	CC_PF                       // French Polynesia
	CC_PG                       // Papua New Guinea
	CC_PH                       // Philippines
	CC_PK                       // Pakistan
	CC_PL                       // Poland
	CC_PM                       // Saint Pierre and Miquelon
	CC_PR                       // Puerto Rico
	CC_PS                       // occupied Palestinian territory, including east Jerusalem
	CC_PT                       // Portugal
	CC_PY                       // Paraguay
	CC_QA                       // Qatar
	CC_RE                       // Réunion
	CC_RO                       // Romania
	CC_RS                       // Serbia
	CC_RU                       // Russian Federation
	CC_RW                       // Rwanda
	CC_SA                       // Saudi Arabia
	CC_SC                       // Seychelles
	CC_SD                       // Sudan
	CC_SE                       // Sweden
	CC_SG                       // Singapore
	CC_SI                       // Slovenia
	CC_SK                       // Slovakia
	CC_SL                       // Sierra Leone
	CC_SM                       // San Marino
	CC_SN                       // Senegal
	CC_SO                       // Somalia
	CC_SR                       // Suriname
	CC_SS                       // South Sudan
	CC_ST                       // Sao Tome and Principe
	CC_SV                       // El Salvador
	CC_SX                       // Sint Maarten
	CC_SY                       // Syrian Arab Republic
	CC_SZ                       // Eswatini
	CC_TC                       // Turks and Caicos Islands
	CC_TD                       // Chad
	CC_TG                       // Togo
	CC_TH                       // Thailand
	CC_TJ                       // Tajikistan
	CC_TL                       // Timor-Leste
	CC_TN                       // Tunisia
	CC_TR                       // Turkey
	CC_TT                       // Trinidad and Tobago
	CC_TZ                       // United Republic of Tanzania
	CC_UA                       // Ukraine
	CC_UG                       // Uganda
	CC_US                       // United States of America
	CC_UY                       // Uruguay
	CC_UZ                       // Uzbekistan
	CC_VA                       // Holy See
	CC_VC                       // Saint Vincent and the Grenadines
	CC_VE                       // Venezuela (Bolivarian Republic of)
	CC_VG                       // British Virgin Islands
	CC_VI                       // United States Virgin Islands
	CC_VN                       // Viet Nam
	CC_XK                       // Kosovo
	CC_YE                       // Yemen
	CC_YT                       // Mayotte
	CC_ZA                       // South Africa
	CC_ZM                       // Zambia
	CC_ZW                       // Zimbabwe
)

var countryCodeMap = map[CountryCode]string{
	CC_AD: "AD", // Andorra
	CC_AE: "AE", // United Arab Emirates
	CC_AF: "AF", // Afghanistan
	CC_AG: "AG", // Antigua and Barbuda
	CC_AI: "AI", // Anguilla
	CC_AL: "AL", // Albania
	CC_AM: "AM", // Armenia
	CC_AO: "AO", // Angola
	CC_AR: "AR", // Argentina
	CC_AT: "AT", // Austria
	CC_AU: "AU", // Australia
	CC_AW: "AW", // Aruba
	CC_AZ: "AZ", // Azerbaijan
	CC_BA: "BA", // Bosnia and Herzegovina
	CC_BB: "BB", // Barbados
	CC_BD: "BD", // Bangladesh
	CC_BE: "BE", // Belgium
	CC_BF: "BF", // Burkina Faso
	CC_BG: "BG", // Bulgaria
	CC_BH: "BH", // Bahrain
	CC_BI: "BI", // Burundi
	CC_BJ: "BJ", // Benin
	CC_BL: "BL", // Saint Barthélemy
	CC_BM: "BM", // Bermuda
	CC_BN: "BN", // Brunei Darussalam
	CC_BO: "BO", // Bolivia (Plurinational State of)
	CC_BQ: "BQ", // Bonaire, Sint Eustatius and Saba
	CC_BR: "BR", // Brazil
	CC_BS: "BS", // Bahamas
	CC_BT: "BT", // Bhutan
	CC_BW: "BW", // Botswana
	CC_BY: "BY", // Belarus
	CC_BZ: "BZ", // Belize
	CC_CA: "CA", // Canada
	CC_CD: "CD", // Democratic Republic of the Congo
	CC_CF: "CF", // Central African Republic
	CC_CG: "CG", // Congo
	CC_CH: "CH", // Switzerland
	CC_CI: "CI", // Côte d’Ivoire
	CC_CL: "CL", // Chile
	CC_CM: "CM", // Cameroon
	CC_CN: "CN", // China
	CC_CO: "CO", // Colombia
	CC_CR: "CR", // Costa Rica
	CC_CU: "CU", // Cuba
	CC_CV: "CV", // Cabo Verde
	CC_CW: "CW", // Curaçao
	CC_CY: "CY", // Cyprus
	CC_CZ: "CZ", // Czechia
	CC_DE: "DE", // Germany
	CC_DJ: "DJ", // Djibouti
	CC_DK: "DK", // Denmark
	CC_DM: "DM", // Dominica
	CC_DO: "DO", // Dominican Republic
	CC_DZ: "DZ", // Algeria
	CC_EC: "EC", // Ecuador
	CC_EE: "EE", // Estonia
	CC_EG: "EG", // Egypt
	CC_ER: "ER", // Eritrea
	CC_ES: "ES", // Spain
	CC_ET: "ET", // Ethiopia
	CC_FI: "FI", // Finland
	CC_FJ: "FJ", // Fiji
	CC_FK: "FK", // Falkland Islands (Malvinas)
	CC_FO: "FO", // Faroe Islands
	CC_FR: "FR", // France
	CC_GA: "GA", // Gabon
	CC_GB: "GB", // The United Kingdom
	CC_GD: "GD", // Grenada
	CC_GE: "GE", // Georgia
	CC_GF: "GF", // French Guiana
	CC_GG: "GG", // Guernsey
	CC_GH: "GH", // Ghana
	CC_GI: "GI", // Gibraltar
	CC_GL: "GL", // Greenland
	CC_GM: "GM", // Gambia
	CC_GN: "GN", // Guinea
	CC_GP: "GP", // Guadeloupe
	CC_GQ: "GQ", // Equatorial Guinea
	CC_GR: "GR", // Greece
	CC_GT: "GT", // Guatemala
	CC_GU: "GU", // Guam
	CC_GW: "GW", // Guinea-Bissau
	CC_GY: "GY", // Guyana
	CC_HN: "HN", // Honduras
	CC_HR: "HR", // Croatia
	CC_HT: "HT", // Haiti
	CC_HU: "HU", // Hungary
	CC_ID: "ID", // Indonesia
	CC_IE: "IE", // Ireland
	CC_IL: "IL", // Israel
	CC_IM: "IM", // Isle of Man
	CC_IN: "IN", // India
	CC_IQ: "IQ", // Iraq
	CC_IR: "IR", // Iran (Islamic Republic of)
	CC_IS: "IS", // Iceland
	CC_IT: "IT", // Italy
	CC_JE: "JE", // Jersey
	CC_JM: "JM", // Jamaica
	CC_JO: "JO", // Jordan
	CC_JP: "JP", // Japan
	CC_KE: "KE", // Kenya
	CC_KG: "KG", // Kyrgyzstan
	CC_KH: "KH", // Cambodia
	CC_KM: "KM", // Comoros
	CC_KN: "KN", // Saint Kitts and Nevis
	CC_KR: "KR", // Republic of Korea
	CC_KW: "KW", // Kuwait
	CC_KY: "KY", // Cayman Islands
	CC_KZ: "KZ", // Kazakhstan
	CC_LA: "LA", // Lao People's Democratic Republic
	CC_LB: "LB", // Lebanon
	CC_LC: "LC", // Saint Lucia
	CC_LI: "LI", // Liechtenstein
	CC_LK: "LK", // Sri Lanka
	CC_LR: "LR", // Liberia
	CC_LS: "LS", // Lesotho
	CC_LT: "LT", // Lithuania
	CC_LU: "LU", // Luxembourg
	CC_LV: "LV", // Latvia
	CC_LY: "LY", // Libya
	CC_MA: "MA", // Morocco
	CC_MC: "MC", // Monaco
	CC_MD: "MD", // Republic of Moldova
	CC_ME: "ME", // Montenegro
	CC_MF: "MF", // Saint Martin
	CC_MG: "MG", // Madagascar
	CC_MK: "MK", // North Macedonia
	CC_ML: "ML", // Mali
	CC_MM: "MM", // Myanmar
	CC_MN: "MN", // Mongolia
	CC_MP: "MP", // Northern Mariana Islands (Commonwealth of the)
	CC_MQ: "MQ", // Martinique
	CC_MR: "MR", // Mauritania
	CC_MS: "MS", // Montserrat
	CC_MT: "MT", // Malta
	CC_MU: "MU", // Mauritius
	CC_MV: "MV", // Maldives
	CC_MW: "MW", // Malawi
	CC_MX: "MX", // Mexico
	CC_MY: "MY", // Malaysia
	CC_MZ: "MZ", // Mozambique
	CC_NA: "NA", // Namibia
	CC_NC: "NC", // New Caledonia
	CC_NE: "NE", // Niger
	CC_NG: "NG", // Nigeria
	CC_NI: "NI", // Nicaragua
	CC_NL: "NL", // Netherlands
	CC_NO: "NO", // Norway
	CC_NP: "NP", // Nepal
	CC_NZ: "NZ", // New Zealand
	CC_OM: "OM", // Oman
	CC_PA: "PA", // Panama
	CC_PE: "PE", // Peru
	CC_PF: "PF", // French Polynesia
	CC_PG: "PG", // Papua New Guinea
	CC_PH: "PH", // Philippines
	CC_PK: "PK", // Pakistan
	CC_PL: "PL", // Poland
	CC_PM: "PM", // Saint Pierre and Miquelon
	CC_PR: "PR", // Puerto Rico
	CC_PS: "PS", // occupied Palestinian territory, including east Jerusalem
	CC_PT: "PT", // Portugal
	CC_PY: "PY", // Paraguay
	CC_QA: "QA", // Qatar
	CC_RE: "RE", // Réunion
	CC_RO: "RO", // Romania
	CC_RS: "RS", // Serbia
	CC_RU: "RU", // Russian Federation
	CC_RW: "RW", // Rwanda
	CC_SA: "SA", // Saudi Arabia
	CC_SC: "SC", // Seychelles
	CC_SD: "SD", // Sudan
	CC_SE: "SE", // Sweden
	CC_SG: "SG", // Singapore
	CC_SI: "SI", // Slovenia
	CC_SK: "SK", // Slovakia
	CC_SL: "SL", // Sierra Leone
	CC_SM: "SM", // San Marino
	CC_SN: "SN", // Senegal
	CC_SO: "SO", // Somalia
	CC_SR: "SR", // Suriname
	CC_SS: "SS", // South Sudan
	CC_ST: "ST", // Sao Tome and Principe
	CC_SV: "SV", // El Salvador
	CC_SX: "SX", // Sint Maarten
	CC_SY: "SY", // Syrian Arab Republic
	CC_SZ: "SZ", // Eswatini
	CC_TC: "TC", // Turks and Caicos Islands
	CC_TD: "TD", // Chad
	CC_TG: "TG", // Togo
	CC_TH: "TH", // Thailand
	CC_TJ: "TJ", // Tajikistan
	CC_TL: "TL", // Timor-Leste
	CC_TN: "TN", // Tunisia
	CC_TR: "TR", // Turkey
	CC_TT: "TT", // Trinidad and Tobago
	CC_TZ: "TZ", // United Republic of Tanzania
	CC_UA: "UA", // Ukraine
	CC_UG: "UG", // Uganda
	CC_US: "US", // United States of America
	CC_UY: "UY", // Uruguay
	CC_UZ: "UZ", // Uzbekistan
	CC_VA: "VA", // Holy See
	CC_VC: "VC", // Saint Vincent and the Grenadines
	CC_VE: "VE", // Venezuela (Bolivarian Republic of)
	CC_VG: "VG", // British Virgin Islands
	CC_VI: "VI", // United States Virgin Islands
	CC_VN: "VN", // Viet Nam
	CC_XK: "XK", // Kosovo[1]
	CC_YE: "YE", // Yemen
	CC_YT: "YT", // Mayotte
	CC_ZA: "ZA", // South Africa
	CC_ZM: "ZM", // Zambia
	CC_ZW: "ZW", // Zimbabwe
}

var countryNameMap = map[CountryCode]string{
	CC_AD: "Andorra",
	CC_AE: "United Arab Emirates",
	CC_AF: "Afghanistan",
	CC_AG: "Antigua and Barbuda",
	CC_AI: "Anguilla",
	CC_AL: "Albania",
	CC_AM: "Armenia",
	CC_AO: "Angola",
	CC_AR: "Argentina",
	CC_AT: "Austria",
	CC_AU: "Australia",
	CC_AW: "Aruba",
	CC_AZ: "Azerbaijan",
	CC_BA: "Bosnia and Herzegovina",
	CC_BB: "Barbados",
	CC_BD: "Bangladesh",
	CC_BE: "Belgium",
	CC_BF: "Burkina Faso",
	CC_BG: "Bulgaria",
	CC_BH: "Bahrain",
	CC_BI: "Burundi",
	CC_BJ: "Benin",
	CC_BL: "Saint Barthélemy",
	CC_BM: "Bermuda",
	CC_BN: "Brunei Darussalam",
	CC_BO: "Bolivia (Plurinational State of)",
	CC_BQ: "Bonaire, Sint Eustatius and Saba",
	CC_BR: "Brazil",
	CC_BS: "Bahamas",
	CC_BT: "Bhutan",
	CC_BW: "Botswana",
	CC_BY: "Belarus",
	CC_BZ: "Belize",
	CC_CA: "Canada",
	CC_CD: "Democratic Republic of the Congo",
	CC_CF: "Central African Republic",
	CC_CG: "Congo",
	CC_CH: "Switzerland",
	CC_CI: "Côte d’Ivoire",
	CC_CL: "Chile",
	CC_CM: "Cameroon",
	CC_CN: "China",
	CC_CO: "Colombia",
	CC_CR: "Costa Rica",
	CC_CU: "Cuba",
	CC_CV: "Cabo Verde",
	CC_CW: "Curaçao",
	CC_CY: "Cyprus",
	CC_CZ: "Czechia",
	CC_DE: "Germany",
	CC_DJ: "Djibouti",
	CC_DK: "Denmark",
	CC_DM: "Dominica",
	CC_DO: "Dominican Republic",
	CC_DZ: "Algeria",
	CC_EC: "Ecuador",
	CC_EE: "Estonia",
	CC_EG: "Egypt",
	CC_ER: "Eritrea",
	CC_ES: "Spain",
	CC_ET: "Ethiopia",
	CC_FI: "Finland",
	CC_FJ: "Fiji",
	CC_FK: "Falkland Islands (Malvinas)",
	CC_FO: "Faroe Islands",
	CC_FR: "France",
	CC_GA: "Gabon",
	CC_GB: "The United Kingdom",
	CC_GD: "Grenada",
	CC_GE: "Georgia",
	CC_GF: "French Guiana",
	CC_GG: "Guernsey",
	CC_GH: "Ghana",
	CC_GI: "Gibraltar",
	CC_GL: "Greenland",
	CC_GM: "Gambia",
	CC_GN: "Guinea",
	CC_GP: "Guadeloupe",
	CC_GQ: "Equatorial Guinea",
	CC_GR: "Greece",
	CC_GT: "Guatemala",
	CC_GU: "Guam",
	CC_GW: "Guinea-Bissau",
	CC_GY: "Guyana",
	CC_HN: "Honduras",
	CC_HR: "Croatia",
	CC_HT: "Haiti",
	CC_HU: "Hungary",
	CC_ID: "Indonesia",
	CC_IE: "Ireland",
	CC_IL: "Israel",
	CC_IM: "Isle of Man",
	CC_IN: "India",
	CC_IQ: "Iraq",
	CC_IR: "Iran (Islamic Republic of)",
	CC_IS: "Iceland",
	CC_IT: "Italy",
	CC_JE: "Jersey",
	CC_JM: "Jamaica",
	CC_JO: "Jordan",
	CC_JP: "Japan",
	CC_KE: "Kenya",
	CC_KG: "Kyrgyzstan",
	CC_KH: "Cambodia",
	CC_KM: "Comoros",
	CC_KN: "Saint Kitts and Nevis",
	CC_KR: "Republic of Korea",
	CC_KW: "Kuwait",
	CC_KY: "Cayman Islands",
	CC_KZ: "Kazakhstan",
	CC_LA: "Lao People's Democratic Republic",
	CC_LB: "Lebanon",
	CC_LC: "Saint Lucia",
	CC_LI: "Liechtenstein",
	CC_LK: "Sri Lanka",
	CC_LR: "Liberia",
	CC_LS: "Lesotho",
	CC_LT: "Lithuania",
	CC_LU: "Luxembourg",
	CC_LV: "Latvia",
	CC_LY: "Libya",
	CC_MA: "Morocco",
	CC_MC: "Monaco",
	CC_MD: "Republic of Moldova",
	CC_ME: "Montenegro",
	CC_MF: "Saint Martin",
	CC_MG: "Madagascar",
	CC_MK: "North Macedonia",
	CC_ML: "Mali",
	CC_MM: "Myanmar",
	CC_MN: "Mongolia",
	CC_MP: "Northern Mariana Islands (Commonwealth of the)",
	CC_MQ: "Martinique",
	CC_MR: "Mauritania",
	CC_MS: "Montserrat",
	CC_MT: "Malta",
	CC_MU: "Mauritius",
	CC_MV: "Maldives",
	CC_MW: "Malawi",
	CC_MX: "Mexico",
	CC_MY: "Malaysia",
	CC_MZ: "Mozambique",
	CC_NA: "Namibia",
	CC_NC: "New Caledonia",
	CC_NE: "Niger",
	CC_NG: "Nigeria",
	CC_NI: "Nicaragua",
	CC_NL: "Netherlands",
	CC_NO: "Norway",
	CC_NP: "Nepal",
	CC_NZ: "New Zealand",
	CC_OM: "Oman",
	CC_PA: "Panama",
	CC_PE: "Peru",
	CC_PF: "French Polynesia",
	CC_PG: "Papua New Guinea",
	CC_PH: "Philippines",
	CC_PK: "Pakistan",
	CC_PL: "Poland",
	CC_PM: "Saint Pierre and Miquelon",
	CC_PR: "Puerto Rico",
	CC_PS: "occupied Palestinian territory, including east Jerusalem",
	CC_PT: "Portugal",
	CC_PY: "Paraguay",
	CC_QA: "Qatar",
	CC_RE: "Réunion",
	CC_RO: "Romania",
	CC_RS: "Serbia",
	CC_RU: "Russian Federation",
	CC_RW: "Rwanda",
	CC_SA: "Saudi Arabia",
	CC_SC: "Seychelles",
	CC_SD: "Sudan",
	CC_SE: "Sweden",
	CC_SG: "Singapore",
	CC_SI: "Slovenia",
	CC_SK: "Slovakia",
	CC_SL: "Sierra Leone",
	CC_SM: "San Marino",
	CC_SN: "Senegal",
	CC_SO: "Somalia",
	CC_SR: "Suriname",
	CC_SS: "South Sudan",
	CC_ST: "Sao Tome and Principe",
	CC_SV: "El Salvador",
	CC_SX: "Sint Maarten",
	CC_SY: "Syrian Arab Republic",
	CC_SZ: "Eswatini",
	CC_TC: "Turks and Caicos Islands",
	CC_TD: "Chad",
	CC_TG: "Togo",
	CC_TH: "Thailand",
	CC_TJ: "Tajikistan",
	CC_TL: "Timor-Leste",
	CC_TN: "Tunisia",
	CC_TR: "Turkey",
	CC_TT: "Trinidad and Tobago",
	CC_TZ: "United Republic of Tanzania",
	CC_UA: "Ukraine",
	CC_UG: "Uganda",
	CC_US: "United States of America",
	CC_UY: "Uruguay",
	CC_UZ: "Uzbekistan",
	CC_VA: "Holy See",
	CC_VC: "Saint Vincent and the Grenadines",
	CC_VE: "Venezuela (Bolivarian Republic of)",
	CC_VG: "British Virgin Islands",
	CC_VI: "United States Virgin Islands",
	CC_VN: "Viet Nam",
	CC_XK: "Kosovo",
	CC_YE: "Yemen",
	CC_YT: "Mayotte",
	CC_ZA: "South Africa",
	CC_ZM: "Zambia",
	CC_ZW: "Zimbabwe",
}

//GetCountryCode function returns CountryCode instrance from string.
func GetCountryCode(s string) CountryCode {
	for k, v := range countryCodeMap {
		if strings.EqualFold(v, s) {
			return k
		}
	}
	return CC_Other
}

func (cc CountryCode) String() string {
	if s, ok := countryCodeMap[cc]; ok {
		return s
	}
	return "Other"
}

//Name method returns country name.
func (cc CountryCode) Name() string {
	if s, ok := countryNameMap[cc]; ok {
		return s
	}
	return "Other"
}

//UnmarshalJSON method returns result of Unmarshal for json.Unmarshal().
func (cc *CountryCode) UnmarshalJSON(b []byte) error {
	s, err := strconv.Unquote(string(b))
	if err != nil {
		s = string(b)
	}
	*cc = GetCountryCode(s)
	return nil
}

//MarshalJSON method returns string for json.Marshal().
func (cc *CountryCode) MarshalJSON() ([]byte, error) {
	if cc == nil {
		return []byte(`""`), nil
	}
	return []byte(strconv.Quote(cc.String())), nil
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
