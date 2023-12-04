package api

import "github.com/gin-gonic/gin"

// GetBinary - возвращает данные карты по номеру
func (s *Server) GetBinary(ctx *gin.Context) {}

// CreateBinary - записывает новую карту в хранилище
func (s *Server) CreateBinary(ctx *gin.Context) {}

// GetAllBinaries - возвращает все карты пользователя
func (s *Server) GetAllBinaries(ctx *gin.Context) {}

// UpdateBinary - обновляет данные карты
func (s *Server) UpdateBinary(ctx *gin.Context) {}

// DeleteBinaries  - удаляет карту по номеру
func (s *Server) DeleteBinaries(ctx *gin.Context) {}

// DeleteAllBinaries - удаляет все карты пользователя
func (s *Server) DeleteAllBinaries(ctx *gin.Context) {}
