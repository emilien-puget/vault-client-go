# SystemMountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | Configuration for this mount, such as default_lease_ttl and max_lease_ttl. | [optional] 
**Description** | Pointer to **string** | User-friendly description for this mount. | [optional] 
**ExternalEntropyAccess** | Pointer to **bool** | Whether to give the mount access to Vault&#39;s external entropy. | [optional] [default to false]
**Local** | Pointer to **bool** | Mark the mount as a local mount, which is not replicated and is unaffected by replication. | [optional] [default to false]
**Options** | Pointer to **map[string]interface{}** | The options to pass into the backend. Should be a json object with string keys and values. | [optional] 
**PluginName** | Pointer to **string** | Name of the plugin to mount based from the name registered in the plugin catalog. | [optional] 
**SealWrap** | Pointer to **bool** | Whether to turn on seal wrapping for the mount. | [optional] [default to false]
**Type** | Pointer to **string** | The type of the backend. Example: \&quot;passthrough\&quot; | [optional] 

## Methods

### NewSystemMountsRequest

`func NewSystemMountsRequest() *SystemMountsRequest`

NewSystemMountsRequest instantiates a new SystemMountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemMountsRequestWithDefaults

`func NewSystemMountsRequestWithDefaults() *SystemMountsRequest`

NewSystemMountsRequestWithDefaults instantiates a new SystemMountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SystemMountsRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SystemMountsRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SystemMountsRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SystemMountsRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *SystemMountsRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemMountsRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemMountsRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemMountsRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalEntropyAccess

`func (o *SystemMountsRequest) GetExternalEntropyAccess() bool`

GetExternalEntropyAccess returns the ExternalEntropyAccess field if non-nil, zero value otherwise.

### GetExternalEntropyAccessOk

`func (o *SystemMountsRequest) GetExternalEntropyAccessOk() (*bool, bool)`

GetExternalEntropyAccessOk returns a tuple with the ExternalEntropyAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalEntropyAccess

`func (o *SystemMountsRequest) SetExternalEntropyAccess(v bool)`

SetExternalEntropyAccess sets ExternalEntropyAccess field to given value.

### HasExternalEntropyAccess

`func (o *SystemMountsRequest) HasExternalEntropyAccess() bool`

HasExternalEntropyAccess returns a boolean if a field has been set.

### GetLocal

`func (o *SystemMountsRequest) GetLocal() bool`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *SystemMountsRequest) GetLocalOk() (*bool, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *SystemMountsRequest) SetLocal(v bool)`

SetLocal sets Local field to given value.

### HasLocal

`func (o *SystemMountsRequest) HasLocal() bool`

HasLocal returns a boolean if a field has been set.

### GetOptions

`func (o *SystemMountsRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SystemMountsRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SystemMountsRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SystemMountsRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPluginName

`func (o *SystemMountsRequest) GetPluginName() string`

GetPluginName returns the PluginName field if non-nil, zero value otherwise.

### GetPluginNameOk

`func (o *SystemMountsRequest) GetPluginNameOk() (*string, bool)`

GetPluginNameOk returns a tuple with the PluginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginName

`func (o *SystemMountsRequest) SetPluginName(v string)`

SetPluginName sets PluginName field to given value.

### HasPluginName

`func (o *SystemMountsRequest) HasPluginName() bool`

HasPluginName returns a boolean if a field has been set.

### GetSealWrap

`func (o *SystemMountsRequest) GetSealWrap() bool`

GetSealWrap returns the SealWrap field if non-nil, zero value otherwise.

### GetSealWrapOk

`func (o *SystemMountsRequest) GetSealWrapOk() (*bool, bool)`

GetSealWrapOk returns a tuple with the SealWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealWrap

`func (o *SystemMountsRequest) SetSealWrap(v bool)`

SetSealWrap sets SealWrap field to given value.

### HasSealWrap

`func (o *SystemMountsRequest) HasSealWrap() bool`

HasSealWrap returns a boolean if a field has been set.

### GetType

`func (o *SystemMountsRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemMountsRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemMountsRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SystemMountsRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

