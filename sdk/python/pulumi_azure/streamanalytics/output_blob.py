# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class OutputBlob(pulumi.CustomResource):
    date_format: pulumi.Output[str]
    """
    The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
    """
    name: pulumi.Output[str]
    """
    The name of the Stream Output. Changing this forces a new resource to be created.
    """
    path_pattern: pulumi.Output[str]
    """
    The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
    """
    serialization: pulumi.Output[dict]
    """
    A `serialization` block as defined below.
    """
    storage_account_key: pulumi.Output[str]
    """
    The Access Key which should be used to connect to this Storage Account.
    """
    storage_account_name: pulumi.Output[str]
    """
    The name of the Storage Account.
    """
    storage_container_name: pulumi.Output[str]
    """
    The name of the Container within the Storage Account.
    """
    stream_analytics_job_name: pulumi.Output[str]
    """
    The name of the Stream Analytics Job. Changing this forces a new resource to be created.
    """
    time_format: pulumi.Output[str]
    """
    The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
    """
    def __init__(__self__, resource_name, opts=None, date_format=None, name=None, path_pattern=None, resource_group_name=None, serialization=None, storage_account_key=None, storage_account_name=None, storage_container_name=None, stream_analytics_job_name=None, time_format=None, __name__=None, __opts__=None):
        """
        Manages a Stream Analytics Output to Blob Storage.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] date_format: The date format. Wherever `{date}` appears in `path_pattern`, the value of this property is used as the date format instead.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] path_pattern: The blob path pattern. Not a regular expression. It represents a pattern against which blob names will be matched to determine whether or not they should be included as input or output to the job.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] storage_account_key: The Access Key which should be used to connect to this Storage Account.
        :param pulumi.Input[str] storage_account_name: The name of the Storage Account.
        :param pulumi.Input[str] storage_container_name: The name of the Container within the Storage Account.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] time_format: The time format. Wherever `{time}` appears in `path_pattern`, the value of this property is used as the time format instead.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if not resource_name:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(resource_name, str):
            raise TypeError('Expected resource name to be a string')
        if opts and not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if date_format is None:
            raise TypeError("Missing required property 'date_format'")
        __props__['date_format'] = date_format

        __props__['name'] = name

        if path_pattern is None:
            raise TypeError("Missing required property 'path_pattern'")
        __props__['path_pattern'] = path_pattern

        if resource_group_name is None:
            raise TypeError("Missing required property 'resource_group_name'")
        __props__['resource_group_name'] = resource_group_name

        if serialization is None:
            raise TypeError("Missing required property 'serialization'")
        __props__['serialization'] = serialization

        if storage_account_key is None:
            raise TypeError("Missing required property 'storage_account_key'")
        __props__['storage_account_key'] = storage_account_key

        if storage_account_name is None:
            raise TypeError("Missing required property 'storage_account_name'")
        __props__['storage_account_name'] = storage_account_name

        if storage_container_name is None:
            raise TypeError("Missing required property 'storage_container_name'")
        __props__['storage_container_name'] = storage_container_name

        if stream_analytics_job_name is None:
            raise TypeError("Missing required property 'stream_analytics_job_name'")
        __props__['stream_analytics_job_name'] = stream_analytics_job_name

        if time_format is None:
            raise TypeError("Missing required property 'time_format'")
        __props__['time_format'] = time_format

        super(OutputBlob, __self__).__init__(
            'azure:streamanalytics/outputBlob:OutputBlob',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

