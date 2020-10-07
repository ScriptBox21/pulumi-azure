// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mariadb

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Sets a MariaDB Configuration value on a MariaDB Server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/mariadb"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("West Europe"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleServer, err := mariadb.NewServer(ctx, "exampleServer", &mariadb.ServerArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			SkuName:           pulumi.String("B_Gen5_2"),
// 			StorageProfile: &mariadb.ServerStorageProfileArgs{
// 				StorageMb:           pulumi.Int(5120),
// 				BackupRetentionDays: pulumi.Int(7),
// 				GeoRedundantBackup:  pulumi.String("Disabled"),
// 			},
// 			AdministratorLogin:         pulumi.String("mariadbadmin"),
// 			AdministratorLoginPassword: pulumi.String("H@Sh1CoR3!"),
// 			Version:                    pulumi.String("10.2"),
// 			SslEnforcement:             pulumi.String("Enabled"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = mariadb.NewConfiguration(ctx, "exampleConfiguration", &mariadb.ConfigurationArgs{
// 			Name:              pulumi.String("interactive_timeout"),
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			ServerName:        exampleServer.Name,
// 			Value:             pulumi.String("600"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Configuration struct {
	pulumi.CustomResourceState

	// Specifies the name of the MariaDB Configuration, which needs [to be a valid MariaDB configuration name](https://mariadb.com/kb/en/library/server-system-variables/). Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values.
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewConfiguration registers a new resource with the given unique name, arguments, and options.
func NewConfiguration(ctx *pulumi.Context,
	name string, args *ConfigurationArgs, opts ...pulumi.ResourceOption) (*Configuration, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServerName == nil {
		return nil, errors.New("missing required argument 'ServerName'")
	}
	if args == nil || args.Value == nil {
		return nil, errors.New("missing required argument 'Value'")
	}
	if args == nil {
		args = &ConfigurationArgs{}
	}
	var resource Configuration
	err := ctx.RegisterResource("azure:mariadb/configuration:Configuration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfiguration gets an existing Configuration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationState, opts ...pulumi.ResourceOption) (*Configuration, error) {
	var resource Configuration
	err := ctx.ReadResource("azure:mariadb/configuration:Configuration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Configuration resources.
type configurationState struct {
	// Specifies the name of the MariaDB Configuration, which needs [to be a valid MariaDB configuration name](https://mariadb.com/kb/en/library/server-system-variables/). Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values.
	Value *string `pulumi:"value"`
}

type ConfigurationState struct {
	// Specifies the name of the MariaDB Configuration, which needs [to be a valid MariaDB configuration name](https://mariadb.com/kb/en/library/server-system-variables/). Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values.
	Value pulumi.StringPtrInput
}

func (ConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationState)(nil)).Elem()
}

type configurationArgs struct {
	// Specifies the name of the MariaDB Configuration, which needs [to be a valid MariaDB configuration name](https://mariadb.com/kb/en/library/server-system-variables/). Changing this forces a new resource to be created.
	Name string `pulumi:"name"`
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Configuration resource.
type ConfigurationArgs struct {
	// Specifies the name of the MariaDB Configuration, which needs [to be a valid MariaDB configuration name](https://mariadb.com/kb/en/library/server-system-variables/). Changing this forces a new resource to be created.
	Name pulumi.StringInput
	// The name of the resource group in which the MariaDB Server exists. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// Specifies the value of the MariaDB Configuration. See the MariaDB documentation for valid values.
	Value pulumi.StringInput
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationArgs)(nil)).Elem()
}
