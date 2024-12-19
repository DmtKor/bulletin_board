package dbmanager

import "errors"

/* Placeholder for DB operations */

type BrowseViewData struct {
	Title     string
	Num       int
	ShortText string
}

type BrowseView struct {
	Page int
	Data []BrowseViewData     
}

func GetBrowseElementsByPage(page uint) BrowseView {
	ids := make([]uint, PAGE_LEN)
	page -= 1
	for i := uint(0); i < PAGE_LEN; i++ {
		ids[i] = i + page * uint(PAGE_LEN)
	}
	res := GetBrowseElements(ids)
	res.Page = int(page) + 1
	return res
}

func GetBrowseElements(ids []uint) BrowseView {
	var res BrowseView
	res.Data = make([]BrowseViewData, 0)

	for _, val := range ids {
		if val <= uint(len(DB)) {
			res.Data = append(res.Data, BrowseViewData{ 
				Num: int(val + 1), 
				Title: DB[val].Title, 
				ShortText: DB[val].Text[:150] })
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