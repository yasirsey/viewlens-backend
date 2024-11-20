package services

import "viewlens/internal/repositories"

type ViewService struct {
	repo *repositories.ViewRepository
}

func NewViewService(repo *repositories.ViewRepository) *ViewService {
	return &ViewService{repo}
}

func (vs *ViewService) ListViews() ([]string, error) {
	return vs.repo.ListViews()
}

func (vs *ViewService) GetViewDetails(viewName string) ([]string, []map[string]interface{}, error) {
	return vs.repo.GetViewDetails(viewName)
}
