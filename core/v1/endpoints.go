package v1

import (
	"github.com/button-tech/utils-price-tool/core/internal/respond"
	"github.com/button-tech/utils-price-tool/pkg/storage/cache"
	"github.com/button-tech/utils-price-tool/services"
	routing "github.com/qiangxue/fasthttp-routing"
)

type Provider struct {
	Store  *cache.Cache
	Prices *services.Prices
}

type controller struct {
	prices            *services.Prices
	store             *cache.Cache
	privateCurrencies map[string][]string
}

func API(g *routing.RouteGroup, p *Provider) {
	c := controller{
		store:             p.Store,
		prices:            p.Prices,
		privateCurrencies: privateCurrencies(),
	}

	g.Get("/info", c.info)
	g.Get("/docs/swagger.json", respond.SwaggerJSONHandler(v1))
	g.Post("/prices", c.courses)
	g.Post("/change", c.privatePrices)

}

func privateCurrencies() map[string][]string {
	return map[string][]string{
		"BTC":   {"0x0000000000000000000000000000000000000000", "Bitcoin"},
		"ETH":   {"0x000000000000000000000000000000000000003c", "Ethereum"},
		"ETC":   {"0x000000000000000000000000000000000000003d", "Ethereum Classic"},
		"BCH":   {"0x0000000000000000000000000000000000000091", "Bitcoin Cash"},
		"LTC":   {"0x0000000000000000000000000000000000000002", "Litecoin"},
		"XLM":   {"0x0000000000000000000000000000000000000094", "Stellar"},
		"WAVES": {"0x0000000000000000000000000000000000579bfc", "Waves"},
	}
}
