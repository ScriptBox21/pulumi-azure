// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Authorization Rule for an Event Hub Namespace.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub_namespace_authorization_rule.html.markdown.
type EventHubNamespaceAuthorizationRule struct {
	pulumi.CustomResourceState

	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrOutput `pulumi:"listen"`
	Location pulumi.StringOutput `pulumi:"location"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrOutput `pulumi:"manage"`
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// The Primary Connection String for the Authorization Rule.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The Primary Key for the Authorization Rule.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Authorization Rule.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the Authorization Rule.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrOutput `pulumi:"send"`
}

// NewEventHubNamespaceAuthorizationRule registers a new resource with the given unique name, arguments, and options.
func NewEventHubNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, args *EventHubNamespaceAuthorizationRuleArgs, opts ...pulumi.ResourceOption) (*EventHubNamespaceAuthorizationRule, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventHubNamespaceAuthorizationRuleArgs{}
	}
	var resource EventHubNamespaceAuthorizationRule
	err := ctx.RegisterResource("azure:eventhub/eventHubNamespaceAuthorizationRule:EventHubNamespaceAuthorizationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHubNamespaceAuthorizationRule gets an existing EventHubNamespaceAuthorizationRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHubNamespaceAuthorizationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubNamespaceAuthorizationRuleState, opts ...pulumi.ResourceOption) (*EventHubNamespaceAuthorizationRule, error) {
	var resource EventHubNamespaceAuthorizationRule
	err := ctx.ReadResource("azure:eventhub/eventHubNamespaceAuthorizationRule:EventHubNamespaceAuthorizationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHubNamespaceAuthorizationRule resources.
type eventHubNamespaceAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	Location *string `pulumi:"location"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// The Primary Connection String for the Authorization Rule.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The Primary Key for the Authorization Rule.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Authorization Rule.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the Authorization Rule.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send *bool `pulumi:"send"`
}

type EventHubNamespaceAuthorizationRuleState struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrInput
	Location pulumi.StringPtrInput
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// The Primary Connection String for the Authorization Rule.
	PrimaryConnectionString pulumi.StringPtrInput
	// The Primary Key for the Authorization Rule.
	PrimaryKey pulumi.StringPtrInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Secondary Connection String for the Authorization Rule.
	SecondaryConnectionString pulumi.StringPtrInput
	// The Secondary Key for the Authorization Rule.
	SecondaryKey pulumi.StringPtrInput
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (EventHubNamespaceAuthorizationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubNamespaceAuthorizationRuleState)(nil)).Elem()
}

type eventHubNamespaceAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen *bool `pulumi:"listen"`
	Location *string `pulumi:"location"`
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage *bool `pulumi:"manage"`
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send *bool `pulumi:"send"`
}

// The set of arguments for constructing a EventHubNamespaceAuthorizationRule resource.
type EventHubNamespaceAuthorizationRuleArgs struct {
	// Grants listen access to this this Authorization Rule. Defaults to `false`.
	Listen pulumi.BoolPtrInput
	Location pulumi.StringPtrInput
	// Grants manage access to this this Authorization Rule. When this property is `true` - both `listen` and `send` must be too. Defaults to `false`.
	Manage pulumi.BoolPtrInput
	// Specifies the name of the Authorization Rule. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// The name of the resource group in which the EventHub Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Grants send access to this this Authorization Rule. Defaults to `false`.
	Send pulumi.BoolPtrInput
}

func (EventHubNamespaceAuthorizationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubNamespaceAuthorizationRuleArgs)(nil)).Elem()
}

