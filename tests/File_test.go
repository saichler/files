package tests

import (
	"github.com/saichler/files/handlers"
	fservice "github.com/saichler/files/service"
	"github.com/saichler/habitat"
	"github.com/saichler/habitat/service"
	"github.com/sirupsen/logrus"
	"time"
	"testing"
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
	sm1:=fs1.ServiceManager()
	logrus.Info("Service 1 created ")

	fs2:=startFileService(t)

	time.Sleep(time.Second)

	sm1.CreateAndSend(fs1,fs2.ServiceID(),handlers.REQUEST_FILE_LIST,[]byte("some directory"))
	sm1.CreateAndSend(fs1,fs2.ServiceID(),handlers.REQUEST_FILE,[]byte("some file"))

	time.Sleep(time.Second*2)
}


func TestRemoteService(t *testing.T){
	fs:=startFileService(t)
	sm:=fs.ServiceManager()

	uplinkHID:=sm.Habitat().Uplink("some ip")

	time.Sleep(time.Second)

	dest:=habitat.NewServiceID(uplinkHID,0,fs.ServiceID().Topic())

	time.Sleep(time.Second*2)
	sm.CreateAndSend(fs,dest,handlers.REQUEST_FILE_LIST,[]byte("some directory"))

	time.Sleep(time.Second*10)

	sm.CreateAndSend(fs,dest,handlers.REQUEST_FILE,[]byte("some large remote file"))

	fs.ServiceManager().WaitForShutdown()
}

func TestRemoteService2(t *testing.T) {
	fs:=startFileService(t)

	logrus.Info("Waiting")
	time.Sleep(time.Second*10)
	logrus.Info("Sending Message")

	sm:=fs.ServiceManager()
	dest:=habitat.NewServiceID(habitat.NewHID("192.168.86.29",52001),0,fs.ServiceID().Topic())

	sm.CreateAndSend(fs,dest,handlers.REQUEST_FILE_LIST,[]byte("/mnt/Vol1/Media/complete"))

	logrus.Info("Waiting")
	time.Sleep(time.Second*10)
	logrus.Info("Sending Message")

	sm.CreateAndSend(fs,dest,handlers.REQUEST_FILE,[]byte("/mnt/Vol1/Media/complete/VID_20181003_184455AA.MP4"))

	sm.WaitForShutdown()
}

func TestWithAdjacent(t *testing.T) {
	fs:=startFileService(t)

	logrus.Info("Waiting")
	time.Sleep(time.Second*10)
	logrus.Info("Sending Message")

	sm:=fs.ServiceManager()
	adjacents:=make([]*habitat.ServiceID,0)
	for;len(adjacents)==0; {
		logrus.Info("no adjected yet")
		time.Sleep(time.Second)
		adjacents = sm.GetAllAdjacents(fs)
	}

	sm.CreateAndSend(fs,adjacents[0],handlers.REQUEST_FILE_LIST,[]byte("/mnt/Vol1/Media/complete"))

	sm.WaitForShutdown()
}