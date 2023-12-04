package api

import "github.com/gin-gonic/gin"

// GetPassword - возвращает пару логин-пароль по логину
func (s *Server) GetPassword(ctx *gin.Context) {}

// CreatePassword - записывает пару логин-пароль в хранилище
func (s *Server) CreatePassword(ctx *gin.Context) {}

// GetAllPasswords - возвращает все пары логин-пароль пользователя
func (s *Server) GetAllPasswords(ctx *gin.Context) {}

// UpdatePassword - обновляет пару логин-пароль
func (s *Server) UpdatePassword(ctx *gin.Context) {}

// DeletePassword - удаляет пару логин-пароль
func (s *Server) DeletePassword(ctx *gin.Context) {}

// DeleteAllPasswords -удаляет все записи логин-пароль пользователя
func (s *Server) DeleteAllPasswords(ctx *gin.Context) {}
