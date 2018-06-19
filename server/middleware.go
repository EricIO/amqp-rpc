package server

/*
MiddlewareFunc represent a function that can be used as a middleware.

For example:

	func myMiddle(next HandlerFunc) HandlerFunc {

		// Preinitialization of middleware here.

		return func(ctx context.Context, rw *ResponseWriter d amqp.Delivery) {
			// Before handler execution here.

			// Execute the handler.
			next(ctx, rw, d)

			// After execution here.
		}
	}

	server := New("url")

	// Add middleware to specific handler.
	server.AddHandler("foobar", myMiddle(HandlerFunc))

	// Add middleware to all handlers on the server.
	server.AddMiddleware(myMiddle)
*/
type MiddlewareFunc func(next HandlerFunc) HandlerFunc

/*
MiddlewareChain will attatch all given middlewares to your HandlerFunc.
The middlewares will be executed in the same order as your input.

For example:

	server := New("url")

	server.AddHandler(
		"foobar",
		MiddlewareChain(
			myHandler,
			middlewareOne,
			middlewareTwo,
			middlewareThree,
		)
	)
*/
func MiddlewareChain(next HandlerFunc, m ...MiddlewareFunc) HandlerFunc {
	if len(m) == 0 {
		// The middleware chain is done. All middlewares have been applied.
		return next
	}

	// Nest the middlewares so that we attatch them in order.
	// The first middleware will have the second middleware applied, and so on.
	return m[0](MiddlewareChain(next, m[1:cap(m)]...))
}
