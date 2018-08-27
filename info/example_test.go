package info_test

import (
	"github.com/skillcoder/go-common-handlers/info"
	"github.com/takama/router"
)

// ExampleHandler is a usage example for info.Handler
func ExampleHandler() {
	r := router.New()

	version := "1.0.0"
	repo := "go-common-handlers"
	commit := "356a192b7913b04c54574d18c28d46e6395428ab"

	r.GET("/info", info.Handler(version, repo, commit))
	r.Listen(":3000")
}
