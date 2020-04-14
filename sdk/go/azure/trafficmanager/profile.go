// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package trafficmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Traffic Manager Profile to which multiple endpoints can be attached.
type Profile struct {
	pulumi.CustomResourceState

	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig ProfileDnsConfigOutput `pulumi:"dnsConfig"`
	// The FQDN of the created Profile.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig ProfileMonitorConfigOutput `pulumi:"monitorConfig"`
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringOutput `pulumi:"profileStatus"`
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod pulumi.StringOutput `pulumi:"trafficRoutingMethod"`
}

// NewProfile registers a new resource with the given unique name, arguments, and options.
func NewProfile(ctx *pulumi.Context,
	name string, args *ProfileArgs, opts ...pulumi.ResourceOption) (*Profile, error) {
	if args == nil || args.DnsConfig == nil {
		return nil, errors.New("missing required argument 'DnsConfig'")
	}
	if args == nil || args.MonitorConfig == nil {
		return nil, errors.New("missing required argument 'MonitorConfig'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TrafficRoutingMethod == nil {
		return nil, errors.New("missing required argument 'TrafficRoutingMethod'")
	}
	if args == nil {
		args = &ProfileArgs{}
	}
	var resource Profile
	err := ctx.RegisterResource("azure:trafficmanager/profile:Profile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProfile gets an existing Profile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProfileState, opts ...pulumi.ResourceOption) (*Profile, error) {
	var resource Profile
	err := ctx.ReadResource("azure:trafficmanager/profile:Profile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Profile resources.
type profileState struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig *ProfileDnsConfig `pulumi:"dnsConfig"`
	// The FQDN of the created Profile.
	Fqdn *string `pulumi:"fqdn"`
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig *ProfileMonitorConfig `pulumi:"monitorConfig"`
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod *string `pulumi:"trafficRoutingMethod"`
}

type ProfileState struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig ProfileDnsConfigPtrInput
	// The FQDN of the created Profile.
	Fqdn pulumi.StringPtrInput
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig ProfileMonitorConfigPtrInput
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod pulumi.StringPtrInput
}

func (ProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*profileState)(nil)).Elem()
}

type profileArgs struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig ProfileDnsConfig `pulumi:"dnsConfig"`
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig ProfileMonitorConfig `pulumi:"monitorConfig"`
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus *string `pulumi:"profileStatus"`
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod string `pulumi:"trafficRoutingMethod"`
}

// The set of arguments for constructing a Profile resource.
type ProfileArgs struct {
	// This block specifies the DNS configuration of the Profile, it supports the fields documented below.
	DnsConfig ProfileDnsConfigInput
	// This block specifies the Endpoint monitoring configuration for the Profile, it supports the fields documented below.
	MonitorConfig ProfileMonitorConfigInput
	// The name of the Traffic Manager profile. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The status of the profile, can be set to either `Enabled` or `Disabled`. Defaults to `Enabled`.
	ProfileStatus pulumi.StringPtrInput
	// The name of the resource group in which to create the Traffic Manager profile.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Specifies the algorithm used to route traffic, possible values are:
	TrafficRoutingMethod pulumi.StringInput
}

func (ProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*profileArgs)(nil)).Elem()
}
