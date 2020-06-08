// Code generated by sysl DO NOT EDIT.
package simple

import (
	"net/http"

	"github.com/anz-bank/sysl-go/common"
	"github.com/anz-bank/sysl-go/core"
	"github.com/anz-bank/sysl-go/restlib"
	"github.com/anz-bank/sysl-go/validator"
	"github.com/anz-bank/sysl-template/gen/jsonplaceholder"
)

// Handler interface for simple
type Handler interface {
	GetHandler(w http.ResponseWriter, r *http.Request)
	GetFoobarListHandler(w http.ResponseWriter, r *http.Request)
	PostBobHandler(w http.ResponseWriter, r *http.Request)
}

// ServiceHandler for simple API
type ServiceHandler struct {
	genCallback                           core.RestGenCallback
	serviceInterface                      *ServiceInterface
	jsonplaceholderjsonplaceholderService jsonplaceholder.Service
}

// NewServiceHandler for simple
func NewServiceHandler(genCallback core.RestGenCallback, serviceInterface *ServiceInterface, jsonplaceholderjsonplaceholderService jsonplaceholder.Service) *ServiceHandler {
	return &ServiceHandler{genCallback, serviceInterface, jsonplaceholderjsonplaceholderService}
}

// GetHandler ...
func (s *ServiceHandler) GetHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.Get == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := GetClient{}

	welcome, err := s.serviceInterface.Get(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, welcome)
}

// GetFoobarListHandler ...
func (s *ServiceHandler) GetFoobarListHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.GetFoobarList == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req GetFoobarListRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := GetFoobarListClient{
		GetTodos: s.jsonplaceholderjsonplaceholderService.GetTodos,
	}

	todosresponse, err := s.serviceInterface.GetFoobarList(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, todosresponse)
}

// PostBobHandler ...
func (s *ServiceHandler) PostBobHandler(w http.ResponseWriter, r *http.Request) {
	if s.serviceInterface.PostBob == nil {
		common.HandleError(r.Context(), w, common.InternalError, "not implemented", nil, s.genCallback.MapError)
		return
	}

	ctx := common.RequestHeaderToContext(r.Context(), r.Header)
	ctx = common.RespHeaderAndStatusToContext(ctx, make(http.Header), http.StatusOK)
	var req PostBobRequest

	ctx, cancel := s.genCallback.DownstreamTimeoutContext(ctx)
	defer cancel()
	valErr := validator.Validate(&req)
	if valErr != nil {
		common.HandleError(ctx, w, common.BadRequestError, "Invalid request", valErr, s.genCallback.MapError)
		return
	}

	client := PostBobClient{}

	welcome, err := s.serviceInterface.PostBob(ctx, &req, client)
	if err != nil {
		common.HandleError(ctx, w, common.DownstreamUnexpectedResponseError, "Downstream failure", err, s.genCallback.MapError)
		return
	}

	headermap, httpstatus := common.RespHeaderAndStatusFromContext(ctx)
	restlib.SetHeaders(w, headermap)
	restlib.SendHTTPResponse(w, httpstatus, welcome)
}