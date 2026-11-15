package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// NewClient connects using a Redis URL (redis://host:port or rediss:// for TLS,
// which managed providers like Upstash require) and verifies the connection.
func NewClient(url string) (*goredis.Client, error) {
	opts, err := goredis.ParseURL(url)
	if err != nil {
		// Deliberately drop the underlying error: url.Parse embeds the raw
		// input, and the input contains the password, so wrapping it would
		// leak the credential into logs.
		return nil, errors.New("parse redis url: invalid format (expected redis:// or rediss://user:pass@host:port, no surrounding quotes)")
	}
	rdb := goredis.NewClient(opts)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		rdb.Close()
		return nil, fmt.Errorf("redis ping %s: %w", opts.Addr, err)
	}
	return rdb, nil
}
