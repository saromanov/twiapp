package cache

// Cache provides interface for data for caching
type Cache interface {
	Put(key, value string)
	Get(key string)
}