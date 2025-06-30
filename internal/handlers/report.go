package handlers

import (
	"bus_depot/internal/models"
	"bus_depot/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ReportHandler struct {
	service *service.ReportService
}

func NewReportHandler(s *service.ReportService) *ReportHandler {
	return &ReportHandler{s}
}


// Create godoc
// @Summary      Создать новый отчет
// @Description  Создает новый отчет, связывая его с текущим механиком (из контекста)
// @Tags         reports
// @Accept       json
// @Produce      json
// @Param        report  body      models.Report  true  "Данные отчета"
// @Success      201     {object}  map[string]string
// @Failure      400     {object}  map[string]string
// @Failure      401     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /reports [post]
func (h *ReportHandler) CreateReport(c *gin.Context) {
	var report models.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}
	report.MechanicID = userID.(uint)

	if err := h.service.CreateReport(&report); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Отчёт успешно создан"})
}

// GetAll godoc
// @Summary      Получить все отчеты
// @Description  Возвращает список всех отчетов
// @Tags         reports
// @Produce      json
// @Success      200   {array}   models.Report
// @Failure      500   {object}  map[string]string
// @Router       /reports [get]
func (h *ReportHandler) GetAllReports(c *gin.Context) {
	reports, err := h.service.GetAllReports()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reports)
}

// GetByID godoc
// @Summary      Получить отчет по ID
// @Description  Возвращает отчет по уникальному идентификатору
// @Tags         reports
// @Produce      json
// @Param        id   path      int  true  "ID отчета"
// @Success      200  {object}  models.Report
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /reports/{id} [get]
func (h *ReportHandler) GetReportByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID отчёта"})
		return
	}

	report, err := h.service.GetReportByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Отчёт не найден"})
		return
	}

	c.JSON(http.StatusOK, report)
}


// Delete godoc
// @Summary      Удалить отчет
// @Description  Удаляет отчет по ID
// @Tags         reports
// @Produce      json
// @Param        id   path      int  true  "ID отчета"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /reports/{id} [delete]
func (h *ReportHandler) DeleteReport(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}

	if err := h.service.DeleteReport(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Отчёт успешно удалён"})
}
