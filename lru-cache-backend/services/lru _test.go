package services

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	var key = 1
	var value = "value"
	var newValue = "newValue"
	var duration = 1000 //17mins

	cache := NewLRUCache(3)
	cache.Set(key, value, time.Duration(duration)*time.Second)
	k, _, _ := cache.Get(1)
	if k != "value" {
		t.Fatalf("expected `value` but got `%s`", k)
	}

	cache.Set(key, newValue, time.Duration(duration)*time.Second)
	k1, _, _ := cache.Get(1)
	if k1 != "newValue" {
		t.Fatalf("expected `newValue` but got `%s`", k)
	}
}


func TestNewLRUCache(t *testing.T) {

	cache := NewLRUCache(3)

	cache.Set(1, "uk", time.Duration(150)*time.Second)
	cache.Set(2, "france", time.Duration(150)*time.Second)
	cache.Set(3, "germany", time.Duration(150)*time.Second)

	// order at this point
	// germany,france,uk

	cache.Set(4, "belgium", time.Duration(150)*time.Second)

	uk, _, _ := cache.Get(1)
	if uk != "" {
		t.Fatalf("uk should no loger exist in the cache")
	}

	// order
	// belgium,germany,france

	f, _, _ := cache.Get(2)
	if f != "france" {
		t.Fatalf("expectesd`france` but got `%s`", f)
	}

	//order is changed
	//france,belgium,germany

	//remove germany
	cache.Set(5, "netherlands", time.Duration(150)*time.Second)
	//order changed
	//netherlands,france,belgium

	n, _, _ := cache.Get(5)
	if n != "netherlands" {
		t.Fatalf("expectes `netherlands` but got `%s`", n)
	}

	fr, _, _ := cache.Get(2)
	if fr != "france" {
		t.Fatalf("expectes `france` but got `%s`", fr)
	}

	b, _, _ := cache.Get(4)
	if b != "belgium" {
		t.Fatalf("expectes `belgium` but got `%s`", b)
	}

}
