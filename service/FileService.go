package service

import (
	. "github.com/saichler/files/handlers"
	. "github.com/saichler/habitat/service"
)

type FileService struct {
	svm *ServiceManager
}

func (s *FileService) SID() uint16 {
	return 11
}

func (s *FileService) Name() string {
	return "File Service"
}

func (s *FileService) ServiceMessageHandlers()[]ServiceMessageHandler {
	return []ServiceMessageHandler{
		&RequestFileHandler{},
		&ReplyFileHandler{},
		&ReplyNoSuchFileHandler{},
		&RequestFileListHandler{},
		&ReplyFileListHandler{}}
}

func(s *FileService) SetManager(svm *ServiceManager) {
	s.svm = svm
}

func(s *FileService) GetManager() *ServiceManager{
	return s.svm
}