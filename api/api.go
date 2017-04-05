package api

import "github.com/joelbraga/aztek/repository"

type ApiOptions struct {
	Repository repository.Repository
}

type ApiResource struct {
	Repository repository.Repository
}

func NewApiResource(options *ApiOptions) *ApiResource{
	repo := options.Repository
	if(repo==nil){
		panic("Repository is required")
	}

	return &ApiResource{
		repo,
	}
}
