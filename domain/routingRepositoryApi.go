package domain

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/eggysetiawan/go-api-gateway/config"
	"github.com/eggysetiawan/go-api-gateway/dto"
	"github.com/eggysetiawan/go-api-gateway/errs"
	"net/http"
)

type DefaultRoutingRepositoryApi struct {
	client  *http.Client
	address string
}

func (a DefaultRoutingRepositoryApi) clientError(resp *http.Response) *errs.Exception {
	if resp.StatusCode >= 400 {
		var response config.Response

		err := json.NewDecoder(resp.Body).Decode(&response)

		if err != nil {
			return errs.NewUnexpectedException("Error while parsing error body " + err.Error())
		}

		return errs.NewException(response.Message, response.Code)
	}
	return nil
}

func (a DefaultRoutingRepositoryApi) FindRoutingBy(uuid string) (*Routing, *errs.Exception) {
	url := a.address + "/api/routings/" + uuid

	mtd := "GET"

	var routing Routing

	request, err := http.NewRequest(mtd, url, nil)

	if err != nil {
		return nil, errs.NewUnexpectedException("Failed to make request Routing API " + err.Error())
	}

	resp, respErr := a.client.Do(request)

	if respErr != nil {
		return nil, errs.NewUnexpectedException("Failed to fetch response Routing API " + err.Error())
	}

	defer resp.Body.Close()

	if clientError := a.clientError(resp); clientError != nil {
		return nil, clientError
	}

	err = json.NewDecoder(resp.Body).Decode(&routing)

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while parsing request " + err.Error())
	}

	return &routing, nil

}

func (a DefaultRoutingRepositoryApi) FindAllRoutings() ([]Routing, *errs.Exception) {
	url := a.address + "/api/routings"

	mtd := "GET"

	routings := make([]Routing, 0)

	request, err := http.NewRequest(mtd, url, nil)

	if err != nil {
		return nil, errs.NewUnexpectedException("Failed to make request Routing API " + err.Error())
	}

	resp, respErr := a.client.Do(request)

	if respErr != nil {
		return nil, errs.NewUnexpectedException("Failed to fetch response Routing API " + err.Error())
	}

	defer resp.Body.Close()

	if clientError := a.clientError(resp); clientError != nil {
		return nil, clientError
	}

	err = json.NewDecoder(resp.Body).Decode(&routings)

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while parsing request " + err.Error())
	}

	return routings, nil

}

func (a DefaultRoutingRepositoryApi) Save(request dto.RoutingRequest) *errs.Exception {
	mtd := "POST"

	url := a.address + "/api/routings"

	var payload *bytes.Buffer = nil

	payload = new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(request)

	if err != nil {
		return errs.NewUnexpectedException("Error while encoding payload " + err.Error())
	}

	newRequest, err := http.NewRequest(mtd, url, payload)

	if err != nil {
		return errs.NewUnexpectedException("Error while making API request " + err.Error())
	}

	response, err := a.client.Do(newRequest)

	defer response.Body.Close()

	if clientError := a.clientError(response); clientError != nil {
		return clientError
	}

	if err != nil {
		return errs.NewUnexpectedException("Error when sending request " + err.Error())
	}

	return nil

}

func (a DefaultRoutingRepositoryApi) Update(request dto.RoutingRequest, uuid string) (*config.Response, *errs.Exception) {
	mtd := "PUT"

	url := a.address + fmt.Sprintf("/api/routings/%v/edit", uuid)

	var payload *bytes.Buffer = nil

	payload = new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(request)

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while encoding payload " + err.Error())
	}

	newRequest, err := http.NewRequest(mtd, url, payload)

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while making API request " + err.Error())
	}

	response, err := a.client.Do(newRequest)

	defer response.Body.Close()

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while sending request to client " + err.Error())
	}

	if clientError := a.clientError(response); clientError != nil {
		return nil, clientError
	}

	var resp config.Response

	json.NewDecoder(response.Body).Decode(&resp)

	return &resp, nil
}

func (a DefaultRoutingRepositoryApi) Destroy(uuid string) (*config.Response, *errs.Exception) {
	url := a.address + fmt.Sprintf("/api/routings/%s/delete", uuid)

	mtd := "DELETE"

	request, err := http.NewRequest(mtd, url, nil)

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while requesting api " + err.Error())
	}

	resp, err := a.client.Do(request)

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while sending request api " + err.Error())
	}

	defer resp.Body.Close()

	var response config.Response

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return nil, errs.NewUnexpectedException("Error while parsing response body " + err.Error())
	}

	if clientError := a.clientError(resp); clientError != nil {
		return nil, clientError
	}

	return &response, nil

}

func NewRoutingRepositoryApi(client *http.Client, address string) *DefaultRoutingRepositoryApi {

	return &DefaultRoutingRepositoryApi{client, address}
}
