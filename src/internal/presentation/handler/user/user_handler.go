package user

import (
	"fmt"
	"net/http"

	usecase "go-gin-domain/internal/application/usecase/user"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindByUID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type userHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(
	userUsecase usecase.UserUsecase,
) UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
	}
}

type CreateUserRequestBody struct {
	LastName  string `json:"last_name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

type UpdateUserRequestBody struct {
	LastName  string `json:"last_name" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

func (h *userHandler) Create(c *gin.Context) {
	// 共通コンテキスト
	ctx := c.Request.Context()

	// バリデーションチェック
	var reqBody CreateUserRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		msg := fmt.Sprintf("バリデーションエラー: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}

	user, err := h.userUsecase.Create(ctx, reqBody.LastName, reqBody.FirstName, reqBody.Email)
	if err != nil {
		msg := fmt.Sprintf("Internal Server Error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *userHandler) FindAll(c *gin.Context) {
	// 共通コンテキスト
	ctx := c.Request.Context()

	users, err := h.userUsecase.FindAll(ctx)
	if err != nil {
		msg := fmt.Sprintf("Internal Server Error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *userHandler) FindByUID(c *gin.Context) {
	// 共通コンテキスト
	ctx := c.Request.Context()

	// バリデーションチェック
	uid := c.Param("uid")
	if uid == "" {
		msg := fmt.Sprintf("バリデーションエラー: %s", "uid is required")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}

	user, err := h.userUsecase.FindByUID(ctx, uid)
	if err != nil {
		msg := fmt.Sprintf("Internal Server Error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
		return
	}

	// userがnilの場合に空のオブジェクトを返す
	if user == nil {
		c.JSON(http.StatusOK, map[string]interface{}{})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Update(c *gin.Context) {
	// 共通コンテキスト
	ctx := c.Request.Context()

	// バリデーションチェック
	uid := c.Param("uid")
	if uid == "" {
		msg := fmt.Sprintf("バリデーションエラー: %s", "uid is required")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}

	var reqBody UpdateUserRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		msg := fmt.Sprintf("バリデーションエラー: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}

	user, err := h.userUsecase.Update(ctx, uid, reqBody.LastName, reqBody.FirstName, reqBody.Email)
	if err != nil {
		msg := fmt.Sprintf("Internal Server Error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) Delete(c *gin.Context) {
	// 共通コンテキスト
	ctx := c.Request.Context()

	// バリデーションチェック
	uid := c.Param("uid")
	if uid == "" {
		msg := fmt.Sprintf("バリデーションエラー: %s", "uid is required")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": msg,
		})
		return
	}

	user, err := h.userUsecase.Delete(ctx, uid)
	if err != nil {
		msg := fmt.Sprintf("Internal Server Error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
		return
	}

	c.JSON(http.StatusOK, user)
}
