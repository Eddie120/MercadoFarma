package controllers

import searchInput "github.com/mercadofarma/services/services/search-input"

type SearchInputController struct {
	svc searchInput.SearchInputService
}

func NewSearchInputController(searchInputService searchInput.SearchInputService) *SearchInputController {
	return &SearchInputController{
		svc: searchInputService,
	}
}
