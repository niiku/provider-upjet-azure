/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this BackupPolicyBlobStorage
func (mg *BackupPolicyBlobStorage) GetTerraformResourceType() string {
	return "azurerm_data_protection_backup_policy_blob_storage"
}

// GetConnectionDetailsMapping for this BackupPolicyBlobStorage
func (tr *BackupPolicyBlobStorage) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupPolicyBlobStorage
func (tr *BackupPolicyBlobStorage) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupPolicyBlobStorage
func (tr *BackupPolicyBlobStorage) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupPolicyBlobStorage
func (tr *BackupPolicyBlobStorage) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupPolicyBlobStorage
func (tr *BackupPolicyBlobStorage) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupPolicyBlobStorage
func (tr *BackupPolicyBlobStorage) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupPolicyBlobStorage using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupPolicyBlobStorage) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupPolicyBlobStorageParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupPolicyBlobStorage) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BackupPolicyDisk
func (mg *BackupPolicyDisk) GetTerraformResourceType() string {
	return "azurerm_data_protection_backup_policy_disk"
}

// GetConnectionDetailsMapping for this BackupPolicyDisk
func (tr *BackupPolicyDisk) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupPolicyDisk
func (tr *BackupPolicyDisk) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupPolicyDisk
func (tr *BackupPolicyDisk) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupPolicyDisk
func (tr *BackupPolicyDisk) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupPolicyDisk
func (tr *BackupPolicyDisk) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupPolicyDisk
func (tr *BackupPolicyDisk) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupPolicyDisk using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupPolicyDisk) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupPolicyDiskParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupPolicyDisk) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BackupPolicyPostgreSQL
func (mg *BackupPolicyPostgreSQL) GetTerraformResourceType() string {
	return "azurerm_data_protection_backup_policy_postgresql"
}

// GetConnectionDetailsMapping for this BackupPolicyPostgreSQL
func (tr *BackupPolicyPostgreSQL) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupPolicyPostgreSQL
func (tr *BackupPolicyPostgreSQL) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupPolicyPostgreSQL
func (tr *BackupPolicyPostgreSQL) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupPolicyPostgreSQL
func (tr *BackupPolicyPostgreSQL) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupPolicyPostgreSQL
func (tr *BackupPolicyPostgreSQL) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupPolicyPostgreSQL
func (tr *BackupPolicyPostgreSQL) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupPolicyPostgreSQL using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupPolicyPostgreSQL) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupPolicyPostgreSQLParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupPolicyPostgreSQL) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this BackupVault
func (mg *BackupVault) GetTerraformResourceType() string {
	return "azurerm_data_protection_backup_vault"
}

// GetConnectionDetailsMapping for this BackupVault
func (tr *BackupVault) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this BackupVault
func (tr *BackupVault) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this BackupVault
func (tr *BackupVault) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this BackupVault
func (tr *BackupVault) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this BackupVault
func (tr *BackupVault) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this BackupVault
func (tr *BackupVault) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this BackupVault using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *BackupVault) LateInitialize(attrs []byte) (bool, error) {
	params := &BackupVaultParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *BackupVault) GetTerraformSchemaVersion() int {
	return 0
}
