// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package notificationhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Notification Hub Namespace.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/notification_hub_namespace.html.markdown.
type Namespace struct {
	pulumi.CustomResourceState

	// Is this Notification Hub Namespace enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The Azure Region in which this Notification Hub Namespace should be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
	NamespaceType pulumi.StringOutput `pulumi:"namespaceType"`
	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ServiceBus Endpoint for this Notification Hub Namespace.
	ServicebusEndpoint pulumi.StringOutput `pulumi:"servicebusEndpoint"`
	// ) A `sku` block as described below.
	Sku NamespaceSkuOutput `pulumi:"sku"`
	// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
	SkuName pulumi.StringOutput `pulumi:"skuName"`
}

// NewNamespace registers a new resource with the given unique name, arguments, and options.
func NewNamespace(ctx *pulumi.Context,
	name string, args *NamespaceArgs, opts ...pulumi.ResourceOption) (*Namespace, error) {
	if args == nil || args.NamespaceType == nil {
		return nil, errors.New("missing required argument 'NamespaceType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NamespaceArgs{}
	}
	var resource Namespace
	err := ctx.RegisterResource("azure:notificationhub/namespace:Namespace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:notificationhub/namespace:Namespace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Namespace resources.
type namespaceState struct {
	// Is this Notification Hub Namespace enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Azure Region in which this Notification Hub Namespace should be created.
	Location *string `pulumi:"location"`
	// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
	NamespaceType *string `pulumi:"namespaceType"`
	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ServiceBus Endpoint for this Notification Hub Namespace.
	ServicebusEndpoint *string `pulumi:"servicebusEndpoint"`
	// ) A `sku` block as described below.
	Sku *NamespaceSku `pulumi:"sku"`
	// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
	SkuName *string `pulumi:"skuName"`
}

type NamespaceState struct {
	// Is this Notification Hub Namespace enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Azure Region in which this Notification Hub Namespace should be created.
	Location pulumi.StringPtrInput
	// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
	NamespaceType pulumi.StringPtrInput
	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ServiceBus Endpoint for this Notification Hub Namespace.
	ServicebusEndpoint pulumi.StringPtrInput
	// ) A `sku` block as described below.
	Sku NamespaceSkuPtrInput
	// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
	SkuName pulumi.StringPtrInput
}

func (NamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceState)(nil)).Elem()
}

type namespaceArgs struct {
	// Is this Notification Hub Namespace enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The Azure Region in which this Notification Hub Namespace should be created.
	Location *string `pulumi:"location"`
	// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
	NamespaceType string `pulumi:"namespaceType"`
	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// ) A `sku` block as described below.
	Sku *NamespaceSku `pulumi:"sku"`
	// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
	SkuName *string `pulumi:"skuName"`
}

// The set of arguments for constructing a Namespace resource.
type NamespaceArgs struct {
	// Is this Notification Hub Namespace enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The Azure Region in which this Notification Hub Namespace should be created.
	Location pulumi.StringPtrInput
	// The name to use for this Notification Hub Namespace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The Type of Namespace - possible values are `Messaging` or `NotificationHub`. Changing this forces a new resource to be created.
	NamespaceType pulumi.StringInput
	// The name of the Resource Group in which the Notification Hub Namespace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// ) A `sku` block as described below.
	Sku NamespaceSkuPtrInput
	// The name of the SKU to use for this Notification Hub Namespace. Possible values are `Free`, `Basic` or `Standard`. Changing this forces a new resource to be created.
	SkuName pulumi.StringPtrInput
}

func (NamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceArgs)(nil)).Elem()
}

