package handlers

import (
	. "github.com/saichler/habitat/service"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)
import . "github.com/saichler/habitat"

type ReplyFileHandler struct {
}


func (h *ReplyFileHandler) Type() uint16 {
	return REPLY_FILE
}

func (h *ReplyFileHandler) HandleMessage(svm *ServiceManager,service Service,m *Message) {
	logrus.Info("Wring file...")
	ioutil.WriteFile("/tmp/aq.mov",m.Data,0777)
	logrus.Info("Done")
	svm.Shutdown()
}
