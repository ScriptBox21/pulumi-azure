// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an API Management Product Policy
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_product_policy.html.markdown.
type ProductPolicy struct {
	pulumi.CustomResourceState

	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent pulumi.StringOutput `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrOutput `pulumi:"xmlLink"`
}

// NewProductPolicy registers a new resource with the given unique name, arguments, and options.
func NewProductPolicy(ctx *pulumi.Context,
	name string, args *ProductPolicyArgs, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ProductPolicyArgs{}
	}
	var resource ProductPolicy
	err := ctx.RegisterResource("azure:apimanagement/productPolicy:ProductPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProductPolicy gets an existing ProductPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProductPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductPolicyState, opts ...pulumi.ResourceOption) (*ProductPolicy, error) {
	var resource ProductPolicy
	err := ctx.ReadResource("azure:apimanagement/productPolicy:ProductPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProductPolicy resources.
type productPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId *string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

type ProductPolicyState struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringPtrInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The XML Content for this Policy.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ProductPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyState)(nil)).Elem()
}

type productPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId string `pulumi:"productId"`
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The XML Content for this Policy.
	XmlContent *string `pulumi:"xmlContent"`
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink *string `pulumi:"xmlLink"`
}

// The set of arguments for constructing a ProductPolicy resource.
type ProductPolicyArgs struct {
	// The name of the API Management Service. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
	ProductId pulumi.StringInput
	// The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The XML Content for this Policy.
	XmlContent pulumi.StringPtrInput
	// A link to a Policy XML Document, which must be publicly available.
	XmlLink pulumi.StringPtrInput
}

func (ProductPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productPolicyArgs)(nil)).Elem()
}

