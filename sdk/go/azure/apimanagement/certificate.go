// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package apimanagement

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Certificate within an API Management Service.
type Certificate struct {
	pulumi.CustomResourceState

	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringOutput `pulumi:"apiManagementName"`
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data pulumi.StringOutput `pulumi:"data"`
	// The Expiration Date of this Certificate, formatted as an RFC3339 string.
	Expiration pulumi.StringOutput `pulumi:"expiration"`
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Subject of this Certificate.
	Subject pulumi.StringOutput `pulumi:"subject"`
	// The Thumbprint of this Certificate.
	Thumbprint pulumi.StringOutput `pulumi:"thumbprint"`
}

// NewCertificate registers a new resource with the given unique name, arguments, and options.
func NewCertificate(ctx *pulumi.Context,
	name string, args *CertificateArgs, opts ...pulumi.ResourceOption) (*Certificate, error) {
	if args == nil || args.ApiManagementName == nil {
		return nil, errors.New("missing required argument 'ApiManagementName'")
	}
	if args == nil || args.Data == nil {
		return nil, errors.New("missing required argument 'Data'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CertificateArgs{}
	}
	var resource Certificate
	err := ctx.RegisterResource("azure:apimanagement/certificate:Certificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertificate gets an existing Certificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertificateState, opts ...pulumi.ResourceOption) (*Certificate, error) {
	var resource Certificate
	err := ctx.ReadResource("azure:apimanagement/certificate:Certificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Certificate resources.
type certificateState struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName *string `pulumi:"apiManagementName"`
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data *string `pulumi:"data"`
	// The Expiration Date of this Certificate, formatted as an RFC3339 string.
	Expiration *string `pulumi:"expiration"`
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Subject of this Certificate.
	Subject *string `pulumi:"subject"`
	// The Thumbprint of this Certificate.
	Thumbprint *string `pulumi:"thumbprint"`
}

type CertificateState struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringPtrInput
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data pulumi.StringPtrInput
	// The Expiration Date of this Certificate, formatted as an RFC3339 string.
	Expiration pulumi.StringPtrInput
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Subject of this Certificate.
	Subject pulumi.StringPtrInput
	// The Thumbprint of this Certificate.
	Thumbprint pulumi.StringPtrInput
}

func (CertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateState)(nil)).Elem()
}

type certificateArgs struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName string `pulumi:"apiManagementName"`
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data string `pulumi:"data"`
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password *string `pulumi:"password"`
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a Certificate resource.
type CertificateArgs struct {
	// The Name of the API Management Service where this Service should be created. Changing this forces a new resource to be created.
	ApiManagementName pulumi.StringInput
	// The base-64 encoded certificate data, which must be a PFX file. Changing this forces a new resource to be created.
	Data pulumi.StringInput
	// The name of the API Management Certificate. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The password used for this certificate. Changing this forces a new resource to be created.
	Password pulumi.StringPtrInput
	// The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certificateArgs)(nil)).Elem()
}
