// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Event Hubs as a nested resource within a Event Hubs namespace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/eventhub.html.markdown.
type EventHub struct {
	pulumi.CustomResourceState

	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionPtrOutput `pulumi:"captureDescription"`
	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention pulumi.IntOutput `pulumi:"messageRetention"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntOutput `pulumi:"partitionCount"`
	// The identifiers for partitions created for Event Hubs.
	PartitionIds pulumi.StringArrayOutput `pulumi:"partitionIds"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewEventHub registers a new resource with the given unique name, arguments, and options.
func NewEventHub(ctx *pulumi.Context,
	name string, args *EventHubArgs, opts ...pulumi.ResourceOption) (*EventHub, error) {
	if args == nil || args.MessageRetention == nil {
		return nil, errors.New("missing required argument 'MessageRetention'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.PartitionCount == nil {
		return nil, errors.New("missing required argument 'PartitionCount'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventHubArgs{}
	}
	var resource EventHub
	err := ctx.RegisterResource("azure:eventhub/eventHub:EventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHub gets an existing EventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubState, opts ...pulumi.ResourceOption) (*EventHub, error) {
	var resource EventHub
	err := ctx.ReadResource("azure:eventhub/eventHub:EventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHub resources.
type eventHubState struct {
	// A `captureDescription` block as defined below.
	CaptureDescription *EventHubCaptureDescription `pulumi:"captureDescription"`
	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention *int `pulumi:"messageRetention"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount *int `pulumi:"partitionCount"`
	// The identifiers for partitions created for Event Hubs.
	PartitionIds []string `pulumi:"partitionIds"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type EventHubState struct {
	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionPtrInput
	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention pulumi.IntPtrInput
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntPtrInput
	// The identifiers for partitions created for Event Hubs.
	PartitionIds pulumi.StringArrayInput
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (EventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubState)(nil)).Elem()
}

type eventHubArgs struct {
	// A `captureDescription` block as defined below.
	CaptureDescription *EventHubCaptureDescription `pulumi:"captureDescription"`
	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention int `pulumi:"messageRetention"`
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount int `pulumi:"partitionCount"`
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a EventHub resource.
type EventHubArgs struct {
	// A `captureDescription` block as defined below.
	CaptureDescription EventHubCaptureDescriptionPtrInput
	// Specifies the number of days to retain the events for this Event Hub. Needs to be between 1 and 7 days; or 1 day when using a Basic SKU for the parent EventHub Namespace.
	MessageRetention pulumi.IntInput
	// Specifies the name of the EventHub Namespace resource. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the EventHub Namespace. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// Specifies the current number of shards on the Event Hub. Changing this forces a new resource to be created.
	PartitionCount pulumi.IntInput
	// The name of the resource group in which the EventHub's parent Namespace exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (EventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubArgs)(nil)).Elem()
}
