package bindings

import (
	"fmt"
	"log/slog"
	"net/http"
	mvpVendors "rest-std-lib/mvp/vendors"
)

type HttpBinding struct {
	server *http.Server
	mux    *http.ServeMux
}

func (s *HttpBinding) Init(config HttpBindingConfig) {
	addr := fmt.Sprintf(":%s", config.Port)
	s.mux = http.NewServeMux()
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}
}

func (s *HttpBinding) Launch(vendors []mvpVendors.IVendor) {
	for _, v := range vendors {
		s.mux.HandleFunc(v.GetEndpoints(), v.ServeHTTP)
		s.mux.HandleFunc(v.GetEndpoints()+"/*", v.ServeHTTP)
	}
	err := s.server.ListenAndServe()
	if err != nil {
		slog.Error("Error while running HTTP Server")
		panic("http server stops working")
	}
}
