package api

import "github.com/jamiegyoung/runemarkers-go/internal/pageio"

type ButtonPage struct {
	Page        pageio.Page
	TilesString string
}

func (p ButtonPage) Data() map[string]interface{} {
	return map[string]interface{}{
		"Page":        &p.Page,
		"TilesString": p.TilesString,
	}
}
