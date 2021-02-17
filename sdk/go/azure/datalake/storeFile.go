// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datalake

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Azure Data Lake Store File.
//
// > **Note:** If you want to change the data in the remote file without changing the `localFilePath`, then
// taint the resource so the `datalake.StoreFile` gets recreated with the new data.
//
// ## Import
//
// Data Lake Store File's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datalake/storeFile:StoreFile txt
// ```
type StoreFile struct {
	pulumi.CustomResourceState

	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath pulumi.StringOutput `pulumi:"localFilePath"`
	// The path created for the file on the Data Lake Store.
	RemoteFilePath pulumi.StringOutput `pulumi:"remoteFilePath"`
}

// NewStoreFile registers a new resource with the given unique name, arguments, and options.
func NewStoreFile(ctx *pulumi.Context,
	name string, args *StoreFileArgs, opts ...pulumi.ResourceOption) (*StoreFile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.LocalFilePath == nil {
		return nil, errors.New("invalid value for required argument 'LocalFilePath'")
	}
	if args.RemoteFilePath == nil {
		return nil, errors.New("invalid value for required argument 'RemoteFilePath'")
	}
	var resource StoreFile
	err := ctx.RegisterResource("azure:datalake/storeFile:StoreFile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStoreFile gets an existing StoreFile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStoreFile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StoreFileState, opts ...pulumi.ResourceOption) (*StoreFile, error) {
	var resource StoreFile
	err := ctx.ReadResource("azure:datalake/storeFile:StoreFile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StoreFile resources.
type storeFileState struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName *string `pulumi:"accountName"`
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath *string `pulumi:"localFilePath"`
	// The path created for the file on the Data Lake Store.
	RemoteFilePath *string `pulumi:"remoteFilePath"`
}

type StoreFileState struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName pulumi.StringPtrInput
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath pulumi.StringPtrInput
	// The path created for the file on the Data Lake Store.
	RemoteFilePath pulumi.StringPtrInput
}

func (StoreFileState) ElementType() reflect.Type {
	return reflect.TypeOf((*storeFileState)(nil)).Elem()
}

type storeFileArgs struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName string `pulumi:"accountName"`
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath string `pulumi:"localFilePath"`
	// The path created for the file on the Data Lake Store.
	RemoteFilePath string `pulumi:"remoteFilePath"`
}

// The set of arguments for constructing a StoreFile resource.
type StoreFileArgs struct {
	// Specifies the name of the Data Lake Store for which the File should created.
	AccountName pulumi.StringInput
	// The path to the local file to be added to the Data Lake Store.
	LocalFilePath pulumi.StringInput
	// The path created for the file on the Data Lake Store.
	RemoteFilePath pulumi.StringInput
}

func (StoreFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storeFileArgs)(nil)).Elem()
}

type StoreFileInput interface {
	pulumi.Input

	ToStoreFileOutput() StoreFileOutput
	ToStoreFileOutputWithContext(ctx context.Context) StoreFileOutput
}

func (*StoreFile) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreFile)(nil))
}

func (i *StoreFile) ToStoreFileOutput() StoreFileOutput {
	return i.ToStoreFileOutputWithContext(context.Background())
}

func (i *StoreFile) ToStoreFileOutputWithContext(ctx context.Context) StoreFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFileOutput)
}

func (i *StoreFile) ToStoreFilePtrOutput() StoreFilePtrOutput {
	return i.ToStoreFilePtrOutputWithContext(context.Background())
}

func (i *StoreFile) ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFilePtrOutput)
}

type StoreFilePtrInput interface {
	pulumi.Input

	ToStoreFilePtrOutput() StoreFilePtrOutput
	ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput
}

type storeFilePtrType StoreFileArgs

func (*storeFilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StoreFile)(nil))
}

func (i *storeFilePtrType) ToStoreFilePtrOutput() StoreFilePtrOutput {
	return i.ToStoreFilePtrOutputWithContext(context.Background())
}

func (i *storeFilePtrType) ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFilePtrOutput)
}

// StoreFileArrayInput is an input type that accepts StoreFileArray and StoreFileArrayOutput values.
// You can construct a concrete instance of `StoreFileArrayInput` via:
//
//          StoreFileArray{ StoreFileArgs{...} }
type StoreFileArrayInput interface {
	pulumi.Input

	ToStoreFileArrayOutput() StoreFileArrayOutput
	ToStoreFileArrayOutputWithContext(context.Context) StoreFileArrayOutput
}

type StoreFileArray []StoreFileInput

func (StoreFileArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*StoreFile)(nil))
}

func (i StoreFileArray) ToStoreFileArrayOutput() StoreFileArrayOutput {
	return i.ToStoreFileArrayOutputWithContext(context.Background())
}

func (i StoreFileArray) ToStoreFileArrayOutputWithContext(ctx context.Context) StoreFileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFileArrayOutput)
}

// StoreFileMapInput is an input type that accepts StoreFileMap and StoreFileMapOutput values.
// You can construct a concrete instance of `StoreFileMapInput` via:
//
//          StoreFileMap{ "key": StoreFileArgs{...} }
type StoreFileMapInput interface {
	pulumi.Input

	ToStoreFileMapOutput() StoreFileMapOutput
	ToStoreFileMapOutputWithContext(context.Context) StoreFileMapOutput
}

type StoreFileMap map[string]StoreFileInput

func (StoreFileMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*StoreFile)(nil))
}

func (i StoreFileMap) ToStoreFileMapOutput() StoreFileMapOutput {
	return i.ToStoreFileMapOutputWithContext(context.Background())
}

func (i StoreFileMap) ToStoreFileMapOutputWithContext(ctx context.Context) StoreFileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StoreFileMapOutput)
}

type StoreFileOutput struct {
	*pulumi.OutputState
}

func (StoreFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StoreFile)(nil))
}

func (o StoreFileOutput) ToStoreFileOutput() StoreFileOutput {
	return o
}

func (o StoreFileOutput) ToStoreFileOutputWithContext(ctx context.Context) StoreFileOutput {
	return o
}

func (o StoreFileOutput) ToStoreFilePtrOutput() StoreFilePtrOutput {
	return o.ToStoreFilePtrOutputWithContext(context.Background())
}

func (o StoreFileOutput) ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput {
	return o.ApplyT(func(v StoreFile) *StoreFile {
		return &v
	}).(StoreFilePtrOutput)
}

type StoreFilePtrOutput struct {
	*pulumi.OutputState
}

func (StoreFilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StoreFile)(nil))
}

func (o StoreFilePtrOutput) ToStoreFilePtrOutput() StoreFilePtrOutput {
	return o
}

func (o StoreFilePtrOutput) ToStoreFilePtrOutputWithContext(ctx context.Context) StoreFilePtrOutput {
	return o
}

type StoreFileArrayOutput struct{ *pulumi.OutputState }

func (StoreFileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StoreFile)(nil))
}

func (o StoreFileArrayOutput) ToStoreFileArrayOutput() StoreFileArrayOutput {
	return o
}

func (o StoreFileArrayOutput) ToStoreFileArrayOutputWithContext(ctx context.Context) StoreFileArrayOutput {
	return o
}

func (o StoreFileArrayOutput) Index(i pulumi.IntInput) StoreFileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StoreFile {
		return vs[0].([]StoreFile)[vs[1].(int)]
	}).(StoreFileOutput)
}

type StoreFileMapOutput struct{ *pulumi.OutputState }

func (StoreFileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StoreFile)(nil))
}

func (o StoreFileMapOutput) ToStoreFileMapOutput() StoreFileMapOutput {
	return o
}

func (o StoreFileMapOutput) ToStoreFileMapOutputWithContext(ctx context.Context) StoreFileMapOutput {
	return o
}

func (o StoreFileMapOutput) MapIndex(k pulumi.StringInput) StoreFileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StoreFile {
		return vs[0].(map[string]StoreFile)[vs[1].(string)]
	}).(StoreFileOutput)
}

func init() {
	pulumi.RegisterOutputType(StoreFileOutput{})
	pulumi.RegisterOutputType(StoreFilePtrOutput{})
	pulumi.RegisterOutputType(StoreFileArrayOutput{})
	pulumi.RegisterOutputType(StoreFileMapOutput{})
}
