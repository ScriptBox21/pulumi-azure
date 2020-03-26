// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package monitoring

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the properties of a LogToMetricAction scheduled query rule.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/monitor_scheduled_query_rules_log.html.markdown.
func LookupScheduledQueryRulesLog(ctx *pulumi.Context, args *LookupScheduledQueryRulesLogArgs, opts ...pulumi.InvokeOption) (*LookupScheduledQueryRulesLogResult, error) {
	var rv LookupScheduledQueryRulesLogResult
	err := ctx.Invoke("azure:monitoring/getScheduledQueryRulesLog:getScheduledQueryRulesLog", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getScheduledQueryRulesLog.
type LookupScheduledQueryRulesLogArgs struct {
	// Specifies the name of the scheduled query rule.
	Name string `pulumi:"name"`
	// Specifies the name of the resource group where the scheduled query rule is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getScheduledQueryRulesLog.
type LookupScheduledQueryRulesLogResult struct {
	AuthorizedResourceIds []string `pulumi:"authorizedResourceIds"`
	// A `criteria` block as defined below.
	Criterias []GetScheduledQueryRulesLogCriteria `pulumi:"criterias"`
	// The resource URI over which log search query is to be run.
	DataSourceId string `pulumi:"dataSourceId"`
	// The description of the scheduled query rule.
	Description string `pulumi:"description"`
	// Whether this scheduled query rule is enabled.
	Enabled bool `pulumi:"enabled"`
	// id is the provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
	// Name of the dimension.
	Name              string            `pulumi:"name"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
}
