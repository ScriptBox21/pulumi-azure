// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.DataFactory.Outputs
{

    [OutputType]
    public sealed class IntegrationRuntimeSsisCustomSetupScript
    {
        /// <summary>
        /// The blob endpoint for the container which contains a custom setup script that will be run on every node on startup. See [https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup](https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup) for more information.
        /// </summary>
        public readonly string BlobContainerUri;
        /// <summary>
        /// A container SAS token that gives access to the files. See [https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup](https://docs.microsoft.com/en-us/azure/data-factory/how-to-configure-azure-ssis-ir-custom-setup) for more information.
        /// </summary>
        public readonly string SasToken;

        [OutputConstructor]
        private IntegrationRuntimeSsisCustomSetupScript(
            string blobContainerUri,

            string sasToken)
        {
            BlobContainerUri = blobContainerUri;
            SasToken = sasToken;
        }
    }
}
