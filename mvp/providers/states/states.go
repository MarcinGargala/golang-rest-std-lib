package states

import "rest-std-lib/mvp/providers"

type IStateProvider interface {
	Init(config providers.ProviderConfig) error
	Upsert(request UpsertRequest) error
	List(request ListRequest) []StateEntry
	Get(request GetRequest) (StateEntry, error)
}

type UpsertRequest struct {
	ID       string
	Body     interface{}
	Metadata map[string]interface{}
}

type ListRequest struct {
	Metadata map[string]interface{}
}

type GetRequest struct {
	ID string
}

type StateEntry struct {
	ID   string
	Body interface{}
}
