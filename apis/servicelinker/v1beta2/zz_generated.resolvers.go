// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta2

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"

	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	apisresolver "github.com/upbound/provider-azure/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *SpringCloudConnection) ResolveReferences( // ResolveReferences of this SpringCloudConnection.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("appplatform.azure.upbound.io", "v1beta2", "SpringCloudJavaDeployment", "SpringCloudJavaDeploymentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.SpringCloudID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.SpringCloudIDRef,
			Selector:     mg.Spec.ForProvider.SpringCloudIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.SpringCloudID")
	}
	mg.Spec.ForProvider.SpringCloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.SpringCloudIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cosmosdb.azure.upbound.io", "v1beta2", "SQLDatabase", "SQLDatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.TargetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.TargetResourceIDRef,
			Selector:     mg.Spec.ForProvider.TargetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.TargetResourceID")
	}
	mg.Spec.ForProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.TargetResourceIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("appplatform.azure.upbound.io", "v1beta2", "SpringCloudJavaDeployment", "SpringCloudJavaDeploymentList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.SpringCloudID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.SpringCloudIDRef,
			Selector:     mg.Spec.InitProvider.SpringCloudIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.SpringCloudID")
	}
	mg.Spec.InitProvider.SpringCloudID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.SpringCloudIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("cosmosdb.azure.upbound.io", "v1beta2", "SQLDatabase", "SQLDatabaseList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.TargetResourceID),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.TargetResourceIDRef,
			Selector:     mg.Spec.InitProvider.TargetResourceIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.TargetResourceID")
	}
	mg.Spec.InitProvider.TargetResourceID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.TargetResourceIDRef = rsp.ResolvedReference

	return nil
}
