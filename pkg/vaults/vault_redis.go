package vaults

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

// redis with ctx with timeout ?
var ctx = context.Background()

const (
	PONG = "PONG"
)

func GetAllEnvsFromRedis() string {
	uriVault, exist := os.LookupEnv("VAULT_URI")
	if !exist {
		log.Panicf("VAULT_URI exist: %v setted to %s", exist, uriVault)
	}
	vaulKeyFrontendEnvs := os.Getenv("VAULT_KEY_FRONTEND_ENVS_PROD")

	if os.Getenv("GO_ENV") == "dev" {
		vaulKeyFrontendEnvs = os.Getenv("VAULT_KEY_FRONTEND_ENVS_DEV")
	}
	if uriVault == "" || vaulKeyFrontendEnvs == "" {
		log.Panic("ERROR | Cannot load initial VAULT_URI")
	}
	opt, err := redis.ParseURL(uriVault)
	if err != nil {
		log.Panicf("ERROR | Cannot parse uri %s error: %v", uriVault, err)
	}
	redisClient := redis.NewClient(opt)
	defer redisClient.Close()

	err = pingRedis(ctx, redisClient)
	if err != nil {
		log.Panic("ERROR | Not possible to ping REDIS vault")
	}

	envsStr, err := getEnvsFromRedis(ctx, redisClient, vaulKeyFrontendEnvs)
	if err != nil {
		log.Panicf("ERROR | Cannot load VAULT_KEY_FRONTEND_ENVS %v DEV? %s", err, os.Getenv("GO_ENV"))
	}
	return envsStr
}

func pingRedis(ctx context.Context, client *redis.Client) error {
	status, err := client.Ping(ctx).Result()
	if err != nil || status != "PONG" {
		log.Panicf("ERROR | Not possible to ping REDIS vault error: %v", err)
	}
	return err
}

func getEnvsFromRedis(ctx context.Context, client *redis.Client, key string) (string, error) {
	envsStr, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Panicf("ERROR | Cannot load VAULT_KEY_FRONTEND_ENVS %v", err)
	}
	return envsStr, err
}
