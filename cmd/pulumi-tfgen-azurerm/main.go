// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package main

import (
	azure "github.com/pulumi/pulumi-azure"
	"github.com/pulumi/pulumi-azure/pkg/version"
	"github.com/pulumi/pulumi-terraform/pkg/tfgen"
)

func main() {
	tfgen.Main("azurerm", version.Version, azure.Provider())
}