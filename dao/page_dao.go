package dao

import (
	"io/ioutil"
	"strings"

	"github.com/teimurjan/go-plain-server/models"
)

func SavePage(p *models.Page) error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile("pages/"+filename, p.Body, 0600)
}

func LoadPage(title string) (*models.Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile("pages/" + filename)
	if err != nil {
		return nil, err
	}
	return &models.Page{Title: title, Body: body}, nil
}

func LoadAllPages() []*models.Page {
	pagesFileNames := getFileNamesWithExt(".txt", "pages")
	pages := make([]*models.Page, len(pagesFileNames))
	for i, fileName := range pagesFileNames {
		title := strings.Replace(fileName, ".txt", "", 1)
		page, _ := LoadPage(title)
		pages[i] = page
	}
	return pages
}
