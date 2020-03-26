// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package relay

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Relay Namespace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/relay_namespace.html.markdown.
type Namespace struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The Identifier for Azure Insights metrics.
	MetricId pulumi.StringOutput `pulumi:"metricId"`
	// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary connection string for the authorization rule `RootManageSharedAccessKey`.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The name of the resource group in which to create the Azure Relay Namespace.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// The name of the SKU to use. At this time the only supported value is `Standard`.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SkuName == nil {
		return nil, errors.New("missing required argument 'SkuName'")
	}
	if args == nil {
		args = &NamespaceArgs{}
	}
	var resource Namespace
	err := ctx.RegisterResource("azure:relay/namespace:Namespace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespace gets an existing Namespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceState, opts ...pulumi.ResourceOption) (*Namespace, error) {
	var resource Namespace
	err := ctx.ReadResource("azure:relay/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The Identifier for Azure Insights metrics.
	MetricId *string `pulumi:"metricId"`
	// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The primary connection string for the authorization rule `RootManageSharedAccessKey`.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group in which to create the Azure Relay Namespace.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The name of the SKU to use. At this time the only supported value is `Standard`.
	SkuName *string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type NamespaceState struct {
	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The Identifier for Azure Insights metrics.
	MetricId pulumi.StringPtrInput
	// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The primary connection string for the authorization rule `RootManageSharedAccessKey`.
	PrimaryConnectionString pulumi.StringPtrInput
	// The primary access key for the authorization rule `RootManageSharedAccessKey`.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group in which to create the Azure Relay Namespace.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary connection string for the authorization rule `RootManageSharedAccessKey`.
	SecondaryConnectionString pulumi.StringPtrInput
	// The secondary access key for the authorization rule `RootManageSharedAccessKey`.
	SecondaryKey pulumi.StringPtrInput
	// The name of the SKU to use. At this time the only supported value is `Standard`.
	SkuName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Azure Relay Namespace.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SKU to use. At this time the only supported value is `Standard`.
	SkuName string `pulumi:"skuName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Specifies the supported Azure location where the Azure Relay Namespace exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Azure Relay Namespace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Azure Relay Namespace.
	ResourceGroupName pulumi.StringInput
	// The name of the SKU to use. At this time the only supported value is `Standard`.
	SkuName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}
