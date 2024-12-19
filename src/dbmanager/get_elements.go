package dbmanager

import (
	"errors"
	"strconv"
)

/* Placeholder for DB operations */

type BrowseViewData struct {
	Title     string
	Num       int
	ShortText string
	Link      string
}

type BrowseView struct {
	NotFirst  bool
	NotLast   bool
	NextLink  string
	PrevLink  string
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
	if page > 0 {
		res.NotFirst = true
		res.PrevLink = "/browse?page=" + strconv.Itoa(int(page))
	} else {
		res.NotFirst = false
	}
	if (page + 1) * uint(PAGE_LEN) < uint(len(DB)) {
		res.NotLast = true
		res.NextLink = "/browse?page=" + strconv.Itoa(int(page + 2))
	} else {
		res.NotLast = false
	}
	res.Page = int(page) + 1
	return res
}

func GetBrowseElements(ids []uint) BrowseView {
	var res BrowseView
	res.Data = make([]BrowseViewData, 0)

	for _, val := range ids {
		if val < uint(len(DB)) {
			res.Data = append(res.Data, BrowseViewData{ 
				Num: int(val + 1), 
				Title: DB[val].Title, 
				ShortText: DB[val].Text[:150],
			    Link: "/examine?num=" + strconv.Itoa(int(val)) })
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