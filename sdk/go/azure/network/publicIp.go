// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manage a Public IP Address.
type PublicIp struct {
	s *pulumi.ResourceState
}

// NewPublicIp registers a new resource with the given unique name, arguments, and options.
func NewPublicIp(ctx *pulumi.Context,
	name string, args *PublicIpArgs, opts ...pulumi.ResourceOpt) (*PublicIp, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allocationMethod"] = nil
		inputs["domainNameLabel"] = nil
		inputs["idleTimeoutInMinutes"] = nil
		inputs["ipVersion"] = nil
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["publicIpAddressAllocation"] = nil
		inputs["publicIpPrefixId"] = nil
		inputs["resourceGroupName"] = nil
		inputs["reverseFqdn"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
		inputs["zones"] = nil
	} else {
		inputs["allocationMethod"] = args.AllocationMethod
		inputs["domainNameLabel"] = args.DomainNameLabel
		inputs["idleTimeoutInMinutes"] = args.IdleTimeoutInMinutes
		inputs["ipVersion"] = args.IpVersion
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["publicIpAddressAllocation"] = args.PublicIpAddressAllocation
		inputs["publicIpPrefixId"] = args.PublicIpPrefixId
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["reverseFqdn"] = args.ReverseFqdn
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
		inputs["zones"] = args.Zones
	}
	inputs["fqdn"] = nil
	inputs["ipAddress"] = nil
	s, err := ctx.RegisterResource("azure:network/publicIp:PublicIp", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PublicIp{s: s}, nil
}

// GetPublicIp gets an existing PublicIp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPublicIp(ctx *pulumi.Context,
	name string, id pulumi.ID, state *PublicIpState, opts ...pulumi.ResourceOpt) (*PublicIp, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allocationMethod"] = state.AllocationMethod
		inputs["domainNameLabel"] = state.DomainNameLabel
		inputs["fqdn"] = state.Fqdn
		inputs["idleTimeoutInMinutes"] = state.IdleTimeoutInMinutes
		inputs["ipAddress"] = state.IpAddress
		inputs["ipVersion"] = state.IpVersion
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["publicIpAddressAllocation"] = state.PublicIpAddressAllocation
		inputs["publicIpPrefixId"] = state.PublicIpPrefixId
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["reverseFqdn"] = state.ReverseFqdn
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
		inputs["zones"] = state.Zones
	}
	s, err := ctx.ReadResource("azure:network/publicIp:PublicIp", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &PublicIp{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *PublicIp) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *PublicIp) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
func (r *PublicIp) AllocationMethod() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["allocationMethod"])
}

// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
func (r *PublicIp) DomainNameLabel() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["domainNameLabel"])
}

// Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone
func (r *PublicIp) Fqdn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["fqdn"])
}

// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
func (r *PublicIp) IdleTimeoutInMinutes() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["idleTimeoutInMinutes"])
}

// The IP address value that was allocated.
func (r *PublicIp) IpAddress() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipAddress"])
}

// The IP Version to use, IPv6 or IPv4.
func (r *PublicIp) IpVersion() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ipVersion"])
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *PublicIp) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Public IP resource . Changing this forces a
// new resource to be created.
func (r *PublicIp) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

func (r *PublicIp) PublicIpAddressAllocation() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicIpAddressAllocation"])
}

// If specified then public IP address allocated will be provided from the public IP prefix resource.
func (r *PublicIp) PublicIpPrefixId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["publicIpPrefixId"])
}

// The name of the resource group in which to
// create the public ip.
func (r *PublicIp) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
func (r *PublicIp) ReverseFqdn() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["reverseFqdn"])
}

// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
func (r *PublicIp) Sku() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sku"])
}

// A mapping of tags to assign to the resource.
func (r *PublicIp) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// A collection containing the availability zone to allocate the Public IP in.
func (r *PublicIp) Zones() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zones"])
}

// Input properties used for looking up and filtering PublicIp resources.
type PublicIpState struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod interface{}
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel interface{}
	// Fully qualified domain name of the A DNS record associated with the public IP. This is the concatenation of the domainNameLabel and the regionalized DNS zone
	Fqdn interface{}
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes interface{}
	// The IP address value that was allocated.
	IpAddress interface{}
	// The IP Version to use, IPv6 or IPv4.
	IpVersion interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name interface{}
	PublicIpAddressAllocation interface{}
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId interface{}
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName interface{}
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn interface{}
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A collection containing the availability zone to allocate the Public IP in.
	Zones interface{}
}

// The set of arguments for constructing a PublicIp resource.
type PublicIpArgs struct {
	// Defines the allocation method for this IP address. Possible values are `Static` or `Dynamic`.
	AllocationMethod interface{}
	// Label for the Domain Name. Will be used to make up the FQDN.  If a domain name label is specified, an A DNS record is created for the public IP in the Microsoft Azure DNS system.
	DomainNameLabel interface{}
	// Specifies the timeout for the TCP idle connection. The value can be set between 4 and 30 minutes.
	IdleTimeoutInMinutes interface{}
	// The IP Version to use, IPv6 or IPv4.
	IpVersion interface{}
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Public IP resource . Changing this forces a
	// new resource to be created.
	Name interface{}
	PublicIpAddressAllocation interface{}
	// If specified then public IP address allocated will be provided from the public IP prefix resource.
	PublicIpPrefixId interface{}
	// The name of the resource group in which to
	// create the public ip.
	ResourceGroupName interface{}
	// A fully qualified domain name that resolves to this public IP address. If the reverseFqdn is specified, then a PTR DNS record is created pointing from the IP address in the in-addr.arpa domain to the reverse FQDN.
	ReverseFqdn interface{}
	// The SKU of the Public IP. Accepted values are `Basic` and `Standard`. Defaults to `Basic`.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// A collection containing the availability zone to allocate the Public IP in.
	Zones interface{}
}
