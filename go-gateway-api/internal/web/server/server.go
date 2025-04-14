package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/romeolacerda/payment-gateway/go-gateway-api/internal/service"
	"github.com/romeolacerda/payment-gateway/go-gateway-api/internal/web/handlers"
	"github.com/romeolacerda/payment-gateway/go-gateway-api/internal/web/middleware"
)

type Server struct {
	router         *chi.Mux
	server         *http.Server
	accountService *service.AccountService
	invoiceService *service.InvoiceService
	port           string
}

func NewServer(accountService *service.AccountService, invoiceService *service.InvoiceService, port string) *Server {
	return &Server{
		router:         chi.NewRouter(),
		accountService: accountService,
		invoiceService: invoiceService,
		port:           port,
	}
}

func (s *Server) ConfigureRoutes() {
	AccountHandler := handlers.NewAccountHandler(s.accountService)
	Invoicehandler := handlers.NewInvoiceHandler(s.invoiceService)
	AuthMiddleware := middleware.NewAuthMiddleware(s.accountService)

	s.router.Post("/accounts", AccountHandler.Create)
	s.router.Get("/accounts", AccountHandler.Get)

	s.router.Group((func(r chi.Router) {

		r.Use(AuthMiddleware.Authenticate)

		s.router.Post("/invoice", Invoicehandler.Create)
		s.router.Get("/invoice/{id}", Invoicehandler.GetById)
		s.router.Get("/invoice", Invoicehandler.ListByAccount)
	}))

}

func (s *Server) Start() error {
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}
	return s.server.ListenAndServe()
}
