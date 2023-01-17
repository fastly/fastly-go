# LoggingKafkaAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** | The Kafka topic to send logs to. Required. | [optional] 
**Brokers** | Pointer to **string** | A comma-separated list of IP addresses or hostnames of Kafka brokers. Required. | [optional] 
**CompressionCodec** | Pointer to **NullableString** | The codec used for compression of your logs. | [optional] 
**RequiredAcks** | Pointer to **int32** | The number of acknowledgements a leader must receive before a write is considered successful. | [optional] [default to 1]
**RequestMaxBytes** | Pointer to **int32** | The maximum number of bytes sent in one request. Defaults `0` (no limit). | [optional] [default to 0]
**ParseLogKeyvals** | Pointer to **bool** | Enables parsing of key&#x3D;value tuples from the beginning of a logline, turning them into [record headers](https://cwiki.apache.org/confluence/display/KAFKA/KIP-82+-+Add+Record+Headers). | [optional] 
**AuthMethod** | Pointer to **string** | SASL authentication method. | [optional] 
**User** | Pointer to **string** | SASL user. | [optional] 
**Password** | Pointer to **string** | SASL password. | [optional] 
**UseTLS** | Pointer to [**LoggingUseTLS**](LoggingUseTLS.md) |  | [optional] [default to LOGGINGUSETLS_no_tls]

## Methods

### NewLoggingKafkaAllOf

`func NewLoggingKafkaAllOf() *LoggingKafkaAllOf`

NewLoggingKafkaAllOf instantiates a new LoggingKafkaAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoggingKafkaAllOfWithDefaults

`func NewLoggingKafkaAllOfWithDefaults() *LoggingKafkaAllOf`

NewLoggingKafkaAllOfWithDefaults instantiates a new LoggingKafkaAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *LoggingKafkaAllOf) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *LoggingKafkaAllOf) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *LoggingKafkaAllOf) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *LoggingKafkaAllOf) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetBrokers

`func (o *LoggingKafkaAllOf) GetBrokers() string`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *LoggingKafkaAllOf) GetBrokersOk() (*string, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *LoggingKafkaAllOf) SetBrokers(v string)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *LoggingKafkaAllOf) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetCompressionCodec

`func (o *LoggingKafkaAllOf) GetCompressionCodec() string`

GetCompressionCodec returns the CompressionCodec field if non-nil, zero value otherwise.

### GetCompressionCodecOk

`func (o *LoggingKafkaAllOf) GetCompressionCodecOk() (*string, bool)`

GetCompressionCodecOk returns a tuple with the CompressionCodec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionCodec

`func (o *LoggingKafkaAllOf) SetCompressionCodec(v string)`

SetCompressionCodec sets CompressionCodec field to given value.

### HasCompressionCodec

`func (o *LoggingKafkaAllOf) HasCompressionCodec() bool`

HasCompressionCodec returns a boolean if a field has been set.

### SetCompressionCodecNil

`func (o *LoggingKafkaAllOf) SetCompressionCodecNil(b bool)`

 SetCompressionCodecNil sets the value for CompressionCodec to be an explicit nil

### UnsetCompressionCodec
`func (o *LoggingKafkaAllOf) UnsetCompressionCodec()`

UnsetCompressionCodec ensures that no value is present for CompressionCodec, not even an explicit nil
### GetRequiredAcks

`func (o *LoggingKafkaAllOf) GetRequiredAcks() int32`

GetRequiredAcks returns the RequiredAcks field if non-nil, zero value otherwise.

### GetRequiredAcksOk

`func (o *LoggingKafkaAllOf) GetRequiredAcksOk() (*int32, bool)`

GetRequiredAcksOk returns a tuple with the RequiredAcks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAcks

`func (o *LoggingKafkaAllOf) SetRequiredAcks(v int32)`

SetRequiredAcks sets RequiredAcks field to given value.

### HasRequiredAcks

`func (o *LoggingKafkaAllOf) HasRequiredAcks() bool`

HasRequiredAcks returns a boolean if a field has been set.

### GetRequestMaxBytes

`func (o *LoggingKafkaAllOf) GetRequestMaxBytes() int32`

GetRequestMaxBytes returns the RequestMaxBytes field if non-nil, zero value otherwise.

### GetRequestMaxBytesOk

`func (o *LoggingKafkaAllOf) GetRequestMaxBytesOk() (*int32, bool)`

GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestMaxBytes

`func (o *LoggingKafkaAllOf) SetRequestMaxBytes(v int32)`

SetRequestMaxBytes sets RequestMaxBytes field to given value.

### HasRequestMaxBytes

`func (o *LoggingKafkaAllOf) HasRequestMaxBytes() bool`

HasRequestMaxBytes returns a boolean if a field has been set.

### GetParseLogKeyvals

`func (o *LoggingKafkaAllOf) GetParseLogKeyvals() bool`

GetParseLogKeyvals returns the ParseLogKeyvals field if non-nil, zero value otherwise.

### GetParseLogKeyvalsOk

`func (o *LoggingKafkaAllOf) GetParseLogKeyvalsOk() (*bool, bool)`

GetParseLogKeyvalsOk returns a tuple with the ParseLogKeyvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseLogKeyvals

`func (o *LoggingKafkaAllOf) SetParseLogKeyvals(v bool)`

SetParseLogKeyvals sets ParseLogKeyvals field to given value.

### HasParseLogKeyvals

`func (o *LoggingKafkaAllOf) HasParseLogKeyvals() bool`

HasParseLogKeyvals returns a boolean if a field has been set.

### GetAuthMethod

`func (o *LoggingKafkaAllOf) GetAuthMethod() string`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *LoggingKafkaAllOf) GetAuthMethodOk() (*string, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *LoggingKafkaAllOf) SetAuthMethod(v string)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *LoggingKafkaAllOf) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetUser

`func (o *LoggingKafkaAllOf) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LoggingKafkaAllOf) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LoggingKafkaAllOf) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *LoggingKafkaAllOf) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *LoggingKafkaAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *LoggingKafkaAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *LoggingKafkaAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *LoggingKafkaAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseTLS

`func (o *LoggingKafkaAllOf) GetUseTLS() LoggingUseTLS`

GetUseTLS returns the UseTLS field if non-nil, zero value otherwise.

### GetUseTLSOk

`func (o *LoggingKafkaAllOf) GetUseTLSOk() (*LoggingUseTLS, bool)`

GetUseTLSOk returns a tuple with the UseTLS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTLS

`func (o *LoggingKafkaAllOf) SetUseTLS(v LoggingUseTLS)`

SetUseTLS sets UseTLS field to given value.

### HasUseTLS

`func (o *LoggingKafkaAllOf) HasUseTLS() bool`

HasUseTLS returns a boolean if a field has been set.


[Back to API list](../README.md#documentation-for-api-endpoints) | [Back to README](../README.md)
