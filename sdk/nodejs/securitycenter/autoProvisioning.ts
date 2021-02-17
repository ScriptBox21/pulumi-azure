// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Enables or disables the Security Center Auto Provisioning feature for the subscription
 *
 * > **NOTE:** There is no resource name required, it will always be "default"
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const example = new azure.securitycenter.AutoProvisioning("example", {
 *     autoProvision: "On",
 * });
 * ```
 *
 * ## Import
 *
 * Security Center Auto Provisioning can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:securitycenter/autoProvisioning:AutoProvisioning example /subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Security/autoProvisioningSettings/default
 * ```
 */
export class AutoProvisioning extends pulumi.CustomResource {
    /**
     * Get an existing AutoProvisioning resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoProvisioningState, opts?: pulumi.CustomResourceOptions): AutoProvisioning {
        return new AutoProvisioning(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:securitycenter/autoProvisioning:AutoProvisioning';

    /**
     * Returns true if the given object is an instance of AutoProvisioning.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoProvisioning {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoProvisioning.__pulumiType;
    }

    /**
     * Should the security agent be automatically provisioned on Virtual Machines in this subscription? Possible values are `On` (to install the security agent automatically, if it's missing) or `Off` (to not install the security agent automatically).
     */
    public readonly autoProvision!: pulumi.Output<string>;

    /**
     * Create a AutoProvisioning resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoProvisioningArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoProvisioningArgs | AutoProvisioningState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoProvisioningState | undefined;
            inputs["autoProvision"] = state ? state.autoProvision : undefined;
        } else {
            const args = argsOrState as AutoProvisioningArgs | undefined;
            if ((!args || args.autoProvision === undefined) && !opts.urn) {
                throw new Error("Missing required property 'autoProvision'");
            }
            inputs["autoProvision"] = args ? args.autoProvision : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AutoProvisioning.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoProvisioning resources.
 */
export interface AutoProvisioningState {
    /**
     * Should the security agent be automatically provisioned on Virtual Machines in this subscription? Possible values are `On` (to install the security agent automatically, if it's missing) or `Off` (to not install the security agent automatically).
     */
    readonly autoProvision?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AutoProvisioning resource.
 */
export interface AutoProvisioningArgs {
    /**
     * Should the security agent be automatically provisioned on Virtual Machines in this subscription? Possible values are `On` (to install the security agent automatically, if it's missing) or `Off` (to not install the security agent automatically).
     */
    readonly autoProvision: pulumi.Input<string>;
}
