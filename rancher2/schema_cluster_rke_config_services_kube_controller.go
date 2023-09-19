package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func clusterRKEConfigServicesKubeControllerExtraArgsArrayFields() map[string]*schema.Schema {
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

func clusterRKEConfigServicesKubeControllerFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"cluster_cidr": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"extra_args": {
			Type:     schema.TypeMap,
			Optional: true,
			Computed: true,
		},
		"extra_args_array": {
			Type:        schema.TypeList,
			Optional:    true,
			MaxItems:    1,
			Description: "Extra Arguments that can be specified multiple times which are added to kube-controller services",
			Elem: &schema.Resource{
				Schema: clusterRKEConfigServicesKubeControllerExtraArgsArrayFields(),
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
		"service_cluster_ip_range": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}
