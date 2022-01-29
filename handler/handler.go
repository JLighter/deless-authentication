package handler

import (
	"glog/services/logger"
	"glog/store"
)

type Handler struct {
  userStore *store.UserStore
  passwordStore *store.PasswordStore
  logger *logger.StandardLogger
}

func NewHandler(userStore *store.UserStore, passwordStore *store.PasswordStore, logger *logger.StandardLogger) *Handler {
  return &Handler{userStore, passwordStore, logger}
}
