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
                        "Cn": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
						"Description": &schema.Schema{
							    Type:     schema.TypeString,
								Optional: true,
						},
						"Setattr": &schema.Schema{
							    Type: schema.TypeArray,
								Optional: true,
								Elem: &schema.Schema{
									    Type: schema.TypeString,
								},
						},
						"Addattr": &schema.Schema{
							    Type: schema.TypeArray,
								Optional: true,
								Elem: &schema.Schema{
									    Type: schema.TypeString,
								},
						},
						"All": &schema.Schema{
							    Type: schema.TypeBool,
								Optional: true,
						},
						"Raw": &schema.Schema{
							    Type: schema.TypeBool,
								Optional: true,
						},
						"Rights": &schema.Schema{
							    Type: schema.TypeBool,
								Optional: true,
						},
						"NoMembers": &schema.Schema{
							    Type: schema.TypeBool,
								Optional: true,
						},
                },
        }
}

func resourceHostgroupCreate(d *schema.ResourceData, m interface{}) error {
        client := m.(c)
		var o &freeipa.HostgroupAddOptionalArgs
		var a &freeipa.HostgroupAddArgs
        
		if v, ok := d.GetOk("Description"); ok {
			    o.Description = v.(string)
		}

		if v, ok := d.GetOk("Setattr"); ok {
				o.Setattr = v.([]string)
	    }

        if v, ok := d.GetOk("Addattr") ok {
			    o.Addattr = v.([]string)
		}

		if v, ok := d.GetOk("All") ok {
		        o.All = v.(bool)
		}

		if v, ok := d.GetOk("Raw") ok {
			    o.Raw = v.(bool)
		}

		if v, ok := d.GetOk("NoMembers") ok {
			    o.NoMembers = v.(bool)
		}

		if v, ok := d.GetOk("Cn") ok {
			    a.Cn = v.(string)
		}

		res, e := client.HostgroupAdd(a, o)
		if e != nil {
			    log.Fatal(e)
		}
        
		d.SetId(d.Cn)
	    return nil
}

func resourceHostgroupRead(d *schema.ResourceData, m interface{}) error {
        client := m.(c)
        var a *HostgroupShowArgs
	    var o *HostgroupShowOptionalArgs

		if v, ok := d.GetOk("Cn") ok {
			    a.Cn = v.(string)
		}

		if v, ok := d.GetOk("All") ok {
		        o.All = v.(bool)
		}

		if v, ok := d.GetOk("Raw") ok {
			    o.Raw = v.(bool)
		}

		if v, ok := d.GetOk("NoMembers") ok {
			    o.NoMembers = v.(bool)
		}

		if v, ok := d.GetOk("Rights") ok {
			    o.Rights = v.(bool)
		}

		res, e := client.HostgroupRead(a, o)
		if e != nil {
			    log.Fatal(e)
		}
        
        d.SetId(res.Value)

        return nil 
}

func resourceHostgroupUpdate(d *schema.ResourceData, m interface{}) error {
        client := m.(c)
		var o &freeipa.HostgroupModOptionalArgs
		var a &freeipa.HostgroupModArgs
        
		if v, ok := d.GetOk("Description"); ok {
			    o.Description = v.(string)
		}

		if v, ok := d.GetOk("Setattr"); ok {
				o.Setattr = v.([]string)
	    }

        if v, ok := d.GetOk("Addattr") ok {
			    o.Addattr = v.([]string)
		}

		if v, ok := d.GetOk("Delattr") ok {
			    o.Delattr = v.([]string)
		}

		if v, ok := d.GetOk("Rights") ok {
			    o.Rights = v.(bool)
	    }

		if v, ok := d.GetOk("All") ok {
		        o.All = v.(bool)
		}

		if v, ok := d.GetOk("Raw") ok {
			    o.Raw = v.(bool)
		}

		if v, ok := d.GetOk("NoMembers") ok {
			    o.NoMembers = v.(bool)
		}

		if v, ok := d.GetOk("Cn") ok {
			    a.Cn = v.(string)
		}

		res, e := client.HostgroupAdd(a, o)
		if e != nil {
			    log.Fatal(e)
		}
		
		d.SetId(d.Cn)
        return nil
}

func resourceHostgroupDelete(d *schema.ResourceData, m interface{}) error {
        client := m.(c)
		var o &freeipa.HostgroupDelOptionalArgs
		var a &freeipa.HostgroupDelArgs
        
		if v, ok := d.GetOk("Continue") ok {
			    o.NoMembers = v.(bool)
		}

		if v, ok := d.GetOk("Cn") ok {
			    a.Cn = v.(string)
		}

		res, e := client.HostgroupAdd(a, o)
		if e != nil {
			    log.Fatal(e)
		}
		
		d.SetId(d.Cn)
	    
        return nil
}
