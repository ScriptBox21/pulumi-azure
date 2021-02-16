# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .dataset_azure_blob import *
from .dataset_cosmos_db_api import *
from .dataset_delimited_text import *
from .dataset_http import *
from .dataset_json import *
from .dataset_mysql import *
from .dataset_postgresql import *
from .dataset_sql_server_table import *
from .factory import *
from .get_factory import *
from .integration_runtime_managed import *
from .integration_runtime_rule import *
from .integration_runtime_self_hosted import *
from .integration_runtime_ssis import *
from .linked_service_azure_blob_storage import *
from .linked_service_azure_file_storage import *
from .linked_service_azure_function import *
from .linked_service_azure_sql_database import *
from .linked_service_azure_table_storage import *
from .linked_service_cosmos_db import *
from .linked_service_data_lake_storage_gen2 import *
from .linked_service_key_vault import *
from .linked_service_mysql import *
from .linked_service_postgresql import *
from .linked_service_sftp import *
from .linked_service_snowflake import *
from .linked_service_sql_server import *
from .linked_service_synapse import *
from .linked_service_web import *
from .pipeline import *
from .trigger_schedule import *
from ._inputs import *
from . import outputs

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "azure:datafactory/datasetAzureBlob:DatasetAzureBlob":
                return DatasetAzureBlob(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/datasetCosmosDBApi:DatasetCosmosDBApi":
                return DatasetCosmosDBApi(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/datasetDelimitedText:DatasetDelimitedText":
                return DatasetDelimitedText(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/datasetHttp:DatasetHttp":
                return DatasetHttp(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/datasetJson:DatasetJson":
                return DatasetJson(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/datasetMysql:DatasetMysql":
                return DatasetMysql(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/datasetPostgresql:DatasetPostgresql":
                return DatasetPostgresql(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/datasetSqlServerTable:DatasetSqlServerTable":
                return DatasetSqlServerTable(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/factory:Factory":
                return Factory(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/integrationRuntimeManaged:IntegrationRuntimeManaged":
                return IntegrationRuntimeManaged(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/integrationRuntimeRule:IntegrationRuntimeRule":
                return IntegrationRuntimeRule(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/integrationRuntimeSelfHosted:IntegrationRuntimeSelfHosted":
                return IntegrationRuntimeSelfHosted(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/integrationRuntimeSsis:IntegrationRuntimeSsis":
                return IntegrationRuntimeSsis(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceAzureBlobStorage:LinkedServiceAzureBlobStorage":
                return LinkedServiceAzureBlobStorage(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceAzureFileStorage:LinkedServiceAzureFileStorage":
                return LinkedServiceAzureFileStorage(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceAzureFunction:LinkedServiceAzureFunction":
                return LinkedServiceAzureFunction(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceAzureSqlDatabase:LinkedServiceAzureSqlDatabase":
                return LinkedServiceAzureSqlDatabase(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceAzureTableStorage:LinkedServiceAzureTableStorage":
                return LinkedServiceAzureTableStorage(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceCosmosDb:LinkedServiceCosmosDb":
                return LinkedServiceCosmosDb(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceDataLakeStorageGen2:LinkedServiceDataLakeStorageGen2":
                return LinkedServiceDataLakeStorageGen2(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceKeyVault:LinkedServiceKeyVault":
                return LinkedServiceKeyVault(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceMysql:LinkedServiceMysql":
                return LinkedServiceMysql(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServicePostgresql:LinkedServicePostgresql":
                return LinkedServicePostgresql(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceSftp:LinkedServiceSftp":
                return LinkedServiceSftp(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceSnowflake:LinkedServiceSnowflake":
                return LinkedServiceSnowflake(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceSqlServer:LinkedServiceSqlServer":
                return LinkedServiceSqlServer(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceSynapse:LinkedServiceSynapse":
                return LinkedServiceSynapse(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/linkedServiceWeb:LinkedServiceWeb":
                return LinkedServiceWeb(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/pipeline:Pipeline":
                return Pipeline(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "azure:datafactory/triggerSchedule:TriggerSchedule":
                return TriggerSchedule(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetAzureBlob", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetCosmosDBApi", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetDelimitedText", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetHttp", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetJson", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetMysql", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetPostgresql", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/datasetSqlServerTable", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/factory", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/integrationRuntimeManaged", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/integrationRuntimeRule", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/integrationRuntimeSelfHosted", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/integrationRuntimeSsis", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceAzureBlobStorage", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceAzureFileStorage", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceAzureFunction", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceAzureSqlDatabase", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceAzureTableStorage", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceCosmosDb", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceDataLakeStorageGen2", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceKeyVault", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceMysql", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServicePostgresql", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceSftp", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceSnowflake", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceSqlServer", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceSynapse", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/linkedServiceWeb", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/pipeline", _module_instance)
    pulumi.runtime.register_resource_module("azure", "datafactory/triggerSchedule", _module_instance)

_register_module()
