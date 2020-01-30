// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package storage

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an Azure Storage Account.
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/storage_account.html.markdown.
type Account struct {
	pulumi.CustomResourceState

	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier pulumi.StringOutput `pulumi:"accessTier"`
	// The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
	AccountEncryptionSource pulumi.StringPtrOutput `pulumi:"accountEncryptionSource"`
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `Storage`.
	AccountKind pulumi.StringPtrOutput `pulumi:"accountKind"`
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
	AccountReplicationType pulumi.StringOutput `pulumi:"accountReplicationType"`
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier pulumi.StringOutput `pulumi:"accountTier"`
	AccountType pulumi.StringOutput `pulumi:"accountType"`
	// A `blobProperties` block as defined below.
	BlobProperties AccountBlobPropertiesOutput `pulumi:"blobProperties"`
	// A `customDomain` block as documented below.
	CustomDomain AccountCustomDomainPtrOutput `pulumi:"customDomain"`
	// Boolean flag which controls if advanced threat protection is enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-advanced-threat-protection) for more information. Defaults to `false`.
	EnableAdvancedThreatProtection pulumi.BoolOutput `pulumi:"enableAdvancedThreatProtection"`
	// Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableBlobEncryption pulumi.BoolPtrOutput `pulumi:"enableBlobEncryption"`
	// Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableFileEncryption pulumi.BoolPtrOutput `pulumi:"enableFileEncryption"`
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information.
	EnableHttpsTrafficOnly pulumi.BoolPtrOutput `pulumi:"enableHttpsTrafficOnly"`
	// A `identity` block as defined below.
	Identity AccountIdentityOutput `pulumi:"identity"`
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled pulumi.BoolPtrOutput `pulumi:"isHnsEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name pulumi.StringOutput `pulumi:"name"`
	// A `networkRules` block as documented below.
	NetworkRules AccountNetworkRulesTypeOutput `pulumi:"networkRules"`
	// The primary access key for the storage account.
	PrimaryAccessKey pulumi.StringOutput `pulumi:"primaryAccessKey"`
	// The connection string associated with the primary blob location.
	PrimaryBlobConnectionString pulumi.StringOutput `pulumi:"primaryBlobConnectionString"`
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint pulumi.StringOutput `pulumi:"primaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the primary location.
	PrimaryBlobHost pulumi.StringOutput `pulumi:"primaryBlobHost"`
	// The connection string associated with the primary location.
	PrimaryConnectionString pulumi.StringOutput `pulumi:"primaryConnectionString"`
	// The endpoint URL for DFS storage in the primary location.
	PrimaryDfsEndpoint pulumi.StringOutput `pulumi:"primaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the primary location.
	PrimaryDfsHost pulumi.StringOutput `pulumi:"primaryDfsHost"`
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint pulumi.StringOutput `pulumi:"primaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the primary location.
	PrimaryFileHost pulumi.StringOutput `pulumi:"primaryFileHost"`
	// The primary location of the storage account.
	PrimaryLocation pulumi.StringOutput `pulumi:"primaryLocation"`
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint pulumi.StringOutput `pulumi:"primaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the primary location.
	PrimaryQueueHost pulumi.StringOutput `pulumi:"primaryQueueHost"`
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint pulumi.StringOutput `pulumi:"primaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the primary location.
	PrimaryTableHost pulumi.StringOutput `pulumi:"primaryTableHost"`
	// The endpoint URL for web storage in the primary location.
	PrimaryWebEndpoint pulumi.StringOutput `pulumi:"primaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the primary location.
	PrimaryWebHost pulumi.StringOutput `pulumi:"primaryWebHost"`
	// A `queueProperties` block as defined below.
	QueueProperties AccountQueuePropertiesOutput `pulumi:"queueProperties"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// The secondary access key for the storage account.
	SecondaryAccessKey pulumi.StringOutput `pulumi:"secondaryAccessKey"`
	// The connection string associated with the secondary blob location.
	SecondaryBlobConnectionString pulumi.StringOutput `pulumi:"secondaryBlobConnectionString"`
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint pulumi.StringOutput `pulumi:"secondaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the secondary location.
	SecondaryBlobHost pulumi.StringOutput `pulumi:"secondaryBlobHost"`
	// The connection string associated with the secondary location.
	SecondaryConnectionString pulumi.StringOutput `pulumi:"secondaryConnectionString"`
	// The endpoint URL for DFS storage in the secondary location.
	SecondaryDfsEndpoint pulumi.StringOutput `pulumi:"secondaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the secondary location.
	SecondaryDfsHost pulumi.StringOutput `pulumi:"secondaryDfsHost"`
	// The endpoint URL for file storage in the secondary location.
	SecondaryFileEndpoint pulumi.StringOutput `pulumi:"secondaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the secondary location.
	SecondaryFileHost pulumi.StringOutput `pulumi:"secondaryFileHost"`
	// The secondary location of the storage account.
	SecondaryLocation pulumi.StringOutput `pulumi:"secondaryLocation"`
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint pulumi.StringOutput `pulumi:"secondaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the secondary location.
	SecondaryQueueHost pulumi.StringOutput `pulumi:"secondaryQueueHost"`
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint pulumi.StringOutput `pulumi:"secondaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the secondary location.
	SecondaryTableHost pulumi.StringOutput `pulumi:"secondaryTableHost"`
	// The endpoint URL for web storage in the secondary location.
	SecondaryWebEndpoint pulumi.StringOutput `pulumi:"secondaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the secondary location.
	SecondaryWebHost pulumi.StringOutput `pulumi:"secondaryWebHost"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAccount registers a new resource with the given unique name, arguments, and options.
func NewAccount(ctx *pulumi.Context,
	name string, args *AccountArgs, opts ...pulumi.ResourceOption) (*Account, error) {
	if args == nil || args.AccountReplicationType == nil {
		return nil, errors.New("missing required argument 'AccountReplicationType'")
	}
	if args == nil || args.AccountTier == nil {
		return nil, errors.New("missing required argument 'AccountTier'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AccountArgs{}
	}
	var resource Account
	err := ctx.RegisterResource("azure:storage/account:Account", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccount gets an existing Account resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccountState, opts ...pulumi.ResourceOption) (*Account, error) {
	var resource Account
	err := ctx.ReadResource("azure:storage/account:Account", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Account resources.
type accountState struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier *string `pulumi:"accessTier"`
	// The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
	AccountEncryptionSource *string `pulumi:"accountEncryptionSource"`
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `Storage`.
	AccountKind *string `pulumi:"accountKind"`
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
	AccountReplicationType *string `pulumi:"accountReplicationType"`
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier *string `pulumi:"accountTier"`
	AccountType *string `pulumi:"accountType"`
	// A `blobProperties` block as defined below.
	BlobProperties *AccountBlobProperties `pulumi:"blobProperties"`
	// A `customDomain` block as documented below.
	CustomDomain *AccountCustomDomain `pulumi:"customDomain"`
	// Boolean flag which controls if advanced threat protection is enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-advanced-threat-protection) for more information. Defaults to `false`.
	EnableAdvancedThreatProtection *bool `pulumi:"enableAdvancedThreatProtection"`
	// Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableBlobEncryption *bool `pulumi:"enableBlobEncryption"`
	// Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableFileEncryption *bool `pulumi:"enableFileEncryption"`
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information.
	EnableHttpsTrafficOnly *bool `pulumi:"enableHttpsTrafficOnly"`
	// A `identity` block as defined below.
	Identity *AccountIdentity `pulumi:"identity"`
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled *bool `pulumi:"isHnsEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name *string `pulumi:"name"`
	// A `networkRules` block as documented below.
	NetworkRules *AccountNetworkRulesType `pulumi:"networkRules"`
	// The primary access key for the storage account.
	PrimaryAccessKey *string `pulumi:"primaryAccessKey"`
	// The connection string associated with the primary blob location.
	PrimaryBlobConnectionString *string `pulumi:"primaryBlobConnectionString"`
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint *string `pulumi:"primaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the primary location.
	PrimaryBlobHost *string `pulumi:"primaryBlobHost"`
	// The connection string associated with the primary location.
	PrimaryConnectionString *string `pulumi:"primaryConnectionString"`
	// The endpoint URL for DFS storage in the primary location.
	PrimaryDfsEndpoint *string `pulumi:"primaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the primary location.
	PrimaryDfsHost *string `pulumi:"primaryDfsHost"`
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint *string `pulumi:"primaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the primary location.
	PrimaryFileHost *string `pulumi:"primaryFileHost"`
	// The primary location of the storage account.
	PrimaryLocation *string `pulumi:"primaryLocation"`
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint *string `pulumi:"primaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the primary location.
	PrimaryQueueHost *string `pulumi:"primaryQueueHost"`
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint *string `pulumi:"primaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the primary location.
	PrimaryTableHost *string `pulumi:"primaryTableHost"`
	// The endpoint URL for web storage in the primary location.
	PrimaryWebEndpoint *string `pulumi:"primaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the primary location.
	PrimaryWebHost *string `pulumi:"primaryWebHost"`
	// A `queueProperties` block as defined below.
	QueueProperties *AccountQueueProperties `pulumi:"queueProperties"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// The secondary access key for the storage account.
	SecondaryAccessKey *string `pulumi:"secondaryAccessKey"`
	// The connection string associated with the secondary blob location.
	SecondaryBlobConnectionString *string `pulumi:"secondaryBlobConnectionString"`
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint *string `pulumi:"secondaryBlobEndpoint"`
	// The hostname with port if applicable for blob storage in the secondary location.
	SecondaryBlobHost *string `pulumi:"secondaryBlobHost"`
	// The connection string associated with the secondary location.
	SecondaryConnectionString *string `pulumi:"secondaryConnectionString"`
	// The endpoint URL for DFS storage in the secondary location.
	SecondaryDfsEndpoint *string `pulumi:"secondaryDfsEndpoint"`
	// The hostname with port if applicable for DFS storage in the secondary location.
	SecondaryDfsHost *string `pulumi:"secondaryDfsHost"`
	// The endpoint URL for file storage in the secondary location.
	SecondaryFileEndpoint *string `pulumi:"secondaryFileEndpoint"`
	// The hostname with port if applicable for file storage in the secondary location.
	SecondaryFileHost *string `pulumi:"secondaryFileHost"`
	// The secondary location of the storage account.
	SecondaryLocation *string `pulumi:"secondaryLocation"`
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint *string `pulumi:"secondaryQueueEndpoint"`
	// The hostname with port if applicable for queue storage in the secondary location.
	SecondaryQueueHost *string `pulumi:"secondaryQueueHost"`
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint *string `pulumi:"secondaryTableEndpoint"`
	// The hostname with port if applicable for table storage in the secondary location.
	SecondaryTableHost *string `pulumi:"secondaryTableHost"`
	// The endpoint URL for web storage in the secondary location.
	SecondaryWebEndpoint *string `pulumi:"secondaryWebEndpoint"`
	// The hostname with port if applicable for web storage in the secondary location.
	SecondaryWebHost *string `pulumi:"secondaryWebHost"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AccountState struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier pulumi.StringPtrInput
	// The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
	AccountEncryptionSource pulumi.StringPtrInput
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `Storage`.
	AccountKind pulumi.StringPtrInput
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
	AccountReplicationType pulumi.StringPtrInput
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier pulumi.StringPtrInput
	AccountType pulumi.StringPtrInput
	// A `blobProperties` block as defined below.
	BlobProperties AccountBlobPropertiesPtrInput
	// A `customDomain` block as documented below.
	CustomDomain AccountCustomDomainPtrInput
	// Boolean flag which controls if advanced threat protection is enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-advanced-threat-protection) for more information. Defaults to `false`.
	EnableAdvancedThreatProtection pulumi.BoolPtrInput
	// Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableBlobEncryption pulumi.BoolPtrInput
	// Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableFileEncryption pulumi.BoolPtrInput
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information.
	EnableHttpsTrafficOnly pulumi.BoolPtrInput
	// A `identity` block as defined below.
	Identity AccountIdentityPtrInput
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name pulumi.StringPtrInput
	// A `networkRules` block as documented below.
	NetworkRules AccountNetworkRulesTypePtrInput
	// The primary access key for the storage account.
	PrimaryAccessKey pulumi.StringPtrInput
	// The connection string associated with the primary blob location.
	PrimaryBlobConnectionString pulumi.StringPtrInput
	// The endpoint URL for blob storage in the primary location.
	PrimaryBlobEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for blob storage in the primary location.
	PrimaryBlobHost pulumi.StringPtrInput
	// The connection string associated with the primary location.
	PrimaryConnectionString pulumi.StringPtrInput
	// The endpoint URL for DFS storage in the primary location.
	PrimaryDfsEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for DFS storage in the primary location.
	PrimaryDfsHost pulumi.StringPtrInput
	// The endpoint URL for file storage in the primary location.
	PrimaryFileEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for file storage in the primary location.
	PrimaryFileHost pulumi.StringPtrInput
	// The primary location of the storage account.
	PrimaryLocation pulumi.StringPtrInput
	// The endpoint URL for queue storage in the primary location.
	PrimaryQueueEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for queue storage in the primary location.
	PrimaryQueueHost pulumi.StringPtrInput
	// The endpoint URL for table storage in the primary location.
	PrimaryTableEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for table storage in the primary location.
	PrimaryTableHost pulumi.StringPtrInput
	// The endpoint URL for web storage in the primary location.
	PrimaryWebEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for web storage in the primary location.
	PrimaryWebHost pulumi.StringPtrInput
	// A `queueProperties` block as defined below.
	QueueProperties AccountQueuePropertiesPtrInput
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringPtrInput
	// The secondary access key for the storage account.
	SecondaryAccessKey pulumi.StringPtrInput
	// The connection string associated with the secondary blob location.
	SecondaryBlobConnectionString pulumi.StringPtrInput
	// The endpoint URL for blob storage in the secondary location.
	SecondaryBlobEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for blob storage in the secondary location.
	SecondaryBlobHost pulumi.StringPtrInput
	// The connection string associated with the secondary location.
	SecondaryConnectionString pulumi.StringPtrInput
	// The endpoint URL for DFS storage in the secondary location.
	SecondaryDfsEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for DFS storage in the secondary location.
	SecondaryDfsHost pulumi.StringPtrInput
	// The endpoint URL for file storage in the secondary location.
	SecondaryFileEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for file storage in the secondary location.
	SecondaryFileHost pulumi.StringPtrInput
	// The secondary location of the storage account.
	SecondaryLocation pulumi.StringPtrInput
	// The endpoint URL for queue storage in the secondary location.
	SecondaryQueueEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for queue storage in the secondary location.
	SecondaryQueueHost pulumi.StringPtrInput
	// The endpoint URL for table storage in the secondary location.
	SecondaryTableEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for table storage in the secondary location.
	SecondaryTableHost pulumi.StringPtrInput
	// The endpoint URL for web storage in the secondary location.
	SecondaryWebEndpoint pulumi.StringPtrInput
	// The hostname with port if applicable for web storage in the secondary location.
	SecondaryWebHost pulumi.StringPtrInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*accountState)(nil)).Elem()
}

type accountArgs struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier *string `pulumi:"accessTier"`
	// The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
	AccountEncryptionSource *string `pulumi:"accountEncryptionSource"`
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `Storage`.
	AccountKind *string `pulumi:"accountKind"`
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
	AccountReplicationType string `pulumi:"accountReplicationType"`
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier string `pulumi:"accountTier"`
	AccountType *string `pulumi:"accountType"`
	// A `blobProperties` block as defined below.
	BlobProperties *AccountBlobProperties `pulumi:"blobProperties"`
	// A `customDomain` block as documented below.
	CustomDomain *AccountCustomDomain `pulumi:"customDomain"`
	// Boolean flag which controls if advanced threat protection is enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-advanced-threat-protection) for more information. Defaults to `false`.
	EnableAdvancedThreatProtection *bool `pulumi:"enableAdvancedThreatProtection"`
	// Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableBlobEncryption *bool `pulumi:"enableBlobEncryption"`
	// Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableFileEncryption *bool `pulumi:"enableFileEncryption"`
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information.
	EnableHttpsTrafficOnly *bool `pulumi:"enableHttpsTrafficOnly"`
	// A `identity` block as defined below.
	Identity *AccountIdentity `pulumi:"identity"`
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled *bool `pulumi:"isHnsEnabled"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name *string `pulumi:"name"`
	// A `networkRules` block as documented below.
	NetworkRules *AccountNetworkRulesType `pulumi:"networkRules"`
	// A `queueProperties` block as defined below.
	QueueProperties *AccountQueueProperties `pulumi:"queueProperties"`
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Account resource.
type AccountArgs struct {
	// Defines the access tier for `BlobStorage`, `FileStorage` and `StorageV2` accounts. Valid options are `Hot` and `Cool`, defaults to `Hot`.
	AccessTier pulumi.StringPtrInput
	// The Encryption Source for this Storage Account. Possible values are `Microsoft.Keyvault` and `Microsoft.Storage`. Defaults to `Microsoft.Storage`.
	AccountEncryptionSource pulumi.StringPtrInput
	// Defines the Kind of account. Valid options are `BlobStorage`, `BlockBlobStorage`, `FileStorage`, `Storage` and `StorageV2`. Changing this forces a new resource to be created. Defaults to `Storage`.
	AccountKind pulumi.StringPtrInput
	// Defines the type of replication to use for this storage account. Valid options are `LRS`, `GRS`, `RAGRS` and `ZRS`.
	AccountReplicationType pulumi.StringInput
	// Defines the Tier to use for this storage account. Valid options are `Standard` and `Premium`. For `FileStorage` accounts only `Premium` is valid. Changing this forces a new resource to be created.
	AccountTier pulumi.StringInput
	AccountType pulumi.StringPtrInput
	// A `blobProperties` block as defined below.
	BlobProperties AccountBlobPropertiesPtrInput
	// A `customDomain` block as documented below.
	CustomDomain AccountCustomDomainPtrInput
	// Boolean flag which controls if advanced threat protection is enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/common/storage-advanced-threat-protection) for more information. Defaults to `false`.
	EnableAdvancedThreatProtection pulumi.BoolPtrInput
	// Boolean flag which controls if Encryption Services are enabled for Blob storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableBlobEncryption pulumi.BoolPtrInput
	// Boolean flag which controls if Encryption Services are enabled for File storage, see [here](https://azure.microsoft.com/en-us/documentation/articles/storage-service-encryption/) for more information. Defaults to `true`.
	EnableFileEncryption pulumi.BoolPtrInput
	// Boolean flag which forces HTTPS if enabled, see [here](https://docs.microsoft.com/en-us/azure/storage/storage-require-secure-transfer/)
	// for more information.
	EnableHttpsTrafficOnly pulumi.BoolPtrInput
	// A `identity` block as defined below.
	Identity AccountIdentityPtrInput
	// Is Hierarchical Namespace enabled? This can be used with Azure Data Lake Storage Gen 2 ([see here for more information](https://docs.microsoft.com/en-us/azure/storage/blobs/data-lake-storage-quickstart-create-account/)). Changing this forces a new resource to be created.
	IsHnsEnabled pulumi.BoolPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// Specifies the name of the storage account. Changing this forces a new resource to be created. This must be unique across the entire Azure service, not just within the resource group.
	Name pulumi.StringPtrInput
	// A `networkRules` block as documented below.
	NetworkRules AccountNetworkRulesTypePtrInput
	// A `queueProperties` block as defined below.
	QueueProperties AccountQueuePropertiesPtrInput
	// The name of the resource group in which to create the storage account. Changing this forces a new resource to be created.
	ResourceGroupName pulumi.StringInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accountArgs)(nil)).Elem()
}

