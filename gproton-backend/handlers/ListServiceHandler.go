package handlers

import (
	"gproton-backend/grpcclient"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListServiceHandler(protoSet *grpcclient.ProtoSet) echo.HandlerFunc {
	return func(c echo.Context) error {
		if protoSet == nil {
			return c.JSON(http.StatusServiceUnavailable, map[string]string{
				"error": "Server still starting up",
			})
		}

		services := []map[string]interface{}{}
		for _, s := range protoSet.Services {
			methods := []string{}
			for _, m := range s.Methods {
				methods = append(methods, m.Name)
			}
			services = append(services, map[string]interface{}{
				"name":    s.Name,
				"methods": methods,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"services": services,
		})
	}
}
