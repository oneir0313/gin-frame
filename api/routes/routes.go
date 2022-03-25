package routes

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes(
	publicRoutes PublicRoutes,
	jobRoutes JobRoutes,
) Routes {
	return Routes{
		publicRoutes,
		jobRoutes,
	}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
