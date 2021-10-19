package forms

import "github.com/shoriwe/metrolinea/internal/data/graph"

type ListForm struct {
	Cookies string
}

type ListTerminalsResponse struct {
	Terminals []string
}

type ListRoutesResponse struct {
	Routes map[string]graph.Route
}
