package appstate

import (
	"bootdev/gopoke/internal/pokecache"
	"context"
	"time"
)

var GlobalCache *pokecache.PokeCache

func InitCache(ctx context.Context) {
	GlobalCache = pokecache.NewCache(5*time.Second, ctx)
}
