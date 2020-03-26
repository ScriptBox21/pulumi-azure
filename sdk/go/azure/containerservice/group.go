// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package containerservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages as an Azure Container Group instance.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/container_group.html.markdown.
type Group struct {
	pulumi.CustomResourceState

	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers GroupContainerArrayOutput `pulumi:"containers"`
	// A `diagnostics` block as documented below.
	Diagnostics GroupDiagnosticsPtrOutput `pulumi:"diagnostics"`
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel pulumi.StringPtrOutput `pulumi:"dnsNameLabel"`
	// The FQDN of the container group derived from `dnsNameLabel`.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// An `identity` block as defined below.
	Identity GroupIdentityOutput `pulumi:"identity"`
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials GroupImageRegistryCredentialArrayOutput `pulumi:"imageRegistryCredentials"`
	// The IP address allocated to the container group.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType pulumi.StringPtrOutput `pulumi:"ipAddressType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network profile ID for deploying to virtual network.
	NetworkProfileId pulumi.StringPtrOutput `pulumi:"networkProfileId"`
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringOutput `pulumi:"osType"`
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy pulumi.StringPtrOutput `pulumi:"restartPolicy"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.Containers == nil {
		return nil, errors.New("missing required argument 'Containers'")
	}
	if args == nil || args.OsType == nil {
		return nil, errors.New("missing required argument 'OsType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("azure:containerservice/group:Group", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroup gets an existing Group resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupState, opts ...pulumi.ResourceOption) (*Group, error) {
	var resource Group
	err := ctx.ReadResource("azure:containerservice/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers []GroupContainer `pulumi:"containers"`
	// A `diagnostics` block as documented below.
	Diagnostics *GroupDiagnostics `pulumi:"diagnostics"`
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel *string `pulumi:"dnsNameLabel"`
	// The FQDN of the container group derived from `dnsNameLabel`.
	Fqdn *string `pulumi:"fqdn"`
	// An `identity` block as defined below.
	Identity *GroupIdentity `pulumi:"identity"`
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials []GroupImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// The IP address allocated to the container group.
	IpAddress *string `pulumi:"ipAddress"`
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType *string `pulumi:"ipAddressType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Network profile ID for deploying to virtual network.
	NetworkProfileId *string `pulumi:"networkProfileId"`
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType *string `pulumi:"osType"`
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags map[string]string `pulumi:"tags"`
}

type GroupState struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers GroupContainerArrayInput
	// A `diagnostics` block as documented below.
	Diagnostics GroupDiagnosticsPtrInput
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel pulumi.StringPtrInput
	// The FQDN of the container group derived from `dnsNameLabel`.
	Fqdn pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity GroupIdentityPtrInput
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials GroupImageRegistryCredentialArrayInput
	// The IP address allocated to the container group.
	IpAddress pulumi.StringPtrInput
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Network profile ID for deploying to virtual network.
	NetworkProfileId pulumi.StringPtrInput
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringPtrInput
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers []GroupContainer `pulumi:"containers"`
	// A `diagnostics` block as documented below.
	Diagnostics *GroupDiagnostics `pulumi:"diagnostics"`
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel *string `pulumi:"dnsNameLabel"`
	// An `identity` block as defined below.
	Identity *GroupIdentity `pulumi:"identity"`
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials []GroupImageRegistryCredential `pulumi:"imageRegistryCredentials"`
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType *string `pulumi:"ipAddressType"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Network profile ID for deploying to virtual network.
	NetworkProfileId *string `pulumi:"networkProfileId"`
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType string `pulumi:"osType"`
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy *string `pulumi:"restartPolicy"`
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The definition of a container that is part of the group as documented in the `container` block below. Changing this forces a new resource to be created.
	Containers GroupContainerArrayInput
	// A `diagnostics` block as documented below.
	Diagnostics GroupDiagnosticsPtrInput
	// The DNS label/name for the container groups IP. Changing this forces a new resource to be created.
	DnsNameLabel pulumi.StringPtrInput
	// An `identity` block as defined below.
	Identity GroupIdentityPtrInput
	// A `imageRegistryCredential` block as documented below. Changing this forces a new resource to be created.
	ImageRegistryCredentials GroupImageRegistryCredentialArrayInput
	// Specifies the ip address type of the container. `Public` or `Private`. Changing this forces a new resource to be created. If set to `Private`, `networkProfileId` also needs to be set.
	IpAddressType pulumi.StringPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Container Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Network profile ID for deploying to virtual network.
	NetworkProfileId pulumi.StringPtrInput
	// The OS for the container group. Allowed values are `Linux` and `Windows`. Changing this forces a new resource to be created.
	OsType pulumi.StringInput
	// The name of the resource group in which to create the Container Group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Restart policy for the container group. Allowed values are `Always`, `Never`, `OnFailure`. Defaults to `Always`. Changing this forces a new resource to be created.
	RestartPolicy pulumi.StringPtrInput
	// A mapping of tags to assign to the resource. Changing this forces a new resource to be created.
	Tags pulumi.StringMapInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
