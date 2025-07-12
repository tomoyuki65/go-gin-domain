package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("新規ユーザー作成", func(t *testing.T) {
		uid := "xxxx-xxxx-xxxx-0001"
		lastName := "田中"
		firstName := "太郎"
		email := "t.tanaka@example.com"

		// 処理実行
		user := NewUser(uid, lastName, firstName, email)

		// 検証
		assert.NotNil(t, user)
		assert.Equal(t, user.ID, int64(0))
		assert.Equal(t, uid, user.UID)
		assert.Equal(t, lastName, user.LastName)
		assert.Equal(t, firstName, user.FirstName)
		assert.Equal(t, email, user.Email)
		assert.True(t, user.CreatedAt.IsZero())
		assert.True(t, user.UpdatedAt.IsZero())
		assert.Nil(t, user.DeletedAt)
	})
}

func TestUser_UpdateProfile(t *testing.T) {
	baseUser := func() *User {
		return &User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}
	}

	t.Run("プロフィール更新処理が正常終了", func(t *testing.T) {
		user := baseUser()
		oldUpdatedAt := user.UpdatedAt

		newLastName := "佐藤"
		newFirstName := "二郎"
		newEmail := "z.satou@example.com"

		// 処理実行
		err := user.UpdateProfile(newLastName, newFirstName, newEmail)

		// 検証
		assert.NoError(t, err)
		assert.Equal(t, newLastName, user.LastName)
		assert.Equal(t, newFirstName, user.FirstName)
		assert.Equal(t, newEmail, user.Email)
		assert.True(t, user.UpdatedAt.After(oldUpdatedAt))
	})

	t.Run("プロフィール更新処理でlast_nameが空の場合エラー", func(t *testing.T) {
		user := baseUser()

		newLastName := ""
		newFirstName := "二郎"
		newEmail := "z.satou@example.com"

		// 処理実行
		err := user.UpdateProfile(newLastName, newFirstName, newEmail)

		// 検証
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "last_nameは必須です。")
	})

	t.Run("プロフィール更新処理でfirst_nameが空の場合エラー", func(t *testing.T) {
		user := baseUser()

		newLastName := "佐藤"
		newFirstName := ""
		newEmail := "z.satou@example.com"

		// 処理実行
		err := user.UpdateProfile(newLastName, newFirstName, newEmail)

		// 検証
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "first_nameは必須です。")
	})

	t.Run("プロフィール更新処理でemailが空の場合エラー", func(t *testing.T) {
		user := baseUser()

		newLastName := "佐藤"
		newFirstName := "二郎"
		newEmail := ""

		// 処理実行
		err := user.UpdateProfile(newLastName, newFirstName, newEmail)

		// 検証
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "emailは必須です。")
	})

	t.Run("プロフィール更新処理で複数項目が空の場合エラー", func(t *testing.T) {
		user := baseUser()

		// 処理実行
		err := user.UpdateProfile("", "", "")

		// 検証
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "last_nameは必須です。")
		assert.Contains(t, err.Error(), "first_nameは必須です。")
		assert.Contains(t, err.Error(), "emailは必須です。")
	})
}

func TestUser_SetDelete(t *testing.T) {
	t.Run("論理削除設定がされること", func(t *testing.T) {
		user := &User{
			ID:        1,
			UID:       "xxxx-xxxx-xxxx-0001",
			LastName:  "田中",
			FirstName: "太郎",
			Email:     "t.tanaka@example.com",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: nil,
		}
		initialEmail := user.Email
		initialUpdatedAt := user.UpdatedAt

		// 処理実行
		user.SetDelete()

		// 検証
		assert.NotEqual(t, initialEmail, user.Email)
		assert.True(t, user.UpdatedAt.After(initialUpdatedAt))
		assert.NotNil(t, user.DeletedAt)
		assert.Equal(t, user.UpdatedAt, *user.DeletedAt)
	})
}
