package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func clusterRKEConfigServicesSchedulerExtraArgsArrayFields() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"extra_arg": {
			Type:     schema.TypeList,
			Required: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"argument": {
						Required: true,
						Type:     schema.TypeString,
					},
					"values": {
						Type:     schema.TypeList,
						Required: true,
						Elem: &schema.Schema{
							Type: schema.TypeString,
						},
					},
				},
			},
		},
	}
}

func clusterRKEConfigServicesSchedulerFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"extra_args": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
		},
		"extra_args_array": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Extra Arguments that can be specified multiple times which are added to scheduler services",
			Elem: &schema.Resource{
				Schema: clusterRKEConfigServicesSchedulerExtraArgsArrayFields(),
			},
		},
		"extra_binds": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"extra_env": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
		},
		"image": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}
