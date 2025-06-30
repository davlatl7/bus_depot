package handlers

import (
	"bus_depot/internal/models"
	"bus_depot/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type WorkScheduleHandler struct {
	service *service.WorkScheduleService
}

func NewWorkScheduleHandler(s *service.WorkScheduleService) *WorkScheduleHandler {
	return &WorkScheduleHandler{s}
}

// Create godoc
// @Summary      Создать график работы
// @Description  Добавляет новый график работы
// @Tags         workschedules
// @Accept       json
// @Produce      json
// @Param        schedule  body      models.WorkSchedule  true  "Данные графика работы"
// @Success      201       {object}  map[string]string
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /schedules [post]
func (h *WorkScheduleHandler) CreateSchedule(c *gin.Context) {
	var schedule models.WorkSchedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных: " + err.Error()})
		return
	}

	if err := h.service.CreateSchedule(&schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать график: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "График успешно создан"})
}


// GetAll godoc
// @Summary      Получить все графики работы
// @Description  Возвращает список всех графиков работы
// @Tags         workschedules
// @Produce      json
// @Success      200  {array}   models.WorkSchedule
// @Failure      500  {object}  map[string]string
// @Router       /schedules [get]
func (h *WorkScheduleHandler) GetAllSchedules(c *gin.Context) {
	schedules, err := h.service.GetAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения графиков: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, schedules)
}


// GetByID godoc
// @Summary      Получить график работы по ID
// @Description  Возвращает график работы по уникальному идентификатору
// @Tags         workschedules
// @Produce      json
// @Param        id   path      int  true  "ID графика работы"
// @Success      200  {object}  models.WorkSchedule
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /schedules/{id} [get]
func (h *WorkScheduleHandler) GetScheduleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	schedule, err := h.service.GetScheduleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "График не найден"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}


// Update godoc
// @Summary      Обновить график работы
// @Description  Обновляет график работы по ID
// @Tags         workschedules
// @Accept       json
// @Produce      json
// @Param        id        path      int                 true  "ID графика работы"
// @Param        schedule  body      models.WorkSchedule true  "Обновленные данные графика работы"
// @Success      200       {object}  map[string]string
// @Failure      400       {object}  map[string]string
// @Failure      500       {object}  map[string]string
// @Router       /schedules/{id} [put]
func (h *WorkScheduleHandler) UpdateSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	var schedule models.WorkSchedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных: " + err.Error()})
		return
	}

	schedule.ID = uint(id)
	if err := h.service.UpdateSchedule(&schedule); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить график: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "График успешно обновлён"})
}


// Delete godoc
// @Summary      Удалить график работы
// @Description  Удаляет график работы по ID
// @Tags         workschedules
// @Produce      json
// @Param        id   path      int  true  "ID графика работы"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /schedules/{id} [delete]
func (h *WorkScheduleHandler) DeleteSchedule(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	if err := h.service.DeleteSchedule(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления графика: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "График успешно удалён"})
}



// GetMySchedule godoc
// @Summary      Получить мой график работы
// @Description  Возвращает график работы для текущего авторизованного водителя
// @Tags         workschedules
// @Produce      json
// @Success      200  {array}   models.MyScheduleResponse
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Security     ApiKeyAuth
// @Router       /schedules/my [get]
func (h *WorkScheduleHandler) GetMySchedule(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		return
	}

	schedules, err := h.service.GetByDriverID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения графика: " + err.Error()})
		return
	}

	// Сформируем облегчённый ответ
	var response []models.MyScheduleResponse
	for _, s := range schedules {
		response = append(response, models.MyScheduleResponse{
			ID:        s.ID,
			TimeRange: s.TimeRange,
			LineName:  s.LineName,
			Bus: models.BusShortInfo{
				Model1: s.Bus.Model1,
				Number: s.Bus.Number,
				Status: s.Bus.Status,
			},
		})
	}

	c.JSON(http.StatusOK, response)
}

