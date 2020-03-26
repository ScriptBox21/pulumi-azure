// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Group.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_group.html.markdown.
type Group struct {
	pulumi.CustomResourceState

	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The description of this API Management Group.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of this API Management Group.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId pulumi.StringPtrOutput `pulumi:"externalId"`
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewGroup registers a new resource with the given unique name, arguments, and options.
func NewGroup(ctx *pulumi.Context,
	name string, args *GroupArgs, opts ...pulumi.ResourceOption) (*Group, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GroupArgs{}
	}
	var resource Group
	err := ctx.RegisterResource("azure:apimanagement/group:Group", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure:apimanagement/group:Group", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Group resources.
type groupState struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The description of this API Management Group.
	Description *string `pulumi:"description"`
	// The display name of this API Management Group.
	DisplayName *string `pulumi:"displayName"`
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId *string `pulumi:"externalId"`
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type *string `pulumi:"type"`
}

type GroupState struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The description of this API Management Group.
	Description pulumi.StringPtrInput
	// The display name of this API Management Group.
	DisplayName pulumi.StringPtrInput
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId pulumi.StringPtrInput
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type pulumi.StringPtrInput
}

func (GroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupState)(nil)).Elem()
}

type groupArgs struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The description of this API Management Group.
	Description *string `pulumi:"description"`
	// The display name of this API Management Group.
	DisplayName string `pulumi:"displayName"`
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId *string `pulumi:"externalId"`
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Group resource.
type GroupArgs struct {
	// The name of the API Management Service in which the API Management Group should exist. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The description of this API Management Group.
	Description pulumi.StringPtrInput
	// The display name of this API Management Group.
	DisplayName pulumi.StringInput
	// The identifier of the external Group. For example, an Azure Active Directory group `aad://<tenant>.onmicrosoft.com/groups/<group object id>`.
	ExternalId pulumi.StringPtrInput
	// The name of the API Management Group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Group should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The type of this API Management Group. Possible values are `custom` and `external`. Default is `custom`.
	Type pulumi.StringPtrInput
}

func (GroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupArgs)(nil)).Elem()
}
