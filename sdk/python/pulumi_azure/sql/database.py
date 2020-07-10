# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class Database(pulumi.CustomResource):
    collation: pulumi.Output[str]
    """
    The name of the collation. Applies only if `create_mode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
    """
    create_mode: pulumi.Output[str]
    """
    Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
    """
    creation_date: pulumi.Output[str]
    """
    The creation date of the SQL Database.
    """
    default_secondary_location: pulumi.Output[str]
    """
    The default secondary location of the SQL Database.
    """
    edition: pulumi.Output[str]
    """
    The edition of the database to be created. Applies only if `create_mode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
    """
    elastic_pool_name: pulumi.Output[str]
    """
    The name of the elastic database pool.
    """
    encryption: pulumi.Output[str]
    extended_auditing_policy: pulumi.Output[dict]
    """
    A `extended_auditing_policy` block as defined below.

      * `retention_in_days` (`float`) - Specifies the number of days to retain logs for in the storage account.
      * `storage_account_access_key` (`str`) - Specifies the access key to use for the auditing storage account.
      * `storageAccountAccessKeyIsSecondary` (`bool`) - Specifies whether `storage_account_access_key` value is the storage's secondary key.
      * `storage_endpoint` (`str`) - Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).
    """
    import_: pulumi.Output[dict]
    """
    A Database Import block as documented below. `create_mode` must be set to `Default`.

      * `administrator_login` (`str`) - Specifies the name of the SQL administrator.
      * `administrator_login_password` (`str`) - Specifies the password of the SQL administrator.
      * `authentication_type` (`str`) - Specifies the type of authentication used to access the server. Valid values are `SQL` or `ADPassword`.
      * `operationMode` (`str`) - Specifies the type of import operation being performed. The only allowable value is `Import`.
      * `storageKey` (`str`) - Specifies the access key for the storage account.
      * `storageKeyType` (`str`) - Specifies the type of access key for the storage account. Valid values are `StorageAccessKey` or `SharedAccessKey`.
      * `storageUri` (`str`) - Specifies the blob URI of the .bacpac file.
    """
    location: pulumi.Output[str]
    """
    Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
    """
    max_size_bytes: pulumi.Output[str]
    """
    The maximum size that the database can grow to. Applies only if `create_mode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
    """
    max_size_gb: pulumi.Output[str]
    name: pulumi.Output[str]
    """
    The name of the database.
    """
    read_scale: pulumi.Output[bool]
    """
    Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
    """
    requested_service_objective_id: pulumi.Output[str]
    """
    A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
    .
    """
    requested_service_objective_name: pulumi.Output[str]
    """
    The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
    """
    restore_point_in_time: pulumi.Output[str]
    """
    The point in time for the restore. Only applies if `create_mode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
    """
    server_name: pulumi.Output[str]
    """
    The name of the SQL Server on which to create the database.
    """
    source_database_deletion_date: pulumi.Output[str]
    """
    The deletion date time of the source database. Only applies to deleted databases where `create_mode` is `PointInTimeRestore`.
    """
    source_database_id: pulumi.Output[str]
    """
    The URI of the source database if `create_mode` value is not `Default`.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the resource.
    """
    threat_detection_policy: pulumi.Output[dict]
    """
    Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.

      * `disabled_alerts` (`list`) - Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
      * `email_account_admins` (`str`) - Should the account administrators be emailed when this alert is triggered?
      * `email_addresses` (`list`) - A list of email addresses which alerts should be sent to.
      * `retention_days` (`float`) - Specifies the number of days to keep in the Threat Detection audit logs.
      * `state` (`str`) - The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
      * `storage_account_access_key` (`str`) - Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
      * `storage_endpoint` (`str`) - Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
      * `useServerDefault` (`str`) - Should the default server policy be used? Defaults to `Disabled`.
    """
    zone_redundant: pulumi.Output[bool]
    """
    Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.
    """
    def __init__(__self__, resource_name, opts=None, collation=None, create_mode=None, edition=None, elastic_pool_name=None, extended_auditing_policy=None, import_=None, location=None, max_size_bytes=None, max_size_gb=None, name=None, read_scale=None, requested_service_objective_id=None, requested_service_objective_name=None, resource_group_name=None, restore_point_in_time=None, server_name=None, source_database_deletion_date=None, source_database_id=None, tags=None, threat_detection_policy=None, zone_redundant=None, __props__=None, __name__=None, __opts__=None):
        """
        Allows you to manage an Azure SQL Database

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West US")
        example_sql_server = azure.sql.SqlServer("exampleSqlServer",
            resource_group_name=example_resource_group.name,
            location="West US",
            version="12.0",
            administrator_login="4dm1n157r470r",
            administrator_login_password="4-v3ry-53cr37-p455w0rd",
            tags={
                "environment": "production",
            })
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="LRS")
        example_database = azure.sql.Database("exampleDatabase",
            resource_group_name=example_resource_group.name,
            location="West US",
            server_name=example_sql_server.name,
            extended_auditing_policy={
                "storage_endpoint": example_account.primary_blob_endpoint,
                "storage_account_access_key": example_account.primary_access_key,
                "storageAccountAccessKeyIsSecondary": True,
                "retention_in_days": 6,
            },
            tags={
                "environment": "production",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collation: The name of the collation. Applies only if `create_mode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_mode: Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
        :param pulumi.Input[str] edition: The edition of the database to be created. Applies only if `create_mode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] elastic_pool_name: The name of the elastic database pool.
        :param pulumi.Input[dict] extended_auditing_policy: A `extended_auditing_policy` block as defined below.
        :param pulumi.Input[dict] import_: A Database Import block as documented below. `create_mode` must be set to `Default`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] max_size_bytes: The maximum size that the database can grow to. Applies only if `create_mode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] name: The name of the database.
        :param pulumi.Input[bool] read_scale: Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
        :param pulumi.Input[str] requested_service_objective_id: A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
               .
        :param pulumi.Input[str] requested_service_objective_name: The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
        :param pulumi.Input[str] restore_point_in_time: The point in time for the restore. Only applies if `create_mode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
        :param pulumi.Input[str] server_name: The name of the SQL Server on which to create the database.
        :param pulumi.Input[str] source_database_deletion_date: The deletion date time of the source database. Only applies to deleted databases where `create_mode` is `PointInTimeRestore`.
        :param pulumi.Input[str] source_database_id: The URI of the source database if `create_mode` value is not `Default`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] threat_detection_policy: Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        :param pulumi.Input[bool] zone_redundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.

        The **extended_auditing_policy** object supports the following:

          * `retention_in_days` (`pulumi.Input[float]`) - Specifies the number of days to retain logs for in the storage account.
          * `storage_account_access_key` (`pulumi.Input[str]`) - Specifies the access key to use for the auditing storage account.
          * `storageAccountAccessKeyIsSecondary` (`pulumi.Input[bool]`) - Specifies whether `storage_account_access_key` value is the storage's secondary key.
          * `storage_endpoint` (`pulumi.Input[str]`) - Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).

        The **import_** object supports the following:

          * `administrator_login` (`pulumi.Input[str]`) - Specifies the name of the SQL administrator.
          * `administrator_login_password` (`pulumi.Input[str]`) - Specifies the password of the SQL administrator.
          * `authentication_type` (`pulumi.Input[str]`) - Specifies the type of authentication used to access the server. Valid values are `SQL` or `ADPassword`.
          * `operationMode` (`pulumi.Input[str]`) - Specifies the type of import operation being performed. The only allowable value is `Import`.
          * `storageKey` (`pulumi.Input[str]`) - Specifies the access key for the storage account.
          * `storageKeyType` (`pulumi.Input[str]`) - Specifies the type of access key for the storage account. Valid values are `StorageAccessKey` or `SharedAccessKey`.
          * `storageUri` (`pulumi.Input[str]`) - Specifies the blob URI of the .bacpac file.

        The **threat_detection_policy** object supports the following:

          * `disabled_alerts` (`pulumi.Input[list]`) - Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
          * `email_account_admins` (`pulumi.Input[str]`) - Should the account administrators be emailed when this alert is triggered?
          * `email_addresses` (`pulumi.Input[list]`) - A list of email addresses which alerts should be sent to.
          * `retention_days` (`pulumi.Input[float]`) - Specifies the number of days to keep in the Threat Detection audit logs.
          * `state` (`pulumi.Input[str]`) - The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
          * `storage_account_access_key` (`pulumi.Input[str]`) - Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
          * `storage_endpoint` (`pulumi.Input[str]`) - Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
          * `useServerDefault` (`pulumi.Input[str]`) - Should the default server policy be used? Defaults to `Disabled`.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['collation'] = collation
            __props__['create_mode'] = create_mode
            __props__['edition'] = edition
            __props__['elastic_pool_name'] = elastic_pool_name
            __props__['extended_auditing_policy'] = extended_auditing_policy
            __props__['import_'] = import_
            __props__['location'] = location
            __props__['max_size_bytes'] = max_size_bytes
            __props__['max_size_gb'] = max_size_gb
            __props__['name'] = name
            __props__['read_scale'] = read_scale
            __props__['requested_service_objective_id'] = requested_service_objective_id
            __props__['requested_service_objective_name'] = requested_service_objective_name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['restore_point_in_time'] = restore_point_in_time
            if server_name is None:
                raise TypeError("Missing required property 'server_name'")
            __props__['server_name'] = server_name
            __props__['source_database_deletion_date'] = source_database_deletion_date
            __props__['source_database_id'] = source_database_id
            __props__['tags'] = tags
            __props__['threat_detection_policy'] = threat_detection_policy
            __props__['zone_redundant'] = zone_redundant
            __props__['creation_date'] = None
            __props__['default_secondary_location'] = None
            __props__['encryption'] = None
        super(Database, __self__).__init__(
            'azure:sql/database:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, collation=None, create_mode=None, creation_date=None, default_secondary_location=None, edition=None, elastic_pool_name=None, encryption=None, extended_auditing_policy=None, import_=None, location=None, max_size_bytes=None, max_size_gb=None, name=None, read_scale=None, requested_service_objective_id=None, requested_service_objective_name=None, resource_group_name=None, restore_point_in_time=None, server_name=None, source_database_deletion_date=None, source_database_id=None, tags=None, threat_detection_policy=None, zone_redundant=None):
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] collation: The name of the collation. Applies only if `create_mode` is `Default`.  Azure default is `SQL_LATIN1_GENERAL_CP1_CI_AS`. Changing this forces a new resource to be created.
        :param pulumi.Input[str] create_mode: Specifies how to create the database. Valid values are: `Default`, `Copy`, `OnlineSecondary`, `NonReadableSecondary`,  `PointInTimeRestore`, `Recovery`, `Restore` or `RestoreLongTermRetentionBackup`. Must be `Default` to create a new database. Defaults to `Default`. Please see [Azure SQL Database REST API](https://docs.microsoft.com/en-us/rest/api/sql/databases/createorupdate#createmode)
        :param pulumi.Input[str] creation_date: The creation date of the SQL Database.
        :param pulumi.Input[str] default_secondary_location: The default secondary location of the SQL Database.
        :param pulumi.Input[str] edition: The edition of the database to be created. Applies only if `create_mode` is `Default`. Valid values are: `Basic`, `Standard`, `Premium`, `DataWarehouse`, `Business`, `BusinessCritical`, `Free`, `GeneralPurpose`, `Hyperscale`, `Premium`, `PremiumRS`, `Standard`, `Stretch`, `System`, `System2`, or `Web`. Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] elastic_pool_name: The name of the elastic database pool.
        :param pulumi.Input[dict] extended_auditing_policy: A `extended_auditing_policy` block as defined below.
        :param pulumi.Input[dict] import_: A Database Import block as documented below. `create_mode` must be set to `Default`.
        :param pulumi.Input[str] location: Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        :param pulumi.Input[str] max_size_bytes: The maximum size that the database can grow to. Applies only if `create_mode` is `Default`.  Please see [Azure SQL Database Service Tiers](https://azure.microsoft.com/en-gb/documentation/articles/sql-database-service-tiers/).
        :param pulumi.Input[str] name: The name of the database.
        :param pulumi.Input[bool] read_scale: Read-only connections will be redirected to a high-available replica. Please see [Use read-only replicas to load-balance read-only query workloads](https://docs.microsoft.com/en-us/azure/sql-database/sql-database-read-scale-out).
        :param pulumi.Input[str] requested_service_objective_id: A GUID/UUID corresponding to a configured Service Level Objective for the Azure SQL database which can be used to configure a performance level.
               .
        :param pulumi.Input[str] requested_service_objective_name: The service objective name for the database. Valid values depend on edition and location and may include `S0`, `S1`, `S2`, `S3`, `P1`, `P2`, `P4`, `P6`, `P11` and `ElasticPool`. You can list the available names with the cli: ```shell az sql db list-editions -l westus -o table ```. For further information please see [Azure CLI - az sql db](https://docs.microsoft.com/en-us/cli/azure/sql/db?view=azure-cli-latest#az-sql-db-list-editions).
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the database.  This must be the same as Database Server resource group currently.
        :param pulumi.Input[str] restore_point_in_time: The point in time for the restore. Only applies if `create_mode` is `PointInTimeRestore` e.g. 2013-11-08T22:00:40Z
        :param pulumi.Input[str] server_name: The name of the SQL Server on which to create the database.
        :param pulumi.Input[str] source_database_deletion_date: The deletion date time of the source database. Only applies to deleted databases where `create_mode` is `PointInTimeRestore`.
        :param pulumi.Input[str] source_database_id: The URI of the source database if `create_mode` value is not `Default`.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the resource.
        :param pulumi.Input[dict] threat_detection_policy: Threat detection policy configuration. The `threat_detection_policy` block supports fields documented below.
        :param pulumi.Input[bool] zone_redundant: Whether or not this database is zone redundant, which means the replicas of this database will be spread across multiple availability zones.

        The **extended_auditing_policy** object supports the following:

          * `retention_in_days` (`pulumi.Input[float]`) - Specifies the number of days to retain logs for in the storage account.
          * `storage_account_access_key` (`pulumi.Input[str]`) - Specifies the access key to use for the auditing storage account.
          * `storageAccountAccessKeyIsSecondary` (`pulumi.Input[bool]`) - Specifies whether `storage_account_access_key` value is the storage's secondary key.
          * `storage_endpoint` (`pulumi.Input[str]`) - Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net).

        The **import_** object supports the following:

          * `administrator_login` (`pulumi.Input[str]`) - Specifies the name of the SQL administrator.
          * `administrator_login_password` (`pulumi.Input[str]`) - Specifies the password of the SQL administrator.
          * `authentication_type` (`pulumi.Input[str]`) - Specifies the type of authentication used to access the server. Valid values are `SQL` or `ADPassword`.
          * `operationMode` (`pulumi.Input[str]`) - Specifies the type of import operation being performed. The only allowable value is `Import`.
          * `storageKey` (`pulumi.Input[str]`) - Specifies the access key for the storage account.
          * `storageKeyType` (`pulumi.Input[str]`) - Specifies the type of access key for the storage account. Valid values are `StorageAccessKey` or `SharedAccessKey`.
          * `storageUri` (`pulumi.Input[str]`) - Specifies the blob URI of the .bacpac file.

        The **threat_detection_policy** object supports the following:

          * `disabled_alerts` (`pulumi.Input[list]`) - Specifies a list of alerts which should be disabled. Possible values include `Access_Anomaly`, `Sql_Injection` and `Sql_Injection_Vulnerability`.
          * `email_account_admins` (`pulumi.Input[str]`) - Should the account administrators be emailed when this alert is triggered?
          * `email_addresses` (`pulumi.Input[list]`) - A list of email addresses which alerts should be sent to.
          * `retention_days` (`pulumi.Input[float]`) - Specifies the number of days to keep in the Threat Detection audit logs.
          * `state` (`pulumi.Input[str]`) - The State of the Policy. Possible values are `Enabled`, `Disabled` or `New`.
          * `storage_account_access_key` (`pulumi.Input[str]`) - Specifies the identifier key of the Threat Detection audit storage account. Required if `state` is `Enabled`.
          * `storage_endpoint` (`pulumi.Input[str]`) - Specifies the blob storage endpoint (e.g. https://MyAccount.blob.core.windows.net). This blob storage will hold all Threat Detection audit logs. Required if `state` is `Enabled`.
          * `useServerDefault` (`pulumi.Input[str]`) - Should the default server policy be used? Defaults to `Disabled`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["collation"] = collation
        __props__["create_mode"] = create_mode
        __props__["creation_date"] = creation_date
        __props__["default_secondary_location"] = default_secondary_location
        __props__["edition"] = edition
        __props__["elastic_pool_name"] = elastic_pool_name
        __props__["encryption"] = encryption
        __props__["extended_auditing_policy"] = extended_auditing_policy
        __props__["import_"] = import_
        __props__["location"] = location
        __props__["max_size_bytes"] = max_size_bytes
        __props__["max_size_gb"] = max_size_gb
        __props__["name"] = name
        __props__["read_scale"] = read_scale
        __props__["requested_service_objective_id"] = requested_service_objective_id
        __props__["requested_service_objective_name"] = requested_service_objective_name
        __props__["resource_group_name"] = resource_group_name
        __props__["restore_point_in_time"] = restore_point_in_time
        __props__["server_name"] = server_name
        __props__["source_database_deletion_date"] = source_database_deletion_date
        __props__["source_database_id"] = source_database_id
        __props__["tags"] = tags
        __props__["threat_detection_policy"] = threat_detection_policy
        __props__["zone_redundant"] = zone_redundant
        return Database(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
