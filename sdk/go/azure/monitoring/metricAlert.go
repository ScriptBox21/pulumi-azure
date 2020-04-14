// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Metric Alert within Azure Monitor.
type MetricAlert struct {
	pulumi.CustomResourceState

	// One or more `action` blocks as defined below.
	Actions MetricAlertActionArrayOutput `pulumi:"actions"`
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate pulumi.BoolPtrOutput `pulumi:"autoMitigate"`
	// One or more `criteria` blocks as defined below.
	Criterias MetricAlertCriteriaArrayOutput `pulumi:"criterias"`
	// The description of this Metric Alert.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency pulumi.StringPtrOutput `pulumi:"frequency"`
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes pulumi.StringOutput `pulumi:"scopes"`
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity pulumi.IntPtrOutput `pulumi:"severity"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize pulumi.StringPtrOutput `pulumi:"windowSize"`
}

// NewMetricAlert registers a new resource with the given unique name, arguments, and options.
func NewMetricAlert(ctx *pulumi.Context,
	name string, args *MetricAlertArgs, opts ...pulumi.ResourceOption) (*MetricAlert, error) {
	if args == nil || args.Criterias == nil {
		return nil, errors.New("missing required argument 'Criterias'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Scopes == nil {
		return nil, errors.New("missing required argument 'Scopes'")
	}
	if args == nil {
		args = &MetricAlertArgs{}
	}
	var resource MetricAlert
	err := ctx.RegisterResource("azure:monitoring/metricAlert:MetricAlert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricAlert gets an existing MetricAlert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricAlertState, opts ...pulumi.ResourceOption) (*MetricAlert, error) {
	var resource MetricAlert
	err := ctx.ReadResource("azure:monitoring/metricAlert:MetricAlert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricAlert resources.
type metricAlertState struct {
	// One or more `action` blocks as defined below.
	Actions []MetricAlertAction `pulumi:"actions"`
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// One or more `criteria` blocks as defined below.
	Criterias []MetricAlertCriteria `pulumi:"criterias"`
	// The description of this Metric Alert.
	Description *string `pulumi:"description"`
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency *string `pulumi:"frequency"`
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes *string `pulumi:"scopes"`
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity *int `pulumi:"severity"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize *string `pulumi:"windowSize"`
}

type MetricAlertState struct {
	// One or more `action` blocks as defined below.
	Actions MetricAlertActionArrayInput
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate pulumi.BoolPtrInput
	// One or more `criteria` blocks as defined below.
	Criterias MetricAlertCriteriaArrayInput
	// The description of this Metric Alert.
	Description pulumi.StringPtrInput
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency pulumi.StringPtrInput
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName pulumi.StringPtrInput
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes pulumi.StringPtrInput
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize pulumi.StringPtrInput
}

func (MetricAlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlertState)(nil)).Elem()
}

type metricAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions []MetricAlertAction `pulumi:"actions"`
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate *bool `pulumi:"autoMitigate"`
	// One or more `criteria` blocks as defined below.
	Criterias []MetricAlertCriteria `pulumi:"criterias"`
	// The description of this Metric Alert.
	Description *string `pulumi:"description"`
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency *string `pulumi:"frequency"`
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes string `pulumi:"scopes"`
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity *int `pulumi:"severity"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize *string `pulumi:"windowSize"`
}

// The set of arguments for constructing a MetricAlert resource.
type MetricAlertArgs struct {
	// One or more `action` blocks as defined below.
	Actions MetricAlertActionArrayInput
	// Should the alerts in this Metric Alert be auto resolved? Defaults to `true`.
	AutoMitigate pulumi.BoolPtrInput
	// One or more `criteria` blocks as defined below.
	Criterias MetricAlertCriteriaArrayInput
	// The description of this Metric Alert.
	Description pulumi.StringPtrInput
	// Should this Metric Alert be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// The evaluation frequency of this Metric Alert, represented in ISO 8601 duration format. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M` and `PT1H`. Defaults to `PT1M`.
	Frequency pulumi.StringPtrInput
	// The name of the Metric Alert. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the Metric Alert instance.
	ResourceGroupName pulumi.StringInput
	// A set of strings of resource IDs at which the metric criteria should be applied.
	Scopes pulumi.StringInput
	// The severity of this Metric Alert. Possible values are `0`, `1`, `2`, `3` and `4`. Defaults to `3`.
	Severity pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The period of time that is used to monitor alert activity, represented in ISO 8601 duration format. This value must be greater than `frequency`. Possible values are `PT1M`, `PT5M`, `PT15M`, `PT30M`, `PT1H`, `PT6H`, `PT12H` and `P1D`. Defaults to `PT5M`.
	WindowSize pulumi.StringPtrInput
}

func (MetricAlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricAlertArgs)(nil)).Elem()
}
