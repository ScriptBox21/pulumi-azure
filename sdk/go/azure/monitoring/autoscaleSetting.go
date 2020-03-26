// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package monitoring

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a AutoScale Setting which can be applied to Virtual Machine Scale Sets, App Services and other scalable resources.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/monitor_autoscale_setting.html.markdown.
type AutoscaleSetting struct {
	pulumi.CustomResourceState

	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies a `notification` block as defined below.
	Notification AutoscaleSettingNotificationPtrOutput `pulumi:"notification"`
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles AutoscaleSettingProfileArrayOutput `pulumi:"profiles"`
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId pulumi.StringOutput `pulumi:"targetResourceId"`
}

// NewAutoscaleSetting registers a new resource with the given unique name, arguments, and options.
func NewAutoscaleSetting(ctx *pulumi.Context,
	name string, args *AutoscaleSettingArgs, opts ...pulumi.ResourceOption) (*AutoscaleSetting, error) {
	if args == nil || args.Profiles == nil {
		return nil, errors.New("missing required argument 'Profiles'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TargetResourceId == nil {
		return nil, errors.New("missing required argument 'TargetResourceId'")
	}
	if args == nil {
		args = &AutoscaleSettingArgs{}
	}
	var resource AutoscaleSetting
	err := ctx.RegisterResource("azure:monitoring/autoscaleSetting:AutoscaleSetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoscaleSetting gets an existing AutoscaleSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoscaleSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoscaleSettingState, opts ...pulumi.ResourceOption) (*AutoscaleSetting, error) {
	var resource AutoscaleSetting
	err := ctx.ReadResource("azure:monitoring/autoscaleSetting:AutoscaleSetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoscaleSetting resources.
type autoscaleSettingState struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies a `notification` block as defined below.
	Notification *AutoscaleSettingNotification `pulumi:"notification"`
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles []AutoscaleSettingProfile `pulumi:"profiles"`
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId *string `pulumi:"targetResourceId"`
}

type AutoscaleSettingState struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies a `notification` block as defined below.
	Notification AutoscaleSettingNotificationPtrInput
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles AutoscaleSettingProfileArrayInput
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId pulumi.StringPtrInput
}

func (AutoscaleSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleSettingState)(nil)).Elem()
}

type autoscaleSettingArgs struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies a `notification` block as defined below.
	Notification *AutoscaleSettingNotification `pulumi:"notification"`
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles []AutoscaleSettingProfile `pulumi:"profiles"`
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId string `pulumi:"targetResourceId"`
}

// The set of arguments for constructing a AutoscaleSetting resource.
type AutoscaleSettingArgs struct {
	// Specifies whether automatic scaling is enabled for the target resource. Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the AutoScale Setting should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the AutoScale Setting. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies a `notification` block as defined below.
	Notification AutoscaleSettingNotificationPtrInput
	// Specifies one or more (up to 20) `profile` blocks as defined below.
	Profiles AutoscaleSettingProfileArrayInput
	// The name of the Resource Group in the AutoScale Setting should be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the resource ID of the resource that the autoscale setting should be added to.
	TargetResourceId pulumi.StringInput
}

func (AutoscaleSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoscaleSettingArgs)(nil)).Elem()
}
