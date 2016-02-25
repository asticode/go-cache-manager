package cachemanager

type Configuration struct {
	prefix string `json:"prefix"`
	ttl int32 `json:"ttl"`
}

type ConfigurationMemcache struct {
	Configuration
	servers string `json:"servers"`
}
