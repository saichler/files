package handlers

import (
	. "github.com/saichler/habitat/service"
	. "github.com/saichler/habitat"
	"io/ioutil"
	. "github.com/saichler/utils/golang"
) 

type ReplyFileHandler struct {
}


func (h *ReplyFileHandler) Type() uint16 {
	return REPLY_FILE
}

func (h *ReplyFileHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	Info("Wring file...")
	ioutil.WriteFile("/tmp/aq.mov",m.Data,0777)
	Info("Done")
	svm.Shutdown()
}
