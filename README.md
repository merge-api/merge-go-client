# Merge Go Library

The Merge Go library provides convenient access to the Merge API from Go.

[![go shield](https://img.shields.io/badge/go-docs-blue)](https://pkg.go.dev/github.com/merge-api/merge-go-client)

## Requirements

This module requires Go version >= 1.19.

## Installation

Run the following command to use the Merge Go library in your Go module:

```sh
go get github.com/merge-api/merge-go-client
```

## Instantiation

```go
import (
  "context"
  "fmt"

  merge "github.com/merge-api/merge-go-client"
  mergeclient "github.com/merge-api/merge-go-client/client"
)

client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)
```

## Categories

This SDK contains both the ATS, HRIS, CRM, Ticketing, and Accounting categories. Even if
you do not plan on using more than one Merge API category right now, the SDK provides
upgrade-flexibility in case you find new Merge API categories useful in the future.

Each category is namespaced:

```go
client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)

client.Ats(). // APIs specific to the ATS category

client.Hris(). // APIs specific to the HRIS category
```

## Usage

### Create Link Token

```go
import (
  "context"
  "fmt"

  merge "github.com/fern-api/merge-go"
  mergeclient "github.com/fern-api/merge-go/client"
  "github.com/fern-api/merge-go/ats"
)

client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)
linkTokenResponse, err := client.Ats().LinkToken().Create(
  context.TODO(),
  &ats.EndUserDetailsRequest{
    EndUserEmailAddress:     "john.smith@gmail.com",
    EndUserOrganizationName: "acme",
    EndUserOriginId:         "1234",
    Categories:              []ats.CategoriesEnum{ats.CategoriesEnumAts},
    LinkExpiryMins:          merge.Int(30),
  },
)
if err != nil {
  return err
}
fmt.Printf("Created link token %s\n", linkTokenResponse.LinkToken)
```

### Get Employee

```go
import (
  "context"
  "fmt"

  merge "github.com/fern-api/merge-go"
  mergeclient "github.com/fern-api/merge-go/client"
  "github.com/fern-api/merge-go/hris"
)

client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)
employee, err := client.Hris().Employees().Retrieve(
  context.TODO(),
  "0958cbc6-6040-430a-848e-aafacbadf4ae",
  &hris.EmployeesRetrieveRequest{
    IncludeRemoteData: merge.Bool(true),
  },
)
if err != nil {
  return err
}
fmt.Printf("Retrieved employee with ID %q\n", *employee.Id)
```

### Get Candidate

```go
import (
  "context"
  "fmt"

  merge "github.com/fern-api/merge-go"
  mergeclient "github.com/fern-api/merge-go/client"
  "github.com/fern-api/merge-go/ats"
)

client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)
candidate, err := client.Ats().Candidates().Retrieve(
  context.TODO(),
  "521b18c2-4d01-4297-b451-19858d07c133",
  &ats.CandidatesRetrieveRequest{
    IncludeRemoteData: merge.Bool(true),
  },
)
if err != nil {
  return err
}
fmt.Printf("Retrieved candidate with ID %q\n", *candidate.Id)
```

### Filter Candidate

```go
import (
  "context"
  "fmt"
  "time"

  merge "github.com/fern-api/merge-go"
  mergeclient "github.com/fern-api/merge-go/client"
  "github.com/fern-api/merge-go/ats"
)

client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)
candidateList, err := client.Ats().Candidates().List(
  context.TODO(),
  &ats.CandidatesListRequest{
    CreatedBefore: merge.Time(time.Now()),
  },
)
if err != nil {
  return err
}
fmt.Printf("Retrieved %d candidates\n", len(candidateList.Results))
```

### Create Ticket

```go
import (
  "context"
  "fmt"
  "time"

  merge "github.com/fern-api/merge-go"
  mergeclient "github.com/fern-api/merge-go/client"
  "github.com/fern-api/merge-go/ticketing"
)

client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHeaderAccountToken("<YOUR_ACCOUNT_TOKEN>"),
)
ticketResponse, err := client.Ticketing().Tickets().Create(
  context.TODO(),
  &ticketing.TicketEndpointRequest{
    Model: &ticketing.TicketRequest{
      Name: merge.String("Please add more integrations"),
      Assignees: []string{
        "17a54124-287f-494d-965e-3c5b330c9a68",
      },
      Creator: merge.String("3fa85f64-5717-4562-b3fc-2c963f66afa6"),
      DueDate: merge.Time(time.Now().Add(24 * time.Hour)),
      Status:  ticketing.NewTicketRequestStatusFromTicketStatusEnum(ticketing.TicketStatusEnumOpen),
    },
  },
)
if err != nil {
  return err
}
fmt.Printf("Successfully created ticket with ID %q\n", *ticketResponse.Model.Id)
```

## Timeouts

Setting a timeout for each individual request is as simple as using the standard
`context` library. Setting a one second timeout for an individual API call looks
like the following:

```go
ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
defer cancel()
employeeList, err := client.Hris().Employees().List(
  ctx,
  &hris.EmployeesListRequest{
    CreatedBefore: merge.Time(time.Now()),
  },
)
if err != nil {
  return err
}
```

## Client Options

A variety of client options are included to adapt the behavior of the library, which includes
configuring authorization tokens to be sent on every request, or providing your own instrumented
`*http.Client`. Both of these options are shown below:

```go
client := mergeclient.NewClient(
  mergeclient.ClientWithAuthApiKey("<YOUR_API_KEY>"),
  mergeclient.ClientWithHTTPClient(
    &http.Client{
      Timeout: 5 * time.Second,
    },
  ),
)
```

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

## Errors

Structured error types are returned from API calls that return non-success status codes. For example,
you can check if the error was due to a bad request (i.e. status code 400) with the following:

```go
employeeList, err := client.Hris().Employees().List(
  context.TODO(),
  &hris.EmployeesListRequest{
    CreatedBefore: merge.Time(time.Now()),
  },
)
if err != nil {
  if apiErr, ok := err.(*core.APIError); ok && apiErr.StatusCode == http.StatusBadRequest {
    // Do something with the bad request ...
  }
  return err
}
```

These errors are also compatible with the `errors.Is` and `errors.As` APIs, so you can access the error
like so:

```go
employeeList, err := client.Hris().Employees().List(
  context.TODO(),
  &hris.EmployeesListRequest{
    CreatedBefore: merge.Time(time.Now()),
  },
)
if err != nil {
  var apiErr *core.APIError
  if errors.As(err, apiError); ok {
    switch apiErr.StatusCode {
      case http.StatusBadRequest:
        // Do something with the bad request ...
    }
  }
  return err
}
```

If you'd like to wrap the errors with additional information and still retain the ability to access the type
with `errors.Is` and `errors.As`, you can use the `%w` directive:

```go
employeeList, err := client.Hris().Employees().List(
  context.TODO(),
  &hris.EmployeesListRequest{
    CreatedBefore: merge.Time(time.Now()),
  },
)
if err != nil {
  return fmt.Errorf("failed to list employees: %w", err)
}
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically. Additions
made directly to this library would have to be moved over to our generation code, otherwise they would be
overwritten upon the next generated release. Feel free to open a PR as a proof of concept, but know that we
will not be able to merge it as-is. We suggest opening an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
