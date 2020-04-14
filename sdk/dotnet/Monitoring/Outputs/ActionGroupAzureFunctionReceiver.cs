// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.Monitoring.Outputs
{

    [OutputType]
    public sealed class ActionGroupAzureFunctionReceiver
    {
        public readonly string FunctionAppResourceId;
        /// <summary>
        /// The function name in the function app.
        /// </summary>
        public readonly string FunctionName;
        /// <summary>
        /// The http trigger url where http request sent to.
        /// </summary>
        public readonly string HttpTriggerUrl;
        /// <summary>
        /// The name of the Azure Function receiver.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Enables or disables the common alert schema.
        /// </summary>
        public readonly bool? UseCommonAlertSchema;

        [OutputConstructor]
        private ActionGroupAzureFunctionReceiver(
            string functionAppResourceId,

            string functionName,

            string httpTriggerUrl,

            string name,

            bool? useCommonAlertSchema)
        {
            FunctionAppResourceId = functionAppResourceId;
            FunctionName = functionName;
            HttpTriggerUrl = httpTriggerUrl;
            Name = name;
            UseCommonAlertSchema = useCommonAlertSchema;
        }
    }
}
