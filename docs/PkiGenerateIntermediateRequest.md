# PkiGenerateIntermediateRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddBasicConstraints** | Pointer to **bool** | Whether to add a Basic Constraints extension with CA: true. Only needed as a workaround in some compatibility scenarios with Active Directory Certificate Services. | [optional] 
**AltNames** | Pointer to **string** | The requested Subject Alternative Names, if any, in a comma-delimited list. May contain both DNS names and email addresses. | [optional] 
**CommonName** | Pointer to **string** | The requested common name; if you want more than one, specify the alternative names in the alt_names map. If not specified when signing, the common name will be taken from the CSR; other names must still be specified in alt_names or ip_sans. | [optional] 
**Country** | Pointer to **[]string** | If set, Country will be set to this value. | [optional] 
**ExcludeCnFromSans** | Pointer to **bool** | If true, the Common Name will not be included in DNS or Email Subject Alternate Names. Defaults to false (CN is included). | [optional] [default to false]
**Format** | Pointer to **string** | Format for returned data. Can be \&quot;pem\&quot;, \&quot;der\&quot;, or \&quot;pem_bundle\&quot;. If \&quot;pem_bundle\&quot;, any private key and issuing cert will be appended to the certificate pem. If \&quot;der\&quot;, the value will be base64 encoded. Defaults to \&quot;pem\&quot;. | [optional] [default to "pem"]
**IpSans** | Pointer to **[]string** | The requested IP SANs, if any, in a comma-delimited list | [optional] 
**KeyBits** | Pointer to **int32** | The number of bits to use. Allowed values are 0 (universal default); with rsa key_type: 2048 (default), 3072, or 4096; with ec key_type: 224, 256 (default), 384, or 521; ignored with ed25519. | [optional] [default to 0]
**KeyName** | Pointer to **string** | Provide a name to the generated or existing key, the name must be unique across all keys and not be the reserved value &#x27;default&#x27; | [optional] 
**KeyRef** | Pointer to **string** | Reference to a existing key; either \&quot;default\&quot; for the configured default key, an identifier or the name assigned to the key. | [optional] [default to "default"]
**KeyType** | Pointer to **string** | The type of key to use; defaults to RSA. \&quot;rsa\&quot; \&quot;ec\&quot; and \&quot;ed25519\&quot; are the only valid values. | [optional] [default to "rsa"]
**Locality** | Pointer to **[]string** | If set, Locality will be set to this value. | [optional] 
**ManagedKeyId** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_name is required. Ignored for other types. | [optional] 
**ManagedKeyName** | Pointer to **string** | The name of the managed key to use when the exported type is kms. When kms type is the key type, this field or managed_key_id is required. Ignored for other types. | [optional] 
**NotAfter** | Pointer to **string** | Set the not after field of the certificate with specified date value. The value format should be given in UTC format YYYY-MM-ddTHH:MM:SSZ | [optional] 
**NotBeforeDuration** | Pointer to **int32** | The duration before now which the certificate needs to be backdated by. | [optional] [default to 30]
**Organization** | Pointer to **[]string** | If set, O (Organization) will be set to this value. | [optional] 
**OtherSans** | Pointer to **[]string** | Requested other SANs, in an array with the format &lt;oid&gt;;UTF8:&lt;utf8 string value&gt; for each entry. | [optional] 
**Ou** | Pointer to **[]string** | If set, OU (OrganizationalUnit) will be set to this value. | [optional] 
**PostalCode** | Pointer to **[]string** | If set, Postal Code will be set to this value. | [optional] 
**PrivateKeyFormat** | Pointer to **string** | Format for the returned private key. Generally the default will be controlled by the \&quot;format\&quot; parameter as either base64-encoded DER or PEM-encoded DER. However, this can be set to \&quot;pkcs8\&quot; to have the returned private key contain base64-encoded pkcs8 or PEM-encoded pkcs8 instead. Defaults to \&quot;der\&quot;. | [optional] [default to "der"]
**Province** | Pointer to **[]string** | If set, Province will be set to this value. | [optional] 
**SerialNumber** | Pointer to **string** | The Subject&#x27;s requested serial number, if any. See RFC 4519 Section 2.31 &#x27;serialNumber&#x27; for a description of this field. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5. This has no impact on the final certificate&#x27;s Serial Number field. | [optional] 
**SignatureBits** | Pointer to **int32** | The number of bits to use in the signature algorithm; accepts 256 for SHA-2-256, 384 for SHA-2-384, and 512 for SHA-2-512. Defaults to 0 to automatically detect based on key length (SHA-2-256 for RSA keys, and matching the curve size for NIST P-Curves). | [optional] [default to 0]
**StreetAddress** | Pointer to **[]string** | If set, Street Address will be set to this value. | [optional] 
**Ttl** | Pointer to **int32** | The requested Time To Live for the certificate; sets the expiration date. If not specified the role default, backend default, or system default TTL is used, in that order. Cannot be larger than the mount max TTL. Note: this only has an effect when generating a CA cert or signing a CA cert, not when generating a CSR for an intermediate CA. | [optional] 
**UriSans** | Pointer to **[]string** | The requested URI SANs, if any, in a comma-delimited list. | [optional] 



## Methods


### NewPkiGenerateIntermediateRequest

`func NewPkiGenerateIntermediateRequest() *PkiGenerateIntermediateRequest`

NewPkiGenerateIntermediateRequest instantiates a new PkiGenerateIntermediateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPkiGenerateIntermediateRequestWithDefaults

`func NewPkiGenerateIntermediateRequestWithDefaults() *PkiGenerateIntermediateRequest`

NewPkiGenerateIntermediateRequestWithDefaults instantiates a new PkiGenerateIntermediateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set


### GetAddBasicConstraints

`func (o *PkiGenerateIntermediateRequest) GetAddBasicConstraints() bool`

GetAddBasicConstraints returns the AddBasicConstraints field if non-nil, zero value otherwise.

### GetAddBasicConstraintsOk

`func (o *PkiGenerateIntermediateRequest) GetAddBasicConstraintsOk() (*bool, bool)`

GetAddBasicConstraintsOk returns a tuple with the AddBasicConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddBasicConstraints

`func (o *PkiGenerateIntermediateRequest) SetAddBasicConstraints(v bool)`

SetAddBasicConstraints sets AddBasicConstraints field to given value.


### HasAddBasicConstraints

`func (o *PkiGenerateIntermediateRequest) HasAddBasicConstraints() bool`

HasAddBasicConstraints returns a boolean if a field has been set.




### GetAltNames

`func (o *PkiGenerateIntermediateRequest) GetAltNames() string`

GetAltNames returns the AltNames field if non-nil, zero value otherwise.

### GetAltNamesOk

`func (o *PkiGenerateIntermediateRequest) GetAltNamesOk() (*string, bool)`

GetAltNamesOk returns a tuple with the AltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltNames

`func (o *PkiGenerateIntermediateRequest) SetAltNames(v string)`

SetAltNames sets AltNames field to given value.


### HasAltNames

`func (o *PkiGenerateIntermediateRequest) HasAltNames() bool`

HasAltNames returns a boolean if a field has been set.




### GetCommonName

`func (o *PkiGenerateIntermediateRequest) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *PkiGenerateIntermediateRequest) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *PkiGenerateIntermediateRequest) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.


### HasCommonName

`func (o *PkiGenerateIntermediateRequest) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.




### GetCountry

`func (o *PkiGenerateIntermediateRequest) GetCountry() []string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PkiGenerateIntermediateRequest) GetCountryOk() (*[]string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PkiGenerateIntermediateRequest) SetCountry(v []string)`

SetCountry sets Country field to given value.


### HasCountry

`func (o *PkiGenerateIntermediateRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.




### GetExcludeCnFromSans

`func (o *PkiGenerateIntermediateRequest) GetExcludeCnFromSans() bool`

GetExcludeCnFromSans returns the ExcludeCnFromSans field if non-nil, zero value otherwise.

### GetExcludeCnFromSansOk

`func (o *PkiGenerateIntermediateRequest) GetExcludeCnFromSansOk() (*bool, bool)`

GetExcludeCnFromSansOk returns a tuple with the ExcludeCnFromSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeCnFromSans

`func (o *PkiGenerateIntermediateRequest) SetExcludeCnFromSans(v bool)`

SetExcludeCnFromSans sets ExcludeCnFromSans field to given value.


### HasExcludeCnFromSans

`func (o *PkiGenerateIntermediateRequest) HasExcludeCnFromSans() bool`

HasExcludeCnFromSans returns a boolean if a field has been set.




### GetFormat

`func (o *PkiGenerateIntermediateRequest) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *PkiGenerateIntermediateRequest) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *PkiGenerateIntermediateRequest) SetFormat(v string)`

SetFormat sets Format field to given value.


### HasFormat

`func (o *PkiGenerateIntermediateRequest) HasFormat() bool`

HasFormat returns a boolean if a field has been set.




### GetIpSans

`func (o *PkiGenerateIntermediateRequest) GetIpSans() []string`

GetIpSans returns the IpSans field if non-nil, zero value otherwise.

### GetIpSansOk

`func (o *PkiGenerateIntermediateRequest) GetIpSansOk() (*[]string, bool)`

GetIpSansOk returns a tuple with the IpSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSans

`func (o *PkiGenerateIntermediateRequest) SetIpSans(v []string)`

SetIpSans sets IpSans field to given value.


### HasIpSans

`func (o *PkiGenerateIntermediateRequest) HasIpSans() bool`

HasIpSans returns a boolean if a field has been set.




### GetKeyBits

`func (o *PkiGenerateIntermediateRequest) GetKeyBits() int32`

GetKeyBits returns the KeyBits field if non-nil, zero value otherwise.

### GetKeyBitsOk

`func (o *PkiGenerateIntermediateRequest) GetKeyBitsOk() (*int32, bool)`

GetKeyBitsOk returns a tuple with the KeyBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyBits

`func (o *PkiGenerateIntermediateRequest) SetKeyBits(v int32)`

SetKeyBits sets KeyBits field to given value.


### HasKeyBits

`func (o *PkiGenerateIntermediateRequest) HasKeyBits() bool`

HasKeyBits returns a boolean if a field has been set.




### GetKeyName

`func (o *PkiGenerateIntermediateRequest) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PkiGenerateIntermediateRequest) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PkiGenerateIntermediateRequest) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### HasKeyName

`func (o *PkiGenerateIntermediateRequest) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.




### GetKeyRef

`func (o *PkiGenerateIntermediateRequest) GetKeyRef() string`

GetKeyRef returns the KeyRef field if non-nil, zero value otherwise.

### GetKeyRefOk

`func (o *PkiGenerateIntermediateRequest) GetKeyRefOk() (*string, bool)`

GetKeyRefOk returns a tuple with the KeyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyRef

`func (o *PkiGenerateIntermediateRequest) SetKeyRef(v string)`

SetKeyRef sets KeyRef field to given value.


### HasKeyRef

`func (o *PkiGenerateIntermediateRequest) HasKeyRef() bool`

HasKeyRef returns a boolean if a field has been set.




### GetKeyType

`func (o *PkiGenerateIntermediateRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *PkiGenerateIntermediateRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *PkiGenerateIntermediateRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### HasKeyType

`func (o *PkiGenerateIntermediateRequest) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.




### GetLocality

`func (o *PkiGenerateIntermediateRequest) GetLocality() []string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *PkiGenerateIntermediateRequest) GetLocalityOk() (*[]string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *PkiGenerateIntermediateRequest) SetLocality(v []string)`

SetLocality sets Locality field to given value.


### HasLocality

`func (o *PkiGenerateIntermediateRequest) HasLocality() bool`

HasLocality returns a boolean if a field has been set.




### GetManagedKeyId

`func (o *PkiGenerateIntermediateRequest) GetManagedKeyId() string`

GetManagedKeyId returns the ManagedKeyId field if non-nil, zero value otherwise.

### GetManagedKeyIdOk

`func (o *PkiGenerateIntermediateRequest) GetManagedKeyIdOk() (*string, bool)`

GetManagedKeyIdOk returns a tuple with the ManagedKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyId

`func (o *PkiGenerateIntermediateRequest) SetManagedKeyId(v string)`

SetManagedKeyId sets ManagedKeyId field to given value.


### HasManagedKeyId

`func (o *PkiGenerateIntermediateRequest) HasManagedKeyId() bool`

HasManagedKeyId returns a boolean if a field has been set.




### GetManagedKeyName

`func (o *PkiGenerateIntermediateRequest) GetManagedKeyName() string`

GetManagedKeyName returns the ManagedKeyName field if non-nil, zero value otherwise.

### GetManagedKeyNameOk

`func (o *PkiGenerateIntermediateRequest) GetManagedKeyNameOk() (*string, bool)`

GetManagedKeyNameOk returns a tuple with the ManagedKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedKeyName

`func (o *PkiGenerateIntermediateRequest) SetManagedKeyName(v string)`

SetManagedKeyName sets ManagedKeyName field to given value.


### HasManagedKeyName

`func (o *PkiGenerateIntermediateRequest) HasManagedKeyName() bool`

HasManagedKeyName returns a boolean if a field has been set.




### GetNotAfter

`func (o *PkiGenerateIntermediateRequest) GetNotAfter() string`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *PkiGenerateIntermediateRequest) GetNotAfterOk() (*string, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *PkiGenerateIntermediateRequest) SetNotAfter(v string)`

SetNotAfter sets NotAfter field to given value.


### HasNotAfter

`func (o *PkiGenerateIntermediateRequest) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.




### GetNotBeforeDuration

`func (o *PkiGenerateIntermediateRequest) GetNotBeforeDuration() int32`

GetNotBeforeDuration returns the NotBeforeDuration field if non-nil, zero value otherwise.

### GetNotBeforeDurationOk

`func (o *PkiGenerateIntermediateRequest) GetNotBeforeDurationOk() (*int32, bool)`

GetNotBeforeDurationOk returns a tuple with the NotBeforeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBeforeDuration

`func (o *PkiGenerateIntermediateRequest) SetNotBeforeDuration(v int32)`

SetNotBeforeDuration sets NotBeforeDuration field to given value.


### HasNotBeforeDuration

`func (o *PkiGenerateIntermediateRequest) HasNotBeforeDuration() bool`

HasNotBeforeDuration returns a boolean if a field has been set.




### GetOrganization

`func (o *PkiGenerateIntermediateRequest) GetOrganization() []string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *PkiGenerateIntermediateRequest) GetOrganizationOk() (*[]string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *PkiGenerateIntermediateRequest) SetOrganization(v []string)`

SetOrganization sets Organization field to given value.


### HasOrganization

`func (o *PkiGenerateIntermediateRequest) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.




### GetOtherSans

`func (o *PkiGenerateIntermediateRequest) GetOtherSans() []string`

GetOtherSans returns the OtherSans field if non-nil, zero value otherwise.

### GetOtherSansOk

`func (o *PkiGenerateIntermediateRequest) GetOtherSansOk() (*[]string, bool)`

GetOtherSansOk returns a tuple with the OtherSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherSans

`func (o *PkiGenerateIntermediateRequest) SetOtherSans(v []string)`

SetOtherSans sets OtherSans field to given value.


### HasOtherSans

`func (o *PkiGenerateIntermediateRequest) HasOtherSans() bool`

HasOtherSans returns a boolean if a field has been set.




### GetOu

`func (o *PkiGenerateIntermediateRequest) GetOu() []string`

GetOu returns the Ou field if non-nil, zero value otherwise.

### GetOuOk

`func (o *PkiGenerateIntermediateRequest) GetOuOk() (*[]string, bool)`

GetOuOk returns a tuple with the Ou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOu

`func (o *PkiGenerateIntermediateRequest) SetOu(v []string)`

SetOu sets Ou field to given value.


### HasOu

`func (o *PkiGenerateIntermediateRequest) HasOu() bool`

HasOu returns a boolean if a field has been set.




### GetPostalCode

`func (o *PkiGenerateIntermediateRequest) GetPostalCode() []string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PkiGenerateIntermediateRequest) GetPostalCodeOk() (*[]string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PkiGenerateIntermediateRequest) SetPostalCode(v []string)`

SetPostalCode sets PostalCode field to given value.


### HasPostalCode

`func (o *PkiGenerateIntermediateRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.




### GetPrivateKeyFormat

`func (o *PkiGenerateIntermediateRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *PkiGenerateIntermediateRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *PkiGenerateIntermediateRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.


### HasPrivateKeyFormat

`func (o *PkiGenerateIntermediateRequest) HasPrivateKeyFormat() bool`

HasPrivateKeyFormat returns a boolean if a field has been set.




### GetProvince

`func (o *PkiGenerateIntermediateRequest) GetProvince() []string`

GetProvince returns the Province field if non-nil, zero value otherwise.

### GetProvinceOk

`func (o *PkiGenerateIntermediateRequest) GetProvinceOk() (*[]string, bool)`

GetProvinceOk returns a tuple with the Province field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvince

`func (o *PkiGenerateIntermediateRequest) SetProvince(v []string)`

SetProvince sets Province field to given value.


### HasProvince

`func (o *PkiGenerateIntermediateRequest) HasProvince() bool`

HasProvince returns a boolean if a field has been set.




### GetSerialNumber

`func (o *PkiGenerateIntermediateRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PkiGenerateIntermediateRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PkiGenerateIntermediateRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### HasSerialNumber

`func (o *PkiGenerateIntermediateRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.




### GetSignatureBits

`func (o *PkiGenerateIntermediateRequest) GetSignatureBits() int32`

GetSignatureBits returns the SignatureBits field if non-nil, zero value otherwise.

### GetSignatureBitsOk

`func (o *PkiGenerateIntermediateRequest) GetSignatureBitsOk() (*int32, bool)`

GetSignatureBitsOk returns a tuple with the SignatureBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureBits

`func (o *PkiGenerateIntermediateRequest) SetSignatureBits(v int32)`

SetSignatureBits sets SignatureBits field to given value.


### HasSignatureBits

`func (o *PkiGenerateIntermediateRequest) HasSignatureBits() bool`

HasSignatureBits returns a boolean if a field has been set.




### GetStreetAddress

`func (o *PkiGenerateIntermediateRequest) GetStreetAddress() []string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *PkiGenerateIntermediateRequest) GetStreetAddressOk() (*[]string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *PkiGenerateIntermediateRequest) SetStreetAddress(v []string)`

SetStreetAddress sets StreetAddress field to given value.


### HasStreetAddress

`func (o *PkiGenerateIntermediateRequest) HasStreetAddress() bool`

HasStreetAddress returns a boolean if a field has been set.




### GetTtl

`func (o *PkiGenerateIntermediateRequest) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PkiGenerateIntermediateRequest) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PkiGenerateIntermediateRequest) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### HasTtl

`func (o *PkiGenerateIntermediateRequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.




### GetUriSans

`func (o *PkiGenerateIntermediateRequest) GetUriSans() []string`

GetUriSans returns the UriSans field if non-nil, zero value otherwise.

### GetUriSansOk

`func (o *PkiGenerateIntermediateRequest) GetUriSansOk() (*[]string, bool)`

GetUriSansOk returns a tuple with the UriSans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriSans

`func (o *PkiGenerateIntermediateRequest) SetUriSans(v []string)`

SetUriSans sets UriSans field to given value.


### HasUriSans

`func (o *PkiGenerateIntermediateRequest) HasUriSans() bool`

HasUriSans returns a boolean if a field has been set.









[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

