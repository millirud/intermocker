package entity

func NewRouteBuilder(method string, path string) *RouteBuilder {
	return &RouteBuilder{
		route: &Route{},
	}
}

type RouteBuilder struct {
	route *Route
}

type Route struct {
	Method string
	Path   string
}
