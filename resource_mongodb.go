package main

import (
    "log"
    "fmt"
    "github.com/hashicorp/terraform/helper/schema"
    "github.com/qualialabs/composeio-go-api-2"
)


func resourceMongodbUser() *schema.Resource {
    return &schema.Resource{
        Create: resourceMongodbUserCreate,
        Read:   resourceMongodbUserRead,
        Update: resourceMongodbUserUpdate,
        Delete: resourceMongodbUserDelete,

        Schema: map[string]*schema.Schema{
            "account": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "deployment": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "db_name": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "db_user": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
            "db_password": &schema.Schema{
                Type:     schema.TypeString,
                Required: true,
            },
        },
    }
}


func resourceMongodbUserCreate(d *schema.ResourceData, meta interface{}) error {

    client := meta.(*composeio.Client)
    // create user
    account := d.Get("account").(string)
    deployment := d.Get("deployment").(string)
    db_name := d.Get("db_name").(string)
    db_user := d.Get("db_user").(string)
    db_password := d.Get("db_password").(string)
    log.Println("[DEBUG] Creating new user" + db_user +" on mongodb " + db_name + " on deployment " + deployment + " under account " + account)

    mongodb := &composeio.Mongodb{
      Account: account,
      Deployment: deployment,
      Name: db_name,
    }
    user := &composeio.User{
      Username: db_user,
      Password: db_password,
      ReadOnly: false,
    }
    err := client.CreateMongodbUser(mongodb, user)

    if err != nil {
      return fmt.Errorf("Failed to create mongodb user: %s", err.Error())
    }

    d.SetId(db_name)
    log.Printf("[INFO] record ID: %s", d.Id())
    
    return nil
}

func resourceMongodbUserRead(d *schema.ResourceData, meta interface{}) error {

  return nil
}

func resourceMongodbUserUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*composeio.Client)

    account := d.Get("account").(string)
    deployment := d.Get("deployment").(string)
    db_name := d.Get("db_name").(string)
    db_user := d.Get("db_user").(string)
    db_password := d.Get("db_password").(string)
    log.Println("[DEBUG] Updating user" + db_user +" on mongodb " + db_name + " on deployment " + deployment + " under account " + account)

    mongodb := &composeio.Mongodb{
      Account: account,
      Deployment: deployment,
      Name: db_name,
    }
    user := &composeio.User{
      Username: db_user,
      Password: db_password,
      ReadOnly: false,
    }
    err := client.UpdateMongodbUser(mongodb, user)

    if err != nil {
      return fmt.Errorf("Failed to update mongodb user: %s", err.Error())
    }

    d.SetId(db_name)
    log.Printf("[INFO] record ID: %s", d.Id())
    
    return nil
}

func resourceMongodbUserDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*composeio.Client)

    account := d.Get("account").(string)
    deployment := d.Get("deployment").(string)
    db_name := d.Get("db_name").(string)
    db_user := d.Get("db_user").(string)
    db_password := d.Get("db_password").(string)
    log.Println("[DEBUG] Deleting user" + db_user +" on mongodb " + db_name + " on deployment " + deployment + " under account " + account)

    mongodb := &composeio.Mongodb{
      Account: account,
      Deployment: deployment,
      Name: db_name,
    }
    user := &composeio.User{
      Username: db_user,
      Password: db_password,
      ReadOnly: false,
    }
    err := client.DeleteMongodbUser(mongodb, user)

    if err != nil {
      return fmt.Errorf("Failed to delete mongodb user: %s", err.Error())
    }

    return nil

}