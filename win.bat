cd src

set GOOS=linux
set GOARCH=arm
set GOARM=7
go mod tidy
go build -o wifi_plugins