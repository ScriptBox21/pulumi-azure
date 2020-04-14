// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.AppService.Outputs
{

    [OutputType]
    public sealed class AppServiceConnectionString
    {
        /// <summary>
        /// The name of the Connection String.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The type of the Connection String. Possible values are `APIHub`, `Custom`, `DocDb`, `EventHub`, `MySQL`, `NotificationHub`, `PostgreSQL`, `RedisCache`, `ServiceBus`, `SQLAzure` and  `SQLServer`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The value for the Connection String.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private AppServiceConnectionString(
            string name,

            string type,

            string value)
        {
            Name = name;
            Type = type;
            Value = value;
        }
    }
}
