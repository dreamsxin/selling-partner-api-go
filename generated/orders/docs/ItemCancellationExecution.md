# ItemCancellationExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelledBy** | Pointer to **string** | The entity that executed the cancellation for this item.   **Possible values**: &#x60;BUYER&#x60;, &#x60;MERCHANT&#x60;, &#x60;AMAZON&#x60;. | [optional] 
**CancelReason** | Pointer to **string** | The provided explanation for why the cancellation occurred. | [optional] 

## Methods

### NewItemCancellationExecution

`func NewItemCancellationExecution() *ItemCancellationExecution`

NewItemCancellationExecution instantiates a new ItemCancellationExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemCancellationExecutionWithDefaults

`func NewItemCancellationExecutionWithDefaults() *ItemCancellationExecution`

NewItemCancellationExecutionWithDefaults instantiates a new ItemCancellationExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelledBy

`func (o *ItemCancellationExecution) GetCancelledBy() string`

GetCancelledBy returns the CancelledBy field if non-nil, zero value otherwise.

### GetCancelledByOk

`func (o *ItemCancellationExecution) GetCancelledByOk() (*string, bool)`

GetCancelledByOk returns a tuple with the CancelledBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledBy

`func (o *ItemCancellationExecution) SetCancelledBy(v string)`

SetCancelledBy sets CancelledBy field to given value.

### HasCancelledBy

`func (o *ItemCancellationExecution) HasCancelledBy() bool`

HasCancelledBy returns a boolean if a field has been set.

### GetCancelReason

`func (o *ItemCancellationExecution) GetCancelReason() string`

GetCancelReason returns the CancelReason field if non-nil, zero value otherwise.

### GetCancelReasonOk

`func (o *ItemCancellationExecution) GetCancelReasonOk() (*string, bool)`

GetCancelReasonOk returns a tuple with the CancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelReason

`func (o *ItemCancellationExecution) SetCancelReason(v string)`

SetCancelReason sets CancelReason field to given value.

### HasCancelReason

`func (o *ItemCancellationExecution) HasCancelReason() bool`

HasCancelReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


