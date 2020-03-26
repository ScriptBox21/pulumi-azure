# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class OutputServicebusTopic(pulumi.CustomResource):
    name: pulumi.Output[str]
    """
    The name of the Stream Output. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
    """
    serialization: pulumi.Output[dict]
    """
    A `serialization` block as defined below.

      * `encoding` (`str`) - The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
      * `fieldDelimiter` (`str`) - The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
      * `format` (`str`) - Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
      * `type` (`str`) - The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
    """
    servicebus_namespace: pulumi.Output[str]
    """
    The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
    """
    shared_access_policy_key: pulumi.Output[str]
    """
    The shared access policy key for the specified shared access policy.
    """
    shared_access_policy_name: pulumi.Output[str]
    """
    The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
    """
    stream_analytics_job_name: pulumi.Output[str]
    """
    The name of the Stream Analytics Job. Changing this forces a new resource to be created.
    """
    topic_name: pulumi.Output[str]
    """
    The name of the Service Bus Topic.
    """
    def __init__(__self__, resource_name, opts=None, name=None, resource_group_name=None, serialization=None, servicebus_namespace=None, shared_access_policy_key=None, shared_access_policy_name=None, stream_analytics_job_name=None, topic_name=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Stream Analytics Output to a ServiceBus Topic.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/stream_analytics_output_servicebus_topic.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the Service Bus Topic.

        The **serialization** object supports the following:

          * `encoding` (`pulumi.Input[str]`) - The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
          * `fieldDelimiter` (`pulumi.Input[str]`) - The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
          * `format` (`pulumi.Input[str]`) - Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
          * `type` (`pulumi.Input[str]`) - The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
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

            __props__['name'] = name
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            if serialization is None:
                raise TypeError("Missing required property 'serialization'")
            __props__['serialization'] = serialization
            if servicebus_namespace is None:
                raise TypeError("Missing required property 'servicebus_namespace'")
            __props__['servicebus_namespace'] = servicebus_namespace
            if shared_access_policy_key is None:
                raise TypeError("Missing required property 'shared_access_policy_key'")
            __props__['shared_access_policy_key'] = shared_access_policy_key
            if shared_access_policy_name is None:
                raise TypeError("Missing required property 'shared_access_policy_name'")
            __props__['shared_access_policy_name'] = shared_access_policy_name
            if stream_analytics_job_name is None:
                raise TypeError("Missing required property 'stream_analytics_job_name'")
            __props__['stream_analytics_job_name'] = stream_analytics_job_name
            if topic_name is None:
                raise TypeError("Missing required property 'topic_name'")
            __props__['topic_name'] = topic_name
        super(OutputServicebusTopic, __self__).__init__(
            'azure:streamanalytics/outputServicebusTopic:OutputServicebusTopic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, name=None, resource_group_name=None, serialization=None, servicebus_namespace=None, shared_access_policy_key=None, shared_access_policy_name=None, stream_analytics_job_name=None, topic_name=None):
        """
        Get an existing OutputServicebusTopic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the Stream Output. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Stream Analytics Job exists. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] serialization: A `serialization` block as defined below.
        :param pulumi.Input[str] servicebus_namespace: The namespace that is associated with the desired Event Hub, Service Bus Topic, Service Bus Topic, etc.
        :param pulumi.Input[str] shared_access_policy_key: The shared access policy key for the specified shared access policy.
        :param pulumi.Input[str] shared_access_policy_name: The shared access policy name for the Event Hub, Service Bus Queue, Service Bus Topic, etc.
        :param pulumi.Input[str] stream_analytics_job_name: The name of the Stream Analytics Job. Changing this forces a new resource to be created.
        :param pulumi.Input[str] topic_name: The name of the Service Bus Topic.

        The **serialization** object supports the following:

          * `encoding` (`pulumi.Input[str]`) - The encoding of the incoming data in the case of input and the encoding of outgoing data in the case of output. It currently can only be set to `UTF8`.
          * `fieldDelimiter` (`pulumi.Input[str]`) - The delimiter that will be used to separate comma-separated value (CSV) records. Possible values are ` ` (space), `,` (comma), `   ` (tab), `|` (pipe) and `;`.
          * `format` (`pulumi.Input[str]`) - Specifies the format of the JSON the output will be written in. Possible values are `Array` and `LineSeparated`.
          * `type` (`pulumi.Input[str]`) - The serialization format used for outgoing data streams. Possible values are `Avro`, `Csv` and `Json`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["resource_group_name"] = resource_group_name
        __props__["serialization"] = serialization
        __props__["servicebus_namespace"] = servicebus_namespace
        __props__["shared_access_policy_key"] = shared_access_policy_key
        __props__["shared_access_policy_name"] = shared_access_policy_name
        __props__["stream_analytics_job_name"] = stream_analytics_job_name
        __props__["topic_name"] = topic_name
        return OutputServicebusTopic(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

