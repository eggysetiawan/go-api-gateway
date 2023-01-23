package dto

type RoutingRequest struct {
	Method string `json:"method"`
	Uri    string `json:"uri"`
	Name   string `json:"name"`
}
