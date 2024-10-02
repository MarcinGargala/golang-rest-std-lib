package vendors

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	apiManagers "rest-std-lib/api/managers"
	"rest-std-lib/api/model"
	"rest-std-lib/api/providers/states"
	mvpManagers "rest-std-lib/mvp/managers"
	mvpProviders "rest-std-lib/mvp/providers"
	"rest-std-lib/mvp/vendors"
	"strings"
	"testing"
)

func TestComponentCRUD(t *testing.T) {

	inMemProvConf := mvpProviders.ProviderConfig{
		Name: "InMemoryTest",
		Type: "providers.state.memory",
	}
	inMemProv := states.InMemoryStateProvider{}
	_ = inMemProv.Init(inMemProvConf)

	compManagerConfig := mvpManagers.ManagerConfig{
		Name: "ComponentsManagerTest",
		Type: "managers.symphony.components",
		Properties: map[string]string{
			"providers.persistentstate": "in-mem-test",
		},
	}
	compManager := apiManagers.ComponentsManager{}
	_ = compManager.Init(compManagerConfig, map[string]mvpProviders.IProvider{
		"in-mem-test": &inMemProv,
	})

	compVendConfig := vendors.VendorConfig{
		Type:  "vendors.components",
		Route: "/components",
	}
	compVendor := ComponentsVendor{}
	_ = compVendor.Init(compVendConfig, []mvpManagers.IManager{&compManager})

	// GET empty list of components
	request1 := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Path:   "/components",
		},
	}
	response1 := httptest.NewRecorder()
	compVendor.ServeHTTP(response1, request1)

	if response1.Code != 200 {
		t.Errorf("Wrong http response code %v", response1.Code)
	}
	var components1 []model.Component
	body, _ := io.ReadAll(response1.Result().Body)
	_ = json.Unmarshal(body, &components1)
	if len(components1) > 0 {
		t.Error("Wrong response body, component name ")
	}

	// CREATE new component
	var bodyStr = `{
		"name": "soft-2",
		"type": "component",
		"constrains": "$(bla...bla)",
		"metadata": {
		"prop1": "value2"
	},
		"properties": {
		"obj3": {
			"key3": "val2"
		}
	}
	}`
	request2 := &http.Request{
		Method: "POST",
		URL: &url.URL{
			Scheme: "http",
			Path:   "/components",
		},
		Body: io.NopCloser(strings.NewReader(bodyStr)),
	}
	response2 := &httptest.ResponseRecorder{}
	compVendor.ServeHTTP(response2, request2)

	if response2.Code != 200 {
		t.Errorf("Wrong http response code for POST %v", response2.Code)
	}

	// GET component by name
	var compName = "soft-2"
	request3 := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Path:   "/components/" + compName,
		},
	}
	response3 := httptest.NewRecorder()
	compVendor.ServeHTTP(response3, request3)

	if response3.Code != 200 {
		t.Errorf("Wrong http response code %v", response3.Code)
	}
	var components3 model.Component
	body3, _ := io.ReadAll(response3.Result().Body)
	_ = json.Unmarshal(body3, &components3)
	if components3.Name != compName {
		t.Errorf("Wrong response body no components found in with name %v", compName)
	}

	// CREATE second component
	var bodyStr2 = `{
		"name": "soft-3",
		"type": "component",
		"constrains": "$(bla...bla)",
		"metadata": {
		"prop1": "value3"
	},
		"properties": {
		"obj3": {
			"key3": "val3"
		}
	}
	}`
	request4 := &http.Request{
		Method: "POST",
		URL: &url.URL{
			Scheme: "http",
			Path:   "/components",
		},
		Body: io.NopCloser(strings.NewReader(bodyStr2)),
	}
	response4 := &httptest.ResponseRecorder{}
	compVendor.ServeHTTP(response4, request4)

	if response4.Code != 200 {
		t.Errorf("Wrong http response code for POST %v", response4.Code)
	}

	// GET list of components
	request5 := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "http",
			Path:   "/components",
		},
	}
	response5 := httptest.NewRecorder()
	compVendor.ServeHTTP(response5, request5)

	if response5.Code != 200 {
		t.Errorf("Wrong http response code %v", response5.Code)
	}
	var components5 []model.Component
	body5, _ := io.ReadAll(response5.Result().Body)
	_ = json.Unmarshal(body5, &components5)
	if len(components5) < 1 {
		t.Error("Wrong response body no components found in response")
	}

}
