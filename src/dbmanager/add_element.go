package dbmanager

import "errors"

/* Placeholder for DB operation */
func AddElement(e Element) (int, error) {
	if len(e.Title) == 0 {
		return -1, errors.New("передан пустой заголовок")
	}
	if len(e.Text) == 0 {
		return -1, errors.New("передано пустое описание")
	}
	if len(e.Title) > MAX_TITLE_LEN {
		return -1, errors.New("длина заголовка превышает максимальное значение")
	}
	if len(e.Text) > MAX_TEXT_LEN {
		return -1, errors.New("длина текста превышает максимальное значение")
	}
	i := len(DB)
	DB = append(DB, e)
	return i, nil
}