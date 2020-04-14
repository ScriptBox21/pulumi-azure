// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) EventHub Data Connection
type EventhubDataConnection struct {
	pulumi.CustomResourceState

	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringOutput `pulumi:"consumerGroup"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrOutput `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId pulumi.StringOutput `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrOutput `pulumi:"mappingRuleName"`
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrOutput `pulumi:"tableName"`
}

// NewEventhubDataConnection registers a new resource with the given unique name, arguments, and options.
func NewEventhubDataConnection(ctx *pulumi.Context,
	name string, args *EventhubDataConnectionArgs, opts ...pulumi.ResourceOption) (*EventhubDataConnection, error) {
	if args == nil || args.ClusterName == nil {
		return nil, errors.New("missing required argument 'ClusterName'")
	}
	if args == nil || args.ConsumerGroup == nil {
		return nil, errors.New("missing required argument 'ConsumerGroup'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.EventhubId == nil {
		return nil, errors.New("missing required argument 'EventhubId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventhubDataConnectionArgs{}
	}
	var resource EventhubDataConnection
	err := ctx.RegisterResource("azure:kusto/eventhubDataConnection:EventhubDataConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventhubDataConnection gets an existing EventhubDataConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventhubDataConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventhubDataConnectionState, opts ...pulumi.ResourceOption) (*EventhubDataConnection, error) {
	var resource EventhubDataConnection
	err := ctx.ReadResource("azure:kusto/eventhubDataConnection:EventhubDataConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventhubDataConnection resources.
type eventhubDataConnectionState struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup *string `pulumi:"consumerGroup"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId *string `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

type EventhubDataConnectionState struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringPtrInput
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId pulumi.StringPtrInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (EventhubDataConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubDataConnectionState)(nil)).Elem()
}

type eventhubDataConnectionArgs struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup string `pulumi:"consumerGroup"`
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat *string `pulumi:"dataFormat"`
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId string `pulumi:"eventhubId"`
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName *string `pulumi:"mappingRuleName"`
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName *string `pulumi:"tableName"`
}

// The set of arguments for constructing a EventhubDataConnection resource.
type EventhubDataConnectionArgs struct {
	// Specifies the name of the Kusto Cluster this data connection will be added to. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// Specifies the EventHub consumer group this data connection will use for ingestion. Changing this forces a new resource to be created.
	ConsumerGroup pulumi.StringInput
	// Specifies the data format of the EventHub messages. Allowed values: `AVRO`, `CSV`, `JSON`, `MULTIJSON`, `PSV`, `RAW`, `SCSV`, `SINGLEJSON`, `SOHSV`, `TSV` and `TXT`
	DataFormat pulumi.StringPtrInput
	// Specifies the name of the Kusto Database this data connection will be added to. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// Specifies the resource id of the EventHub this data connection will use for ingestion. Changing this forces a new resource to be created.
	EventhubId pulumi.StringInput
	// The location where the Kusto Database should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the mapping rule used for the message ingestion. Mapping rule must exist before resource is created.
	MappingRuleName pulumi.StringPtrInput
	// The name of the Kusto EventHub Data Connection to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Database should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the target table name used for the message ingestion. Table must exist before resource is created.
	TableName pulumi.StringPtrInput
}

func (EventhubDataConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventhubDataConnectionArgs)(nil)).Elem()
}
