package rancher2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

//Schemas

func clusterRKEConfigServicesEtcdBackupConfigS3Fields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"access_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
		"bucket_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"custom_ca": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"endpoint": {
			Type:     schema.TypeString,
			Required: true,
		},
		"folder": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"region": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"secret_key": {
			Type:      schema.TypeString,
			Optional:  true,
			Sensitive: true,
		},
	}
	return s
}

func clusterRKEConfigServicesEtcdBackupConfigFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"interval_hours": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  12,
		},
		"retention": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  6,
		},
		"s3_backup_config": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Elem: &schema.Resource{
				Schema: clusterRKEConfigServicesEtcdBackupConfigS3Fields(),
			},
		},
		"safe_timestamp": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"timeout": {
			Type:     schema.TypeInt,
			Optional: true,
			Computed: true,
		},
	}
	return s
}

func clusterRKEConfigServicesEtcdExtraArgsArrayFields() map[string]*schema.Schema {
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

func clusterRKEConfigServicesEtcdFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"backup_config": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: clusterRKEConfigServicesEtcdBackupConfigFields(),
			},
		},
		"ca_cert": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cert": {
			Type:      schema.TypeString,
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},
		"creation": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"external_urls": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString,
			},
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
			Description: "Extra Arguments that can be specified multiple times which are added to etcd services",
			Elem: &schema.Resource{
				Schema: clusterRKEConfigServicesEtcdExtraArgsArrayFields(),
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
		"gid": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"image": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"key": {
			Type:      schema.TypeString,
			Optional:  true,
			Computed:  true,
			Sensitive: true,
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"retention": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"snapshot": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"uid": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
	}
	return s
}
