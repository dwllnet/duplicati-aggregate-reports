export GOPATH=~/golang:~/go:~/git/duplicati-aggregate-reports/src
export GOROOT=/usr/local/go

env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v dar.go 
