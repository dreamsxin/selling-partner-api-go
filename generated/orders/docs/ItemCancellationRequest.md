# ItemCancellationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requester** | Pointer to **string** | Entity that initiated the cancellation request for this item.   **Possible values**: &#x60;BUYER&#x60; | [optional] 
**CancelReason** | Pointer to **string** | Explanation provided for why the cancellation was requested. | [optional] 

## Methods

### NewItemCancellationRequest

`func NewItemCancellationRequest() *ItemCancellationRequest`

NewItemCancellationRequest instantiates a new ItemCancellationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemCancellationRequestWithDefaults

`func NewItemCancellationRequestWithDefaults() *ItemCancellationRequest`

NewItemCancellationRequestWithDefaults instantiates a new ItemCancellationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequester

`func (o *ItemCancellationRequest) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ItemCancellationRequest) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ItemCancellationRequest) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ItemCancellationRequest) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetCancelReason

`func (o *ItemCancellationRequest) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *ItemCancellationRequest) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *ItemCancellationRequest) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *ItemCancellationRequest) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


