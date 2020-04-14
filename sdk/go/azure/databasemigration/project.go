// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package databasemigration

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manage a Azure Database Migration Project.
//
// > **NOTE:** Destroying a Database Migration Project will leave any outstanding tasks untouched. This is to avoid unexpectedly deleting any tasks managed outside of this provider.
type Project struct {
	pulumi.CustomResourceState

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
	SourcePlatform pulumi.StringOutput `pulumi:"sourcePlatform"`
	// A mapping of tags to assigned to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
	TargetPlatform pulumi.StringOutput `pulumi:"targetPlatform"`
}

// NewProject registers a new resource with the given unique name, arguments, and options.
func NewProject(ctx *pulumi.Context,
	name string, args *ProjectArgs, opts ...pulumi.ResourceOption) (*Project, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.SourcePlatform == nil {
		return nil, errors.New("missing required argument 'SourcePlatform'")
	}
	if args == nil || args.TargetPlatform == nil {
		return nil, errors.New("missing required argument 'TargetPlatform'")
	}
	if args == nil {
		args = &ProjectArgs{}
	}
	var resource Project
	err := ctx.RegisterResource("azure:databasemigration/project:Project", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProject gets an existing Project resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectState, opts ...pulumi.ResourceOption) (*Project, error) {
	var resource Project
	err := ctx.ReadResource("azure:databasemigration/project:Project", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Project resources.
type projectState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
	ServiceName *string `pulumi:"serviceName"`
	// The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
	SourcePlatform *string `pulumi:"sourcePlatform"`
	// A mapping of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
	TargetPlatform *string `pulumi:"targetPlatform"`
}

type ProjectState struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
	ServiceName pulumi.StringPtrInput
	// The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
	SourcePlatform pulumi.StringPtrInput
	// A mapping of tags to assigned to the resource.
	Tags pulumi.StringMapInput
	// The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
	TargetPlatform pulumi.StringPtrInput
}

func (ProjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectState)(nil)).Elem()
}

type projectArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
	ServiceName string `pulumi:"serviceName"`
	// The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
	SourcePlatform string `pulumi:"sourcePlatform"`
	// A mapping of tags to assigned to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
	TargetPlatform string `pulumi:"targetPlatform"`
}

// The set of arguments for constructing a Project resource.
type ProjectArgs struct {
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specify the name of the database migration project. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Name of the resource group in which to create the database migration project. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// Name of the database migration service where resource belongs to. Changing this forces a new resource to be created.
	ServiceName pulumi.StringInput
	// The platform type of the migration source. Currently only support: `SQL`(on-premises SQL Server). Changing this forces a new resource to be created.
	SourcePlatform pulumi.StringInput
	// A mapping of tags to assigned to the resource.
	Tags pulumi.StringMapInput
	// The platform type of the migration target. Currently only support: `SQLDB`(Azure SQL Database). Changing this forces a new resource to be created.
	TargetPlatform pulumi.StringInput
}

func (ProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectArgs)(nil)).Elem()
}
