package util

import (
	"context"
	"fmt"
	"net/http"

	fx "go.uber.org/fx"
)

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("OnStart")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("OnStop")
			return nil
		},
	})
	return nil
}

func FxSvc() {
	fmt.Println("start")

	fx.New(fx.Provide(NewHTTPServer), fx.Invoke(func(*http.Server) {
		fmt.Println("invoke")
	})).Run()

	fmt.Println("exit")
}
