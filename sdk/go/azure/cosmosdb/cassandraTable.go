// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Cassandra Table within a Cosmos DB Cassandra Keyspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.LookupResourceGroup(ctx, &core.LookupResourceGroupArgs{
// 			Name: "tflex-cosmosdb-account-rg",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleAccount, err := cosmosdb.NewAccount(ctx, "exampleAccount", &cosmosdb.AccountArgs{
// 			ResourceGroupName: pulumi.String(exampleResourceGroup.Name),
// 			Location:          pulumi.String(exampleResourceGroup.Location),
// 			OfferType:         pulumi.String("Standard"),
// 			Capabilities: cosmosdb.AccountCapabilityArray{
// 				&cosmosdb.AccountCapabilityArgs{
// 					Name: pulumi.String("EnableCassandra"),
// 				},
// 			},
// 			ConsistencyPolicy: &cosmosdb.AccountConsistencyPolicyArgs{
// 				ConsistencyLevel: pulumi.String("Strong"),
// 			},
// 			GeoLocations: cosmosdb.AccountGeoLocationArray{
// 				&cosmosdb.AccountGeoLocationArgs{
// 					Location:         pulumi.String("West US"),
// 					FailoverPriority: pulumi.Int(0),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleCassandraKeyspace, err := cosmosdb.NewCassandraKeyspace(ctx, "exampleCassandraKeyspace", &cosmosdb.CassandraKeyspaceArgs{
// 			ResourceGroupName: exampleAccount.ResourceGroupName,
// 			AccountName:       exampleAccount.Name,
// 			Throughput:        pulumi.Int(400),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewCassandraTable(ctx, "exampleCassandraTable", &cosmosdb.CassandraTableArgs{
// 			CassandraKeyspaceId: exampleCassandraKeyspace.ID(),
// 			Schema: &cosmosdb.CassandraTableSchemaArgs{
// 				Columns: cosmosdb.CassandraTableSchemaColumnArray{
// 					&cosmosdb.CassandraTableSchemaColumnArgs{
// 						Name: pulumi.String("test1"),
// 						Type: pulumi.String("ascii"),
// 					},
// 					&cosmosdb.CassandraTableSchemaColumnArgs{
// 						Name: pulumi.String("test2"),
// 						Type: pulumi.String("int"),
// 					},
// 				},
// 				PartitionKeys: cosmosdb.CassandraTableSchemaPartitionKeyArray{
// 					&cosmosdb.CassandraTableSchemaPartitionKeyArgs{
// 						Name: pulumi.String("test1"),
// 					},
// 				},
// 			},
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
// Cosmos Cassandra Table can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/cassandraTable:CassandraTable ks1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/account1/cassandraKeyspaces/ks1/tables/table1
// ```
type CassandraTable struct {
	pulumi.CustomResourceState

	AutoscaleSettings CassandraTableAutoscaleSettingsPtrOutput `pulumi:"autoscaleSettings"`
	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	CassandraKeyspaceId pulumi.StringOutput `pulumi:"cassandraKeyspaceId"`
	DefaultTtl          pulumi.IntOutput    `pulumi:"defaultTtl"`
	// Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `schema` block as defined below. Changing this forces a new resource to be created.
	Schema     CassandraTableSchemaOutput `pulumi:"schema"`
	Throughput pulumi.IntOutput           `pulumi:"throughput"`
}

// NewCassandraTable registers a new resource with the given unique name, arguments, and options.
func NewCassandraTable(ctx *pulumi.Context,
	name string, args *CassandraTableArgs, opts ...pulumi.ResourceOption) (*CassandraTable, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CassandraKeyspaceId == nil {
		return nil, errors.New("invalid value for required argument 'CassandraKeyspaceId'")
	}
	if args.Schema == nil {
		return nil, errors.New("invalid value for required argument 'Schema'")
	}
	var resource CassandraTable
	err := ctx.RegisterResource("azure:cosmosdb/cassandraTable:CassandraTable", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCassandraTable gets an existing CassandraTable resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCassandraTable(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CassandraTableState, opts ...pulumi.ResourceOption) (*CassandraTable, error) {
	var resource CassandraTable
	err := ctx.ReadResource("azure:cosmosdb/cassandraTable:CassandraTable", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CassandraTable resources.
type cassandraTableState struct {
	AutoscaleSettings *CassandraTableAutoscaleSettings `pulumi:"autoscaleSettings"`
	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	CassandraKeyspaceId *string `pulumi:"cassandraKeyspaceId"`
	DefaultTtl          *int    `pulumi:"defaultTtl"`
	// Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `schema` block as defined below. Changing this forces a new resource to be created.
	Schema     *CassandraTableSchema `pulumi:"schema"`
	Throughput *int                  `pulumi:"throughput"`
}

type CassandraTableState struct {
	AutoscaleSettings CassandraTableAutoscaleSettingsPtrInput
	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	CassandraKeyspaceId pulumi.StringPtrInput
	DefaultTtl          pulumi.IntPtrInput
	// Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `schema` block as defined below. Changing this forces a new resource to be created.
	Schema     CassandraTableSchemaPtrInput
	Throughput pulumi.IntPtrInput
}

func (CassandraTableState) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraTableState)(nil)).Elem()
}

type cassandraTableArgs struct {
	AutoscaleSettings *CassandraTableAutoscaleSettings `pulumi:"autoscaleSettings"`
	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	CassandraKeyspaceId string `pulumi:"cassandraKeyspaceId"`
	DefaultTtl          *int   `pulumi:"defaultTtl"`
	// Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A `schema` block as defined below. Changing this forces a new resource to be created.
	Schema     CassandraTableSchema `pulumi:"schema"`
	Throughput *int                 `pulumi:"throughput"`
}

// The set of arguments for constructing a CassandraTable resource.
type CassandraTableArgs struct {
	AutoscaleSettings CassandraTableAutoscaleSettingsPtrInput
	// The ID of the Cosmos DB Cassandra Keyspace to create the table within. Changing this forces a new resource to be created.
	CassandraKeyspaceId pulumi.StringInput
	DefaultTtl          pulumi.IntPtrInput
	// Specifies the name of the Cosmos DB Cassandra Table. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A `schema` block as defined below. Changing this forces a new resource to be created.
	Schema     CassandraTableSchemaInput
	Throughput pulumi.IntPtrInput
}

func (CassandraTableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cassandraTableArgs)(nil)).Elem()
}

type CassandraTableInput interface {
	pulumi.Input

	ToCassandraTableOutput() CassandraTableOutput
	ToCassandraTableOutputWithContext(ctx context.Context) CassandraTableOutput
}

func (CassandraTable) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTable)(nil)).Elem()
}

func (i CassandraTable) ToCassandraTableOutput() CassandraTableOutput {
	return i.ToCassandraTableOutputWithContext(context.Background())
}

func (i CassandraTable) ToCassandraTableOutputWithContext(ctx context.Context) CassandraTableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableOutput)
}

type CassandraTableOutput struct {
	*pulumi.OutputState
}

func (CassandraTableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableOutput)(nil)).Elem()
}

func (o CassandraTableOutput) ToCassandraTableOutput() CassandraTableOutput {
	return o
}

func (o CassandraTableOutput) ToCassandraTableOutputWithContext(ctx context.Context) CassandraTableOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(CassandraTableOutput{})
}
