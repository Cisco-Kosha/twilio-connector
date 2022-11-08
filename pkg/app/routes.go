package app

func (a *App) initializeRoutes() {
	var apiV1 = "/api/v1"

	// specification routes
	a.Router.HandleFunc(apiV1+"/specification/list", a.listConnectorSpecification).Methods("GET", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/specification/test", a.testConnectorSpecification).Methods("GET", "OPTIONS")

	a.Router.HandleFunc(apiV1+"/message", a.sendWhatsAppMessage).Methods("POST", "OPTIONS")
	a.Router.HandleFunc(apiV1+"/call", a.makePhoneCall).Methods("POST", "OPTIONS")
}
