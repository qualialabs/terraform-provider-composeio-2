package main

import (
    "log"
    "github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{

      Schema: map[string]*schema.Schema{
        "composeio_token": &schema.Schema{
          Type:        schema.TypeString,
          Optional:    true,
          DefaultFunc: schema.EnvDefaultFunc("COMPOSEIO_TOKEN", nil),
        },
      },


      ResourcesMap: map[string]*schema.Resource{
          "composeio_mongodb": resourceMongodb(),
          "composeio_mongodbuser": resourceMongodbUser(),
      },

      ConfigureFunc: providerConfigure,

    }
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

  config := Config{
    COMPOSEIO_TOKEN: d.Get("composeio_token").(string),
  }

  log.Println("[INFO] Initializing Composeio client")
  return config.Client()
}