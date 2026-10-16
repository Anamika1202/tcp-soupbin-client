package routes

import (
	"net/http"

	"tcp-soupbin-client/constants"

	"github.com/rs-commons/pkg/log"
	"github.com/rs-commons/pkg/web"
)

func RoutesList(logger *log.Logger) []web.RouteMapper {
	routes := []web.RouteMapper{
		{
			Uri:    constants.GetVersion.String(),
			Method: http.MethodPost,
			// Add the correct handler field here
			Logger: logger.GetLogger(constants.GetVersion.String()),
		},
	}

	return routes
}
