package post

// エンティティの定義
type Post struct {
	// フィールドはプライベートにし、値オブジェクト型を使用
	text Text
}

// レスポンス用の構造体を定義
type PostResponse struct {
	Text string `json:"text"`
}

// コンストラクタ
func NewPost(text string) (*Post, error) {
	// 値オブジェクトを利用してtextをチェック
	newText, err := NewText(text)
	if err != nil {
		return nil, err
	}

	return &Post{text: newText}, nil
}

// DBから復元するためのコンストラクタ（チェック処理無し）
func ReconstitutePost(text string) *Post {
	return &Post{text: ReconstituteText(text)}
}

// textフィールドの値を返すメソッド
func (p *Post) TextValue() string {
	return p.text.Value()
}

// DTO（Data Transfer Object）用の関数
func ToResponse(p *Post) *PostResponse {
	return &PostResponse{
		Text: p.TextValue(),
	}
}
