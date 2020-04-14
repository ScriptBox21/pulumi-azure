// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Machine Learning Workspace.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure:machinelearning/getWorkspace:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getWorkspace.
type LookupWorkspaceArgs struct {
	// The name of the Machine Learning Workspace exists.
	Name string `pulumi:"name"`
	// The name of the Resource Group where the Machine Learning Workspace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getWorkspace.
type LookupWorkspaceResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The location where the Machine Learning Workspace exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags assigned to the Machine Learning Workspace.
	Tags map[string]string `pulumi:"tags"`
}
