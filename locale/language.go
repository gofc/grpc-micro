package locale

import (
	"golang.org/x/text/language"
)

// AcceptLanguage アプリケーションの使用可能言語を保持する
var AcceptLanguage *Languages
var languageMatcher language.Matcher

// Languages 使用可能言語
type Languages struct {
	Default    string   `yaml:"default"`
	Availables []string `yaml:"availables"`
}

// NewLanguage アプリケーションで使用可能な言語をセットする
func NewLanguage(defaultLang string, availables ...string) *Languages {
	return &Languages{
		Default:    defaultLang,
		Availables: append(make([]string, 0), availables...),
	}
}

// SetGlobal 使用可能言語をグローバル変数にセット
func (a *Languages) SetGlobal() {
	AcceptLanguage = a

	tags := make([]language.Tag, 0)
	tags = append(tags, language.MustParse(AcceptLanguage.Default))
	for _, s := range AcceptLanguage.Availables {
		tags = append(tags, language.MustParse(s))
	}
	languageMatcher = language.NewMatcher(tags)
}

// ExtractLocale 指定されたAcceptLanguage(優先順位)から、使用出来る言語を返却する
func (a *Languages) ExtractLocale(acceptLanguage []string) string {
	tag, _ := language.MatchStrings(languageMatcher, acceptLanguage...)
	return tag.String()
}
