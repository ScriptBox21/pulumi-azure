// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicebus

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SubscriptionRuleCorrelationFilter struct {
	// Content type of the message.
	ContentType *string `pulumi:"contentType"`
	// Identifier of the correlation.
	CorrelationId *string `pulumi:"correlationId"`
	// Application specific label.
	Label *string `pulumi:"label"`
	// Identifier of the message.
	MessageId *string `pulumi:"messageId"`
	// Address of the queue to reply to.
	ReplyTo *string `pulumi:"replyTo"`
	// Session identifier to reply to.
	ReplyToSessionId *string `pulumi:"replyToSessionId"`
	// Session identifier.
	SessionId *string `pulumi:"sessionId"`
	// Address to send to.
	To *string `pulumi:"to"`
}

// SubscriptionRuleCorrelationFilterInput is an input type that accepts SubscriptionRuleCorrelationFilterArgs and SubscriptionRuleCorrelationFilterOutput values.
// You can construct a concrete instance of `SubscriptionRuleCorrelationFilterInput` via:
//
// 		 SubscriptionRuleCorrelationFilterArgs{...}
//
type SubscriptionRuleCorrelationFilterInput interface {
	pulumi.Input

	ToSubscriptionRuleCorrelationFilterOutput() SubscriptionRuleCorrelationFilterOutput
	ToSubscriptionRuleCorrelationFilterOutputWithContext(context.Context) SubscriptionRuleCorrelationFilterOutput
}

type SubscriptionRuleCorrelationFilterArgs struct {
	// Content type of the message.
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	// Identifier of the correlation.
	CorrelationId pulumi.StringPtrInput `pulumi:"correlationId"`
	// Application specific label.
	Label pulumi.StringPtrInput `pulumi:"label"`
	// Identifier of the message.
	MessageId pulumi.StringPtrInput `pulumi:"messageId"`
	// Address of the queue to reply to.
	ReplyTo pulumi.StringPtrInput `pulumi:"replyTo"`
	// Session identifier to reply to.
	ReplyToSessionId pulumi.StringPtrInput `pulumi:"replyToSessionId"`
	// Session identifier.
	SessionId pulumi.StringPtrInput `pulumi:"sessionId"`
	// Address to send to.
	To pulumi.StringPtrInput `pulumi:"to"`
}

func (SubscriptionRuleCorrelationFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionRuleCorrelationFilter)(nil)).Elem()
}

func (i SubscriptionRuleCorrelationFilterArgs) ToSubscriptionRuleCorrelationFilterOutput() SubscriptionRuleCorrelationFilterOutput {
	return i.ToSubscriptionRuleCorrelationFilterOutputWithContext(context.Background())
}

func (i SubscriptionRuleCorrelationFilterArgs) ToSubscriptionRuleCorrelationFilterOutputWithContext(ctx context.Context) SubscriptionRuleCorrelationFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRuleCorrelationFilterOutput)
}

func (i SubscriptionRuleCorrelationFilterArgs) ToSubscriptionRuleCorrelationFilterPtrOutput() SubscriptionRuleCorrelationFilterPtrOutput {
	return i.ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i SubscriptionRuleCorrelationFilterArgs) ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(ctx context.Context) SubscriptionRuleCorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRuleCorrelationFilterOutput).ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(ctx)
}

// SubscriptionRuleCorrelationFilterPtrInput is an input type that accepts SubscriptionRuleCorrelationFilterArgs, SubscriptionRuleCorrelationFilterPtr and SubscriptionRuleCorrelationFilterPtrOutput values.
// You can construct a concrete instance of `SubscriptionRuleCorrelationFilterPtrInput` via:
//
// 		 SubscriptionRuleCorrelationFilterArgs{...}
//
//  or:
//
// 		 nil
//
type SubscriptionRuleCorrelationFilterPtrInput interface {
	pulumi.Input

	ToSubscriptionRuleCorrelationFilterPtrOutput() SubscriptionRuleCorrelationFilterPtrOutput
	ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(context.Context) SubscriptionRuleCorrelationFilterPtrOutput
}

type subscriptionRuleCorrelationFilterPtrType SubscriptionRuleCorrelationFilterArgs

func SubscriptionRuleCorrelationFilterPtr(v *SubscriptionRuleCorrelationFilterArgs) SubscriptionRuleCorrelationFilterPtrInput {
	return (*subscriptionRuleCorrelationFilterPtrType)(v)
}

func (*subscriptionRuleCorrelationFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionRuleCorrelationFilter)(nil)).Elem()
}

func (i *subscriptionRuleCorrelationFilterPtrType) ToSubscriptionRuleCorrelationFilterPtrOutput() SubscriptionRuleCorrelationFilterPtrOutput {
	return i.ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(context.Background())
}

func (i *subscriptionRuleCorrelationFilterPtrType) ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(ctx context.Context) SubscriptionRuleCorrelationFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionRuleCorrelationFilterPtrOutput)
}

type SubscriptionRuleCorrelationFilterOutput struct{ *pulumi.OutputState }

func (SubscriptionRuleCorrelationFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionRuleCorrelationFilter)(nil)).Elem()
}

func (o SubscriptionRuleCorrelationFilterOutput) ToSubscriptionRuleCorrelationFilterOutput() SubscriptionRuleCorrelationFilterOutput {
	return o
}

func (o SubscriptionRuleCorrelationFilterOutput) ToSubscriptionRuleCorrelationFilterOutputWithContext(ctx context.Context) SubscriptionRuleCorrelationFilterOutput {
	return o
}

func (o SubscriptionRuleCorrelationFilterOutput) ToSubscriptionRuleCorrelationFilterPtrOutput() SubscriptionRuleCorrelationFilterPtrOutput {
	return o.ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(context.Background())
}

func (o SubscriptionRuleCorrelationFilterOutput) ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(ctx context.Context) SubscriptionRuleCorrelationFilterPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *SubscriptionRuleCorrelationFilter {
		return &v
	}).(SubscriptionRuleCorrelationFilterPtrOutput)
}

// Content type of the message.
func (o SubscriptionRuleCorrelationFilterOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

// Identifier of the correlation.
func (o SubscriptionRuleCorrelationFilterOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

// Application specific label.
func (o SubscriptionRuleCorrelationFilterOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.Label }).(pulumi.StringPtrOutput)
}

// Identifier of the message.
func (o SubscriptionRuleCorrelationFilterOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

// Address of the queue to reply to.
func (o SubscriptionRuleCorrelationFilterOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

// Session identifier to reply to.
func (o SubscriptionRuleCorrelationFilterOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

// Session identifier.
func (o SubscriptionRuleCorrelationFilterOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

// Address to send to.
func (o SubscriptionRuleCorrelationFilterOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.To }).(pulumi.StringPtrOutput)
}

type SubscriptionRuleCorrelationFilterPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionRuleCorrelationFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionRuleCorrelationFilter)(nil)).Elem()
}

func (o SubscriptionRuleCorrelationFilterPtrOutput) ToSubscriptionRuleCorrelationFilterPtrOutput() SubscriptionRuleCorrelationFilterPtrOutput {
	return o
}

func (o SubscriptionRuleCorrelationFilterPtrOutput) ToSubscriptionRuleCorrelationFilterPtrOutputWithContext(ctx context.Context) SubscriptionRuleCorrelationFilterPtrOutput {
	return o
}

func (o SubscriptionRuleCorrelationFilterPtrOutput) Elem() SubscriptionRuleCorrelationFilterOutput {
	return o.ApplyT(func(v *SubscriptionRuleCorrelationFilter) SubscriptionRuleCorrelationFilter { return *v }).(SubscriptionRuleCorrelationFilterOutput)
}

// Content type of the message.
func (o SubscriptionRuleCorrelationFilterPtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

// Identifier of the correlation.
func (o SubscriptionRuleCorrelationFilterPtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

// Application specific label.
func (o SubscriptionRuleCorrelationFilterPtrOutput) Label() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.Label }).(pulumi.StringPtrOutput)
}

// Identifier of the message.
func (o SubscriptionRuleCorrelationFilterPtrOutput) MessageId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.MessageId }).(pulumi.StringPtrOutput)
}

// Address of the queue to reply to.
func (o SubscriptionRuleCorrelationFilterPtrOutput) ReplyTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.ReplyTo }).(pulumi.StringPtrOutput)
}

// Session identifier to reply to.
func (o SubscriptionRuleCorrelationFilterPtrOutput) ReplyToSessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.ReplyToSessionId }).(pulumi.StringPtrOutput)
}

// Session identifier.
func (o SubscriptionRuleCorrelationFilterPtrOutput) SessionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.SessionId }).(pulumi.StringPtrOutput)
}

// Address to send to.
func (o SubscriptionRuleCorrelationFilterPtrOutput) To() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionRuleCorrelationFilter) *string { return v.To }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionRuleCorrelationFilterOutput{})
	pulumi.RegisterOutputType(SubscriptionRuleCorrelationFilterPtrOutput{})
}
