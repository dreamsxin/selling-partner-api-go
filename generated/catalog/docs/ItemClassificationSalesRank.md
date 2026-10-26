# ItemClassificationSalesRank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassificationId** | **string** | Identifier of the classification that is associated with the sales rank. | 
**Title** | **string** | Name of the sales rank. | 
**Link** | Pointer to **string** | Corresponding Amazon retail website URL for the sales category. | [optional] 
**Rank** | **int32** | Sales rank. | 

## Methods

### NewItemClassificationSalesRank

`func NewItemClassificationSalesRank(classificationId string, title string, rank int32, ) *ItemClassificationSalesRank`

NewItemClassificationSalesRank instantiates a new ItemClassificationSalesRank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemClassificationSalesRankWithDefaults

`func NewItemClassificationSalesRankWithDefaults() *ItemClassificationSalesRank`

NewItemClassificationSalesRankWithDefaults instantiates a new ItemClassificationSalesRank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassificationId

`func (o *ItemClassificationSalesRank) GetClassificationId() string`

GetClassificationId returns the ClassificationId field if non-nil, zero value otherwise.

### GetClassificationIdOk

`func (o *ItemClassificationSalesRank) GetClassificationIdOk() (*string, bool)`

GetClassificationIdOk returns a tuple with the ClassificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassificationId

`func (o *ItemClassificationSalesRank) SetClassificationId(v string)`

SetClassificationId sets ClassificationId field to given value.


### GetTitle

`func (o *ItemClassificationSalesRank) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ItemClassificationSalesRank) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ItemClassificationSalesRank) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetLink

`func (o *ItemClassificationSalesRank) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *ItemClassificationSalesRank) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *ItemClassificationSalesRank) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *ItemClassificationSalesRank) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetRank

`func (o *ItemClassificationSalesRank) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ItemClassificationSalesRank) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ItemClassificationSalesRank) SetRank(v int32)`

SetRank sets Rank field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


