package main

import (
  "log"
  "composeio"
)

type Config struct {
  COMPOSEIO_TOKEN string
}

func (c *Config) Client() (*composeio.Client, error) {
  client := composeio.NewClient(c.COMPOSEIO_TOKEN)

  log.Printf("[INFO] Composeio Client configured ")

  return client, nil
}