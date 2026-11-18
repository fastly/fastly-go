# BillingAddressResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** | Alphanumeric string identifying the billing address. | [optional] [readonly] 
**Attributes** | Pointer to [**BillingAddressAttributes**](BillingAddressAttributes.md) |  | [optional] 
**Type** | Pointer to [**TypeBillingAddress**](TypeBillingAddress.md) |  | [optional] [default to TYPEBILLINGADDRESS_BILLING_ADDRESS]
**Relationships** | Pointer to [**RelationshipCustomer**](RelationshipCustomer.md) |  | [optional] 

## Methods

### NewBillingAddressResponseData

`func NewBillingAddressResponseData() *BillingAddressResponseData`

NewBillingAddressResponseData instantiates a new BillingAddressResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAddressResponseDataWithDefaults

`func NewBillingAddressResponseDataWithDefaults() *BillingAddressResponseData`

NewBillingAddressResponseDataWithDefaults instantiates a new BillingAddressResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *BillingAddressResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *BillingAddressResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *BillingAddressResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *BillingAddressResponseData) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAttributes

`func (o *BillingAddressResponseData) GetAttributes() BillingAddressAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *BillingAddressResponseData) GetAttributesOk() (*BillingAddressAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *BillingAddressResponseData) SetAttributes(v BillingAddressAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *BillingAddressResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *BillingAddressResponseData) GetType() TypeBillingAddress`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillingAddressResponseData) GetTypeOk() (*TypeBillingAddress, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillingAddressResponseData) SetType(v TypeBillingAddress)`

SetType sets Type field to given value.

### HasType

`func (o *BillingAddressResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRelationships

`func (o *BillingAddressResponseData) GetRelationships() RelationshipCustomer`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *BillingAddressResponseData) GetRelationshipsOk() (*RelationshipCustomer, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *BillingAddressResponseData) SetRelationships(v RelationshipCustomer)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *BillingAddressResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
