package service

import (
	. "github.com/saichler/files/handlers"
	. "github.com/saichler/habitat"
	. "github.com/saichler/habitat/service"
)

type FileService struct {
	svm *ServiceManager
	sid *ServiceID
}

const (
	FILE_SERVICE_TOPIC="File Service Topic"
)

func (s *FileService) Name() string {
	return "File Service"
}

func (s *FileService) ServiceID() *ServiceID {
	return s.sid
}

func (s *FileService) ServiceManager() *ServiceManager {
	return s.svm
}

func (s *FileService) Init(svm *ServiceManager,componentID uint16) {
	s.svm = svm
	s.sid = NewServiceID(svm.HID(),componentID,FILE_SERVICE_TOPIC)
}

func (s *FileService) ServiceMessageHandlers()[]ServiceMessageHandler {
	return []ServiceMessageHandler{
		&RequestFileHandler{},
		&ReplyFileHandler{},
		&ReplyNoSuchFileHandler{},
		&RequestFileListHandler{},
		&ReplyFileListHandler{},
	    &StartFileListHandler{}}
}