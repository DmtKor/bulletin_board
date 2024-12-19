package dbmanager

const (
	MAX_TITLE_LEN int = 80
	MAX_TEXT_LEN  int = 300
	PAGE_LEN     uint = 10
)

type Element struct {
	Title string
	Text string
}

/* Database placeholder */
var DB []Element = make([]Element, 0)