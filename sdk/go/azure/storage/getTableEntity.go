// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to access information about an existing Storage Table Entity.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storage.LookupTableEntity(ctx, &storage.LookupTableEntityArgs{
// 			PartitionKey:       "example-parition-key",
// 			RowKey:             "example-row-key",
// 			StorageAccountName: "example-storage-account-name",
// 			TableName:          "example-table-name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupTableEntity(ctx *pulumi.Context, args *LookupTableEntityArgs, opts ...pulumi.InvokeOption) (*LookupTableEntityResult, error) {
	var rv LookupTableEntityResult
	err := ctx.Invoke("azure:storage/getTableEntity:getTableEntity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTableEntity.
type LookupTableEntityArgs struct {
	// The key for the partition where the entity will be retrieved.
	PartitionKey string `pulumi:"partitionKey"`
	// The key for the row where the entity will be retrieved.
	RowKey string `pulumi:"rowKey"`
	// The name of the Storage Account where the Table exists.
	StorageAccountName string `pulumi:"storageAccountName"`
	// The name of the Table.
	TableName string `pulumi:"tableName"`
}

// A collection of values returned by getTableEntity.
type LookupTableEntityResult struct {
	// A map of key/value pairs that describe the entity to be stored in the storage table.
	Entity map[string]string `pulumi:"entity"`
	// The provider-assigned unique ID for this managed resource.
	Id                 string `pulumi:"id"`
	PartitionKey       string `pulumi:"partitionKey"`
	RowKey             string `pulumi:"rowKey"`
	StorageAccountName string `pulumi:"storageAccountName"`
	TableName          string `pulumi:"tableName"`
}
