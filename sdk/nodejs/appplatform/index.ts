// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./getSpringCloudService";
export * from "./springCloudActiveDeployment";
export * from "./springCloudApp";
export * from "./springCloudCertificate";
export * from "./springCloudJavaDeployment";
export * from "./springCloudService";

// Import resources to register:
import { SpringCloudActiveDeployment } from "./springCloudActiveDeployment";
import { SpringCloudApp } from "./springCloudApp";
import { SpringCloudCertificate } from "./springCloudCertificate";
import { SpringCloudJavaDeployment } from "./springCloudJavaDeployment";
import { SpringCloudService } from "./springCloudService";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure:appplatform/springCloudActiveDeployment:SpringCloudActiveDeployment":
                return new SpringCloudActiveDeployment(name, <any>undefined, { urn })
            case "azure:appplatform/springCloudApp:SpringCloudApp":
                return new SpringCloudApp(name, <any>undefined, { urn })
            case "azure:appplatform/springCloudCertificate:SpringCloudCertificate":
                return new SpringCloudCertificate(name, <any>undefined, { urn })
            case "azure:appplatform/springCloudJavaDeployment:SpringCloudJavaDeployment":
                return new SpringCloudJavaDeployment(name, <any>undefined, { urn })
            case "azure:appplatform/springCloudService:SpringCloudService":
                return new SpringCloudService(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure", "appplatform/springCloudActiveDeployment", _module)
pulumi.runtime.registerResourceModule("azure", "appplatform/springCloudApp", _module)
pulumi.runtime.registerResourceModule("azure", "appplatform/springCloudCertificate", _module)
pulumi.runtime.registerResourceModule("azure", "appplatform/springCloudJavaDeployment", _module)
pulumi.runtime.registerResourceModule("azure", "appplatform/springCloudService", _module)
