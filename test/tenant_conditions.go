package test

import (
	"fmt"

	api "github.com/SAP/stewardci-core/pkg/apis/steward/v1alpha1"
	"github.com/SAP/stewardci-core/pkg/k8s"
)

// TenantCheck is a check for a tenant
type TenantCheck func(*api.Tenant) bool

// CreateTenantCondition creates a WaitCondition for a tenant with a dedicated check
func CreateTenantCondition(tenant *api.Tenant, check TenantCheck, desc string) WaitCondition {
	key := fmt.Sprintf("%s/%s", tenant.GetNamespace(), tenant.GetName())
	return NewWaitCondition(func(clientFactory k8s.ClientFactory) (bool, error) {
		fetcher := k8s.NewTenantFetcher(clientFactory)
		tenant, err := fetcher.ByKey(key)
		if err != nil {
			return true, err
		}
		return check(tenant), nil
	},
		fmt.Sprintf("TenantCondition_%s_%s", key, desc))
}

// TenantHasStateResult creates a TenantCheck which checks for a dedicated State
func TenantHasStateResult(result api.TenantResult) TenantCheck {
	return func(tenant *api.Tenant) bool {
		return tenant.Status.Result == result
	}
}
