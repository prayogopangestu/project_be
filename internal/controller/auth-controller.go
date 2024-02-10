package controller

import (
	"net/http"
	"project/internal/dto"
	"project/internal/models"
	"project/internal/service"
	"project/pkg/responsebuilder"
	"strconv"

	"github.com/gin-gonic/gin"
	// "strings"
)

// AuthController interface is a contract what this controller can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	// ChangePassword(ctx *gin.Context)
	// LoginBpjs(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

// NewAuthController creates a new instance of AuthController
func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := responsebuilder.BuildErrorResponse("Failed to process request", errDTO.Error(), responsebuilder.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)
	if v, ok := authResult.(models.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10), v.First_name+" "+v.Last_name, v.Nickname, strconv.FormatUint(uint64(v.Otoritas), 10), v.Email)
		// fmt.Printf(generatedToken)
		Data := struct {
			ID         uint   `json:"id" form:"id"`
			Nickname   string ` json:"nickname" `
			First_name string ` json:"first_name" binding:"required"`
			Last_name  string ` json:"last_name" binding:"required"`
			Email      string ` json:"email" binding:"required"`
			Phone      string ` json:"phone" binding:"required"`
			Otoritas   uint32 ` json:"otoritas" binding:"required"`
			Status     string `json:"status" binding:"required"`
			// DokterM    *models.Dokter_m `json:"dokter_m"`
		}{
			ID:         v.ID,
			Nickname:   v.Nickname,
			First_name: v.First_name,
			Last_name:  v.Last_name,
			Email:      v.Email,
			Phone:      v.Phone,
			Otoritas:   v.Otoritas,
			Status:     v.Status,
			// DokterM:    v.Dokter_m,
		}
		// v.Token = generatedToken
		response := responsebuilder.BuildResponseLogin(true, "OK!", generatedToken, Data)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := responsebuilder.BuildErrorResponse("Please check again your credential", "Invalid Credential", responsebuilder.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := responsebuilder.BuildErrorResponse("Failed to process request", errDTO.Error(), responsebuilder.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateEmail(registerDTO.Email) {
		response := responsebuilder.BuildErrorResponse("Failed to process request", "Duplicate email", responsebuilder.EmptyObj{})
		ctx.JSON(http.StatusConflict, response)
	} else {
		createdUser := c.authService.CreateUser(registerDTO)
		// token := c.jwtService.GenerateToken(strconv.FormatUint(uint64(createdUser.ID), 10))
		// fmt.Printf(token)
		// createdUser.Token = token
		response := responsebuilder.BuildResponse(true, "OK!", createdUser)
		ctx.JSON(http.StatusCreated, response)
	}
}

// func (c *authController) LoginBpjs(ctx *gin.Context) {
// 	// var loginDTO dto.LoginDTO
// 	// errDTO := ctx.ShouldBind(&loginDTO)
// 	// if errDTO != nil {
// 	// 	response := responsebuilder.BuildErrorResponse("Failed to process request", errDTO.Error(), responsebuilder.EmptyObj{})
// 	// 	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
// 	// 	return
// 	// }
// 	fmt.Println(ctx.Request.Header.Get("x-username"))
// 	authResult := c.authService.VerifyCredential(ctx.Request.Header.Get("x-username"), ctx.Request.Header.Get("x-password"))
// 	if v, ok := authResult.(models.User); ok {
// 		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.ID), 10), v.First_name+" "+v.Last_name)
// 		// fmt.Printf(generatedToken)
// 		Data := struct {
// 			Token string `json:"token"`
// 		}{
// 			Token: generatedToken,
// 		}
// 		// v.Token = generatedToken
// 		response := responsebuilder.BuildResponseBpjs("Ok", 200, Data)
// 		ctx.JSON(http.StatusOK, response)
// 		return
// 	}
// 	response := responsebuilder.BuildErrorResponse("Please check again your credential", "Invalid Credential", responsebuilder.EmptyObj{})
// 	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
// }
