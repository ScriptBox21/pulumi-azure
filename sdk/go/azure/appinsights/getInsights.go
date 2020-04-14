// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appinsights

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to access information about an existing Application Insights component.
func LookupInsights(ctx *pulumi.Context, args *LookupInsightsArgs, opts ...pulumi.InvokeOption) (*LookupInsightsResult, error) {
	var rv LookupInsightsResult
	err := ctx.Invoke("azure:appinsights/getInsights:getInsights", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInsights.
type LookupInsightsArgs struct {
	// Specifies the name of the Application Insights component.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Application Insights component is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getInsights.
type LookupInsightsResult struct {
	// The App ID associated with this Application Insights component.
	AppId string `pulumi:"appId"`
	// The type of the component.
	ApplicationType string `pulumi:"applicationType"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The instrumentation key of the Application Insights component.
	InstrumentationKey string `pulumi:"instrumentationKey"`
	// The Azure location where the component exists.
	Location          string `pulumi:"location"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The retention period in days.
	RetentionInDays int `pulumi:"retentionInDays"`
	// Tags applied to the component.
	Tags map[string]string `pulumi:"tags"`
}
