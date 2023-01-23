package service

import (
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/domain"
	"github.com/eggysetiawan/go-api-gateway/dto"
	"github.com/eggysetiawan/go-api-gateway/errs"
)

type IRoutingService interface {
	GetRoutings() ([]domain.Routing, *errs.Exception)
	GetRoutingBy(uuid string) (*domain.Routing, *errs.Exception)
	NewRouting(request dto.RoutingRequest) *errs.Exception
	UpdateRouting(request dto.RoutingRequest, uuid string) (*config.Response, *errs.Exception)
	DeleteRouting(uuid string) (*config.Response, *errs.Exception)
}

type ApiRoutingService struct {
	repo domain.IRoutingRepository
}

func NewApiRoutingService(repository domain.IRoutingRepository) ApiRoutingService {
	return ApiRoutingService{repository}
}

func (s ApiRoutingService) GetRoutings() ([]domain.Routing, *errs.Exception) {
	routings, err := s.repo.FindAllRoutings()

	if err != nil {
		return nil, err
	}

	return routings, nil
}

func (s ApiRoutingService) GetRoutingBy(uuid string) (*domain.Routing, *errs.Exception) {
	routing, err := s.repo.FindRoutingBy(uuid)

	if err != nil {
		return nil, err
	}

	return routing, nil
}

func (s ApiRoutingService) NewRouting(request dto.RoutingRequest) *errs.Exception {
	err := s.repo.Save(request)

	if err != nil {
		return err
	}

	return nil
}

func (s ApiRoutingService) UpdateRouting(request dto.RoutingRequest, uuid string) (*config.Response, *errs.Exception) {
	response, err := s.repo.Update(request, uuid)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s ApiRoutingService) DeleteRouting(uuid string) (*config.Response, *errs.Exception) {
	response, err := s.repo.Destroy(uuid)

	if err != nil {
		return nil, err
	}

	return response, nil
}
