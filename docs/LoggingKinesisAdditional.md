# LoggingKinesisAdditional

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for the real-time logging configuration. | [optional] 
**Placement** | Pointer to [**NullableLoggingPlacement**](LoggingPlacement.md) |  | [optional] 
**Format** | Pointer to **string** | A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/). | [optional] [default to "{\"timestamp\":\"%{begin:%Y-%m-%dT%H:%M:%S}t\",\"time_elapsed\":\"%{time.elapsed.usec}V\",\"is_tls\":\"%{if(req.is_ssl, \\\"true\\\", \\\"false\\\")}V\",\"client_ip\":\"%{req.http.Fastly-Client-IP}V\",\"geo_city\":\"%{client.geo.city}V\",\"geo_country_code\":\"%{client.geo.country_code}V\",\"request\":\"%{req.request}V\",\"host\":\"%{req.http.Fastly-Orig-Host}V\",\"url\":\"%{json.escape(req.url)}V\",\"request_referer\":\"%{json.escape(req.http.Referer)}V\",\"request_user_agent\":\"%{json.escape(req.http.User-Agent)}V\",\"request_accept_language\":\"%{json.escape(req.http.Accept-Language)}V\",\"request_accept_charset\":\"%{json.escape(req.http.Accept-Charset)}V\",\"cache_status\":\"%{regsub(fastly_info.state, \\\"^(HIT-(SYNTH)|(HITPASS|HIT|MISS|PASS|ERROR|PIPE)).*\\\", \\\"\\\\2\\\\3\\\") }V\"}"]
**Topic** | Pointer to **string** | The Amazon Kinesis stream to send logs to. Required. | [optional] 
**Region** | Pointer to [**AwsRegion**](AwsRegion.md) |  | [optional] 
**SecretKey** | Pointer to **NullableString** | The secret key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified. | [optional] 
**AccessKey** | Pointer to **NullableString** | The access key associated with the target Amazon Kinesis stream. Not required if `iam_role` is specified. | [optional] 
**IamRole** | Pointer to **NullableString** | The ARN for an IAM role granting Fastly access to the target Amazon Kinesis stream. Not required if `access_key` and `secret_key` are provided. | [optional] 

## Methods

### NewLoggingKinesisAdditional

`func NewLoggingKinesisAdditional() *LoggingKinesisAdditional`

NewLoggingKinesisAdditional instantiates a new LoggingKinesisAdditional object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingKinesisAdditionalWithDefaults

`func NewLoggingKinesisAdditionalWithDefaults() *LoggingKinesisAdditional`

NewLoggingKinesisAdditionalWithDefaults instantiates a new LoggingKinesisAdditional object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoggingKinesisAdditional) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoggingKinesisAdditional) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoggingKinesisAdditional) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoggingKinesisAdditional) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *LoggingKinesisAdditional) GetPlacement() LoggingPlacement`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *LoggingKinesisAdditional) GetPlacementOk() (*LoggingPlacement, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *LoggingKinesisAdditional) SetPlacement(v LoggingPlacement)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *LoggingKinesisAdditional) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### SetPlacementNil

`func (o *LoggingKinesisAdditional) SetPlacementNil(b bool)`

 SetPlacementNil sets the value for Placement to be an explicit nil

### UnsetPlacement
`func (o *LoggingKinesisAdditional) UnsetPlacement()`

UnsetPlacement ensures that no value is present for Placement, not even an explicit nil
### GetFormat

`func (o *LoggingKinesisAdditional) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *LoggingKinesisAdditional) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *LoggingKinesisAdditional) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *LoggingKinesisAdditional) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetTopic

`func (o *LoggingKinesisAdditional) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *LoggingKinesisAdditional) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *LoggingKinesisAdditional) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *LoggingKinesisAdditional) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetRegion

`func (o *LoggingKinesisAdditional) GetRegion() AwsRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *LoggingKinesisAdditional) GetRegionOk() (*AwsRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *LoggingKinesisAdditional) SetRegion(v AwsRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *LoggingKinesisAdditional) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretKey

`func (o *LoggingKinesisAdditional) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *LoggingKinesisAdditional) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *LoggingKinesisAdditional) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *LoggingKinesisAdditional) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### SetSecretKeyNil

`func (o *LoggingKinesisAdditional) SetSecretKeyNil(b bool)`

 SetSecretKeyNil sets the value for SecretKey to be an explicit nil

### UnsetSecretKey
`func (o *LoggingKinesisAdditional) UnsetSecretKey()`

UnsetSecretKey ensures that no value is present for SecretKey, not even an explicit nil
### GetAccessKey

`func (o *LoggingKinesisAdditional) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *LoggingKinesisAdditional) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *LoggingKinesisAdditional) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *LoggingKinesisAdditional) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### SetAccessKeyNil

`func (o *LoggingKinesisAdditional) SetAccessKeyNil(b bool)`

 SetAccessKeyNil sets the value for AccessKey to be an explicit nil

### UnsetAccessKey
`func (o *LoggingKinesisAdditional) UnsetAccessKey()`

UnsetAccessKey ensures that no value is present for AccessKey, not even an explicit nil
### GetIamRole

`func (o *LoggingKinesisAdditional) GetIamRole() string`

GetIamRole returns the IamRole field if non-nil, zero value otherwise.

### GetIamRoleOk

`func (o *LoggingKinesisAdditional) GetIamRoleOk() (*string, bool)`

GetIamRoleOk returns a tuple with the IamRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIamRole

`func (o *LoggingKinesisAdditional) SetIamRole(v string)`

SetIamRole sets IamRole field to given value.

### HasIamRole

`func (o *LoggingKinesisAdditional) HasIamRole() bool`

HasIamRole returns a boolean if a field has been set.

### SetIamRoleNil

`func (o *LoggingKinesisAdditional) SetIamRoleNil(b bool)`

 SetIamRoleNil sets the value for IamRole to be an explicit nil

### UnsetIamRole
`func (o *LoggingKinesisAdditional) UnsetIamRole()`

UnsetIamRole ensures that no value is present for IamRole, not even an explicit nil

[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)


