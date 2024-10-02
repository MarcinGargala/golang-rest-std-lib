package states

import (
	mvpProviders "rest-std-lib/mvp/providers"
	"rest-std-lib/mvp/providers/states"
	"testing"
)

func TestInMemoryAllFunction(t *testing.T) {
	inMemProvConf := mvpProviders.ProviderConfig{
		Name: "InMemoryTest",
		Type: "providers.state.memory",
	}
	inMemProv := InMemoryStateProvider{}
	_ = inMemProv.Init(inMemProvConf)

	// Upsert first component
	body1 := `{
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
	name1 := "soft-2"
	req1 := states.UpsertRequest{
		ID:   name1,
		Body: body1,
	}
	_ = inMemProv.Upsert(req1)

	// Get component by name
	req2 := states.GetRequest{
		ID: name1,
	}
	comp2, _ := inMemProv.Get(req2)

	if comp2.ID != name1 {
		t.Errorf("Not found component by name %v", name1)
	}

	// Upsert second component
	body3 := `{
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
	name3 := "soft-2"
	req3 := states.UpsertRequest{
		ID:   name3,
		Body: body3,
	}
	_ = inMemProv.Upsert(req3)

	// List components
	req4 := states.ListRequest{}
	comps4 := inMemProv.List(req4)

	if len(comps4) < 1 {
		t.Error("Not found all components")
	}
}
