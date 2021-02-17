// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package authorization

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure:authorization/assignment:Assignment":
		r, err = NewAssignment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:authorization/roleDefinition:RoleDefinition":
		r, err = NewRoleDefinition(ctx, name, nil, pulumi.URN_(urn))
	case "azure:authorization/userAssignedIdentity:UserAssignedIdentity":
		r, err = NewUserAssignedIdentity(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure",
		"authorization/assignment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"authorization/roleDefinition",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"authorization/userAssignedIdentity",
		&module{version},
	)
}
