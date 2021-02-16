// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Iot
{
    /// <summary>
    /// Manages a Iot Security Device Group.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Azure = Pulumi.Azure;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleResourceGroup = new Azure.Core.ResourceGroup("exampleResourceGroup", new Azure.Core.ResourceGroupArgs
    ///         {
    ///             Location = "West Europe",
    ///         });
    ///         var exampleIoTHub = new Azure.Iot.IoTHub("exampleIoTHub", new Azure.Iot.IoTHubArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             Sku = new Azure.Iot.Inputs.IoTHubSkuArgs
    ///             {
    ///                 Name = "S1",
    ///                 Capacity = 1,
    ///             },
    ///         });
    ///         var exampleSecuritySolution = new Azure.Iot.SecuritySolution("exampleSecuritySolution", new Azure.Iot.SecuritySolutionArgs
    ///         {
    ///             ResourceGroupName = exampleResourceGroup.Name,
    ///             Location = exampleResourceGroup.Location,
    ///             DisplayName = "Iot Security Solution",
    ///             IothubIds = 
    ///             {
    ///                 exampleIoTHub.Id,
    ///             },
    ///         });
    ///         var exampleSecurityDeviceGroup = new Azure.Iot.SecurityDeviceGroup("exampleSecurityDeviceGroup", new Azure.Iot.SecurityDeviceGroupArgs
    ///         {
    ///             IothubId = exampleIoTHub.Id,
    ///             AllowRule = new Azure.Iot.Inputs.SecurityDeviceGroupAllowRuleArgs
    ///             {
    ///                 ConnectionToIpNotAlloweds = 
    ///                 {
    ///                     "10.0.0.0/24",
    ///                 },
    ///             },
    ///             RangeRules = 
    ///             {
    ///                 new Azure.Iot.Inputs.SecurityDeviceGroupRangeRuleArgs
    ///                 {
    ///                     Type = "ActiveConnectionsNotInAllowedRange",
    ///                     Min = 0,
    ///                     Max = 30,
    ///                     Duration = "PT5M",
    ///                 },
    ///             },
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 exampleSecuritySolution,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Iot Security Device Group can be imported using the `resource id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import azure:iot/securityDeviceGroup:SecurityDeviceGroup example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.Devices/iotHubs/hub1/providers/Microsoft.Security/deviceSecurityGroups/group1
    /// ```
    /// </summary>
    public partial class SecurityDeviceGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// an `allow_rule` blocks as defined below.
        /// </summary>
        [Output("allowRule")]
        public Output<Outputs.SecurityDeviceGroupAllowRule?> AllowRule { get; private set; } = null!;

        /// <summary>
        /// The ID of the IoT Hub which to link the Security Device Group to. Changing this forces a new resource to be created.
        /// </summary>
        [Output("iothubId")]
        public Output<string> IothubId { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the Device Security Group. Changing this forces a new resource to be created.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// One or more `range_rule` blocks as defined below.
        /// </summary>
        [Output("rangeRules")]
        public Output<ImmutableArray<Outputs.SecurityDeviceGroupRangeRule>> RangeRules { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityDeviceGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityDeviceGroup(string name, SecurityDeviceGroupArgs args, CustomResourceOptions? options = null)
            : base("azure:iot/securityDeviceGroup:SecurityDeviceGroup", name, args ?? new SecurityDeviceGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityDeviceGroup(string name, Input<string> id, SecurityDeviceGroupState? state = null, CustomResourceOptions? options = null)
            : base("azure:iot/securityDeviceGroup:SecurityDeviceGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityDeviceGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityDeviceGroup Get(string name, Input<string> id, SecurityDeviceGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new SecurityDeviceGroup(name, id, state, options);
        }
    }

    public sealed class SecurityDeviceGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// an `allow_rule` blocks as defined below.
        /// </summary>
        [Input("allowRule")]
        public Input<Inputs.SecurityDeviceGroupAllowRuleArgs>? AllowRule { get; set; }

        /// <summary>
        /// The ID of the IoT Hub which to link the Security Device Group to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubId", required: true)]
        public Input<string> IothubId { get; set; } = null!;

        /// <summary>
        /// Specifies the name of the Device Security Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rangeRules")]
        private InputList<Inputs.SecurityDeviceGroupRangeRuleArgs>? _rangeRules;

        /// <summary>
        /// One or more `range_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.SecurityDeviceGroupRangeRuleArgs> RangeRules
        {
            get => _rangeRules ?? (_rangeRules = new InputList<Inputs.SecurityDeviceGroupRangeRuleArgs>());
            set => _rangeRules = value;
        }

        public SecurityDeviceGroupArgs()
        {
        }
    }

    public sealed class SecurityDeviceGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// an `allow_rule` blocks as defined below.
        /// </summary>
        [Input("allowRule")]
        public Input<Inputs.SecurityDeviceGroupAllowRuleGetArgs>? AllowRule { get; set; }

        /// <summary>
        /// The ID of the IoT Hub which to link the Security Device Group to. Changing this forces a new resource to be created.
        /// </summary>
        [Input("iothubId")]
        public Input<string>? IothubId { get; set; }

        /// <summary>
        /// Specifies the name of the Device Security Group. Changing this forces a new resource to be created.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("rangeRules")]
        private InputList<Inputs.SecurityDeviceGroupRangeRuleGetArgs>? _rangeRules;

        /// <summary>
        /// One or more `range_rule` blocks as defined below.
        /// </summary>
        public InputList<Inputs.SecurityDeviceGroupRangeRuleGetArgs> RangeRules
        {
            get => _rangeRules ?? (_rangeRules = new InputList<Inputs.SecurityDeviceGroupRangeRuleGetArgs>());
            set => _rangeRules = value;
        }

        public SecurityDeviceGroupState()
        {
        }
    }
}
