package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_final/models"
	"go_final/repositories"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

type UserHandler interface {
	SignInUser(*gin.Context)
	CreateUser(*gin.Context)
	GetUser(*gin.Context)
	GetAllUsers(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler() UserHandler {
	return &userHandler{
		repo: repositories.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *userHandler) SignInUser(ctx *gin.Context) {
	var user models.UserLogin
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(user)

	dbUser, err := h.repo.GetByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "No Such User Found"})
		return
	}

	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		token := GenerateToken(dbUser.ID, dbUser.IsAdmin)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignedIN", "token": token})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password didn't match"})
	return
}

func (h *userHandler) CreateUser(ctx *gin.Context) {
	var input models.UserRegister
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		IsAdmin:  false,
	}

	hashPassword(&user.Password)
	user, err := h.repo.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = ""
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("user_id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.repo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) GetAllUsers(ctx *gin.Context) {
	user, err := h.repo.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var input models.UserUpdate
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("user_id")
	intID, _ := strconv.Atoi(id)

	user, err := h.repo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No such user id database!"})
		return
	}

	user.ID = uint(intID)
	user.Name = input.Name
	user.Password = input.Password
	hashPassword(&user.Password)

	user, err = h.repo.UpdateUser(user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("user_id")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)

	user, err := h.repo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
