/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this DefaultSupportedIdPConfig
func (mg *DefaultSupportedIdPConfig) GetTerraformResourceType() string {
	return "google_identity_platform_default_supported_idp_config"
}

// GetConnectionDetailsMapping for this DefaultSupportedIdPConfig
func (tr *DefaultSupportedIdPConfig) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_id": "spec.forProvider.clientIdSecretRef", "client_secret": "spec.forProvider.clientSecretSecretRef"}
}

// GetObservation of this DefaultSupportedIdPConfig
func (tr *DefaultSupportedIdPConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DefaultSupportedIdPConfig
func (tr *DefaultSupportedIdPConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DefaultSupportedIdPConfig
func (tr *DefaultSupportedIdPConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DefaultSupportedIdPConfig
func (tr *DefaultSupportedIdPConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DefaultSupportedIdPConfig
func (tr *DefaultSupportedIdPConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this DefaultSupportedIdPConfig
func (tr *DefaultSupportedIdPConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this DefaultSupportedIdPConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DefaultSupportedIdPConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &DefaultSupportedIdPConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DefaultSupportedIdPConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this InboundSAMLConfig
func (mg *InboundSAMLConfig) GetTerraformResourceType() string {
	return "google_identity_platform_inbound_saml_config"
}

// GetConnectionDetailsMapping for this InboundSAMLConfig
func (tr *InboundSAMLConfig) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"idp_config[*].idp_certificates[*].x509_certificate": "spec.forProvider.idpConfig[*].idpCertificates[*].x509CertificateSecretRef"}
}

// GetObservation of this InboundSAMLConfig
func (tr *InboundSAMLConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this InboundSAMLConfig
func (tr *InboundSAMLConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this InboundSAMLConfig
func (tr *InboundSAMLConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this InboundSAMLConfig
func (tr *InboundSAMLConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this InboundSAMLConfig
func (tr *InboundSAMLConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this InboundSAMLConfig
func (tr *InboundSAMLConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this InboundSAMLConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *InboundSAMLConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &InboundSAMLConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *InboundSAMLConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this OAuthIdPConfig
func (mg *OAuthIdPConfig) GetTerraformResourceType() string {
	return "google_identity_platform_oauth_idp_config"
}

// GetConnectionDetailsMapping for this OAuthIdPConfig
func (tr *OAuthIdPConfig) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_id": "spec.forProvider.clientIdSecretRef", "client_secret": "spec.forProvider.clientSecretSecretRef"}
}

// GetObservation of this OAuthIdPConfig
func (tr *OAuthIdPConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this OAuthIdPConfig
func (tr *OAuthIdPConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this OAuthIdPConfig
func (tr *OAuthIdPConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this OAuthIdPConfig
func (tr *OAuthIdPConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this OAuthIdPConfig
func (tr *OAuthIdPConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this OAuthIdPConfig
func (tr *OAuthIdPConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this OAuthIdPConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *OAuthIdPConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &OAuthIdPConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *OAuthIdPConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ProjectDefaultConfig
func (mg *ProjectDefaultConfig) GetTerraformResourceType() string {
	return "google_identity_platform_project_default_config"
}

// GetConnectionDetailsMapping for this ProjectDefaultConfig
func (tr *ProjectDefaultConfig) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ProjectDefaultConfig
func (tr *ProjectDefaultConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ProjectDefaultConfig
func (tr *ProjectDefaultConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ProjectDefaultConfig
func (tr *ProjectDefaultConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ProjectDefaultConfig
func (tr *ProjectDefaultConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ProjectDefaultConfig
func (tr *ProjectDefaultConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this ProjectDefaultConfig
func (tr *ProjectDefaultConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this ProjectDefaultConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ProjectDefaultConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &ProjectDefaultConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ProjectDefaultConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Tenant
func (mg *Tenant) GetTerraformResourceType() string {
	return "google_identity_platform_tenant"
}

// GetConnectionDetailsMapping for this Tenant
func (tr *Tenant) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Tenant
func (tr *Tenant) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Tenant
func (tr *Tenant) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Tenant
func (tr *Tenant) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Tenant
func (tr *Tenant) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Tenant
func (tr *Tenant) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Tenant
func (tr *Tenant) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this Tenant using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Tenant) LateInitialize(attrs []byte) (bool, error) {
	params := &TenantParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Tenant) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this TenantDefaultSupportedIdPConfig
func (mg *TenantDefaultSupportedIdPConfig) GetTerraformResourceType() string {
	return "google_identity_platform_tenant_default_supported_idp_config"
}

// GetConnectionDetailsMapping for this TenantDefaultSupportedIdPConfig
func (tr *TenantDefaultSupportedIdPConfig) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_id": "spec.forProvider.clientIdSecretRef", "client_secret": "spec.forProvider.clientSecretSecretRef"}
}

// GetObservation of this TenantDefaultSupportedIdPConfig
func (tr *TenantDefaultSupportedIdPConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this TenantDefaultSupportedIdPConfig
func (tr *TenantDefaultSupportedIdPConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this TenantDefaultSupportedIdPConfig
func (tr *TenantDefaultSupportedIdPConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this TenantDefaultSupportedIdPConfig
func (tr *TenantDefaultSupportedIdPConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this TenantDefaultSupportedIdPConfig
func (tr *TenantDefaultSupportedIdPConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this TenantDefaultSupportedIdPConfig
func (tr *TenantDefaultSupportedIdPConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this TenantDefaultSupportedIdPConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *TenantDefaultSupportedIdPConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &TenantDefaultSupportedIdPConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *TenantDefaultSupportedIdPConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this TenantInboundSAMLConfig
func (mg *TenantInboundSAMLConfig) GetTerraformResourceType() string {
	return "google_identity_platform_tenant_inbound_saml_config"
}

// GetConnectionDetailsMapping for this TenantInboundSAMLConfig
func (tr *TenantInboundSAMLConfig) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"idp_config[*].idp_certificates[*].x509_certificate": "spec.forProvider.idpConfig[*].idpCertificates[*].x509CertificateSecretRef"}
}

// GetObservation of this TenantInboundSAMLConfig
func (tr *TenantInboundSAMLConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this TenantInboundSAMLConfig
func (tr *TenantInboundSAMLConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this TenantInboundSAMLConfig
func (tr *TenantInboundSAMLConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this TenantInboundSAMLConfig
func (tr *TenantInboundSAMLConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this TenantInboundSAMLConfig
func (tr *TenantInboundSAMLConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this TenantInboundSAMLConfig
func (tr *TenantInboundSAMLConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this TenantInboundSAMLConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *TenantInboundSAMLConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &TenantInboundSAMLConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *TenantInboundSAMLConfig) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this TenantOAuthIdPConfig
func (mg *TenantOAuthIdPConfig) GetTerraformResourceType() string {
	return "google_identity_platform_tenant_oauth_idp_config"
}

// GetConnectionDetailsMapping for this TenantOAuthIdPConfig
func (tr *TenantOAuthIdPConfig) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"client_id": "spec.forProvider.clientIdSecretRef", "client_secret": "spec.forProvider.clientSecretSecretRef"}
}

// GetObservation of this TenantOAuthIdPConfig
func (tr *TenantOAuthIdPConfig) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this TenantOAuthIdPConfig
func (tr *TenantOAuthIdPConfig) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this TenantOAuthIdPConfig
func (tr *TenantOAuthIdPConfig) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this TenantOAuthIdPConfig
func (tr *TenantOAuthIdPConfig) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this TenantOAuthIdPConfig
func (tr *TenantOAuthIdPConfig) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this TenantOAuthIdPConfig
func (tr *TenantOAuthIdPConfig) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// LateInitialize this TenantOAuthIdPConfig using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *TenantOAuthIdPConfig) LateInitialize(attrs []byte) (bool, error) {
	params := &TenantOAuthIdPConfigParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *TenantOAuthIdPConfig) GetTerraformSchemaVersion() int {
	return 0
}
