# PathResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Alphanumeric string identifying the path. Stable across versions of the routing config. | [optional] [readonly] 
**Path** | Pointer to **string** | The URL path pattern, beginning with `/`. Maximum 2048 characters. | [optional] 
**Links** | Pointer to **map[string]string** | HATEOAS links to related resources. | [optional] [readonly] 

## Methods

### NewPathResponseAllOf

`func NewPathResponseAllOf() *PathResponseAllOf`

NewPathResponseAllOf instantiates a new PathResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathResponseAllOfWithDefaults

`func NewPathResponseAllOfWithDefaults() *PathResponseAllOf`

NewPathResponseAllOfWithDefaults instantiates a new PathResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PathResponseAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PathResponseAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PathResponseAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PathResponseAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *PathResponseAllOf) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *PathResponseAllOf) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *PathResponseAllOf) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *PathResponseAllOf) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetLinks

`func (o *PathResponseAllOf) GetLinks() map[string]string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PathResponseAllOf) GetLinksOk() (*map[string]string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PathResponseAllOf) SetLinks(v map[string]string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PathResponseAllOf) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


