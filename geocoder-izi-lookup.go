package main

import (
	"context"
	"encoding/json"
	"log"
	"net"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

// Response - jsust a shortalnd
type Response = events.APIGatewayProxyResponse

var jsonHeaders = map[string]string{
	"Content-Type": "application/json",
}

func asJSON(data interface{}) (resp events.APIGatewayProxyResponse, err error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return Response{StatusCode: 422, Body: err.Error()}, nil
	}

	resp = Response{
		Body:       string(raw),
		Headers:    jsonHeaders,
		StatusCode: 200,
	}
	return
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (Response, error) {
	remote := request.QueryStringParameters["remote"]

	parsedIP := net.ParseIP(remote)

	country, err := db.Country(parsedIP)

	if err != nil {
		return Response{StatusCode: 422, Body: err.Error()}, nil
	}

	return asJSON(country)
}

func main() {
	var err error
	var data []byte

	ctx, seg := xray.BeginSegment(context.Background(), "GeoipInit")

	err = xray.Capture(ctx, "LoadBytes", func(ctx1 context.Context) error {
		data, err = Asset("GeoLite2-Country.mmdb")
		return err
	})

	if err == nil {
		xray.Capture(ctx, "DatabaseFromBytes", func(ctx1 context.Context) error {
			db, _ = geoip2.FromBytes(data)
			return nil
		})
		defer db.Close()

		// Close the segment and subsegment
		seg.Close(nil)

		lambda.Start(handleRequest)
	} else {
		log.Println("Could not load GeoLite2-Country.mmdb. Is it included in the binary?", err)
	}
}
