// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Create a failover group of databases on a collection of Azure SQL servers.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/sql"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleResourceGroup, err := core.NewResourceGroup(ctx, "exampleResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("uksouth"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primary, err := sql.NewSqlServer(ctx, "primary", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   exampleResourceGroup.Location,
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("sqladmin"),
// 			AdministratorLoginPassword: pulumi.String(fmt.Sprintf("%v%v%v%v", "pa", "$", "$", "w0rd")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		secondary, err := sql.NewSqlServer(ctx, "secondary", &sql.SqlServerArgs{
// 			ResourceGroupName:          exampleResourceGroup.Name,
// 			Location:                   pulumi.String("northeurope"),
// 			Version:                    pulumi.String("12.0"),
// 			AdministratorLogin:         pulumi.String("sqladmin"),
// 			AdministratorLoginPassword: pulumi.String(fmt.Sprintf("%v%v%v%v", "pa", "$", "$", "w0rd")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		db1, err := sql.NewDatabase(ctx, "db1", &sql.DatabaseArgs{
// 			ResourceGroupName: primary.ResourceGroupName,
// 			Location:          primary.Location,
// 			ServerName:        primary.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = sql.NewFailoverGroup(ctx, "exampleFailoverGroup", &sql.FailoverGroupArgs{
// 			ResourceGroupName: primary.ResourceGroupName,
// 			ServerName:        primary.Name,
// 			Databases: pulumi.StringArray{
// 				db1.ID(),
// 			},
// 			PartnerServers: sql.FailoverGroupPartnerServerArray{
// 				&sql.FailoverGroupPartnerServerArgs{
// 					Id: secondary.ID(),
// 				},
// 			},
// 			ReadWriteEndpointFailoverPolicy: &sql.FailoverGroupReadWriteEndpointFailoverPolicyArgs{
// 				Mode:         pulumi.String("Automatic"),
// 				GraceMinutes: pulumi.Int(60),
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
// SQL Failover Groups can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:sql/failoverGroup:FailoverGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myresourcegroup/providers/Microsoft.Sql/servers/myserver/failovergroups/group1
// ```
type FailoverGroup struct {
	pulumi.CustomResourceState

	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayOutput `pulumi:"databases"`
	// the location of the failover group.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayOutput `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyOutput `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyOutput `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// local replication role of the failover group instance.
	Role pulumi.StringOutput `pulumi:"role"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringOutput `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewFailoverGroup registers a new resource with the given unique name, arguments, and options.
func NewFailoverGroup(ctx *pulumi.Context,
	name string, args *FailoverGroupArgs, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PartnerServers == nil {
		return nil, errors.New("invalid value for required argument 'PartnerServers'")
	}
	if args.ReadWriteEndpointFailoverPolicy == nil {
		return nil, errors.New("invalid value for required argument 'ReadWriteEndpointFailoverPolicy'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	var resource FailoverGroup
	err := ctx.RegisterResource("azure:sql/failoverGroup:FailoverGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFailoverGroup gets an existing FailoverGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFailoverGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FailoverGroupState, opts ...pulumi.ResourceOption) (*FailoverGroup, error) {
	var resource FailoverGroup
	err := ctx.ReadResource("azure:sql/failoverGroup:FailoverGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FailoverGroup resources.
type failoverGroupState struct {
	// A list of database ids to add to the failover group
	Databases []string `pulumi:"databases"`
	// the location of the failover group.
	Location *string `pulumi:"location"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy *FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy *FailoverGroupReadonlyEndpointFailoverPolicy `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// local replication role of the failover group instance.
	Role *string `pulumi:"role"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName *string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type FailoverGroupState struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput
	// the location of the failover group.
	Location pulumi.StringPtrInput
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyPtrInput
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyPtrInput
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringPtrInput
	// local replication role of the failover group instance.
	Role pulumi.StringPtrInput
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupState)(nil)).Elem()
}

type failoverGroupArgs struct {
	// A list of database ids to add to the failover group
	Databases []string `pulumi:"databases"`
	// The name of the failover group. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A list of secondary servers as documented below
	PartnerServers []FailoverGroupPartnerServer `pulumi:"partnerServers"`
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicy `pulumi:"readWriteEndpointFailoverPolicy"`
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy *FailoverGroupReadonlyEndpointFailoverPolicy `pulumi:"readonlyEndpointFailoverPolicy"`
	// The name of the resource group containing the SQL server
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName string `pulumi:"serverName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a FailoverGroup resource.
type FailoverGroupArgs struct {
	// A list of database ids to add to the failover group
	Databases pulumi.StringArrayInput
	// The name of the failover group. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A list of secondary servers as documented below
	PartnerServers FailoverGroupPartnerServerArrayInput
	// A read/write policy as documented below
	ReadWriteEndpointFailoverPolicy FailoverGroupReadWriteEndpointFailoverPolicyInput
	// a read-only policy as documented below
	ReadonlyEndpointFailoverPolicy FailoverGroupReadonlyEndpointFailoverPolicyPtrInput
	// The name of the resource group containing the SQL server
	ResourceGroupName pulumi.StringInput
	// The name of the primary SQL server. Changing this forces a new resource to be created.
	ServerName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (FailoverGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*failoverGroupArgs)(nil)).Elem()
}

type FailoverGroupInput interface {
	pulumi.Input

	ToFailoverGroupOutput() FailoverGroupOutput
	ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput
}

func (*FailoverGroup) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroup)(nil))
}

func (i *FailoverGroup) ToFailoverGroupOutput() FailoverGroupOutput {
	return i.ToFailoverGroupOutputWithContext(context.Background())
}

func (i *FailoverGroup) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupOutput)
}

func (i *FailoverGroup) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return i.ToFailoverGroupPtrOutputWithContext(context.Background())
}

func (i *FailoverGroup) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupPtrOutput)
}

type FailoverGroupPtrInput interface {
	pulumi.Input

	ToFailoverGroupPtrOutput() FailoverGroupPtrOutput
	ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput
}

type failoverGroupPtrType FailoverGroupArgs

func (*failoverGroupPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil))
}

func (i *failoverGroupPtrType) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return i.ToFailoverGroupPtrOutputWithContext(context.Background())
}

func (i *failoverGroupPtrType) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupPtrOutput)
}

// FailoverGroupArrayInput is an input type that accepts FailoverGroupArray and FailoverGroupArrayOutput values.
// You can construct a concrete instance of `FailoverGroupArrayInput` via:
//
//          FailoverGroupArray{ FailoverGroupArgs{...} }
type FailoverGroupArrayInput interface {
	pulumi.Input

	ToFailoverGroupArrayOutput() FailoverGroupArrayOutput
	ToFailoverGroupArrayOutputWithContext(context.Context) FailoverGroupArrayOutput
}

type FailoverGroupArray []FailoverGroupInput

func (FailoverGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FailoverGroup)(nil))
}

func (i FailoverGroupArray) ToFailoverGroupArrayOutput() FailoverGroupArrayOutput {
	return i.ToFailoverGroupArrayOutputWithContext(context.Background())
}

func (i FailoverGroupArray) ToFailoverGroupArrayOutputWithContext(ctx context.Context) FailoverGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupArrayOutput)
}

// FailoverGroupMapInput is an input type that accepts FailoverGroupMap and FailoverGroupMapOutput values.
// You can construct a concrete instance of `FailoverGroupMapInput` via:
//
//          FailoverGroupMap{ "key": FailoverGroupArgs{...} }
type FailoverGroupMapInput interface {
	pulumi.Input

	ToFailoverGroupMapOutput() FailoverGroupMapOutput
	ToFailoverGroupMapOutputWithContext(context.Context) FailoverGroupMapOutput
}

type FailoverGroupMap map[string]FailoverGroupInput

func (FailoverGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FailoverGroup)(nil))
}

func (i FailoverGroupMap) ToFailoverGroupMapOutput() FailoverGroupMapOutput {
	return i.ToFailoverGroupMapOutputWithContext(context.Background())
}

func (i FailoverGroupMap) ToFailoverGroupMapOutputWithContext(ctx context.Context) FailoverGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverGroupMapOutput)
}

type FailoverGroupOutput struct {
	*pulumi.OutputState
}

func (FailoverGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverGroup)(nil))
}

func (o FailoverGroupOutput) ToFailoverGroupOutput() FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupOutputWithContext(ctx context.Context) FailoverGroupOutput {
	return o
}

func (o FailoverGroupOutput) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return o.ToFailoverGroupPtrOutputWithContext(context.Background())
}

func (o FailoverGroupOutput) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return o.ApplyT(func(v FailoverGroup) *FailoverGroup {
		return &v
	}).(FailoverGroupPtrOutput)
}

type FailoverGroupPtrOutput struct {
	*pulumi.OutputState
}

func (FailoverGroupPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FailoverGroup)(nil))
}

func (o FailoverGroupPtrOutput) ToFailoverGroupPtrOutput() FailoverGroupPtrOutput {
	return o
}

func (o FailoverGroupPtrOutput) ToFailoverGroupPtrOutputWithContext(ctx context.Context) FailoverGroupPtrOutput {
	return o
}

type FailoverGroupArrayOutput struct{ *pulumi.OutputState }

func (FailoverGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverGroup)(nil))
}

func (o FailoverGroupArrayOutput) ToFailoverGroupArrayOutput() FailoverGroupArrayOutput {
	return o
}

func (o FailoverGroupArrayOutput) ToFailoverGroupArrayOutputWithContext(ctx context.Context) FailoverGroupArrayOutput {
	return o
}

func (o FailoverGroupArrayOutput) Index(i pulumi.IntInput) FailoverGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FailoverGroup {
		return vs[0].([]FailoverGroup)[vs[1].(int)]
	}).(FailoverGroupOutput)
}

type FailoverGroupMapOutput struct{ *pulumi.OutputState }

func (FailoverGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FailoverGroup)(nil))
}

func (o FailoverGroupMapOutput) ToFailoverGroupMapOutput() FailoverGroupMapOutput {
	return o
}

func (o FailoverGroupMapOutput) ToFailoverGroupMapOutputWithContext(ctx context.Context) FailoverGroupMapOutput {
	return o
}

func (o FailoverGroupMapOutput) MapIndex(k pulumi.StringInput) FailoverGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FailoverGroup {
		return vs[0].(map[string]FailoverGroup)[vs[1].(string)]
	}).(FailoverGroupOutput)
}

func init() {
	pulumi.RegisterOutputType(FailoverGroupOutput{})
	pulumi.RegisterOutputType(FailoverGroupPtrOutput{})
	pulumi.RegisterOutputType(FailoverGroupArrayOutput{})
	pulumi.RegisterOutputType(FailoverGroupMapOutput{})
}
