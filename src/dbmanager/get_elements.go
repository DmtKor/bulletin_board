package dbmanager

import "errors"

/* Placeholder for DB operations */

func GetElements(ids []uint) []Element {
	res := make([]Element, 0)
	for _, val := range ids {
		if (val <= uint(len(DB))) {
			res = append(res, DB[val])
		}
	}

	return res
}

func GetElement(id uint) (Element, error) {
	if (id <= uint(len(DB))) {
		return DB[id], nil
	}
	return Element{"", ""}, errors.New("неверный номер объявления")
}