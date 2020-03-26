// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Availability Set for Virtual Machines.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/availability_set.html.markdown.
type AvailabilitySet struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed pulumi.BoolPtrOutput `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to `3`.
	PlatformFaultDomainCount pulumi.IntPtrOutput `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to `5`.
	PlatformUpdateDomainCount pulumi.IntPtrOutput `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrOutput `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAvailabilitySet registers a new resource with the given unique name, arguments, and options.
func NewAvailabilitySet(ctx *pulumi.Context,
	name string, args *AvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AvailabilitySetArgs{}
	}
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure:compute/availabilitySet:AvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAvailabilitySet gets an existing AvailabilitySet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	var resource AvailabilitySet
	err := ctx.ReadResource("azure:compute/availabilitySet:AvailabilitySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AvailabilitySet resources.
type availabilitySetState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed *bool `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to `3`.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to `5`.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AvailabilitySetState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed pulumi.BoolPtrInput
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used. Defaults to `3`.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specifies the number of update domains that are used. Defaults to `5`.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AvailabilitySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetState)(nil)).Elem()
}

type availabilitySetArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed *bool `pulumi:"managed"`
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used. Defaults to `3`.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specifies the number of update domains that are used. Defaults to `5`.
	PlatformUpdateDomainCount *int `pulumi:"platformUpdateDomainCount"`
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId *string `pulumi:"proximityPlacementGroupId"`
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AvailabilitySet resource.
type AvailabilitySetArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies whether the availability set is managed or not. Possible values are `true` (to specify aligned) or `false` (to specify classic). Default is `true`.
	Managed pulumi.BoolPtrInput
	// Specifies the name of the availability set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used. Defaults to `3`.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specifies the number of update domains that are used. Defaults to `5`.
	PlatformUpdateDomainCount pulumi.IntPtrInput
	// The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
	ProximityPlacementGroupId pulumi.StringPtrInput
	// The name of the resource group in which to create the availability set. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AvailabilitySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetArgs)(nil)).Elem()
}
