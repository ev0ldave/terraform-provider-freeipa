package main

import (
        "github.com/hashicorp/terraform/helper/schema"
)

func resourceHostgroup() *schema.Resource {
        return &schema.Resource{
                Create: resourceHostgroupCreate,
                Read:   resourceHostgroupRead,
                Update: resourceHostgroupUpdate,
                Delete: resourceHostgroupDelete,

                Schema: map[string]*schema.Schema{
                        "hostgroup": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func resourceHostgroupCreate(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceHostgroupRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceHostgroupUpdate(d *schema.ResourceData, m interface{}) error {
        return nil
}

func resourceHostgroupDelete(d *schema.ResourceData, m interface{}) error {
        return nil
}
