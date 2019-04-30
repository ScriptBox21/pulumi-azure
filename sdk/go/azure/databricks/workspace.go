// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Databricks Workspace
type Workspace struct {
	s *pulumi.ResourceState
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOpt) (*Workspace, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["location"] = nil
		inputs["managedResourceGroupName"] = nil
		inputs["name"] = nil
		inputs["resourceGroupName"] = nil
		inputs["sku"] = nil
		inputs["tags"] = nil
	} else {
		inputs["location"] = args.Location
		inputs["managedResourceGroupName"] = args.ManagedResourceGroupName
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["sku"] = args.Sku
		inputs["tags"] = args.Tags
	}
	inputs["managedResourceGroupId"] = nil
	s, err := ctx.RegisterResource("azure:databricks/workspace:Workspace", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Workspace{s: s}, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.ID, state *WorkspaceState, opts ...pulumi.ResourceOpt) (*Workspace, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["location"] = state.Location
		inputs["managedResourceGroupId"] = state.ManagedResourceGroupId
		inputs["managedResourceGroupName"] = state.ManagedResourceGroupName
		inputs["name"] = state.Name
		inputs["resourceGroupName"] = state.ResourceGroupName
		inputs["sku"] = state.Sku
		inputs["tags"] = state.Tags
	}
	s, err := ctx.ReadResource("azure:databricks/workspace:Workspace", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &Workspace{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *Workspace) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *Workspace) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
func (r *Workspace) Location() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["location"])
}

// The ID of the Managed Resource Group created by the Databricks Workspace.
func (r *Workspace) ManagedResourceGroupId() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["managedResourceGroupId"])
}

// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
func (r *Workspace) ManagedResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["managedResourceGroupName"])
}

// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
func (r *Workspace) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
func (r *Workspace) ResourceGroupName() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["resourceGroupName"])
}

// The `sku` to use for the Databricks Workspace. Possible values are `standard` or `premium`. Changing this forces a new resource to be created.
func (r *Workspace) Sku() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["sku"])
}

// A mapping of tags to assign to the resource.
func (r *Workspace) Tags() *pulumi.MapOutput {
	return (*pulumi.MapOutput)(r.s.State["tags"])
}

// Input properties used for looking up and filtering Workspace resources.
type WorkspaceState struct {
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location interface{}
	// The ID of the Managed Resource Group created by the Databricks Workspace.
	ManagedResourceGroupId interface{}
	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	ManagedResourceGroupName interface{}
	// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The `sku` to use for the Databricks Workspace. Possible values are `standard` or `premium`. Changing this forces a new resource to be created.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// Specifies the supported Azure location where the resource has to be created. Changing this forces a new resource to be created.
	Location interface{}
	// The name of the resource group where Azure should place the managed Databricks resources. Changing this forces a new resource to be created.
	ManagedResourceGroupName interface{}
	// Specifies the name of the Databricks Workspace resource. Changing this forces a new resource to be created.
	Name interface{}
	// The name of the Resource Group in which the Databricks Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName interface{}
	// The `sku` to use for the Databricks Workspace. Possible values are `standard` or `premium`. Changing this forces a new resource to be created.
	Sku interface{}
	// A mapping of tags to assign to the resource.
	Tags interface{}
}
