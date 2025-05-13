# DdosProtectionResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**DdosProtectionResponseProductProduct**](DdosProtectionResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with DDoS Protection enabled. | [optional] 
**Links** | Pointer to [**DdosProtectionResponseLinksGetAllServicesLinks**](DdosProtectionResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewDdosProtectionResponseBodyGetAllServices

`func NewDdosProtectionResponseBodyGetAllServices() *DdosProtectionResponseBodyGetAllServices`

NewDdosProtectionResponseBodyGetAllServices instantiates a new DdosProtectionResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosProtectionResponseBodyGetAllServicesWithDefaults

`func NewDdosProtectionResponseBodyGetAllServicesWithDefaults() *DdosProtectionResponseBodyGetAllServices`

NewDdosProtectionResponseBodyGetAllServicesWithDefaults instantiates a new DdosProtectionResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *DdosProtectionResponseBodyGetAllServices) GetProduct() DdosProtectionResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *DdosProtectionResponseBodyGetAllServices) GetProductOk() (*DdosProtectionResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *DdosProtectionResponseBodyGetAllServices) SetProduct(v DdosProtectionResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *DdosProtectionResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *DdosProtectionResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DdosProtectionResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DdosProtectionResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DdosProtectionResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *DdosProtectionResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DdosProtectionResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DdosProtectionResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *DdosProtectionResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *DdosProtectionResponseBodyGetAllServices) GetLinks() DdosProtectionResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DdosProtectionResponseBodyGetAllServices) GetLinksOk() (*DdosProtectionResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DdosProtectionResponseBodyGetAllServices) SetLinks(v DdosProtectionResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DdosProtectionResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
