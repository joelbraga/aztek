package api

import (
	"github.com/joelbraga/aztek/repository"
	"github.com/joelbraga/aztek/action"
)

type ApiOptions struct {
	Repository  repository.Repository
	ActionEvent *action.ActionEvent
}

type ApiResource struct {
	Repository  repository.Repository
	ActionEvent *action.ActionEvent
}

func NewApiResource(options *ApiOptions) *ApiResource {
	repo := options.Repository
	event := options.ActionEvent

	if repo == nil {
		panic("Repository is required")
	}

	return &ApiResource{
		repo,
		event,
	}
}
