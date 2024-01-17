# DocumentODPReasons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhotoOfScreen** | Pointer to **string** | When the applicant&#39;s document is on a physical screen or device. | [optional] 
**Screenshot** | Pointer to **string** | When the applicant has used their mobile phone, tablet, or computer to take a photo within the device. | [optional] 
**DocumentOnPrintedPaper** | Pointer to **string** | When the applicant has previously captured an image of the document, printed it out, and has now taken a photo of this print out to upload. | [optional] 
**Scan** | Pointer to **string** | When the document has clearly been captured using a scanner and there are visible indicators of this | [optional] 

## Methods

### NewDocumentODPReasons

`func NewDocumentODPReasons() *DocumentODPReasons`

NewDocumentODPReasons instantiates a new DocumentODPReasons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentODPReasonsWithDefaults

`func NewDocumentODPReasonsWithDefaults() *DocumentODPReasons`

NewDocumentODPReasonsWithDefaults instantiates a new DocumentODPReasons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhotoOfScreen

`func (o *DocumentODPReasons) GetPhotoOfScreen() string`

GetPhotoOfScreen returns the PhotoOfScreen field if non-nil, zero value otherwise.

### GetPhotoOfScreenOk

`func (o *DocumentODPReasons) GetPhotoOfScreenOk() (*string, bool)`

GetPhotoOfScreenOk returns a tuple with the PhotoOfScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoOfScreen

`func (o *DocumentODPReasons) SetPhotoOfScreen(v string)`

SetPhotoOfScreen sets PhotoOfScreen field to given value.

### HasPhotoOfScreen

`func (o *DocumentODPReasons) HasPhotoOfScreen() bool`

HasPhotoOfScreen returns a boolean if a field has been set.

### GetScreenshot

`func (o *DocumentODPReasons) GetScreenshot() string`

GetScreenshot returns the Screenshot field if non-nil, zero value otherwise.

### GetScreenshotOk

`func (o *DocumentODPReasons) GetScreenshotOk() (*string, bool)`

GetScreenshotOk returns a tuple with the Screenshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshot

`func (o *DocumentODPReasons) SetScreenshot(v string)`

SetScreenshot sets Screenshot field to given value.

### HasScreenshot

`func (o *DocumentODPReasons) HasScreenshot() bool`

HasScreenshot returns a boolean if a field has been set.

### GetDocumentOnPrintedPaper

`func (o *DocumentODPReasons) GetDocumentOnPrintedPaper() string`

GetDocumentOnPrintedPaper returns the DocumentOnPrintedPaper field if non-nil, zero value otherwise.

### GetDocumentOnPrintedPaperOk

`func (o *DocumentODPReasons) GetDocumentOnPrintedPaperOk() (*string, bool)`

GetDocumentOnPrintedPaperOk returns a tuple with the DocumentOnPrintedPaper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentOnPrintedPaper

`func (o *DocumentODPReasons) SetDocumentOnPrintedPaper(v string)`

SetDocumentOnPrintedPaper sets DocumentOnPrintedPaper field to given value.

### HasDocumentOnPrintedPaper

`func (o *DocumentODPReasons) HasDocumentOnPrintedPaper() bool`

HasDocumentOnPrintedPaper returns a boolean if a field has been set.

### GetScan

`func (o *DocumentODPReasons) GetScan() string`

GetScan returns the Scan field if non-nil, zero value otherwise.

### GetScanOk

`func (o *DocumentODPReasons) GetScanOk() (*string, bool)`

GetScanOk returns a tuple with the Scan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScan

`func (o *DocumentODPReasons) SetScan(v string)`

SetScan sets Scan field to given value.

### HasScan

`func (o *DocumentODPReasons) HasScan() bool`

HasScan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


