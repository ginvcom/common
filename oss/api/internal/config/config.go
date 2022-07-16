package config

import "github.com/zeromicro/go-zero/rest"

type OSS struct {
	Region    string
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

type Config struct {
	rest.RestConf
	AliyunOSS map[string]OSS
}
