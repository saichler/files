package handlers

import (
	. "github.com/saichler/habitat/service"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
)
import . "github.com/saichler/habitat"

type RequestFileHandler struct {
}


func (h *RequestFileHandler) Type() uint16 {
	return REQUEST_FILE
}

func (h *RequestFileHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	filename:=string(m.Data)
	data,err:=ioutil.ReadFile(filename)
	if err!=nil {
		svm.CreateAndReply(service,m,REPLY_NO_SUCH_FILE,m.Data)
	}
	logrus.Info("File size="+strconv.Itoa(len(data)))
	svm.CreateAndReply(service,m,REPLY_FILE,data)
}
