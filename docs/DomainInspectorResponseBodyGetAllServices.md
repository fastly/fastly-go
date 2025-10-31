# DomainInspectorResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**DomainInspectorResponseProductProduct**](DomainInspectorResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with Domain Inspector enabled. | [optional] 
**Links** | Pointer to [**DomainInspectorResponseLinksGetAllServicesLinks**](DomainInspectorResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewDomainInspectorResponseBodyGetAllServices

`func NewDomainInspectorResponseBodyGetAllServices() *DomainInspectorResponseBodyGetAllServices`

NewDomainInspectorResponseBodyGetAllServices instantiates a new DomainInspectorResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainInspectorResponseBodyGetAllServicesWithDefaults

`func NewDomainInspectorResponseBodyGetAllServicesWithDefaults() *DomainInspectorResponseBodyGetAllServices`

NewDomainInspectorResponseBodyGetAllServicesWithDefaults instantiates a new DomainInspectorResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *DomainInspectorResponseBodyGetAllServices) GetProduct() DomainInspectorResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *DomainInspectorResponseBodyGetAllServices) GetProductOk() (*DomainInspectorResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *DomainInspectorResponseBodyGetAllServices) SetProduct(v DomainInspectorResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *DomainInspectorResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *DomainInspectorResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *DomainInspectorResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *DomainInspectorResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *DomainInspectorResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *DomainInspectorResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DomainInspectorResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DomainInspectorResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *DomainInspectorResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *DomainInspectorResponseBodyGetAllServices) GetLinks() DomainInspectorResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainInspectorResponseBodyGetAllServices) GetLinksOk() (*DomainInspectorResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainInspectorResponseBodyGetAllServices) SetLinks(v DomainInspectorResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DomainInspectorResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


