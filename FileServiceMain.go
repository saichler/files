package main

import . "github.com/saichler/habitat/service"
import . "github.com/saichler/files/service"

var ServiceInstance Service = &FileService{}

func main(){
	fs:=&FileService{}
	svm,_:=NewServiceManager()
	svm.AddService(fs)
	svm.WaitForShutdown()
}
