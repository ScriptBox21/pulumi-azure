// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package automation

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Automation Datetime Variable.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/automation_variable_datetime.html.markdown.
func LookupDateTimeVariable(ctx *pulumi.Context, args *LookupDateTimeVariableArgs, opts ...pulumi.InvokeOption) (*LookupDateTimeVariableResult, error) {
	var rv LookupDateTimeVariableResult
	err := ctx.Invoke("azure:automation/getDateTimeVariable:getDateTimeVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDateTimeVariable.
type LookupDateTimeVariableArgs struct {
	// The name of the automation account in which the Automation Variable exists.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of the Automation Variable.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the automation account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A collection of values returned by getDateTimeVariable.
type LookupDateTimeVariableResult struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description string `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted bool `pulumi:"encrypted"`
	// id is the provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The value of the Automation Variable in the [RFC3339 Section 5.6 Internet Date/Time Format](https://tools.ietf.org/html/rfc3339#section-5.6).
	Value string `pulumi:"value"`
}
