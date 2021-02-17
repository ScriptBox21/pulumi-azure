// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Deprecated: ZipBlob resource is deprecated in the 2.0 version of the provider. Use Blob resource instead.
type ZipBlob struct {
	pulumi.CustomResourceState

	AccessTier           pulumi.StringOutput    `pulumi:"accessTier"`
	Content              pulumi.ArchiveOutput   `pulumi:"content"`
	ContentMd5           pulumi.StringPtrOutput `pulumi:"contentMd5"`
	ContentType          pulumi.StringPtrOutput `pulumi:"contentType"`
	Metadata             pulumi.StringMapOutput `pulumi:"metadata"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	Parallelism          pulumi.IntPtrOutput    `pulumi:"parallelism"`
	Size                 pulumi.IntPtrOutput    `pulumi:"size"`
	SourceContent        pulumi.StringPtrOutput `pulumi:"sourceContent"`
	SourceUri            pulumi.StringPtrOutput `pulumi:"sourceUri"`
	StorageAccountName   pulumi.StringOutput    `pulumi:"storageAccountName"`
	StorageContainerName pulumi.StringOutput    `pulumi:"storageContainerName"`
	Type                 pulumi.StringOutput    `pulumi:"type"`
	Url                  pulumi.StringOutput    `pulumi:"url"`
}

// NewZipBlob registers a new resource with the given unique name, arguments, and options.
func NewZipBlob(ctx *pulumi.Context,
	name string, args *ZipBlobArgs, opts ...pulumi.ResourceOption) (*ZipBlob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StorageAccountName == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountName'")
	}
	if args.StorageContainerName == nil {
		return nil, errors.New("invalid value for required argument 'StorageContainerName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource ZipBlob
	err := ctx.RegisterResource("azure:storage/zipBlob:ZipBlob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetZipBlob gets an existing ZipBlob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZipBlob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ZipBlobState, opts ...pulumi.ResourceOption) (*ZipBlob, error) {
	var resource ZipBlob
	err := ctx.ReadResource("azure:storage/zipBlob:ZipBlob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ZipBlob resources.
type zipBlobState struct {
	AccessTier           *string           `pulumi:"accessTier"`
	Content              pulumi.Archive    `pulumi:"content"`
	ContentMd5           *string           `pulumi:"contentMd5"`
	ContentType          *string           `pulumi:"contentType"`
	Metadata             map[string]string `pulumi:"metadata"`
	Name                 *string           `pulumi:"name"`
	Parallelism          *int              `pulumi:"parallelism"`
	Size                 *int              `pulumi:"size"`
	SourceContent        *string           `pulumi:"sourceContent"`
	SourceUri            *string           `pulumi:"sourceUri"`
	StorageAccountName   *string           `pulumi:"storageAccountName"`
	StorageContainerName *string           `pulumi:"storageContainerName"`
	Type                 *string           `pulumi:"type"`
	Url                  *string           `pulumi:"url"`
}

type ZipBlobState struct {
	AccessTier           pulumi.StringPtrInput
	Content              pulumi.ArchiveInput
	ContentMd5           pulumi.StringPtrInput
	ContentType          pulumi.StringPtrInput
	Metadata             pulumi.StringMapInput
	Name                 pulumi.StringPtrInput
	Parallelism          pulumi.IntPtrInput
	Size                 pulumi.IntPtrInput
	SourceContent        pulumi.StringPtrInput
	SourceUri            pulumi.StringPtrInput
	StorageAccountName   pulumi.StringPtrInput
	StorageContainerName pulumi.StringPtrInput
	Type                 pulumi.StringPtrInput
	Url                  pulumi.StringPtrInput
}

func (ZipBlobState) ElementType() reflect.Type {
	return reflect.TypeOf((*zipBlobState)(nil)).Elem()
}

type zipBlobArgs struct {
	AccessTier           *string           `pulumi:"accessTier"`
	Content              pulumi.Archive    `pulumi:"content"`
	ContentMd5           *string           `pulumi:"contentMd5"`
	ContentType          *string           `pulumi:"contentType"`
	Metadata             map[string]string `pulumi:"metadata"`
	Name                 *string           `pulumi:"name"`
	Parallelism          *int              `pulumi:"parallelism"`
	Size                 *int              `pulumi:"size"`
	SourceContent        *string           `pulumi:"sourceContent"`
	SourceUri            *string           `pulumi:"sourceUri"`
	StorageAccountName   string            `pulumi:"storageAccountName"`
	StorageContainerName string            `pulumi:"storageContainerName"`
	Type                 string            `pulumi:"type"`
}

// The set of arguments for constructing a ZipBlob resource.
type ZipBlobArgs struct {
	AccessTier           pulumi.StringPtrInput
	Content              pulumi.ArchiveInput
	ContentMd5           pulumi.StringPtrInput
	ContentType          pulumi.StringPtrInput
	Metadata             pulumi.StringMapInput
	Name                 pulumi.StringPtrInput
	Parallelism          pulumi.IntPtrInput
	Size                 pulumi.IntPtrInput
	SourceContent        pulumi.StringPtrInput
	SourceUri            pulumi.StringPtrInput
	StorageAccountName   pulumi.StringInput
	StorageContainerName pulumi.StringInput
	Type                 pulumi.StringInput
}

func (ZipBlobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*zipBlobArgs)(nil)).Elem()
}

type ZipBlobInput interface {
	pulumi.Input

	ToZipBlobOutput() ZipBlobOutput
	ToZipBlobOutputWithContext(ctx context.Context) ZipBlobOutput
}

func (*ZipBlob) ElementType() reflect.Type {
	return reflect.TypeOf((*ZipBlob)(nil))
}

func (i *ZipBlob) ToZipBlobOutput() ZipBlobOutput {
	return i.ToZipBlobOutputWithContext(context.Background())
}

func (i *ZipBlob) ToZipBlobOutputWithContext(ctx context.Context) ZipBlobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZipBlobOutput)
}

func (i *ZipBlob) ToZipBlobPtrOutput() ZipBlobPtrOutput {
	return i.ToZipBlobPtrOutputWithContext(context.Background())
}

func (i *ZipBlob) ToZipBlobPtrOutputWithContext(ctx context.Context) ZipBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZipBlobPtrOutput)
}

type ZipBlobPtrInput interface {
	pulumi.Input

	ToZipBlobPtrOutput() ZipBlobPtrOutput
	ToZipBlobPtrOutputWithContext(ctx context.Context) ZipBlobPtrOutput
}

type zipBlobPtrType ZipBlobArgs

func (*zipBlobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ZipBlob)(nil))
}

func (i *zipBlobPtrType) ToZipBlobPtrOutput() ZipBlobPtrOutput {
	return i.ToZipBlobPtrOutputWithContext(context.Background())
}

func (i *zipBlobPtrType) ToZipBlobPtrOutputWithContext(ctx context.Context) ZipBlobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZipBlobPtrOutput)
}

// ZipBlobArrayInput is an input type that accepts ZipBlobArray and ZipBlobArrayOutput values.
// You can construct a concrete instance of `ZipBlobArrayInput` via:
//
//          ZipBlobArray{ ZipBlobArgs{...} }
type ZipBlobArrayInput interface {
	pulumi.Input

	ToZipBlobArrayOutput() ZipBlobArrayOutput
	ToZipBlobArrayOutputWithContext(context.Context) ZipBlobArrayOutput
}

type ZipBlobArray []ZipBlobInput

func (ZipBlobArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*ZipBlob)(nil))
}

func (i ZipBlobArray) ToZipBlobArrayOutput() ZipBlobArrayOutput {
	return i.ToZipBlobArrayOutputWithContext(context.Background())
}

func (i ZipBlobArray) ToZipBlobArrayOutputWithContext(ctx context.Context) ZipBlobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZipBlobArrayOutput)
}

// ZipBlobMapInput is an input type that accepts ZipBlobMap and ZipBlobMapOutput values.
// You can construct a concrete instance of `ZipBlobMapInput` via:
//
//          ZipBlobMap{ "key": ZipBlobArgs{...} }
type ZipBlobMapInput interface {
	pulumi.Input

	ToZipBlobMapOutput() ZipBlobMapOutput
	ToZipBlobMapOutputWithContext(context.Context) ZipBlobMapOutput
}

type ZipBlobMap map[string]ZipBlobInput

func (ZipBlobMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*ZipBlob)(nil))
}

func (i ZipBlobMap) ToZipBlobMapOutput() ZipBlobMapOutput {
	return i.ToZipBlobMapOutputWithContext(context.Background())
}

func (i ZipBlobMap) ToZipBlobMapOutputWithContext(ctx context.Context) ZipBlobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ZipBlobMapOutput)
}

type ZipBlobOutput struct {
	*pulumi.OutputState
}

func (ZipBlobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZipBlob)(nil))
}

func (o ZipBlobOutput) ToZipBlobOutput() ZipBlobOutput {
	return o
}

func (o ZipBlobOutput) ToZipBlobOutputWithContext(ctx context.Context) ZipBlobOutput {
	return o
}

func (o ZipBlobOutput) ToZipBlobPtrOutput() ZipBlobPtrOutput {
	return o.ToZipBlobPtrOutputWithContext(context.Background())
}

func (o ZipBlobOutput) ToZipBlobPtrOutputWithContext(ctx context.Context) ZipBlobPtrOutput {
	return o.ApplyT(func(v ZipBlob) *ZipBlob {
		return &v
	}).(ZipBlobPtrOutput)
}

type ZipBlobPtrOutput struct {
	*pulumi.OutputState
}

func (ZipBlobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ZipBlob)(nil))
}

func (o ZipBlobPtrOutput) ToZipBlobPtrOutput() ZipBlobPtrOutput {
	return o
}

func (o ZipBlobPtrOutput) ToZipBlobPtrOutputWithContext(ctx context.Context) ZipBlobPtrOutput {
	return o
}

type ZipBlobArrayOutput struct{ *pulumi.OutputState }

func (ZipBlobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ZipBlob)(nil))
}

func (o ZipBlobArrayOutput) ToZipBlobArrayOutput() ZipBlobArrayOutput {
	return o
}

func (o ZipBlobArrayOutput) ToZipBlobArrayOutputWithContext(ctx context.Context) ZipBlobArrayOutput {
	return o
}

func (o ZipBlobArrayOutput) Index(i pulumi.IntInput) ZipBlobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ZipBlob {
		return vs[0].([]ZipBlob)[vs[1].(int)]
	}).(ZipBlobOutput)
}

type ZipBlobMapOutput struct{ *pulumi.OutputState }

func (ZipBlobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ZipBlob)(nil))
}

func (o ZipBlobMapOutput) ToZipBlobMapOutput() ZipBlobMapOutput {
	return o
}

func (o ZipBlobMapOutput) ToZipBlobMapOutputWithContext(ctx context.Context) ZipBlobMapOutput {
	return o
}

func (o ZipBlobMapOutput) MapIndex(k pulumi.StringInput) ZipBlobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ZipBlob {
		return vs[0].(map[string]ZipBlob)[vs[1].(string)]
	}).(ZipBlobOutput)
}

func init() {
	pulumi.RegisterOutputType(ZipBlobOutput{})
	pulumi.RegisterOutputType(ZipBlobPtrOutput{})
	pulumi.RegisterOutputType(ZipBlobArrayOutput{})
	pulumi.RegisterOutputType(ZipBlobMapOutput{})
}
