// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package loganalytics

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Links a Log Analytics (formally Operational Insights) Workspace to another resource. The (currently) only linkable service is an Azure Automation Account.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/log_analytics_linked_service.html.markdown.
type LinkedService struct {
	pulumi.CustomResourceState

	// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
	LinkedServiceName pulumi.StringPtrOutput `pulumi:"linkedServiceName"`
	// The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
	ResourceId pulumi.StringOutput `pulumi:"resourceId"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
	WorkspaceName pulumi.StringOutput `pulumi:"workspaceName"`
}

// NewLinkedService registers a new resource with the given unique name, arguments, and options.
func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceId == nil {
		return nil, errors.New("missing required argument 'ResourceId'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &LinkedServiceArgs{}
	}
	var resource LinkedService
	err := ctx.RegisterResource("azure:loganalytics/linkedService:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedService gets an existing LinkedService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azure:loganalytics/linkedService:LinkedService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedService resources.
type linkedServiceState struct {
	// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
	Name *string `pulumi:"name"`
	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
	ResourceId *string `pulumi:"resourceId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
	WorkspaceName *string `pulumi:"workspaceName"`
}

type LinkedServiceState struct {
	// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
	LinkedServiceName pulumi.StringPtrInput
	// The automatically generated name of the Linked Service. This cannot be specified. The format is always `<workspace_name>/<linked_service_name>` e.g. `workspace1/Automation`
	Name pulumi.StringPtrInput
	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
	ResourceId pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
	WorkspaceName pulumi.StringPtrInput
}

func (LinkedServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceState)(nil)).Elem()
}

type linkedServiceArgs struct {
	// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
	ResourceId string `pulumi:"resourceId"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a LinkedService resource.
type LinkedServiceArgs struct {
	// Name of the type of linkedServices resource to connect to the Log Analytics Workspace specified in `workspaceName`. Currently it defaults to and only supports `automation` as a value. Changing this forces a new resource to be created.
	LinkedServiceName pulumi.StringPtrInput
	// The name of the resource group in which the Log Analytics Linked Service is created. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// The ID of the Resource that will be linked to the workspace. Changing this forces a new resource to be created.
	ResourceId pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// Name of the Log Analytics Workspace that will contain the linkedServices resource. Changing this forces a new resource to be created.
	WorkspaceName pulumi.StringInput
}

func (LinkedServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceArgs)(nil)).Elem()
}
