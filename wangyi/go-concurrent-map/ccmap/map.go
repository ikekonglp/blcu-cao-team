package ccmap

import "github.com/csimplestring/go-concurrent-map/ccmap/key"

// Map defines the functions that a map should support
type Map interface {
	Put(k key.Key, val interface{}) bool
	Get(k key.Key) (interface{}, bool)
	Delete(k key.Key) bool
}
