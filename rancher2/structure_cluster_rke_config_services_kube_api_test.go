package rancher2

import (
	"reflect"
	"testing"

	managementClient "github.com/rancher/rancher/pkg/client/generated/management/v3"
)

var (
	testClusterRKEConfigServicesKubeAPIAuditLogConfigConf                     *managementClient.AuditLogConfig
	testClusterRKEConfigServicesKubeAPIAuditLogConfigInterface                []interface{}
	testClusterRKEConfigServicesKubeAPIAuditLogConf                           *managementClient.AuditLog
	testClusterRKEConfigServicesKubeAPIAuditLogInterface                      []interface{}
	testClusterRKEConfigServicesKubeAPIEventRateLimitConf                     *managementClient.EventRateLimit
	testClusterRKEConfigServicesKubeAPIEventRateLimitInterface                []interface{}
	testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigConf            *managementClient.SecretsEncryptionConfig
	testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigInterface       []interface{}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationConf             map[string]interface{}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationInterface        []interface{}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsConf      []interface{}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsInterface []interface{}
	testClusterRKEConfigServicesKubeAPIConf                                   *managementClient.KubeAPIService
	testClusterRKEConfigServicesKubeAPIInterface                              []interface{}
)

func init() {
	testClusterRKEConfigServicesKubeAPIAuditLogConfigConf = &managementClient.AuditLogConfig{
		Format:    "format",
		MaxAge:    5,
		MaxBackup: 10,
		MaxSize:   100,
		Path:      "path",
		Policy: map[string]interface{}{
			"apiVersion": clusterRKEConfigServicesKubeAPIAuditLogConfigPolicyAPIDefault,
			"kind":       clusterRKEConfigServicesKubeAPIAuditLogConfigPolicyKindDefault,
		},
	}
	testClusterRKEConfigServicesKubeAPIAuditLogConfigInterface = []interface{}{
		map[string]interface{}{
			"format":     "format",
			"max_age":    5,
			"max_backup": 10,
			"max_size":   100,
			"path":       "path",
			"policy":     "apiVersion: " + clusterRKEConfigServicesKubeAPIAuditLogConfigPolicyAPIDefault + "\nkind: " + clusterRKEConfigServicesKubeAPIAuditLogConfigPolicyKindDefault + "\n",
		},
	}
	testClusterRKEConfigServicesKubeAPIAuditLogConf = &managementClient.AuditLog{
		Enabled:       true,
		Configuration: testClusterRKEConfigServicesKubeAPIAuditLogConfigConf,
	}
	testClusterRKEConfigServicesKubeAPIAuditLogInterface = []interface{}{
		map[string]interface{}{
			"enabled":       true,
			"configuration": testClusterRKEConfigServicesKubeAPIAuditLogConfigInterface,
		},
	}
	testClusterRKEConfigServicesKubeAPIEventRateLimitConf = &managementClient.EventRateLimit{
		Enabled: true,
		Configuration: map[string]interface{}{
			"apiVersion": clusterRKEConfigServicesKubeAPIEventRateLimitConfigAPIDefault,
			"kind":       clusterRKEConfigServicesKubeAPIEventRateLimitConfigKindDefault,
		},
	}
	testClusterRKEConfigServicesKubeAPIEventRateLimitInterface = []interface{}{
		map[string]interface{}{
			"enabled":       true,
			"configuration": "apiVersion: " + clusterRKEConfigServicesKubeAPIEventRateLimitConfigAPIDefault + "\nkind: " + clusterRKEConfigServicesKubeAPIEventRateLimitConfigKindDefault + "\n",
		},
	}
	testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigConf = &managementClient.SecretsEncryptionConfig{
		Enabled: true,
		CustomConfig: &managementClient.EncryptionConfiguration{
			APIVersion: clusterRKEConfigServicesKubeAPIEncryptionConfigAPIDefault,
			Kind:       clusterRKEConfigServicesKubeAPIEncryptionConfigKindDefault,
			Resources: []managementClient.ResourceConfiguration{
				{},
			},
		},
	}
	testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigInterface = []interface{}{
		map[string]interface{}{
			"enabled":       true,
			"custom_config": "apiVersion: " + clusterRKEConfigServicesKubeAPIEncryptionConfigAPIDefault + "\nkind: " + clusterRKEConfigServicesKubeAPIEncryptionConfigKindDefault + "\nresources:\n- {}\n",
		},
	}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsConf = []interface{}{
		map[string]interface{}{
			"name": "EventRateLimit",
			"path": "",
			"configuration": map[string]interface{}{
				"apiVersion": clusterRKEConfigServicesKubeAPIEventRateLimitConfigAPIDefault,
				"kind":       clusterRKEConfigServicesKubeAPIEventRateLimitConfigKindDefault,
				"limits": []interface{}{
					map[string]interface{}{},
				},
			},
		},
	}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsInterface = []interface{}{
		map[string]interface{}{
			"name":          "EventRateLimit",
			"path":          "",
			"configuration": "apiVersion: " + clusterRKEConfigServicesKubeAPIEventRateLimitConfigAPIDefault + "\nkind: " + clusterRKEConfigServicesKubeAPIEventRateLimitConfigKindDefault + "\nlimits:\n- {}\n",
		},
	}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationConf = map[string]interface{}{
		"apiVersion": clusterRKEConfigServicesKubeAPIAdmissionConfigurationAPIDefault,
		"kind":       clusterRKEConfigServicesKubeAPIAdmissionConfigurationKindDefault,
		"plugins":    testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsConf,
	}
	testClusterRKEConfigServicesKubeAPIAdmissionConfigurationInterface = []interface{}{
		map[string]interface{}{
			"api_version": clusterRKEConfigServicesKubeAPIAdmissionConfigurationAPIDefault,
			"kind":        clusterRKEConfigServicesKubeAPIAdmissionConfigurationKindDefault,
			"plugins":     testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsInterface,
		},
	}
	testClusterRKEConfigServicesKubeAPIConf = &managementClient.KubeAPIService{
		AlwaysPullImages:       true,
		AuditLog:               testClusterRKEConfigServicesKubeAPIAuditLogConf,
		EventRateLimit:         testClusterRKEConfigServicesKubeAPIEventRateLimitConf,
		AdmissionConfiguration: testClusterRKEConfigServicesKubeAPIAdmissionConfigurationConf,
		ExtraArgs: map[string]string{
			"arg_one": "one",
			"arg_two": "two",
		},
		ExtraArgsArray: map[string][]string{
			"arg1": {"v1"},
			"arg2": {"v2"},
		},
		ExtraBinds:              []string{"bind_one", "bind_two"},
		ExtraEnv:                []string{"env_one", "env_two"},
		Image:                   "image",
		PodSecurityPolicy:       true,
		SecretsEncryptionConfig: testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigConf,
		ServiceClusterIPRange:   "10.43.0.0/16",
		ServiceNodePortRange:    "30000-32000",
	}
	testClusterRKEConfigServicesKubeAPIInterface = []interface{}{
		map[string]interface{}{
			"always_pull_images":      true,
			"admission_configuration": testClusterRKEConfigServicesKubeAPIAdmissionConfigurationInterface,
			"audit_log":               testClusterRKEConfigServicesKubeAPIAuditLogInterface,
			"event_rate_limit":        testClusterRKEConfigServicesKubeAPIEventRateLimitInterface,
			"extra_args": map[string]interface{}{
				"arg_one": "one",
				"arg_two": "two",
			},
			"extra_args_array": []interface{}{
				map[string]interface{}{
					"extra_arg": []interface{}{
						map[string]interface{}{
							"argument": "arg1",
							"values":   []interface{}{"v1"},
						},
						map[string]interface{}{
							"argument": "arg2",
							"values":   []interface{}{"v2"},
						},
					},
				},
			},
			"extra_binds":               []interface{}{"bind_one", "bind_two"},
			"extra_env":                 []interface{}{"env_one", "env_two"},
			"image":                     "image",
			"pod_security_policy":       true,
			"secrets_encryption_config": testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigInterface,
			"service_cluster_ip_range":  "10.43.0.0/16",
			"service_node_port_range":   "30000-32000",
		},
	}
}

func TestFlattenClusterRKEConfigServicesKubeAPIAuditLogConfig(t *testing.T) {

	cases := []struct {
		Input          *managementClient.AuditLogConfig
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIAuditLogConfigConf,
			testClusterRKEConfigServicesKubeAPIAuditLogConfigInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigServicesKubeAPIAuditLogConfig(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterRKEConfigServicesKubeAPIAuditLog(t *testing.T) {

	cases := []struct {
		Input          *managementClient.AuditLog
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIAuditLogConf,
			testClusterRKEConfigServicesKubeAPIAuditLogInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigServicesKubeAPIAuditLog(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterRKEConfigServicesKubeAPIEventRateLimit(t *testing.T) {

	cases := []struct {
		Input          *managementClient.EventRateLimit
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIEventRateLimitConf,
			testClusterRKEConfigServicesKubeAPIEventRateLimitInterface,
		},
	}

	for _, tc := range cases {
		output := flattenClusterRKEConfigServicesKubeAPIEventRateLimit(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterRKEConfigServicesKubeAPISecretsEncryptionConfig(t *testing.T) {

	cases := []struct {
		Input          *managementClient.SecretsEncryptionConfig
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigConf,
			testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigServicesKubeAPISecretsEncryptionConfig(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterRKEConfigServicesKubeAPI(t *testing.T) {

	cases := []struct {
		Input          *managementClient.KubeAPIService
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIConf,
			testClusterRKEConfigServicesKubeAPIInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigServicesKubeAPI(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterRKEConfigServicesKubeAPIAdmissionConfigurationPlugins(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsConf,
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigServicesKubeAPIAdmissionConfigurationPlugins(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestFlattenClusterRKEConfigServicesKubeAPIAdmissionConfiguration(t *testing.T) {

	cases := []struct {
		Input          map[string]interface{}
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationConf,
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationInterface,
		},
	}

	for _, tc := range cases {
		output, err := flattenClusterRKEConfigServicesKubeAPIAdmissionConfiguration(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigServicesKubeAPIAuditLogConfig(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.AuditLogConfig
	}{
		{
			testClusterRKEConfigServicesKubeAPIAuditLogConfigInterface,
			testClusterRKEConfigServicesKubeAPIAuditLogConfigConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterRKEConfigServicesKubeAPIAuditLogConfig(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on expander: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigServicesKubeAPIAuditLog(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.AuditLog
	}{
		{
			testClusterRKEConfigServicesKubeAPIAuditLogInterface,
			testClusterRKEConfigServicesKubeAPIAuditLogConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterRKEConfigServicesKubeAPIAuditLog(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on expander: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigServicesKubeAPIEventRateLimit(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.EventRateLimit
	}{
		{
			testClusterRKEConfigServicesKubeAPIEventRateLimitInterface,
			testClusterRKEConfigServicesKubeAPIEventRateLimitConf,
		},
	}

	for _, tc := range cases {
		output := expandClusterRKEConfigServicesKubeAPIEventRateLimit(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigServicesKubeAPISecretsEncryptionConfig(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.SecretsEncryptionConfig
	}{
		{
			testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigInterface,
			testClusterRKEConfigServicesKubeAPISecretsEncryptionConfigConf,
		},
	}

	for _, tc := range cases {
		output := expandClusterRKEConfigServicesKubeAPISecretsEncryptionConfig(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigServicesKubeAPI(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput *managementClient.KubeAPIService
	}{
		{
			testClusterRKEConfigServicesKubeAPIInterface,
			testClusterRKEConfigServicesKubeAPIConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterRKEConfigServicesKubeAPI(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on expander: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigServicesKubeAPIAdmissionConfigurationPlugins(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput []interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsInterface,
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationPluginsConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterRKEConfigServicesKubeAPIAdmissionConfigurationPlugins(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandClusterRKEConfigServicesKubeAPIAdmissionConfiguration(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput map[string]interface{}
	}{
		{
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationInterface,
			testClusterRKEConfigServicesKubeAPIAdmissionConfigurationConf,
		},
	}

	for _, tc := range cases {
		output, err := expandClusterRKEConfigServicesKubeAPIAdmissionConfiguration(tc.Input)
		if err != nil {
			t.Fatalf("[ERROR] on flattener: %#v", err)
		}
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
