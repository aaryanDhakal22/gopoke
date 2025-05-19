package pokecache

import (
	"context"
	"testing"
	"time"
)

type Case struct {
	inp string
	act string
	exp string
}

func TestCreatePokeCache(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	myCache := NewCache(5, ctx)
	if myCache.interval != 5 {
		t.Error("Cache wasn't created")
	}
	cancel()
}
func TestAddPokeCache(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	myCache := NewCache(5, ctx)
	myCache.Add("Item1", []byte("This is the data"))
	if len(myCache.cache) != 1 {
		t.Error("Couldn't add to cache")
	}
	cancel()
}
func TestGetPokeCache(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	myCache := NewCache(5, ctx)
	byteData := []byte("This is the data")
	myCache.Add("Item1", byteData)
	if _, ok := myCache.Get("Item1"); !ok {
		t.Error("Unable to retrieve data")
	}
	cancel()
}
func TestDeletesAfterInterval(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	myCache := NewCache(1, ctx)

	myCache.Add("Item1", []byte("This is the data"))
	time.Sleep(1200 * time.Millisecond)
	_, ok1 := myCache.Get("Item1")

	myCache.Add("Item2", []byte("This is the data"))
	time.Sleep(1200 * time.Millisecond)
	_, ok2 := myCache.Get("Item2")

	if ok1 || ok2 {
		t.Error("Not properly being deleted")
	}
	cancel()
}
