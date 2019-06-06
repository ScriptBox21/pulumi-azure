// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage an Application Insights component.
 * 
 * ## Example Usage
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 * 
 * const testResourceGroup = new azure.core.ResourceGroup("test", {
 *     location: "West Europe",
 *     name: "tf-test",
 * });
 * const testInsights = new azure.appinsights.Insights("test", {
 *     applicationType: "Web",
 *     location: "West Europe",
 *     name: "tf-test-appinsights",
 *     resourceGroupName: testResourceGroup.name,
 * });
 * 
 * export const appId = testInsights.appId;
 * export const instrumentationKey = testInsights.instrumentationKey;
 * ```
 */
export class Insights extends pulumi.CustomResource {
    /**
     * Get an existing Insights resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InsightsState, opts?: pulumi.CustomResourceOptions): Insights {
        return new Insights(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:appinsights/insights:Insights';

    /**
     * Returns true if the given object is an instance of Insights.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Insights {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Insights.__pulumiType;
    }

    /**
     * The App ID associated with this Application Insights component.
     */
    public /*out*/ readonly appId!: pulumi.Output<string>;
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    public readonly applicationType!: pulumi.Output<string>;
    /**
     * The Instrumentation Key for this Application Insights component.
     */
    public /*out*/ readonly instrumentationKey!: pulumi.Output<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a Insights resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InsightsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InsightsArgs | InsightsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as InsightsState | undefined;
            inputs["appId"] = state ? state.appId : undefined;
            inputs["applicationType"] = state ? state.applicationType : undefined;
            inputs["instrumentationKey"] = state ? state.instrumentationKey : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as InsightsArgs | undefined;
            if (!args || args.applicationType === undefined) {
                throw new Error("Missing required property 'applicationType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["applicationType"] = args ? args.applicationType : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["appId"] = undefined /*out*/;
            inputs["instrumentationKey"] = undefined /*out*/;
        }
        super(Insights.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Insights resources.
 */
export interface InsightsState {
    /**
     * The App ID associated with this Application Insights component.
     */
    readonly appId?: pulumi.Input<string>;
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    readonly applicationType?: pulumi.Input<string>;
    /**
     * The Instrumentation Key for this Application Insights component.
     */
    readonly instrumentationKey?: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a Insights resource.
 */
export interface InsightsArgs {
    /**
     * Specifies the type of Application Insights to create. Valid values are `ios` for _iOS_, `java` for _Java web_, `MobileCenter` for _App Center_, `Node.JS` for _Node.js_, `other` for _General_, `phone` for _Windows Phone_, `store` for _Windows Store_ and `web` for _ASP.NET_. Please note these values are case sensitive; unmatched values are treated as _ASP.NET_ by Azure. Changing this forces a new resource to be created.
     */
    readonly applicationType: pulumi.Input<string>;
    /**
     * Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * Specifies the name of the Application Insights component. Changing this forces a
     * new resource to be created.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to
     * create the Application Insights component.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: any}>;
}
