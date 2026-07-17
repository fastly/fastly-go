# PathCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | The URL path pattern, beginning with `/`. Maximum 2048 characters. | 

## Methods

### NewPathCreate

`func NewPathCreate(path string, ) *PathCreate`

NewPathCreate instantiates a new PathCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathCreateWithDefaults

`func NewPathCreateWithDefaults() *PathCreate`

NewPathCreateWithDefaults instantiates a new PathCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *PathCreate) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathCreate) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathCreate) SetPath(v string)`

SetPath sets Path field to given value.



[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


