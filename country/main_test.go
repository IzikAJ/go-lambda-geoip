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
	country, err := db.Country(parsedIP)
	t.Logf("country: %v\n err: %v\n", country, err)

	// t.Errorf("GOT: %v\n", country)
}
