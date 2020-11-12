package util

import (
	"sync"
)

// -------------- Exported --------------- //

//Store - a thread-safe map with expiration
type Store struct {
	vals map[string]interface{}
	mux  *sync.RWMutex
}

//Get a value
func (s Store) Get(k string) interface{} {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.vals[k]
}

//GetValue - return the underlying object stored
func (s Store) GetValue(k string) interface{} {
	s.mux.RLock()
	defer s.mux.RUnlock()
	return s.vals[k]
}

//Set a value
func (s Store) Set(k string, v interface{}) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.vals[k] = v

}

//Delete a value
func (s Store) Delete(k string) {
	s.mux.Lock()
	defer s.mux.Unlock()
	delete(s.vals, k)

}

//NewStore constructs a store and starts the janitor if enabled
func NewStore() (s *Store) {
	s = &Store{vals: make(map[string]interface{}), mux: new(sync.RWMutex)}

	return s
}

// -------------- End Exported --------------- //

//GetKeys - Get all keys
func (s Store) GetKeys() (out []string) {
	s.mux.RLock()
	defer s.mux.RUnlock()
	out = make([]string, len(s.vals))
	var i int
	for k := range s.vals {
		out[i] = k
	}
	return out
}
