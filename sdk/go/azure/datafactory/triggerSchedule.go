// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package datafactory

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Trigger Schedule inside a Azure Data Factory.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/data_factory_trigger_schedule.html.markdown.
type TriggerSchedule struct {
	pulumi.CustomResourceState

	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The time the Schedule Trigger should end. The time will be represented in UTC. 
	EndTime pulumi.StringPtrOutput `pulumi:"endTime"`
	// The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
	Frequency pulumi.StringPtrOutput `pulumi:"frequency"`
	// The interval for how often the trigger occurs. This defaults to 1.
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Data Factory Pipeline name that the trigger will act on.
	PipelineName pulumi.StringOutput `pulumi:"pipelineName"`
	// The pipeline parameters that the trigger will act upon.
	PipelineParameters pulumi.StringMapOutput `pulumi:"pipelineParameters"`
	// The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC. 
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewTriggerSchedule registers a new resource with the given unique name, arguments, and options.
func NewTriggerSchedule(ctx *pulumi.Context,
	name string, args *TriggerScheduleArgs, opts ...pulumi.ResourceOption) (*TriggerSchedule, error) {
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.PipelineName == nil {
		return nil, errors.New("missing required argument 'PipelineName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &TriggerScheduleArgs{}
	}
	var resource TriggerSchedule
	err := ctx.RegisterResource("azure:datafactory/triggerSchedule:TriggerSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTriggerSchedule gets an existing TriggerSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTriggerSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TriggerScheduleState, opts ...pulumi.ResourceOption) (*TriggerSchedule, error) {
	var resource TriggerSchedule
	err := ctx.ReadResource("azure:datafactory/triggerSchedule:TriggerSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TriggerSchedule resources.
type triggerScheduleState struct {
	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The time the Schedule Trigger should end. The time will be represented in UTC. 
	EndTime *string `pulumi:"endTime"`
	// The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
	Frequency *string `pulumi:"frequency"`
	// The interval for how often the trigger occurs. This defaults to 1.
	Interval *int `pulumi:"interval"`
	// Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The Data Factory Pipeline name that the trigger will act on.
	PipelineName *string `pulumi:"pipelineName"`
	// The pipeline parameters that the trigger will act upon.
	PipelineParameters map[string]string `pulumi:"pipelineParameters"`
	// The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC. 
	StartTime *string `pulumi:"startTime"`
}

type TriggerScheduleState struct {
	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The time the Schedule Trigger should end. The time will be represented in UTC. 
	EndTime pulumi.StringPtrInput
	// The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
	Frequency pulumi.StringPtrInput
	// The interval for how often the trigger occurs. This defaults to 1.
	Interval pulumi.IntPtrInput
	// Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The Data Factory Pipeline name that the trigger will act on.
	PipelineName pulumi.StringPtrInput
	// The pipeline parameters that the trigger will act upon.
	PipelineParameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC. 
	StartTime pulumi.StringPtrInput
}

func (TriggerScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerScheduleState)(nil)).Elem()
}

type triggerScheduleArgs struct {
	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The time the Schedule Trigger should end. The time will be represented in UTC. 
	EndTime *string `pulumi:"endTime"`
	// The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
	Frequency *string `pulumi:"frequency"`
	// The interval for how often the trigger occurs. This defaults to 1.
	Interval *int `pulumi:"interval"`
	// Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// The Data Factory Pipeline name that the trigger will act on.
	PipelineName string `pulumi:"pipelineName"`
	// The pipeline parameters that the trigger will act upon.
	PipelineParameters map[string]string `pulumi:"pipelineParameters"`
	// The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC. 
	StartTime *string `pulumi:"startTime"`
}

// The set of arguments for constructing a TriggerSchedule resource.
type TriggerScheduleArgs struct {
	// List of tags that can be used for describing the Data Factory Schedule Trigger.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Schedule Trigger with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The time the Schedule Trigger should end. The time will be represented in UTC. 
	EndTime pulumi.StringPtrInput
	// The trigger freqency. Valid values include `Minute`, `Hour`, `Day`, `Week`, `Month`. Defaults to `Minute`.
	Frequency pulumi.StringPtrInput
	// The interval for how often the trigger occurs. This defaults to 1.
	Interval pulumi.IntPtrInput
	// Specifies the name of the Data Factory Schedule Trigger. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// The Data Factory Pipeline name that the trigger will act on.
	PipelineName pulumi.StringInput
	// The pipeline parameters that the trigger will act upon.
	PipelineParameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Schedule Trigger. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// The time the Schedule Trigger will start. This defaults to the current time. The time will be represented in UTC. 
	StartTime pulumi.StringPtrInput
}

func (TriggerScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*triggerScheduleArgs)(nil)).Elem()
}

