// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Google Identity Provider.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_identity_provider_google.html.markdown.
type IdentityProviderGoogle struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// Client Id for Google Sign-in.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client secret for Google Sign-in.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIdentityProviderGoogle registers a new resource with the given unique name, arguments, and options.
func NewIdentityProviderGoogle(ctx *pulumi.Context,
	name string, args *IdentityProviderGoogleArgs, opts ...pulumi.ResourceOption) (*IdentityProviderGoogle, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ClientId == nil {
		return nil, errors.New("missing required argument 'ClientId'")
	}
	if args == nil || args.ClientSecret == nil {
		return nil, errors.New("missing required argument 'ClientSecret'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IdentityProviderGoogleArgs{}
	}
	var resource IdentityProviderGoogle
	err := ctx.RegisterResource("azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIdentityProviderGoogle gets an existing IdentityProviderGoogle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIdentityProviderGoogle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IdentityProviderGoogleState, opts ...pulumi.ResourceOption) (*IdentityProviderGoogle, error) {
	var resource IdentityProviderGoogle
	err := ctx.ReadResource("azure:apimanagement/identityProviderGoogle:IdentityProviderGoogle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IdentityProviderGoogle resources.
type identityProviderGoogleState struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// Client Id for Google Sign-in.
	ClientId *string `pulumi:"clientId"`
	// Client secret for Google Sign-in.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IdentityProviderGoogleState struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// Client Id for Google Sign-in.
	ClientId pulumi.StringPtrInput
	// Client secret for Google Sign-in.
	ClientSecret pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IdentityProviderGoogleState) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderGoogleState)(nil)).Elem()
}

type identityProviderGoogleArgs struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// Client Id for Google Sign-in.
	ClientId string `pulumi:"clientId"`
	// Client secret for Google Sign-in.
	ClientSecret string `pulumi:"clientSecret"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IdentityProviderGoogle resource.
type IdentityProviderGoogleArgs struct {
	// The Name of the API Management Service where this Google Identity Provider should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// Client Id for Google Sign-in.
	ClientId pulumi.StringInput
	// Client secret for Google Sign-in.
	ClientSecret pulumi.StringInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IdentityProviderGoogleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*identityProviderGoogleArgs)(nil)).Elem()
}

