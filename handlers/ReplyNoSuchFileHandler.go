package handlers

import (
	. "github.com/saichler/habitat/service"
	"github.com/sirupsen/logrus"
)
import . "github.com/saichler/habitat"

type ReplyNoSuchFileHandler struct {
}


func (h *ReplyNoSuchFileHandler) Type() uint16 {
	return REPLY_NO_SUCH_FILE
}

func (h *ReplyNoSuchFileHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	filename:=string(m.Data)
	logrus.Error("No Such file/dir:"+filename)
}
