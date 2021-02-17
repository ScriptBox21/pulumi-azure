// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Function App.
 *
 * > **Note:** To connect an Azure Function App and a subnet within the same region `azure.appservice.VirtualNetworkSwiftConnection` can be used.
 * For an example, check the `azure.appservice.VirtualNetworkSwiftConnection` documentation.
 *
 * ## Example Usage
 * ### With App Service Plan)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "westus2"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const examplePlan = new azure.appservice.Plan("examplePlan", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     sku: {
 *         tier: "Standard",
 *         size: "S1",
 *     },
 * });
 * const exampleFunctionApp = new azure.appservice.FunctionApp("exampleFunctionApp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 *     storageAccountName: exampleAccount.name,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 * });
 * ```
 * ### In A Consumption Plan)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "westus2"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const examplePlan = new azure.appservice.Plan("examplePlan", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     kind: "FunctionApp",
 *     sku: {
 *         tier: "Dynamic",
 *         size: "Y1",
 *     },
 * });
 * const exampleFunctionApp = new azure.appservice.FunctionApp("exampleFunctionApp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 *     storageAccountName: exampleAccount.name,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 * });
 * ```
 * ### Linux)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "westus2"});
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const examplePlan = new azure.appservice.Plan("examplePlan", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     kind: "Linux",
 *     reserved: true,
 *     sku: {
 *         tier: "Dynamic",
 *         size: "Y1",
 *     },
 * });
 * const exampleFunctionApp = new azure.appservice.FunctionApp("exampleFunctionApp", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     appServicePlanId: examplePlan.id,
 *     storageAccountName: exampleAccount.name,
 *     storageAccountAccessKey: exampleAccount.primaryAccessKey,
 *     osType: "linux",
 * });
 * ```
 *
 * ## Import
 *
 * Function Apps can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:appservice/functionApp:FunctionApp functionapp1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.Web/sites/functionapp1
 * ```
 */
export class FunctionApp extends pulumi.CustomResource {
    /**
     * Get an existing FunctionApp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionAppState, opts?: pulumi.CustomResourceOptions): FunctionApp {
        return new FunctionApp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appservice/functionApp:FunctionApp';

    /**
     * Returns true if the given object is an instance of FunctionApp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FunctionApp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FunctionApp.__pulumiType;
    }

    /**
     * The ID of the App Service Plan within which to create this Function App.
     */
    public readonly appServicePlanId!: pulumi.Output<string>;
    /**
     * A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
     */
    public readonly appSettings!: pulumi.Output<{[key: string]: string}>;
    /**
     * A `authSettings` block as defined below.
     */
    public readonly authSettings!: pulumi.Output<outputs.appservice.FunctionAppAuthSettings>;
    /**
     * Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
     */
    public readonly clientAffinityEnabled!: pulumi.Output<boolean>;
    /**
     * An `connectionString` block as defined below.
     */
    public readonly connectionStrings!: pulumi.Output<outputs.appservice.FunctionAppConnectionString[]>;
    /**
     * An identifier used by App Service to perform domain ownership verification via DNS TXT record.
     */
    public /*out*/ readonly customDomainVerificationId!: pulumi.Output<string>;
    /**
     * The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
     */
    public readonly dailyMemoryTimeQuota!: pulumi.Output<number | undefined>;
    /**
     * The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
     */
    public /*out*/ readonly defaultHostname!: pulumi.Output<string>;
    /**
     * Should the built-in logging of this Function App be enabled? Defaults to `true`.
     */
    public readonly enableBuiltinLogging!: pulumi.Output<boolean | undefined>;
    /**
     * Is the Function App enabled?
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Can the Function App only be accessed via HTTPS? Defaults to `false`.
     */
    public readonly httpsOnly!: pulumi.Output<boolean | undefined>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.appservice.FunctionAppIdentity>;
    /**
     * The Function App kind - such as `functionapp,linux,container`
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Function App. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A string indicating the Operating System type for this function app.
     */
    public readonly osType!: pulumi.Output<string | undefined>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    public /*out*/ readonly outboundIpAddresses!: pulumi.Output<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
     */
    public /*out*/ readonly possibleOutboundIpAddresses!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the Function App.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    public readonly siteConfig!: pulumi.Output<outputs.appservice.FunctionAppSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    public /*out*/ readonly siteCredentials!: pulumi.Output<outputs.appservice.FunctionAppSiteCredential[]>;
    /**
     * A `sourceControl` block, as defined below.
     */
    public readonly sourceControl!: pulumi.Output<outputs.appservice.FunctionAppSourceControl>;
    /**
     * The access key which will be used to access the backend storage account for the Function App.
     */
    public readonly storageAccountAccessKey!: pulumi.Output<string>;
    /**
     * The backend storage account name which will be used by this Function App (such as the dashboard, logs).
     */
    public readonly storageAccountName!: pulumi.Output<string>;
    /**
     * @deprecated Deprecated in favour of `storage_account_name` and `storage_account_access_key`
     */
    public readonly storageConnectionString!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The runtime version associated with the Function App. Defaults to `~1`.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a FunctionApp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionAppArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionAppArgs | FunctionAppState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionAppState | undefined;
            inputs["appServicePlanId"] = state ? state.appServicePlanId : undefined;
            inputs["appSettings"] = state ? state.appSettings : undefined;
            inputs["authSettings"] = state ? state.authSettings : undefined;
            inputs["clientAffinityEnabled"] = state ? state.clientAffinityEnabled : undefined;
            inputs["connectionStrings"] = state ? state.connectionStrings : undefined;
            inputs["customDomainVerificationId"] = state ? state.customDomainVerificationId : undefined;
            inputs["dailyMemoryTimeQuota"] = state ? state.dailyMemoryTimeQuota : undefined;
            inputs["defaultHostname"] = state ? state.defaultHostname : undefined;
            inputs["enableBuiltinLogging"] = state ? state.enableBuiltinLogging : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["httpsOnly"] = state ? state.httpsOnly : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["kind"] = state ? state.kind : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["osType"] = state ? state.osType : undefined;
            inputs["outboundIpAddresses"] = state ? state.outboundIpAddresses : undefined;
            inputs["possibleOutboundIpAddresses"] = state ? state.possibleOutboundIpAddresses : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["siteConfig"] = state ? state.siteConfig : undefined;
            inputs["siteCredentials"] = state ? state.siteCredentials : undefined;
            inputs["sourceControl"] = state ? state.sourceControl : undefined;
            inputs["storageAccountAccessKey"] = state ? state.storageAccountAccessKey : undefined;
            inputs["storageAccountName"] = state ? state.storageAccountName : undefined;
            inputs["storageConnectionString"] = state ? state.storageConnectionString : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as FunctionAppArgs | undefined;
            if ((!args || args.appServicePlanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appServicePlanId'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["appServicePlanId"] = args ? args.appServicePlanId : undefined;
            inputs["appSettings"] = args ? args.appSettings : undefined;
            inputs["authSettings"] = args ? args.authSettings : undefined;
            inputs["clientAffinityEnabled"] = args ? args.clientAffinityEnabled : undefined;
            inputs["connectionStrings"] = args ? args.connectionStrings : undefined;
            inputs["dailyMemoryTimeQuota"] = args ? args.dailyMemoryTimeQuota : undefined;
            inputs["enableBuiltinLogging"] = args ? args.enableBuiltinLogging : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["httpsOnly"] = args ? args.httpsOnly : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["osType"] = args ? args.osType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["siteConfig"] = args ? args.siteConfig : undefined;
            inputs["sourceControl"] = args ? args.sourceControl : undefined;
            inputs["storageAccountAccessKey"] = args ? args.storageAccountAccessKey : undefined;
            inputs["storageAccountName"] = args ? args.storageAccountName : undefined;
            inputs["storageConnectionString"] = args ? args.storageConnectionString : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["customDomainVerificationId"] = undefined /*out*/;
            inputs["defaultHostname"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["outboundIpAddresses"] = undefined /*out*/;
            inputs["possibleOutboundIpAddresses"] = undefined /*out*/;
            inputs["siteCredentials"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(FunctionApp.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FunctionApp resources.
 */
export interface FunctionAppState {
    /**
     * The ID of the App Service Plan within which to create this Function App.
     */
    readonly appServicePlanId?: pulumi.Input<string>;
    /**
     * A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
     */
    readonly appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `authSettings` block as defined below.
     */
    readonly authSettings?: pulumi.Input<inputs.appservice.FunctionAppAuthSettings>;
    /**
     * Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * An `connectionString` block as defined below.
     */
    readonly connectionStrings?: pulumi.Input<pulumi.Input<inputs.appservice.FunctionAppConnectionString>[]>;
    /**
     * An identifier used by App Service to perform domain ownership verification via DNS TXT record.
     */
    readonly customDomainVerificationId?: pulumi.Input<string>;
    /**
     * The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
     */
    readonly dailyMemoryTimeQuota?: pulumi.Input<number>;
    /**
     * The default hostname associated with the Function App - such as `mysite.azurewebsites.net`
     */
    readonly defaultHostname?: pulumi.Input<string>;
    /**
     * Should the built-in logging of this Function App be enabled? Defaults to `true`.
     */
    readonly enableBuiltinLogging?: pulumi.Input<boolean>;
    /**
     * Is the Function App enabled?
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Can the Function App only be accessed via HTTPS? Defaults to `false`.
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.appservice.FunctionAppIdentity>;
    /**
     * The Function App kind - such as `functionapp,linux,container`
     */
    readonly kind?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Function App. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A string indicating the Operating System type for this function app.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12`
     */
    readonly outboundIpAddresses?: pulumi.Input<string>;
    /**
     * A comma separated list of outbound IP addresses - such as `52.23.25.3,52.143.43.12,52.143.43.17` - not all of which are necessarily in use. Superset of `outboundIpAddresses`.
     */
    readonly possibleOutboundIpAddresses?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Function App.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    readonly siteConfig?: pulumi.Input<inputs.appservice.FunctionAppSiteConfig>;
    /**
     * A `siteCredential` block as defined below, which contains the site-level credentials used to publish to this App Service.
     */
    readonly siteCredentials?: pulumi.Input<pulumi.Input<inputs.appservice.FunctionAppSiteCredential>[]>;
    /**
     * A `sourceControl` block, as defined below.
     */
    readonly sourceControl?: pulumi.Input<inputs.appservice.FunctionAppSourceControl>;
    /**
     * The access key which will be used to access the backend storage account for the Function App.
     */
    readonly storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * The backend storage account name which will be used by this Function App (such as the dashboard, logs).
     */
    readonly storageAccountName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favour of `storage_account_name` and `storage_account_access_key`
     */
    readonly storageConnectionString?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The runtime version associated with the Function App. Defaults to `~1`.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FunctionApp resource.
 */
export interface FunctionAppArgs {
    /**
     * The ID of the App Service Plan within which to create this Function App.
     */
    readonly appServicePlanId: pulumi.Input<string>;
    /**
     * A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
     */
    readonly appSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A `authSettings` block as defined below.
     */
    readonly authSettings?: pulumi.Input<inputs.appservice.FunctionAppAuthSettings>;
    /**
     * Should the Function App send session affinity cookies, which route client requests in the same session to the same instance?
     */
    readonly clientAffinityEnabled?: pulumi.Input<boolean>;
    /**
     * An `connectionString` block as defined below.
     */
    readonly connectionStrings?: pulumi.Input<pulumi.Input<inputs.appservice.FunctionAppConnectionString>[]>;
    /**
     * The amount of memory in gigabyte-seconds that your application is allowed to consume per day. Setting this value only affects function apps under the consumption plan. Defaults to `0`.
     */
    readonly dailyMemoryTimeQuota?: pulumi.Input<number>;
    /**
     * Should the built-in logging of this Function App be enabled? Defaults to `true`.
     */
    readonly enableBuiltinLogging?: pulumi.Input<boolean>;
    /**
     * Is the Function App enabled?
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Can the Function App only be accessed via HTTPS? Defaults to `false`.
     */
    readonly httpsOnly?: pulumi.Input<boolean>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.appservice.FunctionAppIdentity>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Function App. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A string indicating the Operating System type for this function app.
     */
    readonly osType?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the Function App.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `siteConfig` object as defined below.
     */
    readonly siteConfig?: pulumi.Input<inputs.appservice.FunctionAppSiteConfig>;
    /**
     * A `sourceControl` block, as defined below.
     */
    readonly sourceControl?: pulumi.Input<inputs.appservice.FunctionAppSourceControl>;
    /**
     * The access key which will be used to access the backend storage account for the Function App.
     */
    readonly storageAccountAccessKey?: pulumi.Input<string>;
    /**
     * The backend storage account name which will be used by this Function App (such as the dashboard, logs).
     */
    readonly storageAccountName?: pulumi.Input<string>;
    /**
     * @deprecated Deprecated in favour of `storage_account_name` and `storage_account_access_key`
     */
    readonly storageConnectionString?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The runtime version associated with the Function App. Defaults to `~1`.
     */
    readonly version?: pulumi.Input<string>;
}
