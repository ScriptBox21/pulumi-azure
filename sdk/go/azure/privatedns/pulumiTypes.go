// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package privatedns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type LinkServiceNatIpConfiguration struct {
	// Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	// Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
	Primary bool `pulumi:"primary"`
	// Specifies a Private Static IP Address for this IP Configuration.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The version of the IP Protocol which should be used. At this time the only supported value is `IPv4`. Defaults to `IPv4`.
	PrivateIpAddressVersion *string `pulumi:"privateIpAddressVersion"`
	// Specifies the ID of the Subnet which should be used for the Private Link Service.
	SubnetId string `pulumi:"subnetId"`
}

// LinkServiceNatIpConfigurationInput is an input type that accepts LinkServiceNatIpConfigurationArgs and LinkServiceNatIpConfigurationOutput values.
// You can construct a concrete instance of `LinkServiceNatIpConfigurationInput` via:
//
// 		 LinkServiceNatIpConfigurationArgs{...}
//
type LinkServiceNatIpConfigurationInput interface {
	pulumi.Input

	ToLinkServiceNatIpConfigurationOutput() LinkServiceNatIpConfigurationOutput
	ToLinkServiceNatIpConfigurationOutputWithContext(context.Context) LinkServiceNatIpConfigurationOutput
}

type LinkServiceNatIpConfigurationArgs struct {
	// Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
	Name pulumi.StringInput `pulumi:"name"`
	// Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
	Primary pulumi.BoolInput `pulumi:"primary"`
	// Specifies a Private Static IP Address for this IP Configuration.
	PrivateIpAddress pulumi.StringPtrInput `pulumi:"privateIpAddress"`
	// The version of the IP Protocol which should be used. At this time the only supported value is `IPv4`. Defaults to `IPv4`.
	PrivateIpAddressVersion pulumi.StringPtrInput `pulumi:"privateIpAddressVersion"`
	// Specifies the ID of the Subnet which should be used for the Private Link Service.
	SubnetId pulumi.StringInput `pulumi:"subnetId"`
}

func (LinkServiceNatIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (i LinkServiceNatIpConfigurationArgs) ToLinkServiceNatIpConfigurationOutput() LinkServiceNatIpConfigurationOutput {
	return i.ToLinkServiceNatIpConfigurationOutputWithContext(context.Background())
}

func (i LinkServiceNatIpConfigurationArgs) ToLinkServiceNatIpConfigurationOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkServiceNatIpConfigurationOutput)
}

// LinkServiceNatIpConfigurationArrayInput is an input type that accepts LinkServiceNatIpConfigurationArray and LinkServiceNatIpConfigurationArrayOutput values.
// You can construct a concrete instance of `LinkServiceNatIpConfigurationArrayInput` via:
//
// 		 LinkServiceNatIpConfigurationArray{ LinkServiceNatIpConfigurationArgs{...} }
//
type LinkServiceNatIpConfigurationArrayInput interface {
	pulumi.Input

	ToLinkServiceNatIpConfigurationArrayOutput() LinkServiceNatIpConfigurationArrayOutput
	ToLinkServiceNatIpConfigurationArrayOutputWithContext(context.Context) LinkServiceNatIpConfigurationArrayOutput
}

type LinkServiceNatIpConfigurationArray []LinkServiceNatIpConfigurationInput

func (LinkServiceNatIpConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (i LinkServiceNatIpConfigurationArray) ToLinkServiceNatIpConfigurationArrayOutput() LinkServiceNatIpConfigurationArrayOutput {
	return i.ToLinkServiceNatIpConfigurationArrayOutputWithContext(context.Background())
}

func (i LinkServiceNatIpConfigurationArray) ToLinkServiceNatIpConfigurationArrayOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkServiceNatIpConfigurationArrayOutput)
}

type LinkServiceNatIpConfigurationOutput struct{ *pulumi.OutputState }

func (LinkServiceNatIpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (o LinkServiceNatIpConfigurationOutput) ToLinkServiceNatIpConfigurationOutput() LinkServiceNatIpConfigurationOutput {
	return o
}

func (o LinkServiceNatIpConfigurationOutput) ToLinkServiceNatIpConfigurationOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationOutput {
	return o
}

// Specifies the name which should be used for the NAT IP Configuration. Changing this forces a new resource to be created.
func (o LinkServiceNatIpConfigurationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LinkServiceNatIpConfiguration) string { return v.Name }).(pulumi.StringOutput)
}

// Is this is the Primary IP Configuration? Changing this forces a new resource to be created.
func (o LinkServiceNatIpConfigurationOutput) Primary() pulumi.BoolOutput {
	return o.ApplyT(func(v LinkServiceNatIpConfiguration) bool { return v.Primary }).(pulumi.BoolOutput)
}

// Specifies a Private Static IP Address for this IP Configuration.
func (o LinkServiceNatIpConfigurationOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkServiceNatIpConfiguration) *string { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

// The version of the IP Protocol which should be used. At this time the only supported value is `IPv4`. Defaults to `IPv4`.
func (o LinkServiceNatIpConfigurationOutput) PrivateIpAddressVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinkServiceNatIpConfiguration) *string { return v.PrivateIpAddressVersion }).(pulumi.StringPtrOutput)
}

// Specifies the ID of the Subnet which should be used for the Private Link Service.
func (o LinkServiceNatIpConfigurationOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LinkServiceNatIpConfiguration) string { return v.SubnetId }).(pulumi.StringOutput)
}

type LinkServiceNatIpConfigurationArrayOutput struct{ *pulumi.OutputState }

func (LinkServiceNatIpConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkServiceNatIpConfiguration)(nil)).Elem()
}

func (o LinkServiceNatIpConfigurationArrayOutput) ToLinkServiceNatIpConfigurationArrayOutput() LinkServiceNatIpConfigurationArrayOutput {
	return o
}

func (o LinkServiceNatIpConfigurationArrayOutput) ToLinkServiceNatIpConfigurationArrayOutputWithContext(ctx context.Context) LinkServiceNatIpConfigurationArrayOutput {
	return o
}

func (o LinkServiceNatIpConfigurationArrayOutput) Index(i pulumi.IntInput) LinkServiceNatIpConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkServiceNatIpConfiguration {
		return vs[0].([]LinkServiceNatIpConfiguration)[vs[1].(int)]
	}).(LinkServiceNatIpConfigurationOutput)
}

type MxRecordRecord struct {
	// The FQDN of the exchange to MX record points to.
	Exchange string `pulumi:"exchange"`
	// The preference of the MX record.
	Preference int `pulumi:"preference"`
}

// MxRecordRecordInput is an input type that accepts MxRecordRecordArgs and MxRecordRecordOutput values.
// You can construct a concrete instance of `MxRecordRecordInput` via:
//
// 		 MxRecordRecordArgs{...}
//
type MxRecordRecordInput interface {
	pulumi.Input

	ToMxRecordRecordOutput() MxRecordRecordOutput
	ToMxRecordRecordOutputWithContext(context.Context) MxRecordRecordOutput
}

type MxRecordRecordArgs struct {
	// The FQDN of the exchange to MX record points to.
	Exchange pulumi.StringInput `pulumi:"exchange"`
	// The preference of the MX record.
	Preference pulumi.IntInput `pulumi:"preference"`
}

func (MxRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordRecord)(nil)).Elem()
}

func (i MxRecordRecordArgs) ToMxRecordRecordOutput() MxRecordRecordOutput {
	return i.ToMxRecordRecordOutputWithContext(context.Background())
}

func (i MxRecordRecordArgs) ToMxRecordRecordOutputWithContext(ctx context.Context) MxRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordRecordOutput)
}

// MxRecordRecordArrayInput is an input type that accepts MxRecordRecordArray and MxRecordRecordArrayOutput values.
// You can construct a concrete instance of `MxRecordRecordArrayInput` via:
//
// 		 MxRecordRecordArray{ MxRecordRecordArgs{...} }
//
type MxRecordRecordArrayInput interface {
	pulumi.Input

	ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput
	ToMxRecordRecordArrayOutputWithContext(context.Context) MxRecordRecordArrayOutput
}

type MxRecordRecordArray []MxRecordRecordInput

func (MxRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordRecord)(nil)).Elem()
}

func (i MxRecordRecordArray) ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput {
	return i.ToMxRecordRecordArrayOutputWithContext(context.Background())
}

func (i MxRecordRecordArray) ToMxRecordRecordArrayOutputWithContext(ctx context.Context) MxRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MxRecordRecordArrayOutput)
}

type MxRecordRecordOutput struct{ *pulumi.OutputState }

func (MxRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MxRecordRecord)(nil)).Elem()
}

func (o MxRecordRecordOutput) ToMxRecordRecordOutput() MxRecordRecordOutput {
	return o
}

func (o MxRecordRecordOutput) ToMxRecordRecordOutputWithContext(ctx context.Context) MxRecordRecordOutput {
	return o
}

// The FQDN of the exchange to MX record points to.
func (o MxRecordRecordOutput) Exchange() pulumi.StringOutput {
	return o.ApplyT(func(v MxRecordRecord) string { return v.Exchange }).(pulumi.StringOutput)
}

// The preference of the MX record.
func (o MxRecordRecordOutput) Preference() pulumi.IntOutput {
	return o.ApplyT(func(v MxRecordRecord) int { return v.Preference }).(pulumi.IntOutput)
}

type MxRecordRecordArrayOutput struct{ *pulumi.OutputState }

func (MxRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MxRecordRecord)(nil)).Elem()
}

func (o MxRecordRecordArrayOutput) ToMxRecordRecordArrayOutput() MxRecordRecordArrayOutput {
	return o
}

func (o MxRecordRecordArrayOutput) ToMxRecordRecordArrayOutputWithContext(ctx context.Context) MxRecordRecordArrayOutput {
	return o
}

func (o MxRecordRecordArrayOutput) Index(i pulumi.IntInput) MxRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MxRecordRecord {
		return vs[0].([]MxRecordRecord)[vs[1].(int)]
	}).(MxRecordRecordOutput)
}

type SRVRecordRecord struct {
	// The Port the service is listening on.
	Port int `pulumi:"port"`
	// The priority of the SRV record.
	Priority int `pulumi:"priority"`
	// The FQDN of the service.
	Target string `pulumi:"target"`
	// The Weight of the SRV record.
	Weight int `pulumi:"weight"`
}

// SRVRecordRecordInput is an input type that accepts SRVRecordRecordArgs and SRVRecordRecordOutput values.
// You can construct a concrete instance of `SRVRecordRecordInput` via:
//
// 		 SRVRecordRecordArgs{...}
//
type SRVRecordRecordInput interface {
	pulumi.Input

	ToSRVRecordRecordOutput() SRVRecordRecordOutput
	ToSRVRecordRecordOutputWithContext(context.Context) SRVRecordRecordOutput
}

type SRVRecordRecordArgs struct {
	// The Port the service is listening on.
	Port pulumi.IntInput `pulumi:"port"`
	// The priority of the SRV record.
	Priority pulumi.IntInput `pulumi:"priority"`
	// The FQDN of the service.
	Target pulumi.StringInput `pulumi:"target"`
	// The Weight of the SRV record.
	Weight pulumi.IntInput `pulumi:"weight"`
}

func (SRVRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SRVRecordRecord)(nil)).Elem()
}

func (i SRVRecordRecordArgs) ToSRVRecordRecordOutput() SRVRecordRecordOutput {
	return i.ToSRVRecordRecordOutputWithContext(context.Background())
}

func (i SRVRecordRecordArgs) ToSRVRecordRecordOutputWithContext(ctx context.Context) SRVRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SRVRecordRecordOutput)
}

// SRVRecordRecordArrayInput is an input type that accepts SRVRecordRecordArray and SRVRecordRecordArrayOutput values.
// You can construct a concrete instance of `SRVRecordRecordArrayInput` via:
//
// 		 SRVRecordRecordArray{ SRVRecordRecordArgs{...} }
//
type SRVRecordRecordArrayInput interface {
	pulumi.Input

	ToSRVRecordRecordArrayOutput() SRVRecordRecordArrayOutput
	ToSRVRecordRecordArrayOutputWithContext(context.Context) SRVRecordRecordArrayOutput
}

type SRVRecordRecordArray []SRVRecordRecordInput

func (SRVRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SRVRecordRecord)(nil)).Elem()
}

func (i SRVRecordRecordArray) ToSRVRecordRecordArrayOutput() SRVRecordRecordArrayOutput {
	return i.ToSRVRecordRecordArrayOutputWithContext(context.Background())
}

func (i SRVRecordRecordArray) ToSRVRecordRecordArrayOutputWithContext(ctx context.Context) SRVRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SRVRecordRecordArrayOutput)
}

type SRVRecordRecordOutput struct{ *pulumi.OutputState }

func (SRVRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SRVRecordRecord)(nil)).Elem()
}

func (o SRVRecordRecordOutput) ToSRVRecordRecordOutput() SRVRecordRecordOutput {
	return o
}

func (o SRVRecordRecordOutput) ToSRVRecordRecordOutputWithContext(ctx context.Context) SRVRecordRecordOutput {
	return o
}

// The Port the service is listening on.
func (o SRVRecordRecordOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v SRVRecordRecord) int { return v.Port }).(pulumi.IntOutput)
}

// The priority of the SRV record.
func (o SRVRecordRecordOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v SRVRecordRecord) int { return v.Priority }).(pulumi.IntOutput)
}

// The FQDN of the service.
func (o SRVRecordRecordOutput) Target() pulumi.StringOutput {
	return o.ApplyT(func(v SRVRecordRecord) string { return v.Target }).(pulumi.StringOutput)
}

// The Weight of the SRV record.
func (o SRVRecordRecordOutput) Weight() pulumi.IntOutput {
	return o.ApplyT(func(v SRVRecordRecord) int { return v.Weight }).(pulumi.IntOutput)
}

type SRVRecordRecordArrayOutput struct{ *pulumi.OutputState }

func (SRVRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SRVRecordRecord)(nil)).Elem()
}

func (o SRVRecordRecordArrayOutput) ToSRVRecordRecordArrayOutput() SRVRecordRecordArrayOutput {
	return o
}

func (o SRVRecordRecordArrayOutput) ToSRVRecordRecordArrayOutputWithContext(ctx context.Context) SRVRecordRecordArrayOutput {
	return o
}

func (o SRVRecordRecordArrayOutput) Index(i pulumi.IntInput) SRVRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SRVRecordRecord {
		return vs[0].([]SRVRecordRecord)[vs[1].(int)]
	}).(SRVRecordRecordOutput)
}

type TxtRecordRecord struct {
	// The value of the TXT record. Max length: 1024 characters
	Value string `pulumi:"value"`
}

// TxtRecordRecordInput is an input type that accepts TxtRecordRecordArgs and TxtRecordRecordOutput values.
// You can construct a concrete instance of `TxtRecordRecordInput` via:
//
// 		 TxtRecordRecordArgs{...}
//
type TxtRecordRecordInput interface {
	pulumi.Input

	ToTxtRecordRecordOutput() TxtRecordRecordOutput
	ToTxtRecordRecordOutputWithContext(context.Context) TxtRecordRecordOutput
}

type TxtRecordRecordArgs struct {
	// The value of the TXT record. Max length: 1024 characters
	Value pulumi.StringInput `pulumi:"value"`
}

func (TxtRecordRecordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordRecord)(nil)).Elem()
}

func (i TxtRecordRecordArgs) ToTxtRecordRecordOutput() TxtRecordRecordOutput {
	return i.ToTxtRecordRecordOutputWithContext(context.Background())
}

func (i TxtRecordRecordArgs) ToTxtRecordRecordOutputWithContext(ctx context.Context) TxtRecordRecordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordRecordOutput)
}

// TxtRecordRecordArrayInput is an input type that accepts TxtRecordRecordArray and TxtRecordRecordArrayOutput values.
// You can construct a concrete instance of `TxtRecordRecordArrayInput` via:
//
// 		 TxtRecordRecordArray{ TxtRecordRecordArgs{...} }
//
type TxtRecordRecordArrayInput interface {
	pulumi.Input

	ToTxtRecordRecordArrayOutput() TxtRecordRecordArrayOutput
	ToTxtRecordRecordArrayOutputWithContext(context.Context) TxtRecordRecordArrayOutput
}

type TxtRecordRecordArray []TxtRecordRecordInput

func (TxtRecordRecordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordRecord)(nil)).Elem()
}

func (i TxtRecordRecordArray) ToTxtRecordRecordArrayOutput() TxtRecordRecordArrayOutput {
	return i.ToTxtRecordRecordArrayOutputWithContext(context.Background())
}

func (i TxtRecordRecordArray) ToTxtRecordRecordArrayOutputWithContext(ctx context.Context) TxtRecordRecordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TxtRecordRecordArrayOutput)
}

type TxtRecordRecordOutput struct{ *pulumi.OutputState }

func (TxtRecordRecordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TxtRecordRecord)(nil)).Elem()
}

func (o TxtRecordRecordOutput) ToTxtRecordRecordOutput() TxtRecordRecordOutput {
	return o
}

func (o TxtRecordRecordOutput) ToTxtRecordRecordOutputWithContext(ctx context.Context) TxtRecordRecordOutput {
	return o
}

// The value of the TXT record. Max length: 1024 characters
func (o TxtRecordRecordOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TxtRecordRecord) string { return v.Value }).(pulumi.StringOutput)
}

type TxtRecordRecordArrayOutput struct{ *pulumi.OutputState }

func (TxtRecordRecordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TxtRecordRecord)(nil)).Elem()
}

func (o TxtRecordRecordArrayOutput) ToTxtRecordRecordArrayOutput() TxtRecordRecordArrayOutput {
	return o
}

func (o TxtRecordRecordArrayOutput) ToTxtRecordRecordArrayOutputWithContext(ctx context.Context) TxtRecordRecordArrayOutput {
	return o
}

func (o TxtRecordRecordArrayOutput) Index(i pulumi.IntInput) TxtRecordRecordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TxtRecordRecord {
		return vs[0].([]TxtRecordRecord)[vs[1].(int)]
	}).(TxtRecordRecordOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkServiceNatIpConfigurationOutput{})
	pulumi.RegisterOutputType(LinkServiceNatIpConfigurationArrayOutput{})
	pulumi.RegisterOutputType(MxRecordRecordOutput{})
	pulumi.RegisterOutputType(MxRecordRecordArrayOutput{})
	pulumi.RegisterOutputType(SRVRecordRecordOutput{})
	pulumi.RegisterOutputType(SRVRecordRecordArrayOutput{})
	pulumi.RegisterOutputType(TxtRecordRecordOutput{})
	pulumi.RegisterOutputType(TxtRecordRecordArrayOutput{})
}
