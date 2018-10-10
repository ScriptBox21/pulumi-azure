# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import pulumi
import pulumi.runtime
from .. import utilities

class Slot(pulumi.CustomResource):
    """
    Manages an App Service Slot (within an App Service).
    
    -> **Note:** When using Slots - the `app_settings`, `connection_string` and `site_config` blocks on the `azurerm_app_service` resource will be overwritten when promoting a Slot using the `azurerm_app_service_active_slot` resource.
    
    """
    def __init__(__self__, __name__, __opts__=None, app_service_name=None, app_service_plan_id=None, app_settings=None, client_affinity_enabled=None, connection_strings=None, enabled=None, https_only=None, identity=None, location=None, name=None, resource_group_name=None, site_config=None, tags=None):
        """Create a Slot resource with the given unique name, props, and options."""
        if not __name__:
            raise TypeError('Missing resource name argument (for URN creation)')
        if not isinstance(__name__, basestring):
            raise TypeError('Expected resource name to be a string')
        if __opts__ and not isinstance(__opts__, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')

        __props__ = dict()

        if not app_service_name:
            raise TypeError('Missing required property app_service_name')
        elif not isinstance(app_service_name, basestring):
            raise TypeError('Expected property app_service_name to be a basestring')
        __self__.app_service_name = app_service_name
        """
        The name of the App Service within which to create the App Service Slot.  Changing this forces a new resource to be created.
        """
        __props__['appServiceName'] = app_service_name

        if not app_service_plan_id:
            raise TypeError('Missing required property app_service_plan_id')
        elif not isinstance(app_service_plan_id, basestring):
            raise TypeError('Expected property app_service_plan_id to be a basestring')
        __self__.app_service_plan_id = app_service_plan_id
        """
        The ID of the App Service Plan within which to create this App Service Slot. Changing this forces a new resource to be created.
        """
        __props__['appServicePlanId'] = app_service_plan_id

        if app_settings and not isinstance(app_settings, dict):
            raise TypeError('Expected property app_settings to be a dict')
        __self__.app_settings = app_settings
        """
        A key-value pair of App Settings.
        """
        __props__['appSettings'] = app_settings

        if client_affinity_enabled and not isinstance(client_affinity_enabled, bool):
            raise TypeError('Expected property client_affinity_enabled to be a bool')
        __self__.client_affinity_enabled = client_affinity_enabled
        """
        Should the App Service Slot send session affinity cookies, which route client requests in the same session to the same instance?
        """
        __props__['clientAffinityEnabled'] = client_affinity_enabled

        if connection_strings and not isinstance(connection_strings, list):
            raise TypeError('Expected property connection_strings to be a list')
        __self__.connection_strings = connection_strings
        """
        An `connection_string` block as defined below.
        """
        __props__['connectionStrings'] = connection_strings

        if enabled and not isinstance(enabled, bool):
            raise TypeError('Expected property enabled to be a bool')
        __self__.enabled = enabled
        """
        Is the App Service Slot Enabled?
        """
        __props__['enabled'] = enabled

        if https_only and not isinstance(https_only, bool):
            raise TypeError('Expected property https_only to be a bool')
        __self__.https_only = https_only
        """
        Can the App Service Slot only be accessed via HTTPS? Defaults to `false`.
        """
        __props__['httpsOnly'] = https_only

        if identity and not isinstance(identity, dict):
            raise TypeError('Expected property identity to be a dict')
        __self__.identity = identity
        """
        A Managed Service Identity block as defined below.
        """
        __props__['identity'] = identity

        if not location:
            raise TypeError('Missing required property location')
        elif not isinstance(location, basestring):
            raise TypeError('Expected property location to be a basestring')
        __self__.location = location
        """
        Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
        """
        __props__['location'] = location

        if name and not isinstance(name, basestring):
            raise TypeError('Expected property name to be a basestring')
        __self__.name = name
        """
        The name of the Connection String.
        """
        __props__['name'] = name

        if not resource_group_name:
            raise TypeError('Missing required property resource_group_name')
        elif not isinstance(resource_group_name, basestring):
            raise TypeError('Expected property resource_group_name to be a basestring')
        __self__.resource_group_name = resource_group_name
        """
        The name of the resource group in which to create the App Service Slot component.
        """
        __props__['resourceGroupName'] = resource_group_name

        if site_config and not isinstance(site_config, dict):
            raise TypeError('Expected property site_config to be a dict')
        __self__.site_config = site_config
        """
        A `site_config` object as defined below.
        """
        __props__['siteConfig'] = site_config

        if tags and not isinstance(tags, dict):
            raise TypeError('Expected property tags to be a dict')
        __self__.tags = tags
        """
        A mapping of tags to assign to the resource.
        """
        __props__['tags'] = tags

        __self__.default_site_hostname = pulumi.runtime.UNKNOWN
        """
        The Default Hostname associated with the App Service Slot - such as `mysite.azurewebsites.net`
        """

        super(Slot, __self__).__init__(
            'azure:appservice/slot:Slot',
            __name__,
            __props__,
            __opts__)

    def set_outputs(self, outs):
        if 'appServiceName' in outs:
            self.app_service_name = outs['appServiceName']
        if 'appServicePlanId' in outs:
            self.app_service_plan_id = outs['appServicePlanId']
        if 'appSettings' in outs:
            self.app_settings = outs['appSettings']
        if 'clientAffinityEnabled' in outs:
            self.client_affinity_enabled = outs['clientAffinityEnabled']
        if 'connectionStrings' in outs:
            self.connection_strings = outs['connectionStrings']
        if 'defaultSiteHostname' in outs:
            self.default_site_hostname = outs['defaultSiteHostname']
        if 'enabled' in outs:
            self.enabled = outs['enabled']
        if 'httpsOnly' in outs:
            self.https_only = outs['httpsOnly']
        if 'identity' in outs:
            self.identity = outs['identity']
        if 'location' in outs:
            self.location = outs['location']
        if 'name' in outs:
            self.name = outs['name']
        if 'resourceGroupName' in outs:
            self.resource_group_name = outs['resourceGroupName']
        if 'siteConfig' in outs:
            self.site_config = outs['siteConfig']
        if 'tags' in outs:
            self.tags = outs['tags']