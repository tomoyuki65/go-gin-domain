package post

import (
	"errors"
	"fmt"
	"net/http"

	usecase "go-gin-domain/internal/application/usecase/post"
	domain "go-gin-domain/internal/domain/post"

	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
}

type postHandler struct {
	postUsecase usecase.PostUsecase
}

func NewPostHandler(
	postUsecase usecase.PostUsecase,
) PostHandler {
	return &postHandler{
		postUsecase: postUsecase,
	}
}

type CreatePostRequestBody struct {
	Text string `json:"text" binding:"required"`
}

func (h *postHandler) Create(c *gin.Context) {
	// 共通コンテキスト
	ctx := c.Request.Context()

	// バリデーションチェック
	var reqBody CreatePostRequestBody
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		msg := fmt.Sprintf("バリデーションエラー: %s", err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"message": msg,
		})
		return
	}

	post, err := h.postUsecase.Create(ctx, reqBody.Text)
	if err != nil {
		// カスタムエラー判定（バリデーションエラーかを判定）
		var ErrInvalidLength *domain.ErrInvalidLength
		if errors.As(err, &ErrInvalidLength) {
			// バリデーションエラーの場合
			msg := fmt.Sprintf("Unprocessable Entity: %s", err.Error())
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"message": msg,
			})
			return
		} else {
			// サーバーエラーの場合
			msg := fmt.Sprintf("Internal Server Error: %s", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": msg,
			})
			return
		}
	}

	// postをDTO用の関数で変換して返す
	c.JSON(http.StatusCreated, domain.ToResponse(post))
}

func (h *postHandler) FindAll(c *gin.Context) {
	// 共通コンテキスト
	ctx := c.Request.Context()

	posts, err := h.postUsecase.FindAll(ctx)
	if err != nil {
		msg := fmt.Sprintf("Internal Server Error: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": msg,
		})
		return
	}

	// レスポンス用のスライスを定義
	resPosts := make([]*domain.PostResponse, 0, len(posts))

	// ループ処理でpostをDTO用の関数で変換し、レスポンス用のスライスに追加
	for _, post := range posts {
		resPosts = append(resPosts, domain.ToResponse(post))
	}

	c.JSON(http.StatusOK, resPosts)
}
