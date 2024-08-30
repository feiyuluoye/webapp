package routes

import (
	"github.com/gin-gonic/gin"
	"tarot-app/controllers"
	"tarot-app/database"
	"tarot-app/repositories"
	"tarot-app/services"
)

// SetupRoutes 设置Gin路由
func SetupRoutes(r *gin.Engine) {
	// 初始化存储库、服务和控制器
	db := database.DB
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	// 注册路由
	r.POST("/login", authController.Login)
	r.POST("/register", authController.Register)
	{
		spreadRepo := repositories.NewTarotSpreadRepository()
		spreadController := controllers.NewTarotSpreadController(spreadRepo)

		r.POST("/spreads", spreadController.CreateSpread)
		r.GET("/spreads/:id", spreadController.GetSpread)
		r.PUT("/spreads/:id", spreadController.UpdateSpread)
		r.DELETE("/spreads/:id", spreadController.DeleteSpread)
		r.GET("/spreads", spreadController.ListSpreads)
	}
}
