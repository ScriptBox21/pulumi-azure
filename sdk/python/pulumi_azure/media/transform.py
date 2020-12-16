# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['Transform']


class Transform(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 media_services_account_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransformOutputArgs']]]]] = None,
                 resource_group_name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Transform.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_service_account = azure.media.ServiceAccount("exampleServiceAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_accounts=[azure.media.ServiceAccountStorageAccountArgs(
                id=example_account.id,
                is_primary=True,
            )])
        example_transform = azure.media.Transform("exampleTransform",
            resource_group_name=example_resource_group.name,
            media_services_account_name=example_service_account.name,
            description="My transform description",
            outputs=[azure.media.TransformOutputArgs(
                relative_priority="Normal",
                on_error_action="ContinueJob",
                builtin_preset=azure.media.TransformOutputBuiltinPresetArgs(
                    preset_name="AACGoodQualityAudio",
                ),
            )])
        ```
        ### With Multiple Outputs

        ```python
        import pulumi
        import pulumi_azure as azure

        example_resource_group = azure.core.ResourceGroup("exampleResourceGroup", location="West Europe")
        example_account = azure.storage.Account("exampleAccount",
            resource_group_name=example_resource_group.name,
            location=example_resource_group.location,
            account_tier="Standard",
            account_replication_type="GRS")
        example_service_account = azure.media.ServiceAccount("exampleServiceAccount",
            location=example_resource_group.location,
            resource_group_name=example_resource_group.name,
            storage_accounts=[azure.media.ServiceAccountStorageAccountArgs(
                id=example_account.id,
                is_primary=True,
            )])
        example_transform = azure.media.Transform("exampleTransform",
            resource_group_name=example_resource_group.name,
            media_services_account_name=example_service_account.name,
            description="My transform description",
            outputs=[
                azure.media.TransformOutputArgs(
                    relative_priority="Normal",
                    on_error_action="ContinueJob",
                    builtin_preset=azure.media.TransformOutputBuiltinPresetArgs(
                        preset_name="AACGoodQualityAudio",
                    ),
                ),
                azure.media.TransformOutputArgs(
                    relative_priority="Low",
                    on_error_action="ContinueJob",
                    audio_analyzer_preset=azure.media.TransformOutputAudioAnalyzerPresetArgs(
                        audio_language="en-US",
                        audio_analysis_mode="Basic",
                    ),
                ),
                azure.media.TransformOutputArgs(
                    relative_priority="Low",
                    on_error_action="StopProcessingJob",
                    face_detector_preset=azure.media.TransformOutputFaceDetectorPresetArgs(
                        analysis_resolution="StandardDefinition",
                    ),
                ),
            ])
        ```

        ## Import

        Transforms can be imported using the `resource id`, e.g.

        ```sh
         $ pulumi import azure:media/transform:Transform example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Media/mediaservices/media1/transforms/transform1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional verbose description of the Transform.
        :param pulumi.Input[str] media_services_account_name: The Media Services account name. Changing this forces a new Transform to be created.
        :param pulumi.Input[str] name: The name which should be used for this Transform. Changing this forces a new Transform to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransformOutputArgs']]]] outputs: One or more `output` blocks as defined below. At least one `output` must be defined.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Transform should exist. Changing this forces a new Transform to be created.
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
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['description'] = description
            if media_services_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'media_services_account_name'")
            __props__['media_services_account_name'] = media_services_account_name
            __props__['name'] = name
            __props__['outputs'] = outputs
            if resource_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'resource_group_name'")
            __props__['resource_group_name'] = resource_group_name
        super(Transform, __self__).__init__(
            'azure:media/transform:Transform',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            media_services_account_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            outputs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransformOutputArgs']]]]] = None,
            resource_group_name: Optional[pulumi.Input[str]] = None) -> 'Transform':
        """
        Get an existing Transform resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional verbose description of the Transform.
        :param pulumi.Input[str] media_services_account_name: The Media Services account name. Changing this forces a new Transform to be created.
        :param pulumi.Input[str] name: The name which should be used for this Transform. Changing this forces a new Transform to be created.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TransformOutputArgs']]]] outputs: One or more `output` blocks as defined below. At least one `output` must be defined.
        :param pulumi.Input[str] resource_group_name: The name of the Resource Group where the Transform should exist. Changing this forces a new Transform to be created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["description"] = description
        __props__["media_services_account_name"] = media_services_account_name
        __props__["name"] = name
        __props__["outputs"] = outputs
        __props__["resource_group_name"] = resource_group_name
        return Transform(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional verbose description of the Transform.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="mediaServicesAccountName")
    def media_services_account_name(self) -> pulumi.Output[str]:
        """
        The Media Services account name. Changing this forces a new Transform to be created.
        """
        return pulumi.get(self, "media_services_account_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name which should be used for this Transform. Changing this forces a new Transform to be created.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def outputs(self) -> pulumi.Output[Optional[Sequence['outputs.TransformOutput']]]:
        """
        One or more `output` blocks as defined below. At least one `output` must be defined.
        """
        return pulumi.get(self, "outputs")

    @property
    @pulumi.getter(name="resourceGroupName")
    def resource_group_name(self) -> pulumi.Output[str]:
        """
        The name of the Resource Group where the Transform should exist. Changing this forces a new Transform to be created.
        """
        return pulumi.get(self, "resource_group_name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

