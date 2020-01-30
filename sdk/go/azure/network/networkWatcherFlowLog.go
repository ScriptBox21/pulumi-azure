// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Network Watcher Flow Log.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_watcher_flow_log.html.markdown.
type NetworkWatcherFlowLog struct {
	pulumi.CustomResourceState

	// Boolean flag to enable/disable traffic analytics.
	Enabled pulumi.BoolOutput `pulumi:"enabled"`
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringOutput `pulumi:"networkSecurityGroupId"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringOutput `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicyOutput `pulumi:"retentionPolicy"`
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics NetworkWatcherFlowLogTrafficAnalyticsPtrOutput `pulumi:"trafficAnalytics"`
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewNetworkWatcherFlowLog registers a new resource with the given unique name, arguments, and options.
func NewNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, args *NetworkWatcherFlowLogArgs, opts ...pulumi.ResourceOption) (*NetworkWatcherFlowLog, error) {
	if args == nil || args.Enabled == nil {
		return nil, errors.New("missing required argument 'Enabled'")
	}
	if args == nil || args.NetworkSecurityGroupId == nil {
		return nil, errors.New("missing required argument 'NetworkSecurityGroupId'")
	}
	if args == nil || args.NetworkWatcherName == nil {
		return nil, errors.New("missing required argument 'NetworkWatcherName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.RetentionPolicy == nil {
		return nil, errors.New("missing required argument 'RetentionPolicy'")
	}
	if args == nil || args.StorageAccountId == nil {
		return nil, errors.New("missing required argument 'StorageAccountId'")
	}
	if args == nil {
		args = &NetworkWatcherFlowLogArgs{}
	}
	var resource NetworkWatcherFlowLog
	err := ctx.RegisterResource("azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkWatcherFlowLog gets an existing NetworkWatcherFlowLog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkWatcherFlowLog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkWatcherFlowLogState, opts ...pulumi.ResourceOption) (*NetworkWatcherFlowLog, error) {
	var resource NetworkWatcherFlowLog
	err := ctx.ReadResource("azure:network/networkWatcherFlowLog:NetworkWatcherFlowLog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkWatcherFlowLog resources.
type networkWatcherFlowLogState struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled *bool `pulumi:"enabled"`
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId *string `pulumi:"networkSecurityGroupId"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName *string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy *NetworkWatcherFlowLogRetentionPolicy `pulumi:"retentionPolicy"`
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics *NetworkWatcherFlowLogTrafficAnalytics `pulumi:"trafficAnalytics"`
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version *int `pulumi:"version"`
}

type NetworkWatcherFlowLogState struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled pulumi.BoolPtrInput
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringPtrInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringPtrInput
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicyPtrInput
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId pulumi.StringPtrInput
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics NetworkWatcherFlowLogTrafficAnalyticsPtrInput
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version pulumi.IntPtrInput
}

func (NetworkWatcherFlowLogState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherFlowLogState)(nil)).Elem()
}

type networkWatcherFlowLogArgs struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled bool `pulumi:"enabled"`
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId string `pulumi:"networkSecurityGroupId"`
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName string `pulumi:"networkWatcherName"`
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicy `pulumi:"retentionPolicy"`
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId string `pulumi:"storageAccountId"`
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics *NetworkWatcherFlowLogTrafficAnalytics `pulumi:"trafficAnalytics"`
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a NetworkWatcherFlowLog resource.
type NetworkWatcherFlowLogArgs struct {
	// Boolean flag to enable/disable traffic analytics.
	Enabled pulumi.BoolInput
	// The ID of the Network Security Group for which to enable flow logs for. Changing this forces a new resource to be created.
	NetworkSecurityGroupId pulumi.StringInput
	// The name of the Network Watcher. Changing this forces a new resource to be created.
	NetworkWatcherName pulumi.StringInput
	// The name of the resource group in which the Network Watcher was deployed. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `retentionPolicy` block as documented below.
	RetentionPolicy NetworkWatcherFlowLogRetentionPolicyInput
	// The ID of the Storage Account where flow logs are stored.
	StorageAccountId pulumi.StringInput
	// A `trafficAnalytics` block as documented below.
	TrafficAnalytics NetworkWatcherFlowLogTrafficAnalyticsPtrInput
	// The version (revision) of the flow log. Possible values are `1` and `2`.
	Version pulumi.IntPtrInput
}

func (NetworkWatcherFlowLogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkWatcherFlowLogArgs)(nil)).Elem()
}

