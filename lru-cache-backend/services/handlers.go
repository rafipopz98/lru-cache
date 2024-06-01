package services

import (
	"container/list"
	"sync"
	"time"
)

// represents a single item in the cache
type CacheItem struct {
	Key        int
	Value      string
	Expiration time.Time
}

// represents the cache with a specific capacity
type LRUCache struct {
	Capacity int
	Items    map[int]*list.Element
	Order    *list.List
	Mutex    sync.Mutex
}

// creates a new LRU cache
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Items:    make(map[int]*list.Element),
		Order:    list.New(),
	}
}

// retrieves an item from the cache
func (c *LRUCache) Get(key int) (string, time.Time, bool) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	//check the element exist or not by key
	element, exists := c.Items[key]
	if !exists {
		//if not exist return emptry string
		return "", time.Time{}, false
	}
	//else Check if the item has expired
	item := element.Value.(*CacheItem)
	//if expired delete it
	if time.Now().After(item.Expiration) {
		c.Order.Remove(element)
		delete(c.Items, key)
		return "", time.Time{}, false
	}
	//since it's been used move to front and return the value
	c.Order.MoveToFront(element)
	return item.Value, item.Expiration, true
}

// adds item to the cache
func (c *LRUCache) Set(key int, value string, duration time.Duration) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	//add expiry time with current time
	expiration := time.Now().Add(duration)

	// check if key exist, if exist then update the value and expiry time
	if element, exists := c.Items[key]; exists {
		c.Order.MoveToFront(element)
		item := element.Value.(*CacheItem)
		item.Value = value
		item.Expiration = expiration
		return
	}

	//if not we have to add the elem to front
	//before that, check the capacity
	if c.Order.Len() >= c.Capacity {
		//if max capacity is hit then remove the last item
		oldest := c.Order.Back()
		if oldest != nil {
			c.Order.Remove(oldest)
			delete(c.Items, oldest.Value.(*CacheItem).Key)
		}
	}
	//after deleiting add the new cachitem to front

	item := &CacheItem{
		Key:        key,
		Value:      value,
		Expiration: expiration,
	}
	element := c.Order.PushFront(item)
	c.Items[key] = element
}

// returns all non-expired cached items as a slice in insertion order
func (c *LRUCache) GetAll() []CacheItem {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	var cacheItems []CacheItem
	// traverse and keep moving while checking for expiry time
	// if expired then delete
	// else keep adding it to cacheItems[]
	for element := c.Order.Front(); element != nil; element = element.Next() {
		item := element.Value.(*CacheItem)
		if time.Now().Before(item.Expiration) {
			cacheItems = append(cacheItems, *item)
		} else {
			// Remove expired items from cache
			delete(c.Items, item.Key)
			c.Order.Remove(element)
		}
	}
	return cacheItems
}

func (c *LRUCache) DeleteOne(key int) bool {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	// check whether element exist
	// if it doesnot exist then return false and print not dound
	element, exist := c.Items[key]
	if !exist {
		return false
	}
	//if exist then delete
	c.Order.Remove(element)
	delete(c.Items, key)
	return true
}

var cacheInstance *LRUCache

func init() {
	cacheInstance = NewLRUCache(1024)
}
