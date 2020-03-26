// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a policy rule definition on a management group or your provider subscription.
//
// Policy definitions do not take effect until they are assigned to a scope using a Policy Assignment.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/policy_definition.html.markdown.
type Definition struct {
	pulumi.CustomResourceState

	// The description of the policy definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId pulumi.StringPtrOutput `pulumi:"managementGroupId"`
	// The metadata for the policy definition. This
	// is a json object representing additional metadata that should be stored
	// with the policy definition.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The policy mode that allows you to specify which resource
	// types will be evaluated.  The value can be "All", "Indexed" or
	// "NotSpecified". Changing this resource forces a new resource to be
	// created.
	Mode pulumi.StringOutput `pulumi:"mode"`
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters for the policy definition. This field
	// is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// The policy rule for the policy definition. This
	// is a json object representing the rule that contains an if and
	// a then block.
	PolicyRule pulumi.StringPtrOutput `pulumi:"policyRule"`
	// The policy type.  The value can be "BuiltIn", "Custom"
	// or "NotSpecified". Changing this forces a new resource to be created.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
}

// NewDefinition registers a new resource with the given unique name, arguments, and options.
func NewDefinition(ctx *pulumi.Context,
	name string, args *DefinitionArgs, opts ...pulumi.ResourceOption) (*Definition, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Mode == nil {
		return nil, errors.New("missing required argument 'Mode'")
	}
	if args == nil || args.PolicyType == nil {
		return nil, errors.New("missing required argument 'PolicyType'")
	}
	if args == nil {
		args = &DefinitionArgs{}
	}
	var resource Definition
	err := ctx.RegisterResource("azure:policy/definition:Definition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefinition gets an existing Definition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefinitionState, opts ...pulumi.ResourceOption) (*Definition, error) {
	var resource Definition
	err := ctx.ReadResource("azure:policy/definition:Definition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Definition resources.
type definitionState struct {
	// The description of the policy definition.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The metadata for the policy definition. This
	// is a json object representing additional metadata that should be stored
	// with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The policy mode that allows you to specify which resource
	// types will be evaluated.  The value can be "All", "Indexed" or
	// "NotSpecified". Changing this resource forces a new resource to be
	// created.
	Mode *string `pulumi:"mode"`
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy definition. This field
	// is a json object that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// The policy rule for the policy definition. This
	// is a json object representing the rule that contains an if and
	// a then block.
	PolicyRule *string `pulumi:"policyRule"`
	// The policy type.  The value can be "BuiltIn", "Custom"
	// or "NotSpecified". Changing this forces a new resource to be created.
	PolicyType *string `pulumi:"policyType"`
}

type DefinitionState struct {
	// The description of the policy definition.
	Description pulumi.StringPtrInput
	// The display name of the policy definition.
	DisplayName pulumi.StringPtrInput
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId pulumi.StringPtrInput
	// The metadata for the policy definition. This
	// is a json object representing additional metadata that should be stored
	// with the policy definition.
	Metadata pulumi.StringPtrInput
	// The policy mode that allows you to specify which resource
	// types will be evaluated.  The value can be "All", "Indexed" or
	// "NotSpecified". Changing this resource forces a new resource to be
	// created.
	Mode pulumi.StringPtrInput
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy definition. This field
	// is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// The policy rule for the policy definition. This
	// is a json object representing the rule that contains an if and
	// a then block.
	PolicyRule pulumi.StringPtrInput
	// The policy type.  The value can be "BuiltIn", "Custom"
	// or "NotSpecified". Changing this forces a new resource to be created.
	PolicyType pulumi.StringPtrInput
}

func (DefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionState)(nil)).Elem()
}

type definitionArgs struct {
	// The description of the policy definition.
	Description *string `pulumi:"description"`
	// The display name of the policy definition.
	DisplayName string `pulumi:"displayName"`
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The metadata for the policy definition. This
	// is a json object representing additional metadata that should be stored
	// with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The policy mode that allows you to specify which resource
	// types will be evaluated.  The value can be "All", "Indexed" or
	// "NotSpecified". Changing this resource forces a new resource to be
	// created.
	Mode string `pulumi:"mode"`
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy definition. This field
	// is a json object that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// The policy rule for the policy definition. This
	// is a json object representing the rule that contains an if and
	// a then block.
	PolicyRule *string `pulumi:"policyRule"`
	// The policy type.  The value can be "BuiltIn", "Custom"
	// or "NotSpecified". Changing this forces a new resource to be created.
	PolicyType string `pulumi:"policyType"`
}

// The set of arguments for constructing a Definition resource.
type DefinitionArgs struct {
	// The description of the policy definition.
	Description pulumi.StringPtrInput
	// The display name of the policy definition.
	DisplayName pulumi.StringInput
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId pulumi.StringPtrInput
	// The metadata for the policy definition. This
	// is a json object representing additional metadata that should be stored
	// with the policy definition.
	Metadata pulumi.StringPtrInput
	// The policy mode that allows you to specify which resource
	// types will be evaluated.  The value can be "All", "Indexed" or
	// "NotSpecified". Changing this resource forces a new resource to be
	// created.
	Mode pulumi.StringInput
	// The name of the policy definition. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy definition. This field
	// is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// The policy rule for the policy definition. This
	// is a json object representing the rule that contains an if and
	// a then block.
	PolicyRule pulumi.StringPtrInput
	// The policy type.  The value can be "BuiltIn", "Custom"
	// or "NotSpecified". Changing this forces a new resource to be created.
	PolicyType pulumi.StringInput
}

func (DefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*definitionArgs)(nil)).Elem()
}
