//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchangeanalysis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"
)

// x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesByResourceGroup.json
func ExampleChangesClient_ListChangesByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armchangeanalysis.NewChangesClient("<subscription-id>", cred, nil)
	pager := client.ListChangesByResourceGroup("<resource-group-name>",
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-25T12:09:03.141Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T12:09:03.141Z"); return t }(),
		&armchangeanalysis.ChangesListChangesByResourceGroupOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Change.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesBySubscription.json
func ExampleChangesClient_ListChangesBySubscription() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armchangeanalysis.NewChangesClient("<subscription-id>", cred, nil)
	pager := client.ListChangesBySubscription(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-25T12:09:03.141Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T12:09:03.141Z"); return t }(),
		&armchangeanalysis.ChangesListChangesBySubscriptionOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Change.ID: %s\n", *v.ID)
		}
	}
}
