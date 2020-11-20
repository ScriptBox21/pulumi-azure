// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages an Azure HTTP Dataset inside an Azure Data Factory.
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
// 		exampleLinkedServiceWeb, err := datafactory.NewLinkedServiceWeb(ctx, "exampleLinkedServiceWeb", &datafactory.LinkedServiceWebArgs{
// 			ResourceGroupName:  exampleResourceGroup.Name,
// 			DataFactoryName:    exampleFactory.Name,
// 			AuthenticationType: pulumi.String("Anonymous"),
// 			Url:                pulumi.String("https://www.bing.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = datafactory.NewDatasetHttp(ctx, "exampleDatasetHttp", &datafactory.DatasetHttpArgs{
// 			ResourceGroupName: exampleResourceGroup.Name,
// 			DataFactoryName:   exampleFactory.Name,
// 			LinkedServiceName: exampleLinkedServiceWeb.Name,
// 			RelativeUrl:       pulumi.String("http://www.bing.com"),
// 			RequestBody:       pulumi.String("foo=bar"),
// 			RequestMethod:     pulumi.String("POST"),
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
// Data Factory Datasets can be imported using the `resource id`, e.g.
//
// ```sh
//  $ pulumi import azure:datafactory/datasetHttp:DatasetHttp example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example/providers/Microsoft.DataFactory/factories/example/datasets/example
// ```
type DatasetHttp struct {
	pulumi.CustomResourceState

	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapOutput `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayOutput `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringOutput `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrOutput `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringOutput `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringOutput `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapOutput `pulumi:"parameters"`
	// The relative URL based on the URL in the HTTP Linked Service.
	RelativeUrl pulumi.StringPtrOutput `pulumi:"relativeUrl"`
	// The body for the HTTP request.
	RequestBody pulumi.StringPtrOutput `pulumi:"requestBody"`
	// The HTTP method for the HTTP request. (e.g. GET, POST)
	RequestMethod pulumi.StringPtrOutput `pulumi:"requestMethod"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetHttpSchemaColumnArrayOutput `pulumi:"schemaColumns"`
}

// NewDatasetHttp registers a new resource with the given unique name, arguments, and options.
func NewDatasetHttp(ctx *pulumi.Context,
	name string, args *DatasetHttpArgs, opts ...pulumi.ResourceOption) (*DatasetHttp, error) {
	if args == nil || args.DataFactoryName == nil {
		return nil, errors.New("missing required argument 'DataFactoryName'")
	}
	if args == nil || args.LinkedServiceName == nil {
		return nil, errors.New("missing required argument 'LinkedServiceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &DatasetHttpArgs{}
	}
	var resource DatasetHttp
	err := ctx.RegisterResource("azure:datafactory/datasetHttp:DatasetHttp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatasetHttp gets an existing DatasetHttp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatasetHttp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatasetHttpState, opts ...pulumi.ResourceOption) (*DatasetHttp, error) {
	var resource DatasetHttp
	err := ctx.ReadResource("azure:datafactory/datasetHttp:DatasetHttp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatasetHttp resources.
type datasetHttpState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName *string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName *string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The relative URL based on the URL in the HTTP Linked Service.
	RelativeUrl *string `pulumi:"relativeUrl"`
	// The body for the HTTP request.
	RequestBody *string `pulumi:"requestBody"`
	// The HTTP method for the HTTP request. (e.g. GET, POST)
	RequestMethod *string `pulumi:"requestMethod"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetHttpSchemaColumn `pulumi:"schemaColumns"`
}

type DatasetHttpState struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringPtrInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringPtrInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The relative URL based on the URL in the HTTP Linked Service.
	RelativeUrl pulumi.StringPtrInput
	// The body for the HTTP request.
	RequestBody pulumi.StringPtrInput
	// The HTTP method for the HTTP request. (e.g. GET, POST)
	RequestMethod pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringPtrInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetHttpSchemaColumnArrayInput
}

func (DatasetHttpState) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetHttpState)(nil)).Elem()
}

type datasetHttpArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties map[string]string `pulumi:"additionalProperties"`
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations []string `pulumi:"annotations"`
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName string `pulumi:"dataFactoryName"`
	// The description for the Data Factory Dataset.
	Description *string `pulumi:"description"`
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder *string `pulumi:"folder"`
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name *string `pulumi:"name"`
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters map[string]string `pulumi:"parameters"`
	// The relative URL based on the URL in the HTTP Linked Service.
	RelativeUrl *string `pulumi:"relativeUrl"`
	// The body for the HTTP request.
	RequestBody *string `pulumi:"requestBody"`
	// The HTTP method for the HTTP request. (e.g. GET, POST)
	RequestMethod *string `pulumi:"requestMethod"`
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `schemaColumn` block as defined below.
	SchemaColumns []DatasetHttpSchemaColumn `pulumi:"schemaColumns"`
}

// The set of arguments for constructing a DatasetHttp resource.
type DatasetHttpArgs struct {
	// A map of additional properties to associate with the Data Factory Dataset.
	AdditionalProperties pulumi.StringMapInput
	// List of tags that can be used for describing the Data Factory Dataset.
	Annotations pulumi.StringArrayInput
	// The Data Factory name in which to associate the Dataset with. Changing this forces a new resource.
	DataFactoryName pulumi.StringInput
	// The description for the Data Factory Dataset.
	Description pulumi.StringPtrInput
	// The folder that this Dataset is in. If not specified, the Dataset will appear at the root level.
	Folder pulumi.StringPtrInput
	// The Data Factory Linked Service name in which to associate the Dataset with.
	LinkedServiceName pulumi.StringInput
	// Specifies the name of the Data Factory Dataset. Changing this forces a new resource to be created. Must be globally unique. See the [Microsoft documentation](https://docs.microsoft.com/en-us/azure/data-factory/naming-rules) for all restrictions.
	Name pulumi.StringPtrInput
	// A map of parameters to associate with the Data Factory Dataset.
	Parameters pulumi.StringMapInput
	// The relative URL based on the URL in the HTTP Linked Service.
	RelativeUrl pulumi.StringPtrInput
	// The body for the HTTP request.
	RequestBody pulumi.StringPtrInput
	// The HTTP method for the HTTP request. (e.g. GET, POST)
	RequestMethod pulumi.StringPtrInput
	// The name of the resource group in which to create the Data Factory Dataset. Changing this forces a new resource
	ResourceGroupName pulumi.StringInput
	// A `schemaColumn` block as defined below.
	SchemaColumns DatasetHttpSchemaColumnArrayInput
}

func (DatasetHttpArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*datasetHttpArgs)(nil)).Elem()
}

type DatasetHttpInput interface {
	pulumi.Input

	ToDatasetHttpOutput() DatasetHttpOutput
	ToDatasetHttpOutputWithContext(ctx context.Context) DatasetHttpOutput
}

func (DatasetHttp) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetHttp)(nil)).Elem()
}

func (i DatasetHttp) ToDatasetHttpOutput() DatasetHttpOutput {
	return i.ToDatasetHttpOutputWithContext(context.Background())
}

func (i DatasetHttp) ToDatasetHttpOutputWithContext(ctx context.Context) DatasetHttpOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatasetHttpOutput)
}

type DatasetHttpOutput struct {
	*pulumi.OutputState
}

func (DatasetHttpOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatasetHttpOutput)(nil)).Elem()
}

func (o DatasetHttpOutput) ToDatasetHttpOutput() DatasetHttpOutput {
	return o
}

func (o DatasetHttpOutput) ToDatasetHttpOutputWithContext(ctx context.Context) DatasetHttpOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatasetHttpOutput{})
}
