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
    public sealed class VirtualMachineOsProfileLinuxConfig
    {
        /// <summary>
        /// Specifies whether password authentication should be disabled. If set to `false`, an `admin_password` must be specified.
        /// </summary>
        public readonly bool DisablePasswordAuthentication;
        /// <summary>
        /// One or more `ssh_keys` blocks. This field is required if `disable_password_authentication` is set to `true`.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineOsProfileLinuxConfigSshKey> SshKeys;

        [OutputConstructor]
        private VirtualMachineOsProfileLinuxConfig(
            bool disablePasswordAuthentication,

            ImmutableArray<Outputs.VirtualMachineOsProfileLinuxConfigSshKey> sshKeys)
        {
            DisablePasswordAuthentication = disablePasswordAuthentication;
            SshKeys = sshKeys;
        }
    }
}
