package api

import "github.com/gin-gonic/gin"

// GetCard - возвращает данные карты по номеру
func (s *Server) GetCard(ctx *gin.Context) {}

// CreateCard - записывает новую карту в хранилище
func (s *Server) CreateCard(ctx *gin.Context) {}

// GetAllCards - возвращает все карты пользователя
func (s *Server) GetAllCards(ctx *gin.Context) {}

// UpdateCard - обновляет данные карты
func (s *Server) UpdateCard(ctx *gin.Context) {}

// DeleteCard  - удаляет карту по номеру
func (s *Server) DeleteCard(ctx *gin.Context) {}

// DeleteAllCards - удаляет все карты пользователя
func (s *Server) DeleteAllCards(ctx *gin.Context) {}
