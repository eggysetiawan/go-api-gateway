package app

import (
	"encoding/json"
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/dto"
	"github.com/eggysetiawan/go-api-gateway/service"
	"github.com/gorilla/mux"
	"net/http"
)

type RoutingHandler struct {
	service service.IRoutingService
}

func (rh *RoutingHandler) indexRouting(w http.ResponseWriter, r *http.Request) {
	routings, err := rh.service.GetRoutings()

	if err != nil {
		response := config.NewUnexpectedResponse(err.Message)

		config.JsonResponse(w, response.Code, response)

		return
	}

	response := config.NewDefaultResponse()

	response.Data = routings

	config.JsonResponse(w, response.Code, response)

	return
}

func (rh *RoutingHandler) showRouting(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	routing, err := rh.service.GetRoutingBy(uuid)

	if err != nil {
		response := config.NewUnexpectedResponse(err.Message)

		response.Code = err.Code

		config.JsonResponse(w, response.Code, response)

		return
	}

	response := config.NewDefaultResponse()

	response.Data = routing

	config.JsonResponse(w, response.Code, response)

	return
}

func (rh *RoutingHandler) storeRouting(w http.ResponseWriter, r *http.Request) {
	var request dto.RoutingRequest

	errParse := json.NewDecoder(r.Body).Decode(&request)

	if errParse != nil {
		response := config.NewUnexpectedResponse("Error parsing request " + errParse.Error())

		config.JsonResponse(w, response.Code, response)

		return
	}

	err := rh.service.NewRouting(request)

	if err != nil {
		response := config.NewUnexpectedResponse(err.Message)

		config.JsonResponse(w, response.Code, response)

		return
	}

	response := config.NewDefaultResponse()

	config.JsonResponse(w, response.Code, response)

	return
}

func (rh *RoutingHandler) updateRouting(w http.ResponseWriter, r *http.Request) {
	var request dto.RoutingRequest

	uuid := mux.Vars(r)["uuid"]

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		response := config.NewUnexpectedResponse("Error while encoding request " + err.Error())

		config.JsonResponse(w, response.Code, response)

		return
	}

	resp, errApi := rh.service.UpdateRouting(request, uuid)

	if errApi != nil {
		config.JsonResponse(w, resp.Code, resp)

		return
	}

	config.JsonResponse(w, resp.Code, resp)

	return

}

func (rh *RoutingHandler) deleteRouting(w http.ResponseWriter, r *http.Request) {
	uuid := mux.Vars(r)["uuid"]

	resp, err := rh.service.DeleteRouting(uuid)

	if err != nil {
		response := config.NewUnexpectedResponse(err.Message)

		response.Code = err.Code

		config.JsonResponse(w, response.Code, response)

		return
	}

	config.JsonResponse(w, resp.Code, resp)

	return
}
