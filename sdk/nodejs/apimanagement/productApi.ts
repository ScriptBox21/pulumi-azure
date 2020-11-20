// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an API Management API Assignment to a Product.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const exampleService = azure.apimanagement.getService({
 *     name: "example-api",
 *     resourceGroupName: "example-resources",
 * });
 * const exampleApi = Promise.all([exampleService, exampleService]).then(([exampleService, exampleService1]) => azure.apimanagement.getApi({
 *     name: "search-api",
 *     apiManagementName: exampleService.name,
 *     resourceGroupName: exampleService1.resourceGroupName,
 *     revision: "2",
 * }));
 * const exampleProduct = Promise.all([exampleService, exampleService]).then(([exampleService, exampleService1]) => azure.apimanagement.getProduct({
 *     productId: "my-product",
 *     apiManagementName: exampleService.name,
 *     resourceGroupName: exampleService1.resourceGroupName,
 * }));
 * const exampleProductApi = new azure.apimanagement.ProductApi("exampleProductApi", {
 *     apiName: exampleApi.then(exampleApi => exampleApi.name),
 *     productId: exampleProduct.then(exampleProduct => exampleProduct.productId),
 *     apiManagementName: exampleService.then(exampleService => exampleService.name),
 *     resourceGroupName: exampleService.then(exampleService => exampleService.resourceGroupName),
 * });
 * ```
 *
 * ## Import
 *
 * API Management Product API's can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:apimanagement/productApi:ProductApi example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.ApiManagement/service/service1/products/exampleId/apis/apiId
 * ```
 */
export class ProductApi extends pulumi.CustomResource {
    /**
     * Get an existing ProductApi resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProductApiState, opts?: pulumi.CustomResourceOptions): ProductApi {
        return new ProductApi(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:apimanagement/productApi:ProductApi';

    /**
     * Returns true if the given object is an instance of ProductApi.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProductApi {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProductApi.__pulumiType;
    }

    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly apiManagementName!: pulumi.Output<string>;
    /**
     * The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly apiName!: pulumi.Output<string>;
    /**
     * The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
     */
    public readonly productId!: pulumi.Output<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;

    /**
     * Create a ProductApi resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProductApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProductApiArgs | ProductApiState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ProductApiState | undefined;
            inputs["apiManagementName"] = state ? state.apiManagementName : undefined;
            inputs["apiName"] = state ? state.apiName : undefined;
            inputs["productId"] = state ? state.productId : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
        } else {
            const args = argsOrState as ProductApiArgs | undefined;
            if (!args || args.apiManagementName === undefined) {
                throw new Error("Missing required property 'apiManagementName'");
            }
            if (!args || args.apiName === undefined) {
                throw new Error("Missing required property 'apiName'");
            }
            if (!args || args.productId === undefined) {
                throw new Error("Missing required property 'productId'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["apiManagementName"] = args ? args.apiManagementName : undefined;
            inputs["apiName"] = args ? args.apiName : undefined;
            inputs["productId"] = args ? args.productId : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ProductApi.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProductApi resources.
 */
export interface ProductApiState {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName?: pulumi.Input<string>;
    /**
     * The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiName?: pulumi.Input<string>;
    /**
     * The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly productId?: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProductApi resource.
 */
export interface ProductApiArgs {
    /**
     * The name of the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiManagementName: pulumi.Input<string>;
    /**
     * The Name of the API Management API within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly apiName: pulumi.Input<string>;
    /**
     * The ID of the API Management Product within the API Management Service. Changing this forces a new resource to be created.
     */
    readonly productId: pulumi.Input<string>;
    /**
     * The name of the Resource Group in which the API Management Service exists. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
}
