// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package serverspec

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName      *string           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType    *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerDebug          *bool             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce          *bool             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError        *string           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars       map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars  []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	Command              *string           `mapstructure:"command" cty:"command" hcl:"command"`
	ExtraArguments       []string          `mapstructure:"extra_arguments" cty:"extra_arguments" hcl:"extra_arguments"`
	RakeEnvVars          []string          `mapstructure:"rake_env_vars" cty:"rake_env_vars" hcl:"rake_env_vars"`
	RakeFile             *string           `mapstructure:"rake_file" cty:"rake_file" hcl:"rake_file"`
	RakeTask             *string           `mapstructure:"rake_task" cty:"rake_task" hcl:"rake_task"`
	User                 *string           `mapstructure:"user" cty:"user" hcl:"user"`
	LocalPort            *string           `mapstructure:"local_port" cty:"local_port" hcl:"local_port"`
	SSHHostKeyFile       *string           `mapstructure:"ssh_host_key_file" cty:"ssh_host_key_file" hcl:"ssh_host_key_file"`
	SSHAuthorizedKeyFile *string           `mapstructure:"ssh_authorized_key_file" cty:"ssh_authorized_key_file" hcl:"ssh_authorized_key_file"`
	SFTPCmd              *string           `mapstructure:"sftp_command" cty:"sftp_command" hcl:"sftp_command"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"command":                    &hcldec.AttrSpec{Name: "command", Type: cty.String, Required: false},
		"extra_arguments":            &hcldec.AttrSpec{Name: "extra_arguments", Type: cty.List(cty.String), Required: false},
		"rake_env_vars":              &hcldec.AttrSpec{Name: "rake_env_vars", Type: cty.List(cty.String), Required: false},
		"rake_file":                  &hcldec.AttrSpec{Name: "rake_file", Type: cty.String, Required: false},
		"rake_task":                  &hcldec.AttrSpec{Name: "rake_task", Type: cty.String, Required: false},
		"user":                       &hcldec.AttrSpec{Name: "user", Type: cty.String, Required: false},
		"local_port":                 &hcldec.AttrSpec{Name: "local_port", Type: cty.String, Required: false},
		"ssh_host_key_file":          &hcldec.AttrSpec{Name: "ssh_host_key_file", Type: cty.String, Required: false},
		"ssh_authorized_key_file":    &hcldec.AttrSpec{Name: "ssh_authorized_key_file", Type: cty.String, Required: false},
		"sftp_command":               &hcldec.AttrSpec{Name: "sftp_command", Type: cty.String, Required: false},
	}
	return s
}
