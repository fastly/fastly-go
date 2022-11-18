# WafExclusionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafExclusion**](TypeWafExclusion.md) |  | [optional] [default to TYPEWAFEXCLUSION_WAF_EXCLUSION]
**Attributes** | Pointer to [**WafExclusionResponseDataAttributes**](WafExclusionResponseDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**WafExclusionResponseDataRelationships**](WafExclusionResponseDataRelationships.md) |  | [optional] 
**ID** | Pointer to **string** | Alphanumeric string identifying a WAF exclusion. | [optional] [readonly] 

## Methods

### NewWafExclusionResponseData

`func NewWafExclusionResponseData() *WafExclusionResponseData`

NewWafExclusionResponseData instantiates a new WafExclusionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafExclusionResponseDataWithDefaults

`func NewWafExclusionResponseDataWithDefaults() *WafExclusionResponseData`

NewWafExclusionResponseDataWithDefaults instantiates a new WafExclusionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafExclusionResponseData) GetType() TypeWafExclusion`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafExclusionResponseData) GetTypeOk() (*TypeWafExclusion, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafExclusionResponseData) SetType(v TypeWafExclusion)`

SetType sets Type field to given value.

### HasType

`func (o *WafExclusionResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WafExclusionResponseData) GetAttributes() WafExclusionResponseDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafExclusionResponseData) GetAttributesOk() (*WafExclusionResponseDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafExclusionResponseData) SetAttributes(v WafExclusionResponseDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafExclusionResponseData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafExclusionResponseData) GetRelationships() WafExclusionResponseDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafExclusionResponseData) GetRelationshipsOk() (*WafExclusionResponseDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafExclusionResponseData) SetRelationships(v WafExclusionResponseDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafExclusionResponseData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetID

`func (o *WafExclusionResponseData) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *WafExclusionResponseData) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *WafExclusionResponseData) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *WafExclusionResponseData) HasID() bool`

HasID returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
