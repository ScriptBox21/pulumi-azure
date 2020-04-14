// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages an API Management Microsoft Identity Provider.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/api_management_identity_provider_microsoft.html.markdown.
 */
export class IdentityProviderMicrosoft extends pulumi.CustomResource {
    /**
     * Get an existing IdentityProviderMicrosoft resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IdentityProviderMicrosoftState, opts?: pulumi.CustomResourceOptions): IdentityProviderMicrosoft {
        return new IdentityProviderMicrosoft(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/identityProviderMicrosoft:IdentityProviderMicrosoft';

    /**
     * Returns true if the given object is an instance of IdentityProviderMicrosoft.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IdentityProviderMicrosoft {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IdentityProviderMicrosoft.__pulumiType;
    }

    /**
     * The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * Client Id of the Azure AD Application.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * Client secret of the Azure AD Application.
     */
    public readonly clientSecret!: pulumi.Output<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a IdentityProviderMicrosoft resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IdentityProviderMicrosoftArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IdentityProviderMicrosoftArgs | IdentityProviderMicrosoftState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as IdentityProviderMicrosoftState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as IdentityProviderMicrosoftArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.clientSecret === undefined) {
                throw new Error("Missing required property 'clientSecret'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(IdentityProviderMicrosoft.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IdentityProviderMicrosoft resources.
 */
export interface IdentityProviderMicrosoftState {
    /**
     * The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * Client Id of the Azure AD Application.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * Client secret of the Azure AD Application.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IdentityProviderMicrosoft resource.
 */
export interface IdentityProviderMicrosoftArgs {
    /**
     * The Name of the API Management Service where this Microsoft Identity Provider should be created. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * Client Id of the Azure AD Application.
     */
    readonly clientId: pulumi.Input<string>;
    /**
     * Client secret of the Azure AD Application.
     */
    readonly clientSecret: pulumi.Input<string>;
    /**
     * The Name of the Resource Group where the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
