package packs

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title,omitempty"`
	Paragraphs []string `json:"paragraphs,omitempty"`
	Options    []Option `json:"options,omitempty"`
}

type Option struct {
	Text    string `json:"text,omitempty"`
	Chapter string `json:"chapter,omitempty"`
}
