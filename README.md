Repo for easy creating geoip API based on AWS lambda and Maxmind's GeoLite2 database.

Inspired by [tmaiaroto/go-lambda-geoip](github.com/tmaiaroto/go-lambda-geoip).

### Prerequisites
 - An [AWS account](https://console.aws.amazon.com/iam/home#/home) (and configured credentials for CLI)
 - [golang](https://golang.org/)
 - [Maxmind's GeoLite2 City/Country database.](http://dev.maxmind.com/geoip/geoip2/geolite2/)
 - [Serverless Framework](https://serverless.com/)
 - [go-bindata](https://github.com/jteeuwen/go-bindata)

## Instructions
### go language
[Download binary](https://golang.org/dl/)

[Install Guide](https://golang.org/doc/install)

### go-bindata
```go get -u github.com/jteeuwen/go-bindata/go-bindata```

### if all set
Assuming you have AWS credentials configured and you've got the `serverless` and `go-bindata` ready to use.

You can run the following to retrieve the geoip data set, build, and deploy:

Download fresh Maxmind's database:
```
make download
```

Deploy stack:
```
make deploy
```

It should build and deploy the Lambda. It may take a little bit due to the size of the database file.


## TODO list
- build & deploy via [docker](https://www.docker.com/)
- return errors as valid JSON object
- change `geocoder` function to work with [geocoder-izi-lookup](https://github.com/IzikAJ/geocoder-izi-lookup) propely
- split code to modules
- write tests
- refactor code
