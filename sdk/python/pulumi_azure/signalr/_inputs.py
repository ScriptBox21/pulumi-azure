# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'ServiceCorArgs',
    'ServiceFeatureArgs',
    'ServiceSkuArgs',
]

@pulumi.input_type
class ServiceCorArgs:
    def __init__(__self__, *,
                 allowed_origins: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_origins: A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
        """
        pulumi.set(__self__, "allowed_origins", allowed_origins)

    @property
    @pulumi.getter(name="allowedOrigins")
    def allowed_origins(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of origins which should be able to make cross-origin calls. `*` can be used to allow all calls.
        """
        return pulumi.get(self, "allowed_origins")

    @allowed_origins.setter
    def allowed_origins(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_origins", value)


@pulumi.input_type
class ServiceFeatureArgs:
    def __init__(__self__, *,
                 flag: pulumi.Input[str],
                 value: pulumi.Input[str]):
        """
        :param pulumi.Input[str] flag: The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
        :param pulumi.Input[str] value: A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
        """
        pulumi.set(__self__, "flag", flag)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def flag(self) -> pulumi.Input[str]:
        """
        The kind of Feature. Possible values are `EnableConnectivityLogs`, `EnableMessagingLogs`, and `ServiceMode`.
        """
        return pulumi.get(self, "flag")

    @flag.setter
    def flag(self, value: pulumi.Input[str]):
        pulumi.set(self, "flag", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        """
        A value of a feature flag. Possible values are `Classic`, `Default` and `Serverless`.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ServiceSkuArgs:
    def __init__(__self__, *,
                 capacity: pulumi.Input[float],
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[float] capacity: Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
        :param pulumi.Input[str] name: Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
        """
        pulumi.set(__self__, "capacity", capacity)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def capacity(self) -> pulumi.Input[float]:
        """
        Specifies the number of units associated with this SignalR service. Valid values are `1`, `2`, `5`, `10`, `20`, `50` and `100`.
        """
        return pulumi.get(self, "capacity")

    @capacity.setter
    def capacity(self, value: pulumi.Input[float]):
        pulumi.set(self, "capacity", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Specifies which tier to use. Valid values are `Free_F1` and `Standard_S1`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


