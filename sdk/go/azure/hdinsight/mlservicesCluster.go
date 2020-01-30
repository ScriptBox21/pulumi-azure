// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package hdinsight

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages a HDInsight ML Services Cluster.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/hdinsight_ml_services_cluster.html.markdown.
type MLServicesCluster struct {
	pulumi.CustomResourceState

	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringOutput `pulumi:"clusterVersion"`
	// The SSH Connectivity Endpoint for the Edge Node of the HDInsight ML Cluster.
	EdgeSshEndpoint pulumi.StringOutput `pulumi:"edgeSshEndpoint"`
	// A `gateway` block as defined below.
	Gateway MLServicesClusterGatewayOutput `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight ML Services Cluster.
	HttpsEndpoint pulumi.StringOutput `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name for this HDInsight ML Services Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles MLServicesClusterRolesOutput `pulumi:"roles"`
	// Should R Studio community edition for ML Services be installed? Changing this forces a new resource to be created.
	Rstudio pulumi.BoolOutput `pulumi:"rstudio"`
	// The SSH Connectivity Endpoint for this HDInsight ML Services Cluster.
	SshEndpoint pulumi.StringOutput `pulumi:"sshEndpoint"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts MLServicesClusterStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight ML Services Cluster.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight ML Services Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier pulumi.StringOutput `pulumi:"tier"`
}

// NewMLServicesCluster registers a new resource with the given unique name, arguments, and options.
func NewMLServicesCluster(ctx *pulumi.Context,
	name string, args *MLServicesClusterArgs, opts ...pulumi.ResourceOption) (*MLServicesCluster, error) {
	if args == nil || args.ClusterVersion == nil {
		return nil, errors.New("missing required argument 'ClusterVersion'")
	}
	if args == nil || args.Gateway == nil {
		return nil, errors.New("missing required argument 'Gateway'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Roles == nil {
		return nil, errors.New("missing required argument 'Roles'")
	}
	if args == nil || args.Rstudio == nil {
		return nil, errors.New("missing required argument 'Rstudio'")
	}
	if args == nil || args.Tier == nil {
		return nil, errors.New("missing required argument 'Tier'")
	}
	if args == nil {
		args = &MLServicesClusterArgs{}
	}
	var resource MLServicesCluster
	err := ctx.RegisterResource("azure:hdinsight/mLServicesCluster:MLServicesCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMLServicesCluster gets an existing MLServicesCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMLServicesCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MLServicesClusterState, opts ...pulumi.ResourceOption) (*MLServicesCluster, error) {
	var resource MLServicesCluster
	err := ctx.ReadResource("azure:hdinsight/mLServicesCluster:MLServicesCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MLServicesCluster resources.
type mlservicesClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion *string `pulumi:"clusterVersion"`
	// The SSH Connectivity Endpoint for the Edge Node of the HDInsight ML Cluster.
	EdgeSshEndpoint *string `pulumi:"edgeSshEndpoint"`
	// A `gateway` block as defined below.
	Gateway *MLServicesClusterGateway `pulumi:"gateway"`
	// The HTTPS Connectivity Endpoint for this HDInsight ML Services Cluster.
	HttpsEndpoint *string `pulumi:"httpsEndpoint"`
	// Specifies the Azure Region which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name for this HDInsight ML Services Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles *MLServicesClusterRoles `pulumi:"roles"`
	// Should R Studio community edition for ML Services be installed? Changing this forces a new resource to be created.
	Rstudio *bool `pulumi:"rstudio"`
	// The SSH Connectivity Endpoint for this HDInsight ML Services Cluster.
	SshEndpoint *string `pulumi:"sshEndpoint"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []MLServicesClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight ML Services Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight ML Services Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier *string `pulumi:"tier"`
}

type MLServicesClusterState struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringPtrInput
	// The SSH Connectivity Endpoint for the Edge Node of the HDInsight ML Cluster.
	EdgeSshEndpoint pulumi.StringPtrInput
	// A `gateway` block as defined below.
	Gateway MLServicesClusterGatewayPtrInput
	// The HTTPS Connectivity Endpoint for this HDInsight ML Services Cluster.
	HttpsEndpoint pulumi.StringPtrInput
	// Specifies the Azure Region which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name for this HDInsight ML Services Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// A `roles` block as defined below.
	Roles MLServicesClusterRolesPtrInput
	// Should R Studio community edition for ML Services be installed? Changing this forces a new resource to be created.
	Rstudio pulumi.BoolPtrInput
	// The SSH Connectivity Endpoint for this HDInsight ML Services Cluster.
	SshEndpoint pulumi.StringPtrInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts MLServicesClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight ML Services Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight ML Services Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier pulumi.StringPtrInput
}

func (MLServicesClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mlservicesClusterState)(nil)).Elem()
}

type mlservicesClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion string `pulumi:"clusterVersion"`
	// A `gateway` block as defined below.
	Gateway MLServicesClusterGateway `pulumi:"gateway"`
	// Specifies the Azure Region which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name for this HDInsight ML Services Cluster. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// Specifies the name of the Resource Group in which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `roles` block as defined below.
	Roles MLServicesClusterRoles `pulumi:"roles"`
	// Should R Studio community edition for ML Services be installed? Changing this forces a new resource to be created.
	Rstudio bool `pulumi:"rstudio"`
	// One or more `storageAccount` block as defined below.
	StorageAccounts []MLServicesClusterStorageAccount `pulumi:"storageAccounts"`
	// A map of Tags which should be assigned to this HDInsight ML Services Cluster.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the Tier which should be used for this HDInsight ML Services Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier string `pulumi:"tier"`
}

// The set of arguments for constructing a MLServicesCluster resource.
type MLServicesClusterArgs struct {
	// Specifies the Version of HDInsights which should be used for this Cluster. Changing this forces a new resource to be created.
	ClusterVersion pulumi.StringInput
	// A `gateway` block as defined below.
	Gateway MLServicesClusterGatewayInput
	// Specifies the Azure Region which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name for this HDInsight ML Services Cluster. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// Specifies the name of the Resource Group in which this HDInsight ML Services Cluster should exist. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A `roles` block as defined below.
	Roles MLServicesClusterRolesInput
	// Should R Studio community edition for ML Services be installed? Changing this forces a new resource to be created.
	Rstudio pulumi.BoolInput
	// One or more `storageAccount` block as defined below.
	StorageAccounts MLServicesClusterStorageAccountArrayInput
	// A map of Tags which should be assigned to this HDInsight ML Services Cluster.
	Tags pulumi.StringMapInput
	// Specifies the Tier which should be used for this HDInsight ML Services Cluster. Possible values are `Standard` or `Premium`. Changing this forces a new resource to be created.
	Tier pulumi.StringInput
}

func (MLServicesClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mlservicesClusterArgs)(nil)).Elem()
}

