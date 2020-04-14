// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iot

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an IotHub Device Provisioning Service.
type IotHubDps struct {
	pulumi.CustomResourceState

	// The allocation policy of the IoT Device Provisioning Service.
	AllocationPolicy pulumi.StringOutput `pulumi:"allocationPolicy"`
	// The device endpoint of the IoT Device Provisioning Service.
	DeviceProvisioningHostName pulumi.StringOutput `pulumi:"deviceProvisioningHostName"`
	// The unique identifier of the IoT Device Provisioning Service.
	IdScope pulumi.StringOutput `pulumi:"idScope"`
	// A `linkedHub` block as defined below.
	LinkedHubs IotHubDpsLinkedHubArrayOutput `pulumi:"linkedHubs"`
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The service endpoint of the IoT Device Provisioning Service.
	ServiceOperationsHostName pulumi.StringOutput `pulumi:"serviceOperationsHostName"`
	// A `sku` block as defined below.
	Sku IotHubDpsSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewIotHubDps registers a new resource with the given unique name, arguments, and options.
func NewIotHubDps(ctx *pulumi.Context,
	name string, args *IotHubDpsArgs, opts ...pulumi.ResourceOption) (*IotHubDps, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &IotHubDpsArgs{}
	}
	var resource IotHubDps
	err := ctx.RegisterResource("azure:iot/iotHubDps:IotHubDps", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubDps gets an existing IotHubDps resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubDps(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubDpsState, opts ...pulumi.ResourceOption) (*IotHubDps, error) {
	var resource IotHubDps
	err := ctx.ReadResource("azure:iot/iotHubDps:IotHubDps", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubDps resources.
type iotHubDpsState struct {
	// The allocation policy of the IoT Device Provisioning Service.
	AllocationPolicy *string `pulumi:"allocationPolicy"`
	// The device endpoint of the IoT Device Provisioning Service.
	DeviceProvisioningHostName *string `pulumi:"deviceProvisioningHostName"`
	// The unique identifier of the IoT Device Provisioning Service.
	IdScope *string `pulumi:"idScope"`
	// A `linkedHub` block as defined below.
	LinkedHubs []IotHubDpsLinkedHub `pulumi:"linkedHubs"`
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The service endpoint of the IoT Device Provisioning Service.
	ServiceOperationsHostName *string `pulumi:"serviceOperationsHostName"`
	// A `sku` block as defined below.
	Sku *IotHubDpsSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type IotHubDpsState struct {
	// The allocation policy of the IoT Device Provisioning Service.
	AllocationPolicy pulumi.StringPtrInput
	// The device endpoint of the IoT Device Provisioning Service.
	DeviceProvisioningHostName pulumi.StringPtrInput
	// The unique identifier of the IoT Device Provisioning Service.
	IdScope pulumi.StringPtrInput
	// A `linkedHub` block as defined below.
	LinkedHubs IotHubDpsLinkedHubArrayInput
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The service endpoint of the IoT Device Provisioning Service.
	ServiceOperationsHostName pulumi.StringPtrInput
	// A `sku` block as defined below.
	Sku IotHubDpsSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (IotHubDpsState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubDpsState)(nil)).Elem()
}

type iotHubDpsArgs struct {
	// A `linkedHub` block as defined below.
	LinkedHubs []IotHubDpsLinkedHub `pulumi:"linkedHubs"`
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku IotHubDpsSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a IotHubDps resource.
type IotHubDpsArgs struct {
	// A `linkedHub` block as defined below.
	LinkedHubs IotHubDpsLinkedHubArrayInput
	// Specifies the supported Azure location where the resource has to be createc. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Iot Device Provisioning Service resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group under which the Iot Device Provisioning Service resource has to be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `sku` block as defined below.
	Sku IotHubDpsSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (IotHubDpsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubDpsArgs)(nil)).Elem()
}
