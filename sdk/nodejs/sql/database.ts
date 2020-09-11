// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Allows you to manage an Azure SQL Database
 *
 * > **NOTE:** The Database Extended Auditing Policy Can be set inline here as well as with the mssqlDatabaseExtendedAuditingPolicy resource resource. You can only use one or the other and using both will cause a conflict.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West US"});
 * const exampleSqlServer = new azure.sql.SqlServer("exampleSqlServer", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: "West US",
 *     version: "12.0",
 *     administratorLogin: "4dm1n157r470r",
 *     administratorLoginPassword: "4-v3ry-53cr37-p455w0rd",
 *     tags: {
 *         environment: "production",
 *     },
 * });
 * const exampleAccount = new azure.storage.Account("exampleAccount", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     accountTier: "Standard",
 *     accountReplicationType: "LRS",
 * });
 * const exampleDatabase = new azure.sql.Database("exampleDatabase", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: "West US",
 *     serverName: exampleSqlServer.name,
 *     extendedAuditingPolicy: {
 *         storageEndpoint: exampleAccount.primaryBlobEndpoint,
 *         storageAccountAccessKey: exampleAccount.primaryAccessKey,
 *         storageAccountAccessKeyIsSecondary: true,
 *         retentionInDays: 6,
 *     },
 *     tags: {
 *         environment: "production",
 *     },
 * });
 * ```
 */
export class Database extends pulumi.CustomResource {
    /**
     * Get an existing Database resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseState, opts?: pulumi.CustomResourceOptions): Database {
        return new Database(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:sql/database:Database';

    /**
     * Returns true if the given object is an instance of Database.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Database {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Database.__pulumiType;
    }

    /**
     * The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
     */
    public readonly collation!: pulumi.Output<string>;
    /**
     * Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
     */
    public readonly createMode!: pulumi.Output<string | undefined>;
    /**
     * The creation date of the SQL Database.
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * The default secondary location of the SQL Database.
     */
    public /*out*/ readonly defaultSecondaryLocation!: pulumi.Output<string>;
    /**
     * The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
     */
    public readonly edition!: pulumi.Output<string>;
    /**
     * The name of the elastic database pool.
     */
    public readonly elasticPoolName!: pulumi.Output<string>;
    public /*out*/ readonly encryption!: pulumi.Output<string>;
    /**
     * A `extendedAuditingPolicy` block as defined below.
     *
     * @deprecated the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
     */
    public readonly extendedAuditingPolicy!: pulumi.Output<outputs.sql.DatabaseExtendedAuditingPolicy>;
    /**
     * A Database Import block as documented below. `createMode` must be set to `Default`.
     */
    public readonly import!: pulumi.Output<outputs.sql.DatabaseImport | undefined>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
     */
    public readonly maxSizeBytes!: pulumi.Output<string>;
    public readonly maxSizeGb!: pulumi.Output<string>;
    /**
     * The name of the database.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
     */
    public readonly readScale!: pulumi.Output<boolean | undefined>;
    /**
     * A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
     * .
     */
    public readonly requestedServiceObjectiveId!: pulumi.Output<string>;
    /**
     * The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
     */
    public readonly requestedServiceObjectiveName!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
     */
    public readonly restorePointInTime!: pulumi.Output<string>;
    /**
     * The name of the SQL Server on which to create the database.
     */
    public readonly serverName!: pulumi.Output<string>;
    /**
     * The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
     */
    public readonly sourceDatabaseDeletionDate!: pulumi.Output<string>;
    /**
     * The URI of the source database if `createMode` value is not `Default`.
     */
    public readonly sourceDatabaseId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
     */
    public readonly threatDetectionPolicy!: pulumi.Output<outputs.sql.DatabaseThreatDetectionPolicy>;
    /**
     * Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    public readonly zoneRedundant!: pulumi.Output<boolean | undefined>;

    /**
     * Create a Database resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatabaseArgs | DatabaseState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DatabaseState | undefined;
            inputs["collation"] = state ? state.collation : undefined;
            inputs["createMode"] = state ? state.createMode : undefined;
            inputs["creationDate"] = state ? state.creationDate : undefined;
            inputs["defaultSecondaryLocation"] = state ? state.defaultSecondaryLocation : undefined;
            inputs["edition"] = state ? state.edition : undefined;
            inputs["elasticPoolName"] = state ? state.elasticPoolName : undefined;
            inputs["encryption"] = state ? state.encryption : undefined;
            inputs["extendedAuditingPolicy"] = state ? state.extendedAuditingPolicy : undefined;
            inputs["import"] = state ? state.import : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maxSizeBytes"] = state ? state.maxSizeBytes : undefined;
            inputs["maxSizeGb"] = state ? state.maxSizeGb : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["readScale"] = state ? state.readScale : undefined;
            inputs["requestedServiceObjectiveId"] = state ? state.requestedServiceObjectiveId : undefined;
            inputs["requestedServiceObjectiveName"] = state ? state.requestedServiceObjectiveName : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["restorePointInTime"] = state ? state.restorePointInTime : undefined;
            inputs["serverName"] = state ? state.serverName : undefined;
            inputs["sourceDatabaseDeletionDate"] = state ? state.sourceDatabaseDeletionDate : undefined;
            inputs["sourceDatabaseId"] = state ? state.sourceDatabaseId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["threatDetectionPolicy"] = state ? state.threatDetectionPolicy : undefined;
            inputs["zoneRedundant"] = state ? state.zoneRedundant : undefined;
        } else {
            const args = argsOrState as DatabaseArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.serverName === undefined) {
                throw new Error("Missing required property 'serverName'");
            }
            inputs["collation"] = args ? args.collation : undefined;
            inputs["createMode"] = args ? args.createMode : undefined;
            inputs["edition"] = args ? args.edition : undefined;
            inputs["elasticPoolName"] = args ? args.elasticPoolName : undefined;
            inputs["extendedAuditingPolicy"] = args ? args.extendedAuditingPolicy : undefined;
            inputs["import"] = args ? args.import : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxSizeBytes"] = args ? args.maxSizeBytes : undefined;
            inputs["maxSizeGb"] = args ? args.maxSizeGb : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["readScale"] = args ? args.readScale : undefined;
            inputs["requestedServiceObjectiveId"] = args ? args.requestedServiceObjectiveId : undefined;
            inputs["requestedServiceObjectiveName"] = args ? args.requestedServiceObjectiveName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["restorePointInTime"] = args ? args.restorePointInTime : undefined;
            inputs["serverName"] = args ? args.serverName : undefined;
            inputs["sourceDatabaseDeletionDate"] = args ? args.sourceDatabaseDeletionDate : undefined;
            inputs["sourceDatabaseId"] = args ? args.sourceDatabaseId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["threatDetectionPolicy"] = args ? args.threatDetectionPolicy : undefined;
            inputs["zoneRedundant"] = args ? args.zoneRedundant : undefined;
            inputs["creationDate"] = undefined /*out*/;
            inputs["defaultSecondaryLocation"] = undefined /*out*/;
            inputs["encryption"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Database.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Database resources.
 */
export interface DatabaseState {
    /**
     * The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
     */
    readonly collation?: pulumi.Input<string>;
    /**
     * Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
     */
    readonly createMode?: pulumi.Input<string>;
    /**
     * The creation date of the SQL Database.
     */
    readonly creationDate?: pulumi.Input<string>;
    /**
     * The default secondary location of the SQL Database.
     */
    readonly defaultSecondaryLocation?: pulumi.Input<string>;
    /**
     * The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
     */
    readonly edition?: pulumi.Input<string>;
    /**
     * The name of the elastic database pool.
     */
    readonly elasticPoolName?: pulumi.Input<string>;
    readonly encryption?: pulumi.Input<string>;
    /**
     * A `extendedAuditingPolicy` block as defined below.
     *
     * @deprecated the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
     */
    readonly extendedAuditingPolicy?: pulumi.Input<inputs.sql.DatabaseExtendedAuditingPolicy>;
    /**
     * A Database Import block as documented below. `createMode` must be set to `Default`.
     */
    readonly import?: pulumi.Input<inputs.sql.DatabaseImport>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
     */
    readonly maxSizeBytes?: pulumi.Input<string>;
    readonly maxSizeGb?: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
     */
    readonly readScale?: pulumi.Input<boolean>;
    /**
     * A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
     * .
     */
    readonly requestedServiceObjectiveId?: pulumi.Input<string>;
    /**
     * The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
     */
    readonly requestedServiceObjectiveName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
     */
    readonly restorePointInTime?: pulumi.Input<string>;
    /**
     * The name of the SQL Server on which to create the database.
     */
    readonly serverName?: pulumi.Input<string>;
    /**
     * The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
     */
    readonly sourceDatabaseDeletionDate?: pulumi.Input<string>;
    /**
     * The URI of the source database if `createMode` value is not `Default`.
     */
    readonly sourceDatabaseId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
     */
    readonly threatDetectionPolicy?: pulumi.Input<inputs.sql.DatabaseThreatDetectionPolicy>;
    /**
     * Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    readonly zoneRedundant?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Database resource.
 */
export interface DatabaseArgs {
    /**
     * The name of the collation. Applies only if `createMode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
     */
    readonly collation?: pulumi.Input<string>;
    /**
     * Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
     */
    readonly createMode?: pulumi.Input<string>;
    /**
     * The edition of the database to be created. Applies only if `createMode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
     */
    readonly edition?: pulumi.Input<string>;
    /**
     * The name of the elastic database pool.
     */
    readonly elasticPoolName?: pulumi.Input<string>;
    /**
     * A `extendedAuditingPolicy` block as defined below.
     *
     * @deprecated the `extended_auditing_policy` block has been moved to `azurerm_mssql_server_extended_auditing_policy` and `azurerm_mssql_database_extended_auditing_policy`. This block will be removed in version 3.0 of the provider.
     */
    readonly extendedAuditingPolicy?: pulumi.Input<inputs.sql.DatabaseExtendedAuditingPolicy>;
    /**
     * A Database Import block as documented below. `createMode` must be set to `Default`.
     */
    readonly import?: pulumi.Input<inputs.sql.DatabaseImport>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The maximum size that the database can grow to. Applies only if `createMode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
     */
    readonly maxSizeBytes?: pulumi.Input<string>;
    readonly maxSizeGb?: pulumi.Input<string>;
    /**
     * The name of the database.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
     */
    readonly readScale?: pulumi.Input<boolean>;
    /**
     * A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
     * .
     */
    readonly requestedServiceObjectiveId?: pulumi.Input<string>;
    /**
     * The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
     */
    readonly requestedServiceObjectiveName?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The point in time for the restore. Only applies if `createMode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
     */
    readonly restorePointInTime?: pulumi.Input<string>;
    /**
     * The name of the SQL Server on which to create the database.
     */
    readonly serverName: pulumi.Input<string>;
    /**
     * The deletion date time of the source database. Only applies to deleted databases where `createMode` is `PointInTimeRestore`.
     */
    readonly sourceDatabaseDeletionDate?: pulumi.Input<string>;
    /**
     * The URI of the source database if `createMode` value is not `Default`.
     */
    readonly sourceDatabaseId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Threat detection policy configuration. The `threatDetectionPolicy` block supports fields documented below.
     */
    readonly threatDetectionPolicy?: pulumi.Input<inputs.sql.DatabaseThreatDetectionPolicy>;
    /**
     * Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
     */
    readonly zoneRedundant?: pulumi.Input<boolean>;
}
