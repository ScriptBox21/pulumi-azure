// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package monitoring

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Activity Log Alert within Azure Monitor.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/monitor_activity_log_alert.html.markdown.
type ActivityLogAlert struct {
	pulumi.CustomResourceState

	// One or more `action` blocks as defined below.
	Actions ActivityLogAlertActionArrayOutput `pulumi:"actions"`
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteriaOutput `pulumi:"criteria"`
	// The description of this activity log alert.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes pulumi.StringArrayOutput `pulumi:"scopes"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewActivityLogAlert registers a new resource with the given unique name, arguments, and options.
func NewActivityLogAlert(ctx *pulumi.Context,
	name string, args *ActivityLogAlertArgs, opts ...pulumi.ResourceOption) (*ActivityLogAlert, error) {
	if args == nil || args.Criteria == nil {
		return nil, errors.New("missing required argument 'Criteria'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	if args == nil {
		args = &ActivityLogAlertArgs{}
	}
	var resource ActivityLogAlert
	err := ctx.RegisterResource("azure:monitoring/activityLogAlert:ActivityLogAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetActivityLogAlert gets an existing ActivityLogAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetActivityLogAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ActivityLogAlertState, opts ...pulumi.ResourceOption) (*ActivityLogAlert, error) {
	var resource ActivityLogAlert
	err := ctx.ReadResource("azure:monitoring/activityLogAlert:ActivityLogAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ActivityLogAlert resources.
type activityLogAlertState struct {
	// One or more `action` blocks as defined below.
	Actions []ActivityLogAlertAction `pulumi:"actions"`
	// A `criteria` block as defined below.
	Criteria *ActivityLogAlertCriteria `pulumi:"criteria"`
	// The description of this activity log alert.
	Description *string `pulumi:"description"`
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes []string `pulumi:"scopes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type ActivityLogAlertState struct {
	// One or more `action` blocks as defined below.
	Actions ActivityLogAlertActionArrayInput
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteriaPtrInput
	// The description of this activity log alert.
	Description pulumi.StringPtrInput
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName pulumi.StringPtrInput
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActivityLogAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*activityLogAlertState)(nil)).Elem()
}

type activityLogAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions []ActivityLogAlertAction `pulumi:"actions"`
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteria `pulumi:"criteria"`
	// The description of this activity log alert.
	Description *string `pulumi:"description"`
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes []string `pulumi:"scopes"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ActivityLogAlert resource.
type ActivityLogAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions ActivityLogAlertActionArrayInput
	// A `criteria` block as defined below.
	Criteria ActivityLogAlertCriteriaInput
	// The description of this activity log alert.
	Description pulumi.StringPtrInput
	// Should this Activity Log Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The name of the activity log alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the activity log alert instance.
	ResourceGroupName pulumi.StringInput
	// The Scope at which the Activity Log should be applied, for example a the Resource ID of a Subscription or a Resource (such as a Storage Account).
	Scopes pulumi.StringArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ActivityLogAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*activityLogAlertArgs)(nil)).Elem()
}

