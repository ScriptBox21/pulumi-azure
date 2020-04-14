// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a SQL Container within a Cosmos DB Account.
type SqlContainer struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntOutput `pulumi:"defaultTtl"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrOutput `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntOutput `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayOutput `pulumi:"uniqueKeys"`
}

// NewSqlContainer registers a new resource with the given unique name, arguments, and options.
func NewSqlContainer(ctx *pulumi.Context,
	name string, args *SqlContainerArgs, opts ...pulumi.ResourceOption) (*SqlContainer, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.DatabaseName == nil {
		return nil, errors.New("missing required argument 'DatabaseName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &SqlContainerArgs{}
	}
	var resource SqlContainer
	err := ctx.RegisterResource("azure:cosmosdb/sqlContainer:SqlContainer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSqlContainer gets an existing SqlContainer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSqlContainer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SqlContainerState, opts ...pulumi.ResourceOption) (*SqlContainer, error) {
	var resource SqlContainer
	err := ctx.ReadResource("azure:cosmosdb/sqlContainer:SqlContainer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SqlContainer resources.
type sqlContainerState struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName *string `pulumi:"accountName"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName *string `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []SqlContainerUniqueKey `pulumi:"uniqueKeys"`
}

type SqlContainerState struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName pulumi.StringPtrInput
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringPtrInput
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntPtrInput
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayInput
}

func (SqlContainerState) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlContainerState)(nil)).Elem()
}

type sqlContainerArgs struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName string `pulumi:"accountName"`
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName string `pulumi:"databaseName"`
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl *int `pulumi:"defaultTtl"`
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath *string `pulumi:"partitionKeyPath"`
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput *int `pulumi:"throughput"`
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys []SqlContainerUniqueKey `pulumi:"uniqueKeys"`
}

// The set of arguments for constructing a SqlContainer resource.
type SqlContainerArgs struct {
	// The name of the Cosmos DB Account to create the container within. Changing this forces a new resource to be created.
	AccountName pulumi.StringInput
	// The name of the Cosmos DB SQL Database to create the container within. Changing this forces a new resource to be created.
	DatabaseName pulumi.StringInput
	// The default time to live of SQL container. If missing, items are not expired automatically. If present and the value is set to `-1`, it is equal to infinity, and items don’t expire by default. If present and the value is set to some number `n` – items will expire `n` seconds after their last modified time.
	DefaultTtl pulumi.IntPtrInput
	// Specifies the name of the Cosmos DB SQL Database. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Define a partition key. Changing this forces a new resource to be created.
	PartitionKeyPath pulumi.StringPtrInput
	// The name of the resource group in which the Cosmos DB SQL Database is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The throughput of SQL container (RU/s). Must be set in increments of `100`. The minimum value is `400`. This must be set upon container creation otherwise it cannot be updated without a manual resource destroy-apply.
	Throughput pulumi.IntPtrInput
	// One or more `uniqueKey` blocks as defined below. Changing this forces a new resource to be created.
	UniqueKeys SqlContainerUniqueKeyArrayInput
}

func (SqlContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sqlContainerArgs)(nil)).Elem()
}
