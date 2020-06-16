// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Orchestrated Virtual Machine Scale Set.
//
// > **Note:** Orchestrated Virtual Machine Scale Sets are in Public Preview - [more details can be found in the Azure Documentation](https://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/orchestration-modes).
//
// ## Example Usage
//
//
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/compute"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleOrchestratedVirtualMachineScaleSet, err := compute.NewOrchestratedVirtualMachineScaleSet(ctx, "exampleOrchestratedVirtualMachineScaleSet", &compute.OrchestratedVirtualMachineScaleSetArgs{
// 			Location:                 exampleResourceGroup.Location,
// 			ResourceGroupName:        exampleResourceGroup.Name,
// 			PlatformFaultDomainCount: pulumi.Int(5),
// 			SinglePlacementGroup:     pulumi.Bool(true),
// 			Zones: pulumi.String(pulumi.String{
// 				pulumi.String("1"),
// 			}),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type OrchestratedVirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntOutput `pulumi:"platformFaultDomainCount"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
	SinglePlacementGroup pulumi.BoolOutput `pulumi:"singlePlacementGroup"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrOutput `pulumi:"zones"`
}

// NewOrchestratedVirtualMachineScaleSet registers a new resource with the given unique name, arguments, and options.
func NewOrchestratedVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *OrchestratedVirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*OrchestratedVirtualMachineScaleSet, error) {
	if args == nil || args.PlatformFaultDomainCount == nil {
		return nil, errors.New("missing required argument 'PlatformFaultDomainCount'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SinglePlacementGroup == nil {
		return nil, errors.New("missing required argument 'SinglePlacementGroup'")
	}
	if args == nil {
		args = &OrchestratedVirtualMachineScaleSetArgs{}
	}
	var resource OrchestratedVirtualMachineScaleSet
	err := ctx.RegisterResource("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrchestratedVirtualMachineScaleSet gets an existing OrchestratedVirtualMachineScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrchestratedVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrchestratedVirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*OrchestratedVirtualMachineScaleSet, error) {
	var resource OrchestratedVirtualMachineScaleSet
	err := ctx.ReadResource("azure:compute/orchestratedVirtualMachineScaleSet:OrchestratedVirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrchestratedVirtualMachineScaleSet resources.
type orchestratedVirtualMachineScaleSetState struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags map[string]string `pulumi:"tags"`
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId *string `pulumi:"uniqueId"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

type OrchestratedVirtualMachineScaleSetState struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
	SinglePlacementGroup pulumi.BoolPtrInput
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags pulumi.StringMapInput
	// The Unique ID for the Orchestrated Virtual Machine Scale Set.
	UniqueId pulumi.StringPtrInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (OrchestratedVirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratedVirtualMachineScaleSetState)(nil)).Elem()
}

type orchestratedVirtualMachineScaleSetArgs struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount int `pulumi:"platformFaultDomainCount"`
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
	SinglePlacementGroup bool `pulumi:"singlePlacementGroup"`
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags map[string]string `pulumi:"tags"`
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones *string `pulumi:"zones"`
}

// The set of arguments for constructing a OrchestratedVirtualMachineScaleSet resource.
type OrchestratedVirtualMachineScaleSetArgs struct {
	// The Azure location where the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the number of fault domains that are used by this Orchestrated Virtual Machine Scale Set. Changing this forces a new resource to be created.
	PlatformFaultDomainCount pulumi.IntInput
	// The name of the Resource Group in which the Orchestrated Virtual Machine Scale Set should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Should the Orchestrated Virtual Machine Scale Set use single placement group? Changing this forces a new resource to be created.
	SinglePlacementGroup pulumi.BoolInput
	// A mapping of tags which should be assigned to this Orchestrated Virtual Machine Scale Set.
	Tags pulumi.StringMapInput
	// A list of Availability Zones in which the Virtual Machines in this Scale Set should be created in. Changing this forces a new resource to be created.
	Zones pulumi.StringPtrInput
}

func (OrchestratedVirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orchestratedVirtualMachineScaleSetArgs)(nil)).Elem()
}