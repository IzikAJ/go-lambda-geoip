package main

import (
	"github.com/oschwald/geoip2-golang"
)

const version = "0.1.1"

// IziLookup - data struct got geoip-izi-lookup
type IziLookup struct {
	IP                string `json:"ip"`
	ContinentCode     string `json:"continent_code"`
	ContinentName     string `json:"continent_name"`
	CountryCode       string `json:"country_code"`
	CountryName       string `json:"country_name"`
	LocationLatitude  string `json:"location_latitude"`
	LocationLongitude string `json:"location_longitude"`
	LocationTimeZone  string `json:"location_time_zone"`
	PostalCode        string `json:"postal_code"`
	Version           string `json:"version"`
	InEu              bool   `json:"in_eu"`
}

func fromGeoipCountry(ip string, data *geoip2.Country) (res IziLookup) {
	return IziLookup{
		IP:      ip,
		Version: version,

		ContinentCode: data.Continent.Code,
		ContinentName: data.Continent.Names["en"],

		CountryCode: data.Country.IsoCode,
		CountryName: data.Country.Names["en"],

		InEu: isInEu(data),
	}
}
