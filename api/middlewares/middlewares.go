package middlewares

// IMiddleware middleware interface
type Middleware interface {
	Setup()
}

// Middlewares contains multiple middleware
type Middlewares []Middleware

// NewMiddlewares creates new middlewares
// Register the middleware that should be applied directly (globally)
func NewMiddlewares(
	corsMiddleware CorsMiddleware,
) Middlewares {
	return Middlewares{
		corsMiddleware,
	}
}

// Setup sets up middlewares
func (m Middlewares) Setup() {
	for _, middleware := range m {
		middleware.Setup()
	}
}
