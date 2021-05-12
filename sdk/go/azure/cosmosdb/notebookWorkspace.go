// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cosmosdb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an SQL Notebook Workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v4/go/azure/cosmosdb"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
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
// 		exampleAccount, err := cosmosdb.NewAccount(ctx, "exampleAccount", &cosmosdb.AccountArgs{
// 			Location:          exampleResourceGroup.Location,
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			OfferType:         pulumi.String("Standard"),
// 			Kind:              pulumi.String("GlobalDocumentDB"),
// 			ConsistencyPolicy: &cosmosdb.AccountConsistencyPolicyArgs{
// 				ConsistencyLevel: pulumi.String("BoundedStaleness"),
// 			},
// 			GeoLocations: cosmosdb.AccountGeoLocationArray{
// 				&cosmosdb.AccountGeoLocationArgs{
// 					Location:         exampleResourceGroup.Location,
// 					FailoverPriority: pulumi.Int(0),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cosmosdb.NewNotebookWorkspace(ctx, "exampleNotebookWorkspace", &cosmosdb.NotebookWorkspaceArgs{
// 			ResourceGroupName: exampleAccount.ResourceGroupName,
// 			AccountName:       exampleAccount.Name,
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
// =SQL Notebook Workspaces can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:cosmosdb/notebookWorkspace:NotebookWorkspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.DocumentDB/databaseAccounts/account1/notebookWorkspaces/notebookWorkspace1
// ```
type NotebookWorkspace struct {
	pulumi.CustomResourceState

	// The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
	AccountName pulumi.StringOutput `pulumi:"accountName"`
	// The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies the endpoint of Notebook server.
	ServerEndpoint pulumi.StringOutput `pulumi:"serverEndpoint"`
}

// NewNotebookWorkspace registers a new resource with the given unique name, arguments, and options.
func NewNotebookWorkspace(ctx *pulumi.Context,
	name string, args *NotebookWorkspaceArgs, opts ...pulumi.ResourceOption) (*NotebookWorkspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource NotebookWorkspace
	err := ctx.RegisterResource("azure:cosmosdb/notebookWorkspace:NotebookWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebookWorkspace gets an existing NotebookWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebookWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebookWorkspaceState, opts ...pulumi.ResourceOption) (*NotebookWorkspace, error) {
	var resource NotebookWorkspace
	err := ctx.ReadResource("azure:cosmosdb/notebookWorkspace:NotebookWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebookWorkspace resources.
type notebookWorkspaceState struct {
	// The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
	AccountName *string `pulumi:"accountName"`
	// The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies the endpoint of Notebook server.
	ServerEndpoint *string `pulumi:"serverEndpoint"`
}

type NotebookWorkspaceState struct {
	// The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
	AccountName pulumi.StringPtrInput
	// The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies the endpoint of Notebook server.
	ServerEndpoint pulumi.StringPtrInput
}

func (NotebookWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookWorkspaceState)(nil)).Elem()
}

type notebookWorkspaceArgs struct {
	// The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
	AccountName string `pulumi:"accountName"`
	// The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
	Name *string `pulumi:"name"`
	// The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a NotebookWorkspace resource.
type NotebookWorkspaceArgs struct {
	// The name of the Cosmos DB Account to create the SQL Notebook Workspace within. Changing this forces a new SQL Notebook Workspace to be created.
	AccountName pulumi.StringInput
	// The name which should be used for this SQL Notebook Workspace. Possible value is `default`. Changing this forces a new SQL Notebook Workspace to be created.
	Name pulumi.StringPtrInput
	// The name of the Resource Group where the SQL Notebook Workspace should exist. Changing this forces a new SQL Notebook Workspace to be created.
	ResourceGroupName pulumi.StringInput
}

func (NotebookWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebookWorkspaceArgs)(nil)).Elem()
}

type NotebookWorkspaceInput interface {
	pulumi.Input

	ToNotebookWorkspaceOutput() NotebookWorkspaceOutput
	ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput
}

func (*NotebookWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookWorkspace)(nil))
}

func (i *NotebookWorkspace) ToNotebookWorkspaceOutput() NotebookWorkspaceOutput {
	return i.ToNotebookWorkspaceOutputWithContext(context.Background())
}

func (i *NotebookWorkspace) ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookWorkspaceOutput)
}

func (i *NotebookWorkspace) ToNotebookWorkspacePtrOutput() NotebookWorkspacePtrOutput {
	return i.ToNotebookWorkspacePtrOutputWithContext(context.Background())
}

func (i *NotebookWorkspace) ToNotebookWorkspacePtrOutputWithContext(ctx context.Context) NotebookWorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookWorkspacePtrOutput)
}

type NotebookWorkspacePtrInput interface {
	pulumi.Input

	ToNotebookWorkspacePtrOutput() NotebookWorkspacePtrOutput
	ToNotebookWorkspacePtrOutputWithContext(ctx context.Context) NotebookWorkspacePtrOutput
}

type notebookWorkspacePtrType NotebookWorkspaceArgs

func (*notebookWorkspacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookWorkspace)(nil))
}

func (i *notebookWorkspacePtrType) ToNotebookWorkspacePtrOutput() NotebookWorkspacePtrOutput {
	return i.ToNotebookWorkspacePtrOutputWithContext(context.Background())
}

func (i *notebookWorkspacePtrType) ToNotebookWorkspacePtrOutputWithContext(ctx context.Context) NotebookWorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookWorkspacePtrOutput)
}

// NotebookWorkspaceArrayInput is an input type that accepts NotebookWorkspaceArray and NotebookWorkspaceArrayOutput values.
// You can construct a concrete instance of `NotebookWorkspaceArrayInput` via:
//
//          NotebookWorkspaceArray{ NotebookWorkspaceArgs{...} }
type NotebookWorkspaceArrayInput interface {
	pulumi.Input

	ToNotebookWorkspaceArrayOutput() NotebookWorkspaceArrayOutput
	ToNotebookWorkspaceArrayOutputWithContext(context.Context) NotebookWorkspaceArrayOutput
}

type NotebookWorkspaceArray []NotebookWorkspaceInput

func (NotebookWorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*NotebookWorkspace)(nil))
}

func (i NotebookWorkspaceArray) ToNotebookWorkspaceArrayOutput() NotebookWorkspaceArrayOutput {
	return i.ToNotebookWorkspaceArrayOutputWithContext(context.Background())
}

func (i NotebookWorkspaceArray) ToNotebookWorkspaceArrayOutputWithContext(ctx context.Context) NotebookWorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookWorkspaceArrayOutput)
}

// NotebookWorkspaceMapInput is an input type that accepts NotebookWorkspaceMap and NotebookWorkspaceMapOutput values.
// You can construct a concrete instance of `NotebookWorkspaceMapInput` via:
//
//          NotebookWorkspaceMap{ "key": NotebookWorkspaceArgs{...} }
type NotebookWorkspaceMapInput interface {
	pulumi.Input

	ToNotebookWorkspaceMapOutput() NotebookWorkspaceMapOutput
	ToNotebookWorkspaceMapOutputWithContext(context.Context) NotebookWorkspaceMapOutput
}

type NotebookWorkspaceMap map[string]NotebookWorkspaceInput

func (NotebookWorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*NotebookWorkspace)(nil))
}

func (i NotebookWorkspaceMap) ToNotebookWorkspaceMapOutput() NotebookWorkspaceMapOutput {
	return i.ToNotebookWorkspaceMapOutputWithContext(context.Background())
}

func (i NotebookWorkspaceMap) ToNotebookWorkspaceMapOutputWithContext(ctx context.Context) NotebookWorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebookWorkspaceMapOutput)
}

type NotebookWorkspaceOutput struct {
	*pulumi.OutputState
}

func (NotebookWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotebookWorkspace)(nil))
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspaceOutput() NotebookWorkspaceOutput {
	return o
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspaceOutputWithContext(ctx context.Context) NotebookWorkspaceOutput {
	return o
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspacePtrOutput() NotebookWorkspacePtrOutput {
	return o.ToNotebookWorkspacePtrOutputWithContext(context.Background())
}

func (o NotebookWorkspaceOutput) ToNotebookWorkspacePtrOutputWithContext(ctx context.Context) NotebookWorkspacePtrOutput {
	return o.ApplyT(func(v NotebookWorkspace) *NotebookWorkspace {
		return &v
	}).(NotebookWorkspacePtrOutput)
}

type NotebookWorkspacePtrOutput struct {
	*pulumi.OutputState
}

func (NotebookWorkspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebookWorkspace)(nil))
}

func (o NotebookWorkspacePtrOutput) ToNotebookWorkspacePtrOutput() NotebookWorkspacePtrOutput {
	return o
}

func (o NotebookWorkspacePtrOutput) ToNotebookWorkspacePtrOutputWithContext(ctx context.Context) NotebookWorkspacePtrOutput {
	return o
}

type NotebookWorkspaceArrayOutput struct{ *pulumi.OutputState }

func (NotebookWorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NotebookWorkspace)(nil))
}

func (o NotebookWorkspaceArrayOutput) ToNotebookWorkspaceArrayOutput() NotebookWorkspaceArrayOutput {
	return o
}

func (o NotebookWorkspaceArrayOutput) ToNotebookWorkspaceArrayOutputWithContext(ctx context.Context) NotebookWorkspaceArrayOutput {
	return o
}

func (o NotebookWorkspaceArrayOutput) Index(i pulumi.IntInput) NotebookWorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NotebookWorkspace {
		return vs[0].([]NotebookWorkspace)[vs[1].(int)]
	}).(NotebookWorkspaceOutput)
}

type NotebookWorkspaceMapOutput struct{ *pulumi.OutputState }

func (NotebookWorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]NotebookWorkspace)(nil))
}

func (o NotebookWorkspaceMapOutput) ToNotebookWorkspaceMapOutput() NotebookWorkspaceMapOutput {
	return o
}

func (o NotebookWorkspaceMapOutput) ToNotebookWorkspaceMapOutputWithContext(ctx context.Context) NotebookWorkspaceMapOutput {
	return o
}

func (o NotebookWorkspaceMapOutput) MapIndex(k pulumi.StringInput) NotebookWorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) NotebookWorkspace {
		return vs[0].(map[string]NotebookWorkspace)[vs[1].(string)]
	}).(NotebookWorkspaceOutput)
}

func init() {
	pulumi.RegisterOutputType(NotebookWorkspaceOutput{})
	pulumi.RegisterOutputType(NotebookWorkspacePtrOutput{})
	pulumi.RegisterOutputType(NotebookWorkspaceArrayOutput{})
	pulumi.RegisterOutputType(NotebookWorkspaceMapOutput{})
}
