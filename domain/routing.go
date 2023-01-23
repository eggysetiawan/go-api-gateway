package domain

import (
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/dto"
	"github.com/eggysetiawan/go-api-gateway/errs"
)

type Routing struct {
	Id        int    `json:"routingId"`
	Uuid      string `json:"uuid"`
	Method    string `json:"method"`
	Uri       string `json:"uri"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type RoutingResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type IRoutingRepository interface {
	FindAllRoutings() ([]Routing, *errs.Exception)
	FindRoutingBy(uuid string) (*Routing, *errs.Exception)
	Save(request dto.RoutingRequest) *errs.Exception
	Update(request dto.RoutingRequest, uuid string) (*config.Response, *errs.Exception)
	Destroy(uuid string) (*config.Response, *errs.Exception)
}
