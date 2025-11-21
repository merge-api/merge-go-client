# Reference
## Ats AccountDetails
<details><summary><code>client.Ats.AccountDetails.Retrieve() -> *ats.AccountDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details for a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.AccountDetails.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats AccountToken
<details><summary><code>client.Ats.AccountToken.Retrieve(PublicToken) -> *ats.AccountToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account token for the end user with the provided public token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.AccountToken.Retrieve(
        context.TODO(),
        "public_token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**publicToken:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Activities
<details><summary><code>client.Ats.Activities.List() -> *ats.PaginatedActivityList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Activity` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ActivitiesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: ats.ActivitiesListRequestRemoteFieldsActivityType.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        ShowEnumOrigins: ats.ActivitiesListRequestShowEnumOriginsActivityType.Ptr(),
        UserId: merge.String(
            "user_id",
        ),
    }
client.Ats.Activities.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*ats.ActivitiesListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*ats.ActivitiesListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**userId:** `*string` — If provided, will only return activities done by this user.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Activities.Create(request) -> *ats.ActivityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Activity` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ActivityEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ats.ActivityRequest{},
        RemoteUserId: "remote_user_id",
    }
client.Ats.Activities.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ats.ActivityRequest` 
    
</dd>
</dl>

<dl>
<dd>

**remoteUserId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Activities.Retrieve(Id) -> *ats.Activity</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Activity` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ActivitiesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: ats.ActivitiesRetrieveRequestRemoteFieldsActivityType.Ptr(),
        ShowEnumOrigins: ats.ActivitiesRetrieveRequestShowEnumOriginsActivityType.Ptr(),
    }
client.Ats.Activities.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*ats.ActivitiesRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*ats.ActivitiesRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Activities.MetaPostRetrieve() -> *ats.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Activity` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Activities.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Applications
<details><summary><code>client.Ats.Applications.List() -> *ats.PaginatedApplicationList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Application` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ApplicationsListRequest{
        CandidateId: merge.String(
            "candidate_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreditedToId: merge.String(
            "credited_to_id",
        ),
        CurrentStageId: merge.String(
            "current_stage_id",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        JobId: merge.String(
            "job_id",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RejectReasonId: merge.String(
            "reject_reason_id",
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Source: merge.String(
            "source",
        ),
    }
client.Ats.Applications.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**candidateId:** `*string` — If provided, will only return applications for this candidate.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**creditedToId:** `*string` — If provided, will only return applications credited to this user.
    
</dd>
</dl>

<dl>
<dd>

**currentStageId:** `*string` — If provided, will only return applications at this interview stage.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.ApplicationsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**jobId:** `*string` — If provided, will only return applications for this job.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**rejectReasonId:** `*string` — If provided, will only return applications with this reject reason.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**source:** `*string` — If provided, will only return applications with this source.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Applications.Create(request) -> *ats.ApplicationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Application` object with the given values.
For certain integrations, but not all, our API detects duplicate candidates and will associate applications with existing records in the third-party. New candidates are created and automatically linked to the application.

See our [Help Center article](https://help.merge.dev/en/articles/10012366-updates-to-post-applications-oct-2024) for detailed support per integration.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ApplicationEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ats.ApplicationRequest{},
        RemoteUserId: "remote_user_id",
    }
client.Ats.Applications.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ats.ApplicationRequest` 
    
</dd>
</dl>

<dl>
<dd>

**remoteUserId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Applications.Retrieve(Id) -> *ats.Application</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Application` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ApplicationsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Applications.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.ApplicationsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Applications.ChangeStageCreate(Id, request) -> *ats.ApplicationResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates the `current_stage` field of an `Application` object
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.UpdateApplicationStageRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
    }
client.Ats.Applications.ChangeStageCreate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**jobInterviewStage:** `*string` — The interview stage to move the application to.
    
</dd>
</dl>

<dl>
<dd>

**remoteUserId:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Applications.MetaPostRetrieve() -> *ats.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Application` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ApplicationsMetaPostRetrieveRequest{
        ApplicationRemoteTemplateId: merge.String(
            "application_remote_template_id",
        ),
    }
client.Ats.Applications.MetaPostRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**applicationRemoteTemplateId:** `*string` — The template ID associated with the nested application in the request.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats AsyncPassthrough
<details><summary><code>client.Ats.AsyncPassthrough.Create(request) -> *ats.AsyncPassthroughReciept</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Asynchronously pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.DataPassthroughRequest{
        Method: ats.MethodEnumGet,
        Path: "/scooters",
    }
client.Ats.AsyncPassthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*ats.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.AsyncPassthrough.Retrieve(AsyncPassthroughReceiptId) -> *ats.AsyncPassthroughRetrieveResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves data from earlier async-passthrough POST request
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.AsyncPassthrough.Retrieve(
        context.TODO(),
        "async_passthrough_receipt_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**asyncPassthroughReceiptId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Attachments
<details><summary><code>client.Ats.Attachments.List() -> *ats.PaginatedAttachmentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Attachment` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.AttachmentsListRequest{
        CandidateId: merge.String(
            "candidate_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Attachments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**candidateId:** `*string` — If provided, will only return attachments for this candidate.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Attachments.Create(request) -> *ats.AttachmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Attachment` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.AttachmentEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ats.AttachmentRequest{},
        RemoteUserId: "remote_user_id",
    }
client.Ats.Attachments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ats.AttachmentRequest` 
    
</dd>
</dl>

<dl>
<dd>

**remoteUserId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Attachments.Retrieve(Id) -> *ats.Attachment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Attachment` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.AttachmentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Attachments.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Attachments.MetaPostRetrieve() -> *ats.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Attachment` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Attachments.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats AuditTrail
<details><summary><code>client.Ats.AuditTrail.List() -> *ats.PaginatedAuditLogEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a list of audit trail events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.AuditTrailListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EventType: merge.String(
            "event_type",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        UserEmail: merge.String(
            "user_email",
        ),
    }
client.Ats.AuditTrail.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include audit trail events that occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*string` — If included, will only include events with the given event type. Possible values include: `CREATED_REMOTE_PRODUCTION_API_KEY`, `DELETED_REMOTE_PRODUCTION_API_KEY`, `CREATED_TEST_API_KEY`, `DELETED_TEST_API_KEY`, `REGENERATED_PRODUCTION_API_KEY`, `REGENERATED_WEBHOOK_SIGNATURE`, `INVITED_USER`, `TWO_FACTOR_AUTH_ENABLED`, `TWO_FACTOR_AUTH_DISABLED`, `DELETED_LINKED_ACCOUNT`, `DELETED_ALL_COMMON_MODELS_FOR_LINKED_ACCOUNT`, `CREATED_DESTINATION`, `DELETED_DESTINATION`, `CHANGED_DESTINATION`, `CHANGED_SCOPES`, `CHANGED_PERSONAL_INFORMATION`, `CHANGED_ORGANIZATION_SETTINGS`, `ENABLED_INTEGRATION`, `DISABLED_INTEGRATION`, `ENABLED_CATEGORY`, `DISABLED_CATEGORY`, `CHANGED_PASSWORD`, `RESET_PASSWORD`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `CREATED_INTEGRATION_WIDE_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_FIELD_MAPPING`, `CHANGED_INTEGRATION_WIDE_FIELD_MAPPING`, `CHANGED_LINKED_ACCOUNT_FIELD_MAPPING`, `DELETED_INTEGRATION_WIDE_FIELD_MAPPING`, `DELETED_LINKED_ACCOUNT_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `CHANGED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `DELETED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `FORCED_LINKED_ACCOUNT_RESYNC`, `MUTED_ISSUE`, `GENERATED_MAGIC_LINK`, `ENABLED_MERGE_WEBHOOK`, `DISABLED_MERGE_WEBHOOK`, `MERGE_WEBHOOK_TARGET_CHANGED`, `END_USER_CREDENTIALS_ACCESSED`
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include audit trail events that occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**userEmail:** `*string` — If provided, this will return events associated with the specified user email. Please note that the email address reflects the user's email at the time of the event, and may not be their current email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats AvailableActions
<details><summary><code>client.Ats.AvailableActions.Retrieve() -> *ats.AvailableActions</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of models and actions available for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.AvailableActions.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Candidates
<details><summary><code>client.Ats.Candidates.List() -> *ats.PaginatedCandidateList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Candidate` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.CandidatesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmailAddresses: merge.String(
            "email_addresses",
        ),
        FirstName: merge.String(
            "first_name",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        LastName: merge.String(
            "last_name",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Tags: merge.String(
            "tags",
        ),
    }
client.Ats.Candidates.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**emailAddresses:** `*string` — If provided, will only return candidates with these email addresses; multiple addresses can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.CandidatesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `*string` — If provided, will only return candidates with this first name.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `*string` — If provided, will only return candidates with this last name.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**tags:** `*string` — If provided, will only return candidates with these tags; multiple tags can be separated by commas.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Candidates.Create(request) -> *ats.CandidateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Candidate` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.CandidateEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ats.CandidateRequest{},
        RemoteUserId: "remote_user_id",
    }
client.Ats.Candidates.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ats.CandidateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**remoteUserId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Candidates.Retrieve(Id) -> *ats.Candidate</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Candidate` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.CandidatesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Candidates.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.CandidatesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Candidates.PartialUpdate(Id, request) -> *ats.CandidateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a `Candidate` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.PatchedCandidateEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ats.PatchedCandidateRequest{},
        RemoteUserId: "remote_user_id",
    }
client.Ats.Candidates.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ats.PatchedCandidateRequest` 
    
</dd>
</dl>

<dl>
<dd>

**remoteUserId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Candidates.IgnoreCreate(ModelId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.IgnoreCommonModelRequest{
        Reason: &ats.IgnoreCommonModelRequestReason{
            ReasonEnum: ats.ReasonEnumGeneralCustomerRequest,
        },
    }
client.Ats.Candidates.IgnoreCreate(
        context.TODO(),
        "model_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**modelId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*ats.IgnoreCommonModelRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Candidates.MetaPatchRetrieve(Id) -> *ats.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Candidate` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Candidates.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Candidates.MetaPostRetrieve() -> *ats.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Candidate` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Candidates.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Scopes
<details><summary><code>client.Ats.Scopes.DefaultScopesRetrieve() -> *ats.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the default permissions for Merge Common Models and fields across all Linked Accounts of a given category. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Scopes.DefaultScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Scopes.LinkedAccountScopesRetrieve() -> *ats.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all available permissions for Merge Common Models and fields for a single Linked Account. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Scopes.LinkedAccountScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Scopes.LinkedAccountScopesCreate(request) -> *ats.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update permissions for any Common Model or field for a single Linked Account. Any Scopes not set in this POST request will inherit the default Scopes. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.LinkedAccountCommonModelScopeDeserializerRequest{
        CommonModels: []*ats.IndividualCommonModelScopeDeserializerRequest{
            &ats.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Employee",
                ModelPermissions: map[string]*ats.ModelPermissionDeserializerRequest{
                    "READ": &ats.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            true,
                        ),
                    },
                    "WRITE": &ats.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
                FieldPermissions: &ats.FieldPermissionDeserializerRequest{
                    EnabledFields: []any{
                        "avatar",
                        "home_location",
                    },
                    DisabledFields: []any{
                        "work_location",
                    },
                },
            },
            &ats.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Benefit",
                ModelPermissions: map[string]*ats.ModelPermissionDeserializerRequest{
                    "WRITE": &ats.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
            },
        },
    }
client.Ats.Scopes.LinkedAccountScopesCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `[]*ats.IndividualCommonModelScopeDeserializerRequest` — The common models you want to update the scopes for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats DeleteAccount
<details><summary><code>client.Ats.DeleteAccount.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.DeleteAccount.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Departments
<details><summary><code>client.Ats.Departments.List() -> *ats.PaginatedDepartmentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Department` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.DepartmentsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Departments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Departments.Retrieve(Id) -> *ats.Department</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Department` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.DepartmentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Departments.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Eeocs
<details><summary><code>client.Ats.Eeocs.List() -> *ats.PaginatedEeocList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `EEOC` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.EeocsListRequest{
        CandidateId: merge.String(
            "candidate_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: ats.EeocsListRequestRemoteFieldsDisabilityStatus.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        ShowEnumOrigins: ats.EeocsListRequestShowEnumOriginsDisabilityStatus.Ptr(),
    }
client.Ats.Eeocs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**candidateId:** `*string` — If provided, will only return EEOC info for this candidate.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*ats.EeocsListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*ats.EeocsListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Eeocs.Retrieve(Id) -> *ats.Eeoc</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `EEOC` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.EeocsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: ats.EeocsRetrieveRequestRemoteFieldsDisabilityStatus.Ptr(),
        ShowEnumOrigins: ats.EeocsRetrieveRequestShowEnumOriginsDisabilityStatus.Ptr(),
    }
client.Ats.Eeocs.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*ats.EeocsRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*ats.EeocsRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats FieldMapping
<details><summary><code>client.Ats.FieldMapping.FieldMappingsRetrieve() -> *ats.FieldMappingApiInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all Field Mappings for this Linked Account. Field Mappings are mappings between third-party Remote Fields and user defined Merge fields. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.FieldMappingsRetrieveRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
    }
client.Ats.FieldMapping.FieldMappingsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.FieldMapping.FieldMappingsCreate(request) -> *ats.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create new Field Mappings that will be available after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.CreateFieldMappingRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
        TargetFieldName: "example_target_field_name",
        TargetFieldDescription: "this is a example description of the target field",
        RemoteFieldTraversalPath: []any{
            "example_remote_field",
        },
        RemoteMethod: "GET",
        RemoteUrlPath: "/example-url-path",
        CommonModelName: "ExampleCommonModel",
    }
client.Ats.FieldMapping.FieldMappingsCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldName:** `string` — The name of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldDescription:** `string` — The description of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**commonModelName:** `string` — The name of the Common Model that the remote field corresponds to in a given category.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.FieldMapping.FieldMappingsDestroy(FieldMappingId) -> *ats.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes Field Mappings for a Linked Account. All data related to this Field Mapping will be deleted and these changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.FieldMapping.FieldMappingsDestroy(
        context.TODO(),
        "field_mapping_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.FieldMapping.FieldMappingsPartialUpdate(FieldMappingId, request) -> *ats.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update existing Field Mappings for a Linked Account. Changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.PatchedEditFieldMappingRequest{}
client.Ats.FieldMapping.FieldMappingsPartialUpdate(
        context.TODO(),
        "field_mapping_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `*string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `*string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.FieldMapping.RemoteFieldsRetrieve() -> *ats.RemoteFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all remote fields for a Linked Account. Remote fields are third-party fields that are accessible after initial sync if remote_data is enabled. You can use remote fields to override existing Merge fields or map a new Merge field. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.RemoteFieldsRetrieveRequest{
        CommonModels: merge.String(
            "common_models",
        ),
        IncludeExampleValues: merge.String(
            "include_example_values",
        ),
    }
client.Ats.FieldMapping.RemoteFieldsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `*string` — A comma seperated list of Common Model names. If included, will only return Remote Fields for those Common Models.
    
</dd>
</dl>

<dl>
<dd>

**includeExampleValues:** `*string` — If true, will include example values, where available, for remote fields in the 3rd party platform. These examples come from active data from your customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.FieldMapping.TargetFieldsRetrieve() -> *ats.ExternalTargetFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all organization-wide Target Fields, this will not include any Linked Account specific Target Fields. Organization-wide Target Fields are additional fields appended to the Merge Common Model for all Linked Accounts in a category. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/target-fields/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.FieldMapping.TargetFieldsRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats GenerateKey
<details><summary><code>client.Ats.GenerateKey.Create(request) -> *ats.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a remote key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.GenerateRemoteKeyRequest{
        Name: "Remote Deployment Key 1",
    }
client.Ats.GenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Interviews
<details><summary><code>client.Ats.Interviews.List() -> *ats.PaginatedScheduledInterviewList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `ScheduledInterview` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.InterviewsListRequest{
        ApplicationId: merge.String(
            "application_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        JobId: merge.String(
            "job_id",
        ),
        JobInterviewStageId: merge.String(
            "job_interview_stage_id",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OrganizerId: merge.String(
            "organizer_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Interviews.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**applicationId:** `*string` — If provided, will only return interviews for this application.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.InterviewsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**jobId:** `*string` — If provided, wll only return interviews organized for this job.
    
</dd>
</dl>

<dl>
<dd>

**jobInterviewStageId:** `*string` — If provided, will only return interviews at this stage.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**organizerId:** `*string` — If provided, will only return interviews organized by this user.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Interviews.Create(request) -> *ats.ScheduledInterviewResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `ScheduledInterview` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ScheduledInterviewEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ats.ScheduledInterviewRequest{},
        RemoteUserId: "remote_user_id",
    }
client.Ats.Interviews.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ats.ScheduledInterviewRequest` 
    
</dd>
</dl>

<dl>
<dd>

**remoteUserId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Interviews.Retrieve(Id) -> *ats.ScheduledInterview</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `ScheduledInterview` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.InterviewsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Interviews.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.InterviewsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Interviews.MetaPostRetrieve() -> *ats.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `ScheduledInterview` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Interviews.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Issues
<details><summary><code>client.Ats.Issues.List() -> *ats.PaginatedIssueList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets all issues for Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.IssuesListRequest{
        AccountToken: merge.String(
            "account_token",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        FirstIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        FirstIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeMuted: merge.String(
            "include_muted",
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        LastIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LastIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LinkedAccountId: merge.String(
            "linked_account_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        Status: ats.IssuesListRequestStatusOngoing.Ptr(),
    }
client.Ats.Issues.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include issues whose most recent action occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose first incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose first incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**includeMuted:** `*string` — If true, will include muted issues
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose last incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose last incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**linkedAccountId:** `*string` — If provided, will only include issues pertaining to the linked account passed in.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include issues whose most recent action occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**status:** `*ats.IssuesListRequestStatus` 

Status of the issue. Options: ('ONGOING', 'RESOLVED')

* `ONGOING` - ONGOING
* `RESOLVED` - RESOLVED
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Issues.Retrieve(Id) -> *ats.Issue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific issue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.Issues.Retrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats JobInterviewStages
<details><summary><code>client.Ats.JobInterviewStages.List() -> *ats.PaginatedJobInterviewStageList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `JobInterviewStage` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.JobInterviewStagesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        JobId: merge.String(
            "job_id",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.JobInterviewStages.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**jobId:** `*string` — If provided, will only return interview stages for this job.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.JobInterviewStages.Retrieve(Id) -> *ats.JobInterviewStage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `JobInterviewStage` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.JobInterviewStagesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.JobInterviewStages.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats JobPostings
<details><summary><code>client.Ats.JobPostings.List() -> *ats.PaginatedJobPostingList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `JobPosting` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.JobPostingsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Status: ats.JobPostingsListRequestStatusClosed.Ptr(),
    }
client.Ats.JobPostings.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*ats.JobPostingsListRequestStatus` 

If provided, will only return Job Postings with this status. Options: ('PUBLISHED', 'CLOSED', 'DRAFT', 'INTERNAL', 'PENDING')

* `PUBLISHED` - PUBLISHED
* `CLOSED` - CLOSED
* `DRAFT` - DRAFT
* `INTERNAL` - INTERNAL
* `PENDING` - PENDING
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.JobPostings.Retrieve(Id) -> *ats.JobPosting</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `JobPosting` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.JobPostingsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.JobPostings.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Jobs
<details><summary><code>client.Ats.Jobs.List() -> *ats.PaginatedJobList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Job` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.JobsListRequest{
        Code: merge.String(
            "code",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Offices: merge.String(
            "offices",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Status: ats.JobsListRequestStatusArchived.Ptr(),
    }
client.Ats.Jobs.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**code:** `*string` — If provided, will only return jobs with this code.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.JobsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**offices:** `*string` — If provided, will only return jobs for this office; multiple offices can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*ats.JobsListRequestStatus` 

If provided, will only return jobs with this status. Options: ('OPEN', 'CLOSED', 'DRAFT', 'ARCHIVED', 'PENDING')

* `OPEN` - OPEN
* `CLOSED` - CLOSED
* `DRAFT` - DRAFT
* `ARCHIVED` - ARCHIVED
* `PENDING` - PENDING
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Jobs.Retrieve(Id) -> *ats.Job</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Job` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.JobsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Jobs.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.JobsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Jobs.ScreeningQuestionsList(JobId) -> *ats.PaginatedScreeningQuestionList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `ScreeningQuestion` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.JobsScreeningQuestionsListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Ats.Jobs.ScreeningQuestionsList(
        context.TODO(),
        "job_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**jobId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.JobsScreeningQuestionsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats LinkToken
<details><summary><code>client.Ats.LinkToken.Create(request) -> *ats.LinkToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a link token to be used when linking a new end user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.EndUserDetailsRequest{
        EndUserEmailAddress: "example@gmail.com",
        EndUserOrganizationName: "Test Organization",
        EndUserOriginId: "12345",
        Categories: []ats.CategoriesEnum{
            ats.CategoriesEnumHris,
            ats.CategoriesEnumAts,
        },
    }
client.Ats.LinkToken.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**endUserEmailAddress:** `string` — Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `string` — Your end user's organization.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `string` — This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
    
</dd>
</dl>

<dl>
<dd>

**categories:** `[]*ats.CategoriesEnum` — The integration categories to show in Merge Link.
    
</dd>
</dl>

<dl>
<dd>

**integration:** `*string` — The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://docs.merge.dev/guides/merge-link/single-integration/.
    
</dd>
</dl>

<dl>
<dd>

**linkExpiryMins:** `*int` — An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
    
</dd>
</dl>

<dl>
<dd>

**shouldCreateMagicLinkUrl:** `*bool` — Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**hideAdminMagicLink:** `*bool` — Whether to generate a Magic Link URL on the Admin Needed screen during the linking flow. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**commonModels:** `[]*ats.CommonModelScopesBodyRequest` — An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
    
</dd>
</dl>

<dl>
<dd>

**categoryCommonModelScopes:** `map[string][]*ats.IndividualCommonModelScopeDeserializerRequest` — When creating a Link Token, you can set permissions for Common Models that will apply to the account that is going to be linked. Any model or field not specified in link token payload will default to existing settings.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*ats.EndUserDetailsRequestLanguage` 

The following subset of IETF language tags can be used to configure localization.

* `en` - en
* `de` - de
    
</dd>
</dl>

<dl>
<dd>

**areSyncsDisabled:** `*bool` — The boolean that indicates whether initial, periodic, and force syncs will be disabled.
    
</dd>
</dl>

<dl>
<dd>

**integrationSpecificConfig:** `map[string]any` — A JSON object containing integration-specific configuration options.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats LinkedAccounts
<details><summary><code>client.Ats.LinkedAccounts.List() -> *ats.PaginatedAccountDetailsAndActionsList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List linked accounts for your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.LinkedAccountsListRequest{
        Category: ats.LinkedAccountsListRequestCategoryAccounting.Ptr(),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndUserEmailAddress: merge.String(
            "end_user_email_address",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        EndUserOriginId: merge.String(
            "end_user_origin_id",
        ),
        EndUserOriginIds: merge.String(
            "end_user_origin_ids",
        ),
        Id: merge.String(
            "id",
        ),
        Ids: merge.String(
            "ids",
        ),
        IncludeDuplicates: merge.Bool(
            true,
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        IsTestAccount: merge.String(
            "is_test_account",
        ),
        PageSize: merge.Int(
            1,
        ),
        Status: merge.String(
            "status",
        ),
    }
client.Ats.LinkedAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**category:** `*ats.LinkedAccountsListRequestCategory` 

Options: `accounting`, `ats`, `crm`, `filestorage`, `hris`, `mktg`, `ticketing`

* `hris` - hris
* `ats` - ats
* `accounting` - accounting
* `ticketing` - ticketing
* `crm` - crm
* `mktg` - mktg
* `filestorage` - filestorage
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endUserEmailAddress:** `*string` — If provided, will only return linked accounts associated with the given email address.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` — If provided, will only return linked accounts associated with the given organization name.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `*string` — If provided, will only return linked accounts associated with the given origin ID.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginIds:** `*string` — Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
    
</dd>
</dl>

<dl>
<dd>

**includeDuplicates:** `*bool` — If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` — If provided, will only return linked accounts associated with the given integration name.
    
</dd>
</dl>

<dl>
<dd>

**isTestAccount:** `*string` — If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Filter by status. Options: `COMPLETE`, `IDLE`, `INCOMPLETE`, `RELINK_NEEDED`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Offers
<details><summary><code>client.Ats.Offers.List() -> *ats.PaginatedOfferList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Offer` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.OffersListRequest{
        ApplicationId: merge.String(
            "application_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatorId: merge.String(
            "creator_id",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Offers.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**applicationId:** `*string` — If provided, will only return offers for this application.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**creatorId:** `*string` — If provided, will only return offers created by this user.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.OffersListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Offers.Retrieve(Id) -> *ats.Offer</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Offer` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.OffersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Offers.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.OffersRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Offices
<details><summary><code>client.Ats.Offices.List() -> *ats.PaginatedOfficeList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Office` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.OfficesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Offices.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Offices.Retrieve(Id) -> *ats.Office</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Office` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.OfficesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Offices.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Passthrough
<details><summary><code>client.Ats.Passthrough.Create(request) -> *ats.RemoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.DataPassthroughRequest{
        Method: ats.MethodEnumGet,
        Path: "/scooters",
    }
client.Ats.Passthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*ats.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats RegenerateKey
<details><summary><code>client.Ats.RegenerateKey.Create(request) -> *ats.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Exchange remote keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.RemoteKeyForRegenerationRequest{
        Name: "Remote Deployment Key 1",
    }
client.Ats.RegenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats RejectReasons
<details><summary><code>client.Ats.RejectReasons.List() -> *ats.PaginatedRejectReasonList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RejectReason` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.RejectReasonsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.RejectReasons.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.RejectReasons.Retrieve(Id) -> *ats.RejectReason</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `RejectReason` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.RejectReasonsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.RejectReasons.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Scorecards
<details><summary><code>client.Ats.Scorecards.List() -> *ats.PaginatedScorecardList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Scorecard` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ScorecardsListRequest{
        ApplicationId: merge.String(
            "application_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        InterviewId: merge.String(
            "interview_id",
        ),
        InterviewerId: merge.String(
            "interviewer_id",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Scorecards.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**applicationId:** `*string` — If provided, will only return scorecards for this application.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.ScorecardsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**interviewId:** `*string` — If provided, will only return scorecards for this interview.
    
</dd>
</dl>

<dl>
<dd>

**interviewerId:** `*string` — If provided, will only return scorecards for this interviewer.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Scorecards.Retrieve(Id) -> *ats.Scorecard</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Scorecard` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.ScorecardsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Scorecards.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ats.ScorecardsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats SyncStatus
<details><summary><code>client.Ats.SyncStatus.List() -> *ats.PaginatedSyncStatusList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sync status for the current sync and the most recently finished sync. `last_sync_start` represents the most recent time any sync began. `last_sync_finished` represents the most recent time any sync completed. These timestamps may correspond to different sync instances which may result in a sync start time being later than a separate sync completed time. To ensure you are retrieving the latest available data reference the `last_sync_finished` timestamp where `last_sync_result` is `DONE`. Possible values for `status` and `last_sync_result` are `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`. Learn more about sync status in our [Help Center](https://help.merge.dev/en/articles/8184193-merge-sync-statuses).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.SyncStatusListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Ats.SyncStatus.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats ForceResync
<details><summary><code>client.Ats.ForceResync.SyncStatusResyncCreate() -> []*ats.SyncStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Force re-sync of all models. This endpoint is available for monthly, quarterly, and highest sync frequency customers on the Professional or Enterprise plans. Doing so will consume a sync credit for the relevant linked account. Force re-syncs can also be triggered manually in the Merge Dashboard and is available for all customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.ForceResync.SyncStatusResyncCreate(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Tags
<details><summary><code>client.Ats.Tags.List() -> *ats.PaginatedTagList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Tag` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.TagsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Tags.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats Users
<details><summary><code>client.Ats.Users.List() -> *ats.PaginatedRemoteUserList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteUser` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.UsersListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        Email: merge.String(
            "email",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ats.Users.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — If provided, will only return remote users with the given email address
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.Users.Retrieve(Id) -> *ats.RemoteUser</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `RemoteUser` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.UsersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ats.Users.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ats WebhookReceivers
<details><summary><code>client.Ats.WebhookReceivers.List() -> []*ats.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `WebhookReceiver` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ats.WebhookReceivers.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ats.WebhookReceivers.Create(request) -> *ats.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `WebhookReceiver` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ats.WebhookReceiverRequest{
        Event: "event",
        IsActive: true,
    }
client.Ats.WebhookReceivers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**event:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**key:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting AccountDetails
<details><summary><code>client.Accounting.AccountDetails.Retrieve() -> *accounting.AccountDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details for a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.AccountDetails.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting AccountToken
<details><summary><code>client.Accounting.AccountToken.Retrieve(PublicToken) -> *accounting.AccountToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account token for the end user with the provided public token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.AccountToken.Retrieve(
        context.TODO(),
        "public_token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**publicToken:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting AccountingPeriods
<details><summary><code>client.Accounting.AccountingPeriods.List() -> *accounting.PaginatedAccountingPeriodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `AccountingPeriod` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AccountingPeriodsListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.AccountingPeriods.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.AccountingPeriods.Retrieve(Id) -> *accounting.AccountingPeriod</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `AccountingPeriod` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AccountingPeriodsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.AccountingPeriods.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Accounts
<details><summary><code>client.Accounting.Accounts.List() -> *accounting.PaginatedAccountList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Account` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AccountsListRequest{
        AccountType: merge.String(
            "account_type",
        ),
        Classification: accounting.AccountsListRequestClassificationEmpty.Ptr(),
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: accounting.AccountsListRequestRemoteFieldsClassification.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        ShowEnumOrigins: accounting.AccountsListRequestShowEnumOriginsClassification.Ptr(),
        Status: accounting.AccountsListRequestStatusEmpty.Ptr(),
    }
client.Accounting.Accounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountType:** `*string` — If provided, will only return accounts with the passed in enum.
    
</dd>
</dl>

<dl>
<dd>

**classification:** `*accounting.AccountsListRequestClassification` — If provided, will only return accounts with this classification.
    
</dd>
</dl>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return accounts for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return Accounts with this name.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*accounting.AccountsListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*accounting.AccountsListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*accounting.AccountsListRequestStatus` — If provided, will only return accounts with this status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Accounts.Create(request) -> *accounting.AccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Account` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AccountEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.AccountRequest{},
    }
client.Accounting.Accounts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.AccountRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Accounts.Retrieve(Id) -> *accounting.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Account` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AccountsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: accounting.AccountsRetrieveRequestRemoteFieldsClassification.Ptr(),
        ShowEnumOrigins: accounting.AccountsRetrieveRequestShowEnumOriginsClassification.Ptr(),
    }
client.Accounting.Accounts.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*accounting.AccountsRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*accounting.AccountsRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Accounts.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Account` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Accounts.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Addresses
<details><summary><code>client.Accounting.Addresses.Retrieve(Id) -> *accounting.Address</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Address` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AddressesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Addresses.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting AsyncPassthrough
<details><summary><code>client.Accounting.AsyncPassthrough.Create(request) -> *accounting.AsyncPassthroughReciept</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Asynchronously pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.DataPassthroughRequest{
        Method: accounting.MethodEnumGet,
        Path: "/scooters",
    }
client.Accounting.AsyncPassthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*accounting.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.AsyncPassthrough.Retrieve(AsyncPassthroughReceiptId) -> *accounting.AsyncPassthroughRetrieveResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves data from earlier async-passthrough POST request
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.AsyncPassthrough.Retrieve(
        context.TODO(),
        "async_passthrough_receipt_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**asyncPassthroughReceiptId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting AsyncTasks
<details><summary><code>client.Accounting.AsyncTasks.Retrieve(Id) -> *accounting.AsyncPostTask</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `AsyncPostTask` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.AsyncTasks.Retrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Attachments
<details><summary><code>client.Accounting.Attachments.List() -> *accounting.PaginatedAccountingAttachmentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `AccountingAttachment` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AttachmentsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.Attachments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return accounting attachments for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Attachments.Create(request) -> *accounting.AccountingAttachmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `AccountingAttachment` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AccountingAttachmentEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.AccountingAttachmentRequest{},
    }
client.Accounting.Attachments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.AccountingAttachmentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Attachments.Retrieve(Id) -> *accounting.AccountingAttachment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `AccountingAttachment` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AttachmentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Attachments.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Attachments.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `AccountingAttachment` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Attachments.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting AuditTrail
<details><summary><code>client.Accounting.AuditTrail.List() -> *accounting.PaginatedAuditLogEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a list of audit trail events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.AuditTrailListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EventType: merge.String(
            "event_type",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        UserEmail: merge.String(
            "user_email",
        ),
    }
client.Accounting.AuditTrail.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include audit trail events that occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*string` — If included, will only include events with the given event type. Possible values include: `CREATED_REMOTE_PRODUCTION_API_KEY`, `DELETED_REMOTE_PRODUCTION_API_KEY`, `CREATED_TEST_API_KEY`, `DELETED_TEST_API_KEY`, `REGENERATED_PRODUCTION_API_KEY`, `REGENERATED_WEBHOOK_SIGNATURE`, `INVITED_USER`, `TWO_FACTOR_AUTH_ENABLED`, `TWO_FACTOR_AUTH_DISABLED`, `DELETED_LINKED_ACCOUNT`, `DELETED_ALL_COMMON_MODELS_FOR_LINKED_ACCOUNT`, `CREATED_DESTINATION`, `DELETED_DESTINATION`, `CHANGED_DESTINATION`, `CHANGED_SCOPES`, `CHANGED_PERSONAL_INFORMATION`, `CHANGED_ORGANIZATION_SETTINGS`, `ENABLED_INTEGRATION`, `DISABLED_INTEGRATION`, `ENABLED_CATEGORY`, `DISABLED_CATEGORY`, `CHANGED_PASSWORD`, `RESET_PASSWORD`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `CREATED_INTEGRATION_WIDE_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_FIELD_MAPPING`, `CHANGED_INTEGRATION_WIDE_FIELD_MAPPING`, `CHANGED_LINKED_ACCOUNT_FIELD_MAPPING`, `DELETED_INTEGRATION_WIDE_FIELD_MAPPING`, `DELETED_LINKED_ACCOUNT_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `CHANGED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `DELETED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `FORCED_LINKED_ACCOUNT_RESYNC`, `MUTED_ISSUE`, `GENERATED_MAGIC_LINK`, `ENABLED_MERGE_WEBHOOK`, `DISABLED_MERGE_WEBHOOK`, `MERGE_WEBHOOK_TARGET_CHANGED`, `END_USER_CREDENTIALS_ACCESSED`
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include audit trail events that occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**userEmail:** `*string` — If provided, this will return events associated with the specified user email. Please note that the email address reflects the user's email at the time of the event, and may not be their current email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting AvailableActions
<details><summary><code>client.Accounting.AvailableActions.Retrieve() -> *accounting.AvailableActions</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of models and actions available for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.AvailableActions.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting BalanceSheets
<details><summary><code>client.Accounting.BalanceSheets.List() -> *accounting.PaginatedBalanceSheetList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `BalanceSheet` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BalanceSheetsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.BalanceSheets.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return balance sheets for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.BalanceSheets.Retrieve(Id) -> *accounting.BalanceSheet</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `BalanceSheet` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BalanceSheetsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.BalanceSheets.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting BankFeedAccounts
<details><summary><code>client.Accounting.BankFeedAccounts.List() -> *accounting.PaginatedBankFeedAccountList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `BankFeedAccount` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BankFeedAccountsListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.BankFeedAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.BankFeedAccounts.Create(request) -> *accounting.BankFeedAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `BankFeedAccount` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BankFeedAccountEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.BankFeedAccountRequest{},
    }
client.Accounting.BankFeedAccounts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.BankFeedAccountRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.BankFeedAccounts.Retrieve(Id) -> *accounting.BankFeedAccount</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `BankFeedAccount` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BankFeedAccountsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.BankFeedAccounts.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.BankFeedAccounts.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `BankFeedAccount` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.BankFeedAccounts.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting BankFeedTransactions
<details><summary><code>client.Accounting.BankFeedTransactions.List() -> *accounting.PaginatedBankFeedTransactionList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `BankFeedTransaction` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BankFeedTransactionsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsProcessed: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.BankFeedTransactions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isProcessed:** `*bool` — If provided, will only return bank feed transactions with this is_processed value
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.BankFeedTransactions.Create(request) -> *accounting.BankFeedTransactionResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `BankFeedTransaction` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BankFeedTransactionEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.BankFeedTransactionRequestRequest{},
    }
client.Accounting.BankFeedTransactions.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.BankFeedTransactionRequestRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.BankFeedTransactions.Retrieve(Id) -> *accounting.BankFeedTransaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `BankFeedTransaction` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.BankFeedTransactionsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.BankFeedTransactions.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.BankFeedTransactions.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `BankFeedTransaction` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.BankFeedTransactions.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting CashFlowStatements
<details><summary><code>client.Accounting.CashFlowStatements.List() -> *accounting.PaginatedCashFlowStatementList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `CashFlowStatement` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CashFlowStatementsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.CashFlowStatements.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return cash flow statements for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.CashFlowStatements.Retrieve(Id) -> *accounting.CashFlowStatement</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `CashFlowStatement` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CashFlowStatementsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.CashFlowStatements.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting CompanyInfo
<details><summary><code>client.Accounting.CompanyInfo.List() -> *accounting.PaginatedCompanyInfoList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `CompanyInfo` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CompanyInfoListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.CompanyInfo.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.CompanyInfoListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.CompanyInfo.Retrieve(Id) -> *accounting.CompanyInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `CompanyInfo` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CompanyInfoRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.CompanyInfo.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.CompanyInfoRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Contacts
<details><summary><code>client.Accounting.Contacts.List() -> *accounting.PaginatedContactList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Contact` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ContactsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmailAddress: merge.String(
            "email_address",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCustomer: merge.String(
            "is_customer",
        ),
        IsSupplier: merge.String(
            "is_supplier",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Status: accounting.ContactsListRequestStatusEmpty.Ptr(),
    }
client.Accounting.Contacts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return contacts for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — If provided, will only return Contacts that match this email.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ContactsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCustomer:** `*string` — If provided, will only return Contacts that are denoted as customers.
    
</dd>
</dl>

<dl>
<dd>

**isSupplier:** `*string` — If provided, will only return Contacts that are denoted as suppliers.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return Contacts that match this name.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*accounting.ContactsListRequestStatus` — If provided, will only return Contacts that match this status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Contacts.Create(request) -> *accounting.ContactResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Contact` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ContactEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.ContactRequest{},
    }
client.Accounting.Contacts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.ContactRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Contacts.Retrieve(Id) -> *accounting.Contact</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Contact` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ContactsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Contacts.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ContactsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Contacts.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Contact` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Contacts.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Contacts.RemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ContactsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Contacts.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting CreditNotes
<details><summary><code>client.Accounting.CreditNotes.List() -> *accounting.PaginatedCreditNoteList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `CreditNote` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CreditNotesListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: accounting.CreditNotesListRequestRemoteFieldsStatus.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        ShowEnumOrigins: accounting.CreditNotesListRequestShowEnumOriginsStatus.Ptr(),
        TransactionDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TransactionDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Accounting.CreditNotes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return credit notes for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.CreditNotesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*accounting.CreditNotesListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*accounting.CreditNotesListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**transactionDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.CreditNotes.Create(request) -> *accounting.CreditNoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `CreditNote` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CreditNoteEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.CreditNoteRequest{},
    }
client.Accounting.CreditNotes.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.CreditNoteRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.CreditNotes.Retrieve(Id) -> *accounting.CreditNote</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `CreditNote` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CreditNotesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: accounting.CreditNotesRetrieveRequestRemoteFieldsStatus.Ptr(),
        ShowEnumOrigins: accounting.CreditNotesRetrieveRequestShowEnumOriginsStatus.Ptr(),
    }
client.Accounting.CreditNotes.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.CreditNotesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*accounting.CreditNotesRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*accounting.CreditNotesRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.CreditNotes.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `CreditNote` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.CreditNotes.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Scopes
<details><summary><code>client.Accounting.Scopes.DefaultScopesRetrieve() -> *accounting.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the default permissions for Merge Common Models and fields across all Linked Accounts of a given category. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Scopes.DefaultScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Scopes.LinkedAccountScopesRetrieve() -> *accounting.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all available permissions for Merge Common Models and fields for a single Linked Account. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Scopes.LinkedAccountScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Scopes.LinkedAccountScopesCreate(request) -> *accounting.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update permissions for any Common Model or field for a single Linked Account. Any Scopes not set in this POST request will inherit the default Scopes. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.LinkedAccountCommonModelScopeDeserializerRequest{
        CommonModels: []*accounting.IndividualCommonModelScopeDeserializerRequest{
            &accounting.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Employee",
                ModelPermissions: map[string]*accounting.ModelPermissionDeserializerRequest{
                    "READ": &accounting.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            true,
                        ),
                    },
                    "WRITE": &accounting.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
                FieldPermissions: &accounting.FieldPermissionDeserializerRequest{
                    EnabledFields: []any{
                        "avatar",
                        "home_location",
                    },
                    DisabledFields: []any{
                        "work_location",
                    },
                },
            },
            &accounting.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Benefit",
                ModelPermissions: map[string]*accounting.ModelPermissionDeserializerRequest{
                    "WRITE": &accounting.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
            },
        },
    }
client.Accounting.Scopes.LinkedAccountScopesCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `[]*accounting.IndividualCommonModelScopeDeserializerRequest` — The common models you want to update the scopes for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting DeleteAccount
<details><summary><code>client.Accounting.DeleteAccount.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.DeleteAccount.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Employees
<details><summary><code>client.Accounting.Employees.List() -> *accounting.PaginatedEmployeeList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Employee` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.EmployeesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Employees.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Employees.Retrieve(Id) -> *accounting.Employee</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Employee` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.EmployeesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Employees.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Expenses
<details><summary><code>client.Accounting.Expenses.List() -> *accounting.PaginatedExpenseList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Expense` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ExpensesListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        TransactionDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TransactionDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Accounting.Expenses.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return expenses for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ExpensesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Expenses.Create(request) -> *accounting.ExpenseResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Expense` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ExpenseEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.ExpenseRequest{},
    }
client.Accounting.Expenses.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.ExpenseRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Expenses.Retrieve(Id) -> *accounting.Expense</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Expense` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ExpensesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Expenses.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ExpensesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Expenses.LinesRemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ExpensesLinesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Expenses.LinesRemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Expenses.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Expense` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Expenses.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Expenses.RemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ExpensesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Expenses.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting FieldMapping
<details><summary><code>client.Accounting.FieldMapping.FieldMappingsRetrieve() -> *accounting.FieldMappingApiInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all Field Mappings for this Linked Account. Field Mappings are mappings between third-party Remote Fields and user defined Merge fields. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.FieldMappingsRetrieveRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
    }
client.Accounting.FieldMapping.FieldMappingsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.FieldMapping.FieldMappingsCreate(request) -> *accounting.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create new Field Mappings that will be available after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.CreateFieldMappingRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
        TargetFieldName: "example_target_field_name",
        TargetFieldDescription: "this is a example description of the target field",
        RemoteFieldTraversalPath: []any{
            "example_remote_field",
        },
        RemoteMethod: "GET",
        RemoteUrlPath: "/example-url-path",
        CommonModelName: "ExampleCommonModel",
    }
client.Accounting.FieldMapping.FieldMappingsCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldName:** `string` — The name of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldDescription:** `string` — The description of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**commonModelName:** `string` — The name of the Common Model that the remote field corresponds to in a given category.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.FieldMapping.FieldMappingsDestroy(FieldMappingId) -> *accounting.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes Field Mappings for a Linked Account. All data related to this Field Mapping will be deleted and these changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.FieldMapping.FieldMappingsDestroy(
        context.TODO(),
        "field_mapping_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.FieldMapping.FieldMappingsPartialUpdate(FieldMappingId, request) -> *accounting.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update existing Field Mappings for a Linked Account. Changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PatchedEditFieldMappingRequest{}
client.Accounting.FieldMapping.FieldMappingsPartialUpdate(
        context.TODO(),
        "field_mapping_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `*string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `*string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.FieldMapping.RemoteFieldsRetrieve() -> *accounting.RemoteFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all remote fields for a Linked Account. Remote fields are third-party fields that are accessible after initial sync if remote_data is enabled. You can use remote fields to override existing Merge fields or map a new Merge field. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.RemoteFieldsRetrieveRequest{
        CommonModels: merge.String(
            "common_models",
        ),
        IncludeExampleValues: merge.String(
            "include_example_values",
        ),
    }
client.Accounting.FieldMapping.RemoteFieldsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `*string` — A comma seperated list of Common Model names. If included, will only return Remote Fields for those Common Models.
    
</dd>
</dl>

<dl>
<dd>

**includeExampleValues:** `*string` — If true, will include example values, where available, for remote fields in the 3rd party platform. These examples come from active data from your customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.FieldMapping.TargetFieldsRetrieve() -> *accounting.ExternalTargetFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all organization-wide Target Fields, this will not include any Linked Account specific Target Fields. Organization-wide Target Fields are additional fields appended to the Merge Common Model for all Linked Accounts in a category. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/target-fields/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.FieldMapping.TargetFieldsRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting GeneralLedgerTransactions
<details><summary><code>client.Accounting.GeneralLedgerTransactions.List() -> *accounting.PaginatedGeneralLedgerTransactionList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `GeneralLedgerTransaction` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.GeneralLedgerTransactionsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        PostedDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PostedDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.GeneralLedgerTransactions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return general ledger transactions for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.GeneralLedgerTransactionsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**postedDateAfter:** `*time.Time` — If provided, will only return objects posted after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**postedDateBefore:** `*time.Time` — If provided, will only return objects posted before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.GeneralLedgerTransactions.Retrieve(Id) -> *accounting.GeneralLedgerTransaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `GeneralLedgerTransaction` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.GeneralLedgerTransactionsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.GeneralLedgerTransactions.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.GeneralLedgerTransactionsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting GenerateKey
<details><summary><code>client.Accounting.GenerateKey.Create(request) -> *accounting.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a remote key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.GenerateRemoteKeyRequest{
        Name: "Remote Deployment Key 1",
    }
client.Accounting.GenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting IncomeStatements
<details><summary><code>client.Accounting.IncomeStatements.List() -> *accounting.PaginatedIncomeStatementList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `IncomeStatement` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.IncomeStatementsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.IncomeStatements.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return income statements for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.IncomeStatements.Retrieve(Id) -> *accounting.IncomeStatement</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `IncomeStatement` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.IncomeStatementsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.IncomeStatements.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Invoices
<details><summary><code>client.Accounting.Invoices.List() -> *accounting.PaginatedInvoiceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Invoice` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.InvoicesListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        ContactId: merge.String(
            "contact_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IssueDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IssueDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Number: merge.String(
            "number",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Status: accounting.InvoicesListRequestStatusDraft.Ptr(),
        Type: accounting.InvoicesListRequestTypeAccountsPayable.Ptr(),
    }
client.Accounting.Invoices.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return invoices for this company.
    
</dd>
</dl>

<dl>
<dd>

**contactId:** `*string` — If provided, will only return invoices for this contact.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.InvoicesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**issueDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**issueDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**number:** `*string` — If provided, will only return Invoices with this number.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*accounting.InvoicesListRequestStatus` 

If provided, will only return Invoices with this status.

* `PAID` - PAID
* `DRAFT` - DRAFT
* `SUBMITTED` - SUBMITTED
* `PARTIALLY_PAID` - PARTIALLY_PAID
* `OPEN` - OPEN
* `VOID` - VOID
    
</dd>
</dl>

<dl>
<dd>

**type_:** `*accounting.InvoicesListRequestType` 

If provided, will only return Invoices with this type.

* `ACCOUNTS_RECEIVABLE` - ACCOUNTS_RECEIVABLE
* `ACCOUNTS_PAYABLE` - ACCOUNTS_PAYABLE
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Invoices.Create(request) -> *accounting.InvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Invoice` object with the given values.
            Including a `PurchaseOrder` id in the `purchase_orders` property will generate an Accounts Payable Invoice from the specified Purchase Order(s).
            
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.InvoiceEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.InvoiceRequest{},
    }
client.Accounting.Invoices.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.InvoiceRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Invoices.Retrieve(Id) -> *accounting.Invoice</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Invoice` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.InvoicesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Invoices.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.InvoicesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Invoices.PartialUpdate(Id, request) -> *accounting.InvoiceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an `Invoice` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PatchedInvoiceEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.InvoiceRequest{},
    }
client.Accounting.Invoices.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.InvoiceRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Invoices.LineItemsRemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.InvoicesLineItemsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Invoices.LineItemsRemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Invoices.MetaPatchRetrieve(Id) -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Invoice` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Invoices.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Invoices.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Invoice` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Invoices.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Invoices.RemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.InvoicesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Invoices.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Issues
<details><summary><code>client.Accounting.Issues.List() -> *accounting.PaginatedIssueList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets all issues for Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.IssuesListRequest{
        AccountToken: merge.String(
            "account_token",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        FirstIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        FirstIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeMuted: merge.String(
            "include_muted",
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        LastIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LastIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LinkedAccountId: merge.String(
            "linked_account_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        Status: accounting.IssuesListRequestStatusOngoing.Ptr(),
    }
client.Accounting.Issues.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include issues whose most recent action occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose first incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose first incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**includeMuted:** `*string` — If true, will include muted issues
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose last incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose last incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**linkedAccountId:** `*string` — If provided, will only include issues pertaining to the linked account passed in.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include issues whose most recent action occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**status:** `*accounting.IssuesListRequestStatus` 

Status of the issue. Options: ('ONGOING', 'RESOLVED')

* `ONGOING` - ONGOING
* `RESOLVED` - RESOLVED
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Issues.Retrieve(Id) -> *accounting.Issue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific issue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Issues.Retrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Items
<details><summary><code>client.Accounting.Items.List() -> *accounting.PaginatedItemList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Item` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ItemsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.Items.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return items for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ItemsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Items.Create(request) -> *accounting.ItemResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Item` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ItemEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.ItemRequestRequest{},
    }
client.Accounting.Items.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.ItemRequestRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Items.Retrieve(Id) -> *accounting.Item</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Item` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ItemsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Items.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ItemsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Items.PartialUpdate(Id, request) -> *accounting.ItemResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an `Item` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PatchedItemEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.PatchedItemRequestRequest{},
    }
client.Accounting.Items.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.PatchedItemRequestRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Items.MetaPatchRetrieve(Id) -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Item` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Items.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Items.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Item` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Items.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting JournalEntries
<details><summary><code>client.Accounting.JournalEntries.List() -> *accounting.PaginatedJournalEntryList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `JournalEntry` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.JournalEntriesListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        TransactionDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TransactionDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Accounting.JournalEntries.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return journal entries for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.JournalEntriesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.JournalEntries.Create(request) -> *accounting.JournalEntryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `JournalEntry` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.JournalEntryEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.JournalEntryRequest{},
    }
client.Accounting.JournalEntries.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.JournalEntryRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.JournalEntries.Retrieve(Id) -> *accounting.JournalEntry</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `JournalEntry` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.JournalEntriesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.JournalEntries.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.JournalEntriesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.JournalEntries.LinesRemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.JournalEntriesLinesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.JournalEntries.LinesRemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.JournalEntries.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `JournalEntry` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.JournalEntries.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.JournalEntries.RemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.JournalEntriesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.JournalEntries.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting LinkToken
<details><summary><code>client.Accounting.LinkToken.Create(request) -> *accounting.LinkToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a link token to be used when linking a new end user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.EndUserDetailsRequest{
        EndUserEmailAddress: "example@gmail.com",
        EndUserOrganizationName: "Test Organization",
        EndUserOriginId: "12345",
        Categories: []accounting.CategoriesEnum{
            accounting.CategoriesEnumHris,
            accounting.CategoriesEnumAts,
        },
    }
client.Accounting.LinkToken.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**endUserEmailAddress:** `string` — Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `string` — Your end user's organization.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `string` — This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
    
</dd>
</dl>

<dl>
<dd>

**categories:** `[]*accounting.CategoriesEnum` — The integration categories to show in Merge Link.
    
</dd>
</dl>

<dl>
<dd>

**integration:** `*string` — The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://docs.merge.dev/guides/merge-link/single-integration/.
    
</dd>
</dl>

<dl>
<dd>

**linkExpiryMins:** `*int` — An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
    
</dd>
</dl>

<dl>
<dd>

**shouldCreateMagicLinkUrl:** `*bool` — Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**hideAdminMagicLink:** `*bool` — Whether to generate a Magic Link URL on the Admin Needed screen during the linking flow. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**commonModels:** `[]*accounting.CommonModelScopesBodyRequest` — An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
    
</dd>
</dl>

<dl>
<dd>

**categoryCommonModelScopes:** `map[string][]*accounting.IndividualCommonModelScopeDeserializerRequest` — When creating a Link Token, you can set permissions for Common Models that will apply to the account that is going to be linked. Any model or field not specified in link token payload will default to existing settings.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*accounting.EndUserDetailsRequestLanguage` 

The following subset of IETF language tags can be used to configure localization.

* `en` - en
* `de` - de
    
</dd>
</dl>

<dl>
<dd>

**areSyncsDisabled:** `*bool` — The boolean that indicates whether initial, periodic, and force syncs will be disabled.
    
</dd>
</dl>

<dl>
<dd>

**integrationSpecificConfig:** `map[string]any` — A JSON object containing integration-specific configuration options.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting LinkedAccounts
<details><summary><code>client.Accounting.LinkedAccounts.List() -> *accounting.PaginatedAccountDetailsAndActionsList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List linked accounts for your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.LinkedAccountsListRequest{
        Category: accounting.LinkedAccountsListRequestCategoryAccounting.Ptr(),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndUserEmailAddress: merge.String(
            "end_user_email_address",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        EndUserOriginId: merge.String(
            "end_user_origin_id",
        ),
        EndUserOriginIds: merge.String(
            "end_user_origin_ids",
        ),
        Id: merge.String(
            "id",
        ),
        Ids: merge.String(
            "ids",
        ),
        IncludeDuplicates: merge.Bool(
            true,
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        IsTestAccount: merge.String(
            "is_test_account",
        ),
        PageSize: merge.Int(
            1,
        ),
        Status: merge.String(
            "status",
        ),
    }
client.Accounting.LinkedAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**category:** `*accounting.LinkedAccountsListRequestCategory` 

Options: `accounting`, `ats`, `crm`, `filestorage`, `hris`, `mktg`, `ticketing`

* `hris` - hris
* `ats` - ats
* `accounting` - accounting
* `ticketing` - ticketing
* `crm` - crm
* `mktg` - mktg
* `filestorage` - filestorage
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endUserEmailAddress:** `*string` — If provided, will only return linked accounts associated with the given email address.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` — If provided, will only return linked accounts associated with the given organization name.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `*string` — If provided, will only return linked accounts associated with the given origin ID.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginIds:** `*string` — Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
    
</dd>
</dl>

<dl>
<dd>

**includeDuplicates:** `*bool` — If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` — If provided, will only return linked accounts associated with the given integration name.
    
</dd>
</dl>

<dl>
<dd>

**isTestAccount:** `*string` — If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Filter by status. Options: `COMPLETE`, `IDLE`, `INCOMPLETE`, `RELINK_NEEDED`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Passthrough
<details><summary><code>client.Accounting.Passthrough.Create(request) -> *accounting.RemoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.DataPassthroughRequest{
        Method: accounting.MethodEnumGet,
        Path: "/scooters",
    }
client.Accounting.Passthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*accounting.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting PaymentMethods
<details><summary><code>client.Accounting.PaymentMethods.List() -> *accounting.PaginatedPaymentMethodList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `PaymentMethod` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentMethodsListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.PaymentMethods.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.PaymentMethods.Retrieve(Id) -> *accounting.PaymentMethod</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `PaymentMethod` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentMethodsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.PaymentMethods.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting PaymentTerms
<details><summary><code>client.Accounting.PaymentTerms.List() -> *accounting.PaginatedPaymentTermList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `PaymentTerm` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentTermsListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.PaymentTerms.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.PaymentTerms.Retrieve(Id) -> *accounting.PaymentTerm</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `PaymentTerm` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentTermsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.PaymentTerms.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Payments
<details><summary><code>client.Accounting.Payments.List() -> *accounting.PaginatedPaymentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Payment` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentsListRequest{
        AccountId: merge.String(
            "account_id",
        ),
        CompanyId: merge.String(
            "company_id",
        ),
        ContactId: merge.String(
            "contact_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        TransactionDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TransactionDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Accounting.Payments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `*string` — If provided, will only return payments for this account.
    
</dd>
</dl>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return payments for this company.
    
</dd>
</dl>

<dl>
<dd>

**contactId:** `*string` — If provided, will only return payments for this contact.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.PaymentsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Payments.Create(request) -> *accounting.PaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Payment` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.PaymentRequest{},
    }
client.Accounting.Payments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.PaymentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Payments.Retrieve(Id) -> *accounting.Payment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Payment` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Payments.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.PaymentsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Payments.PartialUpdate(Id, request) -> *accounting.PaymentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a `Payment` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PatchedPaymentEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.PatchedPaymentRequest{},
    }
client.Accounting.Payments.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.PatchedPaymentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Payments.LineItemsRemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentsLineItemsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Payments.LineItemsRemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Payments.MetaPatchRetrieve(Id) -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Payment` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Payments.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Payments.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Payment` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.Payments.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Payments.RemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PaymentsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Payments.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting PhoneNumbers
<details><summary><code>client.Accounting.PhoneNumbers.Retrieve(Id) -> *accounting.AccountingPhoneNumber</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `AccountingPhoneNumber` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PhoneNumbersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.PhoneNumbers.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Projects
<details><summary><code>client.Accounting.Projects.List() -> *accounting.PaginatedProjectList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Project` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ProjectsListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.Projects.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ProjectsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Projects.Retrieve(Id) -> *accounting.Project</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Project` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.ProjectsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Projects.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.ProjectsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting PurchaseOrders
<details><summary><code>client.Accounting.PurchaseOrders.List() -> *accounting.PaginatedPurchaseOrderList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `PurchaseOrder` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PurchaseOrdersListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IssueDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IssueDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.PurchaseOrders.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return purchase orders for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.PurchaseOrdersListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**issueDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**issueDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.PurchaseOrders.Create(request) -> *accounting.PurchaseOrderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `PurchaseOrder` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PurchaseOrderEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.PurchaseOrderRequest{},
    }
client.Accounting.PurchaseOrders.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.PurchaseOrderRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.PurchaseOrders.Retrieve(Id) -> *accounting.PurchaseOrder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `PurchaseOrder` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PurchaseOrdersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.PurchaseOrders.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.PurchaseOrdersRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.PurchaseOrders.LineItemsRemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PurchaseOrdersLineItemsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.PurchaseOrders.LineItemsRemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.PurchaseOrders.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `PurchaseOrder` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.PurchaseOrders.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.PurchaseOrders.RemoteFieldClassesList() -> *accounting.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.PurchaseOrdersRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.PurchaseOrders.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting RegenerateKey
<details><summary><code>client.Accounting.RegenerateKey.Create(request) -> *accounting.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Exchange remote keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.RemoteKeyForRegenerationRequest{
        Name: "Remote Deployment Key 1",
    }
client.Accounting.RegenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting SyncStatus
<details><summary><code>client.Accounting.SyncStatus.List() -> *accounting.PaginatedSyncStatusList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sync status for the current sync and the most recently finished sync. `last_sync_start` represents the most recent time any sync began. `last_sync_finished` represents the most recent time any sync completed. These timestamps may correspond to different sync instances which may result in a sync start time being later than a separate sync completed time. To ensure you are retrieving the latest available data reference the `last_sync_finished` timestamp where `last_sync_result` is `DONE`. Possible values for `status` and `last_sync_result` are `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`. Learn more about sync status in our [Help Center](https://help.merge.dev/en/articles/8184193-merge-sync-statuses).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.SyncStatusListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Accounting.SyncStatus.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting ForceResync
<details><summary><code>client.Accounting.ForceResync.SyncStatusResyncCreate() -> []*accounting.SyncStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Force re-sync of all models. This endpoint is available for monthly, quarterly, and highest sync frequency customers on the Professional or Enterprise plans. Doing so will consume a sync credit for the relevant linked account. Force re-syncs can also be triggered manually in the Merge Dashboard and is available for all customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.ForceResync.SyncStatusResyncCreate(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting TaxRates
<details><summary><code>client.Accounting.TaxRates.List() -> *accounting.PaginatedTaxRateList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `TaxRate` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.TaxRatesListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Accounting.TaxRates.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return tax rates for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return TaxRates with this name.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.TaxRates.Retrieve(Id) -> *accounting.TaxRate</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `TaxRate` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.TaxRatesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.TaxRates.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting TrackingCategories
<details><summary><code>client.Accounting.TrackingCategories.List() -> *accounting.PaginatedTrackingCategoryList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `TrackingCategory` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.TrackingCategoriesListRequest{
        CategoryType: accounting.TrackingCategoriesListRequestCategoryTypeEmpty.Ptr(),
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Status: accounting.TrackingCategoriesListRequestStatusEmpty.Ptr(),
    }
client.Accounting.TrackingCategories.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**categoryType:** `*accounting.TrackingCategoriesListRequestCategoryType` — If provided, will only return tracking categories with this type.
    
</dd>
</dl>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return tracking categories for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return tracking categories with this name.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*accounting.TrackingCategoriesListRequestStatus` — If provided, will only return tracking categories with this status.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.TrackingCategories.Retrieve(Id) -> *accounting.TrackingCategory</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `TrackingCategory` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.TrackingCategoriesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.TrackingCategories.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting Transactions
<details><summary><code>client.Accounting.Transactions.List() -> *accounting.PaginatedTransactionList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Transaction` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.TransactionsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        TransactionDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TransactionDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Accounting.Transactions.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return accounting transactions for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.TransactionsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.Transactions.Retrieve(Id) -> *accounting.Transaction</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Transaction` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.TransactionsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.Transactions.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.TransactionsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting VendorCredits
<details><summary><code>client.Accounting.VendorCredits.List() -> *accounting.PaginatedVendorCreditList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `VendorCredit` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.VendorCreditsListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        TransactionDateAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TransactionDateBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Accounting.VendorCredits.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return vendor credits for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.VendorCreditsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**transactionDateBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.VendorCredits.Create(request) -> *accounting.VendorCreditResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `VendorCredit` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.VendorCreditEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &accounting.VendorCreditRequest{},
    }
client.Accounting.VendorCredits.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*accounting.VendorCreditRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.VendorCredits.Retrieve(Id) -> *accounting.VendorCredit</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `VendorCredit` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.VendorCreditsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Accounting.VendorCredits.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*accounting.VendorCreditsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.VendorCredits.MetaPostRetrieve() -> *accounting.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `VendorCredit` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.VendorCredits.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Accounting WebhookReceivers
<details><summary><code>client.Accounting.WebhookReceivers.List() -> []*accounting.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `WebhookReceiver` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Accounting.WebhookReceivers.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Accounting.WebhookReceivers.Create(request) -> *accounting.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `WebhookReceiver` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &accounting.WebhookReceiverRequest{
        Event: "event",
        IsActive: true,
    }
client.Accounting.WebhookReceivers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**event:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**key:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm AccountDetails
<details><summary><code>client.Crm.AccountDetails.Retrieve() -> *crm.AccountDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details for a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.AccountDetails.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm AccountToken
<details><summary><code>client.Crm.AccountToken.Retrieve(PublicToken) -> *crm.AccountToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account token for the end user with the provided public token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.AccountToken.Retrieve(
        context.TODO(),
        "public_token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**publicToken:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Accounts
<details><summary><code>client.Crm.Accounts.List() -> *crm.PaginatedAccountList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Account` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.AccountsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        OwnerId: merge.String(
            "owner_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Accounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return accounts with this name.
    
</dd>
</dl>

<dl>
<dd>

**ownerId:** `*string` — If provided, will only return accounts with this owner.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Accounts.Create(request) -> *crm.CrmAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Account` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CrmAccountEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.AccountRequest{},
    }
client.Crm.Accounts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.AccountRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Accounts.Retrieve(Id) -> *crm.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Account` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.AccountsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Accounts.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Accounts.PartialUpdate(Id, request) -> *crm.CrmAccountResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an `Account` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.PatchedCrmAccountEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.PatchedAccountRequest{},
    }
client.Crm.Accounts.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.PatchedAccountRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Accounts.MetaPatchRetrieve(Id) -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `CRMAccount` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Accounts.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Accounts.MetaPostRetrieve() -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `CRMAccount` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Accounts.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Accounts.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.AccountsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Accounts.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm AsyncPassthrough
<details><summary><code>client.Crm.AsyncPassthrough.Create(request) -> *crm.AsyncPassthroughReciept</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Asynchronously pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.DataPassthroughRequest{
        Method: crm.MethodEnumGet,
        Path: "/scooters",
    }
client.Crm.AsyncPassthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*crm.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.AsyncPassthrough.Retrieve(AsyncPassthroughReceiptId) -> *crm.AsyncPassthroughRetrieveResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves data from earlier async-passthrough POST request
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.AsyncPassthrough.Retrieve(
        context.TODO(),
        "async_passthrough_receipt_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**asyncPassthroughReceiptId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm AuditTrail
<details><summary><code>client.Crm.AuditTrail.List() -> *crm.PaginatedAuditLogEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a list of audit trail events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.AuditTrailListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EventType: merge.String(
            "event_type",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        UserEmail: merge.String(
            "user_email",
        ),
    }
client.Crm.AuditTrail.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include audit trail events that occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*string` — If included, will only include events with the given event type. Possible values include: `CREATED_REMOTE_PRODUCTION_API_KEY`, `DELETED_REMOTE_PRODUCTION_API_KEY`, `CREATED_TEST_API_KEY`, `DELETED_TEST_API_KEY`, `REGENERATED_PRODUCTION_API_KEY`, `REGENERATED_WEBHOOK_SIGNATURE`, `INVITED_USER`, `TWO_FACTOR_AUTH_ENABLED`, `TWO_FACTOR_AUTH_DISABLED`, `DELETED_LINKED_ACCOUNT`, `DELETED_ALL_COMMON_MODELS_FOR_LINKED_ACCOUNT`, `CREATED_DESTINATION`, `DELETED_DESTINATION`, `CHANGED_DESTINATION`, `CHANGED_SCOPES`, `CHANGED_PERSONAL_INFORMATION`, `CHANGED_ORGANIZATION_SETTINGS`, `ENABLED_INTEGRATION`, `DISABLED_INTEGRATION`, `ENABLED_CATEGORY`, `DISABLED_CATEGORY`, `CHANGED_PASSWORD`, `RESET_PASSWORD`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `CREATED_INTEGRATION_WIDE_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_FIELD_MAPPING`, `CHANGED_INTEGRATION_WIDE_FIELD_MAPPING`, `CHANGED_LINKED_ACCOUNT_FIELD_MAPPING`, `DELETED_INTEGRATION_WIDE_FIELD_MAPPING`, `DELETED_LINKED_ACCOUNT_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `CHANGED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `DELETED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `FORCED_LINKED_ACCOUNT_RESYNC`, `MUTED_ISSUE`, `GENERATED_MAGIC_LINK`, `ENABLED_MERGE_WEBHOOK`, `DISABLED_MERGE_WEBHOOK`, `MERGE_WEBHOOK_TARGET_CHANGED`, `END_USER_CREDENTIALS_ACCESSED`
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include audit trail events that occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**userEmail:** `*string` — If provided, this will return events associated with the specified user email. Please note that the email address reflects the user's email at the time of the event, and may not be their current email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm AvailableActions
<details><summary><code>client.Crm.AvailableActions.Retrieve() -> *crm.AvailableActions</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of models and actions available for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.AvailableActions.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Contacts
<details><summary><code>client.Crm.Contacts.List() -> *crm.PaginatedContactList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Contact` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.ContactsListRequest{
        AccountId: merge.String(
            "account_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmailAddresses: merge.String(
            "email_addresses",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        PhoneNumbers: merge.String(
            "phone_numbers",
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Contacts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `*string` — If provided, will only return contacts with this account.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**emailAddresses:** `*string` — If provided, will only return contacts matching the email addresses; multiple email_addresses can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.ContactsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**phoneNumbers:** `*string` — If provided, will only return contacts matching the phone numbers; multiple phone numbers can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Contacts.Create(request) -> *crm.CrmContactResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Contact` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CrmContactEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.ContactRequest{},
    }
client.Crm.Contacts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.ContactRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Contacts.Retrieve(Id) -> *crm.Contact</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Contact` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.ContactsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Contacts.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.ContactsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Contacts.PartialUpdate(Id, request) -> *crm.CrmContactResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a `Contact` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.PatchedCrmContactEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.PatchedContactRequest{},
    }
client.Crm.Contacts.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.PatchedContactRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Contacts.IgnoreCreate(ModelId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.IgnoreCommonModelRequest{
        Reason: &crm.IgnoreCommonModelRequestReason{
            ReasonEnum: crm.ReasonEnumGeneralCustomerRequest,
        },
    }
client.Crm.Contacts.IgnoreCreate(
        context.TODO(),
        "model_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**modelId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*crm.IgnoreCommonModelRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Contacts.MetaPatchRetrieve(Id) -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `CRMContact` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Contacts.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Contacts.MetaPostRetrieve() -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `CRMContact` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Contacts.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Contacts.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.ContactsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Contacts.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm CustomObjectClasses
<details><summary><code>client.Crm.CustomObjectClasses.List() -> *crm.PaginatedCustomObjectClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `CustomObjectClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.CustomObjectClasses.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.CustomObjectClasses.Retrieve(Id) -> *crm.CustomObjectClass</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `CustomObjectClass` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.CustomObjectClasses.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm AssociationTypes
<details><summary><code>client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesList(CustomObjectClassId) -> *crm.PaginatedAssociationTypeList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `AssociationType` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesAssociationTypesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesList(
        context.TODO(),
        "custom_object_class_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesCreate(CustomObjectClassId, request) -> *crm.CrmAssociationTypeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `AssociationType` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CrmAssociationTypeEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.AssociationTypeRequestRequest{
            SourceObjectClass: &crm.ObjectClassDescriptionRequest{
                Id: "id",
                OriginType: crm.OriginTypeEnumCustomObject,
            },
            TargetObjectClasses: []*crm.ObjectClassDescriptionRequest{
                &crm.ObjectClassDescriptionRequest{
                    Id: "id",
                    OriginType: crm.OriginTypeEnumCustomObject,
                },
            },
            RemoteKeyName: "remote_key_name",
        },
    }
client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesCreate(
        context.TODO(),
        "custom_object_class_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.AssociationTypeRequestRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesRetrieve(CustomObjectClassId, Id) -> *crm.AssociationType</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `AssociationType` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesAssociationTypesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesRetrieve(
        context.TODO(),
        "custom_object_class_id",
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesMetaPostRetrieve(CustomObjectClassId) -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `CRMAssociationType` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.AssociationTypes.CustomObjectClassesAssociationTypesMetaPostRetrieve(
        context.TODO(),
        "custom_object_class_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm CustomObjects
<details><summary><code>client.Crm.CustomObjects.CustomObjectClassesCustomObjectsList(CustomObjectClassId) -> *crm.PaginatedCustomObjectList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `CustomObject` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesCustomObjectsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.CustomObjects.CustomObjectClassesCustomObjectsList(
        context.TODO(),
        "custom_object_class_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.CustomObjects.CustomObjectClassesCustomObjectsCreate(CustomObjectClassId, request) -> *crm.CrmCustomObjectResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `CustomObject` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CrmCustomObjectEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.CustomObjectRequest{
            Fields: map[string]any{
                "test_field": "hello",
            },
        },
    }
client.Crm.CustomObjects.CustomObjectClassesCustomObjectsCreate(
        context.TODO(),
        "custom_object_class_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.CustomObjectRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.CustomObjects.CustomObjectClassesCustomObjectsRetrieve(CustomObjectClassId, Id) -> *crm.CustomObject</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `CustomObject` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesCustomObjectsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.CustomObjects.CustomObjectClassesCustomObjectsRetrieve(
        context.TODO(),
        "custom_object_class_id",
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.CustomObjects.CustomObjectClassesCustomObjectsMetaPostRetrieve(CustomObjectClassId) -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `CRMCustomObject` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.CustomObjects.CustomObjectClassesCustomObjectsMetaPostRetrieve(
        context.TODO(),
        "custom_object_class_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.CustomObjects.CustomObjectClassesCustomObjectsRemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesCustomObjectsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.CustomObjects.CustomObjectClassesCustomObjectsRemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Associations
<details><summary><code>client.Crm.Associations.CustomObjectClassesCustomObjectsAssociationsList(CustomObjectClassId, ObjectId) -> *crm.PaginatedAssociationList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Association` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesCustomObjectsAssociationsListRequest{
        AssociationTypeId: merge.String(
            "association_type_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Associations.CustomObjectClassesCustomObjectsAssociationsList(
        context.TODO(),
        "custom_object_class_id",
        "object_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**customObjectClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**objectId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**associationTypeId:** `*string` — If provided, will only return opportunities with this association_type.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Associations.CustomObjectClassesCustomObjectsAssociationsUpdate(SourceClassId, SourceObjectId, TargetClassId, TargetObjectId, AssociationTypeId) -> *crm.Association</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an Association between `source_object_id` and `target_object_id` of type `association_type_id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CustomObjectClassesCustomObjectsAssociationsUpdateRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
    }
client.Crm.Associations.CustomObjectClassesCustomObjectsAssociationsUpdate(
        context.TODO(),
        "source_class_id",
        "source_object_id",
        "target_class_id",
        "target_object_id",
        "association_type_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**sourceClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**sourceObjectId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**targetClassId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**targetObjectId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**associationTypeId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Scopes
<details><summary><code>client.Crm.Scopes.DefaultScopesRetrieve() -> *crm.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the default permissions for Merge Common Models and fields across all Linked Accounts of a given category. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Scopes.DefaultScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Scopes.LinkedAccountScopesRetrieve() -> *crm.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all available permissions for Merge Common Models and fields for a single Linked Account. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Scopes.LinkedAccountScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Scopes.LinkedAccountScopesCreate(request) -> *crm.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update permissions for any Common Model or field for a single Linked Account. Any Scopes not set in this POST request will inherit the default Scopes. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.LinkedAccountCommonModelScopeDeserializerRequest{
        CommonModels: []*crm.IndividualCommonModelScopeDeserializerRequest{
            &crm.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Employee",
                ModelPermissions: map[string]*crm.ModelPermissionDeserializerRequest{
                    "READ": &crm.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            true,
                        ),
                    },
                    "WRITE": &crm.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
                FieldPermissions: &crm.FieldPermissionDeserializerRequest{
                    EnabledFields: []any{
                        "avatar",
                        "home_location",
                    },
                    DisabledFields: []any{
                        "work_location",
                    },
                },
            },
            &crm.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Benefit",
                ModelPermissions: map[string]*crm.ModelPermissionDeserializerRequest{
                    "WRITE": &crm.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
            },
        },
    }
client.Crm.Scopes.LinkedAccountScopesCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `[]*crm.IndividualCommonModelScopeDeserializerRequest` — The common models you want to update the scopes for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm DeleteAccount
<details><summary><code>client.Crm.DeleteAccount.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.DeleteAccount.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm EngagementTypes
<details><summary><code>client.Crm.EngagementTypes.List() -> *crm.PaginatedEngagementTypeList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `EngagementType` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EngagementTypesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.EngagementTypes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.EngagementTypes.Retrieve(Id) -> *crm.EngagementType</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `EngagementType` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EngagementTypesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.EngagementTypes.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.EngagementTypes.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EngagementTypesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.EngagementTypes.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Engagements
<details><summary><code>client.Crm.Engagements.List() -> *crm.PaginatedEngagementList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Engagement` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EngagementsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        StartedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        StartedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Crm.Engagements.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.EngagementsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**startedAfter:** `*time.Time` — If provided, will only return engagements started after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**startedBefore:** `*time.Time` — If provided, will only return engagements started before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Engagements.Create(request) -> *crm.EngagementResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Engagement` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EngagementEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.EngagementRequest{},
    }
client.Crm.Engagements.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.EngagementRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Engagements.Retrieve(Id) -> *crm.Engagement</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Engagement` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EngagementsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Engagements.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.EngagementsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Engagements.PartialUpdate(Id, request) -> *crm.EngagementResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an `Engagement` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.PatchedEngagementEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.PatchedEngagementRequest{},
    }
client.Crm.Engagements.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.PatchedEngagementRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Engagements.MetaPatchRetrieve(Id) -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Engagement` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Engagements.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Engagements.MetaPostRetrieve() -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Engagement` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Engagements.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Engagements.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EngagementsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Engagements.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm FieldMapping
<details><summary><code>client.Crm.FieldMapping.FieldMappingsRetrieve() -> *crm.FieldMappingApiInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all Field Mappings for this Linked Account. Field Mappings are mappings between third-party Remote Fields and user defined Merge fields. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.FieldMappingsRetrieveRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
    }
client.Crm.FieldMapping.FieldMappingsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.FieldMapping.FieldMappingsCreate(request) -> *crm.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create new Field Mappings that will be available after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.CreateFieldMappingRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
        TargetFieldName: "example_target_field_name",
        TargetFieldDescription: "this is a example description of the target field",
        RemoteFieldTraversalPath: []any{
            "example_remote_field",
        },
        RemoteMethod: "GET",
        RemoteUrlPath: "/example-url-path",
        CommonModelName: "ExampleCommonModel",
    }
client.Crm.FieldMapping.FieldMappingsCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldName:** `string` — The name of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldDescription:** `string` — The description of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**commonModelName:** `string` — The name of the Common Model that the remote field corresponds to in a given category.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.FieldMapping.FieldMappingsDestroy(FieldMappingId) -> *crm.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes Field Mappings for a Linked Account. All data related to this Field Mapping will be deleted and these changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.FieldMapping.FieldMappingsDestroy(
        context.TODO(),
        "field_mapping_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.FieldMapping.FieldMappingsPartialUpdate(FieldMappingId, request) -> *crm.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update existing Field Mappings for a Linked Account. Changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.PatchedEditFieldMappingRequest{}
client.Crm.FieldMapping.FieldMappingsPartialUpdate(
        context.TODO(),
        "field_mapping_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `*string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `*string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.FieldMapping.RemoteFieldsRetrieve() -> *crm.RemoteFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all remote fields for a Linked Account. Remote fields are third-party fields that are accessible after initial sync if remote_data is enabled. You can use remote fields to override existing Merge fields or map a new Merge field. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.RemoteFieldsRetrieveRequest{
        CommonModels: merge.String(
            "common_models",
        ),
        IncludeExampleValues: merge.String(
            "include_example_values",
        ),
    }
client.Crm.FieldMapping.RemoteFieldsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `*string` — A comma seperated list of Common Model names. If included, will only return Remote Fields for those Common Models.
    
</dd>
</dl>

<dl>
<dd>

**includeExampleValues:** `*string` — If true, will include example values, where available, for remote fields in the 3rd party platform. These examples come from active data from your customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.FieldMapping.TargetFieldsRetrieve() -> *crm.ExternalTargetFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all organization-wide Target Fields, this will not include any Linked Account specific Target Fields. Organization-wide Target Fields are additional fields appended to the Merge Common Model for all Linked Accounts in a category. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/target-fields/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.FieldMapping.TargetFieldsRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm GenerateKey
<details><summary><code>client.Crm.GenerateKey.Create(request) -> *crm.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a remote key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.GenerateRemoteKeyRequest{
        Name: "Remote Deployment Key 1",
    }
client.Crm.GenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Issues
<details><summary><code>client.Crm.Issues.List() -> *crm.PaginatedIssueList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets all issues for Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.IssuesListRequest{
        AccountToken: merge.String(
            "account_token",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        FirstIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        FirstIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeMuted: merge.String(
            "include_muted",
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        LastIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LastIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LinkedAccountId: merge.String(
            "linked_account_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        Status: crm.IssuesListRequestStatusOngoing.Ptr(),
    }
client.Crm.Issues.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include issues whose most recent action occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose first incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose first incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**includeMuted:** `*string` — If true, will include muted issues
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose last incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose last incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**linkedAccountId:** `*string` — If provided, will only include issues pertaining to the linked account passed in.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include issues whose most recent action occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**status:** `*crm.IssuesListRequestStatus` 

Status of the issue. Options: ('ONGOING', 'RESOLVED')

* `ONGOING` - ONGOING
* `RESOLVED` - RESOLVED
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Issues.Retrieve(Id) -> *crm.Issue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific issue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Issues.Retrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Leads
<details><summary><code>client.Crm.Leads.List() -> *crm.PaginatedLeadList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Lead` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.LeadsListRequest{
        ConvertedAccountId: merge.String(
            "converted_account_id",
        ),
        ConvertedContactId: merge.String(
            "converted_contact_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmailAddresses: merge.String(
            "email_addresses",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OwnerId: merge.String(
            "owner_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        PhoneNumbers: merge.String(
            "phone_numbers",
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Leads.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**convertedAccountId:** `*string` — If provided, will only return leads with this account.
    
</dd>
</dl>

<dl>
<dd>

**convertedContactId:** `*string` — If provided, will only return leads with this contact.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**emailAddresses:** `*string` — If provided, will only return contacts matching the email addresses; multiple email_addresses can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.LeadsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**ownerId:** `*string` — If provided, will only return leads with this owner.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**phoneNumbers:** `*string` — If provided, will only return contacts matching the phone numbers; multiple phone numbers can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Leads.Create(request) -> *crm.LeadResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Lead` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.LeadEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.LeadRequest{},
    }
client.Crm.Leads.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.LeadRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Leads.Retrieve(Id) -> *crm.Lead</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Lead` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.LeadsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Leads.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.LeadsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Leads.MetaPostRetrieve() -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Lead` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Leads.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Leads.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.LeadsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Leads.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm LinkToken
<details><summary><code>client.Crm.LinkToken.Create(request) -> *crm.LinkToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a link token to be used when linking a new end user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.EndUserDetailsRequest{
        EndUserEmailAddress: "example@gmail.com",
        EndUserOrganizationName: "Test Organization",
        EndUserOriginId: "12345",
        Categories: []crm.CategoriesEnum{
            crm.CategoriesEnumHris,
            crm.CategoriesEnumAts,
        },
    }
client.Crm.LinkToken.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**endUserEmailAddress:** `string` — Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `string` — Your end user's organization.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `string` — This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
    
</dd>
</dl>

<dl>
<dd>

**categories:** `[]*crm.CategoriesEnum` — The integration categories to show in Merge Link.
    
</dd>
</dl>

<dl>
<dd>

**integration:** `*string` — The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://docs.merge.dev/guides/merge-link/single-integration/.
    
</dd>
</dl>

<dl>
<dd>

**linkExpiryMins:** `*int` — An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
    
</dd>
</dl>

<dl>
<dd>

**shouldCreateMagicLinkUrl:** `*bool` — Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**hideAdminMagicLink:** `*bool` — Whether to generate a Magic Link URL on the Admin Needed screen during the linking flow. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**commonModels:** `[]*crm.CommonModelScopesBodyRequest` — An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
    
</dd>
</dl>

<dl>
<dd>

**categoryCommonModelScopes:** `map[string][]*crm.IndividualCommonModelScopeDeserializerRequest` — When creating a Link Token, you can set permissions for Common Models that will apply to the account that is going to be linked. Any model or field not specified in link token payload will default to existing settings.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*crm.EndUserDetailsRequestLanguage` 

The following subset of IETF language tags can be used to configure localization.

* `en` - en
* `de` - de
    
</dd>
</dl>

<dl>
<dd>

**areSyncsDisabled:** `*bool` — The boolean that indicates whether initial, periodic, and force syncs will be disabled.
    
</dd>
</dl>

<dl>
<dd>

**integrationSpecificConfig:** `map[string]any` — A JSON object containing integration-specific configuration options.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm LinkedAccounts
<details><summary><code>client.Crm.LinkedAccounts.List() -> *crm.PaginatedAccountDetailsAndActionsList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List linked accounts for your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.LinkedAccountsListRequest{
        Category: crm.LinkedAccountsListRequestCategoryAccounting.Ptr(),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndUserEmailAddress: merge.String(
            "end_user_email_address",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        EndUserOriginId: merge.String(
            "end_user_origin_id",
        ),
        EndUserOriginIds: merge.String(
            "end_user_origin_ids",
        ),
        Id: merge.String(
            "id",
        ),
        Ids: merge.String(
            "ids",
        ),
        IncludeDuplicates: merge.Bool(
            true,
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        IsTestAccount: merge.String(
            "is_test_account",
        ),
        PageSize: merge.Int(
            1,
        ),
        Status: merge.String(
            "status",
        ),
    }
client.Crm.LinkedAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**category:** `*crm.LinkedAccountsListRequestCategory` 

Options: `accounting`, `ats`, `crm`, `filestorage`, `hris`, `mktg`, `ticketing`

* `hris` - hris
* `ats` - ats
* `accounting` - accounting
* `ticketing` - ticketing
* `crm` - crm
* `mktg` - mktg
* `filestorage` - filestorage
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endUserEmailAddress:** `*string` — If provided, will only return linked accounts associated with the given email address.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` — If provided, will only return linked accounts associated with the given organization name.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `*string` — If provided, will only return linked accounts associated with the given origin ID.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginIds:** `*string` — Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
    
</dd>
</dl>

<dl>
<dd>

**includeDuplicates:** `*bool` — If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` — If provided, will only return linked accounts associated with the given integration name.
    
</dd>
</dl>

<dl>
<dd>

**isTestAccount:** `*string` — If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Filter by status. Options: `COMPLETE`, `IDLE`, `INCOMPLETE`, `RELINK_NEEDED`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Notes
<details><summary><code>client.Crm.Notes.List() -> *crm.PaginatedNoteList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Note` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.NotesListRequest{
        AccountId: merge.String(
            "account_id",
        ),
        ContactId: merge.String(
            "contact_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OpportunityId: merge.String(
            "opportunity_id",
        ),
        OwnerId: merge.String(
            "owner_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Notes.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `*string` — If provided, will only return notes with this account.
    
</dd>
</dl>

<dl>
<dd>

**contactId:** `*string` — If provided, will only return notes with this contact.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.NotesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**opportunityId:** `*string` — If provided, will only return notes with this opportunity.
    
</dd>
</dl>

<dl>
<dd>

**ownerId:** `*string` — If provided, will only return notes with this owner.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Notes.Create(request) -> *crm.NoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Note` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.NoteEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.NoteRequest{},
    }
client.Crm.Notes.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.NoteRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Notes.Retrieve(Id) -> *crm.Note</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Note` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.NotesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Notes.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.NotesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Notes.MetaPostRetrieve() -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Note` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Notes.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Notes.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.NotesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Notes.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Opportunities
<details><summary><code>client.Crm.Opportunities.List() -> *crm.PaginatedOpportunityList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Opportunity` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.OpportunitiesListRequest{
        AccountId: merge.String(
            "account_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OwnerId: merge.String(
            "owner_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteCreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        StageId: merge.String(
            "stage_id",
        ),
        Status: crm.OpportunitiesListRequestStatusLost.Ptr(),
    }
client.Crm.Opportunities.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `*string` — If provided, will only return opportunities with this account.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.OpportunitiesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**ownerId:** `*string` — If provided, will only return opportunities with this owner.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteCreatedAfter:** `*time.Time` — If provided, will only return opportunities created in the third party platform after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**stageId:** `*string` — If provided, will only return opportunities with this stage.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*crm.OpportunitiesListRequestStatus` 

If provided, will only return opportunities with this status. Options: ('OPEN', 'WON', 'LOST')

* `OPEN` - OPEN
* `WON` - WON
* `LOST` - LOST
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Opportunities.Create(request) -> *crm.OpportunityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Opportunity` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.OpportunityEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.OpportunityRequest{},
    }
client.Crm.Opportunities.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.OpportunityRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Opportunities.Retrieve(Id) -> *crm.Opportunity</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Opportunity` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.OpportunitiesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Opportunities.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.OpportunitiesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Opportunities.PartialUpdate(Id, request) -> *crm.OpportunityResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates an `Opportunity` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.PatchedOpportunityEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.PatchedOpportunityRequest{},
    }
client.Crm.Opportunities.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.PatchedOpportunityRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Opportunities.MetaPatchRetrieve(Id) -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Opportunity` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Opportunities.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Opportunities.MetaPostRetrieve() -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Opportunity` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Opportunities.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Opportunities.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.OpportunitiesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Opportunities.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Passthrough
<details><summary><code>client.Crm.Passthrough.Create(request) -> *crm.RemoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.DataPassthroughRequest{
        Method: crm.MethodEnumGet,
        Path: "/scooters",
    }
client.Crm.Passthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*crm.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm RegenerateKey
<details><summary><code>client.Crm.RegenerateKey.Create(request) -> *crm.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Exchange remote keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.RemoteKeyForRegenerationRequest{
        Name: "Remote Deployment Key 1",
    }
client.Crm.RegenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Stages
<details><summary><code>client.Crm.Stages.List() -> *crm.PaginatedStageList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Stage` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.StagesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Stages.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Stages.Retrieve(Id) -> *crm.Stage</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Stage` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.StagesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Stages.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Stages.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.StagesRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Stages.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm SyncStatus
<details><summary><code>client.Crm.SyncStatus.List() -> *crm.PaginatedSyncStatusList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sync status for the current sync and the most recently finished sync. `last_sync_start` represents the most recent time any sync began. `last_sync_finished` represents the most recent time any sync completed. These timestamps may correspond to different sync instances which may result in a sync start time being later than a separate sync completed time. To ensure you are retrieving the latest available data reference the `last_sync_finished` timestamp where `last_sync_result` is `DONE`. Possible values for `status` and `last_sync_result` are `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`. Learn more about sync status in our [Help Center](https://help.merge.dev/en/articles/8184193-merge-sync-statuses).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.SyncStatusListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.SyncStatus.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm ForceResync
<details><summary><code>client.Crm.ForceResync.SyncStatusResyncCreate() -> []*crm.SyncStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Force re-sync of all models. This endpoint is available for monthly, quarterly, and highest sync frequency customers on the Professional or Enterprise plans. Doing so will consume a sync credit for the relevant linked account. Force re-syncs can also be triggered manually in the Merge Dashboard and is available for all customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.ForceResync.SyncStatusResyncCreate(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Tasks
<details><summary><code>client.Crm.Tasks.List() -> *crm.PaginatedTaskList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Task` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.TasksListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Tasks.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.TasksListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Tasks.Create(request) -> *crm.TaskResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Task` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.TaskEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.TaskRequest{},
    }
client.Crm.Tasks.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.TaskRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Tasks.Retrieve(Id) -> *crm.Task</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Task` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.TasksRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Tasks.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*crm.TasksRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Tasks.PartialUpdate(Id, request) -> *crm.TaskResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a `Task` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.PatchedTaskEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &crm.PatchedTaskRequest{},
    }
client.Crm.Tasks.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*crm.PatchedTaskRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Tasks.MetaPatchRetrieve(Id) -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Task` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Tasks.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Tasks.MetaPostRetrieve() -> *crm.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Task` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.Tasks.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Tasks.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.TasksRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Tasks.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm Users
<details><summary><code>client.Crm.Users.List() -> *crm.PaginatedUserList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `User` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.UsersListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        Email: merge.String(
            "email",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Crm.Users.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**email:** `*string` — If provided, will only return users with this email.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Users.Retrieve(Id) -> *crm.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `User` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.UsersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Crm.Users.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Users.IgnoreCreate(ModelId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.IgnoreCommonModelRequest{
        Reason: &crm.IgnoreCommonModelRequestReason{
            ReasonEnum: crm.ReasonEnumGeneralCustomerRequest,
        },
    }
client.Crm.Users.IgnoreCreate(
        context.TODO(),
        "model_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**modelId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*crm.IgnoreCommonModelRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.Users.RemoteFieldClassesList() -> *crm.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.UsersRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Crm.Users.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Crm WebhookReceivers
<details><summary><code>client.Crm.WebhookReceivers.List() -> []*crm.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `WebhookReceiver` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Crm.WebhookReceivers.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Crm.WebhookReceivers.Create(request) -> *crm.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `WebhookReceiver` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &crm.WebhookReceiverRequest{
        Event: "event",
        IsActive: true,
    }
client.Crm.WebhookReceivers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**event:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**key:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage AccountDetails
<details><summary><code>client.FileStorage.AccountDetails.Retrieve() -> *filestorage.AccountDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details for a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.AccountDetails.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage AccountToken
<details><summary><code>client.FileStorage.AccountToken.Retrieve(PublicToken) -> *filestorage.AccountToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account token for the end user with the provided public token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.AccountToken.Retrieve(
        context.TODO(),
        "public_token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**publicToken:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage AsyncPassthrough
<details><summary><code>client.FileStorage.AsyncPassthrough.Create(request) -> *filestorage.AsyncPassthroughReciept</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Asynchronously pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.DataPassthroughRequest{
        Method: filestorage.MethodEnumGet,
        Path: "/scooters",
    }
client.FileStorage.AsyncPassthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*filestorage.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.AsyncPassthrough.Retrieve(AsyncPassthroughReceiptId) -> *filestorage.AsyncPassthroughRetrieveResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves data from earlier async-passthrough POST request
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.AsyncPassthrough.Retrieve(
        context.TODO(),
        "async_passthrough_receipt_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**asyncPassthroughReceiptId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage AuditTrail
<details><summary><code>client.FileStorage.AuditTrail.List() -> *filestorage.PaginatedAuditLogEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a list of audit trail events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.AuditTrailListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EventType: merge.String(
            "event_type",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        UserEmail: merge.String(
            "user_email",
        ),
    }
client.FileStorage.AuditTrail.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include audit trail events that occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*string` — If included, will only include events with the given event type. Possible values include: `CREATED_REMOTE_PRODUCTION_API_KEY`, `DELETED_REMOTE_PRODUCTION_API_KEY`, `CREATED_TEST_API_KEY`, `DELETED_TEST_API_KEY`, `REGENERATED_PRODUCTION_API_KEY`, `REGENERATED_WEBHOOK_SIGNATURE`, `INVITED_USER`, `TWO_FACTOR_AUTH_ENABLED`, `TWO_FACTOR_AUTH_DISABLED`, `DELETED_LINKED_ACCOUNT`, `DELETED_ALL_COMMON_MODELS_FOR_LINKED_ACCOUNT`, `CREATED_DESTINATION`, `DELETED_DESTINATION`, `CHANGED_DESTINATION`, `CHANGED_SCOPES`, `CHANGED_PERSONAL_INFORMATION`, `CHANGED_ORGANIZATION_SETTINGS`, `ENABLED_INTEGRATION`, `DISABLED_INTEGRATION`, `ENABLED_CATEGORY`, `DISABLED_CATEGORY`, `CHANGED_PASSWORD`, `RESET_PASSWORD`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `CREATED_INTEGRATION_WIDE_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_FIELD_MAPPING`, `CHANGED_INTEGRATION_WIDE_FIELD_MAPPING`, `CHANGED_LINKED_ACCOUNT_FIELD_MAPPING`, `DELETED_INTEGRATION_WIDE_FIELD_MAPPING`, `DELETED_LINKED_ACCOUNT_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `CHANGED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `DELETED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `FORCED_LINKED_ACCOUNT_RESYNC`, `MUTED_ISSUE`, `GENERATED_MAGIC_LINK`, `ENABLED_MERGE_WEBHOOK`, `DISABLED_MERGE_WEBHOOK`, `MERGE_WEBHOOK_TARGET_CHANGED`, `END_USER_CREDENTIALS_ACCESSED`
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include audit trail events that occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**userEmail:** `*string` — If provided, this will return events associated with the specified user email. Please note that the email address reflects the user's email at the time of the event, and may not be their current email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage AvailableActions
<details><summary><code>client.FileStorage.AvailableActions.Retrieve() -> *filestorage.AvailableActions</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of models and actions available for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.AvailableActions.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Scopes
<details><summary><code>client.FileStorage.Scopes.DefaultScopesRetrieve() -> *filestorage.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the default permissions for Merge Common Models and fields across all Linked Accounts of a given category. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.Scopes.DefaultScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Scopes.LinkedAccountScopesRetrieve() -> *filestorage.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all available permissions for Merge Common Models and fields for a single Linked Account. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.Scopes.LinkedAccountScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Scopes.LinkedAccountScopesCreate(request) -> *filestorage.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update permissions for any Common Model or field for a single Linked Account. Any Scopes not set in this POST request will inherit the default Scopes. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.LinkedAccountCommonModelScopeDeserializerRequest{
        CommonModels: []*filestorage.IndividualCommonModelScopeDeserializerRequest{
            &filestorage.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Employee",
                ModelPermissions: map[string]*filestorage.ModelPermissionDeserializerRequest{
                    "READ": &filestorage.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            true,
                        ),
                    },
                    "WRITE": &filestorage.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
                FieldPermissions: &filestorage.FieldPermissionDeserializerRequest{
                    EnabledFields: []any{
                        "avatar",
                        "home_location",
                    },
                    DisabledFields: []any{
                        "work_location",
                    },
                },
            },
            &filestorage.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Benefit",
                ModelPermissions: map[string]*filestorage.ModelPermissionDeserializerRequest{
                    "WRITE": &filestorage.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
            },
        },
    }
client.FileStorage.Scopes.LinkedAccountScopesCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `[]*filestorage.IndividualCommonModelScopeDeserializerRequest` — The common models you want to update the scopes for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage DeleteAccount
<details><summary><code>client.FileStorage.DeleteAccount.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.DeleteAccount.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Drives
<details><summary><code>client.FileStorage.Drives.List() -> *filestorage.PaginatedDriveList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Drive` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.DrivesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.FileStorage.Drives.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return drives with this name. This performs an exact match.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Drives.Retrieve(Id) -> *filestorage.Drive</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Drive` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.DrivesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.FileStorage.Drives.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage FieldMapping
<details><summary><code>client.FileStorage.FieldMapping.FieldMappingsRetrieve() -> *filestorage.FieldMappingApiInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all Field Mappings for this Linked Account. Field Mappings are mappings between third-party Remote Fields and user defined Merge fields. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FieldMappingsRetrieveRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
    }
client.FileStorage.FieldMapping.FieldMappingsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.FieldMapping.FieldMappingsCreate(request) -> *filestorage.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create new Field Mappings that will be available after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.CreateFieldMappingRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
        TargetFieldName: "example_target_field_name",
        TargetFieldDescription: "this is a example description of the target field",
        RemoteFieldTraversalPath: []any{
            "example_remote_field",
        },
        RemoteMethod: "GET",
        RemoteUrlPath: "/example-url-path",
        CommonModelName: "ExampleCommonModel",
    }
client.FileStorage.FieldMapping.FieldMappingsCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldName:** `string` — The name of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldDescription:** `string` — The description of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**commonModelName:** `string` — The name of the Common Model that the remote field corresponds to in a given category.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.FieldMapping.FieldMappingsDestroy(FieldMappingId) -> *filestorage.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes Field Mappings for a Linked Account. All data related to this Field Mapping will be deleted and these changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.FieldMapping.FieldMappingsDestroy(
        context.TODO(),
        "field_mapping_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.FieldMapping.FieldMappingsPartialUpdate(FieldMappingId, request) -> *filestorage.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update existing Field Mappings for a Linked Account. Changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.PatchedEditFieldMappingRequest{}
client.FileStorage.FieldMapping.FieldMappingsPartialUpdate(
        context.TODO(),
        "field_mapping_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `*string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `*string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.FieldMapping.RemoteFieldsRetrieve() -> *filestorage.RemoteFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all remote fields for a Linked Account. Remote fields are third-party fields that are accessible after initial sync if remote_data is enabled. You can use remote fields to override existing Merge fields or map a new Merge field. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.RemoteFieldsRetrieveRequest{
        CommonModels: merge.String(
            "common_models",
        ),
        IncludeExampleValues: merge.String(
            "include_example_values",
        ),
    }
client.FileStorage.FieldMapping.RemoteFieldsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `*string` — A comma seperated list of Common Model names. If included, will only return Remote Fields for those Common Models.
    
</dd>
</dl>

<dl>
<dd>

**includeExampleValues:** `*string` — If true, will include example values, where available, for remote fields in the 3rd party platform. These examples come from active data from your customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.FieldMapping.TargetFieldsRetrieve() -> *filestorage.ExternalTargetFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all organization-wide Target Fields, this will not include any Linked Account specific Target Fields. Organization-wide Target Fields are additional fields appended to the Merge Common Model for all Linked Accounts in a category. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/target-fields/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.FieldMapping.TargetFieldsRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Files
<details><summary><code>client.FileStorage.Files.List() -> *filestorage.PaginatedFileList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `File` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FilesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        DriveId: merge.String(
            "drive_id",
        ),
        FolderId: merge.String(
            "folder_id",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        MimeType: merge.String(
            "mime_type",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        OrderBy: filestorage.FilesListRequestOrderByCreatedAtDescending.Ptr(),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.FileStorage.Files.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**driveId:** `*string` — Specifying a drive id returns only the files in that drive. Specifying null returns only the files outside the top-level drive.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*filestorage.FilesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**folderId:** `*string` — Specifying a folder id returns only the files in that folder. Specifying null returns only the files in root directory.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**mimeType:** `*string` — If provided, will only return files with these mime_types. Multiple values can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return files with this name. This performs an exact match.
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*filestorage.FilesListRequestOrderBy` — Overrides the default ordering for this endpoint. Possible values include: created_at, -created_at, modified_at, -modified_at.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Files.Create(request) -> *filestorage.FileStorageFileResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `File` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FileStorageFileEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &filestorage.FileRequest{},
    }
client.FileStorage.Files.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*filestorage.FileRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Files.Retrieve(Id) -> *filestorage.File</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `File` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FilesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.FileStorage.Files.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*filestorage.FilesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Files.DownloadRequestMetaRetrieve(Id) -> *filestorage.DownloadRequestMeta</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata to construct an authenticated file download request for a singular file, allowing you to download file directly from the third-party.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FilesDownloadRequestMetaRetrieveRequest{
        MimeType: merge.String(
            "mime_type",
        ),
    }
client.FileStorage.Files.DownloadRequestMetaRetrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**mimeType:** `*string` — If provided, specifies the export format of the file to be downloaded. For information on supported export formats, please refer to our <a href='https://help.merge.dev/en/articles/8615316-file-export-and-download-specification' target='_blank'>export format help center article</a>.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Files.DownloadRequestMetaList() -> *filestorage.PaginatedDownloadRequestMetaList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata to construct authenticated file download requests, allowing you to download files directly from the third-party.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FilesDownloadRequestMetaListRequest{
        CreatedAfter: merge.String(
            "created_after",
        ),
        CreatedBefore: merge.String(
            "created_before",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        MimeTypes: merge.String(
            "mime_types",
        ),
        ModifiedAfter: merge.String(
            "modified_after",
        ),
        ModifiedBefore: merge.String(
            "modified_before",
        ),
        OrderBy: filestorage.FilesDownloadRequestMetaListRequestOrderByCreatedAtDescending.Ptr(),
        PageSize: merge.Int(
            1,
        ),
    }
client.FileStorage.Files.DownloadRequestMetaList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*string` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*string` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**mimeTypes:** `*string` — A comma-separated list of preferred MIME types in order of priority. If supported by the third-party provider, the file(s) will be returned in the first supported MIME type from the list. The default MIME type is PDF. To see supported MIME types by file type, refer to our <a href='https://help.merge.dev/en/articles/8615316-file-export-and-download-specification' target='_blank'>export format help center article</a>.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*string` — If provided, will only return objects modified after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*string` — If provided, will only return objects modified before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*filestorage.FilesDownloadRequestMetaListRequestOrderBy` — Overrides the default ordering for this endpoint. Possible values include: created_at, -created_at, modified_at, -modified_at.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Files.MetaPostRetrieve() -> *filestorage.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `FileStorageFile` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.Files.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Folders
<details><summary><code>client.FileStorage.Folders.List() -> *filestorage.PaginatedFolderList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Folder` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FoldersListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        DriveId: merge.String(
            "drive_id",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        ParentFolderId: merge.String(
            "parent_folder_id",
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.FileStorage.Folders.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**driveId:** `*string` — If provided, will only return folders in this drive.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*filestorage.FoldersListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return folders with this name. This performs an exact match.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**parentFolderId:** `*string` — If provided, will only return folders in this parent folder. If null, will return folders in root directory.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Folders.Create(request) -> *filestorage.FileStorageFolderResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Folder` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FileStorageFolderEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &filestorage.FolderRequest{},
    }
client.FileStorage.Folders.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*filestorage.FolderRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Folders.Retrieve(Id) -> *filestorage.Folder</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Folder` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.FoldersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.FileStorage.Folders.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*filestorage.FoldersRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Folders.MetaPostRetrieve() -> *filestorage.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `FileStorageFolder` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.Folders.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage GenerateKey
<details><summary><code>client.FileStorage.GenerateKey.Create(request) -> *filestorage.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a remote key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.GenerateRemoteKeyRequest{
        Name: "Remote Deployment Key 1",
    }
client.FileStorage.GenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Groups
<details><summary><code>client.FileStorage.Groups.List() -> *filestorage.PaginatedGroupList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Group` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.GroupsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.FileStorage.Groups.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*filestorage.GroupsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Groups.Retrieve(Id) -> *filestorage.Group</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Group` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.GroupsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.FileStorage.Groups.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*filestorage.GroupsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Issues
<details><summary><code>client.FileStorage.Issues.List() -> *filestorage.PaginatedIssueList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets all issues for Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.IssuesListRequest{
        AccountToken: merge.String(
            "account_token",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        FirstIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        FirstIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeMuted: merge.String(
            "include_muted",
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        LastIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LastIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LinkedAccountId: merge.String(
            "linked_account_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        Status: filestorage.IssuesListRequestStatusOngoing.Ptr(),
    }
client.FileStorage.Issues.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include issues whose most recent action occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose first incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose first incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**includeMuted:** `*string` — If true, will include muted issues
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose last incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose last incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**linkedAccountId:** `*string` — If provided, will only include issues pertaining to the linked account passed in.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include issues whose most recent action occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**status:** `*filestorage.IssuesListRequestStatus` 

Status of the issue. Options: ('ONGOING', 'RESOLVED')

* `ONGOING` - ONGOING
* `RESOLVED` - RESOLVED
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Issues.Retrieve(Id) -> *filestorage.Issue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific issue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.Issues.Retrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage LinkToken
<details><summary><code>client.FileStorage.LinkToken.Create(request) -> *filestorage.LinkToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a link token to be used when linking a new end user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.EndUserDetailsRequest{
        EndUserEmailAddress: "example@gmail.com",
        EndUserOrganizationName: "Test Organization",
        EndUserOriginId: "12345",
        Categories: []filestorage.CategoriesEnum{
            filestorage.CategoriesEnumHris,
            filestorage.CategoriesEnumAts,
        },
    }
client.FileStorage.LinkToken.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**endUserEmailAddress:** `string` — Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `string` — Your end user's organization.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `string` — This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
    
</dd>
</dl>

<dl>
<dd>

**categories:** `[]*filestorage.CategoriesEnum` — The integration categories to show in Merge Link.
    
</dd>
</dl>

<dl>
<dd>

**integration:** `*string` — The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://docs.merge.dev/guides/merge-link/single-integration/.
    
</dd>
</dl>

<dl>
<dd>

**linkExpiryMins:** `*int` — An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
    
</dd>
</dl>

<dl>
<dd>

**shouldCreateMagicLinkUrl:** `*bool` — Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**hideAdminMagicLink:** `*bool` — Whether to generate a Magic Link URL on the Admin Needed screen during the linking flow. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**commonModels:** `[]*filestorage.CommonModelScopesBodyRequest` — An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
    
</dd>
</dl>

<dl>
<dd>

**categoryCommonModelScopes:** `map[string][]*filestorage.IndividualCommonModelScopeDeserializerRequest` — When creating a Link Token, you can set permissions for Common Models that will apply to the account that is going to be linked. Any model or field not specified in link token payload will default to existing settings.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*filestorage.EndUserDetailsRequestLanguage` 

The following subset of IETF language tags can be used to configure localization.

* `en` - en
* `de` - de
    
</dd>
</dl>

<dl>
<dd>

**areSyncsDisabled:** `*bool` — The boolean that indicates whether initial, periodic, and force syncs will be disabled.
    
</dd>
</dl>

<dl>
<dd>

**integrationSpecificConfig:** `map[string]any` — A JSON object containing integration-specific configuration options.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage LinkedAccounts
<details><summary><code>client.FileStorage.LinkedAccounts.List() -> *filestorage.PaginatedAccountDetailsAndActionsList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List linked accounts for your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.LinkedAccountsListRequest{
        Category: filestorage.LinkedAccountsListRequestCategoryAccounting.Ptr(),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndUserEmailAddress: merge.String(
            "end_user_email_address",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        EndUserOriginId: merge.String(
            "end_user_origin_id",
        ),
        EndUserOriginIds: merge.String(
            "end_user_origin_ids",
        ),
        Id: merge.String(
            "id",
        ),
        Ids: merge.String(
            "ids",
        ),
        IncludeDuplicates: merge.Bool(
            true,
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        IsTestAccount: merge.String(
            "is_test_account",
        ),
        PageSize: merge.Int(
            1,
        ),
        Status: merge.String(
            "status",
        ),
    }
client.FileStorage.LinkedAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**category:** `*filestorage.LinkedAccountsListRequestCategory` 

Options: `accounting`, `ats`, `crm`, `filestorage`, `hris`, `mktg`, `ticketing`

* `hris` - hris
* `ats` - ats
* `accounting` - accounting
* `ticketing` - ticketing
* `crm` - crm
* `mktg` - mktg
* `filestorage` - filestorage
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endUserEmailAddress:** `*string` — If provided, will only return linked accounts associated with the given email address.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` — If provided, will only return linked accounts associated with the given organization name.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `*string` — If provided, will only return linked accounts associated with the given origin ID.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginIds:** `*string` — Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
    
</dd>
</dl>

<dl>
<dd>

**includeDuplicates:** `*bool` — If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` — If provided, will only return linked accounts associated with the given integration name.
    
</dd>
</dl>

<dl>
<dd>

**isTestAccount:** `*string` — If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Filter by status. Options: `COMPLETE`, `IDLE`, `INCOMPLETE`, `RELINK_NEEDED`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Passthrough
<details><summary><code>client.FileStorage.Passthrough.Create(request) -> *filestorage.RemoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.DataPassthroughRequest{
        Method: filestorage.MethodEnumGet,
        Path: "/scooters",
    }
client.FileStorage.Passthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*filestorage.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage RegenerateKey
<details><summary><code>client.FileStorage.RegenerateKey.Create(request) -> *filestorage.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Exchange remote keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.RemoteKeyForRegenerationRequest{
        Name: "Remote Deployment Key 1",
    }
client.FileStorage.RegenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage SyncStatus
<details><summary><code>client.FileStorage.SyncStatus.List() -> *filestorage.PaginatedSyncStatusList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sync status for the current sync and the most recently finished sync. `last_sync_start` represents the most recent time any sync began. `last_sync_finished` represents the most recent time any sync completed. These timestamps may correspond to different sync instances which may result in a sync start time being later than a separate sync completed time. To ensure you are retrieving the latest available data reference the `last_sync_finished` timestamp where `last_sync_result` is `DONE`. Possible values for `status` and `last_sync_result` are `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`. Learn more about sync status in our [Help Center](https://help.merge.dev/en/articles/8184193-merge-sync-statuses).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.SyncStatusListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.FileStorage.SyncStatus.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage ForceResync
<details><summary><code>client.FileStorage.ForceResync.SyncStatusResyncCreate() -> []*filestorage.SyncStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Force re-sync of all models. This endpoint is available for monthly, quarterly, and highest sync frequency customers on the Professional or Enterprise plans. Doing so will consume a sync credit for the relevant linked account. Force re-syncs can also be triggered manually in the Merge Dashboard and is available for all customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.ForceResync.SyncStatusResyncCreate(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage Users
<details><summary><code>client.FileStorage.Users.List() -> *filestorage.PaginatedUserList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `User` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.UsersListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsMe: merge.String(
            "is_me",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.FileStorage.Users.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isMe:** `*string` — If provided, will only return the user object for requestor.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.Users.Retrieve(Id) -> *filestorage.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `User` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.UsersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.FileStorage.Users.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## FileStorage WebhookReceivers
<details><summary><code>client.FileStorage.WebhookReceivers.List() -> []*filestorage.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `WebhookReceiver` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.FileStorage.WebhookReceivers.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.FileStorage.WebhookReceivers.Create(request) -> *filestorage.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `WebhookReceiver` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &filestorage.WebhookReceiverRequest{
        Event: "event",
        IsActive: true,
    }
client.FileStorage.WebhookReceivers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**event:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**key:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris AccountDetails
<details><summary><code>client.Hris.AccountDetails.Retrieve() -> *hris.AccountDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details for a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.AccountDetails.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris AccountToken
<details><summary><code>client.Hris.AccountToken.Retrieve(PublicToken) -> *hris.AccountToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account token for the end user with the provided public token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.AccountToken.Retrieve(
        context.TODO(),
        "public_token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**publicToken:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris AsyncPassthrough
<details><summary><code>client.Hris.AsyncPassthrough.Create(request) -> *hris.AsyncPassthroughReciept</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Asynchronously pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.DataPassthroughRequest{
        Method: hris.MethodEnumGet,
        Path: "/scooters",
    }
client.Hris.AsyncPassthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*hris.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.AsyncPassthrough.Retrieve(AsyncPassthroughReceiptId) -> *hris.AsyncPassthroughRetrieveResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves data from earlier async-passthrough POST request
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.AsyncPassthrough.Retrieve(
        context.TODO(),
        "async_passthrough_receipt_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**asyncPassthroughReceiptId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris AuditTrail
<details><summary><code>client.Hris.AuditTrail.List() -> *hris.PaginatedAuditLogEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a list of audit trail events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.AuditTrailListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EventType: merge.String(
            "event_type",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        UserEmail: merge.String(
            "user_email",
        ),
    }
client.Hris.AuditTrail.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include audit trail events that occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*string` — If included, will only include events with the given event type. Possible values include: `CREATED_REMOTE_PRODUCTION_API_KEY`, `DELETED_REMOTE_PRODUCTION_API_KEY`, `CREATED_TEST_API_KEY`, `DELETED_TEST_API_KEY`, `REGENERATED_PRODUCTION_API_KEY`, `INVITED_USER`, `TWO_FACTOR_AUTH_ENABLED`, `TWO_FACTOR_AUTH_DISABLED`, `DELETED_LINKED_ACCOUNT`, `DELETED_ALL_COMMON_MODELS_FOR_LINKED_ACCOUNT`, `CREATED_DESTINATION`, `DELETED_DESTINATION`, `CHANGED_DESTINATION`, `CHANGED_SCOPES`, `CHANGED_PERSONAL_INFORMATION`, `CHANGED_ORGANIZATION_SETTINGS`, `ENABLED_INTEGRATION`, `DISABLED_INTEGRATION`, `ENABLED_CATEGORY`, `DISABLED_CATEGORY`, `CHANGED_PASSWORD`, `RESET_PASSWORD`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `CREATED_INTEGRATION_WIDE_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_FIELD_MAPPING`, `CHANGED_INTEGRATION_WIDE_FIELD_MAPPING`, `CHANGED_LINKED_ACCOUNT_FIELD_MAPPING`, `DELETED_INTEGRATION_WIDE_FIELD_MAPPING`, `DELETED_LINKED_ACCOUNT_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `CHANGED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `DELETED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `FORCED_LINKED_ACCOUNT_RESYNC`, `MUTED_ISSUE`, `GENERATED_MAGIC_LINK`, `ENABLED_MERGE_WEBHOOK`, `DISABLED_MERGE_WEBHOOK`, `MERGE_WEBHOOK_TARGET_CHANGED`, `END_USER_CREDENTIALS_ACCESSED`
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include audit trail events that occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**userEmail:** `*string` — If provided, this will return events associated with the specified user email. Please note that the email address reflects the user's email at the time of the event, and may not be their current email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris AvailableActions
<details><summary><code>client.Hris.AvailableActions.Retrieve() -> *hris.AvailableActions</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of models and actions available for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.AvailableActions.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris BankInfo
<details><summary><code>client.Hris.BankInfo.List() -> *hris.PaginatedBankInfoList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `BankInfo` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.BankInfoListRequest{
        AccountType: hris.BankInfoListRequestAccountTypeChecking.Ptr(),
        BankName: merge.String(
            "bank_name",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmployeeId: merge.String(
            "employee_id",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OrderBy: hris.BankInfoListRequestOrderByRemoteCreatedAtDescending.Ptr(),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.BankInfo.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountType:** `*hris.BankInfoListRequestAccountType` 

If provided, will only return BankInfo's with this account type. Options: ('SAVINGS', 'CHECKING')

* `SAVINGS` - SAVINGS
* `CHECKING` - CHECKING
    
</dd>
</dl>

<dl>
<dd>

**bankName:** `*string` — If provided, will only return BankInfo's with this bank name.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**employeeId:** `*string` — If provided, will only return bank accounts for this employee.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*hris.BankInfoListRequestOrderBy` — Overrides the default ordering for this endpoint. Possible values include: remote_created_at, -remote_created_at.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.BankInfo.Retrieve(Id) -> *hris.BankInfo</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `BankInfo` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.BankInfoRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.BankInfo.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Benefits
<details><summary><code>client.Hris.Benefits.List() -> *hris.PaginatedBenefitList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Benefit` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.BenefitsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmployeeId: merge.String(
            "employee_id",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.Benefits.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**employeeId:** `*string` — If provided, will return the benefits associated with the employee.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Benefits.Retrieve(Id) -> *hris.Benefit</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Benefit` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.BenefitsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.Benefits.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Companies
<details><summary><code>client.Hris.Companies.List() -> *hris.PaginatedCompanyList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Company` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.CompaniesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.Companies.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Companies.Retrieve(Id) -> *hris.Company</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Company` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.CompaniesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.Companies.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Scopes
<details><summary><code>client.Hris.Scopes.DefaultScopesRetrieve() -> *hris.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the default permissions for Merge Common Models and fields across all Linked Accounts of a given category. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.Scopes.DefaultScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Scopes.LinkedAccountScopesRetrieve() -> *hris.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all available permissions for Merge Common Models and fields for a single Linked Account. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.Scopes.LinkedAccountScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Scopes.LinkedAccountScopesCreate(request) -> *hris.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update permissions for any Common Model or field for a single Linked Account. Any Scopes not set in this POST request will inherit the default Scopes. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.LinkedAccountCommonModelScopeDeserializerRequest{
        CommonModels: []*hris.IndividualCommonModelScopeDeserializerRequest{
            &hris.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Employee",
                ModelPermissions: map[string]*hris.ModelPermissionDeserializerRequest{
                    "READ": &hris.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            true,
                        ),
                    },
                    "WRITE": &hris.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
                FieldPermissions: &hris.FieldPermissionDeserializerRequest{
                    EnabledFields: []any{
                        "avatar",
                        "home_location",
                    },
                    DisabledFields: []any{
                        "work_location",
                    },
                },
            },
            &hris.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Benefit",
                ModelPermissions: map[string]*hris.ModelPermissionDeserializerRequest{
                    "WRITE": &hris.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
            },
        },
    }
client.Hris.Scopes.LinkedAccountScopesCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `[]*hris.IndividualCommonModelScopeDeserializerRequest` — The common models you want to update the scopes for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris DeleteAccount
<details><summary><code>client.Hris.DeleteAccount.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.DeleteAccount.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Dependents
<details><summary><code>client.Hris.Dependents.List() -> *hris.PaginatedDependentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Dependent` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.DependentsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeSensitiveFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.Dependents.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeSensitiveFields:** `*bool` — Whether to include sensitive fields (such as social security numbers) in the response.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Dependents.Retrieve(Id) -> *hris.Dependent</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Dependent` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.DependentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeSensitiveFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.Dependents.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeSensitiveFields:** `*bool` — Whether to include sensitive fields (such as social security numbers) in the response.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris EmployeePayrollRuns
<details><summary><code>client.Hris.EmployeePayrollRuns.List() -> *hris.PaginatedEmployeePayrollRunList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `EmployeePayrollRun` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmployeePayrollRunsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmployeeId: merge.String(
            "employee_id",
        ),
        EndedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        EndedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        PayrollRunId: merge.String(
            "payroll_run_id",
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        StartedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        StartedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Hris.EmployeePayrollRuns.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**employeeId:** `*string` — If provided, will only return employee payroll runs for this employee.
    
</dd>
</dl>

<dl>
<dd>

**endedAfter:** `*time.Time` — If provided, will only return employee payroll runs ended after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**endedBefore:** `*time.Time` — If provided, will only return employee payroll runs ended before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.EmployeePayrollRunsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**payrollRunId:** `*string` — If provided, will only return employee payroll runs for this employee.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**startedAfter:** `*time.Time` — If provided, will only return employee payroll runs started after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**startedBefore:** `*time.Time` — If provided, will only return employee payroll runs started before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.EmployeePayrollRuns.Retrieve(Id) -> *hris.EmployeePayrollRun</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `EmployeePayrollRun` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmployeePayrollRunsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.EmployeePayrollRuns.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.EmployeePayrollRunsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Employees
<details><summary><code>client.Hris.Employees.List() -> *hris.PaginatedEmployeeList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Employee` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmployeesListRequest{
        CompanyId: merge.String(
            "company_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        DisplayFullName: merge.String(
            "display_full_name",
        ),
        EmploymentStatus: hris.EmployeesListRequestEmploymentStatusActive.Ptr(),
        EmploymentType: merge.String(
            "employment_type",
        ),
        FirstName: merge.String(
            "first_name",
        ),
        Groups: merge.String(
            "groups",
        ),
        HomeLocationId: merge.String(
            "home_location_id",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeSensitiveFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        JobTitle: merge.String(
            "job_title",
        ),
        LastName: merge.String(
            "last_name",
        ),
        ManagerId: merge.String(
            "manager_id",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        PayGroupId: merge.String(
            "pay_group_id",
        ),
        PersonalEmail: merge.String(
            "personal_email",
        ),
        RemoteFields: hris.EmployeesListRequestRemoteFieldsEmploymentStatus.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        ShowEnumOrigins: hris.EmployeesListRequestShowEnumOriginsEmploymentStatus.Ptr(),
        StartedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        StartedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TeamId: merge.String(
            "team_id",
        ),
        TerminatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        TerminatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        WorkEmail: merge.String(
            "work_email",
        ),
        WorkLocationId: merge.String(
            "work_location_id",
        ),
    }
client.Hris.Employees.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**companyId:** `*string` — If provided, will only return employees for this company.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**displayFullName:** `*string` — If provided, will only return employees with this display name.
    
</dd>
</dl>

<dl>
<dd>

**employmentStatus:** `*hris.EmployeesListRequestEmploymentStatus` 

If provided, will only return employees with this employment status.

* `ACTIVE` - ACTIVE
* `PENDING` - PENDING
* `INACTIVE` - INACTIVE
    
</dd>
</dl>

<dl>
<dd>

**employmentType:** `*string` — If provided, will only return employees that have an employment of the specified employment_type.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.EmployeesListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**firstName:** `*string` — If provided, will only return employees with this first name.
    
</dd>
</dl>

<dl>
<dd>

**groups:** `*string` — If provided, will only return employees matching the group ids; multiple groups can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**homeLocationId:** `*string` — If provided, will only return employees for this home location.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeSensitiveFields:** `*bool` — Whether to include sensitive fields (such as social security numbers) in the response.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**jobTitle:** `*string` — If provided, will only return employees that have an employment of the specified job_title.
    
</dd>
</dl>

<dl>
<dd>

**lastName:** `*string` — If provided, will only return employees with this last name.
    
</dd>
</dl>

<dl>
<dd>

**managerId:** `*string` — If provided, will only return employees for this manager.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**payGroupId:** `*string` — If provided, will only return employees for this pay group
    
</dd>
</dl>

<dl>
<dd>

**personalEmail:** `*string` — If provided, will only return Employees with this personal email
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.EmployeesListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.EmployeesListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**startedAfter:** `*time.Time` — If provided, will only return employees that started after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**startedBefore:** `*time.Time` — If provided, will only return employees that started before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**teamId:** `*string` — If provided, will only return employees for this team.
    
</dd>
</dl>

<dl>
<dd>

**terminatedAfter:** `*time.Time` — If provided, will only return employees that were terminated after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**terminatedBefore:** `*time.Time` — If provided, will only return employees that were terminated before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**workEmail:** `*string` — If provided, will only return Employees with this work email
    
</dd>
</dl>

<dl>
<dd>

**workLocationId:** `*string` — If provided, will only return employees for this location.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Employees.Create(request) -> *hris.EmployeeResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Employee` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmployeeEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &hris.EmployeeRequest{},
    }
client.Hris.Employees.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*hris.EmployeeRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Employees.Retrieve(Id) -> *hris.Employee</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Employee` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmployeesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeSensitiveFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: hris.EmployeesRetrieveRequestRemoteFieldsEmploymentStatus.Ptr(),
        ShowEnumOrigins: hris.EmployeesRetrieveRequestShowEnumOriginsEmploymentStatus.Ptr(),
    }
client.Hris.Employees.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.EmployeesRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeSensitiveFields:** `*bool` — Whether to include sensitive fields (such as social security numbers) in the response.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.EmployeesRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.EmployeesRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Employees.IgnoreCreate(ModelId, request) -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.IgnoreCommonModelRequest{
        Reason: &hris.IgnoreCommonModelRequestReason{
            ReasonEnum: hris.ReasonEnumGeneralCustomerRequest,
        },
    }
client.Hris.Employees.IgnoreCreate(
        context.TODO(),
        "model_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**modelId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**request:** `*hris.IgnoreCommonModelRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Employees.MetaPostRetrieve() -> *hris.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Employee` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.Employees.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris EmployerBenefits
<details><summary><code>client.Hris.EmployerBenefits.List() -> *hris.PaginatedEmployerBenefitList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `EmployerBenefit` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmployerBenefitsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.EmployerBenefits.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.EmployerBenefits.Retrieve(Id) -> *hris.EmployerBenefit</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `EmployerBenefit` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmployerBenefitsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.EmployerBenefits.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Employments
<details><summary><code>client.Hris.Employments.List() -> *hris.PaginatedEmploymentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Employment` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmploymentsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmployeeId: merge.String(
            "employee_id",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OrderBy: hris.EmploymentsListRequestOrderByEffectiveDateDescending.Ptr(),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: hris.EmploymentsListRequestRemoteFieldsEmploymentType.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        ShowEnumOrigins: hris.EmploymentsListRequestShowEnumOriginsEmploymentType.Ptr(),
    }
client.Hris.Employments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**employeeId:** `*string` — If provided, will only return employments for this employee.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.EmploymentsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*hris.EmploymentsListRequestOrderBy` — Overrides the default ordering for this endpoint. Possible values include: effective_date, -effective_date.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.EmploymentsListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.EmploymentsListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Employments.Retrieve(Id) -> *hris.Employment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Employment` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EmploymentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: hris.EmploymentsRetrieveRequestRemoteFieldsEmploymentType.Ptr(),
        ShowEnumOrigins: hris.EmploymentsRetrieveRequestShowEnumOriginsEmploymentType.Ptr(),
    }
client.Hris.Employments.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.EmploymentsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.EmploymentsRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.EmploymentsRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris FieldMapping
<details><summary><code>client.Hris.FieldMapping.FieldMappingsRetrieve() -> *hris.FieldMappingApiInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all Field Mappings for this Linked Account. Field Mappings are mappings between third-party Remote Fields and user defined Merge fields. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.FieldMappingsRetrieveRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
    }
client.Hris.FieldMapping.FieldMappingsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.FieldMapping.FieldMappingsCreate(request) -> *hris.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create new Field Mappings that will be available after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.CreateFieldMappingRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
        TargetFieldName: "example_target_field_name",
        TargetFieldDescription: "this is a example description of the target field",
        RemoteFieldTraversalPath: []any{
            "example_remote_field",
        },
        RemoteMethod: "GET",
        RemoteUrlPath: "/example-url-path",
        CommonModelName: "ExampleCommonModel",
    }
client.Hris.FieldMapping.FieldMappingsCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldName:** `string` — The name of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldDescription:** `string` — The description of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**commonModelName:** `string` — The name of the Common Model that the remote field corresponds to in a given category.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.FieldMapping.FieldMappingsDestroy(FieldMappingId) -> *hris.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes Field Mappings for a Linked Account. All data related to this Field Mapping will be deleted and these changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.FieldMapping.FieldMappingsDestroy(
        context.TODO(),
        "field_mapping_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.FieldMapping.FieldMappingsPartialUpdate(FieldMappingId, request) -> *hris.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update existing Field Mappings for a Linked Account. Changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.PatchedEditFieldMappingRequest{}
client.Hris.FieldMapping.FieldMappingsPartialUpdate(
        context.TODO(),
        "field_mapping_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `*string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `*string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.FieldMapping.RemoteFieldsRetrieve() -> *hris.RemoteFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all remote fields for a Linked Account. Remote fields are third-party fields that are accessible after initial sync if remote_data is enabled. You can use remote fields to override existing Merge fields or map a new Merge field. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.RemoteFieldsRetrieveRequest{
        CommonModels: merge.String(
            "common_models",
        ),
        IncludeExampleValues: merge.String(
            "include_example_values",
        ),
    }
client.Hris.FieldMapping.RemoteFieldsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `*string` — A comma seperated list of Common Model names. If included, will only return Remote Fields for those Common Models.
    
</dd>
</dl>

<dl>
<dd>

**includeExampleValues:** `*string` — If true, will include example values, where available, for remote fields in the 3rd party platform. These examples come from active data from your customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.FieldMapping.TargetFieldsRetrieve() -> *hris.ExternalTargetFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all organization-wide Target Fields, this will not include any Linked Account specific Target Fields. Organization-wide Target Fields are additional fields appended to the Merge Common Model for all Linked Accounts in a category. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/target-fields/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.FieldMapping.TargetFieldsRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris GenerateKey
<details><summary><code>client.Hris.GenerateKey.Create(request) -> *hris.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a remote key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.GenerateRemoteKeyRequest{
        Name: "Remote Deployment Key 1",
    }
client.Hris.GenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Groups
<details><summary><code>client.Hris.Groups.List() -> *hris.PaginatedGroupList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Group` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.GroupsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonlyUsedAsTeam: merge.String(
            "is_commonly_used_as_team",
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Names: merge.String(
            "names",
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Types: merge.String(
            "types",
        ),
    }
client.Hris.Groups.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonlyUsedAsTeam:** `*string` — If provided, specifies whether to return only Group objects which refer to a team in the third party platform. Note that this is an opinionated view based on how a team may be represented in the third party platform.
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**names:** `*string` — If provided, will only return groups with these names. Multiple values can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**types:** `*string` — If provided, will only return groups of these types. Multiple values can be separated by commas.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Groups.Retrieve(Id) -> *hris.Group</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Group` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.GroupsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.Groups.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Issues
<details><summary><code>client.Hris.Issues.List() -> *hris.PaginatedIssueList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets all issues for Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.IssuesListRequest{
        AccountToken: merge.String(
            "account_token",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        FirstIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        FirstIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeMuted: merge.String(
            "include_muted",
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        LastIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LastIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LinkedAccountId: merge.String(
            "linked_account_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        Status: hris.IssuesListRequestStatusOngoing.Ptr(),
    }
client.Hris.Issues.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include issues whose most recent action occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose first incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose first incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**includeMuted:** `*string` — If true, will include muted issues
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose last incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose last incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**linkedAccountId:** `*string` — If provided, will only include issues pertaining to the linked account passed in.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include issues whose most recent action occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**status:** `*hris.IssuesListRequestStatus` 

Status of the issue. Options: ('ONGOING', 'RESOLVED')

* `ONGOING` - ONGOING
* `RESOLVED` - RESOLVED
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Issues.Retrieve(Id) -> *hris.Issue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific issue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.Issues.Retrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris LinkToken
<details><summary><code>client.Hris.LinkToken.Create(request) -> *hris.LinkToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a link token to be used when linking a new end user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.EndUserDetailsRequest{
        EndUserEmailAddress: "example@gmail.com",
        EndUserOrganizationName: "Test Organization",
        EndUserOriginId: "12345",
        Categories: []hris.CategoriesEnum{
            hris.CategoriesEnumHris,
            hris.CategoriesEnumAts,
        },
    }
client.Hris.LinkToken.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**endUserEmailAddress:** `string` — Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `string` — Your end user's organization.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `string` — This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
    
</dd>
</dl>

<dl>
<dd>

**categories:** `[]*hris.CategoriesEnum` — The integration categories to show in Merge Link.
    
</dd>
</dl>

<dl>
<dd>

**integration:** `*string` — The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://docs.merge.dev/guides/merge-link/single-integration/.
    
</dd>
</dl>

<dl>
<dd>

**linkExpiryMins:** `*int` — An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
    
</dd>
</dl>

<dl>
<dd>

**shouldCreateMagicLinkUrl:** `*bool` — Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**hideAdminMagicLink:** `*bool` — Whether to generate a Magic Link URL on the Admin Needed screen during the linking flow. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**commonModels:** `[]*hris.CommonModelScopesBodyRequest` — An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
    
</dd>
</dl>

<dl>
<dd>

**categoryCommonModelScopes:** `map[string][]*hris.IndividualCommonModelScopeDeserializerRequest` — When creating a Link Token, you can set permissions for Common Models that will apply to the account that is going to be linked. Any model or field not specified in link token payload will default to existing settings.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*hris.EndUserDetailsRequestLanguage` 

The following subset of IETF language tags can be used to configure localization.

* `en` - en
* `de` - de
    
</dd>
</dl>

<dl>
<dd>

**areSyncsDisabled:** `*bool` — The boolean that indicates whether initial, periodic, and force syncs will be disabled.
    
</dd>
</dl>

<dl>
<dd>

**integrationSpecificConfig:** `map[string]any` — A JSON object containing integration-specific configuration options.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris LinkedAccounts
<details><summary><code>client.Hris.LinkedAccounts.List() -> *hris.PaginatedAccountDetailsAndActionsList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List linked accounts for your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.LinkedAccountsListRequest{
        Category: hris.LinkedAccountsListRequestCategoryAccounting.Ptr(),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndUserEmailAddress: merge.String(
            "end_user_email_address",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        EndUserOriginId: merge.String(
            "end_user_origin_id",
        ),
        EndUserOriginIds: merge.String(
            "end_user_origin_ids",
        ),
        Id: merge.String(
            "id",
        ),
        Ids: merge.String(
            "ids",
        ),
        IncludeDuplicates: merge.Bool(
            true,
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        IsTestAccount: merge.String(
            "is_test_account",
        ),
        PageSize: merge.Int(
            1,
        ),
        Status: merge.String(
            "status",
        ),
    }
client.Hris.LinkedAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**category:** `*hris.LinkedAccountsListRequestCategory` 

Options: `accounting`, `ats`, `crm`, `filestorage`, `hris`, `mktg`, `ticketing`

* `hris` - hris
* `ats` - ats
* `accounting` - accounting
* `ticketing` - ticketing
* `crm` - crm
* `mktg` - mktg
* `filestorage` - filestorage
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endUserEmailAddress:** `*string` — If provided, will only return linked accounts associated with the given email address.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` — If provided, will only return linked accounts associated with the given organization name.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `*string` — If provided, will only return linked accounts associated with the given origin ID.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginIds:** `*string` — Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
    
</dd>
</dl>

<dl>
<dd>

**includeDuplicates:** `*bool` — If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` — If provided, will only return linked accounts associated with the given integration name.
    
</dd>
</dl>

<dl>
<dd>

**isTestAccount:** `*string` — If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Filter by status. Options: `COMPLETE`, `IDLE`, `INCOMPLETE`, `RELINK_NEEDED`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Locations
<details><summary><code>client.Hris.Locations.List() -> *hris.PaginatedLocationList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Location` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.LocationsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        LocationType: hris.LocationsListRequestLocationTypeHome.Ptr(),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: hris.LocationsListRequestRemoteFieldsCountry.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        ShowEnumOrigins: hris.LocationsListRequestShowEnumOriginsCountry.Ptr(),
    }
client.Hris.Locations.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**locationType:** `*hris.LocationsListRequestLocationType` 

If provided, will only return locations with this location_type

* `HOME` - HOME
* `WORK` - WORK
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.LocationsListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.LocationsListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Locations.Retrieve(Id) -> *hris.Location</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Location` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.LocationsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: hris.LocationsRetrieveRequestRemoteFieldsCountry.Ptr(),
        ShowEnumOrigins: hris.LocationsRetrieveRequestShowEnumOriginsCountry.Ptr(),
    }
client.Hris.Locations.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.LocationsRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.LocationsRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Passthrough
<details><summary><code>client.Hris.Passthrough.Create(request) -> *hris.RemoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.DataPassthroughRequest{
        Method: hris.MethodEnumGet,
        Path: "/scooters",
    }
client.Hris.Passthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*hris.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris PayGroups
<details><summary><code>client.Hris.PayGroups.List() -> *hris.PaginatedPayGroupList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `PayGroup` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.PayGroupsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.PayGroups.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.PayGroups.Retrieve(Id) -> *hris.PayGroup</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `PayGroup` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.PayGroupsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.PayGroups.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris PayrollRuns
<details><summary><code>client.Hris.PayrollRuns.List() -> *hris.PaginatedPayrollRunList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `PayrollRun` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.PayrollRunsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        EndedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: hris.PayrollRunsListRequestRemoteFieldsRunState.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        RunType: hris.PayrollRunsListRequestRunTypeCorrection.Ptr(),
        ShowEnumOrigins: hris.PayrollRunsListRequestShowEnumOriginsRunState.Ptr(),
        StartedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        StartedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Hris.PayrollRuns.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endedAfter:** `*time.Time` — If provided, will only return payroll runs ended after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**endedBefore:** `*time.Time` — If provided, will only return payroll runs ended before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.PayrollRunsListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**runType:** `*hris.PayrollRunsListRequestRunType` 

If provided, will only return PayrollRun's with this status. Options: ('REGULAR', 'OFF_CYCLE', 'CORRECTION', 'TERMINATION', 'SIGN_ON_BONUS')

* `REGULAR` - REGULAR
* `OFF_CYCLE` - OFF_CYCLE
* `CORRECTION` - CORRECTION
* `TERMINATION` - TERMINATION
* `SIGN_ON_BONUS` - SIGN_ON_BONUS
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.PayrollRunsListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**startedAfter:** `*time.Time` — If provided, will only return payroll runs started after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**startedBefore:** `*time.Time` — If provided, will only return payroll runs started before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.PayrollRuns.Retrieve(Id) -> *hris.PayrollRun</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `PayrollRun` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.PayrollRunsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: hris.PayrollRunsRetrieveRequestRemoteFieldsRunState.Ptr(),
        ShowEnumOrigins: hris.PayrollRunsRetrieveRequestShowEnumOriginsRunState.Ptr(),
    }
client.Hris.PayrollRuns.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.PayrollRunsRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.PayrollRunsRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris RegenerateKey
<details><summary><code>client.Hris.RegenerateKey.Create(request) -> *hris.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Exchange remote keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.RemoteKeyForRegenerationRequest{
        Name: "Remote Deployment Key 1",
    }
client.Hris.RegenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris SyncStatus
<details><summary><code>client.Hris.SyncStatus.List() -> *hris.PaginatedSyncStatusList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sync status for the current sync and the most recently finished sync. `last_sync_start` represents the most recent time any sync began. `last_sync_finished` represents the most recent time any sync completed. These timestamps may correspond to different sync instances which may result in a sync start time being later than a separate sync completed time. To ensure you are retrieving the latest available data reference the `last_sync_finished` timestamp where `last_sync_result` is `DONE`. Possible values for `status` and `last_sync_result` are `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`. Learn more about sync status in our [Help Center](https://help.merge.dev/en/articles/8184193-merge-sync-statuses).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.SyncStatusListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Hris.SyncStatus.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris ForceResync
<details><summary><code>client.Hris.ForceResync.SyncStatusResyncCreate() -> []*hris.SyncStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Force re-sync of all models. This endpoint is available for monthly, quarterly, and highest sync frequency customers on the Professional or Enterprise plans. Doing so will consume a sync credit for the relevant linked account. Force re-syncs can also be triggered manually in the Merge Dashboard and is available for all customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.ForceResync.SyncStatusResyncCreate(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris Teams
<details><summary><code>client.Hris.Teams.List() -> *hris.PaginatedTeamList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Team` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TeamsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        ParentTeamId: merge.String(
            "parent_team_id",
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.Teams.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**parentTeamId:** `*string` — If provided, will only return teams with this parent team.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.Teams.Retrieve(Id) -> *hris.Team</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Team` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TeamsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.Teams.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris TimeOff
<details><summary><code>client.Hris.TimeOff.List() -> *hris.PaginatedTimeOffList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `TimeOff` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimeOffListRequest{
        ApproverId: merge.String(
            "approver_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmployeeId: merge.String(
            "employee_id",
        ),
        EndedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        EndedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteFields: hris.TimeOffListRequestRemoteFieldsRequestType.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        RequestType: hris.TimeOffListRequestRequestTypeBereavement.Ptr(),
        ShowEnumOrigins: hris.TimeOffListRequestShowEnumOriginsRequestType.Ptr(),
        StartedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        StartedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Status: hris.TimeOffListRequestStatusApproved.Ptr(),
    }
client.Hris.TimeOff.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**approverId:** `*string` — If provided, will only return time off for this approver.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**employeeId:** `*string` — If provided, will only return time off for this employee.
    
</dd>
</dl>

<dl>
<dd>

**endedAfter:** `*time.Time` — If provided, will only return employees that ended after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**endedBefore:** `*time.Time` — If provided, will only return time-offs that ended before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.TimeOffListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.TimeOffListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**requestType:** `*hris.TimeOffListRequestRequestType` 

If provided, will only return TimeOff with this request type. Options: ('VACATION', 'SICK', 'PERSONAL', 'JURY_DUTY', 'VOLUNTEER', 'BEREAVEMENT')

* `VACATION` - VACATION
* `SICK` - SICK
* `PERSONAL` - PERSONAL
* `JURY_DUTY` - JURY_DUTY
* `VOLUNTEER` - VOLUNTEER
* `BEREAVEMENT` - BEREAVEMENT
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.TimeOffListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**startedAfter:** `*time.Time` — If provided, will only return time-offs that started after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**startedBefore:** `*time.Time` — If provided, will only return time-offs that started before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*hris.TimeOffListRequestStatus` 

If provided, will only return TimeOff with this status. Options: ('REQUESTED', 'APPROVED', 'DECLINED', 'CANCELLED', 'DELETED')

* `REQUESTED` - REQUESTED
* `APPROVED` - APPROVED
* `DECLINED` - DECLINED
* `CANCELLED` - CANCELLED
* `DELETED` - DELETED
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.TimeOff.Create(request) -> *hris.TimeOffResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `TimeOff` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimeOffEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &hris.TimeOffRequest{},
    }
client.Hris.TimeOff.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*hris.TimeOffRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.TimeOff.Retrieve(Id) -> *hris.TimeOff</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `TimeOff` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimeOffRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: hris.TimeOffRetrieveRequestRemoteFieldsRequestType.Ptr(),
        ShowEnumOrigins: hris.TimeOffRetrieveRequestShowEnumOriginsRequestType.Ptr(),
    }
client.Hris.TimeOff.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*hris.TimeOffRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*hris.TimeOffRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*hris.TimeOffRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.TimeOff.MetaPostRetrieve() -> *hris.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `TimeOff` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.TimeOff.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris TimeOffBalances
<details><summary><code>client.Hris.TimeOffBalances.List() -> *hris.PaginatedTimeOffBalanceList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `TimeOffBalance` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimeOffBalancesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmployeeId: merge.String(
            "employee_id",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        PolicyType: hris.TimeOffBalancesListRequestPolicyTypeBereavement.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Hris.TimeOffBalances.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**employeeId:** `*string` — If provided, will only return time off balances for this employee.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**policyType:** `*hris.TimeOffBalancesListRequestPolicyType` 

If provided, will only return TimeOffBalance with this policy type. Options: ('VACATION', 'SICK', 'PERSONAL', 'JURY_DUTY', 'VOLUNTEER', 'BEREAVEMENT')

* `VACATION` - VACATION
* `SICK` - SICK
* `PERSONAL` - PERSONAL
* `JURY_DUTY` - JURY_DUTY
* `VOLUNTEER` - VOLUNTEER
* `BEREAVEMENT` - BEREAVEMENT
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.TimeOffBalances.Retrieve(Id) -> *hris.TimeOffBalance</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `TimeOffBalance` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimeOffBalancesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.TimeOffBalances.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris TimesheetEntries
<details><summary><code>client.Hris.TimesheetEntries.List() -> *hris.PaginatedTimesheetEntryList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `TimesheetEntry` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimesheetEntriesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmployeeId: merge.String(
            "employee_id",
        ),
        EndedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        EndedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        OrderBy: hris.TimesheetEntriesListRequestOrderByStartTimeDescending.Ptr(),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        StartedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        StartedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
    }
client.Hris.TimesheetEntries.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**employeeId:** `*string` — If provided, will only return timesheet entries for this employee.
    
</dd>
</dl>

<dl>
<dd>

**endedAfter:** `*time.Time` — If provided, will only return timesheet entries ended after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**endedBefore:** `*time.Time` — If provided, will only return timesheet entries ended before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**orderBy:** `*hris.TimesheetEntriesListRequestOrderBy` — Overrides the default ordering for this endpoint. Possible values include: start_time, -start_time.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**startedAfter:** `*time.Time` — If provided, will only return timesheet entries started after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**startedBefore:** `*time.Time` — If provided, will only return timesheet entries started before this datetime.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.TimesheetEntries.Create(request) -> *hris.TimesheetEntryResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `TimesheetEntry` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimesheetEntryEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &hris.TimesheetEntryRequest{},
    }
client.Hris.TimesheetEntries.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*hris.TimesheetEntryRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.TimesheetEntries.Retrieve(Id) -> *hris.TimesheetEntry</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `TimesheetEntry` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.TimesheetEntriesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Hris.TimesheetEntries.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.TimesheetEntries.MetaPostRetrieve() -> *hris.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `TimesheetEntry` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.TimesheetEntries.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Hris WebhookReceivers
<details><summary><code>client.Hris.WebhookReceivers.List() -> []*hris.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `WebhookReceiver` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Hris.WebhookReceivers.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Hris.WebhookReceivers.Create(request) -> *hris.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `WebhookReceiver` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &hris.WebhookReceiverRequest{
        Event: "event",
        IsActive: true,
    }
client.Hris.WebhookReceivers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**event:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**key:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing AccountDetails
<details><summary><code>client.Ticketing.AccountDetails.Retrieve() -> *ticketing.AccountDetails</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get details for a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.AccountDetails.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing AccountToken
<details><summary><code>client.Ticketing.AccountToken.Retrieve(PublicToken) -> *ticketing.AccountToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns the account token for the end user with the provided public token.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.AccountToken.Retrieve(
        context.TODO(),
        "public_token",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**publicToken:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Accounts
<details><summary><code>client.Ticketing.Accounts.List() -> *ticketing.PaginatedAccountList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Account` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.AccountsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ticketing.Accounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Accounts.Retrieve(Id) -> *ticketing.Account</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Account` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.AccountsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Accounts.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing AsyncPassthrough
<details><summary><code>client.Ticketing.AsyncPassthrough.Create(request) -> *ticketing.AsyncPassthroughReciept</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Asynchronously pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.DataPassthroughRequest{
        Method: ticketing.MethodEnumGet,
        Path: "/scooters",
    }
client.Ticketing.AsyncPassthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*ticketing.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.AsyncPassthrough.Retrieve(AsyncPassthroughReceiptId) -> *ticketing.AsyncPassthroughRetrieveResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Retrieves data from earlier async-passthrough POST request
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.AsyncPassthrough.Retrieve(
        context.TODO(),
        "async_passthrough_receipt_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**asyncPassthroughReceiptId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Attachments
<details><summary><code>client.Ticketing.Attachments.List() -> *ticketing.PaginatedAttachmentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Attachment` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.AttachmentsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteCreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        TicketId: merge.String(
            "ticket_id",
        ),
    }
client.Ticketing.Attachments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteCreatedAfter:** `*time.Time` — If provided, will only return attachments created in the third party platform after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**ticketId:** `*string` — If provided, will only return comments for this ticket.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Attachments.Create(request) -> *ticketing.TicketingAttachmentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates an `Attachment` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketingAttachmentEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ticketing.AttachmentRequest{},
    }
client.Ticketing.Attachments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ticketing.AttachmentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Attachments.Retrieve(Id) -> *ticketing.Attachment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns an `Attachment` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.AttachmentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Attachments.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Attachments.MetaPostRetrieve() -> *ticketing.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `TicketingAttachment` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.Attachments.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing AuditTrail
<details><summary><code>client.Ticketing.AuditTrail.List() -> *ticketing.PaginatedAuditLogEventList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets a list of audit trail events.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.AuditTrailListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EventType: merge.String(
            "event_type",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        UserEmail: merge.String(
            "user_email",
        ),
    }
client.Ticketing.AuditTrail.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include audit trail events that occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**eventType:** `*string` — If included, will only include events with the given event type. Possible values include: `CREATED_REMOTE_PRODUCTION_API_KEY`, `DELETED_REMOTE_PRODUCTION_API_KEY`, `CREATED_TEST_API_KEY`, `DELETED_TEST_API_KEY`, `REGENERATED_PRODUCTION_API_KEY`, `REGENERATED_WEBHOOK_SIGNATURE`, `INVITED_USER`, `TWO_FACTOR_AUTH_ENABLED`, `TWO_FACTOR_AUTH_DISABLED`, `DELETED_LINKED_ACCOUNT`, `DELETED_ALL_COMMON_MODELS_FOR_LINKED_ACCOUNT`, `CREATED_DESTINATION`, `DELETED_DESTINATION`, `CHANGED_DESTINATION`, `CHANGED_SCOPES`, `CHANGED_PERSONAL_INFORMATION`, `CHANGED_ORGANIZATION_SETTINGS`, `ENABLED_INTEGRATION`, `DISABLED_INTEGRATION`, `ENABLED_CATEGORY`, `DISABLED_CATEGORY`, `CHANGED_PASSWORD`, `RESET_PASSWORD`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `ENABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_ORGANIZATION`, `DISABLED_REDACT_UNMAPPED_DATA_FOR_LINKED_ACCOUNT`, `CREATED_INTEGRATION_WIDE_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_FIELD_MAPPING`, `CHANGED_INTEGRATION_WIDE_FIELD_MAPPING`, `CHANGED_LINKED_ACCOUNT_FIELD_MAPPING`, `DELETED_INTEGRATION_WIDE_FIELD_MAPPING`, `DELETED_LINKED_ACCOUNT_FIELD_MAPPING`, `CREATED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `CHANGED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `DELETED_LINKED_ACCOUNT_COMMON_MODEL_OVERRIDE`, `FORCED_LINKED_ACCOUNT_RESYNC`, `MUTED_ISSUE`, `GENERATED_MAGIC_LINK`, `ENABLED_MERGE_WEBHOOK`, `DISABLED_MERGE_WEBHOOK`, `MERGE_WEBHOOK_TARGET_CHANGED`, `END_USER_CREDENTIALS_ACCESSED`
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include audit trail events that occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**userEmail:** `*string` — If provided, this will return events associated with the specified user email. Please note that the email address reflects the user's email at the time of the event, and may not be their current email.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing AvailableActions
<details><summary><code>client.Ticketing.AvailableActions.Retrieve() -> *ticketing.AvailableActions</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of models and actions available for an account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.AvailableActions.Retrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Collections
<details><summary><code>client.Ticketing.Collections.List() -> *ticketing.PaginatedCollectionList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Collection` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.CollectionsListRequest{
        CollectionType: ticketing.CollectionsListRequestCollectionTypeEmpty.Ptr(),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        ParentCollectionId: merge.String(
            "parent_collection_id",
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ticketing.Collections.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionType:** `*ticketing.CollectionsListRequestCollectionType` — If provided, will only return collections of the given type.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return collections with this name.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**parentCollectionId:** `*string` — If provided, will only return collections whose parent collection matches the given id.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Collections.ViewersList(CollectionId) -> *ticketing.PaginatedViewerList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Viewer` objects that point to a User id or Team id that is either an assignee or viewer on a `Collection` with the given id. [Learn more.](https://help.merge.dev/en/articles/10333658-ticketing-access-control-list-acls)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.CollectionsViewersListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Ticketing.Collections.ViewersList(
        context.TODO(),
        "collection_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.CollectionsViewersListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Collections.Retrieve(Id) -> *ticketing.Collection</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Collection` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.CollectionsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Collections.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*string` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*string` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Comments
<details><summary><code>client.Ticketing.Comments.List() -> *ticketing.PaginatedCommentList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Comment` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.CommentsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteCreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        TicketId: merge.String(
            "ticket_id",
        ),
    }
client.Ticketing.Comments.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.CommentsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteCreatedAfter:** `*time.Time` — If provided, will only return Comments created in the third party platform after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**ticketId:** `*string` — If provided, will only return comments for this ticket.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Comments.Create(request) -> *ticketing.CommentResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Comment` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.CommentEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ticketing.CommentRequest{},
    }
client.Ticketing.Comments.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ticketing.CommentRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Comments.Retrieve(Id) -> *ticketing.Comment</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Comment` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.CommentsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Comments.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.CommentsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Comments.MetaPostRetrieve() -> *ticketing.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Comment` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.Comments.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Contacts
<details><summary><code>client.Ticketing.Contacts.List() -> *ticketing.PaginatedContactList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Contact` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.ContactsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmailAddress: merge.String(
            "email_address",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ticketing.Contacts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — If provided, will only return Contacts that match this email.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Contacts.Create(request) -> *ticketing.TicketingContactResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Contact` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketingContactEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ticketing.ContactRequest{},
    }
client.Ticketing.Contacts.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ticketing.ContactRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Contacts.Retrieve(Id) -> *ticketing.Contact</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Contact` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.ContactsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Contacts.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*string` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Contacts.MetaPostRetrieve() -> *ticketing.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `TicketingContact` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.Contacts.MetaPostRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Scopes
<details><summary><code>client.Ticketing.Scopes.DefaultScopesRetrieve() -> *ticketing.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get the default permissions for Merge Common Models and fields across all Linked Accounts of a given category. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.Scopes.DefaultScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Scopes.LinkedAccountScopesRetrieve() -> *ticketing.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all available permissions for Merge Common Models and fields for a single Linked Account. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.Scopes.LinkedAccountScopesRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Scopes.LinkedAccountScopesCreate(request) -> *ticketing.CommonModelScopeApi</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Update permissions for any Common Model or field for a single Linked Account. Any Scopes not set in this POST request will inherit the default Scopes. [Learn more](https://help.merge.dev/en/articles/5950052-common-model-and-field-scopes)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.LinkedAccountCommonModelScopeDeserializerRequest{
        CommonModels: []*ticketing.IndividualCommonModelScopeDeserializerRequest{
            &ticketing.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Employee",
                ModelPermissions: map[string]*ticketing.ModelPermissionDeserializerRequest{
                    "READ": &ticketing.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            true,
                        ),
                    },
                    "WRITE": &ticketing.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
                FieldPermissions: &ticketing.FieldPermissionDeserializerRequest{
                    EnabledFields: []any{
                        "avatar",
                        "home_location",
                    },
                    DisabledFields: []any{
                        "work_location",
                    },
                },
            },
            &ticketing.IndividualCommonModelScopeDeserializerRequest{
                ModelName: "Benefit",
                ModelPermissions: map[string]*ticketing.ModelPermissionDeserializerRequest{
                    "WRITE": &ticketing.ModelPermissionDeserializerRequest{
                        IsEnabled: merge.Bool(
                            false,
                        ),
                    },
                },
            },
        },
    }
client.Ticketing.Scopes.LinkedAccountScopesCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `[]*ticketing.IndividualCommonModelScopeDeserializerRequest` — The common models you want to update the scopes for
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing DeleteAccount
<details><summary><code>client.Ticketing.DeleteAccount.Delete() -> error</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Delete a linked account.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.DeleteAccount.Delete(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing FieldMapping
<details><summary><code>client.Ticketing.FieldMapping.FieldMappingsRetrieve() -> *ticketing.FieldMappingApiInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all Field Mappings for this Linked Account. Field Mappings are mappings between third-party Remote Fields and user defined Merge fields. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.FieldMappingsRetrieveRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
    }
client.Ticketing.FieldMapping.FieldMappingsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.FieldMapping.FieldMappingsCreate(request) -> *ticketing.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create new Field Mappings that will be available after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.CreateFieldMappingRequest{
        ExcludeRemoteFieldMetadata: merge.Bool(
            true,
        ),
        TargetFieldName: "example_target_field_name",
        TargetFieldDescription: "this is a example description of the target field",
        RemoteFieldTraversalPath: []any{
            "example_remote_field",
        },
        RemoteMethod: "GET",
        RemoteUrlPath: "/example-url-path",
        CommonModelName: "ExampleCommonModel",
    }
client.Ticketing.FieldMapping.FieldMappingsCreate(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**excludeRemoteFieldMetadata:** `*bool` — If `true`, remote fields metadata is excluded from each field mapping instance (i.e. `remote_fields.remote_key_name` and `remote_fields.schema` will be null). This will increase the speed of the request since these fields require some calculations.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldName:** `string` — The name of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**targetFieldDescription:** `string` — The description of the target field you want this remote field to map to.
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**commonModelName:** `string` — The name of the Common Model that the remote field corresponds to in a given category.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.FieldMapping.FieldMappingsDestroy(FieldMappingId) -> *ticketing.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Deletes Field Mappings for a Linked Account. All data related to this Field Mapping will be deleted and these changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.FieldMapping.FieldMappingsDestroy(
        context.TODO(),
        "field_mapping_id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.FieldMapping.FieldMappingsPartialUpdate(FieldMappingId, request) -> *ticketing.FieldMappingInstanceResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create or update existing Field Mappings for a Linked Account. Changes will be reflected after the next scheduled sync. This will cause the next sync for this Linked Account to sync **ALL** data from start.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.PatchedEditFieldMappingRequest{}
client.Ticketing.FieldMapping.FieldMappingsPartialUpdate(
        context.TODO(),
        "field_mapping_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**fieldMappingId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**remoteFieldTraversalPath:** `[]any` — The field traversal path of the remote field listed when you hit the GET /remote-fields endpoint.
    
</dd>
</dl>

<dl>
<dd>

**remoteMethod:** `*string` — The method of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>

<dl>
<dd>

**remoteUrlPath:** `*string` — The path of the remote endpoint where the remote field is coming from.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.FieldMapping.RemoteFieldsRetrieve() -> *ticketing.RemoteFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all remote fields for a Linked Account. Remote fields are third-party fields that are accessible after initial sync if remote_data is enabled. You can use remote fields to override existing Merge fields or map a new Merge field. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/overview/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.RemoteFieldsRetrieveRequest{
        CommonModels: merge.String(
            "common_models",
        ),
        IncludeExampleValues: merge.String(
            "include_example_values",
        ),
    }
client.Ticketing.FieldMapping.RemoteFieldsRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**commonModels:** `*string` — A comma seperated list of Common Model names. If included, will only return Remote Fields for those Common Models.
    
</dd>
</dl>

<dl>
<dd>

**includeExampleValues:** `*string` — If true, will include example values, where available, for remote fields in the 3rd party platform. These examples come from active data from your customers.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.FieldMapping.TargetFieldsRetrieve() -> *ticketing.ExternalTargetFieldApiResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get all organization-wide Target Fields, this will not include any Linked Account specific Target Fields. Organization-wide Target Fields are additional fields appended to the Merge Common Model for all Linked Accounts in a category. [Learn more](https://docs.merge.dev/supplemental-data/field-mappings/target-fields/).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.FieldMapping.TargetFieldsRetrieve(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing GenerateKey
<details><summary><code>client.Ticketing.GenerateKey.Create(request) -> *ticketing.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Create a remote key.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.GenerateRemoteKeyRequest{
        Name: "Remote Deployment Key 1",
    }
client.Ticketing.GenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Issues
<details><summary><code>client.Ticketing.Issues.List() -> *ticketing.PaginatedIssueList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Gets all issues for Organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.IssuesListRequest{
        AccountToken: merge.String(
            "account_token",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndDate: merge.String(
            "end_date",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        FirstIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        FirstIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeMuted: merge.String(
            "include_muted",
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        LastIncidentTimeAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LastIncidentTimeBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        LinkedAccountId: merge.String(
            "linked_account_id",
        ),
        PageSize: merge.Int(
            1,
        ),
        StartDate: merge.String(
            "start_date",
        ),
        Status: ticketing.IssuesListRequestStatusOngoing.Ptr(),
    }
client.Ticketing.Issues.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountToken:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endDate:** `*string` — If included, will only include issues whose most recent action occurred before this time
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose first incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**firstIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose first incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**includeMuted:** `*string` — If true, will include muted issues
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeAfter:** `*time.Time` — If provided, will only return issues whose last incident time was after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**lastIncidentTimeBefore:** `*time.Time` — If provided, will only return issues whose last incident time was before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**linkedAccountId:** `*string` — If provided, will only include issues pertaining to the linked account passed in.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**startDate:** `*string` — If included, will only include issues whose most recent action occurred after this time
    
</dd>
</dl>

<dl>
<dd>

**status:** `*ticketing.IssuesListRequestStatus` 

Status of the issue. Options: ('ONGOING', 'RESOLVED')

* `ONGOING` - ONGOING
* `RESOLVED` - RESOLVED
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Issues.Retrieve(Id) -> *ticketing.Issue</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get a specific issue.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.Issues.Retrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing LinkToken
<details><summary><code>client.Ticketing.LinkToken.Create(request) -> *ticketing.LinkToken</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a link token to be used when linking a new end user.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.EndUserDetailsRequest{
        EndUserEmailAddress: "example@gmail.com",
        EndUserOrganizationName: "Test Organization",
        EndUserOriginId: "12345",
        Categories: []ticketing.CategoriesEnum{
            ticketing.CategoriesEnumHris,
            ticketing.CategoriesEnumAts,
        },
    }
client.Ticketing.LinkToken.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**endUserEmailAddress:** `string` — Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `string` — Your end user's organization.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `string` — This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
    
</dd>
</dl>

<dl>
<dd>

**categories:** `[]*ticketing.CategoriesEnum` — The integration categories to show in Merge Link.
    
</dd>
</dl>

<dl>
<dd>

**integration:** `*string` — The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://docs.merge.dev/guides/merge-link/single-integration/.
    
</dd>
</dl>

<dl>
<dd>

**linkExpiryMins:** `*int` — An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
    
</dd>
</dl>

<dl>
<dd>

**shouldCreateMagicLinkUrl:** `*bool` — Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**hideAdminMagicLink:** `*bool` — Whether to generate a Magic Link URL on the Admin Needed screen during the linking flow. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
    
</dd>
</dl>

<dl>
<dd>

**commonModels:** `[]*ticketing.CommonModelScopesBodyRequest` — An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
    
</dd>
</dl>

<dl>
<dd>

**categoryCommonModelScopes:** `map[string][]*ticketing.IndividualCommonModelScopeDeserializerRequest` — When creating a Link Token, you can set permissions for Common Models that will apply to the account that is going to be linked. Any model or field not specified in link token payload will default to existing settings.
    
</dd>
</dl>

<dl>
<dd>

**language:** `*ticketing.EndUserDetailsRequestLanguage` 

The following subset of IETF language tags can be used to configure localization.

* `en` - en
* `de` - de
    
</dd>
</dl>

<dl>
<dd>

**areSyncsDisabled:** `*bool` — The boolean that indicates whether initial, periodic, and force syncs will be disabled.
    
</dd>
</dl>

<dl>
<dd>

**integrationSpecificConfig:** `map[string]any` — A JSON object containing integration-specific configuration options.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing LinkedAccounts
<details><summary><code>client.Ticketing.LinkedAccounts.List() -> *ticketing.PaginatedAccountDetailsAndActionsList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

List linked accounts for your organization.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.LinkedAccountsListRequest{
        Category: ticketing.LinkedAccountsListRequestCategoryAccounting.Ptr(),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EndUserEmailAddress: merge.String(
            "end_user_email_address",
        ),
        EndUserOrganizationName: merge.String(
            "end_user_organization_name",
        ),
        EndUserOriginId: merge.String(
            "end_user_origin_id",
        ),
        EndUserOriginIds: merge.String(
            "end_user_origin_ids",
        ),
        Id: merge.String(
            "id",
        ),
        Ids: merge.String(
            "ids",
        ),
        IncludeDuplicates: merge.Bool(
            true,
        ),
        IntegrationName: merge.String(
            "integration_name",
        ),
        IsTestAccount: merge.String(
            "is_test_account",
        ),
        PageSize: merge.Int(
            1,
        ),
        Status: merge.String(
            "status",
        ),
    }
client.Ticketing.LinkedAccounts.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**category:** `*ticketing.LinkedAccountsListRequestCategory` 

Options: `accounting`, `ats`, `crm`, `filestorage`, `hris`, `mktg`, `ticketing`

* `hris` - hris
* `ats` - ats
* `accounting` - accounting
* `ticketing` - ticketing
* `crm` - crm
* `mktg` - mktg
* `filestorage` - filestorage
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**endUserEmailAddress:** `*string` — If provided, will only return linked accounts associated with the given email address.
    
</dd>
</dl>

<dl>
<dd>

**endUserOrganizationName:** `*string` — If provided, will only return linked accounts associated with the given organization name.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginId:** `*string` — If provided, will only return linked accounts associated with the given origin ID.
    
</dd>
</dl>

<dl>
<dd>

**endUserOriginIds:** `*string` — Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
    
</dd>
</dl>

<dl>
<dd>

**id:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
    
</dd>
</dl>

<dl>
<dd>

**includeDuplicates:** `*bool` — If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
    
</dd>
</dl>

<dl>
<dd>

**integrationName:** `*string` — If provided, will only return linked accounts associated with the given integration name.
    
</dd>
</dl>

<dl>
<dd>

**isTestAccount:** `*string` — If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**status:** `*string` — Filter by status. Options: `COMPLETE`, `IDLE`, `INCOMPLETE`, `RELINK_NEEDED`
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Passthrough
<details><summary><code>client.Ticketing.Passthrough.Create(request) -> *ticketing.RemoteResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Pull data from an endpoint not currently supported by Merge.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.DataPassthroughRequest{
        Method: ticketing.MethodEnumGet,
        Path: "/scooters",
    }
client.Ticketing.Passthrough.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*ticketing.DataPassthroughRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Projects
<details><summary><code>client.Ticketing.Projects.List() -> *ticketing.PaginatedProjectList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Project` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.ProjectsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ticketing.Projects.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Projects.Retrieve(Id) -> *ticketing.Project</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Project` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.ProjectsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Projects.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Projects.UsersList(ParentId) -> *ticketing.PaginatedUserList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `User` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.ProjectsUsersListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Ticketing.Projects.UsersList(
        context.TODO(),
        "parent_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**parentId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.ProjectsUsersListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing RegenerateKey
<details><summary><code>client.Ticketing.RegenerateKey.Create(request) -> *ticketing.RemoteKey</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Exchange remote keys.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.RemoteKeyForRegenerationRequest{
        Name: "Remote Deployment Key 1",
    }
client.Ticketing.RegenerateKey.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**name:** `string` — The name of the remote key
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Roles
<details><summary><code>client.Ticketing.Roles.List() -> *ticketing.PaginatedRoleList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Role` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.RolesListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ticketing.Roles.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Roles.Retrieve(Id) -> *ticketing.Role</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Role` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.RolesRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Roles.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing SyncStatus
<details><summary><code>client.Ticketing.SyncStatus.List() -> *ticketing.PaginatedSyncStatusList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Get sync status for the current sync and the most recently finished sync. `last_sync_start` represents the most recent time any sync began. `last_sync_finished` represents the most recent time any sync completed. These timestamps may correspond to different sync instances which may result in a sync start time being later than a separate sync completed time. To ensure you are retrieving the latest available data reference the `last_sync_finished` timestamp where `last_sync_result` is `DONE`. Possible values for `status` and `last_sync_result` are `DISABLED`, `DONE`, `FAILED`, `PARTIALLY_SYNCED`, `PAUSED`, `SYNCING`. Learn more about sync status in our [Help Center](https://help.merge.dev/en/articles/8184193-merge-sync-statuses).
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.SyncStatusListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Ticketing.SyncStatus.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing ForceResync
<details><summary><code>client.Ticketing.ForceResync.SyncStatusResyncCreate() -> []*ticketing.SyncStatus</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Force re-sync of all models. This endpoint is available for monthly, quarterly, and highest sync frequency customers on the Professional or Enterprise plans. Doing so will consume a sync credit for the relevant linked account. Force re-syncs can also be triggered manually in the Merge Dashboard and is available for all customers.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.ForceResync.SyncStatusResyncCreate(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Tags
<details><summary><code>client.Ticketing.Tags.List() -> *ticketing.PaginatedTagList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Tag` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TagsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ticketing.Tags.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tags.Retrieve(Id) -> *ticketing.Tag</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Tag` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TagsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Tags.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Teams
<details><summary><code>client.Ticketing.Teams.List() -> *ticketing.PaginatedTeamList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Team` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TeamsListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
    }
client.Ticketing.Teams.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Teams.Retrieve(Id) -> *ticketing.Team</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Team` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TeamsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Teams.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Tickets
<details><summary><code>client.Ticketing.Tickets.List() -> *ticketing.PaginatedTicketList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Ticket` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketsListRequest{
        AccountId: merge.String(
            "account_id",
        ),
        AssigneeIds: merge.String(
            "assignee_ids",
        ),
        CollectionIds: merge.String(
            "collection_ids",
        ),
        CompletedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CompletedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ContactId: merge.String(
            "contact_id",
        ),
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatorId: merge.String(
            "creator_id",
        ),
        CreatorIds: merge.String(
            "creator_ids",
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        DueAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        DueBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Name: merge.String(
            "name",
        ),
        PageSize: merge.Int(
            1,
        ),
        ParentTicketId: merge.String(
            "parent_ticket_id",
        ),
        Priority: ticketing.TicketsListRequestPriorityHigh.Ptr(),
        RemoteCreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        RemoteCreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        RemoteFields: ticketing.TicketsListRequestRemoteFieldsPriority.Ptr(),
        RemoteId: merge.String(
            "remote_id",
        ),
        RemoteUpdatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        RemoteUpdatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ShowEnumOrigins: ticketing.TicketsListRequestShowEnumOriginsPriority.Ptr(),
        Status: ticketing.TicketsListRequestStatusEmpty.Ptr(),
        Tags: merge.String(
            "tags",
        ),
        TicketType: merge.String(
            "ticket_type",
        ),
        TicketUrl: merge.String(
            "ticket_url",
        ),
    }
client.Ticketing.Tickets.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**accountId:** `*string` — If provided, will only return tickets for this account.
    
</dd>
</dl>

<dl>
<dd>

**assigneeIds:** `*string` — If provided, will only return tickets assigned to the assignee_ids; multiple assignee_ids can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**collectionIds:** `*string` — If provided, will only return tickets assigned to the collection_ids; multiple collection_ids can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**completedAfter:** `*time.Time` — If provided, will only return tickets completed after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**completedBefore:** `*time.Time` — If provided, will only return tickets completed before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**contactId:** `*string` — If provided, will only return tickets for this contact.
    
</dd>
</dl>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**creatorId:** `*string` — If provided, will only return tickets created by this creator_id.
    
</dd>
</dl>

<dl>
<dd>

**creatorIds:** `*string` — If provided, will only return tickets created by the creator_ids; multiple creator_ids can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**dueAfter:** `*time.Time` — If provided, will only return tickets due after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**dueBefore:** `*time.Time` — If provided, will only return tickets due before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.TicketsListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**name:** `*string` — If provided, will only return tickets with this name.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**parentTicketId:** `*string` — If provided, will only return sub tickets of the parent_ticket_id.
    
</dd>
</dl>

<dl>
<dd>

**priority:** `*ticketing.TicketsListRequestPriority` 

If provided, will only return tickets of this priority.

* `URGENT` - URGENT
* `HIGH` - HIGH
* `NORMAL` - NORMAL
* `LOW` - LOW
    
</dd>
</dl>

<dl>
<dd>

**remoteCreatedAfter:** `*time.Time` — If provided, will only return tickets created in the third party platform after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**remoteCreatedBefore:** `*time.Time` — If provided, will only return tickets created in the third party platform before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*ticketing.TicketsListRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**remoteUpdatedAfter:** `*time.Time` — If provided, will only return tickets updated in the third party platform after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**remoteUpdatedBefore:** `*time.Time` — If provided, will only return tickets updated in the third party platform before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*ticketing.TicketsListRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>

<dl>
<dd>

**status:** `*ticketing.TicketsListRequestStatus` — If provided, will only return tickets of this status.
    
</dd>
</dl>

<dl>
<dd>

**tags:** `*string` — If provided, will only return tickets matching the tags; multiple tags can be separated by commas.
    
</dd>
</dl>

<dl>
<dd>

**ticketType:** `*string` — If provided, will only return tickets of this type.
    
</dd>
</dl>

<dl>
<dd>

**ticketUrl:** `*string` — If provided, will only return tickets where the URL matches or contains the substring
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tickets.Create(request) -> *ticketing.TicketResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `Ticket` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ticketing.TicketRequest{},
    }
client.Ticketing.Tickets.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ticketing.TicketRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tickets.Retrieve(Id) -> *ticketing.Ticket</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `Ticket` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketsRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeRemoteFields: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        RemoteFields: ticketing.TicketsRetrieveRequestRemoteFieldsPriority.Ptr(),
        ShowEnumOrigins: ticketing.TicketsRetrieveRequestShowEnumOriginsPriority.Ptr(),
    }
client.Ticketing.Tickets.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.TicketsRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteFields:** `*bool` — Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**remoteFields:** `*ticketing.TicketsRetrieveRequestRemoteFields` — Deprecated. Use show_enum_origins.
    
</dd>
</dl>

<dl>
<dd>

**showEnumOrigins:** `*ticketing.TicketsRetrieveRequestShowEnumOrigins` — A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tickets.PartialUpdate(Id, request) -> *ticketing.TicketResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Updates a `Ticket` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.PatchedTicketEndpointRequest{
        IsDebugMode: merge.Bool(
            true,
        ),
        RunAsync: merge.Bool(
            true,
        ),
        Model: &ticketing.PatchedTicketRequest{},
    }
client.Ticketing.Tickets.PartialUpdate(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isDebugMode:** `*bool` — Whether to include debug fields (such as log file links) in the response.
    
</dd>
</dl>

<dl>
<dd>

**runAsync:** `*bool` — Whether or not third-party updates should be run asynchronously.
    
</dd>
</dl>

<dl>
<dd>

**model:** `*ticketing.PatchedTicketRequest` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tickets.ViewersList(TicketId) -> *ticketing.PaginatedViewerList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `Viewer` objects that point to a User id or Team id that is either an assignee or viewer on a `Ticket` with the given id. [Learn more.](https://help.merge.dev/en/articles/10333658-ticketing-access-control-list-acls)
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketsViewersListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Ticketing.Tickets.ViewersList(
        context.TODO(),
        "ticket_id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**ticketId:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.TicketsViewersListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tickets.MetaPatchRetrieve(Id) -> *ticketing.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Ticket` PATCHs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.Tickets.MetaPatchRetrieve(
        context.TODO(),
        "id",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tickets.MetaPostRetrieve() -> *ticketing.MetaResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns metadata for `Ticket` POSTs.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketsMetaPostRetrieveRequest{
        CollectionId: merge.String(
            "collection_id",
        ),
        TicketType: merge.String(
            "ticket_type",
        ),
    }
client.Ticketing.Tickets.MetaPostRetrieve(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**collectionId:** `*string` — If provided, will only return tickets for this collection.
    
</dd>
</dl>

<dl>
<dd>

**ticketType:** `*string` — If provided, will only return tickets for this ticket type.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Tickets.RemoteFieldClassesList() -> *ticketing.PaginatedRemoteFieldClassList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `RemoteFieldClass` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.TicketsRemoteFieldClassesListRequest{
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        Ids: merge.String(
            "ids",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        IsCommonModelField: merge.Bool(
            true,
        ),
        IsCustom: merge.Bool(
            true,
        ),
        PageSize: merge.Int(
            1,
        ),
    }
client.Ticketing.Tickets.RemoteFieldClassesList(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**ids:** `*string` — If provided, will only return remote field classes with the `ids` in this list
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**isCommonModelField:** `*bool` — If provided, will only return remote field classes with this is_common_model_field value
    
</dd>
</dl>

<dl>
<dd>

**isCustom:** `*bool` — If provided, will only return remote fields classes with this is_custom value
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing Users
<details><summary><code>client.Ticketing.Users.List() -> *ticketing.PaginatedUserList</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `User` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.UsersListRequest{
        CreatedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        CreatedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        Cursor: merge.String(
            "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw",
        ),
        EmailAddress: merge.String(
            "email_address",
        ),
        IncludeDeletedData: merge.Bool(
            true,
        ),
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
        ModifiedAfter: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        ModifiedBefore: merge.Time(
            merge.MustParseDateTime(
                "2024-01-15T09:30:00Z",
            ),
        ),
        PageSize: merge.Int(
            1,
        ),
        RemoteId: merge.String(
            "remote_id",
        ),
        Team: merge.String(
            "team",
        ),
    }
client.Ticketing.Users.List(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**createdAfter:** `*time.Time` — If provided, will only return objects created after this datetime.
    
</dd>
</dl>

<dl>
<dd>

**createdBefore:** `*time.Time` — If provided, will only return objects created before this datetime.
    
</dd>
</dl>

<dl>
<dd>

**cursor:** `*string` — The pagination cursor value.
    
</dd>
</dl>

<dl>
<dd>

**emailAddress:** `*string` — If provided, will only return users with emails equal to this value (case insensitive).
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.UsersListRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeDeletedData:** `*bool` — Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>

<dl>
<dd>

**modifiedAfter:** `*time.Time` — If provided, only objects synced by Merge after this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**modifiedBefore:** `*time.Time` — If provided, only objects synced by Merge before this date time will be returned.
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*int` — Number of results to return per page.
    
</dd>
</dl>

<dl>
<dd>

**remoteId:** `*string` — The API provider's ID for the given object.
    
</dd>
</dl>

<dl>
<dd>

**team:** `*string` — If provided, will only return users matching in this team.
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.Users.Retrieve(Id) -> *ticketing.User</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a `User` object with the given `id`.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.UsersRetrieveRequest{
        IncludeRemoteData: merge.Bool(
            true,
        ),
        IncludeShellData: merge.Bool(
            true,
        ),
    }
client.Ticketing.Users.Retrieve(
        context.TODO(),
        "id",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**id:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**expand:** `*ticketing.UsersRetrieveRequestExpandItem` — Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
    
</dd>
</dl>

<dl>
<dd>

**includeRemoteData:** `*bool` — Whether to include the original data Merge fetched from the third-party to produce these models.
    
</dd>
</dl>

<dl>
<dd>

**includeShellData:** `*bool` — Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Ticketing WebhookReceivers
<details><summary><code>client.Ticketing.WebhookReceivers.List() -> []*ticketing.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Returns a list of `WebhookReceiver` objects.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Ticketing.WebhookReceivers.List(
        context.TODO(),
    )
}
```
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Ticketing.WebhookReceivers.Create(request) -> *ticketing.WebhookReceiver</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

Creates a `WebhookReceiver` object with the given values.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &ticketing.WebhookReceiverRequest{
        Event: "event",
        IsActive: true,
    }
client.Ticketing.WebhookReceivers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**event:** `string` 
    
</dd>
</dl>

<dl>
<dd>

**isActive:** `bool` 
    
</dd>
</dl>

<dl>
<dd>

**key:** `*string` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>
