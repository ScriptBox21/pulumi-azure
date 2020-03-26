// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a ServiceBus Namespace authorization Rule within a ServiceBus.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/servicebus_namespace_authorization_rule.html.markdown.
type NamespaceAuthorizationRule struct {
	pulumi.CustomResourceState

	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Namespace authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The Primary Key for the ServiceBus Namespace authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Namespace authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the ServiceBus Namespace authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
}

// NewNamespaceAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *NamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NamespaceAuthorizationRuleArgs{}
	}
	var resource NamespaceAuthorizationRule
	err := ctx.RegisterResource("azure:eventhub/namespaceAuthorizationRule:NamespaceAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceAuthorizationRule gets an existing NamespaceAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*NamespaceAuthorizationRule, error) {
	var resource NamespaceAuthorizationRule
	err := ctx.ReadResource("azure:eventhub/namespaceAuthorizationRule:NamespaceAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceAuthorizationRule resources.
type namespaceAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The Primary Connection String for the ServiceBus Namespace authorization Rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The Primary Key for the ServiceBus Namespace authorization Rule.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the ServiceBus Namespace authorization Rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the ServiceBus Namespace authorization Rule.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send *bool `pulumi:"send"`
}

type NamespaceAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The Primary Connection String for the ServiceBus Namespace authorization Rule.
	PrimaryConnectionString pulumi.StringPtrInput
	// The Primary Key for the ServiceBus Namespace authorization Rule.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Connection String for the ServiceBus Namespace authorization Rule.
	SecondaryConnectionString pulumi.StringPtrInput
	// The Secondary Key for the ServiceBus Namespace authorization Rule.
	SecondaryKey pulumi.StringPtrInput
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (NamespaceAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleState)(nil)).Elem()
}

type namespaceAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send *bool `pulumi:"send"`
}

// The set of arguments for constructing a NamespaceAuthorizationRule resource.
type NamespaceAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrInput
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the ServiceBus Namespace Authorization Rule resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the ServiceBus Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the resource group in which the ServiceBus Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (NamespaceAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceAuthorizationRuleArgs)(nil)).Elem()
}
