// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a MariaDB Server.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const exampleResourceGroup = new azure.core.ResourceGroup("example", {
 *     location: "West Europe",
 * });
 * const exampleServer = new azure.mariadb.Server("example", {
 *     administratorLogin: "mariadbadmin",
 *     administratorLoginPassword: "H@Sh1CoR3!",
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     skuName: "B_Gen5_2",
 *     sslEnforcement: "Enabled",
 *     storageProfile: {
 *         backupRetentionDays: 7,
 *         geoRedundantBackup: "Disabled",
 *         storageAutogrow: "Disabled",
 *         storageMb: 5120,
 *     },
 *     version: "10.2",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/mariadb_server.html.markdown.
 */
export class Server extends pulumi.CustomResource {
    /**
     * Get an existing Server resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServerState, opts?: pulumi.CustomResourceOptions): Server {
        return new Server(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mariadb/server:Server';

    /**
     * Returns true if the given object is an instance of Server.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Server {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Server.__pulumiType;
    }

    /**
     * The Administrator Login for the MariaDB Server. Changing this forces a new resource to be created.
     */
    public readonly administratorLogin!: pulumi.Output<string>;
    /**
     * The Password associated with the `administratorLogin` for the MariaDB Server.
     */
    public readonly administratorLoginPassword!: pulumi.Output<string>;
    /**
     * The FQDN of the MariaDB Server.
     */
    public /*out*/ readonly fqdn!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#sku).
     */
    public readonly skuName!: pulumi.Output<string>;
    /**
     * Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
     */
    public readonly sslEnforcement!: pulumi.Output<string>;
    /**
     * A `storageProfile` block as defined below.
     */
    public readonly storageProfile!: pulumi.Output<outputs.mariadb.ServerStorageProfile>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the version of MariaDB to use. Possible values are `10.2` and `10.3`. Changing this forces a new resource to be created.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a Server resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServerArgs | ServerState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ServerState | undefined;
            inputs["administratorLogin"] = state ? state.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = state ? state.administratorLoginPassword : undefined;
            inputs["fqdn"] = state ? state.fqdn : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["skuName"] = state ? state.skuName : undefined;
            inputs["sslEnforcement"] = state ? state.sslEnforcement : undefined;
            inputs["storageProfile"] = state ? state.storageProfile : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ServerArgs | undefined;
            if (!args || args.administratorLogin === undefined) {
                throw new Error("Missing required property 'administratorLogin'");
            }
            if (!args || args.administratorLoginPassword === undefined) {
                throw new Error("Missing required property 'administratorLoginPassword'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.skuName === undefined) {
                throw new Error("Missing required property 'skuName'");
            }
            if (!args || args.sslEnforcement === undefined) {
                throw new Error("Missing required property 'sslEnforcement'");
            }
            if (!args || args.storageProfile === undefined) {
                throw new Error("Missing required property 'storageProfile'");
            }
            if (!args || args.version === undefined) {
                throw new Error("Missing required property 'version'");
            }
            inputs["administratorLogin"] = args ? args.administratorLogin : undefined;
            inputs["administratorLoginPassword"] = args ? args.administratorLoginPassword : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["skuName"] = args ? args.skuName : undefined;
            inputs["sslEnforcement"] = args ? args.sslEnforcement : undefined;
            inputs["storageProfile"] = args ? args.storageProfile : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["fqdn"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Server.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Server resources.
 */
export interface ServerState {
    /**
     * The Administrator Login for the MariaDB Server. Changing this forces a new resource to be created.
     */
    readonly administratorLogin?: pulumi.Input<string>;
    /**
     * The Password associated with the `administratorLogin` for the MariaDB Server.
     */
    readonly administratorLoginPassword?: pulumi.Input<string>;
    /**
     * The FQDN of the MariaDB Server.
     */
    readonly fqdn?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#sku).
     */
    readonly skuName?: pulumi.Input<string>;
    /**
     * Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
     */
    readonly sslEnforcement?: pulumi.Input<string>;
    /**
     * A `storageProfile` block as defined below.
     */
    readonly storageProfile?: pulumi.Input<inputs.mariadb.ServerStorageProfile>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the version of MariaDB to use. Possible values are `10.2` and `10.3`. Changing this forces a new resource to be created.
     */
    readonly version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Server resource.
 */
export interface ServerArgs {
    /**
     * The Administrator Login for the MariaDB Server. Changing this forces a new resource to be created.
     */
    readonly administratorLogin: pulumi.Input<string>;
    /**
     * The Password associated with the `administratorLogin` for the MariaDB Server.
     */
    readonly administratorLoginPassword: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the MariaDB Server. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the MariaDB Server. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Specifies the SKU Name for this MariaDB Server. The name of the SKU, follows the `tier` + `family` + `cores` pattern (e.g. `B_Gen4_1`, `GP_Gen5_8`). For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#sku).
     */
    readonly skuName: pulumi.Input<string>;
    /**
     * Specifies if SSL should be enforced on connections. Possible values are `Enabled` and `Disabled`.
     */
    readonly sslEnforcement: pulumi.Input<string>;
    /**
     * A `storageProfile` block as defined below.
     */
    readonly storageProfile: pulumi.Input<inputs.mariadb.ServerStorageProfile>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the version of MariaDB to use. Possible values are `10.2` and `10.3`. Changing this forces a new resource to be created.
     */
    readonly version: pulumi.Input<string>;
}
