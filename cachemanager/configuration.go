package cachemanager

type Configuration struct {
	Prefix string `json:"prefix"`
	TTL    int64  `json:"ttl"`
}

type ConfigurationMemcache struct {
	Configuration
	Servers string `json:"servers"`
}

type ConfigurationMemory struct {
	CleanupInterval int64 `json:"cleanup_interval"`
	Configuration
	MaxSize int64 `json:"max_size"`
}
