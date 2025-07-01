package handlers

import (
	"bus_depot/internal/models"
	"bus_depot/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handlers struct {
	Auth *AuthHandler
	Bus  *BusHandler
	User *UserHandler
}

type BusHandler struct {
	service *service.BusService
}

func NewBusHandler(service *service.BusService) *BusHandler {
	return &BusHandler{service: service}
}

// CreateBus godoc
// @Summary      Создать новый автобус
// @Description  Добавляет новый автобус в систему
// @Tags         buses
// @Accept       json
// @Produce      json
// @Param        bus   body      models.Bus  true  "Данные автобуса"
// @Success      201   {object}  models.Bus
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /buses [post]
func (h *BusHandler) CreateBus(c *gin.Context) {
	var bus models.Bus
	if err := c.ShouldBindJSON(&bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateBus(&bus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании автобуса"})
		return
	}

	c.JSON(http.StatusCreated, bus)
}

// GetAllBuses godoc
// @Summary      Получить список всех автобусов
// @Description  Возвращает массив всех автобусов
// @Tags         buses
// @Produce      json
// @Success      200   {array}   models.Bus
// @Failure      500   {object}  map[string]string
// @Router       /buses [get]
func (h *BusHandler) GetAllBuses(c *gin.Context) {
	buses, err := h.service.GetAllBuses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении списка автобусов"})
		return
	}

	c.JSON(http.StatusOK, buses)
}

// GetBusByID godoc
// @Summary      Получить автобус по ID
// @Description  Возвращает автобус по уникальному идентификатору
// @Tags         buses
// @Produce      json
// @Param        id    path      int  true  "ID автобуса"
// @Success      200   {object}  models.Bus
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /buses/{id} [get]
func (h *BusHandler) GetBusByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	bus, err := h.service.GetBusByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Автобус не найден"})
		return
	}

	c.JSON(http.StatusOK, bus)
}

// UpdateBus godoc
// @Summary      Обновить данные автобуса
// @Description  Обновляет информацию об автобусе по ID
// @Tags         buses
// @Accept       json
// @Produce      json
// @Param        id    path      int        true  "ID автобуса"
// @Param        bus   body      models.Bus true  "Обновленные данные автобуса"
// @Success      200   {object}  models.Bus
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /buses/{id} [put]
func (h *BusHandler) UpdateBus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	var bus models.Bus
	if err := c.ShouldBindJSON(&bus); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bus.ID = uint(id)
	if err := h.service.UpdateBus(&bus); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении автобуса"})
		return
	}

	c.JSON(http.StatusOK, bus)
}

// DeleteBus godoc
// @Summary      Удалить автобус
// @Description  Удаляет автобус по ID
// @Tags         buses
// @Produce      json
// @Param        id    path      int  true  "ID автобуса"
// @Success      200   {object}  map[string]string
// @Failure      400   {object}  map[string]string
// @Failure      500   {object}  map[string]string
// @Router       /buses/{id} [delete]
func (h *BusHandler) DeleteBus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID"})
		return
	}

	if err := h.service.DeleteBus(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении автобуса"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Автобус удален"})
}

// AssignDriver godoc
// @Summary      Назначить водителя автобусу
// @Description  Назначает водителя указанному автобусу
// @Tags         buses
// @Accept       json
// @Produce      json
// @Param        id         path      int  true  "ID автобуса"
// @Param        driverId   body      int  true  "ID водителя"
// @Success      200        {object}  map[string]string
// @Failure      400        {object}  map[string]string
// @Failure      500        {object}  map[string]string
// @Router       /buses/{id}/assign-driver [post]
func (h *BusHandler) AssignDriver(c *gin.Context) {
	var input struct {
		DriverID uint `json:"driver_id"`
	}
	busIDParam := c.Param("id")

	busID, err := strconv.ParseUint(busIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	err = h.service.AssignDriver(uint(busID), input.DriverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось назначить водителя"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Водитель успешно назначен"})
}

// AssignMechanic godoc
// @Summary      Назначить механика автобусу
// @Description  Назначает механика указанному автобусу
// @Tags         buses
// @Accept       json
// @Produce      json
// @Param        id          path      int  true  "ID автобуса"
// @Param        mechanicId  body      int  true  "ID механика"
// @Success      200         {object}  map[string]string
// @Failure      400         {object}  map[string]string
// @Failure      500         {object}  map[string]string
// @Router       /buses/{id}/assign-mechanic [post]
func (h *BusHandler) AssignMechanic(c *gin.Context) {
	busIDParam := c.Param("id")
	busID, err := strconv.ParseUint(busIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bus ID"})
		return
	}

	var input struct {
		MechanicID uint `json:"mechanic_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = h.service.AssignMechanic(uint(busID), input.MechanicID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось назначить механика"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Механик успешно назначен"})
}
