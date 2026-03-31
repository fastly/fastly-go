# OperationBulkAddTags

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperationIds** | **[]string** | List of operation IDs to add tags to. | 
**TagIds** | **[]string** | List of tag IDs to add to the operations. | 

## Methods

### NewOperationBulkAddTags

`func NewOperationBulkAddTags(operationIds []string, tagIds []string, ) *OperationBulkAddTags`

NewOperationBulkAddTags instantiates a new OperationBulkAddTags object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationBulkAddTagsWithDefaults

`func NewOperationBulkAddTagsWithDefaults() *OperationBulkAddTags`

NewOperationBulkAddTagsWithDefaults instantiates a new OperationBulkAddTags object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperationIds

`func (o *OperationBulkAddTags) GetOperationIds() []string`

GetOperationIds returns the OperationIds field if non-nil, zero value otherwise.

### GetOperationIdsOk

`func (o *OperationBulkAddTags) GetOperationIdsOk() (*[]string, bool)`

GetOperationIdsOk returns a tuple with the OperationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationIds

`func (o *OperationBulkAddTags) SetOperationIds(v []string)`

SetOperationIds sets OperationIds field to given value.


### GetTagIds

`func (o *OperationBulkAddTags) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OperationBulkAddTags) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OperationBulkAddTags) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


