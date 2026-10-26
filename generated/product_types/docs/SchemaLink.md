# SchemaLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Link** | [**SchemaLinkLink**](SchemaLinkLink.md) |  | 
**Checksum** | **string** | Checksum hash of the schema (Base64 MD5). Use this to verify schema contents, identify changes between schema versions, and for caching. | 

## Methods

### NewSchemaLink

`func NewSchemaLink(link SchemaLinkLink, checksum string, ) *SchemaLink`

NewSchemaLink instantiates a new SchemaLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaLinkWithDefaults

`func NewSchemaLinkWithDefaults() *SchemaLink`

NewSchemaLinkWithDefaults instantiates a new SchemaLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLink

`func (o *SchemaLink) GetLink() SchemaLinkLink`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *SchemaLink) GetLinkOk() (*SchemaLinkLink, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *SchemaLink) SetLink(v SchemaLinkLink)`

SetLink sets Link field to given value.


### GetChecksum

`func (o *SchemaLink) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *SchemaLink) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *SchemaLink) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


