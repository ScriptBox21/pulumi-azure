// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sentinel

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Sentinel Scheduled Alert Rule.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/operationalinsights"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sentinel"
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
// 		exampleAnalyticsWorkspace, err := operationalinsights.NewAnalyticsWorkspace(ctx, "exampleAnalyticsWorkspace", &operationalinsights.AnalyticsWorkspaceArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Sku:               pulumi.String("pergb2018"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sentinel.NewAlertRuleScheduled(ctx, "exampleAlertRuleScheduled", &sentinel.AlertRuleScheduledArgs{
// 			LogAnalyticsWorkspaceId: exampleAnalyticsWorkspace.ID(),
// 			DisplayName:             pulumi.String("example"),
// 			Severity:                pulumi.String("High"),
// 			Query:                   pulumi.String(fmt.Sprintf("%v%v%v%v", "AzureActivity |\n", "  where OperationName == \"Create or Update Virtual Machine\" or OperationName ==\"Create Deployment\" |\n", "  where ActivityStatus == \"Succeeded\" |\n", "  make-series dcount(ResourceId) default=0 on EventSubmissionTimestamp in range(ago(7d), now(), 1d) by Caller\n")),
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
// Sentinel Scheduled Alert Rules can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sentinel/alertRuleScheduled:AlertRuleScheduled example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.OperationalInsights/workspaces/workspace1/providers/Microsoft.SecurityInsights/alertRules/rule1
// ```
type AlertRuleScheduled struct {
	pulumi.CustomResourceState

	// The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	AlertRuleTemplateGuid pulumi.StringPtrOutput `pulumi:"alertRuleTemplateGuid"`
	// The description of this Sentinel Scheduled Alert Rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The friendly name of this Sentinel Scheduled Alert Rule.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// A `eventGrouping` block as defined below.
	EventGrouping AlertRuleScheduledEventGroupingPtrOutput `pulumi:"eventGrouping"`
	// A `incidentConfiguration` block as defined below.
	IncidentConfiguration AlertRuleScheduledIncidentConfigurationOutput `pulumi:"incidentConfiguration"`
	// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringOutput `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The query of this Sentinel Scheduled Alert Rule.
	Query pulumi.StringOutput `pulumi:"query"`
	// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
	QueryFrequency pulumi.StringPtrOutput `pulumi:"queryFrequency"`
	// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
	QueryPeriod pulumi.StringPtrOutput `pulumi:"queryPeriod"`
	// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
	Severity pulumi.StringOutput `pulumi:"severity"`
	// If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
	SuppressionDuration pulumi.StringPtrOutput `pulumi:"suppressionDuration"`
	// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
	SuppressionEnabled pulumi.BoolPtrOutput `pulumi:"suppressionEnabled"`
	// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
	Tactics pulumi.StringArrayOutput `pulumi:"tactics"`
	// The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
	TriggerOperator pulumi.StringPtrOutput `pulumi:"triggerOperator"`
	// The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
	TriggerThreshold pulumi.IntPtrOutput `pulumi:"triggerThreshold"`
}

// NewAlertRuleScheduled registers a new resource with the given unique name, arguments, and options.
func NewAlertRuleScheduled(ctx *pulumi.Context,
	name string, args *AlertRuleScheduledArgs, opts ...pulumi.ResourceOption) (*AlertRuleScheduled, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.LogAnalyticsWorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'LogAnalyticsWorkspaceId'")
	}
	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	var resource AlertRuleScheduled
	err := ctx.RegisterResource("azure:sentinel/alertRuleScheduled:AlertRuleScheduled", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlertRuleScheduled gets an existing AlertRuleScheduled resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlertRuleScheduled(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleScheduledState, opts ...pulumi.ResourceOption) (*AlertRuleScheduled, error) {
	var resource AlertRuleScheduled
	err := ctx.ReadResource("azure:sentinel/alertRuleScheduled:AlertRuleScheduled", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlertRuleScheduled resources.
type alertRuleScheduledState struct {
	// The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	AlertRuleTemplateGuid *string `pulumi:"alertRuleTemplateGuid"`
	// The description of this Sentinel Scheduled Alert Rule.
	Description *string `pulumi:"description"`
	// The friendly name of this Sentinel Scheduled Alert Rule.
	DisplayName *string `pulumi:"displayName"`
	// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// A `eventGrouping` block as defined below.
	EventGrouping *AlertRuleScheduledEventGrouping `pulumi:"eventGrouping"`
	// A `incidentConfiguration` block as defined below.
	IncidentConfiguration *AlertRuleScheduledIncidentConfiguration `pulumi:"incidentConfiguration"`
	// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	LogAnalyticsWorkspaceId *string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	Name *string `pulumi:"name"`
	// The query of this Sentinel Scheduled Alert Rule.
	Query *string `pulumi:"query"`
	// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
	QueryFrequency *string `pulumi:"queryFrequency"`
	// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
	QueryPeriod *string `pulumi:"queryPeriod"`
	// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
	Severity *string `pulumi:"severity"`
	// If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
	SuppressionDuration *string `pulumi:"suppressionDuration"`
	// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
	SuppressionEnabled *bool `pulumi:"suppressionEnabled"`
	// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
	Tactics []string `pulumi:"tactics"`
	// The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
	TriggerOperator *string `pulumi:"triggerOperator"`
	// The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
	TriggerThreshold *int `pulumi:"triggerThreshold"`
}

type AlertRuleScheduledState struct {
	// The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	AlertRuleTemplateGuid pulumi.StringPtrInput
	// The description of this Sentinel Scheduled Alert Rule.
	Description pulumi.StringPtrInput
	// The friendly name of this Sentinel Scheduled Alert Rule.
	DisplayName pulumi.StringPtrInput
	// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// A `eventGrouping` block as defined below.
	EventGrouping AlertRuleScheduledEventGroupingPtrInput
	// A `incidentConfiguration` block as defined below.
	IncidentConfiguration AlertRuleScheduledIncidentConfigurationPtrInput
	// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringPtrInput
	// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	Name pulumi.StringPtrInput
	// The query of this Sentinel Scheduled Alert Rule.
	Query pulumi.StringPtrInput
	// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
	QueryFrequency pulumi.StringPtrInput
	// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
	QueryPeriod pulumi.StringPtrInput
	// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
	Severity pulumi.StringPtrInput
	// If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
	SuppressionDuration pulumi.StringPtrInput
	// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
	SuppressionEnabled pulumi.BoolPtrInput
	// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
	Tactics pulumi.StringArrayInput
	// The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
	TriggerOperator pulumi.StringPtrInput
	// The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
	TriggerThreshold pulumi.IntPtrInput
}

func (AlertRuleScheduledState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleScheduledState)(nil)).Elem()
}

type alertRuleScheduledArgs struct {
	// The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	AlertRuleTemplateGuid *string `pulumi:"alertRuleTemplateGuid"`
	// The description of this Sentinel Scheduled Alert Rule.
	Description *string `pulumi:"description"`
	// The friendly name of this Sentinel Scheduled Alert Rule.
	DisplayName string `pulumi:"displayName"`
	// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
	Enabled *bool `pulumi:"enabled"`
	// A `eventGrouping` block as defined below.
	EventGrouping *AlertRuleScheduledEventGrouping `pulumi:"eventGrouping"`
	// A `incidentConfiguration` block as defined below.
	IncidentConfiguration *AlertRuleScheduledIncidentConfiguration `pulumi:"incidentConfiguration"`
	// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	LogAnalyticsWorkspaceId string `pulumi:"logAnalyticsWorkspaceId"`
	// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	Name *string `pulumi:"name"`
	// The query of this Sentinel Scheduled Alert Rule.
	Query string `pulumi:"query"`
	// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
	QueryFrequency *string `pulumi:"queryFrequency"`
	// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
	QueryPeriod *string `pulumi:"queryPeriod"`
	// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
	Severity string `pulumi:"severity"`
	// If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
	SuppressionDuration *string `pulumi:"suppressionDuration"`
	// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
	SuppressionEnabled *bool `pulumi:"suppressionEnabled"`
	// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
	Tactics []string `pulumi:"tactics"`
	// The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
	TriggerOperator *string `pulumi:"triggerOperator"`
	// The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
	TriggerThreshold *int `pulumi:"triggerThreshold"`
}

// The set of arguments for constructing a AlertRuleScheduled resource.
type AlertRuleScheduledArgs struct {
	// The GUID of the alert rule template which is used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	AlertRuleTemplateGuid pulumi.StringPtrInput
	// The description of this Sentinel Scheduled Alert Rule.
	Description pulumi.StringPtrInput
	// The friendly name of this Sentinel Scheduled Alert Rule.
	DisplayName pulumi.StringInput
	// Should the Sentinel Scheduled Alert Rule be enabled? Defaults to `true`.
	Enabled pulumi.BoolPtrInput
	// A `eventGrouping` block as defined below.
	EventGrouping AlertRuleScheduledEventGroupingPtrInput
	// A `incidentConfiguration` block as defined below.
	IncidentConfiguration AlertRuleScheduledIncidentConfigurationPtrInput
	// The ID of the Log Analytics Workspace this Sentinel Scheduled Alert Rule belongs to. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	LogAnalyticsWorkspaceId pulumi.StringInput
	// The name which should be used for this Sentinel Scheduled Alert Rule. Changing this forces a new Sentinel Scheduled Alert Rule to be created.
	Name pulumi.StringPtrInput
	// The query of this Sentinel Scheduled Alert Rule.
	Query pulumi.StringInput
	// The ISO 8601 timespan duration between two consecutive queries. Defaults to `PT5H`.
	QueryFrequency pulumi.StringPtrInput
	// The ISO 8601 timespan duration, which determine the time period of the data covered by the query. For example, it can query the past 10 minutes of data, or the past 6 hours of data. Defaults to `PT5H`.
	QueryPeriod pulumi.StringPtrInput
	// The alert severity of this Sentinel Scheduled Alert Rule. Possible values are `High`, `Medium`, `Low` and `Informational`.
	Severity pulumi.StringInput
	// If `suppressionEnabled` is `true`, this is ISO 8601 timespan duration, which specifies the amount of time the query should stop running after alert is generated. Defaults to `PT5H`.
	SuppressionDuration pulumi.StringPtrInput
	// Should the Sentinel Scheduled Alert Rulea stop running query after alert is generated? Defaults to `false`.
	SuppressionEnabled pulumi.BoolPtrInput
	// A list of categories of attacks by which to classify the rule. Possible values are `Collection`, `CommandAndControl`, `CredentialAccess`, `DefenseEvasion`, `Discovery`, `Execution`, `Exfiltration`, `Impact`, `InitialAccess`, `LateralMovement`, `Persistence` and `PrivilegeEscalation`.
	Tactics pulumi.StringArrayInput
	// The alert trigger operator, combined with `triggerThreshold`, setting alert threshold of this Sentinel Scheduled Alert Rule. Possible values are `Equal`, `GreaterThan`, `LessThan`, `NotEqual`.
	TriggerOperator pulumi.StringPtrInput
	// The baseline number of query results generated, combined with `triggerOperator`, setting alert threshold of this Sentinel Scheduled Alert Rule.
	TriggerThreshold pulumi.IntPtrInput
}

func (AlertRuleScheduledArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleScheduledArgs)(nil)).Elem()
}

type AlertRuleScheduledInput interface {
	pulumi.Input

	ToAlertRuleScheduledOutput() AlertRuleScheduledOutput
	ToAlertRuleScheduledOutputWithContext(ctx context.Context) AlertRuleScheduledOutput
}

func (*AlertRuleScheduled) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleScheduled)(nil))
}

func (i *AlertRuleScheduled) ToAlertRuleScheduledOutput() AlertRuleScheduledOutput {
	return i.ToAlertRuleScheduledOutputWithContext(context.Background())
}

func (i *AlertRuleScheduled) ToAlertRuleScheduledOutputWithContext(ctx context.Context) AlertRuleScheduledOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleScheduledOutput)
}

func (i *AlertRuleScheduled) ToAlertRuleScheduledPtrOutput() AlertRuleScheduledPtrOutput {
	return i.ToAlertRuleScheduledPtrOutputWithContext(context.Background())
}

func (i *AlertRuleScheduled) ToAlertRuleScheduledPtrOutputWithContext(ctx context.Context) AlertRuleScheduledPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleScheduledPtrOutput)
}

type AlertRuleScheduledPtrInput interface {
	pulumi.Input

	ToAlertRuleScheduledPtrOutput() AlertRuleScheduledPtrOutput
	ToAlertRuleScheduledPtrOutputWithContext(ctx context.Context) AlertRuleScheduledPtrOutput
}

type alertRuleScheduledPtrType AlertRuleScheduledArgs

func (*alertRuleScheduledPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleScheduled)(nil))
}

func (i *alertRuleScheduledPtrType) ToAlertRuleScheduledPtrOutput() AlertRuleScheduledPtrOutput {
	return i.ToAlertRuleScheduledPtrOutputWithContext(context.Background())
}

func (i *alertRuleScheduledPtrType) ToAlertRuleScheduledPtrOutputWithContext(ctx context.Context) AlertRuleScheduledPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleScheduledPtrOutput)
}

// AlertRuleScheduledArrayInput is an input type that accepts AlertRuleScheduledArray and AlertRuleScheduledArrayOutput values.
// You can construct a concrete instance of `AlertRuleScheduledArrayInput` via:
//
//          AlertRuleScheduledArray{ AlertRuleScheduledArgs{...} }
type AlertRuleScheduledArrayInput interface {
	pulumi.Input

	ToAlertRuleScheduledArrayOutput() AlertRuleScheduledArrayOutput
	ToAlertRuleScheduledArrayOutputWithContext(context.Context) AlertRuleScheduledArrayOutput
}

type AlertRuleScheduledArray []AlertRuleScheduledInput

func (AlertRuleScheduledArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AlertRuleScheduled)(nil))
}

func (i AlertRuleScheduledArray) ToAlertRuleScheduledArrayOutput() AlertRuleScheduledArrayOutput {
	return i.ToAlertRuleScheduledArrayOutputWithContext(context.Background())
}

func (i AlertRuleScheduledArray) ToAlertRuleScheduledArrayOutputWithContext(ctx context.Context) AlertRuleScheduledArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleScheduledArrayOutput)
}

// AlertRuleScheduledMapInput is an input type that accepts AlertRuleScheduledMap and AlertRuleScheduledMapOutput values.
// You can construct a concrete instance of `AlertRuleScheduledMapInput` via:
//
//          AlertRuleScheduledMap{ "key": AlertRuleScheduledArgs{...} }
type AlertRuleScheduledMapInput interface {
	pulumi.Input

	ToAlertRuleScheduledMapOutput() AlertRuleScheduledMapOutput
	ToAlertRuleScheduledMapOutputWithContext(context.Context) AlertRuleScheduledMapOutput
}

type AlertRuleScheduledMap map[string]AlertRuleScheduledInput

func (AlertRuleScheduledMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AlertRuleScheduled)(nil))
}

func (i AlertRuleScheduledMap) ToAlertRuleScheduledMapOutput() AlertRuleScheduledMapOutput {
	return i.ToAlertRuleScheduledMapOutputWithContext(context.Background())
}

func (i AlertRuleScheduledMap) ToAlertRuleScheduledMapOutputWithContext(ctx context.Context) AlertRuleScheduledMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleScheduledMapOutput)
}

type AlertRuleScheduledOutput struct {
	*pulumi.OutputState
}

func (AlertRuleScheduledOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRuleScheduled)(nil))
}

func (o AlertRuleScheduledOutput) ToAlertRuleScheduledOutput() AlertRuleScheduledOutput {
	return o
}

func (o AlertRuleScheduledOutput) ToAlertRuleScheduledOutputWithContext(ctx context.Context) AlertRuleScheduledOutput {
	return o
}

func (o AlertRuleScheduledOutput) ToAlertRuleScheduledPtrOutput() AlertRuleScheduledPtrOutput {
	return o.ToAlertRuleScheduledPtrOutputWithContext(context.Background())
}

func (o AlertRuleScheduledOutput) ToAlertRuleScheduledPtrOutputWithContext(ctx context.Context) AlertRuleScheduledPtrOutput {
	return o.ApplyT(func(v AlertRuleScheduled) *AlertRuleScheduled {
		return &v
	}).(AlertRuleScheduledPtrOutput)
}

type AlertRuleScheduledPtrOutput struct {
	*pulumi.OutputState
}

func (AlertRuleScheduledPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRuleScheduled)(nil))
}

func (o AlertRuleScheduledPtrOutput) ToAlertRuleScheduledPtrOutput() AlertRuleScheduledPtrOutput {
	return o
}

func (o AlertRuleScheduledPtrOutput) ToAlertRuleScheduledPtrOutputWithContext(ctx context.Context) AlertRuleScheduledPtrOutput {
	return o
}

type AlertRuleScheduledArrayOutput struct{ *pulumi.OutputState }

func (AlertRuleScheduledArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AlertRuleScheduled)(nil))
}

func (o AlertRuleScheduledArrayOutput) ToAlertRuleScheduledArrayOutput() AlertRuleScheduledArrayOutput {
	return o
}

func (o AlertRuleScheduledArrayOutput) ToAlertRuleScheduledArrayOutputWithContext(ctx context.Context) AlertRuleScheduledArrayOutput {
	return o
}

func (o AlertRuleScheduledArrayOutput) Index(i pulumi.IntInput) AlertRuleScheduledOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AlertRuleScheduled {
		return vs[0].([]AlertRuleScheduled)[vs[1].(int)]
	}).(AlertRuleScheduledOutput)
}

type AlertRuleScheduledMapOutput struct{ *pulumi.OutputState }

func (AlertRuleScheduledMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AlertRuleScheduled)(nil))
}

func (o AlertRuleScheduledMapOutput) ToAlertRuleScheduledMapOutput() AlertRuleScheduledMapOutput {
	return o
}

func (o AlertRuleScheduledMapOutput) ToAlertRuleScheduledMapOutputWithContext(ctx context.Context) AlertRuleScheduledMapOutput {
	return o
}

func (o AlertRuleScheduledMapOutput) MapIndex(k pulumi.StringInput) AlertRuleScheduledOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AlertRuleScheduled {
		return vs[0].(map[string]AlertRuleScheduled)[vs[1].(string)]
	}).(AlertRuleScheduledOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertRuleScheduledOutput{})
	pulumi.RegisterOutputType(AlertRuleScheduledPtrOutput{})
	pulumi.RegisterOutputType(AlertRuleScheduledArrayOutput{})
	pulumi.RegisterOutputType(AlertRuleScheduledMapOutput{})
}
