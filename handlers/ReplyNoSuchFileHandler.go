package handlers

import (
	. "github.com/saichler/habitat"
	. "github.com/saichler/habitat/service"
	. "github.com/saichler/utils/golang"
)

type ReplyNoSuchFileHandler struct {
}


func (h *ReplyNoSuchFileHandler) Type() uint16 {
	return REPLY_NO_SUCH_FILE
}

func (h *ReplyNoSuchFileHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	filename:=string(m.Data)
	Error("No Such file/dir:"+filename)
}
