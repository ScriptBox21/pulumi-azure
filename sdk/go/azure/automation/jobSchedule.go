// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Links an Automation Runbook and Schedule.
//
// ## Example Usage
//
// This is an example of just the Job Schedule.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/automation"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := automation.NewJobSchedule(ctx, "example", &automation.JobScheduleArgs{
// 			AutomationAccountName: pulumi.String("tf-automation-account"),
// 			Parameters: pulumi.StringMap{
// 				"resourcegroup": pulumi.String("tf-rgr-vm"),
// 				"vmname":        pulumi.String("TF-VM-01"),
// 			},
// 			ResourceGroupName: pulumi.String("tf-rgr-automation"),
// 			RunbookName:       pulumi.String("Get-VirtualMachine"),
// 			ScheduleName:      pulumi.String("hour"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Automation Job Schedules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:automation/jobSchedule:JobSchedule example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Automation/automationAccounts/account1/jobSchedules/10000000-1001-1001-1001-000000000001
// ```
type JobSchedule struct {
	pulumi.CustomResourceState

	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringOutput `pulumi:"automationAccountName"`
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId pulumi.StringOutput `pulumi:"jobScheduleId"`
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn pulumi.StringPtrOutput `pulumi:"runOn"`
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  pulumi.StringOutput `pulumi:"runbookName"`
	ScheduleName pulumi.StringOutput `pulumi:"scheduleName"`
}

// NewJobSchedule registers a new resource with the given unique name, arguments, and options.
func NewJobSchedule(ctx *pulumi.Context,
	name string, args *JobScheduleArgs, opts ...pulumi.ResourceOption) (*JobSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RunbookName == nil {
		return nil, errors.New("invalid value for required argument 'RunbookName'")
	}
	if args.ScheduleName == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleName'")
	}
	var resource JobSchedule
	err := ctx.RegisterResource("azure:automation/jobSchedule:JobSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobSchedule gets an existing JobSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobScheduleState, opts ...pulumi.ResourceOption) (*JobSchedule, error) {
	var resource JobSchedule
	err := ctx.ReadResource("azure:automation/jobSchedule:JobSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobSchedule resources.
type jobScheduleState struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName *string `pulumi:"automationAccountName"`
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId *string `pulumi:"jobScheduleId"`
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn *string `pulumi:"runOn"`
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  *string `pulumi:"runbookName"`
	ScheduleName *string `pulumi:"scheduleName"`
}

type JobScheduleState struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringPtrInput
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId pulumi.StringPtrInput
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn pulumi.StringPtrInput
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  pulumi.StringPtrInput
	ScheduleName pulumi.StringPtrInput
}

func (JobScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobScheduleState)(nil)).Elem()
}

type jobScheduleArgs struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId *string `pulumi:"jobScheduleId"`
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn *string `pulumi:"runOn"`
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  string `pulumi:"runbookName"`
	ScheduleName string `pulumi:"scheduleName"`
}

// The set of arguments for constructing a JobSchedule resource.
type JobScheduleArgs struct {
	// The name of the Automation Account in which the Job Schedule is created. Changing this forces a new resource to be created.
	AutomationAccountName pulumi.StringInput
	// The UUID identifying the Automation Job Schedule.
	JobScheduleId pulumi.StringPtrInput
	// A map of key/value pairs corresponding to the arguments that can be passed to the Runbook. Changing this forces a new resource to be created.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which the Job Schedule is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Name of a Hybrid Worker Group the Runbook will be executed on. Changing this forces a new resource to be created.
	RunOn pulumi.StringPtrInput
	// The name of a Runbook to link to a Schedule. It needs to be in the same Automation Account as the Schedule and Job Schedule. Changing this forces a new resource to be created.
	RunbookName  pulumi.StringInput
	ScheduleName pulumi.StringInput
}

func (JobScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobScheduleArgs)(nil)).Elem()
}

type JobScheduleInput interface {
	pulumi.Input

	ToJobScheduleOutput() JobScheduleOutput
	ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput
}

func (*JobSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSchedule)(nil))
}

func (i *JobSchedule) ToJobScheduleOutput() JobScheduleOutput {
	return i.ToJobScheduleOutputWithContext(context.Background())
}

func (i *JobSchedule) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleOutput)
}

func (i *JobSchedule) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return i.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (i *JobSchedule) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSchedulePtrOutput)
}

type JobSchedulePtrInput interface {
	pulumi.Input

	ToJobSchedulePtrOutput() JobSchedulePtrOutput
	ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput
}

type jobSchedulePtrType JobScheduleArgs

func (*jobSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil))
}

func (i *jobSchedulePtrType) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return i.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (i *jobSchedulePtrType) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSchedulePtrOutput)
}

// JobScheduleArrayInput is an input type that accepts JobScheduleArray and JobScheduleArrayOutput values.
// You can construct a concrete instance of `JobScheduleArrayInput` via:
//
//          JobScheduleArray{ JobScheduleArgs{...} }
type JobScheduleArrayInput interface {
	pulumi.Input

	ToJobScheduleArrayOutput() JobScheduleArrayOutput
	ToJobScheduleArrayOutputWithContext(context.Context) JobScheduleArrayOutput
}

type JobScheduleArray []JobScheduleInput

func (JobScheduleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*JobSchedule)(nil))
}

func (i JobScheduleArray) ToJobScheduleArrayOutput() JobScheduleArrayOutput {
	return i.ToJobScheduleArrayOutputWithContext(context.Background())
}

func (i JobScheduleArray) ToJobScheduleArrayOutputWithContext(ctx context.Context) JobScheduleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleArrayOutput)
}

// JobScheduleMapInput is an input type that accepts JobScheduleMap and JobScheduleMapOutput values.
// You can construct a concrete instance of `JobScheduleMapInput` via:
//
//          JobScheduleMap{ "key": JobScheduleArgs{...} }
type JobScheduleMapInput interface {
	pulumi.Input

	ToJobScheduleMapOutput() JobScheduleMapOutput
	ToJobScheduleMapOutputWithContext(context.Context) JobScheduleMapOutput
}

type JobScheduleMap map[string]JobScheduleInput

func (JobScheduleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*JobSchedule)(nil))
}

func (i JobScheduleMap) ToJobScheduleMapOutput() JobScheduleMapOutput {
	return i.ToJobScheduleMapOutputWithContext(context.Background())
}

func (i JobScheduleMap) ToJobScheduleMapOutputWithContext(ctx context.Context) JobScheduleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleMapOutput)
}

type JobScheduleOutput struct {
	*pulumi.OutputState
}

func (JobScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSchedule)(nil))
}

func (o JobScheduleOutput) ToJobScheduleOutput() JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return o.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (o JobScheduleOutput) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return o.ApplyT(func(v JobSchedule) *JobSchedule {
		return &v
	}).(JobSchedulePtrOutput)
}

type JobSchedulePtrOutput struct {
	*pulumi.OutputState
}

func (JobSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil))
}

func (o JobSchedulePtrOutput) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return o
}

func (o JobSchedulePtrOutput) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return o
}

type JobScheduleArrayOutput struct{ *pulumi.OutputState }

func (JobScheduleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobSchedule)(nil))
}

func (o JobScheduleArrayOutput) ToJobScheduleArrayOutput() JobScheduleArrayOutput {
	return o
}

func (o JobScheduleArrayOutput) ToJobScheduleArrayOutputWithContext(ctx context.Context) JobScheduleArrayOutput {
	return o
}

func (o JobScheduleArrayOutput) Index(i pulumi.IntInput) JobScheduleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobSchedule {
		return vs[0].([]JobSchedule)[vs[1].(int)]
	}).(JobScheduleOutput)
}

type JobScheduleMapOutput struct{ *pulumi.OutputState }

func (JobScheduleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]JobSchedule)(nil))
}

func (o JobScheduleMapOutput) ToJobScheduleMapOutput() JobScheduleMapOutput {
	return o
}

func (o JobScheduleMapOutput) ToJobScheduleMapOutputWithContext(ctx context.Context) JobScheduleMapOutput {
	return o
}

func (o JobScheduleMapOutput) MapIndex(k pulumi.StringInput) JobScheduleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) JobSchedule {
		return vs[0].(map[string]JobSchedule)[vs[1].(string)]
	}).(JobScheduleOutput)
}

func init() {
	pulumi.RegisterOutputType(JobScheduleOutput{})
	pulumi.RegisterOutputType(JobSchedulePtrOutput{})
	pulumi.RegisterOutputType(JobScheduleArrayOutput{})
	pulumi.RegisterOutputType(JobScheduleMapOutput{})
}
