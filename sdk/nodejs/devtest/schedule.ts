// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Manages automated startup and shutdown schedules for Azure Dev Test Lab.
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/r/dev_test_schedule.html.markdown.
 */
export class Schedule extends pulumi.CustomResource {
    /**
     * Get an existing Schedule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScheduleState, opts?: pulumi.CustomResourceOptions): Schedule {
        return new Schedule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure:devtest/schedule:Schedule';

    /**
     * Returns true if the given object is an instance of Schedule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Schedule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Schedule.__pulumiType;
    }

    public readonly dailyRecurrence!: pulumi.Output<outputs.devtest.ScheduleDailyRecurrence | undefined>;
    public readonly hourlyRecurrence!: pulumi.Output<outputs.devtest.ScheduleHourlyRecurrence | undefined>;
    /**
     * The name of the dev test lab. Changing this forces a new resource to be created.
     */
    public readonly labName!: pulumi.Output<string>;
    /**
     * The location where the schedule is created. Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly notificationSettings!: pulumi.Output<outputs.devtest.ScheduleNotificationSettings>;
    /**
     * The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
     */
    public readonly resourceGroupName!: pulumi.Output<string>;
    /**
     * The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * A mapping of tags to assign to the resource.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
     */
    public readonly taskType!: pulumi.Output<string>;
    /**
     * The time zone ID (e.g. Pacific Standard time).
     */
    public readonly timeZoneId!: pulumi.Output<string>;
    public readonly weeklyRecurrence!: pulumi.Output<outputs.devtest.ScheduleWeeklyRecurrence | undefined>;

    /**
     * Create a Schedule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScheduleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScheduleArgs | ScheduleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ScheduleState | undefined;
            inputs["dailyRecurrence"] = state ? state.dailyRecurrence : undefined;
            inputs["hourlyRecurrence"] = state ? state.hourlyRecurrence : undefined;
            inputs["labName"] = state ? state.labName : undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["notificationSettings"] = state ? state.notificationSettings : undefined;
            inputs["resourceGroupName"] = state ? state.resourceGroupName : undefined;
            inputs["status"] = state ? state.status : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["taskType"] = state ? state.taskType : undefined;
            inputs["timeZoneId"] = state ? state.timeZoneId : undefined;
            inputs["weeklyRecurrence"] = state ? state.weeklyRecurrence : undefined;
        } else {
            const args = argsOrState as ScheduleArgs | undefined;
            if (!args || args.labName === undefined) {
                throw new Error("Missing required property 'labName'");
            }
            if (!args || args.notificationSettings === undefined) {
                throw new Error("Missing required property 'notificationSettings'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.taskType === undefined) {
                throw new Error("Missing required property 'taskType'");
            }
            if (!args || args.timeZoneId === undefined) {
                throw new Error("Missing required property 'timeZoneId'");
            }
            inputs["dailyRecurrence"] = args ? args.dailyRecurrence : undefined;
            inputs["hourlyRecurrence"] = args ? args.hourlyRecurrence : undefined;
            inputs["labName"] = args ? args.labName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["notificationSettings"] = args ? args.notificationSettings : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["taskType"] = args ? args.taskType : undefined;
            inputs["timeZoneId"] = args ? args.timeZoneId : undefined;
            inputs["weeklyRecurrence"] = args ? args.weeklyRecurrence : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Schedule.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Schedule resources.
 */
export interface ScheduleState {
    readonly dailyRecurrence?: pulumi.Input<inputs.devtest.ScheduleDailyRecurrence>;
    readonly hourlyRecurrence?: pulumi.Input<inputs.devtest.ScheduleHourlyRecurrence>;
    /**
     * The name of the dev test lab. Changing this forces a new resource to be created.
     */
    readonly labName?: pulumi.Input<string>;
    /**
     * The location where the schedule is created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
     */
    readonly name?: pulumi.Input<string>;
    readonly notificationSettings?: pulumi.Input<inputs.devtest.ScheduleNotificationSettings>;
    /**
     * The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName?: pulumi.Input<string>;
    /**
     * The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
     */
    readonly taskType?: pulumi.Input<string>;
    /**
     * The time zone ID (e.g. Pacific Standard time).
     */
    readonly timeZoneId?: pulumi.Input<string>;
    readonly weeklyRecurrence?: pulumi.Input<inputs.devtest.ScheduleWeeklyRecurrence>;
}

/**
 * The set of arguments for constructing a Schedule resource.
 */
export interface ScheduleArgs {
    readonly dailyRecurrence?: pulumi.Input<inputs.devtest.ScheduleDailyRecurrence>;
    readonly hourlyRecurrence?: pulumi.Input<inputs.devtest.ScheduleHourlyRecurrence>;
    /**
     * The name of the dev test lab. Changing this forces a new resource to be created.
     */
    readonly labName: pulumi.Input<string>;
    /**
     * The location where the schedule is created. Changing this forces a new resource to be created.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the dev test lab schedule. Valid value for name depends on the `taskType`. For instance for taskType `LabVmsStartupTask` the name needs to be `LabVmAutoStart`.
     */
    readonly name?: pulumi.Input<string>;
    readonly notificationSettings: pulumi.Input<inputs.devtest.ScheduleNotificationSettings>;
    /**
     * The name of the resource group in which to create the dev test lab schedule. Changing this forces a new resource to be created.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The status of this schedule. Possible values are `Enabled` and `Disabled`. Defaults to `Disabled`.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * A mapping of tags to assign to the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The task type of the schedule. Possible values include `LabVmsShutdownTask` and `LabVmAutoStart`.
     */
    readonly taskType: pulumi.Input<string>;
    /**
     * The time zone ID (e.g. Pacific Standard time).
     */
    readonly timeZoneId: pulumi.Input<string>;
    readonly weeklyRecurrence?: pulumi.Input<inputs.devtest.ScheduleWeeklyRecurrence>;
}
