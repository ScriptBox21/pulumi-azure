// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an IotHub Device Provisioning Service Shared Access Policy
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/iothub_dps_shared_access_policy.html.markdown.
type DpsSharedAccessPolicy struct {
	pulumi.CustomResourceState

	// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
	EnrollmentRead pulumi.BoolPtrOutput `pulumi:"enrollmentRead"`
	// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
	EnrollmentWrite pulumi.BoolPtrOutput `pulumi:"enrollmentWrite"`
	// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubDpsName pulumi.StringOutput `pulumi:"iothubDpsName"`
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The primary connection string of the Shared Access Policy.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The primary key used to create the authentication token.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
	RegistrationRead pulumi.BoolPtrOutput `pulumi:"registrationRead"`
	// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
	RegistrationWrite pulumi.BoolPtrOutput `pulumi:"registrationWrite"`
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary connection string of the Shared Access Policy.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The secondary key used to create the authentication token.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	ServiceConfig pulumi.BoolPtrOutput `pulumi:"serviceConfig"`
}

// NewDpsSharedAccessPolicy registers a new resource with the given unique name, arguments, and options.
func NewDpsSharedAccessPolicy(ctx *pulumi.Context,
	name string, args *DpsSharedAccessPolicyArgs, opts ...pulumi.ResourceOption) (*DpsSharedAccessPolicy, error) {
	if args == nil || args.IothubDpsName == nil {
		return nil, errors.New("missing required argument 'IothubDpsName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DpsSharedAccessPolicyArgs{}
	}
	var resource DpsSharedAccessPolicy
	err := ctx.RegisterResource("azure:iot/dpsSharedAccessPolicy:DpsSharedAccessPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDpsSharedAccessPolicy gets an existing DpsSharedAccessPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDpsSharedAccessPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DpsSharedAccessPolicyState, opts ...pulumi.ResourceOption) (*DpsSharedAccessPolicy, error) {
	var resource DpsSharedAccessPolicy
	err := ctx.ReadResource("azure:iot/dpsSharedAccessPolicy:DpsSharedAccessPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DpsSharedAccessPolicy resources.
type dpsSharedAccessPolicyState struct {
	// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
	EnrollmentRead *bool `pulumi:"enrollmentRead"`
	// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
	EnrollmentWrite *bool `pulumi:"enrollmentWrite"`
	// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubDpsName *string `pulumi:"iothubDpsName"`
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The primary connection string of the Shared Access Policy.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The primary key used to create the authentication token.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
	RegistrationRead *bool `pulumi:"registrationRead"`
	// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
	RegistrationWrite *bool `pulumi:"registrationWrite"`
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary connection string of the Shared Access Policy.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The secondary key used to create the authentication token.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	ServiceConfig *bool `pulumi:"serviceConfig"`
}

type DpsSharedAccessPolicyState struct {
	// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
	EnrollmentRead pulumi.BoolPtrInput
	// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
	EnrollmentWrite pulumi.BoolPtrInput
	// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubDpsName pulumi.StringPtrInput
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The primary connection string of the Shared Access Policy.
	PrimaryConnectionString pulumi.StringPtrInput
	// The primary key used to create the authentication token.
	PrimaryKey pulumi.StringPtrInput
	// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
	RegistrationRead pulumi.BoolPtrInput
	// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
	RegistrationWrite pulumi.BoolPtrInput
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary connection string of the Shared Access Policy.
	SecondaryConnectionString pulumi.StringPtrInput
	// The secondary key used to create the authentication token.
	SecondaryKey pulumi.StringPtrInput
	// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	ServiceConfig pulumi.BoolPtrInput
}

func (DpsSharedAccessPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dpsSharedAccessPolicyState)(nil)).Elem()
}

type dpsSharedAccessPolicyArgs struct {
	// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
	EnrollmentRead *bool `pulumi:"enrollmentRead"`
	// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
	EnrollmentWrite *bool `pulumi:"enrollmentWrite"`
	// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubDpsName string `pulumi:"iothubDpsName"`
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
	RegistrationRead *bool `pulumi:"registrationRead"`
	// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
	RegistrationWrite *bool `pulumi:"registrationWrite"`
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	ServiceConfig *bool `pulumi:"serviceConfig"`
}

// The set of arguments for constructing a DpsSharedAccessPolicy resource.
type DpsSharedAccessPolicyArgs struct {
	// Adds `EnrollmentRead` permission to this Shared Access Account. It allows read access to enrollment data.
	EnrollmentRead pulumi.BoolPtrInput
	// Adds `EnrollmentWrite` permission to this Shared Access Account. It allows write access to enrollment data.
	EnrollmentWrite pulumi.BoolPtrInput
	// The name of the IoT Hub Device Provisioning service to which this Shared Access Policy belongs. Changing this forces a new resource to be created.
	IothubDpsName pulumi.StringInput
	// Specifies the name of the IotHub Shared Access Policy resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Adds `RegistrationStatusRead` permission to this Shared Access Account. It allows read access to device registrations.
	RegistrationRead pulumi.BoolPtrInput
	// Adds `RegistrationStatusWrite` permission to this Shared Access Account. It allows write access to device registrations.
	RegistrationWrite pulumi.BoolPtrInput
	// The name of the resource group under which the IotHub Shared Access Policy resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Adds `ServiceConfig` permission to this Shared Access Account. It allows configuration of the Device Provisioning Service.
	ServiceConfig pulumi.BoolPtrInput
}

func (DpsSharedAccessPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dpsSharedAccessPolicyArgs)(nil)).Elem()
}
