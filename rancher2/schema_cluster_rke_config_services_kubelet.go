package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func clusterRKEConfigServicesKubeletExtraArgsArrayFields() map[string]*schema.Schema {
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

func clusterRKEConfigServicesKubeletFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"cluster_dns_server": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cluster_domain": {
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
			Description: "Extra Arguments that can be specified multiple times which are added to kubelet services",
			Elem: &schema.Resource{
				Schema: clusterRKEConfigServicesKubeletExtraArgsArrayFields(),
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
		"fail_swap_on": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"generate_serving_certificate": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"image": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"infra_container_image": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
	return s
}
