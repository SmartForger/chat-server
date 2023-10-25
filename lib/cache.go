package lib

import (
	"fmt"
	"sync"
)

var store = sync.Map{}

func CSet(key string, value string) {
	store.Store(key, value)
}

func CGet(key string) string {
	data, found := store.Load(key)

	if found {
		return fmt.Sprintf("%v", data)
	}

	return ""
}
