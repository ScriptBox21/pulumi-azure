// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package streamanalytics

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Stream Analytics Job.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/stream_analytics_job.html.markdown.
func LookupJob(ctx *pulumi.Context, args *LookupJobArgs, opts ...pulumi.InvokeOption) (*LookupJobResult, error) {
	var rv LookupJobResult
	err := ctx.Invoke("azure:streamanalytics/getJob:getJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getJob.
type LookupJobArgs struct {
	// Specifies the name of the Stream Analytics Job.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group the Stream Analytics Job is located in.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getJob.
type LookupJobResult struct {
	// The compatibility level for this job.
	CompatibilityLevel string `pulumi:"compatibilityLevel"`
	// The Data Locale of the Job.
	DataLocale string `pulumi:"dataLocale"`
	// The maximum tolerable delay in seconds where events arriving late could be included.
	EventsLateArrivalMaxDelayInSeconds int `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	// The maximum tolerable delay in seconds where out-of-order events can be adjusted to be back in order.
	EventsOutOfOrderMaxDelayInSeconds int `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	// The policy which should be applied to events which arrive out of order in the input event stream.
	EventsOutOfOrderPolicy string `pulumi:"eventsOutOfOrderPolicy"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The Job ID assigned by the Stream Analytics Job.
	JobId string `pulumi:"jobId"`
	// The Azure location where the Stream Analytics Job exists.
	Location string `pulumi:"location"`
	Name     string `pulumi:"name"`
	// The policy which should be applied to events which arrive at the output and cannot be written to the external storage due to being malformed (such as missing column values, column values of wrong type or size).
	OutputErrorPolicy string `pulumi:"outputErrorPolicy"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The number of streaming units that the streaming job uses.
	StreamingUnits int `pulumi:"streamingUnits"`
	// The query that will be run in the streaming job, [written in Stream Analytics Query Language (SAQL)](https://msdn.microsoft.com/library/azure/dn834998).
	TransformationQuery string `pulumi:"transformationQuery"`
}
