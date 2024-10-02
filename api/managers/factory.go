package managers

import "rest-std-lib/mvp/managers"

type ManagerFactory struct {
}

func (s *ManagerFactory) Create(config managers.ManagerConfig) (manager managers.IManager, err error) {

	switch config.Type {
	case "managers.symphony.components":
		return &ComponentsManager{}, nil
	}

	return nil, nil
}
