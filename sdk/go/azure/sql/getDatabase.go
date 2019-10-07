// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Use this data source to access information about an existing SQL Azure Database.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/sql_database.html.markdown.
func LookupDatabase(ctx *pulumi.Context, args *GetDatabaseArgs) (*GetDatabaseResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["name"] = args.Name
		inputs["resourceGroupName"] = args.ResourceGroupName
		inputs["serverName"] = args.ServerName
		inputs["tags"] = args.Tags
	}
	outputs, err := ctx.Invoke("azure:sql/getDatabase:getDatabase", inputs)
	if err != nil {
		return nil, err
	}
	return &GetDatabaseResult{
		Collation: outputs["collation"],
		DefaultSecondaryLocation: outputs["defaultSecondaryLocation"],
		Edition: outputs["edition"],
		ElasticPoolName: outputs["elasticPoolName"],
		FailoverGroupId: outputs["failoverGroupId"],
		Location: outputs["location"],
		Name: outputs["name"],
		ReadScale: outputs["readScale"],
		ResourceGroupName: outputs["resourceGroupName"],
		ServerName: outputs["serverName"],
		Tags: outputs["tags"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getDatabase.
type GetDatabaseArgs struct {
	// The name of the SQL Database.
	Name interface{}
	// Specifies the name of the Resource Group where the Azure SQL Database exists.
	ResourceGroupName interface{}
	// The name of the SQL Server.
	ServerName interface{}
	Tags interface{}
}

// A collection of values returned by getDatabase.
type GetDatabaseResult struct {
	// The name of the collation. 
	Collation interface{}
	// The default secondary location of the SQL Database.
	DefaultSecondaryLocation interface{}
	// The edition of the database.
	Edition interface{}
	// The name of the elastic database pool the database belongs to.
	ElasticPoolName interface{}
	// The ID of the failover group the database belongs to.
	FailoverGroupId interface{}
	// The location of the Resource Group in which the SQL Server exists.
	Location interface{}
	// The name of the database.
	Name interface{}
	// Indicate if read-only connections will be redirected to a high-available replica.
	ReadScale interface{}
	// The name of the resource group in which the database resides. This will always be the same resource group as the Database Server.
	ResourceGroupName interface{}
	// The name of the SQL Server on which to create the database.
	ServerName interface{}
	// A mapping of tags assigned to the resource.
	Tags interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}