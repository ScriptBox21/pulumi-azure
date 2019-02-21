# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from .. import utilities, tables

class ConsumerGroup(pulumi.CustomResource):
    eventhub_endpoint_name: pulumi.Output[str]
    """
    The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
    """
    iothub_name: pulumi.Output[str]
    """
    The name of the IoT Hub. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of this Consumer Group. Changing this forces a new resource to be created.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
    """
    def __init__(__self__, resource_name, opts=None, eventhub_endpoint_name=None, iothub_name=None, name=None, resource_group_name=None, __name__=None, __opts__=None):
        """
        Manages a Consumer Group within an IotHub
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eventhub_endpoint_name: The name of the Event Hub-compatible endpoint in the IoT hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] iothub_name: The name of the IoT Hub. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of this Consumer Group. Changing this forces a new resource to be created.
        :param pulumi.Input[str] resource_group_name: The name of the resource group that contains the IoT hub. Changing this forces a new resource to be created.
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

        if eventhub_endpoint_name is None:
            raise TypeError('Missing required property eventhub_endpoint_name')
        __props__['eventhub_endpoint_name'] = eventhub_endpoint_name

        if iothub_name is None:
            raise TypeError('Missing required property iothub_name')
        __props__['iothub_name'] = iothub_name

        __props__['name'] = name

        if resource_group_name is None:
            raise TypeError('Missing required property resource_group_name')
        __props__['resource_group_name'] = resource_group_name

        super(ConsumerGroup, __self__).__init__(
            'azure:iot/consumerGroup:ConsumerGroup',
            resource_name,
            __props__,
            opts)


    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
