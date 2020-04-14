// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.MariaDB.Inputs
{

    public sealed class ServerStorageProfileGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether autogrow is enabled or disabled for the storage. Valid values are `Enabled` or `Disabled`.
        /// </summary>
        [Input("autoGrow")]
        public Input<string>? AutoGrow { get; set; }

        /// <summary>
        /// Backup retention days for the server, supported values are between `7` and `35` days.
        /// </summary>
        [Input("backupRetentionDays")]
        public Input<int>? BackupRetentionDays { get; set; }

        /// <summary>
        /// Enable Geo-redundant or not for server backup. Valid values for this property are `Enabled` or `Disabled`. Changing this forces a new resource to be created.
        /// </summary>
        [Input("geoRedundantBackup")]
        public Input<string>? GeoRedundantBackup { get; set; }

        /// <summary>
        /// Max storage allowed for a server. Possible values are between `5120` MB (5GB) and `1024000`MB (1TB) for the Basic SKU and between `5120` MB (5GB) and `4096000` MB (4TB) for General Purpose/Memory Optimized SKUs. For more information see the [product documentation](https://docs.microsoft.com/en-us/rest/api/mariadb/servers/create#storageprofile).
        /// </summary>
        [Input("storageMb", required: true)]
        public Input<int> StorageMb { get; set; } = null!;

        public ServerStorageProfileGetArgs()
        {
        }
    }
}
