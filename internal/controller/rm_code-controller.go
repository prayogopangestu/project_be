package controller

import (
	"fmt"
	"net/http"
	"project/internal/dto"
	"project/internal/models"
	"project/internal/repository"
	"project/internal/service"
	"project/pkg/responsebuilder"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	_ "github.com/golang-jwt/jwt"
	"github.com/mashingan/smapping"
)

// RmCodeController is a ...
type RmCodeController interface {
	All(context *gin.Context)
	FindByID(context *gin.Context)
	FindByFilter(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type rm_codeController struct {
	rm_codeRepo repository.RmCodeRepository
	jwtService  service.JWTService
}

// NewRmCodeController create a new instances of BoookController
func NewRmCodeController(rm_codeRepo repository.RmCodeRepository, jwtServ service.JWTService) RmCodeController {
	return &rm_codeController{
		rm_codeRepo: rm_codeRepo,
		jwtService:  jwtServ,
	}
}

func (c *rm_codeController) All(context *gin.Context) {
	rm_code, err := c.rm_codeRepo.AllRmCode()
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
		return
	}
	res := responsebuilder.BuildResponse(true, "OK", rm_code)
	context.JSON(http.StatusOK, res)

}

func (c *rm_codeController) FindByID(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("No param id was found", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	rm_code_my, err := c.rm_codeRepo.FindRmCodeByID(id)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
		return
	}
	if reflect.DeepEqual(rm_code_my, models.RmCodeM{}) {
		// if (rm_code_my == transaction.RmCode_t{}) {
		res := responsebuilder.BuildErrorResponse("Data not found", "No data with given id", nil)
		context.JSON(http.StatusNotFound, res)
		return
	}
	res := responsebuilder.BuildResponse(true, "OK", rm_code_my)
	context.JSON(http.StatusOK, res)

}

func (c *rm_codeController) FindByFilter(context *gin.Context) {

	var listRmCodeDTO dto.ListRmCodeDTO
	errDTO := context.ShouldBind(&listRmCodeDTO)
	catchErr := responsebuilder.CatchError(errDTO, context)
	if !catchErr.Status {
		context.AbortWithStatusJSON(http.StatusBadRequest, catchErr)
		return
	}
	fmt.Println(listRmCodeDTO)

	authHeader := context.GetHeader("Authorization")
	userID := c.getUserIDByToken(authHeader)
	_, err := strconv.ParseUint(userID, 10, 64)

	if err == nil {
		// listPendaftaranDTO = uint32(convertedUserID)
	}
	result, count_, err := c.rm_codeRepo.FindRmCodeByFilter(listRmCodeDTO)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
	} else {
		if (reflect.DeepEqual(result, models.RmCodeM{})) {
			res := responsebuilder.BuildErrorResponse("Data not found", "No data with given id", nil)
			context.JSON(http.StatusNotFound, res)
			return
		} else {
			res := responsebuilder.BuildResponse_table(true, "OK", *count_, result)
			context.JSON(http.StatusOK, res)
			return
		}

	}

}

func (c *rm_codeController) Insert(context *gin.Context) {
	var rm_codeCreateDTO dto.RmCodeCreateDTO
	errDTO := context.ShouldBind(&rm_codeCreateDTO)
	catchErr := responsebuilder.CatchError(errDTO, context)
	if !catchErr.Status {
		context.AbortWithStatusJSON(http.StatusBadRequest, catchErr)
		return
	}
	authHeader := context.GetHeader("Authorization")
	rm_codes := models.RmCodeM{}
	userID := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		rm_codeCreateDTO.CreateUserID = uint(convertedUserID)
	}

	err = smapping.FillStruct(&rm_codes, smapping.MapFields(&rm_codeCreateDTO))
	if err != nil {
		// log.Fatalf("Failed map %v: ", err)
		fmt.Println("Failed map : ", err)
	}

	result, err := c.rm_codeRepo.InsertRmCode(rm_codes)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
		return
	}
	response := responsebuilder.BuildResponse(true, "OK", result)
	context.JSON(http.StatusCreated, response)
}

func (c *rm_codeController) Update(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("No param id was found", err.Error(), nil)
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var rm_codeUpdateDTO dto.RmCodeUpdateDTO
	errDTO := context.ShouldBind(&rm_codeUpdateDTO)
	if errDTO != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", errDTO.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")

	userID := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		updateUserID := uint(convertedUserID)
		rm_codeUpdateDTO.UpdateUserID = &updateUserID
	}
	// token, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// _ = fmt.Sprintf("%v", claims["user_id"])
	// if c.rm_codeService.IsAllowedToEdit(userID, rm_codeUpdateDTO.ID) {
	// 	id, errID := strconv.ParseUint(userID, 10, 64)
	// 	if errID == nil {
	// 		rm_codeUpdateDTO.UserID = uint32(id)
	// 	}
	// 	result := c.rm_codeService.Update(rm_codeUpdateDTO)
	// 	response := responsebuilder.BuildResponse(true, "OK", result)
	// 	context.JSON(http.StatusOK, response)
	// } else {
	// 	response := responsebuilder.BuildErrorResponse("You dont have permission", "You are not the owner", responsebuilder.EmptyObj{})
	// 	context.JSON(http.StatusForbidden, response)
	// }
	rm_codes := models.RmCodeM{}
	err = smapping.FillStruct(&rm_codes, smapping.MapFields(&rm_codeUpdateDTO))
	if err != nil {
		// log.Fatalf("Failed map %v: ", err)
		fmt.Println("Failed map : ", err)
	}
	rm_codes.ID = uint(id)
	result, err := c.rm_codeRepo.UpdateRmCode(rm_codes)
	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
		return
	}
	response := responsebuilder.BuildResponse(true, "OK", result)
	context.JSON(http.StatusOK, response)
}

func (c *rm_codeController) Delete(context *gin.Context) {
	var rm_code models.RmCodeM
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := responsebuilder.BuildErrorResponse("Failed tou get id", "No param id were found", nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	rm_code.ID = uint(id)
	authHeader := context.GetHeader("Authorization")

	userID := c.getUserIDByToken(authHeader)
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err == nil {
		deleteUserId := uint(convertedUserID)
		rm_code.DeleteUserID = &deleteUserId
	}
	// token, errToken := c.jwtService.ValidateToken(authHeader)
	// if errToken != nil {
	// 	panic(errToken.Error())
	// }
	// claims := token.Claims.(jwt.MapClaims)
	// _ = fmt.Sprintf("%v", claims["user_id"])
	// if c.rm_codeService.IsAllowedToEdit(userID, uint64(rm_code.ID)) {
	// 	c.rm_codeService.Delete(rm_code)
	// 	res := responsebuilder.BuildResponse(true, "Deleted", responsebuilder.EmptyObj{})
	// 	context.JSON(http.StatusOK, res)
	// } else {
	// 	response := responsebuilder.BuildErrorResponse("You dont have permission", "You are not the owner", responsebuilder.EmptyObj{})
	// 	context.JSON(http.StatusForbidden, response)
	// }
	err = c.rm_codeRepo.DeleteRmCode(rm_code)

	if err != nil {
		res := responsebuilder.BuildErrorResponse("Failed to process delete request", err.Error(), nil)
		context.JSON(http.StatusBadRequest, res)
		return
	}

	res := responsebuilder.BuildResponse(true, "Deleted", nil)
	context.JSON(http.StatusOK, res)
}

func (c *rm_codeController) getUserIDByToken(token string) string {
	aToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := aToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
