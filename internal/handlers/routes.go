package handlers

import (
	"github.com/gin-gonic/gin"
	"bus_depot/internal/middleware"
)

func InitRoutes(router *gin.Engine, authHandler *AuthHandler, userHandler *UserHandler, busHandler *BusHandler, scheduleHandler *WorkScheduleHandler, reportHandler *ReportHandler) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "🚍 Bus Depot Server is running"})
	})

	
	
	authG := router.Group("/auth")
	{
		authG.POST("/register", authHandler.Register)
		authG.POST("/login", authHandler.Login)
	}

	
	busesG := router.Group("/buses", middleware.IsDirector())
	{
		busesG.POST("", busHandler.CreateBus)
		busesG.GET("", busHandler.GetAllBuses)
		busesG.GET("/:id", busHandler.GetBusByID)
		busesG.PUT("/:id", busHandler.UpdateBus)
		busesG.DELETE("/:id", busHandler.DeleteBus)
		busesG.PUT("/:id/assign-driver", busHandler.AssignDriver)
		busesG.PUT("/:id/assign-mechanic", busHandler.AssignMechanic)
	}

	
	usersG := router.Group("/users", middleware.IsDirector())
	{
		usersG.GET("", userHandler.GetAllUsers)
		usersG.POST("", userHandler.CreateUser)
		usersG.PUT("/:id", userHandler.UpdateUser)
		usersG.DELETE("/:id", userHandler.DeleteUser)
	}

	
	schedulesG := router.Group("/schedules", middleware.IsDirector())
	{
		schedulesG.POST("", scheduleHandler.CreateSchedule)
		schedulesG.GET("", scheduleHandler.GetAllSchedules)
		schedulesG.GET("/:id", scheduleHandler.GetScheduleByID)
		schedulesG.PUT("/:id", scheduleHandler.UpdateSchedule)
		schedulesG.DELETE("/:id", scheduleHandler.DeleteSchedule)
	}

	scheduleDriverG := router.Group("/schedules")
	{

		scheduleDriverG.GET("/my", middleware.IsDriver(), scheduleHandler.GetMySchedule)
	
	}

	reportG := router.Group("/reports")
	{
		reportG.GET("", middleware.IsDirector(), reportHandler.GetAllReports)
		reportG.GET(":id", middleware.IsDirector(), reportHandler.GetReportByID)
		reportG.POST("", middleware.IsMaster(), reportHandler.CreateReport)
		reportG.DELETE(":id", middleware.IsDirector(), reportHandler.DeleteReport)
	}
}
