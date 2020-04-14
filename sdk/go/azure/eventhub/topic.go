// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package eventhub

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a ServiceBus Topic.
//
// **Note** Topics can only be created in Namespaces with an SKU of `standard` or higher.
type Topic struct {
	pulumi.CustomResourceState

	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringOutput `pulumi:"autoDeleteOnIdle"`
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl pulumi.StringOutput `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringOutput `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations pulumi.BoolPtrOutput `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress pulumi.BoolPtrOutput `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning pulumi.BoolPtrOutput `pulumi:"enablePartitioning"`
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntOutput `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringOutput `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection pulumi.BoolPtrOutput `pulumi:"requiresDuplicateDetection"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering pulumi.BoolPtrOutput `pulumi:"supportOrdering"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &TopicArgs{}
	}
	var resource Topic
	err := ctx.RegisterResource("azure:eventhub/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azure:eventhub/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName *string `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status *string `pulumi:"status"`
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering *bool `pulumi:"supportOrdering"`
}

type TopicState struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning pulumi.BoolPtrInput
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringPtrInput
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrInput
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering pulumi.BoolPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle *string `pulumi:"autoDeleteOnIdle"`
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl *string `pulumi:"defaultMessageTtl"`
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow *string `pulumi:"duplicateDetectionHistoryTimeWindow"`
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations *bool `pulumi:"enableBatchedOperations"`
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress *bool `pulumi:"enableExpress"`
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning *bool `pulumi:"enablePartitioning"`
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes *int `pulumi:"maxSizeInMegabytes"`
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName string `pulumi:"namespaceName"`
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection *bool `pulumi:"requiresDuplicateDetection"`
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status *string `pulumi:"status"`
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering *bool `pulumi:"supportOrdering"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// The ISO 8601 timespan duration of the idle interval after which the
	// Topic is automatically deleted, minimum of 5 minutes.
	AutoDeleteOnIdle pulumi.StringPtrInput
	// The ISO 8601 timespan duration of TTL of messages sent to this topic if no
	// TTL value is set on the message itself.
	DefaultMessageTtl pulumi.StringPtrInput
	// The ISO 8601 timespan duration during which
	// duplicates can be detected. Defaults to 10 minutes. (`PT10M`)
	DuplicateDetectionHistoryTimeWindow pulumi.StringPtrInput
	// Boolean flag which controls if server-side
	// batched operations are enabled. Defaults to false.
	EnableBatchedOperations pulumi.BoolPtrInput
	// Boolean flag which controls whether Express Entities
	// are enabled. An express topic holds a message in memory temporarily before writing
	// it to persistent storage. Defaults to false.
	EnableExpress pulumi.BoolPtrInput
	// Boolean flag which controls whether to enable
	// the topic to be partitioned across multiple message brokers. Defaults to false.
	// Changing this forces a new resource to be created.
	EnablePartitioning pulumi.BoolPtrInput
	// Integer value which controls the size of
	// memory allocated for the topic. For supported values see the "Queue/topic size"
	// section of [this document](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas).
	MaxSizeInMegabytes pulumi.IntPtrInput
	// Specifies the name of the ServiceBus Topic resource. Changing this forces a
	// new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the ServiceBus Namespace to create
	// this topic in. Changing this forces a new resource to be created.
	NamespaceName pulumi.StringInput
	// Boolean flag which controls whether
	// the Topic requires duplicate detection. Defaults to false. Changing this forces
	// a new resource to be created.
	RequiresDuplicateDetection pulumi.BoolPtrInput
	// The name of the resource group in which to
	// create the namespace. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The Status of the Service Bus Topic. Acceptable values are `Active` or `Disabled`. Defaults to `Active`.
	Status pulumi.StringPtrInput
	// Boolean flag which controls whether the Topic
	// supports ordering. Defaults to false.
	SupportOrdering pulumi.BoolPtrInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}
