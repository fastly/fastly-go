# FanoutResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**FanoutResponseProductProduct**](FanoutResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with Fanout enabled. | [optional] 
**Links** | Pointer to [**FanoutResponseLinksGetAllServicesLinks**](FanoutResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewFanoutResponseBodyGetAllServices

`func NewFanoutResponseBodyGetAllServices() *FanoutResponseBodyGetAllServices`

NewFanoutResponseBodyGetAllServices instantiates a new FanoutResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFanoutResponseBodyGetAllServicesWithDefaults

`func NewFanoutResponseBodyGetAllServicesWithDefaults() *FanoutResponseBodyGetAllServices`

NewFanoutResponseBodyGetAllServicesWithDefaults instantiates a new FanoutResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *FanoutResponseBodyGetAllServices) GetProduct() FanoutResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *FanoutResponseBodyGetAllServices) GetProductOk() (*FanoutResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *FanoutResponseBodyGetAllServices) SetProduct(v FanoutResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *FanoutResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *FanoutResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *FanoutResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *FanoutResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *FanoutResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *FanoutResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *FanoutResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *FanoutResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *FanoutResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *FanoutResponseBodyGetAllServices) GetLinks() FanoutResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FanoutResponseBodyGetAllServices) GetLinksOk() (*FanoutResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FanoutResponseBodyGetAllServices) SetLinks(v FanoutResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FanoutResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


