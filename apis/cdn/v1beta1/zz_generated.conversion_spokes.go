// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	ujconversion "github.com/crossplane/upjet/pkg/controller/conversion"
	"github.com/crossplane/upjet/pkg/resource"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this Endpoint to the hub type.
func (tr *Endpoint) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the Endpoint type.
func (tr *Endpoint) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this FrontdoorCustomDomain to the hub type.
func (tr *FrontdoorCustomDomain) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the FrontdoorCustomDomain type.
func (tr *FrontdoorCustomDomain) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this FrontdoorOrigin to the hub type.
func (tr *FrontdoorOrigin) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the FrontdoorOrigin type.
func (tr *FrontdoorOrigin) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this FrontdoorOriginGroup to the hub type.
func (tr *FrontdoorOriginGroup) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the FrontdoorOriginGroup type.
func (tr *FrontdoorOriginGroup) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this FrontdoorRoute to the hub type.
func (tr *FrontdoorRoute) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the FrontdoorRoute type.
func (tr *FrontdoorRoute) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this FrontdoorRule to the hub type.
func (tr *FrontdoorRule) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the FrontdoorRule type.
func (tr *FrontdoorRule) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}

// ConvertTo converts this FrontdoorSecurityPolicy to the hub type.
func (tr *FrontdoorSecurityPolicy) ConvertTo(dstRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := dstRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(dstRaw.(resource.Terraformed), tr); err != nil {
		return errors.Wrapf(err, "cannot convert from the spoke version %q to the hub version %q", spokeVersion, hubVersion)
	}
	return nil
}

// ConvertFrom converts from the hub type to the FrontdoorSecurityPolicy type.
func (tr *FrontdoorSecurityPolicy) ConvertFrom(srcRaw conversion.Hub) error {
	spokeVersion := tr.GetObjectKind().GroupVersionKind().Version
	hubVersion := srcRaw.GetObjectKind().GroupVersionKind().Version
	if err := ujconversion.RoundTrip(tr, srcRaw.(resource.Terraformed)); err != nil {
		return errors.Wrapf(err, "cannot convert from the hub version %q to the spoke version %q", hubVersion, spokeVersion)
	}
	return nil
}