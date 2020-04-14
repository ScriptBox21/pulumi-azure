// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package policy

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a policy set definition.
//
// > **NOTE:**  Policy set definitions (also known as policy initiatives) do not take effect until they are assigned to a scope using a Policy Set Assignment.
type PolicySetDefinition struct {
	pulumi.CustomResourceState

	// The description of the policy set definition.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId pulumi.StringPtrOutput `pulumi:"managementGroupId"`
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata pulumi.StringOutput `pulumi:"metadata"`
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrOutput `pulumi:"parameters"`
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions .
	PolicyDefinitions pulumi.StringPtrOutput `pulumi:"policyDefinitions"`
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
}

// NewPolicySetDefinition registers a new resource with the given unique name, arguments, and options.
func NewPolicySetDefinition(ctx *pulumi.Context,
	name string, args *PolicySetDefinitionArgs, opts ...pulumi.ResourceOption) (*PolicySetDefinition, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.PolicyType == nil {
		return nil, errors.New("missing required argument 'PolicyType'")
	}
	if args == nil {
		args = &PolicySetDefinitionArgs{}
	}
	var resource PolicySetDefinition
	err := ctx.RegisterResource("azure:policy/policySetDefinition:PolicySetDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicySetDefinition gets an existing PolicySetDefinition resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicySetDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicySetDefinitionState, opts ...pulumi.ResourceOption) (*PolicySetDefinition, error) {
	var resource PolicySetDefinition
	err := ctx.ReadResource("azure:policy/policySetDefinition:PolicySetDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicySetDefinition resources.
type policySetDefinitionState struct {
	// The description of the policy set definition.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName *string `pulumi:"displayName"`
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions .
	PolicyDefinitions *string `pulumi:"policyDefinitions"`
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType *string `pulumi:"policyType"`
}

type PolicySetDefinitionState struct {
	// The description of the policy set definition.
	Description pulumi.StringPtrInput
	// The display name of the policy set definition.
	DisplayName pulumi.StringPtrInput
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId pulumi.StringPtrInput
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata pulumi.StringPtrInput
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions .
	PolicyDefinitions pulumi.StringPtrInput
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringPtrInput
}

func (PolicySetDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionState)(nil)).Elem()
}

type policySetDefinitionArgs struct {
	// The description of the policy set definition.
	Description *string `pulumi:"description"`
	// The display name of the policy set definition.
	DisplayName string `pulumi:"displayName"`
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId *string `pulumi:"managementGroupId"`
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata *string `pulumi:"metadata"`
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters *string `pulumi:"parameters"`
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions .
	PolicyDefinitions *string `pulumi:"policyDefinitions"`
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType string `pulumi:"policyType"`
}

// The set of arguments for constructing a PolicySetDefinition resource.
type PolicySetDefinitionArgs struct {
	// The description of the policy set definition.
	Description pulumi.StringPtrInput
	// The display name of the policy set definition.
	DisplayName pulumi.StringInput
	// The ID of the Management Group where this policy should be defined. Changing this forces a new resource to be created.
	ManagementGroupId pulumi.StringPtrInput
	// The metadata for the policy set definition. This is a json object representing additional metadata that should be stored with the policy definition.
	Metadata pulumi.StringPtrInput
	// The name of the policy set definition. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Parameters for the policy set definition. This field is a json object that allows you to parameterize your policy definition.
	Parameters pulumi.StringPtrInput
	// The policy definitions for the policy set definition. This is a json object representing the bundled policy definitions .
	PolicyDefinitions pulumi.StringPtrInput
	// The policy set type. Possible values are `BuiltIn` or `Custom`. Changing this forces a new resource to be created.
	PolicyType pulumi.StringInput
}

func (PolicySetDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policySetDefinitionArgs)(nil)).Elem()
}
