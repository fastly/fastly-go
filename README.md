# fastly-go

A Go client library for interacting with most facets of the [Fastly API](https://www.fastly.com/documentation/reference/api/).

> [!WARNING]
> This API client is auto-generated from Fastly's OpenAPI specification and may not function correctly when used on the Compute platform. Support for Compute is on the roadmap but has not yet been prioritized.

> [!TIP]
> If you'd like to use the hand-written API client instead, see [go-fastly](https://github.com/fastly/go-fastly).

## Requirements

Go version 1.18

## Installation

Add the following to your project's `go.mod`:

```go.mod
require (
	github.com/fastly/fastly-go 1.0.0-beta.35
)
```

## Usage

> [!NOTE]
> The Fastly API requires an [API token](https://www.fastly.com/documentation/reference/api/#authentication) for most operations.

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fastly/fastly-go/fastly"
)

func main() {
  cfg := fastly.NewConfiguration()
  apiClient := fastly.NewAPIClient(cfg)
  ctx := fastly.NewAPIKeyContextFromEnv("FASTLY_API_TOKEN")

  req := apiClient.ServiceAPI.CreateService(ctx)

  res, httpResp, err := req.Comment("comment_example").Name("name_example").CustomerID(os.Getenv("FASTLY_CUSTOMER_ID")).ResourceType("vcl").Execute()
  if err != nil {
    fmt.Fprintf(os.Stderr, "Error when calling `ServiceAPI.CreateService`: %+v\n", err)
    fmt.Fprintf(os.Stderr, "Full HTTP response: %+v\n", httpResp)
    return
  }

  fmt.Fprintf(os.Stdout, "Response from `ServiceAPI.CreateService`:\n\n%+v\n\n", res)

  serviceID = *res.ID
  fmt.Printf("Service Name: %s\nService ID: %s\n\n", *res.Name, serviceID)
}
```

## API Endpoints

The main documentation for the Fastly API can be found on our [Developer Hub](https://www.fastly.com/documentation/reference/api/).

<details>

<summary>Table of API endpoints</summary>

Class | Method | Description
----- | ------ | -----------
*ACLAPI* | [**CreateACL**](docs/AclAPI.md#createacl) | Create a new ACL
*ACLAPI* | [**DeleteACL**](docs/AclAPI.md#deleteacl) | Delete an ACL
*ACLAPI* | [**GetACL**](docs/AclAPI.md#getacl) | Describe an ACL
*ACLAPI* | [**ListACLs**](docs/AclAPI.md#listacls) | List ACLs
*ACLAPI* | [**UpdateACL**](docs/AclAPI.md#updateacl) | Update an ACL
*ACLEntryAPI* | [**BulkUpdateACLEntries**](docs/AclEntryAPI.md#bulkupdateaclentries) | Update multiple ACL entries
*ACLEntryAPI* | [**CreateACLEntry**](docs/AclEntryAPI.md#createaclentry) | Create an ACL entry
*ACLEntryAPI* | [**DeleteACLEntry**](docs/AclEntryAPI.md#deleteaclentry) | Delete an ACL entry
*ACLEntryAPI* | [**GetACLEntry**](docs/AclEntryAPI.md#getaclentry) | Describe an ACL entry
*ACLEntryAPI* | [**ListACLEntries**](docs/AclEntryAPI.md#listaclentries) | List ACL entries
*ACLEntryAPI* | [**UpdateACLEntry**](docs/AclEntryAPI.md#updateaclentry) | Update an ACL entry
*ACLsInComputeAPI* | [**ComputeACLCreateACLs**](docs/AclsInComputeAPI.md#computeaclcreateacls) | Create a new ACL
*ACLsInComputeAPI* | [**ComputeACLDeleteSACLID**](docs/AclsInComputeAPI.md#computeacldeletesaclid) | Delete an ACL
*ACLsInComputeAPI* | [**ComputeACLListACLEntries**](docs/AclsInComputeAPI.md#computeacllistaclentries) | List an ACL
*ACLsInComputeAPI* | [**ComputeACLListACLs**](docs/AclsInComputeAPI.md#computeacllistacls) | List ACLs
*ACLsInComputeAPI* | [**ComputeACLListACLsSAclID**](docs/AclsInComputeAPI.md#computeacllistaclssaclid) | Describe an ACL
*ACLsInComputeAPI* | [**ComputeACLLookupACLs**](docs/AclsInComputeAPI.md#computeacllookupacls) | Lookup an ACL
*ACLsInComputeAPI* | [**ComputeACLUpdateACLs**](docs/AclsInComputeAPI.md#computeaclupdateacls) | Update an ACL
*ApexRedirectAPI* | [**CreateApexRedirect**](docs/ApexRedirectAPI.md#createapexredirect) | Create an apex redirect
*ApexRedirectAPI* | [**DeleteApexRedirect**](docs/ApexRedirectAPI.md#deleteapexredirect) | Delete an apex redirect
*ApexRedirectAPI* | [**GetApexRedirect**](docs/ApexRedirectAPI.md#getapexredirect) | Get an apex redirect
*ApexRedirectAPI* | [**ListApexRedirects**](docs/ApexRedirectAPI.md#listapexredirects) | List apex redirects
*ApexRedirectAPI* | [**UpdateApexRedirect**](docs/ApexRedirectAPI.md#updateapexredirect) | Update an apex redirect
*AutomationTokensAPI* | [**CreateAutomationToken**](docs/AutomationTokensAPI.md#createautomationtoken) | Create Automation Token
*AutomationTokensAPI* | [**GetAutomationTokenID**](docs/AutomationTokensAPI.md#getautomationtokenid) | Retrieve an Automation Token by ID
*AutomationTokensAPI* | [**GetAutomationTokensIDServices**](docs/AutomationTokensAPI.md#getautomationtokensidservices) | List Automation Token Services
*AutomationTokensAPI* | [**ListAutomationTokens**](docs/AutomationTokensAPI.md#listautomationtokens) | List Customer Automation Tokens
*AutomationTokensAPI* | [**RevokeAutomationTokenID**](docs/AutomationTokensAPI.md#revokeautomationtokenid) | Revoke an Automation Token by ID
*BackendAPI* | [**CreateBackend**](docs/BackendAPI.md#createbackend) | Create a backend
*BackendAPI* | [**DeleteBackend**](docs/BackendAPI.md#deletebackend) | Delete a backend
*BackendAPI* | [**GetBackend**](docs/BackendAPI.md#getbackend) | Describe a backend
*BackendAPI* | [**ListBackends**](docs/BackendAPI.md#listbackends) | List backends
*BackendAPI* | [**UpdateBackend**](docs/BackendAPI.md#updatebackend) | Update a backend
*BillingAPI* | [**GetInvoice**](docs/BillingAPI.md#getinvoice) | Get an invoice
*BillingAPI* | [**GetInvoiceByID**](docs/BillingAPI.md#getinvoicebyid) | Get an invoice
*BillingAPI* | [**GetInvoiceMtd**](docs/BillingAPI.md#getinvoicemtd) | Get month-to-date billing estimate
*BillingAddressAPI* | [**AddBillingAddr**](docs/BillingAddressAPI.md#addbillingaddr) | Add a billing address to a customer
*BillingAddressAPI* | [**DeleteBillingAddr**](docs/BillingAddressAPI.md#deletebillingaddr) | Delete a billing address
*BillingAddressAPI* | [**GetBillingAddr**](docs/BillingAddressAPI.md#getbillingaddr) | Get a billing address
*BillingAddressAPI* | [**UpdateBillingAddr**](docs/BillingAddressAPI.md#updatebillingaddr) | Update a billing address
*BillingInvoicesAPI* | [**GetInvoiceByInvoiceID**](docs/BillingInvoicesAPI.md#getinvoicebyinvoiceid) | Get invoice by ID.
*BillingInvoicesAPI* | [**GetMonthToDateInvoice**](docs/BillingInvoicesAPI.md#getmonthtodateinvoice) | Get month-to-date invoice.
*BillingInvoicesAPI* | [**ListInvoices**](docs/BillingInvoicesAPI.md#listinvoices) | List of invoices.
*BillingUsageMetricsAPI* | [**GetServiceLevelUsage**](docs/BillingUsageMetricsAPI.md#getservicelevelusage) | Retrieve service-level usage metrics for a product.
*BillingUsageMetricsAPI* | [**GetServiceLevelUsageTypes**](docs/BillingUsageMetricsAPI.md#getservicelevelusagetypes) | Retrieve product usage types for a customer.
*CacheSettingsAPI* | [**CreateCacheSettings**](docs/CacheSettingsAPI.md#createcachesettings) | Create a cache settings object
*CacheSettingsAPI* | [**DeleteCacheSettings**](docs/CacheSettingsAPI.md#deletecachesettings) | Delete a cache settings object
*CacheSettingsAPI* | [**GetCacheSettings**](docs/CacheSettingsAPI.md#getcachesettings) | Get a cache settings object
*CacheSettingsAPI* | [**ListCacheSettings**](docs/CacheSettingsAPI.md#listcachesettings) | List cache settings objects
*CacheSettingsAPI* | [**UpdateCacheSettings**](docs/CacheSettingsAPI.md#updatecachesettings) | Update a cache settings object
*ConditionAPI* | [**CreateCondition**](docs/ConditionAPI.md#createcondition) | Create a condition
*ConditionAPI* | [**DeleteCondition**](docs/ConditionAPI.md#deletecondition) | Delete a condition
*ConditionAPI* | [**GetCondition**](docs/ConditionAPI.md#getcondition) | Describe a condition
*ConditionAPI* | [**ListConditions**](docs/ConditionAPI.md#listconditions) | List conditions
*ConditionAPI* | [**UpdateCondition**](docs/ConditionAPI.md#updatecondition) | Update a condition
*ConfigStoreAPI* | [**CreateConfigStore**](docs/ConfigStoreAPI.md#createconfigstore) | Create a config store
*ConfigStoreAPI* | [**DeleteConfigStore**](docs/ConfigStoreAPI.md#deleteconfigstore) | Delete a config store
*ConfigStoreAPI* | [**GetConfigStore**](docs/ConfigStoreAPI.md#getconfigstore) | Describe a config store
*ConfigStoreAPI* | [**GetConfigStoreInfo**](docs/ConfigStoreAPI.md#getconfigstoreinfo) | Get config store metadata
*ConfigStoreAPI* | [**ListConfigStoreServices**](docs/ConfigStoreAPI.md#listconfigstoreservices) | List linked services
*ConfigStoreAPI* | [**ListConfigStores**](docs/ConfigStoreAPI.md#listconfigstores) | List config stores
*ConfigStoreAPI* | [**UpdateConfigStore**](docs/ConfigStoreAPI.md#updateconfigstore) | Update a config store
*ConfigStoreItemAPI* | [**BulkUpdateConfigStoreItem**](docs/ConfigStoreItemAPI.md#bulkupdateconfigstoreitem) | Update multiple entries in a config store
*ConfigStoreItemAPI* | [**CreateConfigStoreItem**](docs/ConfigStoreItemAPI.md#createconfigstoreitem) | Create an entry in a config store
*ConfigStoreItemAPI* | [**DeleteConfigStoreItem**](docs/ConfigStoreItemAPI.md#deleteconfigstoreitem) | Delete an item from a config store
*ConfigStoreItemAPI* | [**GetConfigStoreItem**](docs/ConfigStoreItemAPI.md#getconfigstoreitem) | Get an item from a config store
*ConfigStoreItemAPI* | [**ListConfigStoreItems**](docs/ConfigStoreItemAPI.md#listconfigstoreitems) | List items in a config store
*ConfigStoreItemAPI* | [**UpdateConfigStoreItem**](docs/ConfigStoreItemAPI.md#updateconfigstoreitem) | Update an entry in a config store
*ConfigStoreItemAPI* | [**UpsertConfigStoreItem**](docs/ConfigStoreItemAPI.md#upsertconfigstoreitem) | Insert or update an entry in a config store
*ContactAPI* | [**CreateContacts**](docs/ContactAPI.md#createcontacts) | Add a new customer contact
*ContactAPI* | [**DeleteContact**](docs/ContactAPI.md#deletecontact) | Delete a contact
*ContactAPI* | [**ListContacts**](docs/ContactAPI.md#listcontacts) | List contacts
*ContentAPI* | [**ContentCheck**](docs/ContentAPI.md#contentcheck) | Check status of content in each POP's cache
*CustomerAPI* | [**DeleteCustomer**](docs/CustomerAPI.md#deletecustomer) | Delete a customer
*CustomerAPI* | [**GetCustomer**](docs/CustomerAPI.md#getcustomer) | Get a customer
*CustomerAPI* | [**GetLoggedInCustomer**](docs/CustomerAPI.md#getloggedincustomer) | Get the logged in customer
*CustomerAPI* | [**ListUsers**](docs/CustomerAPI.md#listusers) | List users
*CustomerAPI* | [**UpdateCustomer**](docs/CustomerAPI.md#updatecustomer) | Update a customer
*CustomerAddressesAPI* | [**CreateCustomerAddress**](docs/CustomerAddressesAPI.md#createcustomeraddress) | Creates an address associated with a customer account.
*CustomerAddressesAPI* | [**ListCustomerAddresses**](docs/CustomerAddressesAPI.md#listcustomeraddresses) | Return the list of addresses associated with a customer account.
*CustomerAddressesAPI* | [**UpdateCustomerAddress**](docs/CustomerAddressesAPI.md#updatecustomeraddress) | Updates an address associated with a customer account.
*DictionaryAPI* | [**CreateDictionary**](docs/DictionaryAPI.md#createdictionary) | Create an edge dictionary
*DictionaryAPI* | [**DeleteDictionary**](docs/DictionaryAPI.md#deletedictionary) | Delete an edge dictionary
*DictionaryAPI* | [**GetDictionary**](docs/DictionaryAPI.md#getdictionary) | Get an edge dictionary
*DictionaryAPI* | [**ListDictionaries**](docs/DictionaryAPI.md#listdictionaries) | List edge dictionaries
*DictionaryAPI* | [**UpdateDictionary**](docs/DictionaryAPI.md#updatedictionary) | Update an edge dictionary
*DictionaryInfoAPI* | [**GetDictionaryInfo**](docs/DictionaryInfoAPI.md#getdictionaryinfo) | Get edge dictionary metadata
*DictionaryItemAPI* | [**BulkUpdateDictionaryItem**](docs/DictionaryItemAPI.md#bulkupdatedictionaryitem) | Update multiple entries in an edge dictionary
*DictionaryItemAPI* | [**CreateDictionaryItem**](docs/DictionaryItemAPI.md#createdictionaryitem) | Create an entry in an edge dictionary
*DictionaryItemAPI* | [**DeleteDictionaryItem**](docs/DictionaryItemAPI.md#deletedictionaryitem) | Delete an item from an edge dictionary
*DictionaryItemAPI* | [**GetDictionaryItem**](docs/DictionaryItemAPI.md#getdictionaryitem) | Get an item from an edge dictionary
*DictionaryItemAPI* | [**ListDictionaryItems**](docs/DictionaryItemAPI.md#listdictionaryitems) | List items in an edge dictionary
*DictionaryItemAPI* | [**UpdateDictionaryItem**](docs/DictionaryItemAPI.md#updatedictionaryitem) | Update an entry in an edge dictionary
*DictionaryItemAPI* | [**UpsertDictionaryItem**](docs/DictionaryItemAPI.md#upsertdictionaryitem) | Insert or update an entry in an edge dictionary
*DiffAPI* | [**DiffServiceVersions**](docs/DiffAPI.md#diffserviceversions) | Diff two service versions
*DirectorAPI* | [**CreateDirector**](docs/DirectorAPI.md#createdirector) | Create a director
*DirectorAPI* | [**DeleteDirector**](docs/DirectorAPI.md#deletedirector) | Delete a director
*DirectorAPI* | [**GetDirector**](docs/DirectorAPI.md#getdirector) | Get a director
*DirectorAPI* | [**ListDirectors**](docs/DirectorAPI.md#listdirectors) | List directors
*DirectorAPI* | [**UpdateDirector**](docs/DirectorAPI.md#updatedirector) | Update a director
*DirectorBackendAPI* | [**CreateDirectorBackend**](docs/DirectorBackendAPI.md#createdirectorbackend) | Create a director-backend relationship
*DirectorBackendAPI* | [**DeleteDirectorBackend**](docs/DirectorBackendAPI.md#deletedirectorbackend) | Delete a director-backend relationship
*DirectorBackendAPI* | [**GetDirectorBackend**](docs/DirectorBackendAPI.md#getdirectorbackend) | Get a director-backend relationship
*DomainAPI* | [**CheckDomain**](docs/DomainAPI.md#checkdomain) | Validate DNS configuration for a single domain on a service
*DomainAPI* | [**CheckDomains**](docs/DomainAPI.md#checkdomains) | Validate DNS configuration for all domains on a service
*DomainAPI* | [**CreateDomain**](docs/DomainAPI.md#createdomain) | Add a domain name to a service
*DomainAPI* | [**DeleteDomain**](docs/DomainAPI.md#deletedomain) | Remove a domain from a service
*DomainAPI* | [**GetDomain**](docs/DomainAPI.md#getdomain) | Describe a domain
*DomainAPI* | [**ListDomains**](docs/DomainAPI.md#listdomains) | List domains
*DomainAPI* | [**UpdateDomain**](docs/DomainAPI.md#updatedomain) | Update a domain
*DomainInspectorHistoricalAPI* | [**GetDomainInspectorHistorical**](docs/DomainInspectorHistoricalAPI.md#getdomaininspectorhistorical) | Get historical domain data for a service
*DomainInspectorRealtimeAPI* | [**GetDomainInspectorLast120Seconds**](docs/DomainInspectorRealtimeAPI.md#getdomaininspectorlast120seconds) | Get real-time domain data for the last 120 seconds
*DomainInspectorRealtimeAPI* | [**GetDomainInspectorLastMaxEntries**](docs/DomainInspectorRealtimeAPI.md#getdomaininspectorlastmaxentries) | Get a limited number of real-time domain data entries
*DomainInspectorRealtimeAPI* | [**GetDomainInspectorLastSecond**](docs/DomainInspectorRealtimeAPI.md#getdomaininspectorlastsecond) | Get real-time domain data from a specified time
*DomainOwnershipsAPI* | [**ListDomainOwnerships**](docs/DomainOwnershipsAPI.md#listdomainownerships) | List domain-ownerships
*EnabledProductsAPI* | [**DisableProduct**](docs/EnabledProductsAPI.md#disableproduct) | Disable a product
*EnabledProductsAPI* | [**EnableProduct**](docs/EnabledProductsAPI.md#enableproduct) | Enable a product
*EnabledProductsAPI* | [**GetEnabledProduct**](docs/EnabledProductsAPI.md#getenabledproduct) | Get enabled product
*EnabledProductsAPI* | [**GetProductConfiguration**](docs/EnabledProductsAPI.md#getproductconfiguration) | Get configuration for a product
*EnabledProductsAPI* | [**SetProductConfiguration**](docs/EnabledProductsAPI.md#setproductconfiguration) | Update configuration for a product
*EventsAPI* | [**GetEvent**](docs/EventsAPI.md#getevent) | Get an event
*EventsAPI* | [**ListEvents**](docs/EventsAPI.md#listevents) | List events
*GzipAPI* | [**CreateGzipConfig**](docs/GzipAPI.md#creategzipconfig) | Create a gzip configuration
*GzipAPI* | [**DeleteGzipConfig**](docs/GzipAPI.md#deletegzipconfig) | Delete a gzip configuration
*GzipAPI* | [**GetGzipConfigs**](docs/GzipAPI.md#getgzipconfigs) | Get a gzip configuration
*GzipAPI* | [**ListGzipConfigs**](docs/GzipAPI.md#listgzipconfigs) | List gzip configurations
*GzipAPI* | [**UpdateGzipConfig**](docs/GzipAPI.md#updategzipconfig) | Update a gzip configuration
*HeaderAPI* | [**CreateHeaderObject**](docs/HeaderAPI.md#createheaderobject) | Create a Header object
*HeaderAPI* | [**DeleteHeaderObject**](docs/HeaderAPI.md#deleteheaderobject) | Delete a Header object
*HeaderAPI* | [**GetHeaderObject**](docs/HeaderAPI.md#getheaderobject) | Get a Header object
*HeaderAPI* | [**ListHeaderObjects**](docs/HeaderAPI.md#listheaderobjects) | List Header objects
*HeaderAPI* | [**UpdateHeaderObject**](docs/HeaderAPI.md#updateheaderobject) | Update a Header object
*HealthcheckAPI* | [**CreateHealthcheck**](docs/HealthcheckAPI.md#createhealthcheck) | Create a health check
*HealthcheckAPI* | [**DeleteHealthcheck**](docs/HealthcheckAPI.md#deletehealthcheck) | Delete a health check
*HealthcheckAPI* | [**GetHealthcheck**](docs/HealthcheckAPI.md#gethealthcheck) | Get a health check
*HealthcheckAPI* | [**ListHealthchecks**](docs/HealthcheckAPI.md#listhealthchecks) | List health checks
*HealthcheckAPI* | [**UpdateHealthcheck**](docs/HealthcheckAPI.md#updatehealthcheck) | Update a health check
*HistoricalAPI* | [**GetHistStats**](docs/HistoricalAPI.md#gethiststats) | Get historical stats
*HistoricalAPI* | [**GetHistStatsAggregated**](docs/HistoricalAPI.md#gethiststatsaggregated) | Get aggregated historical stats
*HistoricalAPI* | [**GetHistStatsField**](docs/HistoricalAPI.md#gethiststatsfield) | Get historical stats for a single field
*HistoricalAPI* | [**GetHistStatsService**](docs/HistoricalAPI.md#gethiststatsservice) | Get historical stats for a single service
*HistoricalAPI* | [**GetHistStatsServiceField**](docs/HistoricalAPI.md#gethiststatsservicefield) | Get historical stats for a single service/field combination
*HistoricalAPI* | [**GetRegions**](docs/HistoricalAPI.md#getregions) | Get region codes
*HistoricalAPI* | [**GetUsage**](docs/HistoricalAPI.md#getusage) | Get usage statistics
*HistoricalAPI* | [**GetUsageMonth**](docs/HistoricalAPI.md#getusagemonth) | Get month-to-date usage statistics
*HistoricalAPI* | [**GetUsageService**](docs/HistoricalAPI.md#getusageservice) | Get usage statistics per service
*HTTP3API* | [**CreateHTTP3**](docs/Http3API.md#createhttp3) | Enable support for HTTP/3
*HTTP3API* | [**DeleteHTTP3**](docs/Http3API.md#deletehttp3) | Disable support for HTTP/3
*HTTP3API* | [**GetHTTP3**](docs/Http3API.md#gethttp3) | Get HTTP/3 status
*IamPermissionsAPI* | [**ListPermissions**](docs/IamPermissionsAPI.md#listpermissions) | List permissions
*IamRolesAPI* | [**AddRolePermissions**](docs/IamRolesAPI.md#addrolepermissions) | Add permissions to a role
*IamRolesAPI* | [**CreateARole**](docs/IamRolesAPI.md#createarole) | Create a role
*IamRolesAPI* | [**DeleteARole**](docs/IamRolesAPI.md#deletearole) | Delete a role
*IamRolesAPI* | [**GetARole**](docs/IamRolesAPI.md#getarole) | Get a role
*IamRolesAPI* | [**ListRolePermissions**](docs/IamRolesAPI.md#listrolepermissions) | List permissions in a role
*IamRolesAPI* | [**ListRoles**](docs/IamRolesAPI.md#listroles) | List roles
*IamRolesAPI* | [**RemoveRolePermissions**](docs/IamRolesAPI.md#removerolepermissions) | Remove permissions from a role
*IamRolesAPI* | [**UpdateARole**](docs/IamRolesAPI.md#updatearole) | Update a role
*IamServiceGroupsAPI* | [**AddServiceGroupServices**](docs/IamServiceGroupsAPI.md#addservicegroupservices) | Add services in a service group
*IamServiceGroupsAPI* | [**CreateAServiceGroup**](docs/IamServiceGroupsAPI.md#createaservicegroup) | Create a service group
*IamServiceGroupsAPI* | [**DeleteAServiceGroup**](docs/IamServiceGroupsAPI.md#deleteaservicegroup) | Delete a service group
*IamServiceGroupsAPI* | [**GetAServiceGroup**](docs/IamServiceGroupsAPI.md#getaservicegroup) | Get a service group
*IamServiceGroupsAPI* | [**ListServiceGroupServices**](docs/IamServiceGroupsAPI.md#listservicegroupservices) | List services to a service group
*IamServiceGroupsAPI* | [**ListServiceGroups**](docs/IamServiceGroupsAPI.md#listservicegroups) | List service groups
*IamServiceGroupsAPI* | [**RemoveServiceGroupServices**](docs/IamServiceGroupsAPI.md#removeservicegroupservices) | Remove services from a service group
*IamServiceGroupsAPI* | [**UpdateAServiceGroup**](docs/IamServiceGroupsAPI.md#updateaservicegroup) | Update a service group
*IamUserGroupsAPI* | [**AddUserGroupMembers**](docs/IamUserGroupsAPI.md#addusergroupmembers) | Add members to a user group
*IamUserGroupsAPI* | [**AddUserGroupRoles**](docs/IamUserGroupsAPI.md#addusergrouproles) | Add roles to a user group
*IamUserGroupsAPI* | [**AddUserGroupServiceGroups**](docs/IamUserGroupsAPI.md#addusergroupservicegroups) | Add service groups to a user group
*IamUserGroupsAPI* | [**CreateAUserGroup**](docs/IamUserGroupsAPI.md#createausergroup) | Create a user group
*IamUserGroupsAPI* | [**DeleteAUserGroup**](docs/IamUserGroupsAPI.md#deleteausergroup) | Delete a user group
*IamUserGroupsAPI* | [**GetAUserGroup**](docs/IamUserGroupsAPI.md#getausergroup) | Get a user group
*IamUserGroupsAPI* | [**ListUserGroupMembers**](docs/IamUserGroupsAPI.md#listusergroupmembers) | List members of a user group
*IamUserGroupsAPI* | [**ListUserGroupRoles**](docs/IamUserGroupsAPI.md#listusergrouproles) | List roles in a user group
*IamUserGroupsAPI* | [**ListUserGroupServiceGroups**](docs/IamUserGroupsAPI.md#listusergroupservicegroups) | List service groups in a user group
*IamUserGroupsAPI* | [**ListUserGroups**](docs/IamUserGroupsAPI.md#listusergroups) | List user groups
*IamUserGroupsAPI* | [**RemoveUserGroupMembers**](docs/IamUserGroupsAPI.md#removeusergroupmembers) | Remove members of a user group
*IamUserGroupsAPI* | [**RemoveUserGroupRoles**](docs/IamUserGroupsAPI.md#removeusergrouproles) | Remove roles from a user group
*IamUserGroupsAPI* | [**RemoveUserGroupServiceGroups**](docs/IamUserGroupsAPI.md#removeusergroupservicegroups) | Remove service groups from a user group
*IamUserGroupsAPI* | [**UpdateAUserGroup**](docs/IamUserGroupsAPI.md#updateausergroup) | Update a user group
*ImageOptimizerDefaultSettingsAPI* | [**GetDefaultSettings**](docs/ImageOptimizerDefaultSettingsAPI.md#getdefaultsettings) | Get current Image Optimizer Default Settings
*ImageOptimizerDefaultSettingsAPI* | [**UpdateDefaultSettings**](docs/ImageOptimizerDefaultSettingsAPI.md#updatedefaultsettings) | Update Image Optimizer Default Settings
*InvitationsAPI* | [**CreateInvitation**](docs/InvitationsAPI.md#createinvitation) | Create an invitation
*InvitationsAPI* | [**DeleteInvitation**](docs/InvitationsAPI.md#deleteinvitation) | Delete an invitation
*InvitationsAPI* | [**ListInvitations**](docs/InvitationsAPI.md#listinvitations) | List invitations
*KvStoreAPI* | [**CreateStore**](docs/KvStoreAPI.md#createstore) | Create a KV store.
*KvStoreAPI* | [**DeleteStore**](docs/KvStoreAPI.md#deletestore) | Delete a KV store.
*KvStoreAPI* | [**GetStore**](docs/KvStoreAPI.md#getstore) | Describe a KV store.
*KvStoreAPI* | [**GetStores**](docs/KvStoreAPI.md#getstores) | List KV stores.
*KvStoreItemAPI* | [**DeleteKeyFromStore**](docs/KvStoreItemAPI.md#deletekeyfromstore) | Delete kv store item.
*KvStoreItemAPI* | [**GetKeys**](docs/KvStoreItemAPI.md#getkeys) | List kv store keys.
*KvStoreItemAPI* | [**GetValueForKey**](docs/KvStoreItemAPI.md#getvalueforkey) | Get the value of an kv store item
*KvStoreItemAPI* | [**SetValueForKey**](docs/KvStoreItemAPI.md#setvalueforkey) | Insert an item into an kv store
*LegacyWafConfigurationSetsAPI* | [**ListWafConfigSets**](docs/LegacyWafConfigurationSetsAPI.md#listwafconfigsets) | List configuration sets
*LegacyWafConfigurationSetsAPI* | [**ListWafsConfigSet**](docs/LegacyWafConfigurationSetsAPI.md#listwafsconfigset) | List WAFs currently using a configuration set
*LegacyWafConfigurationSetsAPI* | [**UseWafConfigSet**](docs/LegacyWafConfigurationSetsAPI.md#usewafconfigset) | Apply a configuration set to a WAF
*LegacyWafFirewallAPI* | [**CreateLegacyWafFirewallService**](docs/LegacyWafFirewallAPI.md#createlegacywaffirewallservice) | Create a firewall
*LegacyWafFirewallAPI* | [**DisableLegacyWafFirewall**](docs/LegacyWafFirewallAPI.md#disablelegacywaffirewall) | Disable a firewall
*LegacyWafFirewallAPI* | [**EnableLegacyWafFirewall**](docs/LegacyWafFirewallAPI.md#enablelegacywaffirewall) | Enable a firewall
*LegacyWafFirewallAPI* | [**GetLegacyWafFirewall**](docs/LegacyWafFirewallAPI.md#getlegacywaffirewall) | Get a firewall object
*LegacyWafFirewallAPI* | [**GetLegacyWafFirewallService**](docs/LegacyWafFirewallAPI.md#getlegacywaffirewallservice) | Get a firewall
*LegacyWafFirewallAPI* | [**ListLegacyWafFirewalls**](docs/LegacyWafFirewallAPI.md#listlegacywaffirewalls) | List active firewalls
*LegacyWafFirewallAPI* | [**ListLegacyWafFirewallsService**](docs/LegacyWafFirewallAPI.md#listlegacywaffirewallsservice) | List firewalls
*LegacyWafFirewallAPI* | [**UpdateLegacyWafFirewallService**](docs/LegacyWafFirewallAPI.md#updatelegacywaffirewallservice) | Update a firewall
*LegacyWafOwaspAPI* | [**CreateOwaspSettings**](docs/LegacyWafOwaspAPI.md#createowaspsettings) | Create an OWASP settings object
*LegacyWafOwaspAPI* | [**GetOwaspSettings**](docs/LegacyWafOwaspAPI.md#getowaspsettings) | Get the OWASP settings object
*LegacyWafOwaspAPI* | [**UpdateOwaspSettings**](docs/LegacyWafOwaspAPI.md#updateowaspsettings) | Update the OWASP settings object
*LegacyWafRuleAPI* | [**GetLegacyWafFirewallRuleVcl**](docs/LegacyWafRuleAPI.md#getlegacywaffirewallrulevcl) | Get VCL for a rule associated with a firewall
*LegacyWafRuleAPI* | [**GetLegacyWafRule**](docs/LegacyWafRuleAPI.md#getlegacywafrule) | Get a rule
*LegacyWafRuleAPI* | [**GetLegacyWafRuleVcl**](docs/LegacyWafRuleAPI.md#getlegacywafrulevcl) | Get VCL for a rule
*LegacyWafRuleAPI* | [**ListLegacyWafRules**](docs/LegacyWafRuleAPI.md#listlegacywafrules) | List rules in the latest configuration set
*LegacyWafRuleStatusAPI* | [**GetWafFirewallRuleStatus**](docs/LegacyWafRuleStatusAPI.md#getwaffirewallrulestatus) | Get the status of a rule on a firewall
*LegacyWafRuleStatusAPI* | [**ListWafFirewallRuleStatuses**](docs/LegacyWafRuleStatusAPI.md#listwaffirewallrulestatuses) | List rule statuses
*LegacyWafRuleStatusAPI* | [**UpdateWafFirewallRuleStatus**](docs/LegacyWafRuleStatusAPI.md#updatewaffirewallrulestatus) | Update the status of a rule
*LegacyWafRuleStatusAPI* | [**UpdateWafFirewallRuleStatusesTag**](docs/LegacyWafRuleStatusAPI.md#updatewaffirewallrulestatusestag) | Create or update status of a tagged group of rules
*LegacyWafRulesetAPI* | [**GetWafRuleset**](docs/LegacyWafRulesetAPI.md#getwafruleset) | Get a WAF ruleset
*LegacyWafRulesetAPI* | [**GetWafRulesetVcl**](docs/LegacyWafRulesetAPI.md#getwafrulesetvcl) | Generate WAF ruleset VCL
*LegacyWafRulesetAPI* | [**UpdateWafRuleset**](docs/LegacyWafRulesetAPI.md#updatewafruleset) | Update a WAF ruleset
*LegacyWafTagAPI* | [**ListLegacyWafTags**](docs/LegacyWafTagAPI.md#listlegacywaftags) | List WAF tags
*LegacyWafUpdateStatusAPI* | [**GetWafUpdateStatus**](docs/LegacyWafUpdateStatusAPI.md#getwafupdatestatus) | Get the status of a WAF update
*LegacyWafUpdateStatusAPI* | [**ListWafUpdateStatuses**](docs/LegacyWafUpdateStatusAPI.md#listwafupdatestatuses) | List update statuses
*LoggingAzureblobAPI* | [**CreateLogAzure**](docs/LoggingAzureblobAPI.md#createlogazure) | Create an Azure Blob Storage log endpoint
*LoggingAzureblobAPI* | [**DeleteLogAzure**](docs/LoggingAzureblobAPI.md#deletelogazure) | Delete the Azure Blob Storage log endpoint
*LoggingAzureblobAPI* | [**GetLogAzure**](docs/LoggingAzureblobAPI.md#getlogazure) | Get an Azure Blob Storage log endpoint
*LoggingAzureblobAPI* | [**ListLogAzure**](docs/LoggingAzureblobAPI.md#listlogazure) | List Azure Blob Storage log endpoints
*LoggingAzureblobAPI* | [**UpdateLogAzure**](docs/LoggingAzureblobAPI.md#updatelogazure) | Update an Azure Blob Storage log endpoint
*LoggingBigqueryAPI* | [**CreateLogBigquery**](docs/LoggingBigqueryAPI.md#createlogbigquery) | Create a BigQuery log endpoint
*LoggingBigqueryAPI* | [**DeleteLogBigquery**](docs/LoggingBigqueryAPI.md#deletelogbigquery) | Delete a BigQuery log endpoint
*LoggingBigqueryAPI* | [**GetLogBigquery**](docs/LoggingBigqueryAPI.md#getlogbigquery) | Get a BigQuery log endpoint
*LoggingBigqueryAPI* | [**ListLogBigquery**](docs/LoggingBigqueryAPI.md#listlogbigquery) | List BigQuery log endpoints
*LoggingBigqueryAPI* | [**UpdateLogBigquery**](docs/LoggingBigqueryAPI.md#updatelogbigquery) | Update a BigQuery log endpoint
*LoggingCloudfilesAPI* | [**CreateLogCloudfiles**](docs/LoggingCloudfilesAPI.md#createlogcloudfiles) | Create a Cloud Files log endpoint
*LoggingCloudfilesAPI* | [**DeleteLogCloudfiles**](docs/LoggingCloudfilesAPI.md#deletelogcloudfiles) | Delete the Cloud Files log endpoint
*LoggingCloudfilesAPI* | [**GetLogCloudfiles**](docs/LoggingCloudfilesAPI.md#getlogcloudfiles) | Get a Cloud Files log endpoint
*LoggingCloudfilesAPI* | [**ListLogCloudfiles**](docs/LoggingCloudfilesAPI.md#listlogcloudfiles) | List Cloud Files log endpoints
*LoggingCloudfilesAPI* | [**UpdateLogCloudfiles**](docs/LoggingCloudfilesAPI.md#updatelogcloudfiles) | Update the Cloud Files log endpoint
*LoggingDatadogAPI* | [**CreateLogDatadog**](docs/LoggingDatadogAPI.md#createlogdatadog) | Create a Datadog log endpoint
*LoggingDatadogAPI* | [**DeleteLogDatadog**](docs/LoggingDatadogAPI.md#deletelogdatadog) | Delete a Datadog log endpoint
*LoggingDatadogAPI* | [**GetLogDatadog**](docs/LoggingDatadogAPI.md#getlogdatadog) | Get a Datadog log endpoint
*LoggingDatadogAPI* | [**ListLogDatadog**](docs/LoggingDatadogAPI.md#listlogdatadog) | List Datadog log endpoints
*LoggingDatadogAPI* | [**UpdateLogDatadog**](docs/LoggingDatadogAPI.md#updatelogdatadog) | Update a Datadog log endpoint
*LoggingDigitaloceanAPI* | [**CreateLogDigocean**](docs/LoggingDigitaloceanAPI.md#createlogdigocean) | Create a DigitalOcean Spaces log endpoint
*LoggingDigitaloceanAPI* | [**DeleteLogDigocean**](docs/LoggingDigitaloceanAPI.md#deletelogdigocean) | Delete a DigitalOcean Spaces log endpoint
*LoggingDigitaloceanAPI* | [**GetLogDigocean**](docs/LoggingDigitaloceanAPI.md#getlogdigocean) | Get a DigitalOcean Spaces log endpoint
*LoggingDigitaloceanAPI* | [**ListLogDigocean**](docs/LoggingDigitaloceanAPI.md#listlogdigocean) | List DigitalOcean Spaces log endpoints
*LoggingDigitaloceanAPI* | [**UpdateLogDigocean**](docs/LoggingDigitaloceanAPI.md#updatelogdigocean) | Update a DigitalOcean Spaces log endpoint
*LoggingElasticsearchAPI* | [**CreateLogElasticsearch**](docs/LoggingElasticsearchAPI.md#createlogelasticsearch) | Create an Elasticsearch log endpoint
*LoggingElasticsearchAPI* | [**DeleteLogElasticsearch**](docs/LoggingElasticsearchAPI.md#deletelogelasticsearch) | Delete an Elasticsearch log endpoint
*LoggingElasticsearchAPI* | [**GetLogElasticsearch**](docs/LoggingElasticsearchAPI.md#getlogelasticsearch) | Get an Elasticsearch log endpoint
*LoggingElasticsearchAPI* | [**ListLogElasticsearch**](docs/LoggingElasticsearchAPI.md#listlogelasticsearch) | List Elasticsearch log endpoints
*LoggingElasticsearchAPI* | [**UpdateLogElasticsearch**](docs/LoggingElasticsearchAPI.md#updatelogelasticsearch) | Update an Elasticsearch log endpoint
*LoggingFtpAPI* | [**CreateLogFtp**](docs/LoggingFtpAPI.md#createlogftp) | Create an FTP log endpoint
*LoggingFtpAPI* | [**DeleteLogFtp**](docs/LoggingFtpAPI.md#deletelogftp) | Delete an FTP log endpoint
*LoggingFtpAPI* | [**GetLogFtp**](docs/LoggingFtpAPI.md#getlogftp) | Get an FTP log endpoint
*LoggingFtpAPI* | [**ListLogFtp**](docs/LoggingFtpAPI.md#listlogftp) | List FTP log endpoints
*LoggingFtpAPI* | [**UpdateLogFtp**](docs/LoggingFtpAPI.md#updatelogftp) | Update an FTP log endpoint
*LoggingGcsAPI* | [**CreateLogGcs**](docs/LoggingGcsAPI.md#createloggcs) | Create a GCS log endpoint
*LoggingGcsAPI* | [**DeleteLogGcs**](docs/LoggingGcsAPI.md#deleteloggcs) | Delete a GCS log endpoint
*LoggingGcsAPI* | [**GetLogGcs**](docs/LoggingGcsAPI.md#getloggcs) | Get a GCS log endpoint
*LoggingGcsAPI* | [**ListLogGcs**](docs/LoggingGcsAPI.md#listloggcs) | List GCS log endpoints
*LoggingGcsAPI* | [**UpdateLogGcs**](docs/LoggingGcsAPI.md#updateloggcs) | Update a GCS log endpoint
*LoggingHerokuAPI* | [**CreateLogHeroku**](docs/LoggingHerokuAPI.md#createlogheroku) | Create a Heroku log endpoint
*LoggingHerokuAPI* | [**DeleteLogHeroku**](docs/LoggingHerokuAPI.md#deletelogheroku) | Delete the Heroku log endpoint
*LoggingHerokuAPI* | [**GetLogHeroku**](docs/LoggingHerokuAPI.md#getlogheroku) | Get a Heroku log endpoint
*LoggingHerokuAPI* | [**ListLogHeroku**](docs/LoggingHerokuAPI.md#listlogheroku) | List Heroku log endpoints
*LoggingHerokuAPI* | [**UpdateLogHeroku**](docs/LoggingHerokuAPI.md#updatelogheroku) | Update the Heroku log endpoint
*LoggingHoneycombAPI* | [**CreateLogHoneycomb**](docs/LoggingHoneycombAPI.md#createloghoneycomb) | Create a Honeycomb log endpoint
*LoggingHoneycombAPI* | [**DeleteLogHoneycomb**](docs/LoggingHoneycombAPI.md#deleteloghoneycomb) | Delete the Honeycomb log endpoint
*LoggingHoneycombAPI* | [**GetLogHoneycomb**](docs/LoggingHoneycombAPI.md#getloghoneycomb) | Get a Honeycomb log endpoint
*LoggingHoneycombAPI* | [**ListLogHoneycomb**](docs/LoggingHoneycombAPI.md#listloghoneycomb) | List Honeycomb log endpoints
*LoggingHoneycombAPI* | [**UpdateLogHoneycomb**](docs/LoggingHoneycombAPI.md#updateloghoneycomb) | Update a Honeycomb log endpoint
*LoggingHTTPSAPI* | [**CreateLogHTTPS**](docs/LoggingHTTPSAPI.md#createloghttps) | Create an HTTPS log endpoint
*LoggingHTTPSAPI* | [**DeleteLogHTTPS**](docs/LoggingHTTPSAPI.md#deleteloghttps) | Delete an HTTPS log endpoint
*LoggingHTTPSAPI* | [**GetLogHTTPS**](docs/LoggingHTTPSAPI.md#getloghttps) | Get an HTTPS log endpoint
*LoggingHTTPSAPI* | [**ListLogHTTPS**](docs/LoggingHTTPSAPI.md#listloghttps) | List HTTPS log endpoints
*LoggingHTTPSAPI* | [**UpdateLogHTTPS**](docs/LoggingHTTPSAPI.md#updateloghttps) | Update an HTTPS log endpoint
*LoggingKafkaAPI* | [**CreateLogKafka**](docs/LoggingKafkaAPI.md#createlogkafka) | Create a Kafka log endpoint
*LoggingKafkaAPI* | [**DeleteLogKafka**](docs/LoggingKafkaAPI.md#deletelogkafka) | Delete the Kafka log endpoint
*LoggingKafkaAPI* | [**GetLogKafka**](docs/LoggingKafkaAPI.md#getlogkafka) | Get a Kafka log endpoint
*LoggingKafkaAPI* | [**ListLogKafka**](docs/LoggingKafkaAPI.md#listlogkafka) | List Kafka log endpoints
*LoggingKafkaAPI* | [**UpdateLogKafka**](docs/LoggingKafkaAPI.md#updatelogkafka) | Update the Kafka log endpoint
*LoggingKinesisAPI* | [**CreateLogKinesis**](docs/LoggingKinesisAPI.md#createlogkinesis) | Create  an Amazon Kinesis log endpoint
*LoggingKinesisAPI* | [**DeleteLogKinesis**](docs/LoggingKinesisAPI.md#deletelogkinesis) | Delete the Amazon Kinesis log endpoint
*LoggingKinesisAPI* | [**GetLogKinesis**](docs/LoggingKinesisAPI.md#getlogkinesis) | Get an Amazon Kinesis log endpoint
*LoggingKinesisAPI* | [**ListLogKinesis**](docs/LoggingKinesisAPI.md#listlogkinesis) | List Amazon Kinesis log endpoints
*LoggingKinesisAPI* | [**UpdateLogKinesis**](docs/LoggingKinesisAPI.md#updatelogkinesis) | Update the Amazon Kinesis log endpoint
*LoggingLogentriesAPI* | [**CreateLogLogentries**](docs/LoggingLogentriesAPI.md#createloglogentries) | Create a Logentries log endpoint
*LoggingLogentriesAPI* | [**DeleteLogLogentries**](docs/LoggingLogentriesAPI.md#deleteloglogentries) | Delete a Logentries log endpoint
*LoggingLogentriesAPI* | [**GetLogLogentries**](docs/LoggingLogentriesAPI.md#getloglogentries) | Get a Logentries log endpoint
*LoggingLogentriesAPI* | [**ListLogLogentries**](docs/LoggingLogentriesAPI.md#listloglogentries) | List Logentries log endpoints
*LoggingLogentriesAPI* | [**UpdateLogLogentries**](docs/LoggingLogentriesAPI.md#updateloglogentries) | Update a Logentries log endpoint
*LoggingLogglyAPI* | [**CreateLogLoggly**](docs/LoggingLogglyAPI.md#createlogloggly) | Create a Loggly log endpoint
*LoggingLogglyAPI* | [**DeleteLogLoggly**](docs/LoggingLogglyAPI.md#deletelogloggly) | Delete a Loggly log endpoint
*LoggingLogglyAPI* | [**GetLogLoggly**](docs/LoggingLogglyAPI.md#getlogloggly) | Get a Loggly log endpoint
*LoggingLogglyAPI* | [**ListLogLoggly**](docs/LoggingLogglyAPI.md#listlogloggly) | List Loggly log endpoints
*LoggingLogglyAPI* | [**UpdateLogLoggly**](docs/LoggingLogglyAPI.md#updatelogloggly) | Update a Loggly log endpoint
*LoggingLogshuttleAPI* | [**CreateLogLogshuttle**](docs/LoggingLogshuttleAPI.md#createloglogshuttle) | Create a Log Shuttle log endpoint
*LoggingLogshuttleAPI* | [**DeleteLogLogshuttle**](docs/LoggingLogshuttleAPI.md#deleteloglogshuttle) | Delete a Log Shuttle log endpoint
*LoggingLogshuttleAPI* | [**GetLogLogshuttle**](docs/LoggingLogshuttleAPI.md#getloglogshuttle) | Get a Log Shuttle log endpoint
*LoggingLogshuttleAPI* | [**ListLogLogshuttle**](docs/LoggingLogshuttleAPI.md#listloglogshuttle) | List Log Shuttle log endpoints
*LoggingLogshuttleAPI* | [**UpdateLogLogshuttle**](docs/LoggingLogshuttleAPI.md#updateloglogshuttle) | Update a Log Shuttle log endpoint
*LoggingNewrelicAPI* | [**CreateLogNewrelic**](docs/LoggingNewrelicAPI.md#createlognewrelic) | Create a New Relic log endpoint
*LoggingNewrelicAPI* | [**DeleteLogNewrelic**](docs/LoggingNewrelicAPI.md#deletelognewrelic) | Delete a New Relic log endpoint
*LoggingNewrelicAPI* | [**GetLogNewrelic**](docs/LoggingNewrelicAPI.md#getlognewrelic) | Get a New Relic log endpoint
*LoggingNewrelicAPI* | [**ListLogNewrelic**](docs/LoggingNewrelicAPI.md#listlognewrelic) | List New Relic log endpoints
*LoggingNewrelicAPI* | [**UpdateLogNewrelic**](docs/LoggingNewrelicAPI.md#updatelognewrelic) | Update a New Relic log endpoint
*LoggingNewrelicotlpAPI* | [**CreateLogNewrelicotlp**](docs/LoggingNewrelicotlpAPI.md#createlognewrelicotlp) | Create a New Relic OTLP endpoint
*LoggingNewrelicotlpAPI* | [**DeleteLogNewrelicotlp**](docs/LoggingNewrelicotlpAPI.md#deletelognewrelicotlp) | Delete a New Relic OTLP endpoint
*LoggingNewrelicotlpAPI* | [**GetLogNewrelicotlp**](docs/LoggingNewrelicotlpAPI.md#getlognewrelicotlp) | Get a New Relic OTLP endpoint
*LoggingNewrelicotlpAPI* | [**ListLogNewrelicotlp**](docs/LoggingNewrelicotlpAPI.md#listlognewrelicotlp) | List New Relic OTLP endpoints
*LoggingNewrelicotlpAPI* | [**UpdateLogNewrelicotlp**](docs/LoggingNewrelicotlpAPI.md#updatelognewrelicotlp) | Update a New Relic log endpoint
*LoggingOpenstackAPI* | [**CreateLogOpenstack**](docs/LoggingOpenstackAPI.md#createlogopenstack) | Create an OpenStack log endpoint
*LoggingOpenstackAPI* | [**DeleteLogOpenstack**](docs/LoggingOpenstackAPI.md#deletelogopenstack) | Delete an OpenStack log endpoint
*LoggingOpenstackAPI* | [**GetLogOpenstack**](docs/LoggingOpenstackAPI.md#getlogopenstack) | Get an OpenStack log endpoint
*LoggingOpenstackAPI* | [**ListLogOpenstack**](docs/LoggingOpenstackAPI.md#listlogopenstack) | List OpenStack log endpoints
*LoggingOpenstackAPI* | [**UpdateLogOpenstack**](docs/LoggingOpenstackAPI.md#updatelogopenstack) | Update an OpenStack log endpoint
*LoggingPapertrailAPI* | [**CreateLogPapertrail**](docs/LoggingPapertrailAPI.md#createlogpapertrail) | Create a Papertrail log endpoint
*LoggingPapertrailAPI* | [**DeleteLogPapertrail**](docs/LoggingPapertrailAPI.md#deletelogpapertrail) | Delete a Papertrail log endpoint
*LoggingPapertrailAPI* | [**GetLogPapertrail**](docs/LoggingPapertrailAPI.md#getlogpapertrail) | Get a Papertrail log endpoint
*LoggingPapertrailAPI* | [**ListLogPapertrail**](docs/LoggingPapertrailAPI.md#listlogpapertrail) | List Papertrail log endpoints
*LoggingPapertrailAPI* | [**UpdateLogPapertrail**](docs/LoggingPapertrailAPI.md#updatelogpapertrail) | Update a Papertrail log endpoint
*LoggingPubsubAPI* | [**CreateLogGcpPubsub**](docs/LoggingPubsubAPI.md#createloggcppubsub) | Create a GCP Cloud Pub/Sub log endpoint
*LoggingPubsubAPI* | [**DeleteLogGcpPubsub**](docs/LoggingPubsubAPI.md#deleteloggcppubsub) | Delete a GCP Cloud Pub/Sub log endpoint
*LoggingPubsubAPI* | [**GetLogGcpPubsub**](docs/LoggingPubsubAPI.md#getloggcppubsub) | Get a GCP Cloud Pub/Sub log endpoint
*LoggingPubsubAPI* | [**ListLogGcpPubsub**](docs/LoggingPubsubAPI.md#listloggcppubsub) | List GCP Cloud Pub/Sub log endpoints
*LoggingPubsubAPI* | [**UpdateLogGcpPubsub**](docs/LoggingPubsubAPI.md#updateloggcppubsub) | Update a GCP Cloud Pub/Sub log endpoint
*LoggingS3API* | [**CreateLogAwsS3**](docs/LoggingS3API.md#createlogawss3) | Create an AWS S3 log endpoint
*LoggingS3API* | [**DeleteLogAwsS3**](docs/LoggingS3API.md#deletelogawss3) | Delete an AWS S3 log endpoint
*LoggingS3API* | [**GetLogAwsS3**](docs/LoggingS3API.md#getlogawss3) | Get an AWS S3 log endpoint
*LoggingS3API* | [**ListLogAwsS3**](docs/LoggingS3API.md#listlogawss3) | List AWS S3 log endpoints
*LoggingS3API* | [**UpdateLogAwsS3**](docs/LoggingS3API.md#updatelogawss3) | Update an AWS S3 log endpoint
*LoggingScalyrAPI* | [**CreateLogScalyr**](docs/LoggingScalyrAPI.md#createlogscalyr) | Create a Scalyr log endpoint
*LoggingScalyrAPI* | [**DeleteLogScalyr**](docs/LoggingScalyrAPI.md#deletelogscalyr) | Delete the Scalyr log endpoint
*LoggingScalyrAPI* | [**GetLogScalyr**](docs/LoggingScalyrAPI.md#getlogscalyr) | Get a Scalyr log endpoint
*LoggingScalyrAPI* | [**ListLogScalyr**](docs/LoggingScalyrAPI.md#listlogscalyr) | List Scalyr log endpoints
*LoggingScalyrAPI* | [**UpdateLogScalyr**](docs/LoggingScalyrAPI.md#updatelogscalyr) | Update the Scalyr log endpoint
*LoggingSftpAPI* | [**CreateLogSftp**](docs/LoggingSftpAPI.md#createlogsftp) | Create an SFTP log endpoint
*LoggingSftpAPI* | [**DeleteLogSftp**](docs/LoggingSftpAPI.md#deletelogsftp) | Delete an SFTP log endpoint
*LoggingSftpAPI* | [**GetLogSftp**](docs/LoggingSftpAPI.md#getlogsftp) | Get an SFTP log endpoint
*LoggingSftpAPI* | [**ListLogSftp**](docs/LoggingSftpAPI.md#listlogsftp) | List SFTP log endpoints
*LoggingSftpAPI* | [**UpdateLogSftp**](docs/LoggingSftpAPI.md#updatelogsftp) | Update an SFTP log endpoint
*LoggingSplunkAPI* | [**CreateLogSplunk**](docs/LoggingSplunkAPI.md#createlogsplunk) | Create a Splunk log endpoint
*LoggingSplunkAPI* | [**DeleteLogSplunk**](docs/LoggingSplunkAPI.md#deletelogsplunk) | Delete a Splunk log endpoint
*LoggingSplunkAPI* | [**GetLogSplunk**](docs/LoggingSplunkAPI.md#getlogsplunk) | Get a Splunk log endpoint
*LoggingSplunkAPI* | [**ListLogSplunk**](docs/LoggingSplunkAPI.md#listlogsplunk) | List Splunk log endpoints
*LoggingSplunkAPI* | [**UpdateLogSplunk**](docs/LoggingSplunkAPI.md#updatelogsplunk) | Update a Splunk log endpoint
*LoggingSumologicAPI* | [**CreateLogSumologic**](docs/LoggingSumologicAPI.md#createlogsumologic) | Create a Sumologic log endpoint
*LoggingSumologicAPI* | [**DeleteLogSumologic**](docs/LoggingSumologicAPI.md#deletelogsumologic) | Delete a Sumologic log endpoint
*LoggingSumologicAPI* | [**GetLogSumologic**](docs/LoggingSumologicAPI.md#getlogsumologic) | Get a Sumologic log endpoint
*LoggingSumologicAPI* | [**ListLogSumologic**](docs/LoggingSumologicAPI.md#listlogsumologic) | List Sumologic log endpoints
*LoggingSumologicAPI* | [**UpdateLogSumologic**](docs/LoggingSumologicAPI.md#updatelogsumologic) | Update a Sumologic log endpoint
*LoggingSyslogAPI* | [**CreateLogSyslog**](docs/LoggingSyslogAPI.md#createlogsyslog) | Create a syslog log endpoint
*LoggingSyslogAPI* | [**DeleteLogSyslog**](docs/LoggingSyslogAPI.md#deletelogsyslog) | Delete a syslog log endpoint
*LoggingSyslogAPI* | [**GetLogSyslog**](docs/LoggingSyslogAPI.md#getlogsyslog) | Get a syslog log endpoint
*LoggingSyslogAPI* | [**ListLogSyslog**](docs/LoggingSyslogAPI.md#listlogsyslog) | List Syslog log endpoints
*LoggingSyslogAPI* | [**UpdateLogSyslog**](docs/LoggingSyslogAPI.md#updatelogsyslog) | Update a syslog log endpoint
*MutualAuthenticationAPI* | [**CreateMutualTLSAuthentication**](docs/MutualAuthenticationAPI.md#createmutualtlsauthentication) | Create a Mutual Authentication
*MutualAuthenticationAPI* | [**DeleteMutualTLS**](docs/MutualAuthenticationAPI.md#deletemutualtls) | Delete a Mutual TLS
*MutualAuthenticationAPI* | [**GetMutualAuthentication**](docs/MutualAuthenticationAPI.md#getmutualauthentication) | Get a Mutual Authentication
*MutualAuthenticationAPI* | [**ListMutualAuthentications**](docs/MutualAuthenticationAPI.md#listmutualauthentications) | List Mutual Authentications
*MutualAuthenticationAPI* | [**PatchMutualAuthentication**](docs/MutualAuthenticationAPI.md#patchmutualauthentication) | Update a Mutual Authentication
*ObservabilityCustomDashboardsAPI* | [**CreateDashboard**](docs/ObservabilityCustomDashboardsAPI.md#createdashboard) | Create a new dashboard
*ObservabilityCustomDashboardsAPI* | [**DeleteDashboard**](docs/ObservabilityCustomDashboardsAPI.md#deletedashboard) | Delete an existing dashboard
*ObservabilityCustomDashboardsAPI* | [**GetDashboard**](docs/ObservabilityCustomDashboardsAPI.md#getdashboard) | Retrieve a dashboard by ID
*ObservabilityCustomDashboardsAPI* | [**ListDashboards**](docs/ObservabilityCustomDashboardsAPI.md#listdashboards) | List all custom dashboards
*ObservabilityCustomDashboardsAPI* | [**UpdateDashboard**](docs/ObservabilityCustomDashboardsAPI.md#updatedashboard) | Update an existing dashboard
*OriginInspectorHistoricalAPI* | [**GetOriginInspectorHistorical**](docs/OriginInspectorHistoricalAPI.md#getorigininspectorhistorical) | Get historical origin data for a service
*OriginInspectorRealtimeAPI* | [**GetOriginInspectorLast120Seconds**](docs/OriginInspectorRealtimeAPI.md#getorigininspectorlast120seconds) | Get real-time origin data for the last 120 seconds
*OriginInspectorRealtimeAPI* | [**GetOriginInspectorLastMaxEntries**](docs/OriginInspectorRealtimeAPI.md#getorigininspectorlastmaxentries) | Get a limited number of real-time origin data entries
*OriginInspectorRealtimeAPI* | [**GetOriginInspectorLastSecond**](docs/OriginInspectorRealtimeAPI.md#getorigininspectorlastsecond) | Get real-time origin data from specific time.
*PackageAPI* | [**GetPackage**](docs/PackageAPI.md#getpackage) | Get details of the service's Compute package.
*PackageAPI* | [**PutPackage**](docs/PackageAPI.md#putpackage) | Upload a Compute package.
*PoolAPI* | [**CreateServerPool**](docs/PoolAPI.md#createserverpool) | Create a server pool
*PoolAPI* | [**DeleteServerPool**](docs/PoolAPI.md#deleteserverpool) | Delete a server pool
*PoolAPI* | [**GetServerPool**](docs/PoolAPI.md#getserverpool) | Get a server pool
*PoolAPI* | [**ListServerPools**](docs/PoolAPI.md#listserverpools) | List server pools
*PoolAPI* | [**UpdateServerPool**](docs/PoolAPI.md#updateserverpool) | Update a server pool
*PopAPI* | [**ListPops**](docs/PopAPI.md#listpops) | List Fastly POPs
*PublicIPListAPI* | [**ListFastlyIps**](docs/PublicIPListAPI.md#listfastlyips) | List Fastly's public IPs
*PublishAPI* | [**Publish**](docs/PublishAPI.md#publish) | Send messages to Fanout subscribers
*PurgeAPI* | [**BulkPurgeTag**](docs/PurgeAPI.md#bulkpurgetag) | Purge multiple surrogate key tags
*PurgeAPI* | [**PurgeAll**](docs/PurgeAPI.md#purgeall) | Purge everything from a service
*PurgeAPI* | [**PurgeSingleURL**](docs/PurgeAPI.md#purgesingleurl) | Purge a URL
*PurgeAPI* | [**PurgeTag**](docs/PurgeAPI.md#purgetag) | Purge by surrogate key tag
*RateLimiterAPI* | [**CreateRateLimiter**](docs/RateLimiterAPI.md#createratelimiter) | Create a rate limiter
*RateLimiterAPI* | [**DeleteRateLimiter**](docs/RateLimiterAPI.md#deleteratelimiter) | Delete a rate limiter
*RateLimiterAPI* | [**GetRateLimiter**](docs/RateLimiterAPI.md#getratelimiter) | Get a rate limiter
*RateLimiterAPI* | [**ListRateLimiters**](docs/RateLimiterAPI.md#listratelimiters) | List rate limiters
*RateLimiterAPI* | [**UpdateRateLimiter**](docs/RateLimiterAPI.md#updateratelimiter) | Update a rate limiter
*RealtimeAPI* | [**GetStatsLast120Seconds**](docs/RealtimeAPI.md#getstatslast120seconds) | Get real-time data for the last 120 seconds
*RealtimeAPI* | [**GetStatsLast120SecondsLimitEntries**](docs/RealtimeAPI.md#getstatslast120secondslimitentries) | Get a limited number of real-time data entries
*RealtimeAPI* | [**GetStatsLastSecond**](docs/RealtimeAPI.md#getstatslastsecond) | Get real-time data from specified time
*RequestSettingsAPI* | [**CreateRequestSettings**](docs/RequestSettingsAPI.md#createrequestsettings) | Create a Request Settings object
*RequestSettingsAPI* | [**DeleteRequestSettings**](docs/RequestSettingsAPI.md#deleterequestsettings) | Delete a Request Settings object
*RequestSettingsAPI* | [**GetRequestSettings**](docs/RequestSettingsAPI.md#getrequestsettings) | Get a Request Settings object
*RequestSettingsAPI* | [**ListRequestSettings**](docs/RequestSettingsAPI.md#listrequestsettings) | List Request Settings objects
*RequestSettingsAPI* | [**UpdateRequestSettings**](docs/RequestSettingsAPI.md#updaterequestsettings) | Update a Request Settings object
*ResourceAPI* | [**CreateResource**](docs/ResourceAPI.md#createresource) | Create a resource link
*ResourceAPI* | [**DeleteResource**](docs/ResourceAPI.md#deleteresource) | Delete a resource link
*ResourceAPI* | [**GetResource**](docs/ResourceAPI.md#getresource) | Display a resource link
*ResourceAPI* | [**ListResources**](docs/ResourceAPI.md#listresources) | List resource links
*ResourceAPI* | [**UpdateResource**](docs/ResourceAPI.md#updateresource) | Update a resource link
*ResponseObjectAPI* | [**CreateResponseObject**](docs/ResponseObjectAPI.md#createresponseobject) | Create a Response object
*ResponseObjectAPI* | [**DeleteResponseObject**](docs/ResponseObjectAPI.md#deleteresponseobject) | Delete a Response Object
*ResponseObjectAPI* | [**GetResponseObject**](docs/ResponseObjectAPI.md#getresponseobject) | Get a Response object
*ResponseObjectAPI* | [**ListResponseObjects**](docs/ResponseObjectAPI.md#listresponseobjects) | List Response objects
*ResponseObjectAPI* | [**UpdateResponseObject**](docs/ResponseObjectAPI.md#updateresponseobject) | Update a Response object
*SecretStoreAPI* | [**ClientKey**](docs/SecretStoreAPI.md#clientkey) | Create new client key
*SecretStoreAPI* | [**CreateSecretStore**](docs/SecretStoreAPI.md#createsecretstore) | Create new secret store
*SecretStoreAPI* | [**DeleteSecretStore**](docs/SecretStoreAPI.md#deletesecretstore) | Delete secret store
*SecretStoreAPI* | [**GetSecretStore**](docs/SecretStoreAPI.md#getsecretstore) | Get secret store by ID
*SecretStoreAPI* | [**GetSecretStores**](docs/SecretStoreAPI.md#getsecretstores) | Get all secret stores
*SecretStoreAPI* | [**SigningKey**](docs/SecretStoreAPI.md#signingkey) | Get public key
*SecretStoreItemAPI* | [**CreateSecret**](docs/SecretStoreItemAPI.md#createsecret) | Create a new secret in a store.
*SecretStoreItemAPI* | [**DeleteSecret**](docs/SecretStoreItemAPI.md#deletesecret) | Delete a secret from a store.
*SecretStoreItemAPI* | [**GetSecret**](docs/SecretStoreItemAPI.md#getsecret) | Get secret metadata.
*SecretStoreItemAPI* | [**GetSecrets**](docs/SecretStoreItemAPI.md#getsecrets) | List secrets within a store.
*SecretStoreItemAPI* | [**MustRecreateSecret**](docs/SecretStoreItemAPI.md#mustrecreatesecret) | Recreate a secret in a store.
*SecretStoreItemAPI* | [**RecreateSecret**](docs/SecretStoreItemAPI.md#recreatesecret) | Create or recreate a secret in a store.
*ServerAPI* | [**CreatePoolServer**](docs/ServerAPI.md#createpoolserver) | Add a server to a pool
*ServerAPI* | [**DeletePoolServer**](docs/ServerAPI.md#deletepoolserver) | Delete a server from a pool
*ServerAPI* | [**GetPoolServer**](docs/ServerAPI.md#getpoolserver) | Get a pool server
*ServerAPI* | [**ListPoolServers**](docs/ServerAPI.md#listpoolservers) | List servers in a pool
*ServerAPI* | [**UpdatePoolServer**](docs/ServerAPI.md#updatepoolserver) | Update a server
*ServiceAPI* | [**CreateService**](docs/ServiceAPI.md#createservice) | Create a service
*ServiceAPI* | [**DeleteService**](docs/ServiceAPI.md#deleteservice) | Delete a service
*ServiceAPI* | [**GetService**](docs/ServiceAPI.md#getservice) | Get a service
*ServiceAPI* | [**GetServiceDetail**](docs/ServiceAPI.md#getservicedetail) | Get service details
*ServiceAPI* | [**ListServiceDomains**](docs/ServiceAPI.md#listservicedomains) | List the domains within a service
*ServiceAPI* | [**ListServices**](docs/ServiceAPI.md#listservices) | List services
*ServiceAPI* | [**SearchService**](docs/ServiceAPI.md#searchservice) | Search for a service by name
*ServiceAPI* | [**UpdateService**](docs/ServiceAPI.md#updateservice) | Update a service
*ServiceAuthorizationsAPI* | [**CreateServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#createserviceauthorization) | Create service authorization
*ServiceAuthorizationsAPI* | [**DeleteServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#deleteserviceauthorization) | Delete service authorization
*ServiceAuthorizationsAPI* | [**DeleteServiceAuthorization2**](docs/ServiceAuthorizationsAPI.md#deleteserviceauthorization2) | Delete service authorizations
*ServiceAuthorizationsAPI* | [**ListServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#listserviceauthorization) | List service authorizations
*ServiceAuthorizationsAPI* | [**ShowServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#showserviceauthorization) | Show service authorization
*ServiceAuthorizationsAPI* | [**UpdateServiceAuthorization**](docs/ServiceAuthorizationsAPI.md#updateserviceauthorization) | Update service authorization
*ServiceAuthorizationsAPI* | [**UpdateServiceAuthorization2**](docs/ServiceAuthorizationsAPI.md#updateserviceauthorization2) | Update service authorizations
*SettingsAPI* | [**GetServiceSettings**](docs/SettingsAPI.md#getservicesettings) | Get service settings
*SettingsAPI* | [**UpdateServiceSettings**](docs/SettingsAPI.md#updateservicesettings) | Update service settings
*SnippetAPI* | [**CreateSnippet**](docs/SnippetAPI.md#createsnippet) | Create a snippet
*SnippetAPI* | [**DeleteSnippet**](docs/SnippetAPI.md#deletesnippet) | Delete a snippet
*SnippetAPI* | [**GetSnippet**](docs/SnippetAPI.md#getsnippet) | Get a versioned snippet
*SnippetAPI* | [**GetSnippetDynamic**](docs/SnippetAPI.md#getsnippetdynamic) | Get a dynamic snippet
*SnippetAPI* | [**ListSnippets**](docs/SnippetAPI.md#listsnippets) | List snippets
*SnippetAPI* | [**UpdateSnippet**](docs/SnippetAPI.md#updatesnippet) | Update a versioned snippet
*SnippetAPI* | [**UpdateSnippetDynamic**](docs/SnippetAPI.md#updatesnippetdynamic) | Update a dynamic snippet
*StarAPI* | [**CreateServiceStar**](docs/StarAPI.md#createservicestar) | Create a star
*StarAPI* | [**DeleteServiceStar**](docs/StarAPI.md#deleteservicestar) | Delete a star
*StarAPI* | [**GetServiceStar**](docs/StarAPI.md#getservicestar) | Get a star
*StarAPI* | [**ListServiceStars**](docs/StarAPI.md#listservicestars) | List stars
*StatsAPI* | [**GetServiceStats**](docs/StatsAPI.md#getservicestats) | Get stats for a service
*SudoAPI* | [**RequestSudoAccess**](docs/SudoAPI.md#requestsudoaccess) | Request Sudo access
*TLSActivationsAPI* | [**CreateTLSActivation**](docs/TlsActivationsAPI.md#createtlsactivation) | Enable TLS for a domain using a custom certificate
*TLSActivationsAPI* | [**DeleteTLSActivation**](docs/TlsActivationsAPI.md#deletetlsactivation) | Disable TLS on a domain
*TLSActivationsAPI* | [**GetTLSActivation**](docs/TlsActivationsAPI.md#gettlsactivation) | Get a TLS activation
*TLSActivationsAPI* | [**ListTLSActivations**](docs/TlsActivationsAPI.md#listtlsactivations) | List TLS activations
*TLSActivationsAPI* | [**UpdateTLSActivation**](docs/TlsActivationsAPI.md#updatetlsactivation) | Update a certificate
*TLSBulkCertificatesAPI* | [**DeleteBulkTLSCert**](docs/TlsBulkCertificatesAPI.md#deletebulktlscert) | Delete a certificate
*TLSBulkCertificatesAPI* | [**GetTLSBulkCert**](docs/TlsBulkCertificatesAPI.md#gettlsbulkcert) | Get a certificate
*TLSBulkCertificatesAPI* | [**ListTLSBulkCerts**](docs/TlsBulkCertificatesAPI.md#listtlsbulkcerts) | List certificates
*TLSBulkCertificatesAPI* | [**UpdateBulkTLSCert**](docs/TlsBulkCertificatesAPI.md#updatebulktlscert) | Update a certificate
*TLSBulkCertificatesAPI* | [**UploadTLSBulkCert**](docs/TlsBulkCertificatesAPI.md#uploadtlsbulkcert) | Upload a certificate
*TLSCertificatesAPI* | [**CreateTLSCert**](docs/TlsCertificatesAPI.md#createtlscert) | Create a TLS certificate
*TLSCertificatesAPI* | [**DeleteTLSCert**](docs/TlsCertificatesAPI.md#deletetlscert) | Delete a TLS certificate
*TLSCertificatesAPI* | [**GetTLSCert**](docs/TlsCertificatesAPI.md#gettlscert) | Get a TLS certificate
*TLSCertificatesAPI* | [**GetTLSCertBlob**](docs/TlsCertificatesAPI.md#gettlscertblob) | Get a TLS certificate blob (Limited Availability)
*TLSCertificatesAPI* | [**ListTLSCerts**](docs/TlsCertificatesAPI.md#listtlscerts) | List TLS certificates
*TLSCertificatesAPI* | [**UpdateTLSCert**](docs/TlsCertificatesAPI.md#updatetlscert) | Update a TLS certificate
*TLSConfigurationsAPI* | [**GetTLSConfig**](docs/TlsConfigurationsAPI.md#gettlsconfig) | Get a TLS configuration
*TLSConfigurationsAPI* | [**ListTLSConfigs**](docs/TlsConfigurationsAPI.md#listtlsconfigs) | List TLS configurations
*TLSConfigurationsAPI* | [**UpdateTLSConfig**](docs/TlsConfigurationsAPI.md#updatetlsconfig) | Update a TLS configuration
*TLSCsrsAPI* | [**CreateCsr**](docs/TlsCsrsAPI.md#createcsr) | Create CSR
*TLSDomainsAPI* | [**ListTLSDomains**](docs/TlsDomainsAPI.md#listtlsdomains) | List TLS domains
*TLSPrivateKeysAPI* | [**CreateTLSKey**](docs/TlsPrivateKeysAPI.md#createtlskey) | Create a TLS private key
*TLSPrivateKeysAPI* | [**DeleteTLSKey**](docs/TlsPrivateKeysAPI.md#deletetlskey) | Delete a TLS private key
*TLSPrivateKeysAPI* | [**GetTLSKey**](docs/TlsPrivateKeysAPI.md#gettlskey) | Get a TLS private key
*TLSPrivateKeysAPI* | [**ListTLSKeys**](docs/TlsPrivateKeysAPI.md#listtlskeys) | List TLS private keys
*TLSSubscriptionsAPI* | [**CreateGlobalsignEmailChallenge**](docs/TlsSubscriptionsAPI.md#createglobalsignemailchallenge) | Creates a GlobalSign email challenge.
*TLSSubscriptionsAPI* | [**CreateTLSSub**](docs/TlsSubscriptionsAPI.md#createtlssub) | Create a TLS subscription
*TLSSubscriptionsAPI* | [**DeleteGlobalsignEmailChallenge**](docs/TlsSubscriptionsAPI.md#deleteglobalsignemailchallenge) | Delete a GlobalSign email challenge
*TLSSubscriptionsAPI* | [**DeleteTLSSub**](docs/TlsSubscriptionsAPI.md#deletetlssub) | Delete a TLS subscription
*TLSSubscriptionsAPI* | [**GetTLSSub**](docs/TlsSubscriptionsAPI.md#gettlssub) | Get a TLS subscription
*TLSSubscriptionsAPI* | [**ListTLSSubs**](docs/TlsSubscriptionsAPI.md#listtlssubs) | List TLS subscriptions
*TLSSubscriptionsAPI* | [**PatchTLSSub**](docs/TlsSubscriptionsAPI.md#patchtlssub) | Update a TLS subscription
*TokensAPI* | [**BulkRevokeTokens**](docs/TokensAPI.md#bulkrevoketokens) | Revoke multiple tokens
*TokensAPI* | [**CreateToken**](docs/TokensAPI.md#createtoken) | Create a token
*TokensAPI* | [**GetToken**](docs/TokensAPI.md#gettoken) | Get a token
*TokensAPI* | [**GetTokenCurrent**](docs/TokensAPI.md#gettokencurrent) | Get the current token
*TokensAPI* | [**ListTokensCustomer**](docs/TokensAPI.md#listtokenscustomer) | List tokens for a customer
*TokensAPI* | [**ListTokensUser**](docs/TokensAPI.md#listtokensuser) | List tokens for the authenticated user
*TokensAPI* | [**RevokeToken**](docs/TokensAPI.md#revoketoken) | Revoke a token
*TokensAPI* | [**RevokeTokenCurrent**](docs/TokensAPI.md#revoketokencurrent) | Revoke the current token
*UserAPI* | [**CreateUser**](docs/UserAPI.md#createuser) | Create a user
*UserAPI* | [**DeleteUser**](docs/UserAPI.md#deleteuser) | Delete a user
*UserAPI* | [**GetCurrentUser**](docs/UserAPI.md#getcurrentuser) | Get the current user
*UserAPI* | [**GetUser**](docs/UserAPI.md#getuser) | Get a user
*UserAPI* | [**RequestPasswordReset**](docs/UserAPI.md#requestpasswordreset) | Request a password reset
*UserAPI* | [**UpdateUser**](docs/UserAPI.md#updateuser) | Update a user
*UserAPI* | [**UpdateUserPassword**](docs/UserAPI.md#updateuserpassword) | Update the user's password
*VclAPI* | [**CreateCustomVcl**](docs/VclAPI.md#createcustomvcl) | Create a custom VCL file
*VclAPI* | [**DeleteCustomVcl**](docs/VclAPI.md#deletecustomvcl) | Delete a custom VCL file
*VclAPI* | [**GetCustomVcl**](docs/VclAPI.md#getcustomvcl) | Get a custom VCL file
*VclAPI* | [**GetCustomVclBoilerplate**](docs/VclAPI.md#getcustomvclboilerplate) | Get boilerplate VCL
*VclAPI* | [**GetCustomVclGenerated**](docs/VclAPI.md#getcustomvclgenerated) | Get the generated VCL for a service
*VclAPI* | [**GetCustomVclGeneratedHighlighted**](docs/VclAPI.md#getcustomvclgeneratedhighlighted) | Get the generated VCL with syntax highlighting
*VclAPI* | [**GetCustomVclHighlighted**](docs/VclAPI.md#getcustomvclhighlighted) | Get a custom VCL file with syntax highlighting
*VclAPI* | [**GetCustomVclRaw**](docs/VclAPI.md#getcustomvclraw) | Download a custom VCL file
*VclAPI* | [**LintVclDefault**](docs/VclAPI.md#lintvcldefault) | Lint (validate) VCL using a default set of flags.
*VclAPI* | [**LintVclForService**](docs/VclAPI.md#lintvclforservice) | Lint (validate) VCL using flags set for the service.
*VclAPI* | [**ListCustomVcl**](docs/VclAPI.md#listcustomvcl) | List custom VCL files
*VclAPI* | [**SetCustomVclMain**](docs/VclAPI.md#setcustomvclmain) | Set a custom VCL file as main
*VclAPI* | [**UpdateCustomVcl**](docs/VclAPI.md#updatecustomvcl) | Update a custom VCL file
*VclDiffAPI* | [**VclDiffServiceVersions**](docs/VclDiffAPI.md#vcldiffserviceversions) | Get a comparison of the VCL changes between two service versions
*VersionAPI* | [**ActivateServiceVersion**](docs/VersionAPI.md#activateserviceversion) | Activate a service version
*VersionAPI* | [**ActivateServiceVersionEnvironment**](docs/VersionAPI.md#activateserviceversionenvironment) | Activate a service version on the specified environment
*VersionAPI* | [**CloneServiceVersion**](docs/VersionAPI.md#cloneserviceversion) | Clone a service version
*VersionAPI* | [**CreateServiceVersion**](docs/VersionAPI.md#createserviceversion) | Create a service version
*VersionAPI* | [**DeactivateServiceVersion**](docs/VersionAPI.md#deactivateserviceversion) | Deactivate a service version
*VersionAPI* | [**DeactivateServiceVersionEnvironment**](docs/VersionAPI.md#deactivateserviceversionenvironment) | Deactivate a service version on an environment
*VersionAPI* | [**GetServiceVersion**](docs/VersionAPI.md#getserviceversion) | Get a version of a service
*VersionAPI* | [**ListServiceVersions**](docs/VersionAPI.md#listserviceversions) | List versions of a service
*VersionAPI* | [**LockServiceVersion**](docs/VersionAPI.md#lockserviceversion) | Lock a service version
*VersionAPI* | [**UpdateServiceVersion**](docs/VersionAPI.md#updateserviceversion) | Update a service version
*VersionAPI* | [**ValidateServiceVersion**](docs/VersionAPI.md#validateserviceversion) | Validate a service version
*WafActiveRulesAPI* | [**BulkDeleteWafActiveRules**](docs/WafActiveRulesAPI.md#bulkdeletewafactiverules) | Delete multiple active rules from a WAF
*WafActiveRulesAPI* | [**BulkUpdateWafActiveRules**](docs/WafActiveRulesAPI.md#bulkupdatewafactiverules) | Update multiple active rules
*WafActiveRulesAPI* | [**CreateWafActiveRule**](docs/WafActiveRulesAPI.md#createwafactiverule) | Add a rule to a WAF as an active rule
*WafActiveRulesAPI* | [**CreateWafActiveRulesTag**](docs/WafActiveRulesAPI.md#createwafactiverulestag) | Create active rules by tag
*WafActiveRulesAPI* | [**DeleteWafActiveRule**](docs/WafActiveRulesAPI.md#deletewafactiverule) | Delete an active rule
*WafActiveRulesAPI* | [**GetWafActiveRule**](docs/WafActiveRulesAPI.md#getwafactiverule) | Get an active WAF rule object
*WafActiveRulesAPI* | [**ListWafActiveRules**](docs/WafActiveRulesAPI.md#listwafactiverules) | List active rules on a WAF
*WafActiveRulesAPI* | [**UpdateWafActiveRule**](docs/WafActiveRulesAPI.md#updatewafactiverule) | Update an active rule
*WafExclusionsAPI* | [**CreateWafRuleExclusion**](docs/WafExclusionsAPI.md#createwafruleexclusion) | Create a WAF rule exclusion
*WafExclusionsAPI* | [**DeleteWafRuleExclusion**](docs/WafExclusionsAPI.md#deletewafruleexclusion) | Delete a WAF rule exclusion
*WafExclusionsAPI* | [**GetWafRuleExclusion**](docs/WafExclusionsAPI.md#getwafruleexclusion) | Get a WAF rule exclusion
*WafExclusionsAPI* | [**ListWafRuleExclusions**](docs/WafExclusionsAPI.md#listwafruleexclusions) | List WAF rule exclusions
*WafExclusionsAPI* | [**UpdateWafRuleExclusion**](docs/WafExclusionsAPI.md#updatewafruleexclusion) | Update a WAF rule exclusion
*WafFirewallVersionsAPI* | [**CloneWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#clonewaffirewallversion) | Clone a firewall version
*WafFirewallVersionsAPI* | [**CreateWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#createwaffirewallversion) | Create a firewall version
*WafFirewallVersionsAPI* | [**DeployActivateWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#deployactivatewaffirewallversion) | Deploy or activate a firewall version
*WafFirewallVersionsAPI* | [**GetWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#getwaffirewallversion) | Get a firewall version
*WafFirewallVersionsAPI* | [**ListWafFirewallVersions**](docs/WafFirewallVersionsAPI.md#listwaffirewallversions) | List firewall versions
*WafFirewallVersionsAPI* | [**UpdateWafFirewallVersion**](docs/WafFirewallVersionsAPI.md#updatewaffirewallversion) | Update a firewall version
*WafFirewallsAPI* | [**CreateWafFirewall**](docs/WafFirewallsAPI.md#createwaffirewall) | Create a firewall
*WafFirewallsAPI* | [**DeleteWafFirewall**](docs/WafFirewallsAPI.md#deletewaffirewall) | Delete a firewall
*WafFirewallsAPI* | [**GetWafFirewall**](docs/WafFirewallsAPI.md#getwaffirewall) | Get a firewall
*WafFirewallsAPI* | [**ListWafFirewalls**](docs/WafFirewallsAPI.md#listwaffirewalls) | List firewalls
*WafFirewallsAPI* | [**UpdateWafFirewall**](docs/WafFirewallsAPI.md#updatewaffirewall) | Update a firewall
*WafRuleRevisionsAPI* | [**GetWafRuleRevision**](docs/WafRuleRevisionsAPI.md#getwafrulerevision) | Get a revision of a rule
*WafRuleRevisionsAPI* | [**ListWafRuleRevisions**](docs/WafRuleRevisionsAPI.md#listwafrulerevisions) | List revisions for a rule
*WafRulesAPI* | [**GetWafRule**](docs/WafRulesAPI.md#getwafrule) | Get a rule
*WafRulesAPI* | [**ListWafRules**](docs/WafRulesAPI.md#listwafrules) | List available WAF rules
*WafTagsAPI* | [**ListWafTags**](docs/WafTagsAPI.md#listwaftags) | List tags
*WholePlatformDdosHistoricalAPI* | [**GetPlatformDdosHistorical**](docs/WholePlatformDdosHistoricalAPI.md#getplatformddoshistorical) | Get historical DDoS metrics for the entire Fastly platform


</details>

## Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Issues

The fastly-go API client currently does not support the following endpoints:

- [`/alerts/definitions/{definition_id}`](https://www.fastly.com/documentation/reference/api/observability/alerts/definitions) (DELETE, GET, PUT)
- [`/alerts/definitions`](https://www.fastly.com/documentation/reference/api/observability/alerts/definitions) (GET, POST)
- [`/alerts/history`](https://www.fastly.com/documentation/reference/api/observability/alerts/history) (GET)
- [`/dns/configurations/{dns_configuration_id}`](https://www.fastly.com/documentation/reference/api/) (DELETE, GET, PATCH)
- [`/dns/configurations`](https://www.fastly.com/documentation/reference/api/) (GET, POST)
- [`/domains/{domain_id}`](https://www.fastly.com/documentation/reference/api/) (DELETE, GET, PATCH)
- [`/domains`](https://www.fastly.com/documentation/reference/api/) (GET, POST)
- [`/notifications/integration-types`](https://developer.fastly.com/reference/api/observability/notification) (GET)
- [`/notifications/integrations/{integration_id}/rotateSigningKey`](https://developer.fastly.com/reference/api/observability/notification) (POST)
- [`/notifications/integrations/{integration_id}/signingKey`](https://developer.fastly.com/reference/api/observability/notification) (GET)
- [`/notifications/integrations/{integration_id}`](https://developer.fastly.com/reference/api/observability/notification) (DELETE, GET, PATCH)
- [`/notifications/integrations`](https://developer.fastly.com/reference/api/observability/notification) (GET, POST)
- [`/notifications/mailinglist-confirmations`](https://developer.fastly.com/reference/api/observability/notification) (POST)
- [`/resources/stores/kv/{store_id}/batch`](https://www.fastly.com/documentation/reference/api/services/resources/kv-store-item) (PUT)
- [`/security/workspaces/{workspace_id}/events/{event_id}`](https://docs.fastly.com/en/ngwaf/) (GET, PATCH)
- [`/security/workspaces/{workspace_id}/events`](https://docs.fastly.com/en/ngwaf/) (GET)
- [`/security/workspaces/{workspace_id}/redactions/{redaction_id}`](https://docs.fastly.com/en/ngwaf/) (DELETE, GET, PATCH)
- [`/security/workspaces/{workspace_id}/redactions`](https://docs.fastly.com/en/ngwaf/) (GET, POST)
- [`/security/workspaces/{workspace_id}/requests/{request_id}`](https://docs.fastly.com/en/ngwaf/) (GET)
- [`/security/workspaces/{workspace_id}/requests`](https://docs.fastly.com/en/ngwaf/) (GET)
- [`/security/workspaces/{workspace_id}/rules/{rule_id}`](https://docs.fastly.com/en/ngwaf/) (DELETE, GET, PATCH)
- [`/security/workspaces/{workspace_id}/rules`](https://docs.fastly.com/en/ngwaf/) (GET, POST)
- [`/security/workspaces/{workspace_id}/timeseries`](https://docs.fastly.com/en/ngwaf/) (GET)
- [`/security/workspaces/{workspace_id}/virtual-patches/{virtual_patch_id}`](https://docs.fastly.com/en/ngwaf/) (GET, PATCH)
- [`/security/workspaces/{workspace_id}/virtual-patches`](https://docs.fastly.com/en/ngwaf/) (GET)
- [`/security/workspaces/{workspace_id}`](https://docs.fastly.com/en/ngwaf/) (DELETE, GET, PATCH)
- [`/security/workspaces`](https://docs.fastly.com/en/ngwaf/) (GET, POST)
- [`/tls/activations/{tls_activation_id}`](https://www.fastly.com/documentation/reference/api/tls/mutual-tls/activations) (GET, PATCH)
- [`/tls/activations`](https://www.fastly.com/documentation/reference/api/tls/mutual-tls/activations) (GET)
- [`/tls/configurations/{tls_configuration_id}`](https://www.fastly.com/documentation/reference/api/) (DELETE, GET, PATCH)
- [`/tls/configurations`](https://www.fastly.com/documentation/reference/api/) (GET, POST)
- [`/v1/channel/{service_id}/ts/h/limit/{max_entries}`](https://www.fastly.com/documentation/reference/api/metrics-stats/origin-insights) (GET)
- [`/v1/channel/{service_id}/ts/h`](https://www.fastly.com/documentation/reference/api/metrics-stats/origin-insights) (GET)
- [`/v1/channel/{service_id}/ts/{start_timestamp}`](https://www.fastly.com/documentation/reference/api/metrics-stats/origin-insights) (GET)


If you encounter any non-security-related bug or unexpected behavior, please [file an issue][bug]
using the bug report template.

[bug]: https://github.com/fastly/fastly-go/issues/new?labels=bug

### Security issues

Please see our [SECURITY.md](./SECURITY.md) for guidance on reporting security-related issues.

## License

[MIT](./LICENSE).
