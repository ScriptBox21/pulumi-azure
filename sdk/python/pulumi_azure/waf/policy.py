# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Policy(pulumi.CustomResource):
    custom_rules: pulumi.Output[list]
    """
    One or more `custom_rule` blocks as defined below.

      * `action` (`str`) - Type of Actions
      * `matchConditions` (`list`) - One or more `match_condition` block defined below.
        * `matchValues` (`list`) - Match value
        * `matchVariables` (`list`) - One or more `match_variable` block defined below.
          * `selector` (`str`) - Describes field of the matchVariable collection
          * `variableName` (`str`) - The name of the Match Variable

        * `negationCondition` (`bool`) - Describes if this is negate condition or not
        * `operator` (`str`) - Describes operator to be matched

      * `name` (`str`) - The name of the policy. Changing this forces a new resource to be created.
      * `priority` (`float`) - Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value
      * `ruleType` (`str`) - Describes the type of rule
    """
    location: pulumi.Output[str]
    """
    Resource location. Changing this forces a new resource to be created.
    """
    name: pulumi.Output[str]
    """
    The name of the policy. Changing this forces a new resource to be created.
    """
    policy_settings: pulumi.Output[dict]
    """
    A `policy_setting` block as defined below.

      * `enabled` (`bool`) - Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
      * `mode` (`str`) - Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
    """
    resource_group_name: pulumi.Output[str]
    """
    The name of the resource group. Changing this forces a new resource to be created.
    """
    tags: pulumi.Output[dict]
    """
    A mapping of tags to assign to the Web Application Firewall Policy.
    """
    def __init__(__self__, resource_name, opts=None, custom_rules=None, location=None, name=None, policy_settings=None, resource_group_name=None, tags=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages a Azure Web Application Firewall Policy instance.



        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] custom_rules: One or more `custom_rule` blocks as defined below.
        :param pulumi.Input[str] location: Resource location. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] policy_settings: A `policy_setting` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Web Application Firewall Policy.

        The **custom_rules** object supports the following:

          * `action` (`pulumi.Input[str]`) - Type of Actions
          * `matchConditions` (`pulumi.Input[list]`) - One or more `match_condition` block defined below.
            * `matchValues` (`pulumi.Input[list]`) - Match value
            * `matchVariables` (`pulumi.Input[list]`) - One or more `match_variable` block defined below.
              * `selector` (`pulumi.Input[str]`) - Describes field of the matchVariable collection
              * `variableName` (`pulumi.Input[str]`) - The name of the Match Variable

            * `negationCondition` (`pulumi.Input[bool]`) - Describes if this is negate condition or not
            * `operator` (`pulumi.Input[str]`) - Describes operator to be matched

          * `name` (`pulumi.Input[str]`) - The name of the policy. Changing this forces a new resource to be created.
          * `priority` (`pulumi.Input[float]`) - Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value
          * `ruleType` (`pulumi.Input[str]`) - Describes the type of rule

        The **policy_settings** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
          * `mode` (`pulumi.Input[str]`) - Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
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

            __props__['custom_rules'] = custom_rules
            __props__['location'] = location
            __props__['name'] = name
            __props__['policy_settings'] = policy_settings
            if resource_group_name is None:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
            __props__['tags'] = tags
        super(Policy, __self__).__init__(
            'azure:waf/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, custom_rules=None, location=None, name=None, policy_settings=None, resource_group_name=None, tags=None):
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] custom_rules: One or more `custom_rule` blocks as defined below.
        :param pulumi.Input[str] location: Resource location. Changing this forces a new resource to be created.
        :param pulumi.Input[str] name: The name of the policy. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] policy_settings: A `policy_setting` block as defined below.
        :param pulumi.Input[str] resource_group_name: The name of the resource group. Changing this forces a new resource to be created.
        :param pulumi.Input[dict] tags: A mapping of tags to assign to the Web Application Firewall Policy.

        The **custom_rules** object supports the following:

          * `action` (`pulumi.Input[str]`) - Type of Actions
          * `matchConditions` (`pulumi.Input[list]`) - One or more `match_condition` block defined below.
            * `matchValues` (`pulumi.Input[list]`) - Match value
            * `matchVariables` (`pulumi.Input[list]`) - One or more `match_variable` block defined below.
              * `selector` (`pulumi.Input[str]`) - Describes field of the matchVariable collection
              * `variableName` (`pulumi.Input[str]`) - The name of the Match Variable

            * `negationCondition` (`pulumi.Input[bool]`) - Describes if this is negate condition or not
            * `operator` (`pulumi.Input[str]`) - Describes operator to be matched

          * `name` (`pulumi.Input[str]`) - The name of the policy. Changing this forces a new resource to be created.
          * `priority` (`pulumi.Input[float]`) - Describes priority of the rule. Rules with a lower value will be evaluated before rules with a higher value
          * `ruleType` (`pulumi.Input[str]`) - Describes the type of rule

        The **policy_settings** object supports the following:

          * `enabled` (`pulumi.Input[bool]`) - Describes if the policy is in enabled state or disabled state Defaults to `Enabled`.
          * `mode` (`pulumi.Input[str]`) - Describes if it is in detection mode  or prevention mode at the policy level Defaults to `Prevention`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["custom_rules"] = custom_rules
        __props__["location"] = location
        __props__["name"] = name
        __props__["policy_settings"] = policy_settings
        __props__["resource_group_name"] = resource_group_name
        __props__["tags"] = tags
        return Policy(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

