package pokecache

import (
	"testing"
	"time"
)

type Case struct {
	inp string
	act string
	exp string
}

func TestCreatePokeCache(t *testing.T) {
	myCache := NewCache(5)
	if myCache.interval == 5 {

		t.Log("Cache created successfully")
		return
	} else {
		t.Error("Cache wasn't created")
	}
}
func TestAddPokeCache(t *testing.T) {
	myCache := NewCache(5)
	myCache.Add("Item1", []byte("This is the data"))
	if len(myCache.cache) == 1 {
		return
	} else {
		t.Error("Couldn't add to cache")
	}
}
func TestGetPokeCache(t *testing.T) {
	myCache := NewCache(5)
	byteData := []byte("This is the data")
	myCache.Add("Item1", byteData)
	if _, ok := myCache.Get("Item1"); !ok {
		t.Error("Unable to retrieve data")
	}
}
func TestDeletesAfterInterval(t *testing.T) {
	myCache := NewCache(1)

	myCache.Add("Item1", []byte("This is the data"))
	time.Sleep(1200 * time.Millisecond)
	_, ok1 := myCache.Get("Item1")

	myCache.Add("Item2", []byte("This is the data"))
	time.Sleep(1200 * time.Millisecond)
	_, ok2 := myCache.Get("Item2")

	if !ok1 && !ok2 {
		return
	} else {
		t.Error("Not properly being deleted")
		return
	}
}
