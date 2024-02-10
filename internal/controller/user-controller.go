package controller

import (
	"fmt"
	"net/http"
	"project/internal/dto"
	"project/internal/service"
	"project/pkg/responsebuilder"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// UserController is a ....
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	FindByFilter(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

// NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", errDTO.Error(), responsebuilder.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	id, err := strconv.Atoi(userID)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	userUpdateDTO.ID = uint(id)
	u, err := c.userService.Update(userUpdateDTO)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Database Error", err.Error(), responsebuilder.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := responsebuilder.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")
	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		// responsebuilder.WriteError(err.Error())
		panic(err.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := responsebuilder.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}

func (c *userController) FindByFilter(context *gin.Context) {

	var ListUserDTO dto.ListUserDTO
	errDTO := context.ShouldBind(&ListUserDTO)

	if errDTO != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", errDTO.Error(), responsebuilder.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
		return
	}

	fmt.Println(ListUserDTO)

	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	_, err := strconv.ParseUint(userID, 10, 64)

	// if err == nil {
	// 	// listPendaftaranDTO = uint32(convertedUserID)
	// }
	result, count_, err := c.userService.ListUserByFilter(ListUserDTO)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
	} else {
		// if (reflect.DeepEqual(result, models.DokterRuangan_m{})) {
		// 	res := responsebuilder.BuildErrorResponse("Data not found", "No data with given id", nil)
		// 	context.JSON(http.StatusNotFound, res)
		// 	return
		// } else {
		res := responsebuilder.BuildResponse_table(true, "OK", *count_, result)
		context.JSON(http.StatusOK, res)
		return
		// }

	}

}

func (c *userController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		// responsebuilder.WriteError(err.Error())
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
