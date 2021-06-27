package routes

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes() Routes {
	return Routes{}
}

// Setup all the route

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
