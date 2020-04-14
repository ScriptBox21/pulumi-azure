// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to obtain a Shared Access Signature (SAS Token) for an existing Storage Account Blob Container.
 * 
 * Shared access signatures allow fine-grained, ephemeral access control to various aspects of an Azure Storage Account Blob Container.
 * 
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-azurerm/blob/master/website/docs/d/storage_account_blob_container_sas.html.markdown.
 */
export function getAccountBlobContainerSAS(args: GetAccountBlobContainerSASArgs, opts?: pulumi.InvokeOptions): Promise<GetAccountBlobContainerSASResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure:storage/getAccountBlobContainerSAS:getAccountBlobContainerSAS", {
        "cacheControl": args.cacheControl,
        "connectionString": args.connectionString,
        "containerName": args.containerName,
        "contentDisposition": args.contentDisposition,
        "contentEncoding": args.contentEncoding,
        "contentLanguage": args.contentLanguage,
        "contentType": args.contentType,
        "expiry": args.expiry,
        "httpsOnly": args.httpsOnly,
        "ipAddress": args.ipAddress,
        "permissions": args.permissions,
        "start": args.start,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccountBlobContainerSAS.
 */
export interface GetAccountBlobContainerSASArgs {
    /**
     * The `Cache-Control` response header that is sent when this SAS token is used.
     */
    readonly cacheControl?: string;
    /**
     * The connection string for the storage account to which this SAS applies. Typically directly from the `primaryConnectionString` attribute of an `azure.storage.Account` resource.
     */
    readonly connectionString: string;
    /**
     * Name of the container.
     */
    readonly containerName: string;
    /**
     * The `Content-Disposition` response header that is sent when this SAS token is used.
     */
    readonly contentDisposition?: string;
    /**
     * The `Content-Encoding` response header that is sent when this SAS token is used.
     */
    readonly contentEncoding?: string;
    /**
     * The `Content-Language` response header that is sent when this SAS token is used.
     */
    readonly contentLanguage?: string;
    /**
     * The `Content-Type` response header that is sent when this SAS token is used.
     */
    readonly contentType?: string;
    /**
     * The expiration time and date of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    readonly expiry: string;
    /**
     * Only permit `https` access. If `false`, both `http` and `https` are permitted. Defaults to `true`.
     */
    readonly httpsOnly?: boolean;
    /**
     * Single ipv4 address or range (connected with a dash) of ipv4 addresses.
     */
    readonly ipAddress?: string;
    /**
     * A `permissions` block as defined below.
     */
    readonly permissions: inputs.storage.GetAccountBlobContainerSASPermissions;
    /**
     * The starting time and date of validity of this SAS. Must be a valid ISO-8601 format time/date string.
     */
    readonly start: string;
}

/**
 * A collection of values returned by getAccountBlobContainerSAS.
 */
export interface GetAccountBlobContainerSASResult {
    readonly cacheControl?: string;
    readonly connectionString: string;
    readonly containerName: string;
    readonly contentDisposition?: string;
    readonly contentEncoding?: string;
    readonly contentLanguage?: string;
    readonly contentType?: string;
    readonly expiry: string;
    readonly httpsOnly?: boolean;
    readonly ipAddress?: string;
    readonly permissions: outputs.storage.GetAccountBlobContainerSASPermissions;
    /**
     * The computed Blob Container Shared Access Signature (SAS).
     */
    readonly sas: string;
    readonly start: string;
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
