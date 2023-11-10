package minio

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	EndPoint        string
	AccessKeyID     string
	SecretAccessKey string
}

func NewMinio(cfg *Config) *minio.Client {
	c, err := minio.New(cfg.EndPoint, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
	})
	if err != nil {
		log.Error("minio.NewMinio error: %v", err)
		panic(err)
	}
	return c
}
