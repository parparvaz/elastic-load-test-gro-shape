package cli_router

import (
	"context"
)

var routes []map[string]func(ctx context.Context)

func AddRoute(route string, callback func(ctx context.Context)) {
	routes = append(routes, map[string]func(ctx context.Context){route: callback})
}

func Run(args string) {
	var (
		callback func(ctx context.Context)
		ok       bool
	)
	for _, route := range routes {
		callback, ok = route[args]
		if ok {
			break
		}
	}
	if callback == nil {
		return
	}
	callback(context.Background())
}
