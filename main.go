package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"rest-std-lib/api/managers"
	"rest-std-lib/api/providers"
	"rest-std-lib/api/vendors"
	"rest-std-lib/mvp/host"
)

func main() {
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer logFile.Close()
	logger := slog.New(slog.NewJSONHandler(logFile, nil))
	slog.SetDefault(logger)

	file, err := os.ReadFile("api-prod.json")
	if err != nil {
		panic("No profile found")
	}
	var profile = make(map[string]interface{})
	err = json.Unmarshal(file, &profile)
	if err != nil {
		panic("Can not read profile properly!")
	}
	apiJson, err := json.Marshal(profile["api"])
	var apiConfig = host.APIConfig{}
	err2 := json.Unmarshal(apiJson, &apiConfig)
	if err != nil || err2 != nil {
		panic("Can not extract API object from profile!")
	}
	bindingsJson, err := json.Marshal(profile["bindings"])
	var bindingsConfig = host.BindingConfig{}
	err2 = json.Unmarshal(bindingsJson, &bindingsConfig)
	if err != nil || err2 != nil {
		panic("Can not extract bindings from profile!")
	}
	var hostConfig = host.Config{
		API:      apiConfig,
		Bindings: bindingsConfig,
	}
	var hostAPI = host.APIHost{}
	err = hostAPI.Init(hostConfig, &vendors.VendorFactory{}, &managers.ManagerFactory{}, &providers.ProviderFactory{})
	if err != nil {
		panic("Can not Launch App error when init hosts!")
	}
	_ = hostAPI.Launch()
}
