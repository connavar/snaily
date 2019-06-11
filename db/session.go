package db

import (
	"sync"
	"time"
)

var GlobalSessionCache = buildSessionCache()

func buildSessionCache() *SessionCache {
	result := &SessionCache{
		store: make(map[string]*SessionRow),
	}
	go func(cache *SessionCache) {
		for {
			time.Sleep(1 * time.Minute)
			cache.m.Lock()
			for key, row := range cache.store {
				if row.Expired() {
					delete(cache.store, key)
				}
			}
			cache.m.Unlock()
		}
	}(result)
	return result
}

type SessionCache struct {
	store map[string]*SessionRow
	m     sync.Mutex
}

type SessionRow struct {
	Key        string
	UserID     string
	ValidUntil time.Time
}

func (s *SessionRow) Expired() bool {
	return !s.ValidUntil.After(time.Now())
}

func (s *SessionCache) SessionValid(key string) (bool, string) {
	s.m.Lock()
	defer s.m.Unlock()
	if r, ok := s.store[key]; ok {
		return !r.Expired(), r.UserID
	}
	return false, ""

}

func (s *SessionCache) SetSession(key, ID string, duration time.Duration) {
	s.m.Lock()
	defer s.m.Unlock()

	s.store[key] = &SessionRow{
		Key:        key,
		UserID:     ID,
		ValidUntil: time.Now().Add(duration),
	}
	return
}
