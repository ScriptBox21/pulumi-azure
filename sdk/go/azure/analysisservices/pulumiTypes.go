// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package analysisservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

type ServerIpv4FirewallRule struct {
	// Specifies the name of the firewall rule.
	Name string `pulumi:"name"`
	// End of the firewall rule range as IPv4 address.
	RangeEnd string `pulumi:"rangeEnd"`
	// Start of the firewall rule range as IPv4 address.
	RangeStart string `pulumi:"rangeStart"`
}

type ServerIpv4FirewallRuleInput interface {
	pulumi.Input

	ToServerIpv4FirewallRuleOutput() ServerIpv4FirewallRuleOutput
	ToServerIpv4FirewallRuleOutputWithContext(context.Context) ServerIpv4FirewallRuleOutput
}

type ServerIpv4FirewallRuleArgs struct {
	// Specifies the name of the firewall rule.
	Name pulumi.StringInput `pulumi:"name"`
	// End of the firewall rule range as IPv4 address.
	RangeEnd pulumi.StringInput `pulumi:"rangeEnd"`
	// Start of the firewall rule range as IPv4 address.
	RangeStart pulumi.StringInput `pulumi:"rangeStart"`
}

func (ServerIpv4FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerIpv4FirewallRule)(nil)).Elem()
}

func (i ServerIpv4FirewallRuleArgs) ToServerIpv4FirewallRuleOutput() ServerIpv4FirewallRuleOutput {
	return i.ToServerIpv4FirewallRuleOutputWithContext(context.Background())
}

func (i ServerIpv4FirewallRuleArgs) ToServerIpv4FirewallRuleOutputWithContext(ctx context.Context) ServerIpv4FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerIpv4FirewallRuleOutput)
}

type ServerIpv4FirewallRuleArrayInput interface {
	pulumi.Input

	ToServerIpv4FirewallRuleArrayOutput() ServerIpv4FirewallRuleArrayOutput
	ToServerIpv4FirewallRuleArrayOutputWithContext(context.Context) ServerIpv4FirewallRuleArrayOutput
}

type ServerIpv4FirewallRuleArray []ServerIpv4FirewallRuleInput

func (ServerIpv4FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerIpv4FirewallRule)(nil)).Elem()
}

func (i ServerIpv4FirewallRuleArray) ToServerIpv4FirewallRuleArrayOutput() ServerIpv4FirewallRuleArrayOutput {
	return i.ToServerIpv4FirewallRuleArrayOutputWithContext(context.Background())
}

func (i ServerIpv4FirewallRuleArray) ToServerIpv4FirewallRuleArrayOutputWithContext(ctx context.Context) ServerIpv4FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerIpv4FirewallRuleArrayOutput)
}

type ServerIpv4FirewallRuleOutput struct{ *pulumi.OutputState }

func (ServerIpv4FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerIpv4FirewallRule)(nil)).Elem()
}

func (o ServerIpv4FirewallRuleOutput) ToServerIpv4FirewallRuleOutput() ServerIpv4FirewallRuleOutput {
	return o
}

func (o ServerIpv4FirewallRuleOutput) ToServerIpv4FirewallRuleOutputWithContext(ctx context.Context) ServerIpv4FirewallRuleOutput {
	return o
}

// Specifies the name of the firewall rule.
func (o ServerIpv4FirewallRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServerIpv4FirewallRule) string { return v.Name }).(pulumi.StringOutput)
}

// End of the firewall rule range as IPv4 address.
func (o ServerIpv4FirewallRuleOutput) RangeEnd() pulumi.StringOutput {
	return o.ApplyT(func(v ServerIpv4FirewallRule) string { return v.RangeEnd }).(pulumi.StringOutput)
}

// Start of the firewall rule range as IPv4 address.
func (o ServerIpv4FirewallRuleOutput) RangeStart() pulumi.StringOutput {
	return o.ApplyT(func(v ServerIpv4FirewallRule) string { return v.RangeStart }).(pulumi.StringOutput)
}

type ServerIpv4FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (ServerIpv4FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServerIpv4FirewallRule)(nil)).Elem()
}

func (o ServerIpv4FirewallRuleArrayOutput) ToServerIpv4FirewallRuleArrayOutput() ServerIpv4FirewallRuleArrayOutput {
	return o
}

func (o ServerIpv4FirewallRuleArrayOutput) ToServerIpv4FirewallRuleArrayOutputWithContext(ctx context.Context) ServerIpv4FirewallRuleArrayOutput {
	return o
}

func (o ServerIpv4FirewallRuleArrayOutput) Index(i pulumi.IntInput) ServerIpv4FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServerIpv4FirewallRule {
		return vs[0].([]ServerIpv4FirewallRule)[vs[1].(int)]
	}).(ServerIpv4FirewallRuleOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerIpv4FirewallRuleOutput{})
	pulumi.RegisterOutputType(ServerIpv4FirewallRuleArrayOutput{})
}
