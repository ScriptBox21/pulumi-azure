// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an AzureNetwork DDoS Protection Plan.
//
// > **NOTE** Azure only allows `one` DDoS Protection Plan per region.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_ddos_protection_plan.html.markdown.
type DdosProtectionPlan struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A list of Virtual Network ID's associated with the DDoS Protection Plan.
	VirtualNetworkIds pulumi.StringArrayOutput `pulumi:"virtualNetworkIds"`
}

// NewDdosProtectionPlan registers a new resource with the given unique name, arguments, and options.
func NewDdosProtectionPlan(ctx *pulumi.Context,
	name string, args *DdosProtectionPlanArgs, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DdosProtectionPlanArgs{}
	}
	var resource DdosProtectionPlan
	err := ctx.RegisterResource("azure:network/ddosProtectionPlan:DdosProtectionPlan", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDdosProtectionPlan gets an existing DdosProtectionPlan resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDdosProtectionPlan(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DdosProtectionPlanState, opts ...pulumi.ResourceOption) (*DdosProtectionPlan, error) {
	var resource DdosProtectionPlan
	err := ctx.ReadResource("azure:network/ddosProtectionPlan:DdosProtectionPlan", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DdosProtectionPlan resources.
type ddosProtectionPlanState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// A list of Virtual Network ID's associated with the DDoS Protection Plan.
	VirtualNetworkIds []string `pulumi:"virtualNetworkIds"`
}

type DdosProtectionPlanState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// A list of Virtual Network ID's associated with the DDoS Protection Plan.
	VirtualNetworkIds pulumi.StringArrayInput
}

func (DdosProtectionPlanState) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosProtectionPlanState)(nil)).Elem()
}

type ddosProtectionPlanArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DdosProtectionPlan resource.
type DdosProtectionPlanArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the Network DDoS Protection Plan. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (DdosProtectionPlanArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ddosProtectionPlanArgs)(nil)).Elem()
}
