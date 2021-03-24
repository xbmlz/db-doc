@echo off
echo Build Windows...
go build -ldflags="-w -s" -o build\DbDoc-1.1.0-win\DbDoc.exe main.go
echo Build Windows Successfully!

echo Build Linux...
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-w -s" -o build\DbDoc-1.1.0-linux\DbDoc main.go
echo Build Linux Successfully!


echo Build Mac...
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-w -s" -o build\DbDoc-1.1.0-mac\DbDoc main.go
echo Build Mac Successfully!

pause