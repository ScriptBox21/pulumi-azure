# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

warnings.warn("azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue", DeprecationWarning)


class Queue(pulumi.CustomResource):
    auto_delete_on_idle: pulumi.Output[str]
    """
    The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
    """
    dead_lettering_on_message_expiration: pulumi.Output[bool]
    """
    Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
    """
    default_message_ttl: pulumi.Output[str]
    """
    The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
    """
    duplicate_detection_history_time_window: pulumi.Output[str]
    """
    The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
    """
    enable_batched_operations: pulumi.Output[bool]
    """
    Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
    """
    enable_express: pulumi.Output[bool]
    """
    Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
    """
    enable_partitioning: pulumi.Output[bool]
    """
    Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
    """
    forward_dead_lettered_messages_to: pulumi.Output[str]
    """
    The name of a Queue or Topic to automatically forward dead lettered messages to.
    """
    forward_to: pulumi.Output[str]
    """
    The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
    """
    lock_duration: pulumi.Output[str]
    """
    The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
    """
    max_delivery_count: pulumi.Output[float]
    """
    Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
    """
    max_size_in_megabytes: pulumi.Output[float]
    """
    Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
    """
    name: pulumi.Output[str]
    """
    Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
    """
    namespace_name: pulumi.Output[str]
    """
    The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
    """
    requires_duplicate_detection: pulumi.Output[bool]
    """
    Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
    """
    requires_session: pulumi.Output[bool]
    """
    Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
    """
    status: pulumi.Output[str]
    """
    The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
    """
    warnings.warn("azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue", DeprecationWarning)

    def __init__(__self__, resource_name, opts=None, auto_delete_on_idle=None, dead_lettering_on_message_expiration=None, default_message_ttl=None, duplicate_detection_history_time_window=None, enable_batched_operations=None, enable_express=None, enable_partitioning=None, forward_dead_lettered_messages_to=None, forward_to=None, lock_duration=None, max_delivery_count=None, max_size_in_megabytes=None, name=None, namespace_name=None, requires_duplicate_detection=None, requires_session=None, resource_group_name=None, status=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a ServiceBus Queue.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_namespace = azure.servicebus.Namespace("exampleNamespace",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            sku="Standard",
            tags={
                "source": "example",
            })
        example_queue = azure.servicebus.Queue("exampleQueue",
            resource_group_name=example_resource_group.name,
            namespace_name=example_namespace.name,
            enable_partitioning=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_delete_on_idle: The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
        :param pulumi.Input[bool] dead_lettering_on_message_expiration: Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        :param pulumi.Input[str] default_message_ttl: The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
        :param pulumi.Input[str] duplicate_detection_history_time_window: The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
        :param pulumi.Input[bool] enable_batched_operations: Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
        :param pulumi.Input[bool] enable_express: Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
        :param pulumi.Input[bool] enable_partitioning: Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
        :param pulumi.Input[str] forward_dead_lettered_messages_to: The name of a Queue or Topic to automatically forward dead lettered messages to.
        :param pulumi.Input[str] forward_to: The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
        :param pulumi.Input[str] lock_duration: The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
        :param pulumi.Input[float] max_delivery_count: Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
        :param pulumi.Input[float] max_size_in_megabytes: Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_duplicate_detection: Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[bool] requires_session: Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
        """
        pulumi.log.warn("Queue is deprecated: azure.eventhub.Queue has been deprecated in favor of azure.servicebus.Queue")
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

            __props__['auto_delete_on_idle'] = auto_delete_on_idle
            __props__['dead_lettering_on_message_expiration'] = dead_lettering_on_message_expiration
            __props__['default_message_ttl'] = default_message_ttl
            __props__['duplicate_detection_history_time_window'] = duplicate_detection_history_time_window
            __props__['enable_batched_operations'] = enable_batched_operations
            __props__['enable_express'] = enable_express
            __props__['enable_partitioning'] = enable_partitioning
            __props__['forward_dead_lettered_messages_to'] = forward_dead_lettered_messages_to
            __props__['forward_to'] = forward_to
            __props__['lock_duration'] = lock_duration
            __props__['max_delivery_count'] = max_delivery_count
            __props__['max_size_in_megabytes'] = max_size_in_megabytes
            __props__['name'] = name
            if namespace_name is None:
                raise TypeError("Missing required property 'namespace_name'")
            __props__['namespace_name'] = namespace_name
            __props__['requires_duplicate_detection'] = requires_duplicate_detection
            __props__['requires_session'] = requires_session
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['status'] = status
        super(Queue, __self__).__init__(
            'azure:eventhub/queue:Queue',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, auto_delete_on_idle=None, dead_lettering_on_message_expiration=None, default_message_ttl=None, duplicate_detection_history_time_window=None, enable_batched_operations=None, enable_express=None, enable_partitioning=None, forward_dead_lettered_messages_to=None, forward_to=None, lock_duration=None, max_delivery_count=None, max_size_in_megabytes=None, name=None, namespace_name=None, requires_duplicate_detection=None, requires_session=None, resource_group_name=None, status=None):
        """
        Get an existing Queue resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auto_delete_on_idle: The ISO 8601 timespan duration of the idle interval after which the Queue is automatically deleted, minimum of 5 minutes.
        :param pulumi.Input[bool] dead_lettering_on_message_expiration: Boolean flag which controls whether the Queue has dead letter support when a message expires. Defaults to `false`.
        :param pulumi.Input[str] default_message_ttl: The ISO 8601 timespan duration of the TTL of messages sent to this queue. This is the default value used when TTL is not set on message itself.
        :param pulumi.Input[str] duplicate_detection_history_time_window: The ISO 8601 timespan duration during which duplicates can be detected. Defaults to 10 minutes (`PT10M`).
        :param pulumi.Input[bool] enable_batched_operations: Boolean flag which controls whether server-side batched operations are enabled. Defaults to `true`.
        :param pulumi.Input[bool] enable_express: Boolean flag which controls whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `false`.
        :param pulumi.Input[bool] enable_partitioning: Boolean flag which controls whether to enable the queue to be partitioned across multiple message brokers. Changing this forces a new resource to be created. Defaults to `false` for Basic and Standard. For Premium, it MUST be set to `true`.
        :param pulumi.Input[str] forward_dead_lettered_messages_to: The name of a Queue or Topic to automatically forward dead lettered messages to.
        :param pulumi.Input[str] forward_to: The name of a Queue or Topic to automatically forward messages to. Please [see the documentation](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-auto-forwarding) for more information.
        :param pulumi.Input[str] lock_duration: The ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. Maximum value is 5 minutes. Defaults to 1 minute (`PT1M`).
        :param pulumi.Input[float] max_delivery_count: Integer value which controls when a message is automatically dead lettered. Defaults to `10`.
        :param pulumi.Input[float] max_size_in_megabytes: Integer value which controls the size of memory allocated for the queue. For supported values see the "Queue or topic size" section of [Service Bus Quotas](https://docs.microsoft.com/en-us/azure/service-bus-messaging/service-bus-quotas). Defaults to `1024`.
        :param pulumi.Input[str] name: Specifies the name of the ServiceBus Queue resource. Changing this forces a new resource to be created.
        :param pulumi.Input[str] namespace_name: The name of the ServiceBus Namespace to create this queue in. Changing this forces a new resource to be created.
        :param pulumi.Input[bool] requires_duplicate_detection: Boolean flag which controls whether the Queue requires duplicate detection. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[bool] requires_session: Boolean flag which controls whether the Queue requires sessions. This will allow ordered handling of unbounded sequences of related messages. With sessions enabled a queue can guarantee first-in-first-out delivery of messages. Changing this forces a new resource to be created. Defaults to `false`.
        :param pulumi.Input[str] resource_group_name: The name of the resource group in which to create the namespace. Changing this forces a new resource to be created.
        :param pulumi.Input[str] status: The status of the Queue. Possible values are `Active`, `Creating`, `Deleting`, `Disabled`, `ReceiveDisabled`, `Renaming`, `SendDisabled`, `Unknown`. Note that `Restoring` is not accepted. Defaults to `Active`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["auto_delete_on_idle"] = auto_delete_on_idle
        __props__["dead_lettering_on_message_expiration"] = dead_lettering_on_message_expiration
        __props__["default_message_ttl"] = default_message_ttl
        __props__["duplicate_detection_history_time_window"] = duplicate_detection_history_time_window
        __props__["enable_batched_operations"] = enable_batched_operations
        __props__["enable_express"] = enable_express
        __props__["enable_partitioning"] = enable_partitioning
        __props__["forward_dead_lettered_messages_to"] = forward_dead_lettered_messages_to
        __props__["forward_to"] = forward_to
        __props__["lock_duration"] = lock_duration
        __props__["max_delivery_count"] = max_delivery_count
        __props__["max_size_in_megabytes"] = max_size_in_megabytes
        __props__["name"] = name
        __props__["namespace_name"] = namespace_name
        __props__["requires_duplicate_detection"] = requires_duplicate_detection
        __props__["requires_session"] = requires_session
        __props__["resource_group_name"] = resource_group_name
        __props__["status"] = status
        return Queue(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
