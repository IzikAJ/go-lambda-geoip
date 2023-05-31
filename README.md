Repo for easy creating geoip API based on AWS lambda and Maxmind's GeoLite2 database.

Inspired by [tmaiaroto/go-lambda-geoip](github.com/tmaiaroto/go-lambda-geoip).

### Prerequisites
 - An [AWS account](https://console.aws.amazon.com/iam/home#/home) (and configured credentials for CLI)
 - [golang](https://golang.org/)
 - [Maxmind's GeoLite2 City/Country database.](http://dev.maxmind.com/geoip/geoip2/geolite2/)
 - [Serverless Framework](https://serverless.com/)
 - [go-bindata](https://github.com/jteeuwen/go-bindata)

### Useful links
 - [AWS FreeTier usage stats](https://console.aws.amazon.com/billing/home?#/freetier)

## Instructions
### Deploy with Docker

Build image
```
docker build . -t izikaj/geoip --build-arg MAXMIND_LICENSE_KEY=$MAXMIND_LICENSE_KEY
```

Deploy it to AWS
```
docker run -e AWS_ACCESS_KEY_ID="XXXXXXXXX" -e AWS_SECRET_ACCESS_KEY="YYYYYYYYYYYYY" -it izikaj/geoip
```

To run shell on image:
```
docker run -e AWS_ACCESS_KEY_ID="XXXXXXXXX" -e AWS_SECRET_ACCESS_KEY="YYYYYYYYYYYYY" -it izikaj/geoip sh
```

### go language
[Download binary](https://golang.org/dl/)

[Install Guide](https://golang.org/doc/install)

### go-bindata
```go install github.com/jteeuwen/go-bindata/go-bindata```

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

## Features:
- build & deploy via [docker](https://www.docker.com/)
- `geocoder` function should work with [geocoder-izi-lookup](https://github.com/IzikAJ/geocoder-izi-lookup) propely

## TODO:
- write tests
- return errors more detailed errors
- refactor code

```
arch -x86_64 zsh
```
