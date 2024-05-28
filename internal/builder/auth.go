package builder

import (
	"github.com/Kevinmajesta/tes-capstone-1/internal/http/router"
	"github.com/Kevinmajesta/tes-capstone-1/pkg/route"
)

func BuildAuthPublicRoutes() []*route.Route {
	return router.PublicRoutes()
}

func BuildAuthPrivateRoutes() []*route.Route {
	return router.PrivateRoutes()
}
