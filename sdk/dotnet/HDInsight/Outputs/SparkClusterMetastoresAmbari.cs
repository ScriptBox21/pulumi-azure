// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Azure.HDInsight.Outputs
{

    [OutputType]
    public sealed class SparkClusterMetastoresAmbari
    {
        /// <summary>
        /// The external Hive metastore's existing SQL database.  Changing this forces a new resource to be created.
        /// </summary>
        public readonly string DatabaseName;
        /// <summary>
        /// The external Ambari metastore's existing SQL server admin password.  Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The fully-qualified domain name (FQDN) of the SQL server to use for the external Ambari metastore.  Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Server;
        /// <summary>
        /// The external Ambari metastore's existing SQL server admin username.  Changing this forces a new resource to be created.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private SparkClusterMetastoresAmbari(
            string databaseName,

            string password,

            string server,

            string username)
        {
            DatabaseName = databaseName;
            Password = password;
            Server = server;
            Username = username;
        }
    }
}
