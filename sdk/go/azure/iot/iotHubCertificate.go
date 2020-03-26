// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IotHub Device Provisioning Service Certificate.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/iothub_dps_certificate.html.markdown.
type IotHubCertificate struct {
	pulumi.CustomResourceState

	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringOutput `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringOutput `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewIotHubCertificate registers a new resource with the given unique name, arguments, and options.
func NewIotHubCertificate(ctx *pulumi.Context,
	name string, args *IotHubCertificateArgs, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	if args == nil || args.CertificateContent == nil {
		return nil, errors.New("missing required argument 'CertificateContent'")
	}
	if args == nil || args.IotDpsName == nil {
		return nil, errors.New("missing required argument 'IotDpsName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &IotHubCertificateArgs{}
	}
	var resource IotHubCertificate
	err := ctx.RegisterResource("azure:iot/iotHubCertificate:IotHubCertificate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubCertificate gets an existing IotHubCertificate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubCertificate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubCertificateState, opts ...pulumi.ResourceOption) (*IotHubCertificate, error) {
	var resource IotHubCertificate
	err := ctx.ReadResource("azure:iot/iotHubCertificate:IotHubCertificate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubCertificate resources.
type iotHubCertificateState struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent *string `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName *string `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type IotHubCertificateState struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringPtrInput
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringPtrInput
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (IotHubCertificateState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateState)(nil)).Elem()
}

type iotHubCertificateArgs struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent string `pulumi:"certificateContent"`
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName string `pulumi:"iotDpsName"`
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a IotHubCertificate resource.
type IotHubCertificateArgs struct {
	// The Base-64 representation of the X509 leaf certificate .cer file or just a .pem file content.
	CertificateContent pulumi.StringInput
	// The name of the IoT Device Provisioning Service that this certificate will be attached to. Changing this forces a new resource to be created.
	IotDpsName pulumi.StringInput
	// Specifies the name of the Iot Device Provisioning Service Certificate resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service Certificate resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (IotHubCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubCertificateArgs)(nil)).Elem()
}
