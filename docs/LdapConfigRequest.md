# LdapConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousGroupSearch** | Pointer to **bool** | Use anonymous binds when performing LDAP group searches (if true the initial credentials will still be used for the initial connection test). | [optional] [default to false]
**Binddn** | Pointer to **string** | LDAP DN for searching for the user DN (optional) | [optional] 
**Bindpass** | Pointer to **string** | LDAP password for searching for the user DN (optional) | [optional] 
**CaseSensitiveNames** | Pointer to **bool** | If true, case sensitivity will be used when comparing usernames and groups for matching policies. | [optional] 
**Certificate** | Pointer to **string** | CA certificate to use when verifying LDAP server certificate, must be x509 PEM encoded (optional) | [optional] 
**ClientTlsCert** | Pointer to **string** | Client certificate to provide to the LDAP server, must be x509 PEM encoded (optional) | [optional] 
**ClientTlsKey** | Pointer to **string** | Client certificate key to provide to the LDAP server, must be x509 PEM encoded (optional) | [optional] 
**DenyNullBind** | Pointer to **bool** | Denies an unauthenticated LDAP bind request if the user&#39;s password is empty; defaults to true | [optional] [default to true]
**Discoverdn** | Pointer to **bool** | Use anonymous bind to discover the bind DN of a user (optional) | [optional] 
**Groupattr** | Pointer to **string** | LDAP attribute to follow on objects returned by &lt;groupfilter&gt; in order to enumerate user group membership. Examples: \&quot;cn\&quot; or \&quot;memberOf\&quot;, etc. Default: cn | [optional] [default to "cn"]
**Groupdn** | Pointer to **string** | LDAP search base to use for group membership search (eg: ou&#x3D;Groups,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Groupfilter** | Pointer to **string** | Go template for querying group membership of user (optional) The template can access the following context variables: UserDN, Username Example: (&amp;(objectClass&#x3D;group)(member:1.2.840.113556.1.4.1941:&#x3D;{{.UserDN}})) Default: (|(memberUid&#x3D;{{.Username}})(member&#x3D;{{.UserDN}})(uniqueMember&#x3D;{{.UserDN}})) | [optional] [default to "(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))"]
**InsecureTls** | Pointer to **bool** | Skip LDAP server SSL Certificate verification - VERY insecure (optional) | [optional] 
**RequestTimeout** | Pointer to **int32** | Timeout, in seconds, for the connection when making requests against the server before returning back an error. | [optional] 
**Starttls** | Pointer to **bool** | Issue a StartTLS command after establishing unencrypted connection (optional) | [optional] 
**TlsMaxVersion** | Pointer to **string** | Maximum TLS version to use. Accepted values are &#39;tls10&#39;, &#39;tls11&#39;, &#39;tls12&#39; or &#39;tls13&#39;. Defaults to &#39;tls12&#39; | [optional] [default to "tls12"]
**TlsMinVersion** | Pointer to **string** | Minimum TLS version to use. Accepted values are &#39;tls10&#39;, &#39;tls11&#39;, &#39;tls12&#39; or &#39;tls13&#39;. Defaults to &#39;tls12&#39; | [optional] [default to "tls12"]
**TokenBoundCidrs** | Pointer to **[]string** | Comma separated string or JSON list of CIDR blocks. If set, specifies the blocks of IP addresses which are allowed to use the generated token. | [optional] 
**TokenExplicitMaxTtl** | Pointer to **int32** | If set, tokens created via this role carry an explicit maximum TTL. During renewal, the current maximum TTL values of the role and the mount are not checked for changes, and any updates to these values will have no effect on the token being renewed. | [optional] 
**TokenMaxTtl** | Pointer to **int32** | The maximum lifetime of the generated token | [optional] 
**TokenNoDefaultPolicy** | Pointer to **bool** | If true, the &#39;default&#39; policy will not automatically be added to generated tokens | [optional] 
**TokenNumUses** | Pointer to **int32** | The maximum number of times a token may be used, a value of zero means unlimited | [optional] 
**TokenPeriod** | Pointer to **int32** | If set, tokens created via this role will have no max lifetime; instead, their renewal period will be fixed to this value. This takes an integer number of seconds, or a string duration (e.g. \&quot;24h\&quot;). | [optional] 
**TokenPolicies** | Pointer to **[]string** | Comma-separated list of policies. This will apply to all tokens generated by this auth method, in addition to any configured for specific users/groups. | [optional] 
**TokenTtl** | Pointer to **int32** | The initial ttl of the token to generate | [optional] 
**TokenType** | Pointer to **string** | The type of token to generate, service or batch | [optional] [default to "default-service"]
**Upndomain** | Pointer to **string** | Enables userPrincipalDomain login with [username]@UPNDomain (optional) | [optional] 
**Url** | Pointer to **string** | LDAP URL to connect to (default: ldap://127.0.0.1). Multiple URLs can be specified by concatenating them with commas; they will be tried in-order. | [optional] [default to "ldap://127.0.0.1"]
**UsePre111GroupCnBehavior** | Pointer to **bool** | In Vault 1.1.1 a fix for handling group CN values of different cases unfortunately introduced a regression that could cause previously defined groups to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for matching group CNs will be used. This is only needed in some upgrade scenarios for backwards compatibility. It is enabled by default if the config is upgraded but disabled by default on new configurations. | [optional] 
**UseTokenGroups** | Pointer to **bool** | If true, use the Active Directory tokenGroups constructed attribute of the user to find the group memberships. This will find all security groups including nested ones. | [optional] [default to false]
**Userattr** | Pointer to **string** | Attribute used for users (default: cn) | [optional] [default to "cn"]
**Userdn** | Pointer to **string** | LDAP domain to use for users (eg: ou&#x3D;People,dc&#x3D;example,dc&#x3D;org) | [optional] 
**Userfilter** | Pointer to **string** | Go template for LDAP user search filer (optional) The template can access the following context variables: UserAttr, Username Default: ({{.UserAttr}}&#x3D;{{.Username}}) | [optional] [default to "({{.UserAttr}}={{.Username}})"]
**UsernameAsAlias** | Pointer to **bool** | If true, sets the alias name to the username | [optional] [default to false]

## Methods

### NewLdapConfigRequest

`func NewLdapConfigRequest() *LdapConfigRequest`

NewLdapConfigRequest instantiates a new LdapConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigRequestWithDefaults

`func NewLdapConfigRequestWithDefaults() *LdapConfigRequest`

NewLdapConfigRequestWithDefaults instantiates a new LdapConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnonymousGroupSearch

`func (o *LdapConfigRequest) GetAnonymousGroupSearch() bool`

GetAnonymousGroupSearch returns the AnonymousGroupSearch field if non-nil, zero value otherwise.

### GetAnonymousGroupSearchOk

`func (o *LdapConfigRequest) GetAnonymousGroupSearchOk() (*bool, bool)`

GetAnonymousGroupSearchOk returns a tuple with the AnonymousGroupSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnonymousGroupSearch

`func (o *LdapConfigRequest) SetAnonymousGroupSearch(v bool)`

SetAnonymousGroupSearch sets AnonymousGroupSearch field to given value.

### HasAnonymousGroupSearch

`func (o *LdapConfigRequest) HasAnonymousGroupSearch() bool`

HasAnonymousGroupSearch returns a boolean if a field has been set.

### GetBinddn

`func (o *LdapConfigRequest) GetBinddn() string`

GetBinddn returns the Binddn field if non-nil, zero value otherwise.

### GetBinddnOk

`func (o *LdapConfigRequest) GetBinddnOk() (*string, bool)`

GetBinddnOk returns a tuple with the Binddn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinddn

`func (o *LdapConfigRequest) SetBinddn(v string)`

SetBinddn sets Binddn field to given value.

### HasBinddn

`func (o *LdapConfigRequest) HasBinddn() bool`

HasBinddn returns a boolean if a field has been set.

### GetBindpass

`func (o *LdapConfigRequest) GetBindpass() string`

GetBindpass returns the Bindpass field if non-nil, zero value otherwise.

### GetBindpassOk

`func (o *LdapConfigRequest) GetBindpassOk() (*string, bool)`

GetBindpassOk returns a tuple with the Bindpass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindpass

`func (o *LdapConfigRequest) SetBindpass(v string)`

SetBindpass sets Bindpass field to given value.

### HasBindpass

`func (o *LdapConfigRequest) HasBindpass() bool`

HasBindpass returns a boolean if a field has been set.

### GetCaseSensitiveNames

`func (o *LdapConfigRequest) GetCaseSensitiveNames() bool`

GetCaseSensitiveNames returns the CaseSensitiveNames field if non-nil, zero value otherwise.

### GetCaseSensitiveNamesOk

`func (o *LdapConfigRequest) GetCaseSensitiveNamesOk() (*bool, bool)`

GetCaseSensitiveNamesOk returns a tuple with the CaseSensitiveNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitiveNames

`func (o *LdapConfigRequest) SetCaseSensitiveNames(v bool)`

SetCaseSensitiveNames sets CaseSensitiveNames field to given value.

### HasCaseSensitiveNames

`func (o *LdapConfigRequest) HasCaseSensitiveNames() bool`

HasCaseSensitiveNames returns a boolean if a field has been set.

### GetCertificate

`func (o *LdapConfigRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *LdapConfigRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *LdapConfigRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *LdapConfigRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetClientTlsCert

`func (o *LdapConfigRequest) GetClientTlsCert() string`

GetClientTlsCert returns the ClientTlsCert field if non-nil, zero value otherwise.

### GetClientTlsCertOk

`func (o *LdapConfigRequest) GetClientTlsCertOk() (*string, bool)`

GetClientTlsCertOk returns a tuple with the ClientTlsCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsCert

`func (o *LdapConfigRequest) SetClientTlsCert(v string)`

SetClientTlsCert sets ClientTlsCert field to given value.

### HasClientTlsCert

`func (o *LdapConfigRequest) HasClientTlsCert() bool`

HasClientTlsCert returns a boolean if a field has been set.

### GetClientTlsKey

`func (o *LdapConfigRequest) GetClientTlsKey() string`

GetClientTlsKey returns the ClientTlsKey field if non-nil, zero value otherwise.

### GetClientTlsKeyOk

`func (o *LdapConfigRequest) GetClientTlsKeyOk() (*string, bool)`

GetClientTlsKeyOk returns a tuple with the ClientTlsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTlsKey

`func (o *LdapConfigRequest) SetClientTlsKey(v string)`

SetClientTlsKey sets ClientTlsKey field to given value.

### HasClientTlsKey

`func (o *LdapConfigRequest) HasClientTlsKey() bool`

HasClientTlsKey returns a boolean if a field has been set.

### GetDenyNullBind

`func (o *LdapConfigRequest) GetDenyNullBind() bool`

GetDenyNullBind returns the DenyNullBind field if non-nil, zero value otherwise.

### GetDenyNullBindOk

`func (o *LdapConfigRequest) GetDenyNullBindOk() (*bool, bool)`

GetDenyNullBindOk returns a tuple with the DenyNullBind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyNullBind

`func (o *LdapConfigRequest) SetDenyNullBind(v bool)`

SetDenyNullBind sets DenyNullBind field to given value.

### HasDenyNullBind

`func (o *LdapConfigRequest) HasDenyNullBind() bool`

HasDenyNullBind returns a boolean if a field has been set.

### GetDiscoverdn

`func (o *LdapConfigRequest) GetDiscoverdn() bool`

GetDiscoverdn returns the Discoverdn field if non-nil, zero value otherwise.

### GetDiscoverdnOk

`func (o *LdapConfigRequest) GetDiscoverdnOk() (*bool, bool)`

GetDiscoverdnOk returns a tuple with the Discoverdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverdn

`func (o *LdapConfigRequest) SetDiscoverdn(v bool)`

SetDiscoverdn sets Discoverdn field to given value.

### HasDiscoverdn

`func (o *LdapConfigRequest) HasDiscoverdn() bool`

HasDiscoverdn returns a boolean if a field has been set.

### GetGroupattr

`func (o *LdapConfigRequest) GetGroupattr() string`

GetGroupattr returns the Groupattr field if non-nil, zero value otherwise.

### GetGroupattrOk

`func (o *LdapConfigRequest) GetGroupattrOk() (*string, bool)`

GetGroupattrOk returns a tuple with the Groupattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupattr

`func (o *LdapConfigRequest) SetGroupattr(v string)`

SetGroupattr sets Groupattr field to given value.

### HasGroupattr

`func (o *LdapConfigRequest) HasGroupattr() bool`

HasGroupattr returns a boolean if a field has been set.

### GetGroupdn

`func (o *LdapConfigRequest) GetGroupdn() string`

GetGroupdn returns the Groupdn field if non-nil, zero value otherwise.

### GetGroupdnOk

`func (o *LdapConfigRequest) GetGroupdnOk() (*string, bool)`

GetGroupdnOk returns a tuple with the Groupdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupdn

`func (o *LdapConfigRequest) SetGroupdn(v string)`

SetGroupdn sets Groupdn field to given value.

### HasGroupdn

`func (o *LdapConfigRequest) HasGroupdn() bool`

HasGroupdn returns a boolean if a field has been set.

### GetGroupfilter

`func (o *LdapConfigRequest) GetGroupfilter() string`

GetGroupfilter returns the Groupfilter field if non-nil, zero value otherwise.

### GetGroupfilterOk

`func (o *LdapConfigRequest) GetGroupfilterOk() (*string, bool)`

GetGroupfilterOk returns a tuple with the Groupfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupfilter

`func (o *LdapConfigRequest) SetGroupfilter(v string)`

SetGroupfilter sets Groupfilter field to given value.

### HasGroupfilter

`func (o *LdapConfigRequest) HasGroupfilter() bool`

HasGroupfilter returns a boolean if a field has been set.

### GetInsecureTls

`func (o *LdapConfigRequest) GetInsecureTls() bool`

GetInsecureTls returns the InsecureTls field if non-nil, zero value otherwise.

### GetInsecureTlsOk

`func (o *LdapConfigRequest) GetInsecureTlsOk() (*bool, bool)`

GetInsecureTlsOk returns a tuple with the InsecureTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecureTls

`func (o *LdapConfigRequest) SetInsecureTls(v bool)`

SetInsecureTls sets InsecureTls field to given value.

### HasInsecureTls

`func (o *LdapConfigRequest) HasInsecureTls() bool`

HasInsecureTls returns a boolean if a field has been set.

### GetRequestTimeout

`func (o *LdapConfigRequest) GetRequestTimeout() int32`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *LdapConfigRequest) GetRequestTimeoutOk() (*int32, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *LdapConfigRequest) SetRequestTimeout(v int32)`

SetRequestTimeout sets RequestTimeout field to given value.

### HasRequestTimeout

`func (o *LdapConfigRequest) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.

### GetStarttls

`func (o *LdapConfigRequest) GetStarttls() bool`

GetStarttls returns the Starttls field if non-nil, zero value otherwise.

### GetStarttlsOk

`func (o *LdapConfigRequest) GetStarttlsOk() (*bool, bool)`

GetStarttlsOk returns a tuple with the Starttls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarttls

`func (o *LdapConfigRequest) SetStarttls(v bool)`

SetStarttls sets Starttls field to given value.

### HasStarttls

`func (o *LdapConfigRequest) HasStarttls() bool`

HasStarttls returns a boolean if a field has been set.

### GetTlsMaxVersion

`func (o *LdapConfigRequest) GetTlsMaxVersion() string`

GetTlsMaxVersion returns the TlsMaxVersion field if non-nil, zero value otherwise.

### GetTlsMaxVersionOk

`func (o *LdapConfigRequest) GetTlsMaxVersionOk() (*string, bool)`

GetTlsMaxVersionOk returns a tuple with the TlsMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMaxVersion

`func (o *LdapConfigRequest) SetTlsMaxVersion(v string)`

SetTlsMaxVersion sets TlsMaxVersion field to given value.

### HasTlsMaxVersion

`func (o *LdapConfigRequest) HasTlsMaxVersion() bool`

HasTlsMaxVersion returns a boolean if a field has been set.

### GetTlsMinVersion

`func (o *LdapConfigRequest) GetTlsMinVersion() string`

GetTlsMinVersion returns the TlsMinVersion field if non-nil, zero value otherwise.

### GetTlsMinVersionOk

`func (o *LdapConfigRequest) GetTlsMinVersionOk() (*string, bool)`

GetTlsMinVersionOk returns a tuple with the TlsMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsMinVersion

`func (o *LdapConfigRequest) SetTlsMinVersion(v string)`

SetTlsMinVersion sets TlsMinVersion field to given value.

### HasTlsMinVersion

`func (o *LdapConfigRequest) HasTlsMinVersion() bool`

HasTlsMinVersion returns a boolean if a field has been set.

### GetTokenBoundCidrs

`func (o *LdapConfigRequest) GetTokenBoundCidrs() []string`

GetTokenBoundCidrs returns the TokenBoundCidrs field if non-nil, zero value otherwise.

### GetTokenBoundCidrsOk

`func (o *LdapConfigRequest) GetTokenBoundCidrsOk() (*[]string, bool)`

GetTokenBoundCidrsOk returns a tuple with the TokenBoundCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBoundCidrs

`func (o *LdapConfigRequest) SetTokenBoundCidrs(v []string)`

SetTokenBoundCidrs sets TokenBoundCidrs field to given value.

### HasTokenBoundCidrs

`func (o *LdapConfigRequest) HasTokenBoundCidrs() bool`

HasTokenBoundCidrs returns a boolean if a field has been set.

### GetTokenExplicitMaxTtl

`func (o *LdapConfigRequest) GetTokenExplicitMaxTtl() int32`

GetTokenExplicitMaxTtl returns the TokenExplicitMaxTtl field if non-nil, zero value otherwise.

### GetTokenExplicitMaxTtlOk

`func (o *LdapConfigRequest) GetTokenExplicitMaxTtlOk() (*int32, bool)`

GetTokenExplicitMaxTtlOk returns a tuple with the TokenExplicitMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExplicitMaxTtl

`func (o *LdapConfigRequest) SetTokenExplicitMaxTtl(v int32)`

SetTokenExplicitMaxTtl sets TokenExplicitMaxTtl field to given value.

### HasTokenExplicitMaxTtl

`func (o *LdapConfigRequest) HasTokenExplicitMaxTtl() bool`

HasTokenExplicitMaxTtl returns a boolean if a field has been set.

### GetTokenMaxTtl

`func (o *LdapConfigRequest) GetTokenMaxTtl() int32`

GetTokenMaxTtl returns the TokenMaxTtl field if non-nil, zero value otherwise.

### GetTokenMaxTtlOk

`func (o *LdapConfigRequest) GetTokenMaxTtlOk() (*int32, bool)`

GetTokenMaxTtlOk returns a tuple with the TokenMaxTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMaxTtl

`func (o *LdapConfigRequest) SetTokenMaxTtl(v int32)`

SetTokenMaxTtl sets TokenMaxTtl field to given value.

### HasTokenMaxTtl

`func (o *LdapConfigRequest) HasTokenMaxTtl() bool`

HasTokenMaxTtl returns a boolean if a field has been set.

### GetTokenNoDefaultPolicy

`func (o *LdapConfigRequest) GetTokenNoDefaultPolicy() bool`

GetTokenNoDefaultPolicy returns the TokenNoDefaultPolicy field if non-nil, zero value otherwise.

### GetTokenNoDefaultPolicyOk

`func (o *LdapConfigRequest) GetTokenNoDefaultPolicyOk() (*bool, bool)`

GetTokenNoDefaultPolicyOk returns a tuple with the TokenNoDefaultPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNoDefaultPolicy

`func (o *LdapConfigRequest) SetTokenNoDefaultPolicy(v bool)`

SetTokenNoDefaultPolicy sets TokenNoDefaultPolicy field to given value.

### HasTokenNoDefaultPolicy

`func (o *LdapConfigRequest) HasTokenNoDefaultPolicy() bool`

HasTokenNoDefaultPolicy returns a boolean if a field has been set.

### GetTokenNumUses

`func (o *LdapConfigRequest) GetTokenNumUses() int32`

GetTokenNumUses returns the TokenNumUses field if non-nil, zero value otherwise.

### GetTokenNumUsesOk

`func (o *LdapConfigRequest) GetTokenNumUsesOk() (*int32, bool)`

GetTokenNumUsesOk returns a tuple with the TokenNumUses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenNumUses

`func (o *LdapConfigRequest) SetTokenNumUses(v int32)`

SetTokenNumUses sets TokenNumUses field to given value.

### HasTokenNumUses

`func (o *LdapConfigRequest) HasTokenNumUses() bool`

HasTokenNumUses returns a boolean if a field has been set.

### GetTokenPeriod

`func (o *LdapConfigRequest) GetTokenPeriod() int32`

GetTokenPeriod returns the TokenPeriod field if non-nil, zero value otherwise.

### GetTokenPeriodOk

`func (o *LdapConfigRequest) GetTokenPeriodOk() (*int32, bool)`

GetTokenPeriodOk returns a tuple with the TokenPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPeriod

`func (o *LdapConfigRequest) SetTokenPeriod(v int32)`

SetTokenPeriod sets TokenPeriod field to given value.

### HasTokenPeriod

`func (o *LdapConfigRequest) HasTokenPeriod() bool`

HasTokenPeriod returns a boolean if a field has been set.

### GetTokenPolicies

`func (o *LdapConfigRequest) GetTokenPolicies() []string`

GetTokenPolicies returns the TokenPolicies field if non-nil, zero value otherwise.

### GetTokenPoliciesOk

`func (o *LdapConfigRequest) GetTokenPoliciesOk() (*[]string, bool)`

GetTokenPoliciesOk returns a tuple with the TokenPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenPolicies

`func (o *LdapConfigRequest) SetTokenPolicies(v []string)`

SetTokenPolicies sets TokenPolicies field to given value.

### HasTokenPolicies

`func (o *LdapConfigRequest) HasTokenPolicies() bool`

HasTokenPolicies returns a boolean if a field has been set.

### GetTokenTtl

`func (o *LdapConfigRequest) GetTokenTtl() int32`

GetTokenTtl returns the TokenTtl field if non-nil, zero value otherwise.

### GetTokenTtlOk

`func (o *LdapConfigRequest) GetTokenTtlOk() (*int32, bool)`

GetTokenTtlOk returns a tuple with the TokenTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtl

`func (o *LdapConfigRequest) SetTokenTtl(v int32)`

SetTokenTtl sets TokenTtl field to given value.

### HasTokenTtl

`func (o *LdapConfigRequest) HasTokenTtl() bool`

HasTokenTtl returns a boolean if a field has been set.

### GetTokenType

`func (o *LdapConfigRequest) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *LdapConfigRequest) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *LdapConfigRequest) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *LdapConfigRequest) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUpndomain

`func (o *LdapConfigRequest) GetUpndomain() string`

GetUpndomain returns the Upndomain field if non-nil, zero value otherwise.

### GetUpndomainOk

`func (o *LdapConfigRequest) GetUpndomainOk() (*string, bool)`

GetUpndomainOk returns a tuple with the Upndomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpndomain

`func (o *LdapConfigRequest) SetUpndomain(v string)`

SetUpndomain sets Upndomain field to given value.

### HasUpndomain

`func (o *LdapConfigRequest) HasUpndomain() bool`

HasUpndomain returns a boolean if a field has been set.

### GetUrl

`func (o *LdapConfigRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *LdapConfigRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *LdapConfigRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *LdapConfigRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsePre111GroupCnBehavior

`func (o *LdapConfigRequest) GetUsePre111GroupCnBehavior() bool`

GetUsePre111GroupCnBehavior returns the UsePre111GroupCnBehavior field if non-nil, zero value otherwise.

### GetUsePre111GroupCnBehaviorOk

`func (o *LdapConfigRequest) GetUsePre111GroupCnBehaviorOk() (*bool, bool)`

GetUsePre111GroupCnBehaviorOk returns a tuple with the UsePre111GroupCnBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePre111GroupCnBehavior

`func (o *LdapConfigRequest) SetUsePre111GroupCnBehavior(v bool)`

SetUsePre111GroupCnBehavior sets UsePre111GroupCnBehavior field to given value.

### HasUsePre111GroupCnBehavior

`func (o *LdapConfigRequest) HasUsePre111GroupCnBehavior() bool`

HasUsePre111GroupCnBehavior returns a boolean if a field has been set.

### GetUseTokenGroups

`func (o *LdapConfigRequest) GetUseTokenGroups() bool`

GetUseTokenGroups returns the UseTokenGroups field if non-nil, zero value otherwise.

### GetUseTokenGroupsOk

`func (o *LdapConfigRequest) GetUseTokenGroupsOk() (*bool, bool)`

GetUseTokenGroupsOk returns a tuple with the UseTokenGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTokenGroups

`func (o *LdapConfigRequest) SetUseTokenGroups(v bool)`

SetUseTokenGroups sets UseTokenGroups field to given value.

### HasUseTokenGroups

`func (o *LdapConfigRequest) HasUseTokenGroups() bool`

HasUseTokenGroups returns a boolean if a field has been set.

### GetUserattr

`func (o *LdapConfigRequest) GetUserattr() string`

GetUserattr returns the Userattr field if non-nil, zero value otherwise.

### GetUserattrOk

`func (o *LdapConfigRequest) GetUserattrOk() (*string, bool)`

GetUserattrOk returns a tuple with the Userattr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserattr

`func (o *LdapConfigRequest) SetUserattr(v string)`

SetUserattr sets Userattr field to given value.

### HasUserattr

`func (o *LdapConfigRequest) HasUserattr() bool`

HasUserattr returns a boolean if a field has been set.

### GetUserdn

`func (o *LdapConfigRequest) GetUserdn() string`

GetUserdn returns the Userdn field if non-nil, zero value otherwise.

### GetUserdnOk

`func (o *LdapConfigRequest) GetUserdnOk() (*string, bool)`

GetUserdnOk returns a tuple with the Userdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserdn

`func (o *LdapConfigRequest) SetUserdn(v string)`

SetUserdn sets Userdn field to given value.

### HasUserdn

`func (o *LdapConfigRequest) HasUserdn() bool`

HasUserdn returns a boolean if a field has been set.

### GetUserfilter

`func (o *LdapConfigRequest) GetUserfilter() string`

GetUserfilter returns the Userfilter field if non-nil, zero value otherwise.

### GetUserfilterOk

`func (o *LdapConfigRequest) GetUserfilterOk() (*string, bool)`

GetUserfilterOk returns a tuple with the Userfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfilter

`func (o *LdapConfigRequest) SetUserfilter(v string)`

SetUserfilter sets Userfilter field to given value.

### HasUserfilter

`func (o *LdapConfigRequest) HasUserfilter() bool`

HasUserfilter returns a boolean if a field has been set.

### GetUsernameAsAlias

`func (o *LdapConfigRequest) GetUsernameAsAlias() bool`

GetUsernameAsAlias returns the UsernameAsAlias field if non-nil, zero value otherwise.

### GetUsernameAsAliasOk

`func (o *LdapConfigRequest) GetUsernameAsAliasOk() (*bool, bool)`

GetUsernameAsAliasOk returns a tuple with the UsernameAsAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameAsAlias

`func (o *LdapConfigRequest) SetUsernameAsAlias(v bool)`

SetUsernameAsAlias sets UsernameAsAlias field to given value.

### HasUsernameAsAlias

`func (o *LdapConfigRequest) HasUsernameAsAlias() bool`

HasUsernameAsAlias returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

