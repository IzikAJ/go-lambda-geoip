package main

import (
	"context"
	"net"

	"github.com/aws/aws-lambda-go/events"
	"github.com/izikaj/go-lambda-geoip/shared"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (shared.Response, error) {
	remote := request.QueryStringParameters["remote"]

	parsedIP := net.ParseIP(remote)
	country, err := db.City(parsedIP)

	if err != nil {
		return shared.AsError(remote, err)
	}

	return shared.AsData(remote, country)
}
