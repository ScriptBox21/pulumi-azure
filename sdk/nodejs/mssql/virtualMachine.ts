// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a Microsoft SQL Virtual Machine
 *
 * ## Example Usage
 *
 * This example provisions a brief Managed MsSql Virtual Machine.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleVirtualMachine = azure.compute.getVirtualMachine({
 *     name: "example-vm",
 *     resourceGroupName: "example-resources",
 * });
 * const exampleMssql_virtualMachineVirtualMachine = new azure.mssql.VirtualMachine("exampleMssql/virtualMachineVirtualMachine", {
 *     virtualMachineId: exampleVirtualMachine.then(exampleVirtualMachine => exampleVirtualMachine.id),
 *     sqlLicenseType: "PAYG",
 *     rServicesEnabled: true,
 *     sqlConnectivityPort: 1433,
 *     sqlConnectivityType: "PRIVATE",
 *     sqlConnectivityUpdatePassword: "Password1234!",
 *     sqlConnectivityUpdateUsername: "sqllogin",
 *     autoPatching: {
 *         dayOfWeek: "Sunday",
 *         maintenanceWindowDurationInMinutes: 60,
 *         maintenanceWindowStartingHour: 2,
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Sql Virtual Machines can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:mssql/virtualMachine:VirtualMachine example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.SqlVirtualMachine/sqlVirtualMachines/example1
 * ```
 */
export class VirtualMachine extends pulumi.CustomResource {
    /**
     * Get an existing VirtualMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VirtualMachineState, opts?: pulumi.CustomResourceOptions): VirtualMachine {
        return new VirtualMachine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:mssql/virtualMachine:VirtualMachine';

    /**
     * Returns true if the given object is an instance of VirtualMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualMachine.__pulumiType;
    }

    /**
     * An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
     */
    public readonly autoBackup!: pulumi.Output<outputs.mssql.VirtualMachineAutoBackup | undefined>;
    /**
     * An `autoPatching` block as defined below.
     */
    public readonly autoPatching!: pulumi.Output<outputs.mssql.VirtualMachineAutoPatching | undefined>;
    /**
     * (Optional) An `keyVaultCredential` block as defined below.
     */
    public readonly keyVaultCredential!: pulumi.Output<outputs.mssql.VirtualMachineKeyVaultCredential | undefined>;
    /**
     * Should R Services be enabled?
     */
    public readonly rServicesEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * The SQL Server port. Defaults to `1433`.
     */
    public readonly sqlConnectivityPort!: pulumi.Output<number | undefined>;
    /**
     * The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
     */
    public readonly sqlConnectivityType!: pulumi.Output<string | undefined>;
    /**
     * The SQL Server sysadmin login password.
     */
    public readonly sqlConnectivityUpdatePassword!: pulumi.Output<string | undefined>;
    /**
     * The SQL Server sysadmin login to create.
     */
    public readonly sqlConnectivityUpdateUsername!: pulumi.Output<string | undefined>;
    /**
     * The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
     */
    public readonly sqlLicenseType!: pulumi.Output<string>;
    /**
     * An `storageConfiguration` block as defined below.
     */
    public readonly storageConfiguration!: pulumi.Output<outputs.mssql.VirtualMachineStorageConfiguration | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created.
     */
    public readonly virtualMachineId!: pulumi.Output<string>;

    /**
     * Create a VirtualMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualMachineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualMachineArgs | VirtualMachineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VirtualMachineState | undefined;
            inputs["autoBackup"] = state ? state.autoBackup : undefined;
            inputs["autoPatching"] = state ? state.autoPatching : undefined;
            inputs["keyVaultCredential"] = state ? state.keyVaultCredential : undefined;
            inputs["rServicesEnabled"] = state ? state.rServicesEnabled : undefined;
            inputs["sqlConnectivityPort"] = state ? state.sqlConnectivityPort : undefined;
            inputs["sqlConnectivityType"] = state ? state.sqlConnectivityType : undefined;
            inputs["sqlConnectivityUpdatePassword"] = state ? state.sqlConnectivityUpdatePassword : undefined;
            inputs["sqlConnectivityUpdateUsername"] = state ? state.sqlConnectivityUpdateUsername : undefined;
            inputs["sqlLicenseType"] = state ? state.sqlLicenseType : undefined;
            inputs["storageConfiguration"] = state ? state.storageConfiguration : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
        } else {
            const args = argsOrState as VirtualMachineArgs | undefined;
            if ((!args || args.sqlLicenseType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sqlLicenseType'");
            }
            if ((!args || args.virtualMachineId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'virtualMachineId'");
            }
            inputs["autoBackup"] = args ? args.autoBackup : undefined;
            inputs["autoPatching"] = args ? args.autoPatching : undefined;
            inputs["keyVaultCredential"] = args ? args.keyVaultCredential : undefined;
            inputs["rServicesEnabled"] = args ? args.rServicesEnabled : undefined;
            inputs["sqlConnectivityPort"] = args ? args.sqlConnectivityPort : undefined;
            inputs["sqlConnectivityType"] = args ? args.sqlConnectivityType : undefined;
            inputs["sqlConnectivityUpdatePassword"] = args ? args.sqlConnectivityUpdatePassword : undefined;
            inputs["sqlConnectivityUpdateUsername"] = args ? args.sqlConnectivityUpdateUsername : undefined;
            inputs["sqlLicenseType"] = args ? args.sqlLicenseType : undefined;
            inputs["storageConfiguration"] = args ? args.storageConfiguration : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualMachineId"] = args ? args.virtualMachineId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(VirtualMachine.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VirtualMachine resources.
 */
export interface VirtualMachineState {
    /**
     * An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
     */
    autoBackup?: pulumi.Input<inputs.mssql.VirtualMachineAutoBackup>;
    /**
     * An `autoPatching` block as defined below.
     */
    autoPatching?: pulumi.Input<inputs.mssql.VirtualMachineAutoPatching>;
    /**
     * (Optional) An `keyVaultCredential` block as defined below.
     */
    keyVaultCredential?: pulumi.Input<inputs.mssql.VirtualMachineKeyVaultCredential>;
    /**
     * Should R Services be enabled?
     */
    rServicesEnabled?: pulumi.Input<boolean>;
    /**
     * The SQL Server port. Defaults to `1433`.
     */
    sqlConnectivityPort?: pulumi.Input<number>;
    /**
     * The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
     */
    sqlConnectivityType?: pulumi.Input<string>;
    /**
     * The SQL Server sysadmin login password.
     */
    sqlConnectivityUpdatePassword?: pulumi.Input<string>;
    /**
     * The SQL Server sysadmin login to create.
     */
    sqlConnectivityUpdateUsername?: pulumi.Input<string>;
    /**
     * The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
     */
    sqlLicenseType?: pulumi.Input<string>;
    /**
     * An `storageConfiguration` block as defined below.
     */
    storageConfiguration?: pulumi.Input<inputs.mssql.VirtualMachineStorageConfiguration>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created.
     */
    virtualMachineId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VirtualMachine resource.
 */
export interface VirtualMachineArgs {
    /**
     * An `autoBackup` block as defined below. This block can be added to an existing resource, but removing this block forces a new resource to be created.
     */
    autoBackup?: pulumi.Input<inputs.mssql.VirtualMachineAutoBackup>;
    /**
     * An `autoPatching` block as defined below.
     */
    autoPatching?: pulumi.Input<inputs.mssql.VirtualMachineAutoPatching>;
    /**
     * (Optional) An `keyVaultCredential` block as defined below.
     */
    keyVaultCredential?: pulumi.Input<inputs.mssql.VirtualMachineKeyVaultCredential>;
    /**
     * Should R Services be enabled?
     */
    rServicesEnabled?: pulumi.Input<boolean>;
    /**
     * The SQL Server port. Defaults to `1433`.
     */
    sqlConnectivityPort?: pulumi.Input<number>;
    /**
     * The connectivity type used for this SQL Server. Defaults to `PRIVATE`.
     */
    sqlConnectivityType?: pulumi.Input<string>;
    /**
     * The SQL Server sysadmin login password.
     */
    sqlConnectivityUpdatePassword?: pulumi.Input<string>;
    /**
     * The SQL Server sysadmin login to create.
     */
    sqlConnectivityUpdateUsername?: pulumi.Input<string>;
    /**
     * The SQL Server license type. Possible values are `AHUB` (Azure Hybrid Benefit) and `PAYG` (Pay-As-You-Go). Changing this forces a new resource to be created.
     */
    sqlLicenseType: pulumi.Input<string>;
    /**
     * An `storageConfiguration` block as defined below.
     */
    storageConfiguration?: pulumi.Input<inputs.mssql.VirtualMachineStorageConfiguration>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the Virtual Machine. Changing this forces a new resource to be created.
     */
    virtualMachineId: pulumi.Input<string>;
}
