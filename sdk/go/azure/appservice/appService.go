// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package appservice

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Manages an App Service (within an App Service Plan).
//
// > **Note:** When using Slots - the `appSettings`, `connectionString` and `siteConfig` blocks on the `appservice.AppService` resource will be overwritten when promoting a Slot using the `appservice.ActiveSlot` resource.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/app_service.html.markdown.
type AppService struct {
	pulumi.CustomResourceState

	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId pulumi.StringOutput `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapOutput `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings AppServiceAuthSettingsOutput `pulumi:"authSettings"`
	// A `backup` block as defined below.
	Backup AppServiceBackupPtrOutput `pulumi:"backup"`
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolOutput `pulumi:"clientAffinityEnabled"`
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled pulumi.BoolPtrOutput `pulumi:"clientCertEnabled"`
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings AppServiceConnectionStringArrayOutput `pulumi:"connectionStrings"`
	// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringOutput `pulumi:"defaultSiteHostname"`
	// Is the App Service Enabled?
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrOutput `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity AppServiceIdentityOutput `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// A `logs` block as defined below.
	Logs AppServiceLogsOutput `pulumi:"logs"`
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses pulumi.StringOutput `pulumi:"outboundIpAddresses"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses pulumi.StringOutput `pulumi:"possibleOutboundIpAddresses"`
	// The name of the resource group in which to create the App Service.
	ResourceGroupName pulumi.StringOutput `pulumi:"resourceGroupName"`
	// A `siteConfig` block as defined below.
	SiteConfig AppServiceSiteConfigOutput `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials AppServiceSiteCredentialArrayOutput `pulumi:"siteCredentials"`
	// A `sourceControl` block as defined below, which contains the Source Control information when `scmType` is set to `LocalGit`.
	SourceControls AppServiceSourceControlArrayOutput `pulumi:"sourceControls"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts AppServiceStorageAccountArrayOutput `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
}

// NewAppService registers a new resource with the given unique name, arguments, and options.
func NewAppService(ctx *pulumi.Context,
	name string, args *AppServiceArgs, opts ...pulumi.ResourceOption) (*AppService, error) {
	if args == nil || args.AppServicePlanId == nil {
		return nil, errors.New("missing required argument 'AppServicePlanId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &AppServiceArgs{}
	}
	var resource AppService
	err := ctx.RegisterResource("azure:appservice/appService:AppService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAppService gets an existing AppService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAppService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppServiceState, opts ...pulumi.ResourceOption) (*AppService, error) {
	var resource AppService
	err := ctx.ReadResource("azure:appservice/appService:AppService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AppService resources.
type appServiceState struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId *string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *AppServiceAuthSettings `pulumi:"authSettings"`
	// A `backup` block as defined below.
	Backup *AppServiceBackup `pulumi:"backup"`
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings []AppServiceConnectionString `pulumi:"connectionStrings"`
	// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
	DefaultSiteHostname *string `pulumi:"defaultSiteHostname"`
	// Is the App Service Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity *AppServiceIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `logs` block as defined below.
	Logs *AppServiceLogs `pulumi:"logs"`
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses *string `pulumi:"outboundIpAddresses"`
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses *string `pulumi:"possibleOutboundIpAddresses"`
	// The name of the resource group in which to create the App Service.
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	// A `siteConfig` block as defined below.
	SiteConfig *AppServiceSiteConfig `pulumi:"siteConfig"`
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials []AppServiceSiteCredential `pulumi:"siteCredentials"`
	// A `sourceControl` block as defined below, which contains the Source Control information when `scmType` is set to `LocalGit`.
	SourceControls []AppServiceSourceControl `pulumi:"sourceControls"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts []AppServiceStorageAccount `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

type AppServiceState struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId pulumi.StringPtrInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings AppServiceAuthSettingsPtrInput
	// A `backup` block as defined below.
	Backup AppServiceBackupPtrInput
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled pulumi.BoolPtrInput
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings AppServiceConnectionStringArrayInput
	// The Default Hostname associated with the App Service - such as `mysite.azurewebsites.net`
	DefaultSiteHostname pulumi.StringPtrInput
	// Is the App Service Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// A Managed Service Identity block as defined below.
	Identity AppServiceIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `logs` block as defined below.
	Logs AppServiceLogsPtrInput
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
	OutboundIpAddresses pulumi.StringPtrInput
	// A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
	PossibleOutboundIpAddresses pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service.
	ResourceGroupName pulumi.StringPtrInput
	// A `siteConfig` block as defined below.
	SiteConfig AppServiceSiteConfigPtrInput
	// A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
	SiteCredentials AppServiceSiteCredentialArrayInput
	// A `sourceControl` block as defined below, which contains the Source Control information when `scmType` is set to `LocalGit`.
	SourceControls AppServiceSourceControlArrayInput
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts AppServiceStorageAccountArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AppServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceState)(nil)).Elem()
}

type appServiceArgs struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId string `pulumi:"appServicePlanId"`
	// A key-value pair of App Settings.
	AppSettings map[string]string `pulumi:"appSettings"`
	// A `authSettings` block as defined below.
	AuthSettings *AppServiceAuthSettings `pulumi:"authSettings"`
	// A `backup` block as defined below.
	Backup *AppServiceBackup `pulumi:"backup"`
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled *bool `pulumi:"clientAffinityEnabled"`
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled *bool `pulumi:"clientCertEnabled"`
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings []AppServiceConnectionString `pulumi:"connectionStrings"`
	// Is the App Service Enabled?
	Enabled *bool `pulumi:"enabled"`
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly *bool `pulumi:"httpsOnly"`
	// A Managed Service Identity block as defined below.
	Identity *AppServiceIdentity `pulumi:"identity"`
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `pulumi:"location"`
	// A `logs` block as defined below.
	Logs *AppServiceLogs `pulumi:"logs"`
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The name of the resource group in which to create the App Service.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// A `siteConfig` block as defined below.
	SiteConfig *AppServiceSiteConfig `pulumi:"siteConfig"`
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts []AppServiceStorageAccount `pulumi:"storageAccounts"`
	// A mapping of tags to assign to the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a AppService resource.
type AppServiceArgs struct {
	// The ID of the App Service Plan within which to create this App Service.
	AppServicePlanId pulumi.StringInput
	// A key-value pair of App Settings.
	AppSettings pulumi.StringMapInput
	// A `authSettings` block as defined below.
	AuthSettings AppServiceAuthSettingsPtrInput
	// A `backup` block as defined below.
	Backup AppServiceBackupPtrInput
	// Should the App Service send session affinity cookies, which route client requests in the same session to the same instance?
	ClientAffinityEnabled pulumi.BoolPtrInput
	// Does the App Service require client certificates for incoming requests? Defaults to `false`.
	ClientCertEnabled pulumi.BoolPtrInput
	// One or more `connectionString` blocks as defined below.
	ConnectionStrings AppServiceConnectionStringArrayInput
	// Is the App Service Enabled?
	Enabled pulumi.BoolPtrInput
	// Can the App Service only be accessed via HTTPS? Defaults to `false`.
	HttpsOnly pulumi.BoolPtrInput
	// A Managed Service Identity block as defined below.
	Identity AppServiceIdentityPtrInput
	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location pulumi.StringPtrInput
	// A `logs` block as defined below.
	Logs AppServiceLogsPtrInput
	// Specifies the name of the App Service. Changing this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The name of the resource group in which to create the App Service.
	ResourceGroupName pulumi.StringInput
	// A `siteConfig` block as defined below.
	SiteConfig AppServiceSiteConfigPtrInput
	// One or more `storageAccount` blocks as defined below.
	StorageAccounts AppServiceStorageAccountArrayInput
	// A mapping of tags to assign to the resource.
	Tags pulumi.StringMapInput
}

func (AppServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appServiceArgs)(nil)).Elem()
}
