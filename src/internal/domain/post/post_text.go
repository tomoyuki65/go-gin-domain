package post

// カスタムエラー用の構造体を定義
type ErrInvalidLength struct{}

func (e *ErrInvalidLength) Error() string {
	return "文字数は10文字以下にして下さい。"
}

// 値オブジェクトの定義
type Text struct {
	value string
}

// コンストラクタ
func NewText(value string) (Text, error) {
	// 文字数チェック
	if len(value) > 10 {
		return Text{}, &ErrInvalidLength{}
	}

	return Text{value: value}, nil
}

// DBから復元するためのコンストラクタ（チェック処理無し）
func ReconstituteText(value string) Text {
	return Text{value: value}
}

// 値を返すメソッド
func (t Text) Value() string {
	return t.value
}
