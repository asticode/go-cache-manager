package cachemanager

type Configuration struct {
	Prefix string `json:"prefix"`
	TTL    int32  `json:"ttl"`
}

type ConfigurationMemcache struct {
	Configuration
	Servers string `json:"servers"`
}
