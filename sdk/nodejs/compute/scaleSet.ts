// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a virtual machine scale set.
 * 
 * ## Disclaimers
 * 
 * > **Note:** The `azure.compute.ScaleSet` resource has been superseded by the `azure.compute.LinuxVirtualMachineScaleSet` and `azure.compute.WindowsVirtualMachineScaleSet` resources. The existing `azure.compute.ScaleSet` resource will continue to be available throughout the 2.x releases however is in a feature-frozen state to maintain compatibility - new functionality will instead be added to the `azure.compute.LinuxVirtualMachineScaleSet` and `azure.compute.WindowsVirtualMachineScaleSet` resources.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/virtual_machine_scale_set.html.markdown.
 */
export class ScaleSet extends pulumi.CustomResource {
    /**
     * Get an existing ScaleSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScaleSetState, opts?: pulumi.CustomResourceOptions): ScaleSet {
        return new ScaleSet(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/scaleSet:ScaleSet';

    /**
     * Returns true if the given object is an instance of ScaleSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ScaleSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ScaleSet.__pulumiType;
    }

    /**
     * Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
     */
    public readonly automaticOsUpgrade!: pulumi.Output<boolean | undefined>;
    /**
     * A boot diagnostics profile block as referenced below.
     */
    public readonly bootDiagnostics!: pulumi.Output<outputs.compute.ScaleSetBootDiagnostics | undefined>;
    /**
     * Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
     */
    public readonly evictionPolicy!: pulumi.Output<string | undefined>;
    /**
     * Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
     */
    public readonly extensions!: pulumi.Output<outputs.compute.ScaleSetExtension[] | undefined>;
    /**
     * Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
     */
    public readonly healthProbeId!: pulumi.Output<string | undefined>;
    public readonly identity!: pulumi.Output<outputs.compute.ScaleSetIdentity>;
    /**
     * Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
     */
    public readonly licenseType!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A collection of network profile block as documented below.
     */
    public readonly networkProfiles!: pulumi.Output<outputs.compute.ScaleSetNetworkProfile[]>;
    /**
     * A Virtual Machine OS Profile block as documented below.
     */
    public readonly osProfile!: pulumi.Output<outputs.compute.ScaleSetOsProfile>;
    /**
     * A Linux config block as documented below.
     */
    public readonly osProfileLinuxConfig!: pulumi.Output<outputs.compute.ScaleSetOsProfileLinuxConfig>;
    /**
     * A collection of Secret blocks as documented below.
     */
    public readonly osProfileSecrets!: pulumi.Output<outputs.compute.ScaleSetOsProfileSecret[] | undefined>;
    /**
     * A Windows config block as documented below.
     */
    public readonly osProfileWindowsConfig!: pulumi.Output<outputs.compute.ScaleSetOsProfileWindowsConfig | undefined>;
    /**
     * Specifies whether the virtual machine scale set should be overprovisioned.
     */
    public readonly overprovision!: pulumi.Output<boolean | undefined>;
    /**
     * A plan block as documented below.
     */
    public readonly plan!: pulumi.Output<outputs.compute.ScaleSetPlan | undefined>;
    /**
     * Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
     */
    public readonly proximityPlacementGroupId!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
     */
    public readonly rollingUpgradePolicy!: pulumi.Output<outputs.compute.ScaleSetRollingUpgradePolicy | undefined>;
    /**
     * Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
     */
    public readonly singlePlacementGroup!: pulumi.Output<boolean | undefined>;
    /**
     * A sku block as documented below.
     */
    public readonly sku!: pulumi.Output<outputs.compute.ScaleSetSku>;
    /**
     * A storage profile data disk block as documented below
     */
    public readonly storageProfileDataDisks!: pulumi.Output<outputs.compute.ScaleSetStorageProfileDataDisk[] | undefined>;
    /**
     * A storage profile image reference block as documented below.
     */
    public readonly storageProfileImageReference!: pulumi.Output<outputs.compute.ScaleSetStorageProfileImageReference>;
    /**
     * A storage profile os disk block as documented below
     */
    public readonly storageProfileOsDisk!: pulumi.Output<outputs.compute.ScaleSetStorageProfileOsDisk>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
     */
    public readonly upgradePolicyMode!: pulumi.Output<string>;
    /**
     * A collection of availability zones to spread the Virtual Machines over.
     */
    public readonly zones!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ScaleSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScaleSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScaleSetArgs | ScaleSetState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScaleSetState | undefined;
            inputs["automaticOsUpgrade"] = state ? state.automaticOsUpgrade : undefined;
            inputs["bootDiagnostics"] = state ? state.bootDiagnostics : undefined;
            inputs["evictionPolicy"] = state ? state.evictionPolicy : undefined;
            inputs["extensions"] = state ? state.extensions : undefined;
            inputs["healthProbeId"] = state ? state.healthProbeId : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["licenseType"] = state ? state.licenseType : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkProfiles"] = state ? state.networkProfiles : undefined;
            inputs["osProfile"] = state ? state.osProfile : undefined;
            inputs["osProfileLinuxConfig"] = state ? state.osProfileLinuxConfig : undefined;
            inputs["osProfileSecrets"] = state ? state.osProfileSecrets : undefined;
            inputs["osProfileWindowsConfig"] = state ? state.osProfileWindowsConfig : undefined;
            inputs["overprovision"] = state ? state.overprovision : undefined;
            inputs["plan"] = state ? state.plan : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["proximityPlacementGroupId"] = state ? state.proximityPlacementGroupId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["rollingUpgradePolicy"] = state ? state.rollingUpgradePolicy : undefined;
            inputs["singlePlacementGroup"] = state ? state.singlePlacementGroup : undefined;
            inputs["sku"] = state ? state.sku : undefined;
            inputs["storageProfileDataDisks"] = state ? state.storageProfileDataDisks : undefined;
            inputs["storageProfileImageReference"] = state ? state.storageProfileImageReference : undefined;
            inputs["storageProfileOsDisk"] = state ? state.storageProfileOsDisk : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["upgradePolicyMode"] = state ? state.upgradePolicyMode : undefined;
            inputs["zones"] = state ? state.zones : undefined;
        } else {
            const args = argsOrState as ScaleSetArgs | undefined;
            if (!args || args.networkProfiles === undefined) {
                throw new Error("Missing required property 'networkProfiles'");
            }
            if (!args || args.osProfile === undefined) {
                throw new Error("Missing required property 'osProfile'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            if (!args || args.storageProfileOsDisk === undefined) {
                throw new Error("Missing required property 'storageProfileOsDisk'");
            }
            if (!args || args.upgradePolicyMode === undefined) {
                throw new Error("Missing required property 'upgradePolicyMode'");
            }
            inputs["automaticOsUpgrade"] = args ? args.automaticOsUpgrade : undefined;
            inputs["bootDiagnostics"] = args ? args.bootDiagnostics : undefined;
            inputs["evictionPolicy"] = args ? args.evictionPolicy : undefined;
            inputs["extensions"] = args ? args.extensions : undefined;
            inputs["healthProbeId"] = args ? args.healthProbeId : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["licenseType"] = args ? args.licenseType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkProfiles"] = args ? args.networkProfiles : undefined;
            inputs["osProfile"] = args ? args.osProfile : undefined;
            inputs["osProfileLinuxConfig"] = args ? args.osProfileLinuxConfig : undefined;
            inputs["osProfileSecrets"] = args ? args.osProfileSecrets : undefined;
            inputs["osProfileWindowsConfig"] = args ? args.osProfileWindowsConfig : undefined;
            inputs["overprovision"] = args ? args.overprovision : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["proximityPlacementGroupId"] = args ? args.proximityPlacementGroupId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["rollingUpgradePolicy"] = args ? args.rollingUpgradePolicy : undefined;
            inputs["singlePlacementGroup"] = args ? args.singlePlacementGroup : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["storageProfileDataDisks"] = args ? args.storageProfileDataDisks : undefined;
            inputs["storageProfileImageReference"] = args ? args.storageProfileImageReference : undefined;
            inputs["storageProfileOsDisk"] = args ? args.storageProfileOsDisk : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["upgradePolicyMode"] = args ? args.upgradePolicyMode : undefined;
            inputs["zones"] = args ? args.zones : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ScaleSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ScaleSet resources.
 */
export interface ScaleSetState {
    /**
     * Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
     */
    readonly automaticOsUpgrade?: pulumi.Input<boolean>;
    /**
     * A boot diagnostics profile block as referenced below.
     */
    readonly bootDiagnostics?: pulumi.Input<inputs.compute.ScaleSetBootDiagnostics>;
    /**
     * Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
     */
    readonly extensions?: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetExtension>[]>;
    /**
     * Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
     */
    readonly healthProbeId?: pulumi.Input<string>;
    readonly identity?: pulumi.Input<inputs.compute.ScaleSetIdentity>;
    /**
     * Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A collection of network profile block as documented below.
     */
    readonly networkProfiles?: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetNetworkProfile>[]>;
    /**
     * A Virtual Machine OS Profile block as documented below.
     */
    readonly osProfile?: pulumi.Input<inputs.compute.ScaleSetOsProfile>;
    /**
     * A Linux config block as documented below.
     */
    readonly osProfileLinuxConfig?: pulumi.Input<inputs.compute.ScaleSetOsProfileLinuxConfig>;
    /**
     * A collection of Secret blocks as documented below.
     */
    readonly osProfileSecrets?: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetOsProfileSecret>[]>;
    /**
     * A Windows config block as documented below.
     */
    readonly osProfileWindowsConfig?: pulumi.Input<inputs.compute.ScaleSetOsProfileWindowsConfig>;
    /**
     * Specifies whether the virtual machine scale set should be overprovisioned.
     */
    readonly overprovision?: pulumi.Input<boolean>;
    /**
     * A plan block as documented below.
     */
    readonly plan?: pulumi.Input<inputs.compute.ScaleSetPlan>;
    /**
     * Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
     */
    readonly proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
     */
    readonly rollingUpgradePolicy?: pulumi.Input<inputs.compute.ScaleSetRollingUpgradePolicy>;
    /**
     * Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
     */
    readonly singlePlacementGroup?: pulumi.Input<boolean>;
    /**
     * A sku block as documented below.
     */
    readonly sku?: pulumi.Input<inputs.compute.ScaleSetSku>;
    /**
     * A storage profile data disk block as documented below
     */
    readonly storageProfileDataDisks?: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetStorageProfileDataDisk>[]>;
    /**
     * A storage profile image reference block as documented below.
     */
    readonly storageProfileImageReference?: pulumi.Input<inputs.compute.ScaleSetStorageProfileImageReference>;
    /**
     * A storage profile os disk block as documented below
     */
    readonly storageProfileOsDisk?: pulumi.Input<inputs.compute.ScaleSetStorageProfileOsDisk>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
     */
    readonly upgradePolicyMode?: pulumi.Input<string>;
    /**
     * A collection of availability zones to spread the Virtual Machines over.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ScaleSet resource.
 */
export interface ScaleSetArgs {
    /**
     * Automatic OS patches can be applied by Azure to your scaleset. This is particularly useful when `upgradePolicyMode` is set to `Rolling`. Defaults to `false`.
     */
    readonly automaticOsUpgrade?: pulumi.Input<boolean>;
    /**
     * A boot diagnostics profile block as referenced below.
     */
    readonly bootDiagnostics?: pulumi.Input<inputs.compute.ScaleSetBootDiagnostics>;
    /**
     * Specifies the eviction policy for Virtual Machines in this Scale Set. Possible values are `Deallocate` and `Delete`.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * Can be specified multiple times to add extension profiles to the scale set. Each `extension` block supports the fields documented below.
     */
    readonly extensions?: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetExtension>[]>;
    /**
     * Specifies the identifier for the load balancer health probe. Required when using `Rolling` as your `upgradePolicyMode`.
     */
    readonly healthProbeId?: pulumi.Input<string>;
    readonly identity?: pulumi.Input<inputs.compute.ScaleSetIdentity>;
    /**
     * Specifies the Windows OS license type. If supplied, the only allowed values are `Windows_Client` and `Windows_Server`.
     */
    readonly licenseType?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the virtual machine scale set resource. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * A collection of network profile block as documented below.
     */
    readonly networkProfiles: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetNetworkProfile>[]>;
    /**
     * A Virtual Machine OS Profile block as documented below.
     */
    readonly osProfile: pulumi.Input<inputs.compute.ScaleSetOsProfile>;
    /**
     * A Linux config block as documented below.
     */
    readonly osProfileLinuxConfig?: pulumi.Input<inputs.compute.ScaleSetOsProfileLinuxConfig>;
    /**
     * A collection of Secret blocks as documented below.
     */
    readonly osProfileSecrets?: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetOsProfileSecret>[]>;
    /**
     * A Windows config block as documented below.
     */
    readonly osProfileWindowsConfig?: pulumi.Input<inputs.compute.ScaleSetOsProfileWindowsConfig>;
    /**
     * Specifies whether the virtual machine scale set should be overprovisioned.
     */
    readonly overprovision?: pulumi.Input<boolean>;
    /**
     * A plan block as documented below.
     */
    readonly plan?: pulumi.Input<inputs.compute.ScaleSetPlan>;
    /**
     * Specifies the priority for the Virtual Machines in the Scale Set. Defaults to `Regular`. Possible values are `Low` and `Regular`.
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * The ID of the Proximity Placement Group to which this Virtual Machine should be assigned. Changing this forces a new resource to be created
     */
    readonly proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the virtual machine scale set. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A `rollingUpgradePolicy` block as defined below. This is only applicable when the `upgradePolicyMode` is `Rolling`.
     */
    readonly rollingUpgradePolicy?: pulumi.Input<inputs.compute.ScaleSetRollingUpgradePolicy>;
    /**
     * Specifies whether the scale set is limited to a single placement group with a maximum size of 100 virtual machines. If set to false, managed disks must be used. Default is true. Changing this forces a new resource to be created. See [documentation](http://docs.microsoft.com/en-us/azure/virtual-machine-scale-sets/virtual-machine-scale-sets-placement-groups) for more information.
     */
    readonly singlePlacementGroup?: pulumi.Input<boolean>;
    /**
     * A sku block as documented below.
     */
    readonly sku: pulumi.Input<inputs.compute.ScaleSetSku>;
    /**
     * A storage profile data disk block as documented below
     */
    readonly storageProfileDataDisks?: pulumi.Input<pulumi.Input<inputs.compute.ScaleSetStorageProfileDataDisk>[]>;
    /**
     * A storage profile image reference block as documented below.
     */
    readonly storageProfileImageReference?: pulumi.Input<inputs.compute.ScaleSetStorageProfileImageReference>;
    /**
     * A storage profile os disk block as documented below
     */
    readonly storageProfileOsDisk: pulumi.Input<inputs.compute.ScaleSetStorageProfileOsDisk>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the mode of an upgrade to virtual machines in the scale set. Possible values, `Rolling`, `Manual`, or `Automatic`. When choosing `Rolling`, you will need to set a health probe.
     */
    readonly upgradePolicyMode: pulumi.Input<string>;
    /**
     * A collection of availability zones to spread the Virtual Machines over.
     */
    readonly zones?: pulumi.Input<pulumi.Input<string>[]>;
}
