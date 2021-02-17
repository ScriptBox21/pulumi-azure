// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Linked Service (connection) between Azure SQL Database and Azure Data Factory.
//
// > **Note:** All arguments including the connectionString will be stored in the raw state as plain-text. [Read more about sensitive data in state](https://www.terraform.io/docs/state/sensitive-data.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/datafactory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("northeurope"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleFactory, err := datafactory.NewFactory(ctx, "exampleFactory", &datafactory.FactoryArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewLinkedServiceAzureSqlDatabase(ctx, "exampleLinkedServiceAzureSqlDatabase", &datafactory.LinkedServiceAzureSqlDatabaseArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			ConnectionString:  pulumi.String("data source=serverhostname;initial catalog=master;user id=testUser;Password=test;integrated security=False;encrypt=True;connection timeout=30"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Data Factory Azure SQL Database Linked Service's can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/linkedServiceAzureSqlDatabase:LinkedServiceAzureSqlDatabase example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/linkedservices/example
// ```
type LinkedServiceAzureSqlDatabase struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Linked Service Azure SQL Database.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Azure SQL Database.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The connection string in which to authenticate with Azure SQL Database.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Azure SQL Database.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Azure SQL Database.
	IntegrationRuntimeName pulumi.StringPtrOutput `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Azure SQL Database.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewLinkedServiceAzureSqlDatabase registers a new resource with the given unique name, arguments, and options.
func NewLinkedServiceAzureSqlDatabase(ctx *pulumi.Context,
	name string, args *LinkedServiceAzureSqlDatabaseArgs, opts ...pulumi.ResourceOption) (*LinkedServiceAzureSqlDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionString == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionString'")
	}
	if args.DataFactoryName == nil {
		return nil, errors.New("invalid value for required argument 'DataFactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource LinkedServiceAzureSqlDatabase
	err := ctx.RegisterResource("azure:datafactory/linkedServiceAzureSqlDatabase:LinkedServiceAzureSqlDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedServiceAzureSqlDatabase gets an existing LinkedServiceAzureSqlDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedServiceAzureSqlDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceAzureSqlDatabaseState, opts ...pulumi.ResourceOption) (*LinkedServiceAzureSqlDatabase, error) {
	var resource LinkedServiceAzureSqlDatabase
	err := ctx.ReadResource("azure:datafactory/linkedServiceAzureSqlDatabase:LinkedServiceAzureSqlDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedServiceAzureSqlDatabase resources.
type linkedServiceAzureSqlDatabaseState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Azure SQL Database.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Azure SQL Database.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with Azure SQL Database.
	ConnectionString *string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource to be created.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Azure SQL Database.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Azure SQL Database.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Azure SQL Database.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type LinkedServiceAzureSqlDatabaseState struct {
	// A map of additional properties to associate with the Data Factory Linked Service Azure SQL Database.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Azure SQL Database.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with Azure SQL Database.
	ConnectionString pulumi.StringPtrInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Linked Service Azure SQL Database.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Azure SQL Database.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Azure SQL Database.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (LinkedServiceAzureSqlDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureSqlDatabaseState)(nil)).Elem()
}

type linkedServiceAzureSqlDatabaseArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Azure SQL Database.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Linked Service Azure SQL Database.
	Annotations []string `pulumi:"annotations"`
	// The connection string in which to authenticate with Azure SQL Database.
	ConnectionString string `pulumi:"connectionString"`
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource to be created.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Linked Service Azure SQL Database.
	Description *string `pulumi:"description"`
	// The integration runtime reference to associate with the Data Factory Linked Service Azure SQL Database.
	IntegrationRuntimeName *string `pulumi:"integrationRuntimeName"`
	// Specifies the name of the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Linked Service Azure SQL Database.
	Parameters map[string]string `pulumi:"parameters"`
	// The name of the resource group in which to create the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a LinkedServiceAzureSqlDatabase resource.
type LinkedServiceAzureSqlDatabaseArgs struct {
	// A map of additional properties to associate with the Data Factory Linked Service Azure SQL Database.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Linked Service Azure SQL Database.
	Annotations pulumi.StringArrayInput
	// The connection string in which to authenticate with Azure SQL Database.
	ConnectionString pulumi.StringInput
	// The Data Factory name in which to associate the Linked Service with. Changing this forces a new resource to be created.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Linked Service Azure SQL Database.
	Description pulumi.StringPtrInput
	// The integration runtime reference to associate with the Data Factory Linked Service Azure SQL Database.
	IntegrationRuntimeName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created. Must be unique within a data
	// factory. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Linked Service Azure SQL Database.
	Parameters pulumi.StringMapInput
	// The name of the resource group in which to create the Data Factory Linked Service Azure SQL Database. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (LinkedServiceAzureSqlDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceAzureSqlDatabaseArgs)(nil)).Elem()
}

type LinkedServiceAzureSqlDatabaseInput interface {
	pulumi.Input

	ToLinkedServiceAzureSqlDatabaseOutput() LinkedServiceAzureSqlDatabaseOutput
	ToLinkedServiceAzureSqlDatabaseOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabaseOutput
}

func (*LinkedServiceAzureSqlDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceAzureSqlDatabase)(nil))
}

func (i *LinkedServiceAzureSqlDatabase) ToLinkedServiceAzureSqlDatabaseOutput() LinkedServiceAzureSqlDatabaseOutput {
	return i.ToLinkedServiceAzureSqlDatabaseOutputWithContext(context.Background())
}

func (i *LinkedServiceAzureSqlDatabase) ToLinkedServiceAzureSqlDatabaseOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureSqlDatabaseOutput)
}

func (i *LinkedServiceAzureSqlDatabase) ToLinkedServiceAzureSqlDatabasePtrOutput() LinkedServiceAzureSqlDatabasePtrOutput {
	return i.ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(context.Background())
}

func (i *LinkedServiceAzureSqlDatabase) ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureSqlDatabasePtrOutput)
}

type LinkedServiceAzureSqlDatabasePtrInput interface {
	pulumi.Input

	ToLinkedServiceAzureSqlDatabasePtrOutput() LinkedServiceAzureSqlDatabasePtrOutput
	ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabasePtrOutput
}

type linkedServiceAzureSqlDatabasePtrType LinkedServiceAzureSqlDatabaseArgs

func (*linkedServiceAzureSqlDatabasePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceAzureSqlDatabase)(nil))
}

func (i *linkedServiceAzureSqlDatabasePtrType) ToLinkedServiceAzureSqlDatabasePtrOutput() LinkedServiceAzureSqlDatabasePtrOutput {
	return i.ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(context.Background())
}

func (i *linkedServiceAzureSqlDatabasePtrType) ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabasePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureSqlDatabasePtrOutput)
}

// LinkedServiceAzureSqlDatabaseArrayInput is an input type that accepts LinkedServiceAzureSqlDatabaseArray and LinkedServiceAzureSqlDatabaseArrayOutput values.
// You can construct a concrete instance of `LinkedServiceAzureSqlDatabaseArrayInput` via:
//
//          LinkedServiceAzureSqlDatabaseArray{ LinkedServiceAzureSqlDatabaseArgs{...} }
type LinkedServiceAzureSqlDatabaseArrayInput interface {
	pulumi.Input

	ToLinkedServiceAzureSqlDatabaseArrayOutput() LinkedServiceAzureSqlDatabaseArrayOutput
	ToLinkedServiceAzureSqlDatabaseArrayOutputWithContext(context.Context) LinkedServiceAzureSqlDatabaseArrayOutput
}

type LinkedServiceAzureSqlDatabaseArray []LinkedServiceAzureSqlDatabaseInput

func (LinkedServiceAzureSqlDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*LinkedServiceAzureSqlDatabase)(nil))
}

func (i LinkedServiceAzureSqlDatabaseArray) ToLinkedServiceAzureSqlDatabaseArrayOutput() LinkedServiceAzureSqlDatabaseArrayOutput {
	return i.ToLinkedServiceAzureSqlDatabaseArrayOutputWithContext(context.Background())
}

func (i LinkedServiceAzureSqlDatabaseArray) ToLinkedServiceAzureSqlDatabaseArrayOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureSqlDatabaseArrayOutput)
}

// LinkedServiceAzureSqlDatabaseMapInput is an input type that accepts LinkedServiceAzureSqlDatabaseMap and LinkedServiceAzureSqlDatabaseMapOutput values.
// You can construct a concrete instance of `LinkedServiceAzureSqlDatabaseMapInput` via:
//
//          LinkedServiceAzureSqlDatabaseMap{ "key": LinkedServiceAzureSqlDatabaseArgs{...} }
type LinkedServiceAzureSqlDatabaseMapInput interface {
	pulumi.Input

	ToLinkedServiceAzureSqlDatabaseMapOutput() LinkedServiceAzureSqlDatabaseMapOutput
	ToLinkedServiceAzureSqlDatabaseMapOutputWithContext(context.Context) LinkedServiceAzureSqlDatabaseMapOutput
}

type LinkedServiceAzureSqlDatabaseMap map[string]LinkedServiceAzureSqlDatabaseInput

func (LinkedServiceAzureSqlDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*LinkedServiceAzureSqlDatabase)(nil))
}

func (i LinkedServiceAzureSqlDatabaseMap) ToLinkedServiceAzureSqlDatabaseMapOutput() LinkedServiceAzureSqlDatabaseMapOutput {
	return i.ToLinkedServiceAzureSqlDatabaseMapOutputWithContext(context.Background())
}

func (i LinkedServiceAzureSqlDatabaseMap) ToLinkedServiceAzureSqlDatabaseMapOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedServiceAzureSqlDatabaseMapOutput)
}

type LinkedServiceAzureSqlDatabaseOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceAzureSqlDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedServiceAzureSqlDatabase)(nil))
}

func (o LinkedServiceAzureSqlDatabaseOutput) ToLinkedServiceAzureSqlDatabaseOutput() LinkedServiceAzureSqlDatabaseOutput {
	return o
}

func (o LinkedServiceAzureSqlDatabaseOutput) ToLinkedServiceAzureSqlDatabaseOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabaseOutput {
	return o
}

func (o LinkedServiceAzureSqlDatabaseOutput) ToLinkedServiceAzureSqlDatabasePtrOutput() LinkedServiceAzureSqlDatabasePtrOutput {
	return o.ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(context.Background())
}

func (o LinkedServiceAzureSqlDatabaseOutput) ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabasePtrOutput {
	return o.ApplyT(func(v LinkedServiceAzureSqlDatabase) *LinkedServiceAzureSqlDatabase {
		return &v
	}).(LinkedServiceAzureSqlDatabasePtrOutput)
}

type LinkedServiceAzureSqlDatabasePtrOutput struct {
	*pulumi.OutputState
}

func (LinkedServiceAzureSqlDatabasePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinkedServiceAzureSqlDatabase)(nil))
}

func (o LinkedServiceAzureSqlDatabasePtrOutput) ToLinkedServiceAzureSqlDatabasePtrOutput() LinkedServiceAzureSqlDatabasePtrOutput {
	return o
}

func (o LinkedServiceAzureSqlDatabasePtrOutput) ToLinkedServiceAzureSqlDatabasePtrOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabasePtrOutput {
	return o
}

type LinkedServiceAzureSqlDatabaseArrayOutput struct{ *pulumi.OutputState }

func (LinkedServiceAzureSqlDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LinkedServiceAzureSqlDatabase)(nil))
}

func (o LinkedServiceAzureSqlDatabaseArrayOutput) ToLinkedServiceAzureSqlDatabaseArrayOutput() LinkedServiceAzureSqlDatabaseArrayOutput {
	return o
}

func (o LinkedServiceAzureSqlDatabaseArrayOutput) ToLinkedServiceAzureSqlDatabaseArrayOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabaseArrayOutput {
	return o
}

func (o LinkedServiceAzureSqlDatabaseArrayOutput) Index(i pulumi.IntInput) LinkedServiceAzureSqlDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LinkedServiceAzureSqlDatabase {
		return vs[0].([]LinkedServiceAzureSqlDatabase)[vs[1].(int)]
	}).(LinkedServiceAzureSqlDatabaseOutput)
}

type LinkedServiceAzureSqlDatabaseMapOutput struct{ *pulumi.OutputState }

func (LinkedServiceAzureSqlDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]LinkedServiceAzureSqlDatabase)(nil))
}

func (o LinkedServiceAzureSqlDatabaseMapOutput) ToLinkedServiceAzureSqlDatabaseMapOutput() LinkedServiceAzureSqlDatabaseMapOutput {
	return o
}

func (o LinkedServiceAzureSqlDatabaseMapOutput) ToLinkedServiceAzureSqlDatabaseMapOutputWithContext(ctx context.Context) LinkedServiceAzureSqlDatabaseMapOutput {
	return o
}

func (o LinkedServiceAzureSqlDatabaseMapOutput) MapIndex(k pulumi.StringInput) LinkedServiceAzureSqlDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) LinkedServiceAzureSqlDatabase {
		return vs[0].(map[string]LinkedServiceAzureSqlDatabase)[vs[1].(string)]
	}).(LinkedServiceAzureSqlDatabaseOutput)
}

func init() {
	pulumi.RegisterOutputType(LinkedServiceAzureSqlDatabaseOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureSqlDatabasePtrOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureSqlDatabaseArrayOutput{})
	pulumi.RegisterOutputType(LinkedServiceAzureSqlDatabaseMapOutput{})
}
