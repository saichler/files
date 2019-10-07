package handlers

import (
	. "github.com/saichler/habitat/golang/habitat"
	. "github.com/saichler/habitat/golang/service"
	. "github.com/saichler/utils/golang"
)

type StartFileListHandler struct {
}

func (h *StartFileListHandler) Type() uint16 {
	return Message_Type_Service_START
}

func (h *StartFileListHandler) HandleMessage(svm *ServiceManager, service Service, m *Message) {
	Info("File Service was started!")
}
