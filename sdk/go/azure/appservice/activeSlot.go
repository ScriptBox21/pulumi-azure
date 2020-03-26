// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Promotes an App Service Slot to Production within an App Service.
//
// > **Note:** When using Slots - the `appSettings`, `connectionString` and `siteConfig` blocks on the `appservice.AppService` resource will be overwritten when promoting a Slot using the `appservice.ActiveSlot` resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_service_active_slot.html.markdown.
type ActiveSlot struct {
	pulumi.CustomResourceState

	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringOutput `pulumi:"appServiceName"`
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName pulumi.StringOutput `pulumi:"appServiceSlotName"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewActiveSlot registers a new resource with the given unique name, arguments, and options.
func NewActiveSlot(ctx *pulumi.Context,
	name string, args *ActiveSlotArgs, opts ...pulumi.ResourceOption) (*ActiveSlot, error) {
	if args == nil || args.AppServiceName == nil {
		return nil, errors.New("missing required argument 'AppServiceName'")
	}
	if args == nil || args.AppServiceSlotName == nil {
		return nil, errors.New("missing required argument 'AppServiceSlotName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ActiveSlotArgs{}
	}
	var resource ActiveSlot
	err := ctx.RegisterResource("azure:appservice/activeSlot:ActiveSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActiveSlot gets an existing ActiveSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActiveSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActiveSlotState, opts ...pulumi.ResourceOption) (*ActiveSlot, error) {
	var resource ActiveSlot
	err := ctx.ReadResource("azure:appservice/activeSlot:ActiveSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActiveSlot resources.
type activeSlotState struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName *string `pulumi:"appServiceName"`
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName *string `pulumi:"appServiceSlotName"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type ActiveSlotState struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringPtrInput
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName pulumi.StringPtrInput
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (ActiveSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*activeSlotState)(nil)).Elem()
}

type activeSlotArgs struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName string `pulumi:"appServiceName"`
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName string `pulumi:"appServiceSlotName"`
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a ActiveSlot resource.
type ActiveSlotArgs struct {
	// The name of the App Service within which the Slot exists.  Changing this forces a new resource to be created.
	AppServiceName pulumi.StringInput
	// The name of the App Service Slot which should be promoted to the Production Slot within the App Service.
	AppServiceSlotName pulumi.StringInput
	// The name of the resource group in which the App Service exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (ActiveSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activeSlotArgs)(nil)).Elem()
}
