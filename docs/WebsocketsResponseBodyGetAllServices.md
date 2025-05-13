# WebsocketsResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**WebsocketsResponseProductProduct**](WebsocketsResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**BotManagementResponseCustomerCustomer**](BotManagementResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services with Websockets enabled. | [optional] 
**Links** | Pointer to [**WebsocketsResponseLinksGetAllServicesLinks**](WebsocketsResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewWebsocketsResponseBodyGetAllServices

`func NewWebsocketsResponseBodyGetAllServices() *WebsocketsResponseBodyGetAllServices`

NewWebsocketsResponseBodyGetAllServices instantiates a new WebsocketsResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsocketsResponseBodyGetAllServicesWithDefaults

`func NewWebsocketsResponseBodyGetAllServicesWithDefaults() *WebsocketsResponseBodyGetAllServices`

NewWebsocketsResponseBodyGetAllServicesWithDefaults instantiates a new WebsocketsResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *WebsocketsResponseBodyGetAllServices) GetProduct() WebsocketsResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *WebsocketsResponseBodyGetAllServices) GetProductOk() (*WebsocketsResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *WebsocketsResponseBodyGetAllServices) SetProduct(v WebsocketsResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *WebsocketsResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *WebsocketsResponseBodyGetAllServices) GetCustomer() BotManagementResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *WebsocketsResponseBodyGetAllServices) GetCustomerOk() (*BotManagementResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *WebsocketsResponseBodyGetAllServices) SetCustomer(v BotManagementResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *WebsocketsResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *WebsocketsResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *WebsocketsResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *WebsocketsResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *WebsocketsResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *WebsocketsResponseBodyGetAllServices) GetLinks() WebsocketsResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *WebsocketsResponseBodyGetAllServices) GetLinksOk() (*WebsocketsResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *WebsocketsResponseBodyGetAllServices) SetLinks(v WebsocketsResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *WebsocketsResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
