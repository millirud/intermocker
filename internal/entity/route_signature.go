package entity

func NewRouteSignatureBuilder(method string, path string) *RouteSignatureBuilder {
	return &RouteSignatureBuilder{
		route: &RouteSignature{
			Method: method,
			Path:   path,
		},
	}
}

type RouteSignatureBuilder struct {
	route *RouteSignature
}

func (b *RouteSignatureBuilder) New() *RouteSignature {
	return b.route
}

type RouteSignature struct {
	Method string
	Path   string
}
