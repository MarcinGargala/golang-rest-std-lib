package vendors

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	apiManagers "rest-std-lib/api/managers"
	"rest-std-lib/api/model"
	mvnManagers "rest-std-lib/mvp/managers"
	"rest-std-lib/mvp/vendors"
	"strings"
)

type ComponentsVendor struct {
	componentManager *apiManagers.ComponentsManager
	pattern          string
}

func (s *ComponentsVendor) Init(config vendors.VendorConfig, managers []mvnManagers.IManager) error {

	for _, manager := range managers {
		if m, ok := manager.(*apiManagers.ComponentsManager); ok {
			s.componentManager = m
		}
	}
	s.pattern = config.Route
	return nil
}

func (s *ComponentsVendor) GetEndpoints() string {
	return s.pattern
}

func (s *ComponentsVendor) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		var pathVars = strings.Split(request.URL.Path, "/")
		slog.Info("Processing GET request..", "path var", pathVars)
		var responseJson []byte
		var err, err2 error
		if len(pathVars) > 2 {
			var name = pathVars[2]
			var component model.Component
			component, err2 = s.componentManager.Get(name)
			responseJson, err = json.Marshal(component)
		} else {
			components := s.componentManager.List()
			if len(components) < 1 {
				responseJson = []byte("[]")
			} else {
				responseJson, err = json.Marshal(components)
			}
		}
		response.Header().Set("Content-Type", "application/json")
		if err != nil || err2 != nil {
			_, _ = response.Write([]byte("{}"))
			return
		}
		_, _ = response.Write(responseJson)
	case http.MethodPost:
		slog.Info("Processing POST request..")
		body, err := io.ReadAll(request.Body)
		if err != nil {
			slog.Error("Error occurred when reading request body")
		}
		var comp = model.Component{}
		err = json.Unmarshal(body, &comp)
		err = s.componentManager.Upsert(comp)

		response.Header().Set("Content-Type", "application/json")
		_, _ = response.Write([]byte("{}"))
	}
}
