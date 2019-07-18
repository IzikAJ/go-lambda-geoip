package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/oschwald/geoip2-golang"
)

const dbName = "GeoLite2-City.mmdb"

var db *geoip2.Reader

func main() {
	var err error
	var data []byte

	ctx, seg := xray.BeginSegment(context.Background(), "GeoipInit")

	err = xray.Capture(ctx, "LoadBytes", func(ctx1 context.Context) error {
		data, err = Asset(dbName)
		return err
	})

	if err == nil {
		xray.Capture(ctx, "DatabaseFromBytes", func(ctx1 context.Context) error {
			db, _ = geoip2.FromBytes(data)
			return nil
		})
		defer db.Close()
		seg.Close(nil)

		lambda.Start(handleRequest)
	} else {
		log.Println("Could not load "+dbName+". Is it included in the binary?", err)
	}
}
