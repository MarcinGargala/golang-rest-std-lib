package states

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"rest-std-lib/mvp/providers"
	"rest-std-lib/mvp/providers/states"
)

type InMemoryStateProvider struct {
	Name string
	data map[string]interface{}
}

func (s *InMemoryStateProvider) Init(config providers.ProviderConfig) error {
	s.Name = config.Name
	s.data = make(map[string]interface{})
	return nil
}

func (s *InMemoryStateProvider) Upsert(request states.UpsertRequest) error {
	slog.Info("Start Upsert func in InMemoryStateProvider")
	bodyJson, _ := json.Marshal(request.Body)
	if _, ok := s.data[request.ID]; ok {
		slog.Info("Entity already exist overriding", "id", request.ID)
	} else {
		slog.Info("Added new entity to data store", "id", request.ID)
	}
	s.data[request.ID] = bodyJson
	return nil
}

func (s *InMemoryStateProvider) List(request states.ListRequest) []states.StateEntry {
	slog.Info("Start List func in InMemoryStateProvider")
	var result []states.StateEntry
	for k, v := range s.data {
		result = append(result, states.StateEntry{
			ID:   k,
			Body: v,
		})
	}
	return result
}

func (s *InMemoryStateProvider) Get(request states.GetRequest) (states.StateEntry, error) {
	if v, ok := s.data[request.ID]; ok {
		return states.StateEntry{
			ID:   request.ID,
			Body: v,
		}, nil
	}
	return states.StateEntry{}, errors.New(fmt.Sprintf("no entity found by id %s", request.ID))
}
