// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package automation

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing Automation Int Variable.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/automation_variable_int.html.markdown.
func LookupIntVariable(ctx *pulumi.Context, args *LookupIntVariableArgs, opts ...pulumi.InvokeOption) (*LookupIntVariableResult, error) {
	var rv LookupIntVariableResult
	err := ctx.Invoke("azure:automation/getIntVariable:getIntVariable", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIntVariable.
type LookupIntVariableArgs struct {
	// The name of the automation account in which the Automation Variable exists.
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The name of the Automation Variable.
	Name string `pulumi:"name"`
	// The Name of the Resource Group where the automation account exists.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


// A collection of values returned by getIntVariable.
type LookupIntVariableResult struct {
	AutomationAccountName string `pulumi:"automationAccountName"`
	// The description of the Automation Variable.
	Description string `pulumi:"description"`
	// Specifies if the Automation Variable is encrypted. Defaults to `false`.
	Encrypted bool `pulumi:"encrypted"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	Name string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The value of the Automation Variable as a `integer`.
	Value int `pulumi:"value"`
}

