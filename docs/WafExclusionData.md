# WafExclusionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**TypeWafExclusion**](TypeWafExclusion.md) |  | [optional] [default to TYPEWAFEXCLUSION_WAF_EXCLUSION]
**Attributes** | Pointer to [**WafExclusionDataAttributes**](WafExclusionDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**RelationshipsForWafExclusion**](RelationshipsForWafExclusion.md) |  | [optional] 

## Methods

### NewWafExclusionData

`func NewWafExclusionData() *WafExclusionData`

NewWafExclusionData instantiates a new WafExclusionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafExclusionDataWithDefaults

`func NewWafExclusionDataWithDefaults() *WafExclusionData`

NewWafExclusionDataWithDefaults instantiates a new WafExclusionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WafExclusionData) GetType() TypeWafExclusion`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WafExclusionData) GetTypeOk() (*TypeWafExclusion, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WafExclusionData) SetType(v TypeWafExclusion)`

SetType sets Type field to given value.

### HasType

`func (o *WafExclusionData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAttributes

`func (o *WafExclusionData) GetAttributes() WafExclusionDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *WafExclusionData) GetAttributesOk() (*WafExclusionDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *WafExclusionData) SetAttributes(v WafExclusionDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *WafExclusionData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *WafExclusionData) GetRelationships() RelationshipsForWafExclusion`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *WafExclusionData) GetRelationshipsOk() (*RelationshipsForWafExclusion, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *WafExclusionData) SetRelationships(v RelationshipsForWafExclusion)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *WafExclusionData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
