// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MediaServices.Inputs
{

    public sealed class AccountStorageAccountGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the ID of the Storage Account that will be associated with the Media Services instance.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Specifies whether the storage account should be the primary account or not. Defaults to `false`.
        /// </summary>
        [Input("isPrimary")]
        public Input<bool>? IsPrimary { get; set; }

        public AccountStorageAccountGetArgs()
        {
        }
    }
}
