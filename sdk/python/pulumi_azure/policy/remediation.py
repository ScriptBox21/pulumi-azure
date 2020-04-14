# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Remediation(pulumi.CustomResource):
    location_filters: pulumi.Output[list]
    """
    A list of the resource locations that will be remediated.
    """
    name: pulumi.Output[str]
    """
    The name of the Policy Remediation. Changing this forces a new resource to be created.
    """
    policy_assignment_id: pulumi.Output[str]
    """
    The resource ID of the policy assignment that should be remediated.
    """
    policy_definition_reference_id: pulumi.Output[str]
    """
    The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
    """
    scope: pulumi.Output[str]
    """
    The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
    """
    def __init__(__self__, resource_name, opts=None, location_filters=None, name=None, policy_assignment_id=None, policy_definition_reference_id=None, scope=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an Azure Policy Remediation at the specified Scope.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] location_filters: A list of the resource locations that will be remediated.
        :param pulumi.Input[str] name: The name of the Policy Remediation. Changing this forces a new resource to be created.
        :param pulumi.Input[str] policy_assignment_id: The resource ID of the policy assignment that should be remediated.
        :param pulumi.Input[str] policy_definition_reference_id: The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        :param pulumi.Input[str] scope: The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
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

            __props__['location_filters'] = location_filters
            __props__['name'] = name
            if policy_assignment_id is None:
                raise TypeError("Missing required property 'policy_assignment_id'")
            __props__['policy_assignment_id'] = policy_assignment_id
            __props__['policy_definition_reference_id'] = policy_definition_reference_id
            if scope is None:
                raise TypeError("Missing required property 'scope'")
            __props__['scope'] = scope
        super(Remediation, __self__).__init__(
            'azure:policy/remediation:Remediation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, location_filters=None, name=None, policy_assignment_id=None, policy_definition_reference_id=None, scope=None):
        """
        Get an existing Remediation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] location_filters: A list of the resource locations that will be remediated.
        :param pulumi.Input[str] name: The name of the Policy Remediation. Changing this forces a new resource to be created.
        :param pulumi.Input[str] policy_assignment_id: The resource ID of the policy assignment that should be remediated.
        :param pulumi.Input[str] policy_definition_reference_id: The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
        :param pulumi.Input[str] scope: The Scope at which the Policy Remediation should be applied. Changing this forces a new resource to be created. A scope must be a Resource ID out of one of the following list:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["location_filters"] = location_filters
        __props__["name"] = name
        __props__["policy_assignment_id"] = policy_assignment_id
        __props__["policy_definition_reference_id"] = policy_definition_reference_id
        __props__["scope"] = scope
        return Remediation(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

