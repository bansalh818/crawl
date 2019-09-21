package utils

import "sync"

// Set the set of Strings
type Set struct {
	URL  map[string]bool
	lock sync.RWMutex
}

// Add adds a new url to the URL Set. Returns a pointer to the Set.
func (s *Set) Add(newURL string) *Set {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.URL == nil {
		s.URL = make(map[string]bool)
	}
	_, ok := s.URL[newURL]
	if !ok {
		s.URL[newURL] = true
	}
	return s
}

// Delete removes the string from the Set and returns Has(string)
func (s *Set) Delete(newurl string) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := s.URL[newurl]
	if ok {
		delete(s.URL, newurl)
	}
	return ok
}

// Has returns true if the Set contains the string
func (s *Set) Has(newurl string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, ok := s.URL[newurl]
	return ok
}

// // Strings returns the string(s) stored
// func (s *Set) Strings() []string {
// 	s.lock.RLock()
// 	defer s.lock.RUnlock()
// 	items := []string{}
// 	for i := range s.URL {
// 		items = append(items, i)
// 	}
// 	return items
// }

// Size returns the size of the set
func (s *Set) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.URL)
}
