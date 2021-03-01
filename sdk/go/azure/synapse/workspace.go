// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Synapse Workspace.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/synapse"
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
// 		exampleAccount, err := storage.NewAccount(ctx, "exampleAccount", &storage.AccountArgs{
// 			ResourceGroupName:      exampleResourceGroup.Name,
// 			Location:               exampleResourceGroup.Location,
// 			AccountTier:            pulumi.String("Standard"),
// 			AccountReplicationType: pulumi.String("LRS"),
// 			AccountKind:            pulumi.String("StorageV2"),
// 			IsHnsEnabled:           pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleDataLakeGen2Filesystem, err := storage.NewDataLakeGen2Filesystem(ctx, "exampleDataLakeGen2Filesystem", &storage.DataLakeGen2FilesystemArgs{
// 			StorageAccountId: exampleAccount.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = synapse.NewWorkspace(ctx, "exampleWorkspace", &synapse.WorkspaceArgs{
// 			ResourceGroupName:               exampleResourceGroup.Name,
// 			Location:                        exampleResourceGroup.Location,
// 			StorageDataLakeGen2FilesystemId: exampleDataLakeGen2Filesystem.ID(),
// 			SqlAdministratorLogin:           pulumi.String("sqladminuser"),
// 			SqlAdministratorLoginPassword:   pulumi.String("H@Sh1CoR3!"),
// 			AadAdmin: &synapse.WorkspaceAadAdminArgs{
// 				Login:    pulumi.String("AzureAD Admin"),
// 				ObjectId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 				TenantId: pulumi.String("00000000-0000-0000-0000-000000000000"),
// 			},
// 			Tags: pulumi.StringMap{
// 				"Env": pulumi.String("production"),
// 			},
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
// Synapse Workspace can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:synapse/workspace:Workspace example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Synapse/workspaces/workspace1
// ```
type Workspace struct {
	pulumi.CustomResourceState

	// An `aadAdmin` block as defined below.
	AadAdmin WorkspaceAadAdminOutput `pulumi:"aadAdmin"`
	// An `azureDevopsRepo` block as defined below.
	AzureDevopsRepo WorkspaceAzureDevopsRepoPtrOutput `pulumi:"azureDevopsRepo"`
	// A list of Connectivity endpoints for this Synapse Workspace.
	ConnectivityEndpoints pulumi.StringMapOutput `pulumi:"connectivityEndpoints"`
	// A `githubRepo` block as defined below.
	GithubRepo WorkspaceGithubRepoPtrOutput `pulumi:"githubRepo"`
	// An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
	Identities WorkspaceIdentityArrayOutput `pulumi:"identities"`
	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Workspace managed resource group.
	ManagedResourceGroupName pulumi.StringOutput `pulumi:"managedResourceGroupName"`
	// Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
	ManagedVirtualNetworkEnabled pulumi.BoolPtrOutput `pulumi:"managedVirtualNetworkEnabled"`
	// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
	SqlAdministratorLogin pulumi.StringOutput `pulumi:"sqlAdministratorLogin"`
	// The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
	SqlAdministratorLoginPassword pulumi.StringOutput `pulumi:"sqlAdministratorLoginPassword"`
	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	SqlIdentityControlEnabled pulumi.BoolPtrOutput `pulumi:"sqlIdentityControlEnabled"`
	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	StorageDataLakeGen2FilesystemId pulumi.StringOutput `pulumi:"storageDataLakeGen2FilesystemId"`
	// A mapping of tags which should be assigned to the Synapse Workspace.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SqlAdministratorLogin == nil {
		return nil, errors.New("invalid value for required argument 'SqlAdministratorLogin'")
	}
	if args.SqlAdministratorLoginPassword == nil {
		return nil, errors.New("invalid value for required argument 'SqlAdministratorLoginPassword'")
	}
	if args.StorageDataLakeGen2FilesystemId == nil {
		return nil, errors.New("invalid value for required argument 'StorageDataLakeGen2FilesystemId'")
	}
	var resource Workspace
	err := ctx.RegisterResource("azure:synapse/workspace:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure:synapse/workspace:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// An `aadAdmin` block as defined below.
	AadAdmin *WorkspaceAadAdmin `pulumi:"aadAdmin"`
	// An `azureDevopsRepo` block as defined below.
	AzureDevopsRepo *WorkspaceAzureDevopsRepo `pulumi:"azureDevopsRepo"`
	// A list of Connectivity endpoints for this Synapse Workspace.
	ConnectivityEndpoints map[string]string `pulumi:"connectivityEndpoints"`
	// A `githubRepo` block as defined below.
	GithubRepo *WorkspaceGithubRepo `pulumi:"githubRepo"`
	// An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
	Identities []WorkspaceIdentity `pulumi:"identities"`
	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Workspace managed resource group.
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
	ManagedVirtualNetworkEnabled *bool `pulumi:"managedVirtualNetworkEnabled"`
	// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
	SqlAdministratorLogin *string `pulumi:"sqlAdministratorLogin"`
	// The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
	SqlAdministratorLoginPassword *string `pulumi:"sqlAdministratorLoginPassword"`
	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	SqlIdentityControlEnabled *bool `pulumi:"sqlIdentityControlEnabled"`
	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	StorageDataLakeGen2FilesystemId *string `pulumi:"storageDataLakeGen2FilesystemId"`
	// A mapping of tags which should be assigned to the Synapse Workspace.
	Tags map[string]string `pulumi:"tags"`
}

type WorkspaceState struct {
	// An `aadAdmin` block as defined below.
	AadAdmin WorkspaceAadAdminPtrInput
	// An `azureDevopsRepo` block as defined below.
	AzureDevopsRepo WorkspaceAzureDevopsRepoPtrInput
	// A list of Connectivity endpoints for this Synapse Workspace.
	ConnectivityEndpoints pulumi.StringMapInput
	// A `githubRepo` block as defined below.
	GithubRepo WorkspaceGithubRepoPtrInput
	// An `identity` block as defined below, which contains the Managed Service Identity information for this Synapse Workspace.
	Identities WorkspaceIdentityArrayInput
	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Workspace managed resource group.
	ManagedResourceGroupName pulumi.StringPtrInput
	// Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
	ManagedVirtualNetworkEnabled pulumi.BoolPtrInput
	// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
	SqlAdministratorLogin pulumi.StringPtrInput
	// The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
	SqlAdministratorLoginPassword pulumi.StringPtrInput
	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	SqlIdentityControlEnabled pulumi.BoolPtrInput
	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	StorageDataLakeGen2FilesystemId pulumi.StringPtrInput
	// A mapping of tags which should be assigned to the Synapse Workspace.
	Tags pulumi.StringMapInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// An `aadAdmin` block as defined below.
	AadAdmin *WorkspaceAadAdmin `pulumi:"aadAdmin"`
	// An `azureDevopsRepo` block as defined below.
	AzureDevopsRepo *WorkspaceAzureDevopsRepo `pulumi:"azureDevopsRepo"`
	// A `githubRepo` block as defined below.
	GithubRepo *WorkspaceGithubRepo `pulumi:"githubRepo"`
	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Workspace managed resource group.
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
	ManagedVirtualNetworkEnabled *bool `pulumi:"managedVirtualNetworkEnabled"`
	// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
	SqlAdministratorLogin string `pulumi:"sqlAdministratorLogin"`
	// The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
	SqlAdministratorLoginPassword string `pulumi:"sqlAdministratorLoginPassword"`
	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	SqlIdentityControlEnabled *bool `pulumi:"sqlIdentityControlEnabled"`
	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	StorageDataLakeGen2FilesystemId string `pulumi:"storageDataLakeGen2FilesystemId"`
	// A mapping of tags which should be assigned to the Synapse Workspace.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// An `aadAdmin` block as defined below.
	AadAdmin WorkspaceAadAdminPtrInput
	// An `azureDevopsRepo` block as defined below.
	AzureDevopsRepo WorkspaceAzureDevopsRepoPtrInput
	// A `githubRepo` block as defined below.
	GithubRepo WorkspaceGithubRepoPtrInput
	// Specifies the Azure Region where the synapse Workspace should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Workspace managed resource group.
	ManagedResourceGroupName pulumi.StringPtrInput
	// Is Virtual Network enabled for all computes in this workspace? Defaults to `false`. Changing this forces a new resource to be created.
	ManagedVirtualNetworkEnabled pulumi.BoolPtrInput
	// Specifies the name which should be used for this synapse Workspace. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group where the synapse Workspace should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Specifies The Login Name of the SQL administrator. Changing this forces a new resource to be created.
	SqlAdministratorLogin pulumi.StringInput
	// The Password associated with the `sqlAdministratorLogin` for the SQL administrator.
	SqlAdministratorLoginPassword pulumi.StringInput
	// Are pipelines (running as workspace's system assigned identity) allowed to access SQL pools?
	SqlIdentityControlEnabled pulumi.BoolPtrInput
	// Specifies the ID of storage data lake gen2 filesystem resource. Changing this forces a new resource to be created.
	StorageDataLakeGen2FilesystemId pulumi.StringInput
	// A mapping of tags which should be assigned to the Synapse Workspace.
	Tags pulumi.StringMapInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

func (i *Workspace) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return i.ToWorkspacePtrOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePtrOutput)
}

type WorkspacePtrInput interface {
	pulumi.Input

	ToWorkspacePtrOutput() WorkspacePtrOutput
	ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput
}

type workspacePtrType WorkspaceArgs

func (*workspacePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil))
}

func (i *workspacePtrType) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return i.ToWorkspacePtrOutputWithContext(context.Background())
}

func (i *workspacePtrType) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspacePtrOutput)
}

// WorkspaceArrayInput is an input type that accepts WorkspaceArray and WorkspaceArrayOutput values.
// You can construct a concrete instance of `WorkspaceArrayInput` via:
//
//          WorkspaceArray{ WorkspaceArgs{...} }
type WorkspaceArrayInput interface {
	pulumi.Input

	ToWorkspaceArrayOutput() WorkspaceArrayOutput
	ToWorkspaceArrayOutputWithContext(context.Context) WorkspaceArrayOutput
}

type WorkspaceArray []WorkspaceInput

func (WorkspaceArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Workspace)(nil))
}

func (i WorkspaceArray) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return i.ToWorkspaceArrayOutputWithContext(context.Background())
}

func (i WorkspaceArray) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceArrayOutput)
}

// WorkspaceMapInput is an input type that accepts WorkspaceMap and WorkspaceMapOutput values.
// You can construct a concrete instance of `WorkspaceMapInput` via:
//
//          WorkspaceMap{ "key": WorkspaceArgs{...} }
type WorkspaceMapInput interface {
	pulumi.Input

	ToWorkspaceMapOutput() WorkspaceMapOutput
	ToWorkspaceMapOutputWithContext(context.Context) WorkspaceMapOutput
}

type WorkspaceMap map[string]WorkspaceInput

func (WorkspaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Workspace)(nil))
}

func (i WorkspaceMap) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return i.ToWorkspaceMapOutputWithContext(context.Background())
}

func (i WorkspaceMap) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceMapOutput)
}

type WorkspaceOutput struct {
	*pulumi.OutputState
}

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil))
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return o.ToWorkspacePtrOutputWithContext(context.Background())
}

func (o WorkspaceOutput) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return o.ApplyT(func(v Workspace) *Workspace {
		return &v
	}).(WorkspacePtrOutput)
}

type WorkspacePtrOutput struct {
	*pulumi.OutputState
}

func (WorkspacePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil))
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutput() WorkspacePtrOutput {
	return o
}

func (o WorkspacePtrOutput) ToWorkspacePtrOutputWithContext(ctx context.Context) WorkspacePtrOutput {
	return o
}

type WorkspaceArrayOutput struct{ *pulumi.OutputState }

func (WorkspaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Workspace)(nil))
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutput() WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) ToWorkspaceArrayOutputWithContext(ctx context.Context) WorkspaceArrayOutput {
	return o
}

func (o WorkspaceArrayOutput) Index(i pulumi.IntInput) WorkspaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Workspace {
		return vs[0].([]Workspace)[vs[1].(int)]
	}).(WorkspaceOutput)
}

type WorkspaceMapOutput struct{ *pulumi.OutputState }

func (WorkspaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Workspace)(nil))
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutput() WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) ToWorkspaceMapOutputWithContext(ctx context.Context) WorkspaceMapOutput {
	return o
}

func (o WorkspaceMapOutput) MapIndex(k pulumi.StringInput) WorkspaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Workspace {
		return vs[0].(map[string]Workspace)[vs[1].(string)]
	}).(WorkspaceOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
	pulumi.RegisterOutputType(WorkspacePtrOutput{})
	pulumi.RegisterOutputType(WorkspaceArrayOutput{})
	pulumi.RegisterOutputType(WorkspaceMapOutput{})
}
