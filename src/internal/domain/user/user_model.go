package user

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID        int64      `json:"-"`
	UID       string     `json:"uid"`
	LastName  string     `json:"last_name"`
	FirstName string     `json:"first_name"`
	Email     string     `json:"email"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func NewUser(uid, lastName, firstName, email string) *User {
	return &User{
		ID:        0,
		UID:       uid,
		LastName:  lastName,
		FirstName: firstName,
		Email:     email,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
}

// プロフィール更新
func (u *User) UpdateProfile(lastName, firstName, email string) error {
	// パラメータチェック
	var errMsg []string
	if lastName == "" {
		errMsg = append(errMsg, "last_nameは必須です。")
	}
	if firstName == "" {
		errMsg = append(errMsg, "first_nameは必須です。")
	}
	if email == "" {
		errMsg = append(errMsg, "emailは必須です。")
	}
	if len(errMsg) > 0 {
		msg := fmt.Sprintf("バリデーションエラー: %s", strings.Join(errMsg, ", "))
		return fmt.Errorf("%s", msg)
	}

	// 更新
	u.LastName = lastName
	u.FirstName = firstName
	u.Email = email
	u.UpdatedAt = time.Now()

	return nil
}

// 論理削除設定
func (u *User) SetDelete() {
	// 現在の日時を文字列で取得
	date := time.Now()
	dateString := date.Format("2006-01-02 15:04:05")

	// 更新用のemailの値を設定
	updateEmail := u.Email + dateString

	// 更新
	u.Email = updateEmail
	u.UpdatedAt = date
	u.DeletedAt = &date
}
