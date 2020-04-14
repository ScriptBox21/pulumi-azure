# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetNamespaceAuthorizationRuleResult:
    """
    A collection of values returned by getNamespaceAuthorizationRule.
    """
    def __init__(__self__, id=None, name=None, namespace_name=None, primary_connection_string=None, primary_key=None, resource_group_name=None, secondary_connection_string=None, secondary_key=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if namespace_name and not isinstance(namespace_name, str):
            raise TypeError("Expected argument 'namespace_name' to be a str")
        __self__.namespace_name = namespace_name
        if primary_connection_string and not isinstance(primary_connection_string, str):
            raise TypeError("Expected argument 'primary_connection_string' to be a str")
        __self__.primary_connection_string = primary_connection_string
        """
        The primary connection string for the authorization rule.
        """
        if primary_key and not isinstance(primary_key, str):
            raise TypeError("Expected argument 'primary_key' to be a str")
        __self__.primary_key = primary_key
        """
        The primary access key for the authorization rule.
        """
        if resource_group_name and not isinstance(resource_group_name, str):
            raise TypeError("Expected argument 'resource_group_name' to be a str")
        __self__.resource_group_name = resource_group_name
        if secondary_connection_string and not isinstance(secondary_connection_string, str):
            raise TypeError("Expected argument 'secondary_connection_string' to be a str")
        __self__.secondary_connection_string = secondary_connection_string
        """
        The secondary connection string for the authorization rule.
        """
        if secondary_key and not isinstance(secondary_key, str):
            raise TypeError("Expected argument 'secondary_key' to be a str")
        __self__.secondary_key = secondary_key
        """
        The secondary access key for the authorization rule.
        """
class AwaitableGetNamespaceAuthorizationRuleResult(GetNamespaceAuthorizationRuleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNamespaceAuthorizationRuleResult(
            id=self.id,
            name=self.name,
            namespace_name=self.namespace_name,
            primary_connection_string=self.primary_connection_string,
            primary_key=self.primary_key,
            resource_group_name=self.resource_group_name,
            secondary_connection_string=self.secondary_connection_string,
            secondary_key=self.secondary_key)

def get_namespace_authorization_rule(name=None,namespace_name=None,resource_group_name=None,opts=None):
    """
    Use this data source to access information about an existing ServiceBus Namespace Authorization Rule.




    :param str name: Specifies the name of the ServiceBus Namespace Authorization Rule.
    :param str namespace_name: Specifies the name of the ServiceBus Namespace.
    :param str resource_group_name: Specifies the name of the Resource Group where the ServiceBus Namespace exists.
    """
    __args__ = dict()


    __args__['name'] = name
    __args__['namespaceName'] = namespace_name
    __args__['resourceGroupName'] = resource_group_name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('azure:servicebus/getNamespaceAuthorizationRule:getNamespaceAuthorizationRule', __args__, opts=opts).value

    return AwaitableGetNamespaceAuthorizationRuleResult(
        id=__ret__.get('id'),
        name=__ret__.get('name'),
        namespace_name=__ret__.get('namespaceName'),
        primary_connection_string=__ret__.get('primaryConnectionString'),
        primary_key=__ret__.get('primaryKey'),
        resource_group_name=__ret__.get('resourceGroupName'),
        secondary_connection_string=__ret__.get('secondaryConnectionString'),
        secondary_key=__ret__.get('secondaryKey'))
