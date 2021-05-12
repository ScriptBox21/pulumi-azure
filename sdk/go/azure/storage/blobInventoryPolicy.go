// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a Storage Blob Inventory Policy.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			BlobProperties: &storage.AccountBlobPropertiesArgs{
// 				VersioningEnabled: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBlobInventoryPolicy(ctx, "exampleBlobInventoryPolicy", &storage.BlobInventoryPolicyArgs{
// 			StorageAccountId:     exampleAccount.ID(),
// 			StorageContainerName: exampleContainer.Name,
// 			Rules: storage.BlobInventoryPolicyRuleArray{
// 				&storage.BlobInventoryPolicyRuleArgs{
// 					Name: pulumi.String("rule1"),
// 					Filter: &storage.BlobInventoryPolicyRuleFilterArgs{
// 						BlobTypes: pulumi.StringArray{
// 							pulumi.String("blockBlob"),
// 						},
// 						IncludeBlobVersions: pulumi.Bool(true),
// 						IncludeSnapshots:    pulumi.Bool(true),
// 						PrefixMatches: pulumi.StringArray{
// 							pulumi.String("*/example"),
// 						},
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
// Storage Blob Inventory Policys can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:storage/blobInventoryPolicy:BlobInventoryPolicy example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/storageAccount1/inventoryPolicies/inventoryPolicy1
// ```
type BlobInventoryPolicy struct {
	pulumi.CustomResourceState

	// One or more `rules` blocks as defined below.
	Rules BlobInventoryPolicyRuleArrayOutput `pulumi:"rules"`
	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageAccountId pulumi.StringOutput `pulumi:"storageAccountId"`
	// The storage container name to store the blob inventory files. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageContainerName pulumi.StringOutput `pulumi:"storageContainerName"`
}

// NewBlobInventoryPolicy registers a new resource with the given unique name, arguments, and options.
func NewBlobInventoryPolicy(ctx *pulumi.Context,
	name string, args *BlobInventoryPolicyArgs, opts ...pulumi.ResourceOption) (*BlobInventoryPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.StorageAccountId == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountId'")
	}
	if args.StorageContainerName == nil {
		return nil, errors.New("invalid value for required argument 'StorageContainerName'")
	}
	var resource BlobInventoryPolicy
	err := ctx.RegisterResource("azure:storage/blobInventoryPolicy:BlobInventoryPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobInventoryPolicy gets an existing BlobInventoryPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobInventoryPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobInventoryPolicyState, opts ...pulumi.ResourceOption) (*BlobInventoryPolicy, error) {
	var resource BlobInventoryPolicy
	err := ctx.ReadResource("azure:storage/blobInventoryPolicy:BlobInventoryPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobInventoryPolicy resources.
type blobInventoryPolicyState struct {
	// One or more `rules` blocks as defined below.
	Rules []BlobInventoryPolicyRule `pulumi:"rules"`
	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageAccountId *string `pulumi:"storageAccountId"`
	// The storage container name to store the blob inventory files. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageContainerName *string `pulumi:"storageContainerName"`
}

type BlobInventoryPolicyState struct {
	// One or more `rules` blocks as defined below.
	Rules BlobInventoryPolicyRuleArrayInput
	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageAccountId pulumi.StringPtrInput
	// The storage container name to store the blob inventory files. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageContainerName pulumi.StringPtrInput
}

func (BlobInventoryPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobInventoryPolicyState)(nil)).Elem()
}

type blobInventoryPolicyArgs struct {
	// One or more `rules` blocks as defined below.
	Rules []BlobInventoryPolicyRule `pulumi:"rules"`
	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageAccountId string `pulumi:"storageAccountId"`
	// The storage container name to store the blob inventory files. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageContainerName string `pulumi:"storageContainerName"`
}

// The set of arguments for constructing a BlobInventoryPolicy resource.
type BlobInventoryPolicyArgs struct {
	// One or more `rules` blocks as defined below.
	Rules BlobInventoryPolicyRuleArrayInput
	// The ID of the storage account to apply this Blob Inventory Policy to. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageAccountId pulumi.StringInput
	// The storage container name to store the blob inventory files. Changing this forces a new Storage Blob Inventory Policy to be created.
	StorageContainerName pulumi.StringInput
}

func (BlobInventoryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobInventoryPolicyArgs)(nil)).Elem()
}

type BlobInventoryPolicyInput interface {
	pulumi.Input

	ToBlobInventoryPolicyOutput() BlobInventoryPolicyOutput
	ToBlobInventoryPolicyOutputWithContext(ctx context.Context) BlobInventoryPolicyOutput
}

func (*BlobInventoryPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicy)(nil))
}

func (i *BlobInventoryPolicy) ToBlobInventoryPolicyOutput() BlobInventoryPolicyOutput {
	return i.ToBlobInventoryPolicyOutputWithContext(context.Background())
}

func (i *BlobInventoryPolicy) ToBlobInventoryPolicyOutputWithContext(ctx context.Context) BlobInventoryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyOutput)
}

func (i *BlobInventoryPolicy) ToBlobInventoryPolicyPtrOutput() BlobInventoryPolicyPtrOutput {
	return i.ToBlobInventoryPolicyPtrOutputWithContext(context.Background())
}

func (i *BlobInventoryPolicy) ToBlobInventoryPolicyPtrOutputWithContext(ctx context.Context) BlobInventoryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyPtrOutput)
}

type BlobInventoryPolicyPtrInput interface {
	pulumi.Input

	ToBlobInventoryPolicyPtrOutput() BlobInventoryPolicyPtrOutput
	ToBlobInventoryPolicyPtrOutputWithContext(ctx context.Context) BlobInventoryPolicyPtrOutput
}

type blobInventoryPolicyPtrType BlobInventoryPolicyArgs

func (*blobInventoryPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicy)(nil))
}

func (i *blobInventoryPolicyPtrType) ToBlobInventoryPolicyPtrOutput() BlobInventoryPolicyPtrOutput {
	return i.ToBlobInventoryPolicyPtrOutputWithContext(context.Background())
}

func (i *blobInventoryPolicyPtrType) ToBlobInventoryPolicyPtrOutputWithContext(ctx context.Context) BlobInventoryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyPtrOutput)
}

// BlobInventoryPolicyArrayInput is an input type that accepts BlobInventoryPolicyArray and BlobInventoryPolicyArrayOutput values.
// You can construct a concrete instance of `BlobInventoryPolicyArrayInput` via:
//
//          BlobInventoryPolicyArray{ BlobInventoryPolicyArgs{...} }
type BlobInventoryPolicyArrayInput interface {
	pulumi.Input

	ToBlobInventoryPolicyArrayOutput() BlobInventoryPolicyArrayOutput
	ToBlobInventoryPolicyArrayOutputWithContext(context.Context) BlobInventoryPolicyArrayOutput
}

type BlobInventoryPolicyArray []BlobInventoryPolicyInput

func (BlobInventoryPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*BlobInventoryPolicy)(nil))
}

func (i BlobInventoryPolicyArray) ToBlobInventoryPolicyArrayOutput() BlobInventoryPolicyArrayOutput {
	return i.ToBlobInventoryPolicyArrayOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyArray) ToBlobInventoryPolicyArrayOutputWithContext(ctx context.Context) BlobInventoryPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyArrayOutput)
}

// BlobInventoryPolicyMapInput is an input type that accepts BlobInventoryPolicyMap and BlobInventoryPolicyMapOutput values.
// You can construct a concrete instance of `BlobInventoryPolicyMapInput` via:
//
//          BlobInventoryPolicyMap{ "key": BlobInventoryPolicyArgs{...} }
type BlobInventoryPolicyMapInput interface {
	pulumi.Input

	ToBlobInventoryPolicyMapOutput() BlobInventoryPolicyMapOutput
	ToBlobInventoryPolicyMapOutputWithContext(context.Context) BlobInventoryPolicyMapOutput
}

type BlobInventoryPolicyMap map[string]BlobInventoryPolicyInput

func (BlobInventoryPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*BlobInventoryPolicy)(nil))
}

func (i BlobInventoryPolicyMap) ToBlobInventoryPolicyMapOutput() BlobInventoryPolicyMapOutput {
	return i.ToBlobInventoryPolicyMapOutputWithContext(context.Background())
}

func (i BlobInventoryPolicyMap) ToBlobInventoryPolicyMapOutputWithContext(ctx context.Context) BlobInventoryPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BlobInventoryPolicyMapOutput)
}

type BlobInventoryPolicyOutput struct {
	*pulumi.OutputState
}

func (BlobInventoryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BlobInventoryPolicy)(nil))
}

func (o BlobInventoryPolicyOutput) ToBlobInventoryPolicyOutput() BlobInventoryPolicyOutput {
	return o
}

func (o BlobInventoryPolicyOutput) ToBlobInventoryPolicyOutputWithContext(ctx context.Context) BlobInventoryPolicyOutput {
	return o
}

func (o BlobInventoryPolicyOutput) ToBlobInventoryPolicyPtrOutput() BlobInventoryPolicyPtrOutput {
	return o.ToBlobInventoryPolicyPtrOutputWithContext(context.Background())
}

func (o BlobInventoryPolicyOutput) ToBlobInventoryPolicyPtrOutputWithContext(ctx context.Context) BlobInventoryPolicyPtrOutput {
	return o.ApplyT(func(v BlobInventoryPolicy) *BlobInventoryPolicy {
		return &v
	}).(BlobInventoryPolicyPtrOutput)
}

type BlobInventoryPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (BlobInventoryPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BlobInventoryPolicy)(nil))
}

func (o BlobInventoryPolicyPtrOutput) ToBlobInventoryPolicyPtrOutput() BlobInventoryPolicyPtrOutput {
	return o
}

func (o BlobInventoryPolicyPtrOutput) ToBlobInventoryPolicyPtrOutputWithContext(ctx context.Context) BlobInventoryPolicyPtrOutput {
	return o
}

type BlobInventoryPolicyArrayOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BlobInventoryPolicy)(nil))
}

func (o BlobInventoryPolicyArrayOutput) ToBlobInventoryPolicyArrayOutput() BlobInventoryPolicyArrayOutput {
	return o
}

func (o BlobInventoryPolicyArrayOutput) ToBlobInventoryPolicyArrayOutputWithContext(ctx context.Context) BlobInventoryPolicyArrayOutput {
	return o
}

func (o BlobInventoryPolicyArrayOutput) Index(i pulumi.IntInput) BlobInventoryPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BlobInventoryPolicy {
		return vs[0].([]BlobInventoryPolicy)[vs[1].(int)]
	}).(BlobInventoryPolicyOutput)
}

type BlobInventoryPolicyMapOutput struct{ *pulumi.OutputState }

func (BlobInventoryPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]BlobInventoryPolicy)(nil))
}

func (o BlobInventoryPolicyMapOutput) ToBlobInventoryPolicyMapOutput() BlobInventoryPolicyMapOutput {
	return o
}

func (o BlobInventoryPolicyMapOutput) ToBlobInventoryPolicyMapOutputWithContext(ctx context.Context) BlobInventoryPolicyMapOutput {
	return o
}

func (o BlobInventoryPolicyMapOutput) MapIndex(k pulumi.StringInput) BlobInventoryPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) BlobInventoryPolicy {
		return vs[0].(map[string]BlobInventoryPolicy)[vs[1].(string)]
	}).(BlobInventoryPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(BlobInventoryPolicyOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyPtrOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyArrayOutput{})
	pulumi.RegisterOutputType(BlobInventoryPolicyMapOutput{})
}
