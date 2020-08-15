package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

//APIserver ...
type APIserver struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIserver) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Logger...
func (s *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

// Congigure Router
func (s *APIserver) configureRouter() {
	s.router.Handle("/",http.FileServer(http.Dir("./views/")))
	s.router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
}


