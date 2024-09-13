# ListCustomerAddressesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CustomerAddress**](CustomerAddress.md) |  | [optional] 

## Methods

### NewListCustomerAddressesResponse

`func NewListCustomerAddressesResponse() *ListCustomerAddressesResponse`

NewListCustomerAddressesResponse instantiates a new ListCustomerAddressesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomerAddressesResponseWithDefaults

`func NewListCustomerAddressesResponseWithDefaults() *ListCustomerAddressesResponse`

NewListCustomerAddressesResponseWithDefaults instantiates a new ListCustomerAddressesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCustomerAddressesResponse) GetData() []CustomerAddress`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCustomerAddressesResponse) GetDataOk() (*[]CustomerAddress, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCustomerAddressesResponse) SetData(v []CustomerAddress)`

SetData sets Data field to given value.

### HasData

`func (o *ListCustomerAddressesResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
