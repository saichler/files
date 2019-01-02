package handlers

import (
	. "github.com/saichler/habitat/service"
	"github.com/sirupsen/logrus"
)
import . "github.com/saichler/habitat"

type StartFileListHandler struct {
}


func (h *StartFileListHandler) Type() uint16 {
	return Message_Type_Service_START
}

func (h *StartFileListHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	logrus.Info("File Service was started!")
}
