package main

import (
  "log"
  "github.com/qualialabs/composeio-go-api-2"
)

type Config struct {
  ADMIN_MONGODB_URL string
  SSL_PEM_PATH string
}

func (c *Config) Client() (*composeio.Client, error) {
  client := composeio.NewClient(c.ADMIN_MONGODB_URL, c.SSL_PEM_PATH)

  log.Printf("[INFO] Composeio Client configured ")

  return client, nil
}