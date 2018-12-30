package tests

import (
	"github.com/saichler/files/handlers"
	fservice "github.com/saichler/files/service"
	"github.com/saichler/habitat"
	"github.com/saichler/habitat/service"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

func startFileService(t *testing.T) *fservice.FileService{
	sm,err:=service.NewServiceManager()
	if err!=nil {
		t.Fail()
		logrus.Error(err)
		return nil
	}

	fs:=&fservice.FileService{}
	sm.AddService(fs)
	return fs
}


func TestFileService(t *testing.T){
	fs1:=startFileService(t)
	sm1:=fs1.GetManager()
	logrus.Info("Service 1 created ")

	fs2:=startFileService(t)
	sm2:=fs2.GetManager()

	time.Sleep(time.Second)

	sm1.CreateAndSend(fs1,sm2.Source(fs2),handlers.REQUEST_FILE_LIST,[]byte("some directory"))
	sm1.CreateAndSend(fs1,sm2.Source(fs2),handlers.REQUEST_FILE,[]byte("some file"))

	time.Sleep(time.Second*2)
}


func TestRemoteService(t *testing.T){

	fs:=startFileService(t)
	sm:=fs.GetManager()

	uplinkHID:=sm.Habitat().Uplink("some ip")

	time.Sleep(time.Second)

	dest:=habitat.NewSID(uplinkHID,fs.SID())

	time.Sleep(time.Second*2)
	sm.CreateAndSend(fs,dest,handlers.REQUEST_FILE_LIST,[]byte("some directory"))

	time.Sleep(time.Second*10)

	sm.CreateAndSend(fs,dest,handlers.REQUEST_FILE,[]byte("some large remote file"))

	fs.GetManager().WaitForShutdown()
}