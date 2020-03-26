// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package waf

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Azure Web Application Firewall Policy instance.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/web_application_firewall_policy.html.markdown.
type Policy struct {
	pulumi.CustomResourceState

	// One or more `customRule` blocks as defined below.
	CustomRules PolicyCustomRuleArrayOutput `pulumi:"customRules"`
	// Resource location. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `policySetting` block as defined below.
	PolicySettings PolicyPolicySettingsPtrOutput `pulumi:"policySettings"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PolicyArgs{}
	}
	var resource Policy
	err := ctx.RegisterResource("azure:waf/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("azure:waf/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// One or more `customRule` blocks as defined below.
	CustomRules []PolicyCustomRule `pulumi:"customRules"`
	// Resource location. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `policySetting` block as defined below.
	PolicySettings *PolicyPolicySettings `pulumi:"policySettings"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

type PolicyState struct {
	// One or more `customRule` blocks as defined below.
	CustomRules PolicyCustomRuleArrayInput
	// Resource location. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `policySetting` block as defined below.
	PolicySettings PolicyPolicySettingsPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// One or more `customRule` blocks as defined below.
	CustomRules []PolicyCustomRule `pulumi:"customRules"`
	// Resource location. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the policy. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `policySetting` block as defined below.
	PolicySettings *PolicyPolicySettings `pulumi:"policySettings"`
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// One or more `customRule` blocks as defined below.
	CustomRules PolicyCustomRuleArrayInput
	// Resource location. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the policy. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `policySetting` block as defined below.
	PolicySettings PolicyPolicySettingsPtrInput
	// The name of the resource group. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the Web Application Firewall Policy.
	Tags pulumi.StringMapInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}
