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
    public sealed class FunctionAppSiteCredential
    {
        /// <summary>
        /// The password associated with the username, which can be used to publish to this App Service.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// The username which can be used to publish to this App Service
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private FunctionAppSiteCredential(
            string? password,

            string? username)
        {
            Password = password;
            Username = username;
        }
    }
}
