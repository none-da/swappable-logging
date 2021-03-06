package service

import (
	logger "github.com/none-da/swappable-logging/pkg/logger"
)

type Service struct {
	log logger.Logger
}

func CreateService(loggerName string) Service {
	svc := Service{
		log: logger.CreateLogger(loggerName),
	}
	return svc
}

func DoSomething(svc *Service, action string) {
	svc.log.Debug(action)
}
