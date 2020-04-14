// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.ApiManagement.Outputs
{

    [OutputType]
    public sealed class ApiOperationRequestHeader
    {
        /// <summary>
        /// The default value for this Header.
        /// </summary>
        public readonly string? DefaultValue;
        /// <summary>
        /// A description of this Header.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The Name of this Header.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Is this Header Required?
        /// </summary>
        public readonly bool Required;
        /// <summary>
        /// The Type of this Header, such as a `string`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// One or more acceptable values for this Header.
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private ApiOperationRequestHeader(
            string? defaultValue,

            string? description,

            string name,

            bool required,

            string type,

            ImmutableArray<string> values)
        {
            DefaultValue = defaultValue;
            Description = description;
            Name = name;
            Required = required;
            Type = type;
            Values = values;
        }
    }
}
