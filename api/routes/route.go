package routes

// Routes contains multiple routes
type Routes []IRoute

// Route interface
type IRoute interface {
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
