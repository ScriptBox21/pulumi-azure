// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Enables you to manage DNS zones within Azure DNS. These zones are hosted on Azure's name servers to which you can delegate the zone from the parent domain.
type Zone struct {
	s *pulumi.ResourceState
}

// NewZone registers a new resource with the given unique name, arguments, and options.
func NewZone(ctx *pulumi.Context,
	name string, args *ZoneArgs, opts ...pulumi.ResourceOpt) (*Zone, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["name"] = nil
		inputs["registrationVirtualNetworkIds"] = nil
		inputs["resolutionVirtualNetworkIds"] = nil
		inputs["resourceGroupName"] = nil
		inputs["tags"] = nil
		inputs["zoneType"] = nil
	} else {
		inputs["name"] = args.Name
		inputs["registrationVirtualNetworkIds"] = args.RegistrationVirtualNetworkIds
		inputs["resolutionVirtualNetworkIds"] = args.ResolutionVirtualNetworkIds
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["tags"] = args.Tags
		inputs["zoneType"] = args.ZoneType
	}
	inputs["maxNumberOfRecordSets"] = nil
	inputs["nameServers"] = nil
	inputs["numberOfRecordSets"] = nil
	s, err := ctx.RegisterResource("azure:dns/zone:Zone", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Zone{s: s}, nil
}

// GetZone gets an existing Zone resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetZone(ctx *pulumi.Context,
	name string, id pulumi.ID, state *ZoneState, opts ...pulumi.ResourceOpt) (*Zone, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["maxNumberOfRecordSets"] = state.MaxNumberOfRecordSets
		inputs["name"] = state.Name
		inputs["nameServers"] = state.NameServers
		inputs["numberOfRecordSets"] = state.NumberOfRecordSets
		inputs["registrationVirtualNetworkIds"] = state.RegistrationVirtualNetworkIds
		inputs["resolutionVirtualNetworkIds"] = state.ResolutionVirtualNetworkIds
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["tags"] = state.Tags
		inputs["zoneType"] = state.ZoneType
	}
	s, err := ctx.ReadResource("azure:dns/zone:Zone", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Zone{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Zone) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Zone) ID() *pulumi.IDOutput {
	return r.s.ID
}

// (Optional) Maximum number of Records in the zone. Defaults to `1000`.
func (r *Zone) MaxNumberOfRecordSets() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["maxNumberOfRecordSets"])
}

// The name of the DNS Zone. Must be a valid domain name.
func (r *Zone) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// (Optional) A list of values that make up the NS record for the zone.
func (r *Zone) NameServers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["nameServers"])
}

// (Optional) The number of records already in the zone.
func (r *Zone) NumberOfRecordSets() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["numberOfRecordSets"])
}

// A list of Virtual Network ID's that register hostnames in this DNS zone. This field can only be set when `zone_type` is set to `Private`.
func (r *Zone) RegistrationVirtualNetworkIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["registrationVirtualNetworkIds"])
}

// A list of Virtual Network ID's that resolve records in this DNS zone. This field can only be set when `zone_type` is set to `Private`.
func (r *Zone) ResolutionVirtualNetworkIds() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["resolutionVirtualNetworkIds"])
}

// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
func (r *Zone) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// A mapping of tags to assign to the resource.
func (r *Zone) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Specifies the type of this DNS zone. Possible values are `Public` or `Private` (Defaults to `Public`).
func (r *Zone) ZoneType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["zoneType"])
}

// Input properties used for looking up and filtering Zone resources.
type ZoneState struct {
	// (Optional) Maximum number of Records in the zone. Defaults to `1000`.
	MaxNumberOfRecordSets interface{}
	// The name of the DNS Zone. Must be a valid domain name.
	Name interface{}
	// (Optional) A list of values that make up the NS record for the zone.
	NameServers interface{}
	// (Optional) The number of records already in the zone.
	NumberOfRecordSets interface{}
	// A list of Virtual Network ID's that register hostnames in this DNS zone. This field can only be set when `zone_type` is set to `Private`.
	RegistrationVirtualNetworkIds interface{}
	// A list of Virtual Network ID's that resolve records in this DNS zone. This field can only be set when `zone_type` is set to `Private`.
	ResolutionVirtualNetworkIds interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the type of this DNS zone. Possible values are `Public` or `Private` (Defaults to `Public`).
	ZoneType interface{}
}

// The set of arguments for constructing a Zone resource.
type ZoneArgs struct {
	// The name of the DNS Zone. Must be a valid domain name.
	Name interface{}
	// A list of Virtual Network ID's that register hostnames in this DNS zone. This field can only be set when `zone_type` is set to `Private`.
	RegistrationVirtualNetworkIds interface{}
	// A list of Virtual Network ID's that resolve records in this DNS zone. This field can only be set when `zone_type` is set to `Private`.
	ResolutionVirtualNetworkIds interface{}
	// Specifies the resource group where the resource exists. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
	// Specifies the type of this DNS zone. Possible values are `Public` or `Private` (Defaults to `Public`).
	ZoneType interface{}
}