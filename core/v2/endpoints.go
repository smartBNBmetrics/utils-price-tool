package v2

import (
	"github.com/button-tech/utils-price-tool/pkg/storage/cache"
	routing "github.com/qiangxue/fasthttp-routing"
)

type Provider struct {
	Store *cache.Cache
}

type controller struct {
	store *cache.Cache
}

func API(g *routing.RouteGroup, p *Provider) {
	c := controller{
		store: p.Store,
	}

	g.Post("/prices", c.courses)
	g.Get("/info", c.info)
}
