// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages the association between a Network Interface and a Load Balancer's Backend Address Pool.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/network_interface_backend_address_pool_association.html.markdown.
 */
export class NetworkInterfaceBackendAddressPoolAssociation extends pulumi.CustomResource {
    /**
     * Get an existing NetworkInterfaceBackendAddressPoolAssociation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkInterfaceBackendAddressPoolAssociationState, opts?: pulumi.CustomResourceOptions): NetworkInterfaceBackendAddressPoolAssociation {
        return new NetworkInterfaceBackendAddressPoolAssociation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:network/networkInterfaceBackendAddressPoolAssociation:NetworkInterfaceBackendAddressPoolAssociation';

    /**
     * Returns true if the given object is an instance of NetworkInterfaceBackendAddressPoolAssociation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkInterfaceBackendAddressPoolAssociation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkInterfaceBackendAddressPoolAssociation.__pulumiType;
    }

    /**
     * The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    public readonly backendAddressPoolId!: pulumi.Output<string>;
    /**
     * The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
     */
    public readonly ipConfigurationName!: pulumi.Output<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    public readonly networkInterfaceId!: pulumi.Output<string>;

    /**
     * Create a NetworkInterfaceBackendAddressPoolAssociation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkInterfaceBackendAddressPoolAssociationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkInterfaceBackendAddressPoolAssociationArgs | NetworkInterfaceBackendAddressPoolAssociationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NetworkInterfaceBackendAddressPoolAssociationState | undefined;
            inputs["backendAddressPoolId"] = state ? state.backendAddressPoolId : undefined;
            inputs["ipConfigurationName"] = state ? state.ipConfigurationName : undefined;
            inputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
        } else {
            const args = argsOrState as NetworkInterfaceBackendAddressPoolAssociationArgs | undefined;
            if (!args || args.backendAddressPoolId === undefined) {
                throw new Error("Missing required property 'backendAddressPoolId'");
            }
            if (!args || args.ipConfigurationName === undefined) {
                throw new Error("Missing required property 'ipConfigurationName'");
            }
            if (!args || args.networkInterfaceId === undefined) {
                throw new Error("Missing required property 'networkInterfaceId'");
            }
            inputs["backendAddressPoolId"] = args ? args.backendAddressPoolId : undefined;
            inputs["ipConfigurationName"] = args ? args.ipConfigurationName : undefined;
            inputs["networkInterfaceId"] = args ? args.networkInterfaceId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(NetworkInterfaceBackendAddressPoolAssociation.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkInterfaceBackendAddressPoolAssociation resources.
 */
export interface NetworkInterfaceBackendAddressPoolAssociationState {
    /**
     * The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    readonly backendAddressPoolId?: pulumi.Input<string>;
    /**
     * The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
     */
    readonly ipConfigurationName?: pulumi.Input<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    readonly networkInterfaceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkInterfaceBackendAddressPoolAssociation resource.
 */
export interface NetworkInterfaceBackendAddressPoolAssociationArgs {
    /**
     * The ID of the Load Balancer Backend Address Pool which this Network Interface which should be connected to. Changing this forces a new resource to be created.
     */
    readonly backendAddressPoolId: pulumi.Input<string>;
    /**
     * The Name of the IP Configuration within the Network Interface which should be connected to the Backend Address Pool. Changing this forces a new resource to be created.
     */
    readonly ipConfigurationName: pulumi.Input<string>;
    /**
     * The ID of the Network Interface. Changing this forces a new resource to be created.
     */
    readonly networkInterfaceId: pulumi.Input<string>;
}
