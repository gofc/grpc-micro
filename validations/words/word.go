package words

// Word 単語
type Word int

// LocalizeKey ローカライズ用のメッセージID
func (i Word) LocalizeKey() string {
	return i.String()
}

//String コードを文字列へ変更
func (i Word) String() string {
	return ""
}
