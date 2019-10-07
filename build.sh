cd ../habitat/golang/MyHabitat
go build
cp MyHabitat ../../../files/.
cd ../../../files/
go build -buildmode=plugin -o FileService.so FileServiceMain.go

#env GOOS=freebsd GOARCH=amd64 go build
