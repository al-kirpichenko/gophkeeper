package api

import "github.com/gin-gonic/gin"

// GetText - возвращает текст по заголовку
func (s *Server) GetText(ctx *gin.Context) {}

// CreateText - записывает произвольный текст в хранилище
func (s *Server) CreateText(ctx *gin.Context) {}

// GetAllTexts - возвращает все тексты пользователя
func (s *Server) GetAllTexts(ctx *gin.Context) {}

// UpdateText - обновляет текст
func (s *Server) UpdateText(ctx *gin.Context) {}

// DeleteText  - удаляет текст по заголовку
func (s *Server) DeleteText(ctx *gin.Context) {}

// DeleteAllTexts - удаляет все тексты пользователя
func (s *Server) DeleteAllTexts(ctx *gin.Context) {}
