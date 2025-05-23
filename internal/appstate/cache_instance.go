package appstate

import (
	"bootdev/gopoke/internal/pokecache"
	"context"
	"time"
)

func NewCache(ctx context.Context) *pokecache.PokeCache {
	return pokecache.NewCache(10*time.Second, ctx)
}
