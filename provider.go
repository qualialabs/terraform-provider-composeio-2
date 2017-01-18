package main

import (
    "log"
    "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{

      Schema: map[string]*schema.Schema{
        "admin_mongodb_url": &schema.Schema{
          Type:        schema.TypeString,
          Optional:    true,
          DefaultFunc: schema.EnvDefaultFunc("ADMIN_MONGODB_URL", nil),
        },
        "ssl_pem_path": &schema.Schema{
          Type:        schema.TypeString,
          Optional:    true,
          DefaultFunc: schema.EnvDefaultFunc("SSL_PEM_PATH", nil),
        },
      },


      ResourcesMap: map[string]*schema.Resource{
          "composeio_mongodbuser": resourceMongodbUser(),
      },

      ConfigureFunc: providerConfigure,

    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

  config := Config{
    ADMIN_MONGODB_URL: d.Get("admin_mongodb_url").(string),
    SSL_PEM_PATH: d.Get("ssl_pem_path").(string),
  }

  log.Println("[INFO] Initializing Composeio client")
  return config.Client()
}