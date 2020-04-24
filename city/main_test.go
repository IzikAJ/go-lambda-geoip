package main

import (
	"net"
	"testing"

	"github.com/oschwald/geoip2-golang"
)

func TestData(t *testing.T) {
	data, _ := Asset(dbName)
	db, _ := geoip2.FromBytes(data)
	defer db.Close()

	parsedIP := net.ParseIP("1.1.1.1")
	t.Logf("parsedIP: %v\n", parsedIP)
	city, err := db.City(parsedIP)
	t.Logf("city: %v\n err: %v\n", city, err)
	t.Logf("TimeZone: %v\n", city.Location.TimeZone)
	t.Logf("lat: %v, lon: %v\n", city.Location.Latitude, city.Location.Longitude)

	// t.Errorf("GOT: %v\n", city)
}
