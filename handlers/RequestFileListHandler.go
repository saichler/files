package handlers

import (
	. "github.com/saichler/habitat/service"
	. "github.com/saichler/files/model"
	"github.com/sirupsen/logrus"
)
import . "github.com/saichler/habitat"

type RequestFileListHandler struct {
}


func (h *RequestFileListHandler) Type() uint16 {
	return REQUEST_FILE_LIST
}

func (h *RequestFileListHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	dirname:=string(m.Data)
	logrus.Info("Requested Directory:"+dirname)
	fi,err:=NewDirInfo(dirname)
	if err!=nil {
		logrus.Error("No Such Directory:"+dirname)
		svm.CreateAndReply(service,m,REPLY_NO_SUCH_FILE,m.Data)
		return
	}
	logrus.Info("Replying with data for directory:"+dirname)
	svm.CreateAndReply(service,m,REPLY_FILE_LIST,fi.Bytes())
}
