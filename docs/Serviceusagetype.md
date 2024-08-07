# Serviceusagetype

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductID** | Pointer to **string** | The product identifier associated with the usage type. This corresponds to a Fastly product offering. | [optional] 
**Name** | Pointer to **string** | Full name of the product usage type as it might appear on a customer&#39;s invoice. | [optional] 

## Methods

### NewServiceusagetype

`func NewServiceusagetype() *Serviceusagetype`

NewServiceusagetype instantiates a new Serviceusagetype object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceusagetypeWithDefaults

`func NewServiceusagetypeWithDefaults() *Serviceusagetype`

NewServiceusagetypeWithDefaults instantiates a new Serviceusagetype object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductID

`func (o *Serviceusagetype) GetProductID() string`

GetProductID returns the ProductID field if non-nil, zero value otherwise.

### GetProductIDOk

`func (o *Serviceusagetype) GetProductIDOk() (*string, bool)`

GetProductIDOk returns a tuple with the ProductID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductID

`func (o *Serviceusagetype) SetProductID(v string)`

SetProductID sets ProductID field to given value.

### HasProductID

`func (o *Serviceusagetype) HasProductID() bool`

HasProductID returns a boolean if a field has been set.

### GetName

`func (o *Serviceusagetype) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Serviceusagetype) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Serviceusagetype) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Serviceusagetype) HasName() bool`

HasName returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
