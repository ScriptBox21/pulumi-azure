// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package monitoring

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about the Monitor Diagnostics Categories supported by an existing Resource.
func GetDiagnosticCategories(ctx *pulumi.Context, args *GetDiagnosticCategoriesArgs, opts ...pulumi.InvokeOption) (*GetDiagnosticCategoriesResult, error) {
	var rv GetDiagnosticCategoriesResult
	err := ctx.Invoke("azure:monitoring/getDiagnosticCategories:getDiagnosticCategories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDiagnosticCategories.
type GetDiagnosticCategoriesArgs struct {
	// The ID of an existing Resource which Monitor Diagnostics Categories should be retrieved for.
	ResourceId string `pulumi:"resourceId"`
}

// A collection of values returned by getDiagnosticCategories.
type GetDiagnosticCategoriesResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of the Log Categories supported for this Resource.
	Logs []string `pulumi:"logs"`
	// A list of the Metric Categories supported for this Resource.
	Metrics    []string `pulumi:"metrics"`
	ResourceId string   `pulumi:"resourceId"`
}
