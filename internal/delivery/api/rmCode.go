package api

import (
	"project/internal/controller"
	"project/internal/middleware"
	"project/internal/repository"
	"project/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RmCodeRoutes(r *gin.Engine, db *gorm.DB, jwtService service.JWTService) {
	var (
		repo       repository.RmCodeRepository = repository.NewRmCodeRepository(db)
		controller controller.RmCodeController = controller.NewRmCodeController(repo, jwtService)
	)

	rm_codeGroup := r.Group("api/rm_code", middleware.AuthorizeJWT(jwtService))
	r.GET("api/rm_code_/FindSelect/:cari", controller.FindByFilter)

	{
		rm_codeGroup.GET("/", controller.All)
		rm_codeGroup.POST("/", controller.Insert)
		rm_codeGroup.GET("/Find/:id", controller.FindByID)
		rm_codeGroup.POST("/cari", controller.FindByFilter)
		rm_codeGroup.PUT("/:id", controller.Update)
		rm_codeGroup.DELETE("/:id", controller.Delete)
	}
}
