// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kusto

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Attached Database Configuration
//
// ## Import
//
// Kusto Attached Database Configurations can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/AttachedDatabaseConfigurations/configuration1
// ```
type AttachedDatabaseConfiguration struct {
	pulumi.CustomResourceState

	// The list of databases from the `clusterResourceId` which are currently attached to the cluster.
	AttachedDatabaseNames pulumi.StringArrayOutput `pulumi:"attachedDatabaseNames"`
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringOutput `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringOutput `pulumi:"databaseName"`
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind pulumi.StringPtrOutput `pulumi:"defaultPrincipalModificationKind"`
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
}

// NewAttachedDatabaseConfiguration registers a new resource with the given unique name, arguments, and options.
func NewAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, args *AttachedDatabaseConfigurationArgs, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.ClusterResourceId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterResourceId'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	var resource AttachedDatabaseConfiguration
	err := ctx.RegisterResource("azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedDatabaseConfiguration gets an existing AttachedDatabaseConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedDatabaseConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedDatabaseConfigurationState, opts ...pulumi.ResourceOption) (*AttachedDatabaseConfiguration, error) {
	var resource AttachedDatabaseConfiguration
	err := ctx.ReadResource("azure:kusto/attachedDatabaseConfiguration:AttachedDatabaseConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedDatabaseConfiguration resources.
type attachedDatabaseConfigurationState struct {
	// The list of databases from the `clusterResourceId` which are currently attached to the cluster.
	AttachedDatabaseNames []string `pulumi:"attachedDatabaseNames"`
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName *string `pulumi:"clusterName"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId *string `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName *string `pulumi:"databaseName"`
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind *string `pulumi:"defaultPrincipalModificationKind"`
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
}

type AttachedDatabaseConfigurationState struct {
	// The list of databases from the `clusterResourceId` which are currently attached to the cluster.
	AttachedDatabaseNames pulumi.StringArrayInput
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName pulumi.StringPtrInput
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringPtrInput
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringPtrInput
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind pulumi.StringPtrInput
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
}

func (AttachedDatabaseConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationState)(nil)).Elem()
}

type attachedDatabaseConfigurationArgs struct {
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName string `pulumi:"clusterName"`
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId string `pulumi:"clusterResourceId"`
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName string `pulumi:"databaseName"`
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind *string `pulumi:"defaultPrincipalModificationKind"`
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a AttachedDatabaseConfiguration resource.
type AttachedDatabaseConfigurationArgs struct {
	// Specifies the name of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ClusterName pulumi.StringInput
	// The resource id of the cluster where the databases you would like to attach reside.
	ClusterResourceId pulumi.StringInput
	// The name of the database which you would like to attach, use * if you want to follow all current and future databases.
	DatabaseName pulumi.StringInput
	// The default principals modification kind. Valid values are: `None` (default), `Replace` and `Union`.
	DefaultPrincipalModificationKind pulumi.StringPtrInput
	// Specifies the location of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Attached Database Configuration to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the resource group of the Kusto Cluster for which the configuration will be created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
}

func (AttachedDatabaseConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedDatabaseConfigurationArgs)(nil)).Elem()
}

type AttachedDatabaseConfigurationInput interface {
	pulumi.Input

	ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput
	ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput
}

func (*AttachedDatabaseConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachedDatabaseConfiguration)(nil))
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return i.ToAttachedDatabaseConfigurationOutputWithContext(context.Background())
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationOutput)
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationPtrOutput() AttachedDatabaseConfigurationPtrOutput {
	return i.ToAttachedDatabaseConfigurationPtrOutputWithContext(context.Background())
}

func (i *AttachedDatabaseConfiguration) ToAttachedDatabaseConfigurationPtrOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationPtrOutput)
}

type AttachedDatabaseConfigurationPtrInput interface {
	pulumi.Input

	ToAttachedDatabaseConfigurationPtrOutput() AttachedDatabaseConfigurationPtrOutput
	ToAttachedDatabaseConfigurationPtrOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationPtrOutput
}

type attachedDatabaseConfigurationPtrType AttachedDatabaseConfigurationArgs

func (*attachedDatabaseConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDatabaseConfiguration)(nil))
}

func (i *attachedDatabaseConfigurationPtrType) ToAttachedDatabaseConfigurationPtrOutput() AttachedDatabaseConfigurationPtrOutput {
	return i.ToAttachedDatabaseConfigurationPtrOutputWithContext(context.Background())
}

func (i *attachedDatabaseConfigurationPtrType) ToAttachedDatabaseConfigurationPtrOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationPtrOutput)
}

// AttachedDatabaseConfigurationArrayInput is an input type that accepts AttachedDatabaseConfigurationArray and AttachedDatabaseConfigurationArrayOutput values.
// You can construct a concrete instance of `AttachedDatabaseConfigurationArrayInput` via:
//
//          AttachedDatabaseConfigurationArray{ AttachedDatabaseConfigurationArgs{...} }
type AttachedDatabaseConfigurationArrayInput interface {
	pulumi.Input

	ToAttachedDatabaseConfigurationArrayOutput() AttachedDatabaseConfigurationArrayOutput
	ToAttachedDatabaseConfigurationArrayOutputWithContext(context.Context) AttachedDatabaseConfigurationArrayOutput
}

type AttachedDatabaseConfigurationArray []AttachedDatabaseConfigurationInput

func (AttachedDatabaseConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AttachedDatabaseConfiguration)(nil))
}

func (i AttachedDatabaseConfigurationArray) ToAttachedDatabaseConfigurationArrayOutput() AttachedDatabaseConfigurationArrayOutput {
	return i.ToAttachedDatabaseConfigurationArrayOutputWithContext(context.Background())
}

func (i AttachedDatabaseConfigurationArray) ToAttachedDatabaseConfigurationArrayOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationArrayOutput)
}

// AttachedDatabaseConfigurationMapInput is an input type that accepts AttachedDatabaseConfigurationMap and AttachedDatabaseConfigurationMapOutput values.
// You can construct a concrete instance of `AttachedDatabaseConfigurationMapInput` via:
//
//          AttachedDatabaseConfigurationMap{ "key": AttachedDatabaseConfigurationArgs{...} }
type AttachedDatabaseConfigurationMapInput interface {
	pulumi.Input

	ToAttachedDatabaseConfigurationMapOutput() AttachedDatabaseConfigurationMapOutput
	ToAttachedDatabaseConfigurationMapOutputWithContext(context.Context) AttachedDatabaseConfigurationMapOutput
}

type AttachedDatabaseConfigurationMap map[string]AttachedDatabaseConfigurationInput

func (AttachedDatabaseConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AttachedDatabaseConfiguration)(nil))
}

func (i AttachedDatabaseConfigurationMap) ToAttachedDatabaseConfigurationMapOutput() AttachedDatabaseConfigurationMapOutput {
	return i.ToAttachedDatabaseConfigurationMapOutputWithContext(context.Background())
}

func (i AttachedDatabaseConfigurationMap) ToAttachedDatabaseConfigurationMapOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedDatabaseConfigurationMapOutput)
}

type AttachedDatabaseConfigurationOutput struct {
	*pulumi.OutputState
}

func (AttachedDatabaseConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AttachedDatabaseConfiguration)(nil))
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutput() AttachedDatabaseConfigurationOutput {
	return o
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationOutput {
	return o
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationPtrOutput() AttachedDatabaseConfigurationPtrOutput {
	return o.ToAttachedDatabaseConfigurationPtrOutputWithContext(context.Background())
}

func (o AttachedDatabaseConfigurationOutput) ToAttachedDatabaseConfigurationPtrOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationPtrOutput {
	return o.ApplyT(func(v AttachedDatabaseConfiguration) *AttachedDatabaseConfiguration {
		return &v
	}).(AttachedDatabaseConfigurationPtrOutput)
}

type AttachedDatabaseConfigurationPtrOutput struct {
	*pulumi.OutputState
}

func (AttachedDatabaseConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedDatabaseConfiguration)(nil))
}

func (o AttachedDatabaseConfigurationPtrOutput) ToAttachedDatabaseConfigurationPtrOutput() AttachedDatabaseConfigurationPtrOutput {
	return o
}

func (o AttachedDatabaseConfigurationPtrOutput) ToAttachedDatabaseConfigurationPtrOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationPtrOutput {
	return o
}

type AttachedDatabaseConfigurationArrayOutput struct{ *pulumi.OutputState }

func (AttachedDatabaseConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AttachedDatabaseConfiguration)(nil))
}

func (o AttachedDatabaseConfigurationArrayOutput) ToAttachedDatabaseConfigurationArrayOutput() AttachedDatabaseConfigurationArrayOutput {
	return o
}

func (o AttachedDatabaseConfigurationArrayOutput) ToAttachedDatabaseConfigurationArrayOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationArrayOutput {
	return o
}

func (o AttachedDatabaseConfigurationArrayOutput) Index(i pulumi.IntInput) AttachedDatabaseConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AttachedDatabaseConfiguration {
		return vs[0].([]AttachedDatabaseConfiguration)[vs[1].(int)]
	}).(AttachedDatabaseConfigurationOutput)
}

type AttachedDatabaseConfigurationMapOutput struct{ *pulumi.OutputState }

func (AttachedDatabaseConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AttachedDatabaseConfiguration)(nil))
}

func (o AttachedDatabaseConfigurationMapOutput) ToAttachedDatabaseConfigurationMapOutput() AttachedDatabaseConfigurationMapOutput {
	return o
}

func (o AttachedDatabaseConfigurationMapOutput) ToAttachedDatabaseConfigurationMapOutputWithContext(ctx context.Context) AttachedDatabaseConfigurationMapOutput {
	return o
}

func (o AttachedDatabaseConfigurationMapOutput) MapIndex(k pulumi.StringInput) AttachedDatabaseConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AttachedDatabaseConfiguration {
		return vs[0].(map[string]AttachedDatabaseConfiguration)[vs[1].(string)]
	}).(AttachedDatabaseConfigurationOutput)
}

func init() {
	pulumi.RegisterOutputType(AttachedDatabaseConfigurationOutput{})
	pulumi.RegisterOutputType(AttachedDatabaseConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AttachedDatabaseConfigurationArrayOutput{})
	pulumi.RegisterOutputType(AttachedDatabaseConfigurationMapOutput{})
}
