// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./budgetResourceGroup";
export * from "./budgetSubscription";

// Import resources to register:
import { BudgetResourceGroup } from "./budgetResourceGroup";
import { BudgetSubscription } from "./budgetSubscription";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:consumption/budgetResourceGroup:BudgetResourceGroup":
                return new BudgetResourceGroup(name, <any>undefined, { urn })
            case "azure:consumption/budgetSubscription:BudgetSubscription":
                return new BudgetSubscription(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "consumption/budgetResourceGroup", _module)
pulumi.runtime.registerResourceModule("azure", "consumption/budgetSubscription", _module)
