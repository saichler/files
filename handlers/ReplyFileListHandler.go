package handlers

import (
	. "github.com/saichler/habitat/service"
	. "github.com/saichler/files/model"
	"github.com/saichler/utils/golang"
	"github.com/sirupsen/logrus"
)
import . "github.com/saichler/habitat"

type ReplyFileListHandler struct {
}


func (h *ReplyFileListHandler) Type() uint16 {
	return REPLY_FILE_LIST
}

func (h *ReplyFileListHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	logrus.Info("Received File List")
	ba:=utils.NewByteSliceWithData(m.Data,0)
	fi:=FromBytes(ba)
	fi.Print()
}
