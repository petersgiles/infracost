package azure

import (
	"strings"

	"github.com/infracost/infracost/internal/resources/azure"
	"github.com/infracost/infracost/internal/schema"
)

func getAzureRMStorageAccountRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:      "azurerm_storage_account",
		CoreRFunc: newAzureRMStorageAccount,
		CustomRefIDFunc: func(d *schema.ResourceData) []string {
			return []string{d.Get("name").String()}
		},
	}
}

func newAzureRMStorageAccount(d *schema.ResourceData) schema.CoreResource {
	region := lookupRegion(d, []string{})

	accountKind := "StorageV2"
	if !d.IsEmpty("account_kind") {
		accountKind = d.Get("account_kind").String()
	}

	accountReplicationType := d.Get("account_replication_type").String()
	switch strings.ToLower(accountReplicationType) {
	case "ragrs":
		accountReplicationType = "RA-GRS"
	case "ragzrs":
		accountReplicationType = "RA-GZRS"
	}

	accountTier := d.Get("account_tier").String()

	accessTier := "Hot"
	if !d.IsEmpty("access_tier") {
		accessTier = d.Get("access_tier").String()
	}

	nfsv3 := false
	if !d.IsEmpty("nfsv3_enabled") {
		nfsv3 = d.Get("nfsv3_enabled").Bool()
	}

	return &azure.StorageAccount{
		Address:                d.Address,
		Region:                 region,
		AccessTier:             accessTier,
		AccountKind:            accountKind,
		AccountReplicationType: accountReplicationType,
		AccountTier:            accountTier,
		NFSv3:                  nfsv3,
	}
}
