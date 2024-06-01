package services

// import (
// 	"container/list"
// 	"time"
// )

// // Cache interface defines the methods for the cache
// type Cache interface {
// 	Get(key string) string
// 	Set(key, value string, ttl int64)
// }

// // KeyValue is a struct to hold key-value pairs along with expiration time
// type KeyValue struct {
// 	key        string
// 	value      string
// 	expiration int64
// }

// // LRUCache is a struct to hold the cache data
// type LRUCache struct {
// 	m   map[string]*list.Element
// 	cap int
// 	l   *list.List
// }

// // NewLRU initializes a new LRUCache with a given capacity
// func NewLRU(cap int) *LRUCache {
// 	return &LRUCache{
// 		m:   make(map[string]*list.Element),
// 		cap: cap,
// 		l:   list.New(),
// 	}
// }

// // Get retrieves a value from the cache
// func (c *LRUCache) Get(key string) string {
// 	// Find the element by key
// 	el, ok := c.m[key]
// 	// If not present, return empty string
// 	if !ok {
// 		return ""
// 	}
// 	kv := el.Value.(*KeyValue)
// 	// Check if the item has expired
// 	if kv.expiration > 0 && time.Now().Unix() > kv.expiration {
// 		// Remove the expired item
// 		c.l.Remove(el)
// 		delete(c.m, key)
// 		return ""
// 	}
// 	// Move the element to the front of the list
// 	c.l.MoveToFront(el)
// 	// Return the value
// 	return kv.value
// }

// // Set adds a key-value pair to the cache with an expiration time (ttl in seconds)
// func (c *LRUCache) Set(key, value string, ttl int64) {
// 	var expiration int64
// 	if ttl > 0 {
// 		expiration = time.Now().Unix() + ttl
// 	}
// 	// Check if the key is already present
// 	if el, ok := c.m[key]; ok {
// 		// If present, update the value and expiration time, and move to front
// 		kv := el.Value.(*KeyValue)
// 		kv.value = value
// 		kv.expiration = expiration
// 		c.l.MoveToFront(el)
// 	} else {
// 		// If not present, add new key-value pair to the front
// 		newElem := &KeyValue{key, value, expiration}
// 		el := c.l.PushFront(newElem)
// 		c.m[key] = el

// 		
// 		if c.l.Len() > c.cap {
// 			backElem := c.l.Back()
// 			if backElem != nil {
// 				kv := backElem.Value.(*KeyValue)
// 				delete(c.m, kv.key)
// 				c.l.Remove(backElem)
// 			}
// 		}
// 	}
// }
