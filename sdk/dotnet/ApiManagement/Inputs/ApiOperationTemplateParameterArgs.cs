// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Inputs
{

    public sealed class ApiOperationTemplateParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default value for this Template Parameter.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// A description of this Template Parameter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Name of this Template Parameter.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Is this Template Parameter Required?
        /// </summary>
        [Input("required", required: true)]
        public Input<bool> Required { get; set; } = null!;

        /// <summary>
        /// The Type of this Template Parameter, such as a `string`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// One or more acceptable values for this Template Parameter.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ApiOperationTemplateParameterArgs()
        {
        }
    }
}
