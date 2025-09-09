package handlers

import (
	"fmt"
	"gproton-backend/grpcclient"
	"net/http"

	"github.com/labstack/echo/v4"
)

// CallServiceHandler handles /call endpoint
func CallServiceHandler(protoSet *grpcclient.ProtoSet) echo.HandlerFunc {
	return func(c echo.Context) error {
		svcName := c.QueryParam("service")
		methodName := c.QueryParam("method")
		// fmt.Println(svcName, methodName)
		if svcName == "" || methodName == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "service and method required"})
		}

		svc, method := protoSet.FindServiceAndMethod(svcName, methodName)
		if svc == nil || method == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "service or method not found"})
		}
		// fmt.Printf("Available methed is: %d\n", method)

		payload := make(map[string]interface{})
		if err := c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		fmt.Println("svc is in CallServiceHandler : %d\n", svc)
		fmt.Println("method is in CallServiceHandler : %d\n", method)

		resp, err := grpcclient.CallRPCJSON("stage-dispatch-core.u-id.ir:80", svc, method, payload, protoSet)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, resp)
	}
}
