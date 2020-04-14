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
    public sealed class LoggerEventhub
    {
        /// <summary>
        /// The connection string of an EventHub Namespace.
        /// </summary>
        public readonly string ConnectionString;
        /// <summary>
        /// The name of an EventHub.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private LoggerEventhub(
            string connectionString,

            string name)
        {
            ConnectionString = connectionString;
            Name = name;
        }
    }
}
