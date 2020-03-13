// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package search

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ServiceQueryKey struct {
	// The value of this Query Key.
	Key *string `pulumi:"key"`
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name *string `pulumi:"name"`
}

type ServiceQueryKeyInput interface {
	pulumi.Input

	ToServiceQueryKeyOutput() ServiceQueryKeyOutput
	ToServiceQueryKeyOutputWithContext(context.Context) ServiceQueryKeyOutput
}

type ServiceQueryKeyArgs struct {
	// The value of this Query Key.
	Key pulumi.StringPtrInput `pulumi:"key"`
	// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ServiceQueryKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQueryKey)(nil)).Elem()
}

func (i ServiceQueryKeyArgs) ToServiceQueryKeyOutput() ServiceQueryKeyOutput {
	return i.ToServiceQueryKeyOutputWithContext(context.Background())
}

func (i ServiceQueryKeyArgs) ToServiceQueryKeyOutputWithContext(ctx context.Context) ServiceQueryKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQueryKeyOutput)
}

type ServiceQueryKeyArrayInput interface {
	pulumi.Input

	ToServiceQueryKeyArrayOutput() ServiceQueryKeyArrayOutput
	ToServiceQueryKeyArrayOutputWithContext(context.Context) ServiceQueryKeyArrayOutput
}

type ServiceQueryKeyArray []ServiceQueryKeyInput

func (ServiceQueryKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQueryKey)(nil)).Elem()
}

func (i ServiceQueryKeyArray) ToServiceQueryKeyArrayOutput() ServiceQueryKeyArrayOutput {
	return i.ToServiceQueryKeyArrayOutputWithContext(context.Background())
}

func (i ServiceQueryKeyArray) ToServiceQueryKeyArrayOutputWithContext(ctx context.Context) ServiceQueryKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceQueryKeyArrayOutput)
}

type ServiceQueryKeyOutput struct { *pulumi.OutputState }

func (ServiceQueryKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceQueryKey)(nil)).Elem()
}

func (o ServiceQueryKeyOutput) ToServiceQueryKeyOutput() ServiceQueryKeyOutput {
	return o
}

func (o ServiceQueryKeyOutput) ToServiceQueryKeyOutputWithContext(ctx context.Context) ServiceQueryKeyOutput {
	return o
}

// The value of this Query Key.
func (o ServiceQueryKeyOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceQueryKey) *string { return v.Key }).(pulumi.StringPtrOutput)
}

// The Name which should be used for this Search Service. Changing this forces a new Search Service to be created.
func (o ServiceQueryKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func (v ServiceQueryKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ServiceQueryKeyArrayOutput struct { *pulumi.OutputState}

func (ServiceQueryKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceQueryKey)(nil)).Elem()
}

func (o ServiceQueryKeyArrayOutput) ToServiceQueryKeyArrayOutput() ServiceQueryKeyArrayOutput {
	return o
}

func (o ServiceQueryKeyArrayOutput) ToServiceQueryKeyArrayOutputWithContext(ctx context.Context) ServiceQueryKeyArrayOutput {
	return o
}

func (o ServiceQueryKeyArrayOutput) Index(i pulumi.IntInput) ServiceQueryKeyOutput {
	return pulumi.All(o, i).ApplyT(func (vs []interface{}) ServiceQueryKey {
		return vs[0].([]ServiceQueryKey)[vs[1].(int)]
	}).(ServiceQueryKeyOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceQueryKeyOutput{})
	pulumi.RegisterOutputType(ServiceQueryKeyArrayOutput{})
}