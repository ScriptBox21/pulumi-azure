// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Queue within an Azure Storage Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_queue.html.markdown.
type Queue struct {
	pulumi.CustomResourceState

	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringOutput `pulumi:"storageAccountName"`
}

// NewQueue registers a new resource with the given unique name, arguments, and options.
func NewQueue(ctx *pulumi.Context,
	name string, args *QueueArgs, opts ...pulumi.ResourceOption) (*Queue, error) {
	if args == nil || args.StorageAccountName == nil {
		return nil, errors.New("missing required argument 'StorageAccountName'")
	}
	if args == nil {
		args = &QueueArgs{}
	}
	var resource Queue
	err := ctx.RegisterResource("azure:storage/queue:Queue", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQueue gets an existing Queue resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQueue(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QueueState, opts ...pulumi.ResourceOption) (*Queue, error) {
	var resource Queue
	err := ctx.ReadResource("azure:storage/queue:Queue", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Queue resources.
type queueState struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name *string `pulumi:"name"`
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName *string `pulumi:"storageAccountName"`
}

type QueueState struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata pulumi.StringMapInput
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name pulumi.StringPtrInput
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringPtrInput
}

func (QueueState) ElementType() reflect.Type {
	return reflect.TypeOf((*queueState)(nil)).Elem()
}

type queueArgs struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name *string `pulumi:"name"`
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName string `pulumi:"storageAccountName"`
}

// The set of arguments for constructing a Queue resource.
type QueueArgs struct {
	// A mapping of MetaData which should be assigned to this Storage Queue.
	Metadata pulumi.StringMapInput
	// The name of the Queue which should be created within the Storage Account. Must be unique within the storage account the queue is located.
	Name pulumi.StringPtrInput
	// Specifies the Storage Account in which the Storage Queue should exist. Changing this forces a new resource to be created.
	StorageAccountName pulumi.StringInput
}

func (QueueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*queueArgs)(nil)).Elem()
}
