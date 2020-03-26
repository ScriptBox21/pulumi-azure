// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kusto

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a Kusto (also known as Azure Data Explorer) Cluster
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/kusto_cluster.html.markdown.
type Cluster struct {
	pulumi.CustomResourceState

	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri pulumi.StringOutput `pulumi:"dataIngestionUri"`
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrOutput `pulumi:"enableDiskEncryption"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrOutput `pulumi:"enableStreamingIngest"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ClusterSkuOutput `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The FQDN of the Azure Kusto Cluster.
	Uri pulumi.StringOutput `pulumi:"uri"`
}

// NewCluster registers a new resource with the given unique name, arguments, and options.
func NewCluster(ctx *pulumi.Context,
	name string, args *ClusterArgs, opts ...pulumi.ResourceOption) (*Cluster, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &ClusterArgs{}
	}
	var resource Cluster
	err := ctx.RegisterResource("azure:kusto/cluster:Cluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCluster gets an existing Cluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterState, opts ...pulumi.ResourceOption) (*Cluster, error) {
	var resource Cluster
	err := ctx.ReadResource("azure:kusto/cluster:Cluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cluster resources.
type clusterState struct {
	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri *string `pulumi:"dataIngestionUri"`
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku *ClusterSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
	// The FQDN of the Azure Kusto Cluster.
	Uri *string `pulumi:"uri"`
}

type ClusterState struct {
	// The Kusto Cluster URI to be used for data ingestion.
	DataIngestionUri pulumi.StringPtrInput
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `sku` block as defined below.
	Sku ClusterSkuPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
	// The FQDN of the Azure Kusto Cluster.
	Uri pulumi.StringPtrInput
}

func (ClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterState)(nil)).Elem()
}

type clusterArgs struct {
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption *bool `pulumi:"enableDiskEncryption"`
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest *bool `pulumi:"enableStreamingIngest"`
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `sku` block as defined below.
	Sku ClusterSku `pulumi:"sku"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Cluster resource.
type ClusterArgs struct {
	// Specifies if the cluster's disks are encrypted.
	EnableDiskEncryption pulumi.BoolPtrInput
	// Specifies if the streaming ingest is enabled.
	EnableStreamingIngest pulumi.BoolPtrInput
	// The location where the Kusto Cluster should be created. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// The name of the Kusto Cluster to create. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the Resource Group where the Kusto Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `sku` block as defined below.
	Sku ClusterSkuInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (ClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterArgs)(nil)).Elem()
}
