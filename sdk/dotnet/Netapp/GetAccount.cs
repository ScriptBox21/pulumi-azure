// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.NetApp
{
    public static partial class Invokes
    {
        /// <summary>
        /// Uses this data source to access information about an existing NetApp Account.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/netapp_account.html.markdown.
        /// </summary>
        [Obsolete("Use GetAccount.InvokeAsync() instead")]
        public static Task<GetAccountResult> GetAccount(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure:netapp/getAccount:getAccount", args ?? InvokeArgs.Empty, options.WithVersion());
    }
    public static class GetAccount
    {
        /// <summary>
        /// Uses this data source to access information about an existing NetApp Account.
        /// 
        /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/netapp_account.html.markdown.
        /// </summary>
        public static Task<GetAccountResult> InvokeAsync(GetAccountArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAccountResult>("azure:netapp/getAccount:getAccount", args ?? InvokeArgs.Empty, options.WithVersion());
    }

    public sealed class GetAccountArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the NetApp Account.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Name of the Resource Group where the NetApp Account exists.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAccountArgs()
        {
        }
    }

    [OutputType]
    public sealed class GetAccountResult
    {
        /// <summary>
        /// The Azure Region where the NetApp Account exists.
        /// </summary>
        public readonly string Location;
        public readonly string Name;
        public readonly string ResourceGroupName;
        /// <summary>
        /// id is the provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetAccountResult(
            string location,
            string name,
            string resourceGroupName,
            string id)
        {
            Location = location;
            Name = name;
            ResourceGroupName = resourceGroupName;
            Id = id;
        }
    }
}
