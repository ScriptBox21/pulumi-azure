# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime

class GetClientConfigResult(object):
    """
    A collection of values returned by getClientConfig.
    """
    def __init__(__self__, client_id=None, service_principal_application_id=None, service_principal_object_id=None, subscription_id=None, tenant_id=None, id=None):
        if client_id and not isinstance(client_id, basestring):
            raise TypeError('Expected argument client_id to be a basestring')
        __self__.client_id = client_id
        if service_principal_application_id and not isinstance(service_principal_application_id, basestring):
            raise TypeError('Expected argument service_principal_application_id to be a basestring')
        __self__.service_principal_application_id = service_principal_application_id
        if service_principal_object_id and not isinstance(service_principal_object_id, basestring):
            raise TypeError('Expected argument service_principal_object_id to be a basestring')
        __self__.service_principal_object_id = service_principal_object_id
        if subscription_id and not isinstance(subscription_id, basestring):
            raise TypeError('Expected argument subscription_id to be a basestring')
        __self__.subscription_id = subscription_id
        if tenant_id and not isinstance(tenant_id, basestring):
            raise TypeError('Expected argument tenant_id to be a basestring')
        __self__.tenant_id = tenant_id
        if id and not isinstance(id, basestring):
            raise TypeError('Expected argument id to be a basestring')
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """

def get_client_config():
    """
    Use this data source to access the configuration of the Azure Resource Manager
    provider.
    """
    __args__ = dict()

    __ret__ = pulumi.runtime.invoke('azure:core/getClientConfig:getClientConfig', __args__)

    return GetClientConfigResult(
        client_id=__ret__.get('clientId'),
        service_principal_application_id=__ret__.get('servicePrincipalApplicationId'),
        service_principal_object_id=__ret__.get('servicePrincipalObjectId'),
        subscription_id=__ret__.get('subscriptionId'),
        tenant_id=__ret__.get('tenantId'),
        id=__ret__.get('id'))
