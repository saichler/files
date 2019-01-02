package main

import (
	. "github.com/saichler/files/service"
	. "github.com/saichler/habitat/service"
)

var ServiceInstance Service = &FileService{}

func main(){
	svm,_:=NewServiceManager()

	fs:=&FileService{}
	svm.AddService(fs)

	svm.WaitForShutdown()
}
