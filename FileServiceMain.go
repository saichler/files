package main

import (
	. "github.com/saichler/files/service"
	. "github.com/saichler/habitat/service"
	. "github.com/saichler/utils/golang"
	"os"
	"strings"
)

var ServiceInstance Service = &FileService{}

func main() {
	s,err:=NewServiceManager()
	if err!=nil {
		Error("Failed to load habitat",err)
		return
	}

	s.AddService(&FileService{})

	args := os.Args[1:]

	for _,arg:=range args {
		if strings.Contains(arg,".so") {
			s.InstallService(arg)
		} else {
			s.Habitat().Uplink(arg)
		}
	}

	s.WaitForShutdown()
}
