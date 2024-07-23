// package mw is a package that bundles a few middlewares that can be used to wrap http.Handlers.
package mw

import "net/http"

// Middleware is a type that wraps an http.Handler
type Middleware interface {
	Handle(http.Handler) http.HandlerFunc
}

// Chain is a function that chains multiple Middleware
func Chain(f http.Handler, mws ...Middleware) http.HandlerFunc {
	for _, m := range mws {
		f = m.Handle(f)
	}
	return f.ServeHTTP
}
