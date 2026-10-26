# ItemIdentifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierType** | **string** | Type of identifier, such as UPC, EAN, or ISBN. | 
**Identifier** | **string** | Identifier of the item. | 

## Methods

### NewItemIdentifier

`func NewItemIdentifier(identifierType string, identifier string, ) *ItemIdentifier`

NewItemIdentifier instantiates a new ItemIdentifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemIdentifierWithDefaults

`func NewItemIdentifierWithDefaults() *ItemIdentifier`

NewItemIdentifierWithDefaults instantiates a new ItemIdentifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierType

`func (o *ItemIdentifier) GetIdentifierType() string`

GetIdentifierType returns the IdentifierType field if non-nil, zero value otherwise.

### GetIdentifierTypeOk

`func (o *ItemIdentifier) GetIdentifierTypeOk() (*string, bool)`

GetIdentifierTypeOk returns a tuple with the IdentifierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierType

`func (o *ItemIdentifier) SetIdentifierType(v string)`

SetIdentifierType sets IdentifierType field to given value.


### GetIdentifier

`func (o *ItemIdentifier) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ItemIdentifier) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ItemIdentifier) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


