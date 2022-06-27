package interside

import "sync"

// Map container structure definition.
type Map struct {
	store sync.Map
}

// NewMap returns a new Map container instance.
func NewMap() *Map {
	return &Map{store: sync.Map{}}
}

// Append appends key value pair into Map instance.
func (custom *Map) Append(key string, value interface{}) *Map {
	custom.store.Store(key, value)

	return custom
}

// ToMap returns Map container value map[string]interface{}
func (custom *Map) ToMap() map[string]interface{} {
	var sets map[string]interface{}
	custom.store.Range(func(key, value any) bool {
		if sets == nil {
			sets = make(map[string]interface{})
		}

		sets[key.(string)] = value

		return true
	})

	return sets
}
