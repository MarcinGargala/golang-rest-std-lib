package managers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"rest-std-lib/api/model"
	"rest-std-lib/mvp/managers"
	"rest-std-lib/mvp/providers"
	"rest-std-lib/mvp/providers/states"
)

type ComponentsManager struct {
	stateProvider states.IStateProvider
	Name          string
}

func (s *ComponentsManager) Init(config managers.ManagerConfig, providers map[string]providers.IProvider) error {
	s.Name = config.Name
	stateProviderName := config.Properties["providers.persistentstate"]
	provider, ok := providers[stateProviderName]
	if !ok {
		return errors.New("not found StateProvider name in props")
	}
	s.stateProvider, ok = provider.(states.IStateProvider)
	if !ok {
		return errors.New("can not convert founded provider to IStateProvider")
	}
	return nil
}

func (s *ComponentsManager) List() []model.Component {
	var request = states.ListRequest{
		Metadata: map[string]interface{}{},
	}
	results := s.stateProvider.List(request)
	var comps []model.Component
	for _, r := range results {
		var comp model.Component
		err := json.Unmarshal(r.Body.([]byte), &comp)
		if err != nil {
			slog.Warn("Error while unmarshal comp", "id", r.ID, "err", err)
			continue
		}
		comps = append(comps, comp)
	}
	return comps
}

func (s *ComponentsManager) Upsert(component model.Component) error {
	var req = states.UpsertRequest{
		ID:   component.Name,
		Body: component,
	}
	err := s.stateProvider.Upsert(req)
	return err
}

func (s *ComponentsManager) Get(name string) (model.Component, error) {
	var req = states.GetRequest{
		ID: name,
	}
	resp, err := s.stateProvider.Get(req)
	if err != nil {
		return model.Component{}, errors.New(fmt.Sprintf("No comp found by id %v", name))
	}
	var result = model.Component{}
	_ = json.Unmarshal(resp.Body.([]byte), &result)
	return result, nil
}
