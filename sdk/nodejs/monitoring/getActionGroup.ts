// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to access the properties of an Action Group.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/monitor_action_group.html.markdown.
 */
export function getActionGroup(args: GetActionGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetActionGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:monitoring/getActionGroup:getActionGroup", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking getActionGroup.
 */
export interface GetActionGroupArgs {
    /**
     * Specifies the name of the Action Group.
     */
    readonly name: string;
    /**
     * Specifies the name of the resource group the Action Group is located in.
     */
    readonly resourceGroupName: string;
}

/**
 * A collection of values returned by getActionGroup.
 */
export interface GetActionGroupResult {
    /**
     * One or more `armRoleReceiver` blocks as defined below.
     */
    readonly armRoleReceivers: outputs.monitoring.GetActionGroupArmRoleReceiver[];
    /**
     * One or more `automationRunbookReceiver` blocks as defined below.
     */
    readonly automationRunbookReceivers: outputs.monitoring.GetActionGroupAutomationRunbookReceiver[];
    /**
     * One or more `azureAppPushReceiver` blocks as defined below.
     */
    readonly azureAppPushReceivers: outputs.monitoring.GetActionGroupAzureAppPushReceiver[];
    /**
     * One or more `azureFunctionReceiver` blocks as defined below.
     */
    readonly azureFunctionReceivers: outputs.monitoring.GetActionGroupAzureFunctionReceiver[];
    /**
     * One or more `emailReceiver` blocks as defined below.
     */
    readonly emailReceivers: outputs.monitoring.GetActionGroupEmailReceiver[];
    /**
     * Whether this action group is enabled.
     */
    readonly enabled: boolean;
    /**
     * One or more `itsmReceiver` blocks as defined below.
     */
    readonly itsmReceivers: outputs.monitoring.GetActionGroupItsmReceiver[];
    /**
     * One or more `logicAppReceiver` blocks as defined below.
     */
    readonly logicAppReceivers: outputs.monitoring.GetActionGroupLogicAppReceiver[];
    /**
     * The name of the webhook receiver.
     */
    readonly name: string;
    readonly resourceGroupName: string;
    /**
     * The short name of the action group.
     */
    readonly shortName: string;
    /**
     * One or more `smsReceiver` blocks as defined below.
     */
    readonly smsReceivers: outputs.monitoring.GetActionGroupSmsReceiver[];
    /**
     * One or more `voiceReceiver` blocks as defined below.
     */
    readonly voiceReceivers: outputs.monitoring.GetActionGroupVoiceReceiver[];
    /**
     * One or more `webhookReceiver` blocks as defined below.
     */
    readonly webhookReceivers: outputs.monitoring.GetActionGroupWebhookReceiver[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
