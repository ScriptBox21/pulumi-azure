// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an API Management Service Diagnostic.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_diagnostic.html.markdown.
 */
export class Diagnostic extends pulumi.CustomResource {
    /**
     * Get an existing Diagnostic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DiagnosticState, opts?: pulumi.CustomResourceOptions): Diagnostic {
        return new Diagnostic(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/diagnostic:Diagnostic';

    /**
     * Returns true if the given object is an instance of Diagnostic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Diagnostic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Diagnostic.__pulumiType;
    }

    /**
     * The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * Indicates whether a Diagnostic should receive data or not.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
     */
    public readonly identifier!: pulumi.Output<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a Diagnostic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DiagnosticArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DiagnosticArgs | DiagnosticState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as DiagnosticState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["identifier"] = state ? state.identifier : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as DiagnosticArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.enabled === undefined) {
                throw new Error("Missing required property 'enabled'");
            }
            if (!args || args.identifier === undefined) {
                throw new Error("Missing required property 'identifier'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["identifier"] = args ? args.identifier : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Diagnostic.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Diagnostic resources.
 */
export interface DiagnosticState {
    /**
     * The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * Indicates whether a Diagnostic should receive data or not.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
     */
    readonly identifier?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Diagnostic resource.
 */
export interface DiagnosticArgs {
    /**
     * The Name of the API Management Service where this Diagnostic should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * Indicates whether a Diagnostic should receive data or not.
     */
    readonly enabled: pulumi.Input<boolean>;
    /**
     * The diagnostic identifier for the API Management Service. At this time the only supported value is `applicationinsights`. Changing this forces a new resource to be created.
     */
    readonly identifier: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
