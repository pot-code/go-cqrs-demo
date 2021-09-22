package web

import (
	"net/http"

	"github.com/pot-code/gobit/pkg/api"
)

func NewEndpoint(hs *HttpServer) *api.Endpoint {
	handlers := hs.Handlers
	return &api.Endpoint{
		Prefix: "api",
		Groups: []*api.ApiGroup{
			{
				Prefix: "/order",
				Routes: []*api.Route{
					{Path: "", Method: http.MethodGet, Handler: handlers.HandleQueryById},
					{Path: "", Method: http.MethodPost, Handler: handlers.HandleCreateOrder},
					{Path: "/cancel/:id", Method: http.MethodPut, Handler: handlers.HandleCancelOrder},
					{Path: "/confirm/:id", Method: http.MethodPut, Handler: handlers.HandleConfirmOrder},
					{Path: "/:id", Method: http.MethodDelete, Handler: handlers.HandleDeleteOrder},
				},
			},
		},
	}
}
