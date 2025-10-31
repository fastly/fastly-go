# NgwafResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**NgwafResponseProductProduct**](NgwafResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with the Next-Gen WAF enabled. | [optional] 
**Links** | Pointer to [**NgwafResponseLinksGetAllServicesLinks**](NgwafResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewNgwafResponseBodyGetAllServices

`func NewNgwafResponseBodyGetAllServices() *NgwafResponseBodyGetAllServices`

NewNgwafResponseBodyGetAllServices instantiates a new NgwafResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNgwafResponseBodyGetAllServicesWithDefaults

`func NewNgwafResponseBodyGetAllServicesWithDefaults() *NgwafResponseBodyGetAllServices`

NewNgwafResponseBodyGetAllServicesWithDefaults instantiates a new NgwafResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *NgwafResponseBodyGetAllServices) GetProduct() NgwafResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *NgwafResponseBodyGetAllServices) GetProductOk() (*NgwafResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *NgwafResponseBodyGetAllServices) SetProduct(v NgwafResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *NgwafResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *NgwafResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *NgwafResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *NgwafResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *NgwafResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *NgwafResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *NgwafResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *NgwafResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *NgwafResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *NgwafResponseBodyGetAllServices) GetLinks() NgwafResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NgwafResponseBodyGetAllServices) GetLinksOk() (*NgwafResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NgwafResponseBodyGetAllServices) SetLinks(v NgwafResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NgwafResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


