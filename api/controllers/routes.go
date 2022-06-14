package controllers

import "github.com/Arka-cell/goassignment/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Shops routes
	s.Router.HandleFunc("/signup", middlewares.SetMiddlewareJSON(s.CreateShop)).Methods("POST")
	s.Router.HandleFunc("/shops", middlewares.SetMiddlewareJSON(s.GetShops)).Methods("GET")
	s.Router.HandleFunc("/shops/{id}", middlewares.SetMiddlewareJSON(s.GetShop)).Methods("GET")
	s.Router.HandleFunc("/shops/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateShop))).Methods("PUT")
	s.Router.HandleFunc("/shops/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteShop)).Methods("DELETE")

	//Products routes
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.CreateProduct)).Methods("POST")
	s.Router.HandleFunc("/products", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(s.GetProduct)).Methods("GET")
	s.Router.HandleFunc("/products/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProduct))).Methods("PUT")
	s.Router.HandleFunc("/products/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteProduct)).Methods("DELETE")
}
