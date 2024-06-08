package controllers

import (
	"go-crud/pkg/models"
	"go-crud/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("name")
	
	user,err := uc.UserService.GetUser((&username))
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) UpdateUser( ctx *gin.Context) {
	var user models.User
	if err := ctx.Copy().ShouldBindJSON(&user);
	err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message": err.Error()})
		return
	}

	err := uc.UserService.UpdateUser(&user)
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("name")
	err := uc.UserService.DeleteUser((&username))
	if err != nil{
		ctx.JSON(http.StatusBadGateway,gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (uc *UserController) RegisterUserRoutes(router *gin.RouterGroup) {
	userroute := router.Group("/user")
	userroute.POST("/create", uc.CreateUser)
	userroute.GET("/:name", uc.GetUser)
	userroute.GET("/allusers", uc.GetAllUsers)
	userroute.PATCH("/update", uc.UpdateUser)
	userroute.DELETE("/:name", uc.DeleteUser)
}