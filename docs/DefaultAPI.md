# \DefaultAPI

All URIs are relative to *https://api.eu.onfido.com/v3.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelReport**](DefaultAPI.md#CancelReport) | **Post** /reports/{report_id}/cancel | This endpoint is for cancelling individual paused reports.
[**CompleteTask**](DefaultAPI.md#CompleteTask) | **Post** /workflow_runs/{workflow_run_id}/tasks/{task_id}/complete | Completes a Send / Receive Data task.
[**CreateApplicant**](DefaultAPI.md#CreateApplicant) | **Post** /applicants | Create Applicant
[**CreateCheck**](DefaultAPI.md#CreateCheck) | **Post** /checks | Create a check
[**CreateWebhook**](DefaultAPI.md#CreateWebhook) | **Post** /webhooks | Create a webhook
[**CreateWorkflowRun**](DefaultAPI.md#CreateWorkflowRun) | **Post** /workflow_runs | Create a Workflow Run.
[**DeleteWebhook**](DefaultAPI.md#DeleteWebhook) | **Delete** /webhooks/{webhook_id} | Delete a webhook
[**DestroyApplicant**](DefaultAPI.md#DestroyApplicant) | **Delete** /applicants/{applicant_id} | Delete Applicant
[**DownloadCheck**](DefaultAPI.md#DownloadCheck) | **Get** /checks/{check_id}/download | Download a PDF of the check
[**DownloadDocument**](DefaultAPI.md#DownloadDocument) | **Get** /documents/{document_id}/download | Download a documents raw data
[**DownloadLivePhoto**](DefaultAPI.md#DownloadLivePhoto) | **Get** /live_photos/{live_photo_id}/download | Download live photo
[**DownloadLiveVideo**](DefaultAPI.md#DownloadLiveVideo) | **Get** /live_videos/{live_video_id}/download | Download live video
[**DownloadLiveVideoFrame**](DefaultAPI.md#DownloadLiveVideoFrame) | **Get** /live_videos/{live_video_id}/frame | Download live video frame
[**EditWebhook**](DefaultAPI.md#EditWebhook) | **Put** /webhooks/{webhook_id} | Edit a webhook
[**FindAddresses**](DefaultAPI.md#FindAddresses) | **Get** /addresses/pick | Search for addresses by postcode
[**FindApplicant**](DefaultAPI.md#FindApplicant) | **Get** /applicants/{applicant_id} | Retrieve Applicant
[**FindCheck**](DefaultAPI.md#FindCheck) | **Get** /checks/{check_id} | Retrieve a Check
[**FindDocument**](DefaultAPI.md#FindDocument) | **Get** /documents/{document_id} | A single document can be retrieved by calling this endpoint with the document’s unique identifier.
[**FindLivePhoto**](DefaultAPI.md#FindLivePhoto) | **Get** /live_photos/{live_photo_id} | Retrieve live photo
[**FindLiveVideo**](DefaultAPI.md#FindLiveVideo) | **Get** /live_videos/{live_video_id} | Retrieve live video
[**FindReport**](DefaultAPI.md#FindReport) | **Get** /reports/{report_id} | A single report can be retrieved using this endpoint with the corresponding unique identifier.
[**FindWebhook**](DefaultAPI.md#FindWebhook) | **Get** /webhooks/{webhook_id} | Retrieve a Webhook
[**GenerateSdkToken**](DefaultAPI.md#GenerateSdkToken) | **Post** /sdk_token | Generate a SDK token
[**ListApplicants**](DefaultAPI.md#ListApplicants) | **Get** /applicants | List Applicants
[**ListChecks**](DefaultAPI.md#ListChecks) | **Get** /checks | Retrieve Checks
[**ListDocuments**](DefaultAPI.md#ListDocuments) | **Get** /documents | List documents
[**ListLivePhotos**](DefaultAPI.md#ListLivePhotos) | **Get** /live_photos | List live photos
[**ListLiveVideos**](DefaultAPI.md#ListLiveVideos) | **Get** /live_videos | List live videos
[**ListReports**](DefaultAPI.md#ListReports) | **Get** /reports | All the reports belonging to a particular check can be listed from this endpoint.
[**ListTasks**](DefaultAPI.md#ListTasks) | **Get** /workflow_runs/{workflow_run_id}/tasks | The tasks of a Workflow can be retrieved by calling this endpoint with the unique identifier of the Workflow Run.
[**ListWebhooks**](DefaultAPI.md#ListWebhooks) | **Get** /webhooks | List webhooks
[**ListWorkflowRuns**](DefaultAPI.md#ListWorkflowRuns) | **Get** /workflow_runs | List Workflow Runs.
[**Ping**](DefaultAPI.md#Ping) | **Get** /ping | Run a health check on the Onfido API
[**RestoreApplicant**](DefaultAPI.md#RestoreApplicant) | **Post** /applicants/{applicant_id}/restore | Restore Applicant
[**ResumeCheck**](DefaultAPI.md#ResumeCheck) | **Post** /checks/{check_id}/resume | Resume a Check
[**ResumeReport**](DefaultAPI.md#ResumeReport) | **Post** /reports/{report_id}/resume | This endpoint is for resuming individual paused reports.
[**RetrieveTask**](DefaultAPI.md#RetrieveTask) | **Get** /workflow_runs/{workflow_run_id}/tasks/{task_id} | A single task can be retrieved by calling this endpoint with the unique identifier of the Task and Workflow Run.
[**RetrieveWorkflowRun**](DefaultAPI.md#RetrieveWorkflowRun) | **Get** /workflow_runs/{workflow_run_id} | A single workflow run can be retrieved by calling this endpoint with the unique identifier of the Workflow Run.
[**UpdateApplicant**](DefaultAPI.md#UpdateApplicant) | **Put** /applicants/{applicant_id} | Update Applicant
[**UploadDocument**](DefaultAPI.md#UploadDocument) | **Post** /documents | Upload a document
[**UploadLivePhoto**](DefaultAPI.md#UploadLivePhoto) | **Post** /live_photos | Upload live photo



## CancelReport

> CancelReport(ctx, reportId).Execute()

This endpoint is for cancelling individual paused reports.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.CancelReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CancelReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CompleteTask

> CompleteTask(ctx, workflowRunId, taskId).CompleteTaskRequest(completeTaskRequest).Execute()

Completes a Send / Receive Data task.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	workflowRunId := "workflowRunId_example" // string | The unique identifier of the Workflow Run to which the Task belongs.
	taskId := "taskId_example" // string | The identifier of the Task you want to complete.
	completeTaskRequest := *openapiclient.NewCompleteTaskRequest() // CompleteTaskRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.CompleteTask(context.Background(), workflowRunId, taskId).CompleteTaskRequest(completeTaskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CompleteTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run to which the Task belongs. | 
**taskId** | **string** | The identifier of the Task you want to complete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **completeTaskRequest** | [**CompleteTaskRequest**](CompleteTaskRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicant

> Applicant CreateApplicant(ctx).Applicant(applicant).Execute()

Create Applicant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicant := *openapiclient.NewApplicant() // Applicant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateApplicant(context.Background()).Applicant(applicant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicant`: Applicant
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateApplicant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicant** | [**Applicant**](Applicant.md) |  | 

### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCheck

> Check CreateCheck(ctx).Check(check).Execute()

Create a check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	check := *openapiclient.NewCheck() // Check | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateCheck(context.Background()).Check(check).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCheck`: Check
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateCheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **check** | [**Check**](Check.md) |  | 

### Return type

[**Check**](Check.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWebhook

> Webhook CreateWebhook(ctx).Webhook(webhook).Execute()

Create a webhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	webhook := *openapiclient.NewWebhook("Url_example") // Webhook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateWebhook(context.Background()).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflowRun

> WorkflowRun CreateWorkflowRun(ctx).CreateWorkflowRunRequest(createWorkflowRunRequest).Execute()

Create a Workflow Run.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	createWorkflowRunRequest := *openapiclient.NewCreateWorkflowRunRequest("WorkflowId_example", "ApplicantId_example") // CreateWorkflowRunRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.CreateWorkflowRun(context.Background()).CreateWorkflowRunRequest(createWorkflowRunRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.CreateWorkflowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflowRun`: WorkflowRun
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.CreateWorkflowRun`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowRunRequest** | [**CreateWorkflowRunRequest**](CreateWorkflowRunRequest.md) |  | 

### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWebhook

> DeleteWebhook(ctx, webhookId).Execute()

Delete a webhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DeleteWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyApplicant

> DestroyApplicant(ctx, applicantId).Execute()

Delete Applicant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DestroyApplicant(context.Background(), applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DestroyApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadCheck

> DownloadCheck(ctx, checkId).Execute()

Download a PDF of the check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	checkId := "checkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.DownloadCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDocument

> *os.File DownloadDocument(ctx, documentId).Execute()

Download a documents raw data

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	documentId := "documentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadDocument(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDocument`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLivePhoto

> *os.File DownloadLivePhoto(ctx, livePhotoId).Execute()

Download live photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	livePhotoId := "livePhotoId_example" // string | The live photo’s unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadLivePhoto(context.Background(), livePhotoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadLivePhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLivePhoto`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadLivePhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**livePhotoId** | **string** | The live photo’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLiveVideo

> *os.File DownloadLiveVideo(ctx, liveVideoId).Execute()

Download live video



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	liveVideoId := "liveVideoId_example" // string | The live video’s unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadLiveVideo(context.Background(), liveVideoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadLiveVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLiveVideo`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadLiveVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLiveVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLiveVideoFrame

> *os.File DownloadLiveVideoFrame(ctx, liveVideoId).Execute()

Download live video frame



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	liveVideoId := "liveVideoId_example" // string | The live video’s unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.DownloadLiveVideoFrame(context.Background(), liveVideoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DownloadLiveVideoFrame``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadLiveVideoFrame`: *os.File
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.DownloadLiveVideoFrame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLiveVideoFrameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditWebhook

> Webhook EditWebhook(ctx, webhookId).Webhook(webhook).Execute()

Edit a webhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	webhookId := "webhookId_example" // string | 
	webhook := *openapiclient.NewWebhook("Url_example") // Webhook | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.EditWebhook(context.Background(), webhookId).Webhook(webhook).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.EditWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.EditWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webhook** | [**Webhook**](Webhook.md) |  | 

### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAddresses

> AddressesList FindAddresses(ctx).Postcode(postcode).Execute()

Search for addresses by postcode

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	postcode := "postcode_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindAddresses(context.Background()).Postcode(postcode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAddresses`: AddressesList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postcode** | **string** |  | 

### Return type

[**AddressesList**](AddressesList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindApplicant

> Applicant FindApplicant(ctx, applicantId).Execute()

Retrieve Applicant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindApplicant(context.Background(), applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindApplicant`: Applicant
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindApplicant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindCheck

> Check FindCheck(ctx, checkId).Execute()

Retrieve a Check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	checkId := "checkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindCheck`: Check
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Check**](Check.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindDocument

> Document FindDocument(ctx, documentId).Execute()

A single document can be retrieved by calling this endpoint with the document’s unique identifier.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	documentId := "documentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindDocument(context.Background(), documentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindDocument`: Document
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**documentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Document**](Document.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindLivePhoto

> LivePhoto FindLivePhoto(ctx, livePhotoId).Execute()

Retrieve live photo

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	livePhotoId := "livePhotoId_example" // string | The live photo’s unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindLivePhoto(context.Background(), livePhotoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindLivePhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindLivePhoto`: LivePhoto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindLivePhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**livePhotoId** | **string** | The live photo’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LivePhoto**](LivePhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindLiveVideo

> LiveVideo FindLiveVideo(ctx, liveVideoId).Execute()

Retrieve live video

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	liveVideoId := "liveVideoId_example" // string | The live video’s unique identifier.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindLiveVideo(context.Background(), liveVideoId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindLiveVideo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindLiveVideo`: LiveVideo
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindLiveVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**liveVideoId** | **string** | The live video’s unique identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindLiveVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LiveVideo**](LiveVideo.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindReport

> Report FindReport(ctx, reportId).Execute()

A single report can be retrieved using this endpoint with the corresponding unique identifier.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindReport`: Report
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Report**](Report.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindWebhook

> Webhook FindWebhook(ctx, webhookId).Execute()

Retrieve a Webhook

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	webhookId := "webhookId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.FindWebhook(context.Background(), webhookId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.FindWebhook``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindWebhook`: Webhook
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.FindWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**webhookId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Webhook**](Webhook.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSdkToken

> SdkToken GenerateSdkToken(ctx).SdkToken(sdkToken).Execute()

Generate a SDK token

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	sdkToken := *openapiclient.NewSdkToken() // SdkToken | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GenerateSdkToken(context.Background()).SdkToken(sdkToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GenerateSdkToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSdkToken`: SdkToken
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GenerateSdkToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSdkTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sdkToken** | [**SdkToken**](SdkToken.md) |  | 

### Return type

[**SdkToken**](SdkToken.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicants

> ApplicantsList ListApplicants(ctx).Page(page).PerPage(perPage).IncludeDeleted(includeDeleted).Execute()

List Applicants

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	page := int32(56) // int32 | The page to return. The first page is `page=1` (optional) (default to 1)
	perPage := int32(56) // int32 | The number of objects per page. (optional) (default to 20)
	includeDeleted := true // bool | Whether to also include applicants scheduled for deletion. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListApplicants(context.Background()).Page(page).PerPage(perPage).IncludeDeleted(includeDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListApplicants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListApplicants`: ApplicantsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListApplicants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page to return. The first page is &#x60;page&#x3D;1&#x60; | [default to 1]
 **perPage** | **int32** | The number of objects per page. | [default to 20]
 **includeDeleted** | **bool** | Whether to also include applicants scheduled for deletion. | [default to false]

### Return type

[**ApplicantsList**](ApplicantsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListChecks

> ChecksList ListChecks(ctx).ApplicantId(applicantId).Execute()

Retrieve Checks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListChecks(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListChecks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListChecks`: ChecksList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListChecks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListChecksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** |  | 

### Return type

[**ChecksList**](ChecksList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> DocumentsList ListDocuments(ctx).ApplicantId(applicantId).Execute()

List documents



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListDocuments(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDocuments`: DocumentsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** |  | 

### Return type

[**DocumentsList**](DocumentsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLivePhotos

> LivePhotosList ListLivePhotos(ctx).ApplicantId(applicantId).Execute()

List live photos

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | The id of the applicant the live photos belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListLivePhotos(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLivePhotos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLivePhotos`: LivePhotosList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListLivePhotos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLivePhotosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the live photos belong to. | 

### Return type

[**LivePhotosList**](LivePhotosList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLiveVideos

> LiveVideosList ListLiveVideos(ctx).ApplicantId(applicantId).Execute()

List live videos

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | The id of the applicant the live videos belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListLiveVideos(context.Background()).ApplicantId(applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListLiveVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLiveVideos`: LiveVideosList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListLiveVideos`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLiveVideosRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The id of the applicant the live videos belong to. | 

### Return type

[**LiveVideosList**](LiveVideosList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReports

> ReportsList ListReports(ctx).CheckId(checkId).Execute()

All the reports belonging to a particular check can be listed from this endpoint.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	checkId := "checkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListReports(context.Background()).CheckId(checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListReports``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReports`: ReportsList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListReports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkId** | **string** |  | 

### Return type

[**ReportsList**](ReportsList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> ListTasks200Response ListTasks(ctx, workflowRunId).Execute()

The tasks of a Workflow can be retrieved by calling this endpoint with the unique identifier of the Workflow Run.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	workflowRunId := "workflowRunId_example" // string | The unique identifier of the Workflow Run to which the Tasks belong.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListTasks(context.Background(), workflowRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListTasks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTasks`: ListTasks200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run to which the Tasks belong. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListTasks200Response**](ListTasks200Response.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebhooks

> WebhooksList ListWebhooks(ctx).Execute()

List webhooks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWebhooks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWebhooks`: WebhooksList
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


### Return type

[**WebhooksList**](WebhooksList.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowRuns

> []WorkflowRun ListWorkflowRuns(ctx).Page(page).Status(status).CreatedAtGt(createdAtGt).CreatedAtLt(createdAtLt).Sort(sort).Execute()

List Workflow Runs.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	page := int32(56) // int32 | The number of the page to be retrieved. If not specified, defaults to 1. (optional) (default to 1)
	status := "status_example" // string | A list of comma separated status values to filter the results. Possible values are 'processing', 'awaiting_input', 'approved', 'declined', 'review', 'abandoned' and 'error'. (optional)
	createdAtGt := time.Now() // time.Time | A ISO-8601 date to filter results with a created date greater than (after) the one provided. (optional)
	createdAtLt := time.Now() // time.Time | A ISO-8601 date to filter results with a created date less than (before) the one provided. (optional)
	sort := "sort_example" // string | A string with the value 'desc' or 'asc' that allows to sort the returned list by the completed datetime either descending or ascending, respectively. If not specified, defaults to 'desc'. (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.ListWorkflowRuns(context.Background()).Page(page).Status(status).CreatedAtGt(createdAtGt).CreatedAtLt(createdAtLt).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ListWorkflowRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflowRuns`: []WorkflowRun
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.ListWorkflowRuns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The number of the page to be retrieved. If not specified, defaults to 1. | [default to 1]
 **status** | **string** | A list of comma separated status values to filter the results. Possible values are &#39;processing&#39;, &#39;awaiting_input&#39;, &#39;approved&#39;, &#39;declined&#39;, &#39;review&#39;, &#39;abandoned&#39; and &#39;error&#39;. | 
 **createdAtGt** | **time.Time** | A ISO-8601 date to filter results with a created date greater than (after) the one provided. | 
 **createdAtLt** | **time.Time** | A ISO-8601 date to filter results with a created date less than (before) the one provided. | 
 **sort** | **string** | A string with the value &#39;desc&#39; or &#39;asc&#39; that allows to sort the returned list by the completed datetime either descending or ascending, respectively. If not specified, defaults to &#39;desc&#39;. | [default to &quot;desc&quot;]

### Return type

[**[]WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Ping

> string Ping(ctx).Execute()

Run a health check on the Onfido API

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.Ping(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.Ping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Ping`: string
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.Ping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPingRequest struct via the builder pattern


### Return type

**string**

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreApplicant

> RestoreApplicant(ctx, applicantId).Execute()

Restore Applicant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.RestoreApplicant(context.Background(), applicantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RestoreApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeCheck

> ResumeCheck(ctx, checkId).Execute()

Resume a Check

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	checkId := "checkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResumeCheck(context.Background(), checkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResumeCheck``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**checkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeReport

> ResumeReport(ctx, reportId).Execute()

This endpoint is for resuming individual paused reports.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	reportId := "reportId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultAPI.ResumeReport(context.Background(), reportId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ResumeReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveTask

> Task RetrieveTask(ctx, workflowRunId, taskId).Execute()

A single task can be retrieved by calling this endpoint with the unique identifier of the Task and Workflow Run.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	workflowRunId := "workflowRunId_example" // string | The unique identifier of the Workflow Run to which the Task belongs.
	taskId := "taskId_example" // string | The identifier of the Task you want to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RetrieveTask(context.Background(), workflowRunId, taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RetrieveTask``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveTask`: Task
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RetrieveTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run to which the Task belongs. | 
**taskId** | **string** | The identifier of the Task you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Task**](Task.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveWorkflowRun

> WorkflowRun RetrieveWorkflowRun(ctx, workflowRunId).Execute()

A single workflow run can be retrieved by calling this endpoint with the unique identifier of the Workflow Run.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	workflowRunId := "workflowRunId_example" // string | The unique identifier of the Workflow Run.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.RetrieveWorkflowRun(context.Background(), workflowRunId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.RetrieveWorkflowRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveWorkflowRun`: WorkflowRun
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.RetrieveWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowRunId** | **string** | The unique identifier of the Workflow Run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkflowRun**](WorkflowRun.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicant

> Applicant UpdateApplicant(ctx, applicantId).Applicant(applicant).Execute()

Update Applicant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | 
	applicant := *openapiclient.NewApplicant() // Applicant | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UpdateApplicant(context.Background(), applicantId).Applicant(applicant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UpdateApplicant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateApplicant`: Applicant
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UpdateApplicant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicant** | [**Applicant**](Applicant.md) |  | 

### Return type

[**Applicant**](Applicant.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadDocument

> Document UploadDocument(ctx).ApplicantId(applicantId).Type_(type_).File(file).Side(side).IssuingCountry(issuingCountry).ValidateImageQuality(validateImageQuality).Location(location).Execute()

Upload a document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | The ID of the applicant whose document is being uploaded.
	type_ := "type__example" // string | The type of document.
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded.
	side := "side_example" // string | Either the `front` or `back` of the document. (optional)
	issuingCountry := "issuingCountry_example" // string | The issuing country of the document, a 3-letter ISO code. (optional)
	validateImageQuality := true // bool | Defaults to false. When true the submitted image will undergo an image quality validation which may take up to 5 seconds. (optional)
	location := *openapiclient.NewLocation() // Location |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UploadDocument(context.Background()).ApplicantId(applicantId).Type_(type_).File(file).Side(side).IssuingCountry(issuingCountry).ValidateImageQuality(validateImageQuality).Location(location).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UploadDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadDocument`: Document
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UploadDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** | The ID of the applicant whose document is being uploaded. | 
 **type_** | **string** | The type of document. | 
 **file** | ***os.File** | The file to be uploaded. | 
 **side** | **string** | Either the &#x60;front&#x60; or &#x60;back&#x60; of the document. | 
 **issuingCountry** | **string** | The issuing country of the document, a 3-letter ISO code. | 
 **validateImageQuality** | **bool** | Defaults to false. When true the submitted image will undergo an image quality validation which may take up to 5 seconds. | 
 **location** | [**Location**](Location.md) |  | 

### Return type

[**Document**](Document.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLivePhoto

> LivePhoto UploadLivePhoto(ctx).ApplicantId(applicantId).File(file).AdvancedValidation(advancedValidation).Execute()

Upload live photo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/wallester/go-onfido-openapi"
)

func main() {
	applicantId := "applicantId_example" // string | 
	file := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded.
	advancedValidation := true // bool | Validates that the live photo contains exactly one face. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.UploadLivePhoto(context.Background()).ApplicantId(applicantId).File(file).AdvancedValidation(advancedValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.UploadLivePhoto``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLivePhoto`: LivePhoto
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.UploadLivePhoto`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadLivePhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicantId** | **string** |  | 
 **file** | ***os.File** | The file to be uploaded. | 
 **advancedValidation** | **bool** | Validates that the live photo contains exactly one face. | [default to true]

### Return type

[**LivePhoto**](LivePhoto.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

