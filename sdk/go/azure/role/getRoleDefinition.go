// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package role

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access the properties of a custom Role Definition. To access information about a built-in Role Definition, [please see the `azurerm_builtin_role_definition` data source](builtin_role_definition.html) instead.
func LookupRoleDefinition(ctx *pulumi.Context, args *GetRoleDefinitionArgs) (*GetRoleDefinitionResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["roleDefinitionId"] = args.RoleDefinitionId
		inputs["scope"] = args.Scope
	}
	outputs, err := ctx.Invoke("azure:role/getRoleDefinition:getRoleDefinition", inputs)
	if err != nil {
		return nil, err
	}
	return &GetRoleDefinitionResult{
		AssignableScopes: outputs["assignableScopes"],
		Description: outputs["description"],
		Name: outputs["name"],
		Permissions: outputs["permissions"],
		Type: outputs["type"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getRoleDefinition.
type GetRoleDefinitionArgs struct {
	// Specifies the ID of the Role Definition as a UUID/GUID.
	RoleDefinitionId interface{}
	// Specifies the Scope at which the Custom Role Definition exists.
	Scope interface{}
}

// A collection of values returned by getRoleDefinition.
type GetRoleDefinitionResult struct {
	// One or more assignable scopes for this Role Definition, such as `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333`, `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup`, or `/subscriptions/0b1f6471-1bf0-4dda-aec3-111122223333/resourceGroups/myGroup/providers/Microsoft.Compute/virtualMachines/myVM`.
	AssignableScopes interface{}
	// the Description of the built-in Role.
	Description interface{}
	Name interface{}
	// a `permissions` block as documented below.
	Permissions interface{}
	// the Type of the Role.
	Type interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
