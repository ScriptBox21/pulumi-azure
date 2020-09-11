// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages a Linux Virtual Machine.
 *
 * ## Disclaimers
 *
 * > **Note** This provider will automatically remove the OS Disk by default - this behaviour can be configured using the `features` configuration within the Provider configuration block.
 *
 * > **Note** This resource does not support Unmanaged Disks. If you need to use Unmanaged Disks you can continue to use the `azure.compute.VirtualMachine` resource instead.
 *
 * > **Note** This resource does not support attaching existing OS Disks. You can instead capture an image of the OS Disk or continue to use the `azure.compute.VirtualMachine` resource instead.
 *
 * > In this release there's a known issue where the `publicIpAddress` and `publicIpAddresses` fields may not be fully populated for Dynamic Public IP's.
 *
 * ## Example Usage
 *
 * This example provisions a basic Linux Virtual Machine on an internal network.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * import * from "fs";
 *
 * const exampleResourceGroup = new azure.core.ResourceGroup("exampleResourceGroup", {location: "West Europe"});
 * const exampleVirtualNetwork = new azure.network.VirtualNetwork("exampleVirtualNetwork", {
 *     addressSpaces: ["10.0.0.0/16"],
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 * });
 * const exampleSubnet = new azure.network.Subnet("exampleSubnet", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     virtualNetworkName: exampleVirtualNetwork.name,
 *     addressPrefix: "10.0.2.0/24",
 * });
 * const exampleNetworkInterface = new azure.network.NetworkInterface("exampleNetworkInterface", {
 *     location: exampleResourceGroup.location,
 *     resourceGroupName: exampleResourceGroup.name,
 *     ipConfigurations: [{
 *         name: "internal",
 *         subnetId: exampleSubnet.id,
 *         privateIpAddressAllocation: "Dynamic",
 *     }],
 * });
 * const exampleLinuxVirtualMachine = new azure.compute.LinuxVirtualMachine("exampleLinuxVirtualMachine", {
 *     resourceGroupName: exampleResourceGroup.name,
 *     location: exampleResourceGroup.location,
 *     size: "Standard_F2",
 *     adminUsername: "adminuser",
 *     networkInterfaceIds: [exampleNetworkInterface.id],
 *     adminSshKeys: [{
 *         username: "adminuser",
 *         publicKey: fs.readFileSync("~/.ssh/id_rsa.pub"),
 *     }],
 *     osDisk: {
 *         caching: "ReadWrite",
 *         storageAccountType: "Standard_LRS",
 *     },
 *     sourceImageReference: {
 *         publisher: "Canonical",
 *         offer: "UbuntuServer",
 *         sku: "16.04-LTS",
 *         version: "latest",
 *     },
 * });
 * ```
 */
export class LinuxVirtualMachine extends pulumi.CustomResource {
    /**
     * Get an existing LinuxVirtualMachine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LinuxVirtualMachineState, opts?: pulumi.CustomResourceOptions): LinuxVirtualMachine {
        return new LinuxVirtualMachine(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:compute/linuxVirtualMachine:LinuxVirtualMachine';

    /**
     * Returns true if the given object is an instance of LinuxVirtualMachine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LinuxVirtualMachine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LinuxVirtualMachine.__pulumiType;
    }

    /**
     * A `additionalCapabilities` block as defined below.
     */
    public readonly additionalCapabilities!: pulumi.Output<outputs.compute.LinuxVirtualMachineAdditionalCapabilities | undefined>;
    /**
     * The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
     */
    public readonly adminPassword!: pulumi.Output<string | undefined>;
    /**
     * One or more `adminSshKey` blocks as defined below.
     */
    public readonly adminSshKeys!: pulumi.Output<outputs.compute.LinuxVirtualMachineAdminSshKey[] | undefined>;
    /**
     * The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
     */
    public readonly adminUsername!: pulumi.Output<string>;
    /**
     * Should Extension Operations be allowed on this Virtual Machine?
     */
    public readonly allowExtensionOperations!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
     */
    public readonly availabilitySetId!: pulumi.Output<string | undefined>;
    /**
     * A `bootDiagnostics` block as defined below.
     */
    public readonly bootDiagnostics!: pulumi.Output<outputs.compute.LinuxVirtualMachineBootDiagnostics | undefined>;
    /**
     * Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
     */
    public readonly computerName!: pulumi.Output<string>;
    /**
     * The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
     */
    public readonly customData!: pulumi.Output<string | undefined>;
    /**
     * The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
     */
    public readonly dedicatedHostId!: pulumi.Output<string | undefined>;
    /**
     * Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
     */
    public readonly disablePasswordAuthentication!: pulumi.Output<boolean | undefined>;
    /**
     * Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
     */
    public readonly encryptionAtHostEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
     */
    public readonly evictionPolicy!: pulumi.Output<string | undefined>;
    /**
     * An `identity` block as defined below.
     */
    public readonly identity!: pulumi.Output<outputs.compute.LinuxVirtualMachineIdentity | undefined>;
    /**
     * The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
     */
    public readonly maxBidPrice!: pulumi.Output<number | undefined>;
    /**
     * The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
     */
    public readonly networkInterfaceIds!: pulumi.Output<string[]>;
    /**
     * A `osDisk` block as defined below.
     */
    public readonly osDisk!: pulumi.Output<outputs.compute.LinuxVirtualMachineOsDisk>;
    /**
     * A `plan` block as defined below. Changing this forces a new resource to be created.
     */
    public readonly plan!: pulumi.Output<outputs.compute.LinuxVirtualMachinePlan | undefined>;
    /**
     * Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
     */
    public readonly priority!: pulumi.Output<string | undefined>;
    /**
     * The Primary Private IP Address assigned to this Virtual Machine.
     */
    public /*out*/ readonly privateIpAddress!: pulumi.Output<string>;
    /**
     * A list of Private IP Addresses assigned to this Virtual Machine.
     */
    public /*out*/ readonly privateIpAddresses!: pulumi.Output<string[]>;
    /**
     * Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
     */
    public readonly provisionVmAgent!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
     */
    public readonly proximityPlacementGroupId!: pulumi.Output<string | undefined>;
    /**
     * The Primary Public IP Address assigned to this Virtual Machine.
     */
    public /*out*/ readonly publicIpAddress!: pulumi.Output<string>;
    /**
     * A list of the Public IP Addresses assigned to this Virtual Machine.
     */
    public /*out*/ readonly publicIpAddresses!: pulumi.Output<string[]>;
    /**
     * The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * One or more `secret` blocks as defined below.
     */
    public readonly secrets!: pulumi.Output<outputs.compute.LinuxVirtualMachineSecret[] | undefined>;
    /**
     * The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
     */
    public readonly size!: pulumi.Output<string>;
    /**
     * The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
     */
    public readonly sourceImageId!: pulumi.Output<string | undefined>;
    /**
     * A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
     */
    public readonly sourceImageReference!: pulumi.Output<outputs.compute.LinuxVirtualMachineSourceImageReference | undefined>;
    /**
     * A mapping of tags which should be assigned to this Virtual Machine.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A 128-bit identifier which uniquely identifies this Virtual Machine.
     */
    public /*out*/ readonly virtualMachineId!: pulumi.Output<string>;
    /**
     * Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
     */
    public readonly virtualMachineScaleSetId!: pulumi.Output<string | undefined>;
    /**
     * The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a LinuxVirtualMachine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LinuxVirtualMachineArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LinuxVirtualMachineArgs | LinuxVirtualMachineState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as LinuxVirtualMachineState | undefined;
            inputs["additionalCapabilities"] = state ? state.additionalCapabilities : undefined;
            inputs["adminPassword"] = state ? state.adminPassword : undefined;
            inputs["adminSshKeys"] = state ? state.adminSshKeys : undefined;
            inputs["adminUsername"] = state ? state.adminUsername : undefined;
            inputs["allowExtensionOperations"] = state ? state.allowExtensionOperations : undefined;
            inputs["availabilitySetId"] = state ? state.availabilitySetId : undefined;
            inputs["bootDiagnostics"] = state ? state.bootDiagnostics : undefined;
            inputs["computerName"] = state ? state.computerName : undefined;
            inputs["customData"] = state ? state.customData : undefined;
            inputs["dedicatedHostId"] = state ? state.dedicatedHostId : undefined;
            inputs["disablePasswordAuthentication"] = state ? state.disablePasswordAuthentication : undefined;
            inputs["encryptionAtHostEnabled"] = state ? state.encryptionAtHostEnabled : undefined;
            inputs["evictionPolicy"] = state ? state.evictionPolicy : undefined;
            inputs["identity"] = state ? state.identity : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["maxBidPrice"] = state ? state.maxBidPrice : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["networkInterfaceIds"] = state ? state.networkInterfaceIds : undefined;
            inputs["osDisk"] = state ? state.osDisk : undefined;
            inputs["plan"] = state ? state.plan : undefined;
            inputs["priority"] = state ? state.priority : undefined;
            inputs["privateIpAddress"] = state ? state.privateIpAddress : undefined;
            inputs["privateIpAddresses"] = state ? state.privateIpAddresses : undefined;
            inputs["provisionVmAgent"] = state ? state.provisionVmAgent : undefined;
            inputs["proximityPlacementGroupId"] = state ? state.proximityPlacementGroupId : undefined;
            inputs["publicIpAddress"] = state ? state.publicIpAddress : undefined;
            inputs["publicIpAddresses"] = state ? state.publicIpAddresses : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["secrets"] = state ? state.secrets : undefined;
            inputs["size"] = state ? state.size : undefined;
            inputs["sourceImageId"] = state ? state.sourceImageId : undefined;
            inputs["sourceImageReference"] = state ? state.sourceImageReference : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["virtualMachineId"] = state ? state.virtualMachineId : undefined;
            inputs["virtualMachineScaleSetId"] = state ? state.virtualMachineScaleSetId : undefined;
            inputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as LinuxVirtualMachineArgs | undefined;
            if (!args || args.adminUsername === undefined) {
                throw new Error("Missing required property 'adminUsername'");
            }
            if (!args || args.networkInterfaceIds === undefined) {
                throw new Error("Missing required property 'networkInterfaceIds'");
            }
            if (!args || args.osDisk === undefined) {
                throw new Error("Missing required property 'osDisk'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.size === undefined) {
                throw new Error("Missing required property 'size'");
            }
            inputs["additionalCapabilities"] = args ? args.additionalCapabilities : undefined;
            inputs["adminPassword"] = args ? args.adminPassword : undefined;
            inputs["adminSshKeys"] = args ? args.adminSshKeys : undefined;
            inputs["adminUsername"] = args ? args.adminUsername : undefined;
            inputs["allowExtensionOperations"] = args ? args.allowExtensionOperations : undefined;
            inputs["availabilitySetId"] = args ? args.availabilitySetId : undefined;
            inputs["bootDiagnostics"] = args ? args.bootDiagnostics : undefined;
            inputs["computerName"] = args ? args.computerName : undefined;
            inputs["customData"] = args ? args.customData : undefined;
            inputs["dedicatedHostId"] = args ? args.dedicatedHostId : undefined;
            inputs["disablePasswordAuthentication"] = args ? args.disablePasswordAuthentication : undefined;
            inputs["encryptionAtHostEnabled"] = args ? args.encryptionAtHostEnabled : undefined;
            inputs["evictionPolicy"] = args ? args.evictionPolicy : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxBidPrice"] = args ? args.maxBidPrice : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaceIds"] = args ? args.networkInterfaceIds : undefined;
            inputs["osDisk"] = args ? args.osDisk : undefined;
            inputs["plan"] = args ? args.plan : undefined;
            inputs["priority"] = args ? args.priority : undefined;
            inputs["provisionVmAgent"] = args ? args.provisionVmAgent : undefined;
            inputs["proximityPlacementGroupId"] = args ? args.proximityPlacementGroupId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["secrets"] = args ? args.secrets : undefined;
            inputs["size"] = args ? args.size : undefined;
            inputs["sourceImageId"] = args ? args.sourceImageId : undefined;
            inputs["sourceImageReference"] = args ? args.sourceImageReference : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualMachineScaleSetId"] = args ? args.virtualMachineScaleSetId : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["privateIpAddress"] = undefined /*out*/;
            inputs["privateIpAddresses"] = undefined /*out*/;
            inputs["publicIpAddress"] = undefined /*out*/;
            inputs["publicIpAddresses"] = undefined /*out*/;
            inputs["virtualMachineId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(LinuxVirtualMachine.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LinuxVirtualMachine resources.
 */
export interface LinuxVirtualMachineState {
    /**
     * A `additionalCapabilities` block as defined below.
     */
    readonly additionalCapabilities?: pulumi.Input<inputs.compute.LinuxVirtualMachineAdditionalCapabilities>;
    /**
     * The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly adminPassword?: pulumi.Input<string>;
    /**
     * One or more `adminSshKey` blocks as defined below.
     */
    readonly adminSshKeys?: pulumi.Input<pulumi.Input<inputs.compute.LinuxVirtualMachineAdminSshKey>[]>;
    /**
     * The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly adminUsername?: pulumi.Input<string>;
    /**
     * Should Extension Operations be allowed on this Virtual Machine?
     */
    readonly allowExtensionOperations?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
     */
    readonly availabilitySetId?: pulumi.Input<string>;
    /**
     * A `bootDiagnostics` block as defined below.
     */
    readonly bootDiagnostics?: pulumi.Input<inputs.compute.LinuxVirtualMachineBootDiagnostics>;
    /**
     * Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
     */
    readonly computerName?: pulumi.Input<string>;
    /**
     * The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly customData?: pulumi.Input<string>;
    /**
     * The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
     */
    readonly dedicatedHostId?: pulumi.Input<string>;
    /**
     * Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
     */
    readonly disablePasswordAuthentication?: pulumi.Input<boolean>;
    /**
     * Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
     */
    readonly encryptionAtHostEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.compute.LinuxVirtualMachineIdentity>;
    /**
     * The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
     */
    readonly maxBidPrice?: pulumi.Input<number>;
    /**
     * The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
     */
    readonly networkInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `osDisk` block as defined below.
     */
    readonly osDisk?: pulumi.Input<inputs.compute.LinuxVirtualMachineOsDisk>;
    /**
     * A `plan` block as defined below. Changing this forces a new resource to be created.
     */
    readonly plan?: pulumi.Input<inputs.compute.LinuxVirtualMachinePlan>;
    /**
     * Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * The Primary Private IP Address assigned to this Virtual Machine.
     */
    readonly privateIpAddress?: pulumi.Input<string>;
    /**
     * A list of Private IP Addresses assigned to this Virtual Machine.
     */
    readonly privateIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
     */
    readonly provisionVmAgent?: pulumi.Input<boolean>;
    /**
     * The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
     */
    readonly proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The Primary Public IP Address assigned to this Virtual Machine.
     */
    readonly publicIpAddress?: pulumi.Input<string>;
    /**
     * A list of the Public IP Addresses assigned to this Virtual Machine.
     */
    readonly publicIpAddresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * One or more `secret` blocks as defined below.
     */
    readonly secrets?: pulumi.Input<pulumi.Input<inputs.compute.LinuxVirtualMachineSecret>[]>;
    /**
     * The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
     */
    readonly size?: pulumi.Input<string>;
    /**
     * The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
     */
    readonly sourceImageId?: pulumi.Input<string>;
    /**
     * A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
     */
    readonly sourceImageReference?: pulumi.Input<inputs.compute.LinuxVirtualMachineSourceImageReference>;
    /**
     * A mapping of tags which should be assigned to this Virtual Machine.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A 128-bit identifier which uniquely identifies this Virtual Machine.
     */
    readonly virtualMachineId?: pulumi.Input<string>;
    /**
     * Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
     */
    readonly virtualMachineScaleSetId?: pulumi.Input<string>;
    /**
     * The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
     */
    readonly zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LinuxVirtualMachine resource.
 */
export interface LinuxVirtualMachineArgs {
    /**
     * A `additionalCapabilities` block as defined below.
     */
    readonly additionalCapabilities?: pulumi.Input<inputs.compute.LinuxVirtualMachineAdditionalCapabilities>;
    /**
     * The Password which should be used for the local-administrator on this Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly adminPassword?: pulumi.Input<string>;
    /**
     * One or more `adminSshKey` blocks as defined below.
     */
    readonly adminSshKeys?: pulumi.Input<pulumi.Input<inputs.compute.LinuxVirtualMachineAdminSshKey>[]>;
    /**
     * The username of the local administrator used for the Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly adminUsername: pulumi.Input<string>;
    /**
     * Should Extension Operations be allowed on this Virtual Machine?
     */
    readonly allowExtensionOperations?: pulumi.Input<boolean>;
    /**
     * Specifies the ID of the Availability Set in which the Virtual Machine should exist. Changing this forces a new resource to be created.
     */
    readonly availabilitySetId?: pulumi.Input<string>;
    /**
     * A `bootDiagnostics` block as defined below.
     */
    readonly bootDiagnostics?: pulumi.Input<inputs.compute.LinuxVirtualMachineBootDiagnostics>;
    /**
     * Specifies the Hostname which should be used for this Virtual Machine. If unspecified this defaults to the value for the `name` field. If the value of the `name` field is not a valid `computerName`, then you must specify `computerName`. Changing this forces a new resource to be created.
     */
    readonly computerName?: pulumi.Input<string>;
    /**
     * The Base64-Encoded Custom Data which should be used for this Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly customData?: pulumi.Input<string>;
    /**
     * The ID of a Dedicated Host where this machine should be run on. Changing this forces a new resource to be created.
     */
    readonly dedicatedHostId?: pulumi.Input<string>;
    /**
     * Should Password Authentication be disabled on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
     */
    readonly disablePasswordAuthentication?: pulumi.Input<boolean>;
    /**
     * Should all of the disks (including the temp disk) attached to this Virtual Machine be encrypted by enabling Encryption at Host?
     */
    readonly encryptionAtHostEnabled?: pulumi.Input<boolean>;
    /**
     * Specifies what should happen when the Virtual Machine is evicted for price reasons when using a Spot instance. At this time the only supported value is `Deallocate`. Changing this forces a new resource to be created.
     */
    readonly evictionPolicy?: pulumi.Input<string>;
    /**
     * An `identity` block as defined below.
     */
    readonly identity?: pulumi.Input<inputs.compute.LinuxVirtualMachineIdentity>;
    /**
     * The Azure location where the Linux Virtual Machine should exist. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The maximum price you're willing to pay for this Virtual Machine, in US Dollars; which must be greater than the current spot price. If this bid price falls below the current spot price the Virtual Machine will be evicted using the `evictionPolicy`. Defaults to `-1`, which means that the Virtual Machine should not be evicted for price reasons.
     */
    readonly maxBidPrice?: pulumi.Input<number>;
    /**
     * The name of the Linux Virtual Machine. Changing this forces a new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * . A list of Network Interface ID's which should be attached to this Virtual Machine. The first Network Interface ID in this list will be the Primary Network Interface on the Virtual Machine.
     */
    readonly networkInterfaceIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A `osDisk` block as defined below.
     */
    readonly osDisk: pulumi.Input<inputs.compute.LinuxVirtualMachineOsDisk>;
    /**
     * A `plan` block as defined below. Changing this forces a new resource to be created.
     */
    readonly plan?: pulumi.Input<inputs.compute.LinuxVirtualMachinePlan>;
    /**
     * Specifies the priority of this Virtual Machine. Possible values are `Regular` and `Spot`. Defaults to `Regular`. Changing this forces a new resource to be created.
     */
    readonly priority?: pulumi.Input<string>;
    /**
     * Should the Azure VM Agent be provisioned on this Virtual Machine? Defaults to `true`. Changing this forces a new resource to be created.
     */
    readonly provisionVmAgent?: pulumi.Input<boolean>;
    /**
     * The ID of the Proximity Placement Group which the Virtual Machine should be assigned to. Changing this forces a new resource to be created.
     */
    readonly proximityPlacementGroupId?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the Linux Virtual Machine should be exist. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * One or more `secret` blocks as defined below.
     */
    readonly secrets?: pulumi.Input<pulumi.Input<inputs.compute.LinuxVirtualMachineSecret>[]>;
    /**
     * The SKU which should be used for this Virtual Machine, such as `Standard_F2`.
     */
    readonly size: pulumi.Input<string>;
    /**
     * The ID of the Image which this Virtual Machine should be created from. Changing this forces a new resource to be created.
     */
    readonly sourceImageId?: pulumi.Input<string>;
    /**
     * A `sourceImageReference` block as defined below. Changing this forces a new resource to be created.
     */
    readonly sourceImageReference?: pulumi.Input<inputs.compute.LinuxVirtualMachineSourceImageReference>;
    /**
     * A mapping of tags which should be assigned to this Virtual Machine.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the Orchestrated Virtual Machine Scale Set that this Virtual Machine should be created within. Changing this forces a new resource to be created.
     */
    readonly virtualMachineScaleSetId?: pulumi.Input<string>;
    /**
     * The Zone in which this Virtual Machine should be created. Changing this forces a new resource to be created.
     */
    readonly zone?: pulumi.Input<string>;
}
