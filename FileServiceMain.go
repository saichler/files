package main

import (
	. "github.com/saichler/files/service"
	. "github.com/saichler/habitat/service"
)

var ServiceInstance Service = &FileService{}

func main(){
	svm,_:=NewServiceManager()

	/*
	uplinkHID:=svm.Habitat().Uplink("192.168.86.29")
	fmt.Println(uplinkHID.String())
*/
	fs:=&FileService{}
	svm.AddService(fs)

	svm.WaitForShutdown()
}
