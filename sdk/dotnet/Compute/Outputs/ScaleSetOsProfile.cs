// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Compute.Outputs
{

    [OutputType]
    public sealed class ScaleSetOsProfile
    {
        /// <summary>
        /// Specifies the administrator password to use for all the instances of virtual machines in a scale set.
        /// </summary>
        public readonly string? AdminPassword;
        /// <summary>
        /// Specifies the administrator account name to use for all the instances of virtual machines in the scale set.
        /// </summary>
        public readonly string AdminUsername;
        /// <summary>
        /// Specifies the computer name prefix for all of the virtual machines in the scale set. Computer name prefixes must be 1 to 9 characters long for windows images and 1 - 58 for linux. Changing this forces a new resource to be created.
        /// </summary>
        public readonly string ComputerNamePrefix;
        /// <summary>
        /// Specifies custom data to supply to the machine. On linux-based systems, this can be used as a cloud-init script. On other systems, this will be copied as a file on disk. Internally, this provider will base64 encode this value before sending it to the API. The maximum length of the binary array is 65535 bytes.
        /// </summary>
        public readonly string? CustomData;

        [OutputConstructor]
        private ScaleSetOsProfile(
            string? adminPassword,

            string adminUsername,

            string computerNamePrefix,

            string? customData)
        {
            AdminPassword = adminPassword;
            AdminUsername = adminUsername;
            ComputerNamePrefix = computerNamePrefix;
            CustomData = customData;
        }
    }
}
