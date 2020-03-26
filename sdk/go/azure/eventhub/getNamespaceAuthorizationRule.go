// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package eventhub

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an Authorization Rule for an Event Hub Namespace.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/eventhub_namespace_authorization_rule.html.markdown.
func LookupNamespaceAuthorizationRule(ctx *pulumi.Context, args *LookupNamespaceAuthorizationRuleArgs, opts ...pulumi.InvokeOption) (*LookupNamespaceAuthorizationRuleResult, error) {
	var rv LookupNamespaceAuthorizationRuleResult
	err := ctx.Invoke("azure:eventhub/getNamespaceAuthorizationRule:getNamespaceAuthorizationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNamespaceAuthorizationRule.
type LookupNamespaceAuthorizationRuleArgs struct {
	// The name of the EventHub Authorization Rule resource.
	Name string `pulumi:"name"`
	// The name of the EventHub Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The name of the resource group in which the EventHub Namespace exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getNamespaceAuthorizationRule.
type LookupNamespaceAuthorizationRuleResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Does this Authorization Rule have permissions to Listen to the Event Hub?
	Listen bool `pulumi:"listen"`
	// Does this Authorization Rule have permissions to Manage to the Event Hub?
	Manage bool   `pulumi:"manage"`
	Name   string `pulumi:"name"`
	// The name of the EventHub Namespace.
	NamespaceName string `pulumi:"namespaceName"`
	// The Primary Connection String for the Event Hubs authorization Rule.
	PrimaryConnectionString string `pulumi:"primaryConnectionString"`
	// The Primary Key for the Event Hubs authorization Rule.
	PrimaryKey        string `pulumi:"primaryKey"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The Secondary Connection String for the Event Hubs authorization Rule.
	SecondaryConnectionString string `pulumi:"secondaryConnectionString"`
	// The Secondary Key for the Event Hubs authorization Rule.
	SecondaryKey string `pulumi:"secondaryKey"`
	// Does this Authorization Rule have permissions to Send to the Event Hub?
	Send bool `pulumi:"send"`
}
