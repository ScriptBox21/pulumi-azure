// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appservice

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
	case "azure:appservice/activeSlot:ActiveSlot":
		r, err = NewActiveSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/appService:AppService":
		r, err = NewAppService(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/certificate:Certificate":
		r, err = NewCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/certificateBinding:CertificateBinding":
		r, err = NewCertificateBinding(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/certificateOrder:CertificateOrder":
		r, err = NewCertificateOrder(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/customHostnameBinding:CustomHostnameBinding":
		r, err = NewCustomHostnameBinding(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/environment:Environment":
		r, err = NewEnvironment(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/functionApp:FunctionApp":
		r, err = NewFunctionApp(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/functionAppSlot:FunctionAppSlot":
		r, err = NewFunctionAppSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/hybridConnection:HybridConnection":
		r, err = NewHybridConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/managedCertificate:ManagedCertificate":
		r, err = NewManagedCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/plan:Plan":
		r, err = NewPlan(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/slot:Slot":
		r, err = NewSlot(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/slotVirtualNetworkSwiftConnection:SlotVirtualNetworkSwiftConnection":
		r, err = NewSlotVirtualNetworkSwiftConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/sourceCodeToken:SourceCodeToken":
		r, err = NewSourceCodeToken(ctx, name, nil, pulumi.URN_(urn))
	case "azure:appservice/virtualNetworkSwiftConnection:VirtualNetworkSwiftConnection":
		r, err = NewVirtualNetworkSwiftConnection(ctx, name, nil, pulumi.URN_(urn))
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
		"appservice/activeSlot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/appService",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/certificateBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/certificateOrder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/customHostnameBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/environment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/functionApp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/functionAppSlot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/hybridConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/managedCertificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/plan",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/slot",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/slotVirtualNetworkSwiftConnection",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/sourceCodeToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"azure",
		"appservice/virtualNetworkSwiftConnection",
		&module{version},
	)
}