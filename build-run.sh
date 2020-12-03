export GOPATH=~/golang:~/go:~/git/duplicati-aggregate-reports/src
export GOROOT=/usr/local/go
go build -v -a -o ./bin/dar dar.go
./bin/dar
