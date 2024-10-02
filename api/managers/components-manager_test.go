package managers

import (
	"rest-std-lib/api/model"
	"rest-std-lib/api/providers/states"
	apiManagers "rest-std-lib/mvp/managers"
	mvpProviders "rest-std-lib/mvp/providers"
	"testing"
)

func TestAllComponentsFunctions(t *testing.T) {
	inMemProvConf := mvpProviders.ProviderConfig{
		Name: "InMemoryTest",
		Type: "providers.state.memory",
	}
	inMemProv := states.InMemoryStateProvider{}
	_ = inMemProv.Init(inMemProvConf)

	compManagerConfig := apiManagers.ManagerConfig{
		Name: "ComponentsManagerTest",
		Type: "managers.symphony.components",
		Properties: map[string]string{
			"providers.persistentstate": "in-mem-test",
		},
	}
	compManager := ComponentsManager{}
	_ = compManager.Init(compManagerConfig, map[string]mvpProviders.IProvider{
		"in-mem-test": &inMemProv,
	})

	// List components - empty
	comps1 := compManager.List()

	if len(comps1) > 0 {
		t.Error("List should be empty but actual size greater than zero")
	}

	// Create first component
	comp2 := model.Component{
		Name:       "soft-1",
		Type:       "component",
		Constrains: "$(bla...bla)",
		Metadata: map[string]string{
			"prop1": "value1",
		},
		Properties: map[string]interface{}{
			"obj1": map[string]string{
				"key1": "val1",
			},
		},
	}
	err2 := compManager.Upsert(comp2)

	if err2 != nil {
		t.Error("Error while creating first component")
	}

	// Get component by name
	name3 := "soft-1"
	comp3, err3 := compManager.Get(name3)

	if err3 != nil {
		t.Errorf("Cannot find component by name %v", name3)
	}

	if comp3.Name != name3 {
		t.Error("Get component with wrong name")
	}

	// Create second component
	comp4 := model.Component{
		Name:       "soft-2",
		Type:       "component",
		Constrains: "$(bla...bla)",
		Metadata: map[string]string{
			"prop2": "value2",
		},
		Properties: map[string]interface{}{
			"obj2": map[string]string{
				"key2": "val2",
			},
		},
	}
	err4 := compManager.Upsert(comp4)

	if err4 != nil {
		t.Error("Error while creating second component")
	}

	// List all components

	comps5 := compManager.List()

	if len(comps5) != 2 {
		t.Errorf("Wrong number of components expected 2 get %v", len(comps5))
	}

}
