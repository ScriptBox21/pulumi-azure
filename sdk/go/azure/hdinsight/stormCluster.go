// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package hdinsight

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a HDInsight Storm Cluster.
//
// !> **Note:** [HDInsight 3.6 is deprecated and will be retired on 2020-12-31 - HDInsight 4.0 no longer supports Storm Clusters](https://docs.microsoft.com/en-us/azure/hdinsight/hdinsight-component-versioning#available-versions) - as such this resource is deprecated and will be removed in the next major version of the Provider.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/hdinsight"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/storage"
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
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		exampleContainer, err := storage.NewContainer(ctx, "exampleContainer", &storage.ContainerArgs{
// 			StorageAccountName:  exampleAccount.Name,
// 			ContainerAccessType: pulumi.String("private"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = hdinsight.NewStormCluster(ctx, "exampleStormCluster", &hdinsight.StormClusterArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			Location:          exampleResourceGroup.Location,
// 			ClusterVersion:    pulumi.String("3.6"),
// 			Tier:              pulumi.String("Standard"),
// 			ComponentVersion: &hdinsight.StormClusterComponentVersionArgs{
// 				Storm: pulumi.String("1.1"),
// 			},
// 			Gateway: &hdinsight.StormClusterGatewayArgs{
// 				Enabled:  pulumi.Bool(true),
// 				Username: pulumi.String("acctestusrgw"),
// 				Password: pulumi.String("Password123!"),
// 			},
// 			StorageAccounts: hdinsight.StormClusterStorageAccountArray{
// 				&hdinsight.StormClusterStorageAccountArgs{
// 					StorageContainerId: exampleContainer.ID(),
// 					StorageAccountKey:  exampleAccount.PrimaryAccessKey,
// 					IsDefault:          pulumi.Bool(true),
// 				},
// 			},
// 			Roles: &hdinsight.StormClusterRolesArgs{
// 				HeadNode: &hdinsight.StormClusterRolesHeadNodeArgs{
// 					VmSize:   pulumi.String("Standard_A3"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
// 				WorkerNode: &hdinsight.StormClusterRolesWorkerNodeArgs{
// 					VmSize:              pulumi.String("Standard_D3_V2"),
// 					Username:            pulumi.String("acctestusrvm"),
// 					Password:            pulumi.String("AccTestvdSC4daf986!"),
// 					TargetInstanceCount: pulumi.Int(3),
// 				},
// 				ZookeeperNode: &hdinsight.StormClusterRolesZookeeperNodeArgs{
// 					VmSize:   pulumi.String("Standard_A4_V2"),
// 					Username: pulumi.String("acctestusrvm"),
// 					Password: pulumi.String("AccTestvdSC4daf986!"),
// 				},
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
// HDInsight Storm Clusters can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:hdinsight/stormCluster:StormCluster example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.HDInsight/clusters/cluster1
// ```
type StormCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion StormClusterComponentVersionOutput `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway StormClusterGatewayOutput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Storm Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores StormClusterMetastoresPtrOutput `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor StormClusterMonitorPtrOutput `pulumi:"monitor"`
	// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles StormClusterRolesOutput `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight Storm Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts StormClusterStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Storm Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringOutput    `pulumi:"tier"`
	TlsMinVersion pulumi.StringPtrOutput `pulumi:"tlsMinVersion"`
}

// NewStormCluster registers a new resource with the given unique name, arguments, and options.
func NewStormCluster(ctx *pulumi.Context,
	name string, args *StormClusterArgs, opts ...pulumi.ResourceOption) (*StormCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterVersion == nil {
		return nil, errors.New("invalid value for required argument 'ClusterVersion'")
	}
	if args.ComponentVersion == nil {
		return nil, errors.New("invalid value for required argument 'ComponentVersion'")
	}
	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Roles == nil {
		return nil, errors.New("invalid value for required argument 'Roles'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	var resource StormCluster
	err := ctx.RegisterResource("azure:hdinsight/stormCluster:StormCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStormCluster gets an existing StormCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStormCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StormClusterState, opts ...pulumi.ResourceOption) (*StormCluster, error) {
	var resource StormCluster
	err := ctx.ReadResource("azure:hdinsight/stormCluster:StormCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StormCluster resources.
type stormClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion *StormClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway *StormClusterGateway `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight Storm Cluster.
	HttpsEndpoint *string `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *StormClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *StormClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles *StormClusterRoles `pulumi:"roles"`
	// The SSH Connectivity Endpoint for this HDInsight Storm Cluster.
	SshEndpoint *string `pulumi:"sshEndpoint"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []StormClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Storm Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          *string `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

type StormClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringPtrInput
	// A `componentVersion` block as defined below.
	ComponentVersion StormClusterComponentVersionPtrInput
	// A `gateway` block as defined below.
	Gateway StormClusterGatewayPtrInput
	// The HTTPS Connectivity Endpoint for this HDInsight Storm Cluster.
	HttpsEndpoint pulumi.StringPtrInput
	// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores StormClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor StormClusterMonitorPtrInput
	// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roles` block as defined below.
	Roles StormClusterRolesPtrInput
	// The SSH Connectivity Endpoint for this HDInsight Storm Cluster.
	SshEndpoint pulumi.StringPtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts StormClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Storm Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringPtrInput
	TlsMinVersion pulumi.StringPtrInput
}

func (StormClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*stormClusterState)(nil)).Elem()
}

type stormClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A `componentVersion` block as defined below.
	ComponentVersion StormClusterComponentVersion `pulumi:"componentVersion"`
	// A `gateway` block as defined below.
	Gateway StormClusterGateway `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `metastores` block as defined below.
	Metastores *StormClusterMetastores `pulumi:"metastores"`
	// A `monitor` block as defined below.
	Monitor *StormClusterMonitor `pulumi:"monitor"`
	// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles StormClusterRoles `pulumi:"roles"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []StormClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight Storm Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          string  `pulumi:"tier"`
	TlsMinVersion *string `pulumi:"tlsMinVersion"`
}

// The set of arguments for constructing a StormCluster resource.
type StormClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput
	// A `componentVersion` block as defined below.
	ComponentVersion StormClusterComponentVersionInput
	// A `gateway` block as defined below.
	Gateway StormClusterGatewayInput
	// Specifies the Azure Region which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `metastores` block as defined below.
	Metastores StormClusterMetastoresPtrInput
	// A `monitor` block as defined below.
	Monitor StormClusterMonitorPtrInput
	// Specifies the name for this HDInsight Storm Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight Storm Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roles` block as defined below.
	Roles StormClusterRolesInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts StormClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight Storm Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight Storm Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier          pulumi.StringInput
	TlsMinVersion pulumi.StringPtrInput
}

func (StormClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stormClusterArgs)(nil)).Elem()
}

type StormClusterInput interface {
	pulumi.Input

	ToStormClusterOutput() StormClusterOutput
	ToStormClusterOutputWithContext(ctx context.Context) StormClusterOutput
}

func (*StormCluster) ElementType() reflect.Type {
	return reflect.TypeOf((*StormCluster)(nil))
}

func (i *StormCluster) ToStormClusterOutput() StormClusterOutput {
	return i.ToStormClusterOutputWithContext(context.Background())
}

func (i *StormCluster) ToStormClusterOutputWithContext(ctx context.Context) StormClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StormClusterOutput)
}

func (i *StormCluster) ToStormClusterPtrOutput() StormClusterPtrOutput {
	return i.ToStormClusterPtrOutputWithContext(context.Background())
}

func (i *StormCluster) ToStormClusterPtrOutputWithContext(ctx context.Context) StormClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StormClusterPtrOutput)
}

type StormClusterPtrInput interface {
	pulumi.Input

	ToStormClusterPtrOutput() StormClusterPtrOutput
	ToStormClusterPtrOutputWithContext(ctx context.Context) StormClusterPtrOutput
}

type stormClusterPtrType StormClusterArgs

func (*stormClusterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StormCluster)(nil))
}

func (i *stormClusterPtrType) ToStormClusterPtrOutput() StormClusterPtrOutput {
	return i.ToStormClusterPtrOutputWithContext(context.Background())
}

func (i *stormClusterPtrType) ToStormClusterPtrOutputWithContext(ctx context.Context) StormClusterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StormClusterPtrOutput)
}

// StormClusterArrayInput is an input type that accepts StormClusterArray and StormClusterArrayOutput values.
// You can construct a concrete instance of `StormClusterArrayInput` via:
//
//          StormClusterArray{ StormClusterArgs{...} }
type StormClusterArrayInput interface {
	pulumi.Input

	ToStormClusterArrayOutput() StormClusterArrayOutput
	ToStormClusterArrayOutputWithContext(context.Context) StormClusterArrayOutput
}

type StormClusterArray []StormClusterInput

func (StormClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*StormCluster)(nil))
}

func (i StormClusterArray) ToStormClusterArrayOutput() StormClusterArrayOutput {
	return i.ToStormClusterArrayOutputWithContext(context.Background())
}

func (i StormClusterArray) ToStormClusterArrayOutputWithContext(ctx context.Context) StormClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StormClusterArrayOutput)
}

// StormClusterMapInput is an input type that accepts StormClusterMap and StormClusterMapOutput values.
// You can construct a concrete instance of `StormClusterMapInput` via:
//
//          StormClusterMap{ "key": StormClusterArgs{...} }
type StormClusterMapInput interface {
	pulumi.Input

	ToStormClusterMapOutput() StormClusterMapOutput
	ToStormClusterMapOutputWithContext(context.Context) StormClusterMapOutput
}

type StormClusterMap map[string]StormClusterInput

func (StormClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*StormCluster)(nil))
}

func (i StormClusterMap) ToStormClusterMapOutput() StormClusterMapOutput {
	return i.ToStormClusterMapOutputWithContext(context.Background())
}

func (i StormClusterMap) ToStormClusterMapOutputWithContext(ctx context.Context) StormClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StormClusterMapOutput)
}

type StormClusterOutput struct {
	*pulumi.OutputState
}

func (StormClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StormCluster)(nil))
}

func (o StormClusterOutput) ToStormClusterOutput() StormClusterOutput {
	return o
}

func (o StormClusterOutput) ToStormClusterOutputWithContext(ctx context.Context) StormClusterOutput {
	return o
}

func (o StormClusterOutput) ToStormClusterPtrOutput() StormClusterPtrOutput {
	return o.ToStormClusterPtrOutputWithContext(context.Background())
}

func (o StormClusterOutput) ToStormClusterPtrOutputWithContext(ctx context.Context) StormClusterPtrOutput {
	return o.ApplyT(func(v StormCluster) *StormCluster {
		return &v
	}).(StormClusterPtrOutput)
}

type StormClusterPtrOutput struct {
	*pulumi.OutputState
}

func (StormClusterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StormCluster)(nil))
}

func (o StormClusterPtrOutput) ToStormClusterPtrOutput() StormClusterPtrOutput {
	return o
}

func (o StormClusterPtrOutput) ToStormClusterPtrOutputWithContext(ctx context.Context) StormClusterPtrOutput {
	return o
}

type StormClusterArrayOutput struct{ *pulumi.OutputState }

func (StormClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StormCluster)(nil))
}

func (o StormClusterArrayOutput) ToStormClusterArrayOutput() StormClusterArrayOutput {
	return o
}

func (o StormClusterArrayOutput) ToStormClusterArrayOutputWithContext(ctx context.Context) StormClusterArrayOutput {
	return o
}

func (o StormClusterArrayOutput) Index(i pulumi.IntInput) StormClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StormCluster {
		return vs[0].([]StormCluster)[vs[1].(int)]
	}).(StormClusterOutput)
}

type StormClusterMapOutput struct{ *pulumi.OutputState }

func (StormClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]StormCluster)(nil))
}

func (o StormClusterMapOutput) ToStormClusterMapOutput() StormClusterMapOutput {
	return o
}

func (o StormClusterMapOutput) ToStormClusterMapOutputWithContext(ctx context.Context) StormClusterMapOutput {
	return o
}

func (o StormClusterMapOutput) MapIndex(k pulumi.StringInput) StormClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) StormCluster {
		return vs[0].(map[string]StormCluster)[vs[1].(string)]
	}).(StormClusterOutput)
}

func init() {
	pulumi.RegisterOutputType(StormClusterOutput{})
	pulumi.RegisterOutputType(StormClusterPtrOutput{})
	pulumi.RegisterOutputType(StormClusterArrayOutput{})
	pulumi.RegisterOutputType(StormClusterMapOutput{})
}
