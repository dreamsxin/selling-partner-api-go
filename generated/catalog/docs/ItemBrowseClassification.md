# ItemBrowseClassification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name for the classification. | 
**ClassificationId** | **string** | Identifier of the classification. | 
**Parent** | Pointer to [**ItemBrowseClassification**](ItemBrowseClassification.md) |  | [optional] 

## Methods

### NewItemBrowseClassification

`func NewItemBrowseClassification(displayName string, classificationId string, ) *ItemBrowseClassification`

NewItemBrowseClassification instantiates a new ItemBrowseClassification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemBrowseClassificationWithDefaults

`func NewItemBrowseClassificationWithDefaults() *ItemBrowseClassification`

NewItemBrowseClassificationWithDefaults instantiates a new ItemBrowseClassification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ItemBrowseClassification) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ItemBrowseClassification) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ItemBrowseClassification) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetClassificationId

`func (o *ItemBrowseClassification) GetClassificationId() string`

GetClassificationId returns the ClassificationId field if non-nil, zero value otherwise.

### GetClassificationIdOk

`func (o *ItemBrowseClassification) GetClassificationIdOk() (*string, bool)`

GetClassificationIdOk returns a tuple with the ClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationId

`func (o *ItemBrowseClassification) SetClassificationId(v string)`

SetClassificationId sets ClassificationId field to given value.


### GetParent

`func (o *ItemBrowseClassification) GetParent() ItemBrowseClassification`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ItemBrowseClassification) GetParentOk() (*ItemBrowseClassification, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ItemBrowseClassification) SetParent(v ItemBrowseClassification)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ItemBrowseClassification) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


