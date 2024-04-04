// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this SpringCloudService
func (mg *SpringCloudService) GetTerraformResourceType() string {
	return "azurerm_spring_cloud_service"
}

// GetConnectionDetailsMapping for this SpringCloudService
func (tr *SpringCloudService) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"config_server_git_setting[*].http_basic_auth[*].password": "spec.forProvider.configServerGitSetting[*].httpBasicAuth[*].passwordSecretRef", "config_server_git_setting[*].repository[*].http_basic_auth[*].password": "spec.forProvider.configServerGitSetting[*].repository[*].httpBasicAuth[*].passwordSecretRef", "config_server_git_setting[*].repository[*].ssh_auth[*].host_key": "spec.forProvider.configServerGitSetting[*].repository[*].sshAuth[*].hostKeySecretRef", "config_server_git_setting[*].repository[*].ssh_auth[*].private_key": "spec.forProvider.configServerGitSetting[*].repository[*].sshAuth[*].privateKeySecretRef", "config_server_git_setting[*].ssh_auth[*].host_key": "spec.forProvider.configServerGitSetting[*].sshAuth[*].hostKeySecretRef", "config_server_git_setting[*].ssh_auth[*].private_key": "spec.forProvider.configServerGitSetting[*].sshAuth[*].privateKeySecretRef", "container_registry[*].password": "spec.forProvider.containerRegistry[*].passwordSecretRef"}
}

// GetObservation of this SpringCloudService
func (tr *SpringCloudService) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SpringCloudService
func (tr *SpringCloudService) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SpringCloudService
func (tr *SpringCloudService) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SpringCloudService
func (tr *SpringCloudService) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SpringCloudService
func (tr *SpringCloudService) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this SpringCloudService
func (tr *SpringCloudService) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this SpringCloudService
func (tr *SpringCloudService) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this SpringCloudService using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SpringCloudService) LateInitialize(attrs []byte) (bool, error) {
	params := &SpringCloudServiceParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SpringCloudService) GetTerraformSchemaVersion() int {
	return 1
}