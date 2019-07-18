package main

import (
	"github.com/oschwald/geoip2-golang"
)

var in_eu = []string{
	"AT",
	"AX",
	"BE",
	"BG",
	"CY",
	"CZ",
	"DE",
	"DK",
	"EE",
	"ES",
	"FI",
	"FO",
	"FR",
	"GB",
	"GF",
	"GI",
	"GR",
	"HR",
	"HU",
	"IE",
	"IM",
	"IT",
	"LT",
	"LU",
	"LV",
	"MT",
	"NL",
	"PL",
	"PT",
	"RO",
	"SE",
	"SI",
	"SK",
}

func isInEu(country *geoip2.Country) bool {
	code := country.Country.IsoCode
	for _, a := range in_eu {
		if a == code {
			return true
		}
	}
	return false
}
