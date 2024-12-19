package dbmanager

import "errors"

/* Placeholder for DB operation */
func AddElement(e Element) error {
	if len(e.Text) <= MAX_TEXT_LEN && len(e.Title) <= MAX_TITLE_LEN {
		DB = append(DB, e)
		return nil
	} else {
		return errors.New("длина заголовка или текста превышает максимальное значение")
	}

}