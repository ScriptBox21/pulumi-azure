// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages a Kusto Cluster Principal Assignment.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure from "@pulumi/azure";
 *
 * const current = azure.core.getClientConfig({});
 * const rg = new azure.core.ResourceGroup("rg", {location: "East US"});
 * const exampleCluster = new azure.kusto.Cluster("exampleCluster", {
 *     location: rg.location,
 *     resourceGroupName: rg.name,
 *     sku: {
 *         name: "Standard_D13_v2",
 *         capacity: 2,
 *     },
 * });
 * const exampleClusterPrincipalAssignment = new azure.kusto.ClusterPrincipalAssignment("exampleClusterPrincipalAssignment", {
 *     resourceGroupName: rg.name,
 *     clusterName: exampleCluster.name,
 *     tenantId: current.then(current => current.tenantId),
 *     principalId: current.then(current => current.clientId),
 *     principalType: "App",
 *     role: "AllDatabasesAdmin",
 * });
 * ```
 *
 * ## Import
 *
 * Data Explorer Cluster Principal Assignments can be imported using the `resource id`, e.g.
 *
 * ```sh
 *  $ pulumi import azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Kusto/Clusters/cluster1/PrincipalAssignments/assignment1
 * ```
 */
export class ClusterPrincipalAssignment extends pulumi.CustomResource {
    /**
     * Get an existing ClusterPrincipalAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClusterPrincipalAssignmentState, opts?: pulumi.CustomResourceOptions): ClusterPrincipalAssignment {
        return new ClusterPrincipalAssignment(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:kusto/clusterPrincipalAssignment:ClusterPrincipalAssignment';

    /**
     * Returns true if the given object is an instance of ClusterPrincipalAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClusterPrincipalAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClusterPrincipalAssignment.__pulumiType;
    }

    /**
     * The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
     */
    public readonly clusterName!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * The object id of the principal. Changing this forces a new resource to be created.
     */
    public readonly principalId!: pulumi.Output<string>;
    /**
     * The name of the principal.
     */
    public /*out*/ readonly principalName!: pulumi.Output<string>;
    /**
     * The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
     */
    public readonly principalType!: pulumi.Output<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The tenant id in which the principal resides. Changing this forces a new resource to be created.
     */
    public readonly tenantId!: pulumi.Output<string>;
    /**
     * The name of the tenant.
     */
    public /*out*/ readonly tenantName!: pulumi.Output<string>;

    /**
     * Create a ClusterPrincipalAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClusterPrincipalAssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClusterPrincipalAssignmentArgs | ClusterPrincipalAssignmentState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClusterPrincipalAssignmentState | undefined;
            inputs["clusterName"] = state ? state.clusterName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["principalId"] = state ? state.principalId : undefined;
            inputs["principalName"] = state ? state.principalName : undefined;
            inputs["principalType"] = state ? state.principalType : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
            inputs["tenantName"] = state ? state.tenantName : undefined;
        } else {
            const args = argsOrState as ClusterPrincipalAssignmentArgs | undefined;
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.principalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalId'");
            }
            if ((!args || args.principalType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'principalType'");
            }
            if ((!args || args.resourceGroupName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["principalId"] = args ? args.principalId : undefined;
            inputs["principalType"] = args ? args.principalType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
            inputs["principalName"] = undefined /*out*/;
            inputs["tenantName"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ClusterPrincipalAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClusterPrincipalAssignment resources.
 */
export interface ClusterPrincipalAssignmentState {
    /**
     * The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly clusterName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The object id of the principal. Changing this forces a new resource to be created.
     */
    readonly principalId?: pulumi.Input<string>;
    /**
     * The name of the principal.
     */
    readonly principalName?: pulumi.Input<string>;
    /**
     * The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
     */
    readonly principalType?: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
     */
    readonly role?: pulumi.Input<string>;
    /**
     * The tenant id in which the principal resides. Changing this forces a new resource to be created.
     */
    readonly tenantId?: pulumi.Input<string>;
    /**
     * The name of the tenant.
     */
    readonly tenantName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClusterPrincipalAssignment resource.
 */
export interface ClusterPrincipalAssignmentArgs {
    /**
     * The name of the cluster in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly clusterName: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    /**
     * The object id of the principal. Changing this forces a new resource to be created.
     */
    readonly principalId: pulumi.Input<string>;
    /**
     * The type of the principal. Valid values include `App`, `Group`, `User`. Changing this forces a new resource to be created.
     */
    readonly principalType: pulumi.Input<string>;
    /**
     * The name of the resource group in which to create the resource. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The cluster role assigned to the principal. Valid values include `AllDatabasesAdmin` and `AllDatabasesViewer`. Changing this forces a new resource to be created.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The tenant id in which the principal resides. Changing this forces a new resource to be created.
     */
    readonly tenantId: pulumi.Input<string>;
}
