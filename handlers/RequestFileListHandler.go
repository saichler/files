package handlers

import (
	. "github.com/saichler/files/model"
	. "github.com/saichler/habitat"
	. "github.com/saichler/habitat/service"
	. "github.com/saichler/utils/golang"
)

type RequestFileListHandler struct {
}


func (h *RequestFileListHandler) Type() uint16 {
	return REQUEST_FILE_LIST
}

func (h *RequestFileListHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	dirname:=string(m.Data)
	Info("Requested Directory:"+dirname)
	fi,err:=NewDirInfo(dirname)
	if err!=nil {
		Error("No Such Directory:"+dirname)
		svm.CreateAndReply(service,m,REPLY_NO_SUCH_FILE,m.Data)
		return
	}
	Info("Replying with data for directory:"+dirname)
	svm.CreateAndReply(service,m,REPLY_FILE_LIST,fi.Bytes())
}
