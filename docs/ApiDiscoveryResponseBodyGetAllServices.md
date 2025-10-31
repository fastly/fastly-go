# ApiDiscoveryResponseBodyGetAllServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | Pointer to [**ApiDiscoveryResponseProductProduct**](ApiDiscoveryResponseProductProduct.md) |  | [optional] 
**Customer** | Pointer to [**ApiDiscoveryResponseCustomerCustomer**](ApiDiscoveryResponseCustomerCustomer.md) |  | [optional] 
**Services** | Pointer to **[]string** | A list of services for a customer with API Discovery enabled. | [optional] 
**Links** | Pointer to [**ApiDiscoveryResponseLinksGetAllServicesLinks**](ApiDiscoveryResponseLinksGetAllServicesLinks.md) |  | [optional] 

## Methods

### NewApiDiscoveryResponseBodyGetAllServices

`func NewApiDiscoveryResponseBodyGetAllServices() *ApiDiscoveryResponseBodyGetAllServices`

NewApiDiscoveryResponseBodyGetAllServices instantiates a new ApiDiscoveryResponseBodyGetAllServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDiscoveryResponseBodyGetAllServicesWithDefaults

`func NewApiDiscoveryResponseBodyGetAllServicesWithDefaults() *ApiDiscoveryResponseBodyGetAllServices`

NewApiDiscoveryResponseBodyGetAllServicesWithDefaults instantiates a new ApiDiscoveryResponseBodyGetAllServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetProduct() ApiDiscoveryResponseProductProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetProductOk() (*ApiDiscoveryResponseProductProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *ApiDiscoveryResponseBodyGetAllServices) SetProduct(v ApiDiscoveryResponseProductProduct)`

SetProduct sets Product field to given value.

### HasProduct

`func (o *ApiDiscoveryResponseBodyGetAllServices) HasProduct() bool`

HasProduct returns a boolean if a field has been set.

### GetCustomer

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetCustomer() ApiDiscoveryResponseCustomerCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetCustomerOk() (*ApiDiscoveryResponseCustomerCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *ApiDiscoveryResponseBodyGetAllServices) SetCustomer(v ApiDiscoveryResponseCustomerCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *ApiDiscoveryResponseBodyGetAllServices) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetServices

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *ApiDiscoveryResponseBodyGetAllServices) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *ApiDiscoveryResponseBodyGetAllServices) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetLinks

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetLinks() ApiDiscoveryResponseLinksGetAllServicesLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApiDiscoveryResponseBodyGetAllServices) GetLinksOk() (*ApiDiscoveryResponseLinksGetAllServicesLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApiDiscoveryResponseBodyGetAllServices) SetLinks(v ApiDiscoveryResponseLinksGetAllServicesLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApiDiscoveryResponseBodyGetAllServices) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


