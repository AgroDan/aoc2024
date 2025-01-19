package utils

import "sync"

// This will allow for similar behiavor to Python's functools.cache
// feature, allowing computation to be cached and not recomputed.
// Memoization baby.

// For this particular implementation, I'm going to use a Map from the
// sync standard library. After reading the documentation it seems that
// this is the best opportunity to do exactly this. Now I can create
// a function that, if accessed many many times, it will simply return
// the cache of a previous computation that is using a map that is
// optimized for write few/read many. There we go.

type Cache struct {
	store sync.Map
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Get(key string, computeFunc func() interface{}) interface{} {
	// load from the cache first
	if value, exists := c.store.Load(key); exists {
		return value
	}

	// otherwise, execute the function and store the result after returning it.
	value := computeFunc()
	c.store.Store(key, value)
	return value
}

// Now to use this, you would execute it like so:

// func add(x, y int) int {
// 	return x + y
// }
// cache := NewCache()
// a, b := 5, 10
//
// cacheKey := fmt.Sprintf("sum_%d_%d", a, b)
// result := cache.Get(cacheKey, func() interface{} {
// 	return add(a, b)
// }).(int)
//
// fmt.Printf("Cached result: %d\n", result)

// Note the type assertion at the end. This is because the function
// returns an interface{} type, which is the most general type in Go.
