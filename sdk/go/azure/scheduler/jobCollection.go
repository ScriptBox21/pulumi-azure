// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package scheduler

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Scheduler Job Collection.
type JobCollection struct {
	s *pulumi.ResourceState
}

// NewJobCollection registers a new resource with the given unique name, arguments, and options.
func NewJobCollection(ctx *pulumi.Context,
	name string, args *JobCollectionArgs, opts ...pulumi.ResourceOpt) (*JobCollection, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["name"] = nil
		inputs["quota"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["state"] = nil
		inputs["tags"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["name"] = args.Name
		inputs["quota"] = args.Quota
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["state"] = args.State
		inputs["tags"] = args.Tags
	}
	s, err := ctx.RegisterResource("azure:scheduler/jobCollection:JobCollection", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobCollection{s: s}, nil
}

// GetJobCollection gets an existing JobCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobCollection(ctx *pulumi.Context,
	name string, id pulumi.ID, state *JobCollectionState, opts ...pulumi.ResourceOpt) (*JobCollection, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["name"] = state.Name
		inputs["quota"] = state.Quota
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sku"] = state.Sku
		inputs["state"] = state.State
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:scheduler/jobCollection:JobCollection", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &JobCollection{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *JobCollection) URN() *pulumi.URNOutput {
	return r.s.URN
}

// ID is this resource's unique identifier assigned by its provider.
func (r *JobCollection) ID() *pulumi.IDOutput {
	return r.s.ID
}

// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
func (r *JobCollection) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
func (r *JobCollection) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Configures the Job collection quotas as documented in the `quota` block below. 
func (r *JobCollection) Quota() *pulumi.Output {
	return r.s.State["quota"]
}

// The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
func (r *JobCollection) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
func (r *JobCollection) Sku() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sku"])
}

// Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
func (r *JobCollection) State() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["state"])
}

// A mapping of tags to assign to the resource.
func (r *JobCollection) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering JobCollection resources.
type JobCollectionState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
	Name interface{}
	// Configures the Job collection quotas as documented in the `quota` block below. 
	Quota interface{}
	// The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
	Sku interface{}
	// Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
	State interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a JobCollection resource.
type JobCollectionArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location interface{}
	// Specifies the name of the Scheduler Job Collection. Changing this forces a new resource to be created.
	Name interface{}
	// Configures the Job collection quotas as documented in the `quota` block below. 
	Quota interface{}
	// The name of the resource group in which to create the Scheduler Job Collection. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// Sets the Job Collection's pricing level's SKU. Possible values include: `Standard`, `Free`, `P10Premium`, `P20Premium`.
	Sku interface{}
	// Sets Job Collection's state. Possible values include: `Enabled`, `Disabled`, `Suspended`.
	State interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}