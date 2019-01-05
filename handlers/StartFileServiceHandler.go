package handlers

import (
	. "github.com/saichler/habitat"
	. "github.com/saichler/habitat/service"
	. "github.com/saichler/utils/golang"
)

type StartFileListHandler struct {
}


func (h *StartFileListHandler) Type() uint16 {
	return Message_Type_Service_START
}

func (h *StartFileListHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	Info("File Service was started!")
}
