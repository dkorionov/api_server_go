package store

import (
	"fmt"
	"strings"
	"testing"
)

// TestStore
func TestStore(t *testing.T, databaseURL string) (*Store, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			_, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")))
			if err != nil {
				t.Fatal()
			}
		}
		s.Close()
	}
}
