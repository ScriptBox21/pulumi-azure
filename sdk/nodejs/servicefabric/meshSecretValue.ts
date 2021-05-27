// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Service Fabric Mesh Secret Value can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:servicefabric/meshSecretValue:MeshSecretValue value1 /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/mygroup1/providers/Microsoft.ServiceFabricMesh/secrets/secret1/values/value1
 * ```
 */
export class MeshSecretValue extends pulumi.CustomResource {
    /**
     * Get an existing MeshSecretValue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MeshSecretValueState, opts?: pulumi.CustomResourceOptions): MeshSecretValue {
        return new MeshSecretValue(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:servicefabric/meshSecretValue:MeshSecretValue';

    /**
     * Returns true if the given object is an instance of MeshSecretValue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MeshSecretValue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MeshSecretValue.__pulumiType;
    }

    /**
     * Specifies the Azure Region where the Service Fabric Mesh Secret Value should exist. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the Service Fabric Mesh Secret Value. Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The id of the Service Fabric Mesh Secret in which the value will be applied to. Changing this forces a new resource to be created.
     */
    public readonly serviceFabricMeshSecretId!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Specifies the value that will be applied to the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
     */
    public readonly value!: pulumi.Output<string>;

    /**
     * Create a MeshSecretValue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MeshSecretValueArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MeshSecretValueArgs | MeshSecretValueState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MeshSecretValueState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["serviceFabricMeshSecretId"] = state ? state.serviceFabricMeshSecretId : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["value"] = state ? state.value : undefined;
        } else {
            const args = argsOrState as MeshSecretValueArgs | undefined;
            if ((!args || args.serviceFabricMeshSecretId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceFabricMeshSecretId'");
            }
            if ((!args || args.value === undefined) && !opts.urn) {
                throw new Error("Missing required property 'value'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["serviceFabricMeshSecretId"] = args ? args.serviceFabricMeshSecretId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["value"] = args ? args.value : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MeshSecretValue.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MeshSecretValue resources.
 */
export interface MeshSecretValueState {
    /**
     * Specifies the Azure Region where the Service Fabric Mesh Secret Value should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Service Fabric Mesh Secret Value. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the Service Fabric Mesh Secret in which the value will be applied to. Changing this forces a new resource to be created.
     */
    serviceFabricMeshSecretId?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the value that will be applied to the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
     */
    value?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MeshSecretValue resource.
 */
export interface MeshSecretValueArgs {
    /**
     * Specifies the Azure Region where the Service Fabric Mesh Secret Value should exist. Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * The name of the Service Fabric Mesh Secret Value. Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The id of the Service Fabric Mesh Secret in which the value will be applied to. Changing this forces a new resource to be created.
     */
    serviceFabricMeshSecretId: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Specifies the value that will be applied to the Service Fabric Mesh Secret. Changing this forces a new resource to be created.
     */
    value: pulumi.Input<string>;
}
