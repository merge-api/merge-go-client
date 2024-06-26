// This file was auto-generated by Fern from our API Definition.

package ats

import (
	fmt "fmt"
	time "time"
)

type UpdateApplicationStageRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool `json:"-"`
	// The interview stage to move the application to.
	JobInterviewStage *string `json:"job_interview_stage,omitempty"`
	RemoteUserId      *string `json:"remote_user_id,omitempty"`
}

type ApplicationEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync     *bool               `json:"-"`
	Model        *ApplicationRequest `json:"model,omitempty"`
	RemoteUserId string              `json:"remote_user_id"`
}

type ApplicationsListRequest struct {
	// If provided, will only return applications for this candidate.
	CandidateId *string `json:"-"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// If provided, will only return applications credited to this user.
	CreditedToId *string `json:"-"`
	// If provided, will only return applications at this interview stage.
	CurrentStageId *string `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *ApplicationsListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// If provided, will only return applications for this job.
	JobId *string `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// If provided, will only return applications with this reject reason.
	RejectReasonId *string `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return applications with this source.
	Source *string `json:"-"`
}

type ApplicationsMetaPostRetrieveRequest struct {
	// The template ID associated with the nested application in the request.
	ApplicationRemoteTemplateId *string `json:"-"`
}

type ApplicationsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *ApplicationsRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}

type ApplicationsListRequestExpand string

const (
	ApplicationsListRequestExpandCandidate                                            ApplicationsListRequestExpand = "candidate"
	ApplicationsListRequestExpandCandidateCreditedTo                                  ApplicationsListRequestExpand = "candidate,credited_to"
	ApplicationsListRequestExpandCandidateCreditedToCurrentStage                      ApplicationsListRequestExpand = "candidate,credited_to,current_stage"
	ApplicationsListRequestExpandCandidateCreditedToCurrentStageRejectReason          ApplicationsListRequestExpand = "candidate,credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandCandidateCreditedToRejectReason                      ApplicationsListRequestExpand = "candidate,credited_to,reject_reason"
	ApplicationsListRequestExpandCandidateCurrentStage                                ApplicationsListRequestExpand = "candidate,current_stage"
	ApplicationsListRequestExpandCandidateCurrentStageRejectReason                    ApplicationsListRequestExpand = "candidate,current_stage,reject_reason"
	ApplicationsListRequestExpandCandidateJob                                         ApplicationsListRequestExpand = "candidate,job"
	ApplicationsListRequestExpandCandidateJobCreditedTo                               ApplicationsListRequestExpand = "candidate,job,credited_to"
	ApplicationsListRequestExpandCandidateJobCreditedToCurrentStage                   ApplicationsListRequestExpand = "candidate,job,credited_to,current_stage"
	ApplicationsListRequestExpandCandidateJobCreditedToCurrentStageRejectReason       ApplicationsListRequestExpand = "candidate,job,credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandCandidateJobCreditedToRejectReason                   ApplicationsListRequestExpand = "candidate,job,credited_to,reject_reason"
	ApplicationsListRequestExpandCandidateJobCurrentStage                             ApplicationsListRequestExpand = "candidate,job,current_stage"
	ApplicationsListRequestExpandCandidateJobCurrentStageRejectReason                 ApplicationsListRequestExpand = "candidate,job,current_stage,reject_reason"
	ApplicationsListRequestExpandCandidateJobRejectReason                             ApplicationsListRequestExpand = "candidate,job,reject_reason"
	ApplicationsListRequestExpandCandidateRejectReason                                ApplicationsListRequestExpand = "candidate,reject_reason"
	ApplicationsListRequestExpandCreditedTo                                           ApplicationsListRequestExpand = "credited_to"
	ApplicationsListRequestExpandCreditedToCurrentStage                               ApplicationsListRequestExpand = "credited_to,current_stage"
	ApplicationsListRequestExpandCreditedToCurrentStageRejectReason                   ApplicationsListRequestExpand = "credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandCreditedToRejectReason                               ApplicationsListRequestExpand = "credited_to,reject_reason"
	ApplicationsListRequestExpandCurrentStage                                         ApplicationsListRequestExpand = "current_stage"
	ApplicationsListRequestExpandCurrentStageRejectReason                             ApplicationsListRequestExpand = "current_stage,reject_reason"
	ApplicationsListRequestExpandJob                                                  ApplicationsListRequestExpand = "job"
	ApplicationsListRequestExpandJobCreditedTo                                        ApplicationsListRequestExpand = "job,credited_to"
	ApplicationsListRequestExpandJobCreditedToCurrentStage                            ApplicationsListRequestExpand = "job,credited_to,current_stage"
	ApplicationsListRequestExpandJobCreditedToCurrentStageRejectReason                ApplicationsListRequestExpand = "job,credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandJobCreditedToRejectReason                            ApplicationsListRequestExpand = "job,credited_to,reject_reason"
	ApplicationsListRequestExpandJobCurrentStage                                      ApplicationsListRequestExpand = "job,current_stage"
	ApplicationsListRequestExpandJobCurrentStageRejectReason                          ApplicationsListRequestExpand = "job,current_stage,reject_reason"
	ApplicationsListRequestExpandJobRejectReason                                      ApplicationsListRequestExpand = "job,reject_reason"
	ApplicationsListRequestExpandOffers                                               ApplicationsListRequestExpand = "offers"
	ApplicationsListRequestExpandOffersCandidate                                      ApplicationsListRequestExpand = "offers,candidate"
	ApplicationsListRequestExpandOffersCandidateCreditedTo                            ApplicationsListRequestExpand = "offers,candidate,credited_to"
	ApplicationsListRequestExpandOffersCandidateCreditedToCurrentStage                ApplicationsListRequestExpand = "offers,candidate,credited_to,current_stage"
	ApplicationsListRequestExpandOffersCandidateCreditedToCurrentStageRejectReason    ApplicationsListRequestExpand = "offers,candidate,credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersCandidateCreditedToRejectReason                ApplicationsListRequestExpand = "offers,candidate,credited_to,reject_reason"
	ApplicationsListRequestExpandOffersCandidateCurrentStage                          ApplicationsListRequestExpand = "offers,candidate,current_stage"
	ApplicationsListRequestExpandOffersCandidateCurrentStageRejectReason              ApplicationsListRequestExpand = "offers,candidate,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersCandidateJob                                   ApplicationsListRequestExpand = "offers,candidate,job"
	ApplicationsListRequestExpandOffersCandidateJobCreditedTo                         ApplicationsListRequestExpand = "offers,candidate,job,credited_to"
	ApplicationsListRequestExpandOffersCandidateJobCreditedToCurrentStage             ApplicationsListRequestExpand = "offers,candidate,job,credited_to,current_stage"
	ApplicationsListRequestExpandOffersCandidateJobCreditedToCurrentStageRejectReason ApplicationsListRequestExpand = "offers,candidate,job,credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersCandidateJobCreditedToRejectReason             ApplicationsListRequestExpand = "offers,candidate,job,credited_to,reject_reason"
	ApplicationsListRequestExpandOffersCandidateJobCurrentStage                       ApplicationsListRequestExpand = "offers,candidate,job,current_stage"
	ApplicationsListRequestExpandOffersCandidateJobCurrentStageRejectReason           ApplicationsListRequestExpand = "offers,candidate,job,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersCandidateJobRejectReason                       ApplicationsListRequestExpand = "offers,candidate,job,reject_reason"
	ApplicationsListRequestExpandOffersCandidateRejectReason                          ApplicationsListRequestExpand = "offers,candidate,reject_reason"
	ApplicationsListRequestExpandOffersCreditedTo                                     ApplicationsListRequestExpand = "offers,credited_to"
	ApplicationsListRequestExpandOffersCreditedToCurrentStage                         ApplicationsListRequestExpand = "offers,credited_to,current_stage"
	ApplicationsListRequestExpandOffersCreditedToCurrentStageRejectReason             ApplicationsListRequestExpand = "offers,credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersCreditedToRejectReason                         ApplicationsListRequestExpand = "offers,credited_to,reject_reason"
	ApplicationsListRequestExpandOffersCurrentStage                                   ApplicationsListRequestExpand = "offers,current_stage"
	ApplicationsListRequestExpandOffersCurrentStageRejectReason                       ApplicationsListRequestExpand = "offers,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersJob                                            ApplicationsListRequestExpand = "offers,job"
	ApplicationsListRequestExpandOffersJobCreditedTo                                  ApplicationsListRequestExpand = "offers,job,credited_to"
	ApplicationsListRequestExpandOffersJobCreditedToCurrentStage                      ApplicationsListRequestExpand = "offers,job,credited_to,current_stage"
	ApplicationsListRequestExpandOffersJobCreditedToCurrentStageRejectReason          ApplicationsListRequestExpand = "offers,job,credited_to,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersJobCreditedToRejectReason                      ApplicationsListRequestExpand = "offers,job,credited_to,reject_reason"
	ApplicationsListRequestExpandOffersJobCurrentStage                                ApplicationsListRequestExpand = "offers,job,current_stage"
	ApplicationsListRequestExpandOffersJobCurrentStageRejectReason                    ApplicationsListRequestExpand = "offers,job,current_stage,reject_reason"
	ApplicationsListRequestExpandOffersJobRejectReason                                ApplicationsListRequestExpand = "offers,job,reject_reason"
	ApplicationsListRequestExpandOffersRejectReason                                   ApplicationsListRequestExpand = "offers,reject_reason"
	ApplicationsListRequestExpandRejectReason                                         ApplicationsListRequestExpand = "reject_reason"
)

func NewApplicationsListRequestExpandFromString(s string) (ApplicationsListRequestExpand, error) {
	switch s {
	case "candidate":
		return ApplicationsListRequestExpandCandidate, nil
	case "candidate,credited_to":
		return ApplicationsListRequestExpandCandidateCreditedTo, nil
	case "candidate,credited_to,current_stage":
		return ApplicationsListRequestExpandCandidateCreditedToCurrentStage, nil
	case "candidate,credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandCandidateCreditedToCurrentStageRejectReason, nil
	case "candidate,credited_to,reject_reason":
		return ApplicationsListRequestExpandCandidateCreditedToRejectReason, nil
	case "candidate,current_stage":
		return ApplicationsListRequestExpandCandidateCurrentStage, nil
	case "candidate,current_stage,reject_reason":
		return ApplicationsListRequestExpandCandidateCurrentStageRejectReason, nil
	case "candidate,job":
		return ApplicationsListRequestExpandCandidateJob, nil
	case "candidate,job,credited_to":
		return ApplicationsListRequestExpandCandidateJobCreditedTo, nil
	case "candidate,job,credited_to,current_stage":
		return ApplicationsListRequestExpandCandidateJobCreditedToCurrentStage, nil
	case "candidate,job,credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandCandidateJobCreditedToCurrentStageRejectReason, nil
	case "candidate,job,credited_to,reject_reason":
		return ApplicationsListRequestExpandCandidateJobCreditedToRejectReason, nil
	case "candidate,job,current_stage":
		return ApplicationsListRequestExpandCandidateJobCurrentStage, nil
	case "candidate,job,current_stage,reject_reason":
		return ApplicationsListRequestExpandCandidateJobCurrentStageRejectReason, nil
	case "candidate,job,reject_reason":
		return ApplicationsListRequestExpandCandidateJobRejectReason, nil
	case "candidate,reject_reason":
		return ApplicationsListRequestExpandCandidateRejectReason, nil
	case "credited_to":
		return ApplicationsListRequestExpandCreditedTo, nil
	case "credited_to,current_stage":
		return ApplicationsListRequestExpandCreditedToCurrentStage, nil
	case "credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandCreditedToCurrentStageRejectReason, nil
	case "credited_to,reject_reason":
		return ApplicationsListRequestExpandCreditedToRejectReason, nil
	case "current_stage":
		return ApplicationsListRequestExpandCurrentStage, nil
	case "current_stage,reject_reason":
		return ApplicationsListRequestExpandCurrentStageRejectReason, nil
	case "job":
		return ApplicationsListRequestExpandJob, nil
	case "job,credited_to":
		return ApplicationsListRequestExpandJobCreditedTo, nil
	case "job,credited_to,current_stage":
		return ApplicationsListRequestExpandJobCreditedToCurrentStage, nil
	case "job,credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandJobCreditedToCurrentStageRejectReason, nil
	case "job,credited_to,reject_reason":
		return ApplicationsListRequestExpandJobCreditedToRejectReason, nil
	case "job,current_stage":
		return ApplicationsListRequestExpandJobCurrentStage, nil
	case "job,current_stage,reject_reason":
		return ApplicationsListRequestExpandJobCurrentStageRejectReason, nil
	case "job,reject_reason":
		return ApplicationsListRequestExpandJobRejectReason, nil
	case "offers":
		return ApplicationsListRequestExpandOffers, nil
	case "offers,candidate":
		return ApplicationsListRequestExpandOffersCandidate, nil
	case "offers,candidate,credited_to":
		return ApplicationsListRequestExpandOffersCandidateCreditedTo, nil
	case "offers,candidate,credited_to,current_stage":
		return ApplicationsListRequestExpandOffersCandidateCreditedToCurrentStage, nil
	case "offers,candidate,credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateCreditedToCurrentStageRejectReason, nil
	case "offers,candidate,credited_to,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateCreditedToRejectReason, nil
	case "offers,candidate,current_stage":
		return ApplicationsListRequestExpandOffersCandidateCurrentStage, nil
	case "offers,candidate,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateCurrentStageRejectReason, nil
	case "offers,candidate,job":
		return ApplicationsListRequestExpandOffersCandidateJob, nil
	case "offers,candidate,job,credited_to":
		return ApplicationsListRequestExpandOffersCandidateJobCreditedTo, nil
	case "offers,candidate,job,credited_to,current_stage":
		return ApplicationsListRequestExpandOffersCandidateJobCreditedToCurrentStage, nil
	case "offers,candidate,job,credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateJobCreditedToCurrentStageRejectReason, nil
	case "offers,candidate,job,credited_to,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateJobCreditedToRejectReason, nil
	case "offers,candidate,job,current_stage":
		return ApplicationsListRequestExpandOffersCandidateJobCurrentStage, nil
	case "offers,candidate,job,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateJobCurrentStageRejectReason, nil
	case "offers,candidate,job,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateJobRejectReason, nil
	case "offers,candidate,reject_reason":
		return ApplicationsListRequestExpandOffersCandidateRejectReason, nil
	case "offers,credited_to":
		return ApplicationsListRequestExpandOffersCreditedTo, nil
	case "offers,credited_to,current_stage":
		return ApplicationsListRequestExpandOffersCreditedToCurrentStage, nil
	case "offers,credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersCreditedToCurrentStageRejectReason, nil
	case "offers,credited_to,reject_reason":
		return ApplicationsListRequestExpandOffersCreditedToRejectReason, nil
	case "offers,current_stage":
		return ApplicationsListRequestExpandOffersCurrentStage, nil
	case "offers,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersCurrentStageRejectReason, nil
	case "offers,job":
		return ApplicationsListRequestExpandOffersJob, nil
	case "offers,job,credited_to":
		return ApplicationsListRequestExpandOffersJobCreditedTo, nil
	case "offers,job,credited_to,current_stage":
		return ApplicationsListRequestExpandOffersJobCreditedToCurrentStage, nil
	case "offers,job,credited_to,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersJobCreditedToCurrentStageRejectReason, nil
	case "offers,job,credited_to,reject_reason":
		return ApplicationsListRequestExpandOffersJobCreditedToRejectReason, nil
	case "offers,job,current_stage":
		return ApplicationsListRequestExpandOffersJobCurrentStage, nil
	case "offers,job,current_stage,reject_reason":
		return ApplicationsListRequestExpandOffersJobCurrentStageRejectReason, nil
	case "offers,job,reject_reason":
		return ApplicationsListRequestExpandOffersJobRejectReason, nil
	case "offers,reject_reason":
		return ApplicationsListRequestExpandOffersRejectReason, nil
	case "reject_reason":
		return ApplicationsListRequestExpandRejectReason, nil
	}
	var t ApplicationsListRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApplicationsListRequestExpand) Ptr() *ApplicationsListRequestExpand {
	return &a
}

type ApplicationsRetrieveRequestExpand string

const (
	ApplicationsRetrieveRequestExpandCandidate                                            ApplicationsRetrieveRequestExpand = "candidate"
	ApplicationsRetrieveRequestExpandCandidateCreditedTo                                  ApplicationsRetrieveRequestExpand = "candidate,credited_to"
	ApplicationsRetrieveRequestExpandCandidateCreditedToCurrentStage                      ApplicationsRetrieveRequestExpand = "candidate,credited_to,current_stage"
	ApplicationsRetrieveRequestExpandCandidateCreditedToCurrentStageRejectReason          ApplicationsRetrieveRequestExpand = "candidate,credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandCandidateCreditedToRejectReason                      ApplicationsRetrieveRequestExpand = "candidate,credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandCandidateCurrentStage                                ApplicationsRetrieveRequestExpand = "candidate,current_stage"
	ApplicationsRetrieveRequestExpandCandidateCurrentStageRejectReason                    ApplicationsRetrieveRequestExpand = "candidate,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandCandidateJob                                         ApplicationsRetrieveRequestExpand = "candidate,job"
	ApplicationsRetrieveRequestExpandCandidateJobCreditedTo                               ApplicationsRetrieveRequestExpand = "candidate,job,credited_to"
	ApplicationsRetrieveRequestExpandCandidateJobCreditedToCurrentStage                   ApplicationsRetrieveRequestExpand = "candidate,job,credited_to,current_stage"
	ApplicationsRetrieveRequestExpandCandidateJobCreditedToCurrentStageRejectReason       ApplicationsRetrieveRequestExpand = "candidate,job,credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandCandidateJobCreditedToRejectReason                   ApplicationsRetrieveRequestExpand = "candidate,job,credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandCandidateJobCurrentStage                             ApplicationsRetrieveRequestExpand = "candidate,job,current_stage"
	ApplicationsRetrieveRequestExpandCandidateJobCurrentStageRejectReason                 ApplicationsRetrieveRequestExpand = "candidate,job,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandCandidateJobRejectReason                             ApplicationsRetrieveRequestExpand = "candidate,job,reject_reason"
	ApplicationsRetrieveRequestExpandCandidateRejectReason                                ApplicationsRetrieveRequestExpand = "candidate,reject_reason"
	ApplicationsRetrieveRequestExpandCreditedTo                                           ApplicationsRetrieveRequestExpand = "credited_to"
	ApplicationsRetrieveRequestExpandCreditedToCurrentStage                               ApplicationsRetrieveRequestExpand = "credited_to,current_stage"
	ApplicationsRetrieveRequestExpandCreditedToCurrentStageRejectReason                   ApplicationsRetrieveRequestExpand = "credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandCreditedToRejectReason                               ApplicationsRetrieveRequestExpand = "credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandCurrentStage                                         ApplicationsRetrieveRequestExpand = "current_stage"
	ApplicationsRetrieveRequestExpandCurrentStageRejectReason                             ApplicationsRetrieveRequestExpand = "current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandJob                                                  ApplicationsRetrieveRequestExpand = "job"
	ApplicationsRetrieveRequestExpandJobCreditedTo                                        ApplicationsRetrieveRequestExpand = "job,credited_to"
	ApplicationsRetrieveRequestExpandJobCreditedToCurrentStage                            ApplicationsRetrieveRequestExpand = "job,credited_to,current_stage"
	ApplicationsRetrieveRequestExpandJobCreditedToCurrentStageRejectReason                ApplicationsRetrieveRequestExpand = "job,credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandJobCreditedToRejectReason                            ApplicationsRetrieveRequestExpand = "job,credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandJobCurrentStage                                      ApplicationsRetrieveRequestExpand = "job,current_stage"
	ApplicationsRetrieveRequestExpandJobCurrentStageRejectReason                          ApplicationsRetrieveRequestExpand = "job,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandJobRejectReason                                      ApplicationsRetrieveRequestExpand = "job,reject_reason"
	ApplicationsRetrieveRequestExpandOffers                                               ApplicationsRetrieveRequestExpand = "offers"
	ApplicationsRetrieveRequestExpandOffersCandidate                                      ApplicationsRetrieveRequestExpand = "offers,candidate"
	ApplicationsRetrieveRequestExpandOffersCandidateCreditedTo                            ApplicationsRetrieveRequestExpand = "offers,candidate,credited_to"
	ApplicationsRetrieveRequestExpandOffersCandidateCreditedToCurrentStage                ApplicationsRetrieveRequestExpand = "offers,candidate,credited_to,current_stage"
	ApplicationsRetrieveRequestExpandOffersCandidateCreditedToCurrentStageRejectReason    ApplicationsRetrieveRequestExpand = "offers,candidate,credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCandidateCreditedToRejectReason                ApplicationsRetrieveRequestExpand = "offers,candidate,credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCandidateCurrentStage                          ApplicationsRetrieveRequestExpand = "offers,candidate,current_stage"
	ApplicationsRetrieveRequestExpandOffersCandidateCurrentStageRejectReason              ApplicationsRetrieveRequestExpand = "offers,candidate,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCandidateJob                                   ApplicationsRetrieveRequestExpand = "offers,candidate,job"
	ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedTo                         ApplicationsRetrieveRequestExpand = "offers,candidate,job,credited_to"
	ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedToCurrentStage             ApplicationsRetrieveRequestExpand = "offers,candidate,job,credited_to,current_stage"
	ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedToCurrentStageRejectReason ApplicationsRetrieveRequestExpand = "offers,candidate,job,credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedToRejectReason             ApplicationsRetrieveRequestExpand = "offers,candidate,job,credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCandidateJobCurrentStage                       ApplicationsRetrieveRequestExpand = "offers,candidate,job,current_stage"
	ApplicationsRetrieveRequestExpandOffersCandidateJobCurrentStageRejectReason           ApplicationsRetrieveRequestExpand = "offers,candidate,job,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCandidateJobRejectReason                       ApplicationsRetrieveRequestExpand = "offers,candidate,job,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCandidateRejectReason                          ApplicationsRetrieveRequestExpand = "offers,candidate,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCreditedTo                                     ApplicationsRetrieveRequestExpand = "offers,credited_to"
	ApplicationsRetrieveRequestExpandOffersCreditedToCurrentStage                         ApplicationsRetrieveRequestExpand = "offers,credited_to,current_stage"
	ApplicationsRetrieveRequestExpandOffersCreditedToCurrentStageRejectReason             ApplicationsRetrieveRequestExpand = "offers,credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCreditedToRejectReason                         ApplicationsRetrieveRequestExpand = "offers,credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandOffersCurrentStage                                   ApplicationsRetrieveRequestExpand = "offers,current_stage"
	ApplicationsRetrieveRequestExpandOffersCurrentStageRejectReason                       ApplicationsRetrieveRequestExpand = "offers,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersJob                                            ApplicationsRetrieveRequestExpand = "offers,job"
	ApplicationsRetrieveRequestExpandOffersJobCreditedTo                                  ApplicationsRetrieveRequestExpand = "offers,job,credited_to"
	ApplicationsRetrieveRequestExpandOffersJobCreditedToCurrentStage                      ApplicationsRetrieveRequestExpand = "offers,job,credited_to,current_stage"
	ApplicationsRetrieveRequestExpandOffersJobCreditedToCurrentStageRejectReason          ApplicationsRetrieveRequestExpand = "offers,job,credited_to,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersJobCreditedToRejectReason                      ApplicationsRetrieveRequestExpand = "offers,job,credited_to,reject_reason"
	ApplicationsRetrieveRequestExpandOffersJobCurrentStage                                ApplicationsRetrieveRequestExpand = "offers,job,current_stage"
	ApplicationsRetrieveRequestExpandOffersJobCurrentStageRejectReason                    ApplicationsRetrieveRequestExpand = "offers,job,current_stage,reject_reason"
	ApplicationsRetrieveRequestExpandOffersJobRejectReason                                ApplicationsRetrieveRequestExpand = "offers,job,reject_reason"
	ApplicationsRetrieveRequestExpandOffersRejectReason                                   ApplicationsRetrieveRequestExpand = "offers,reject_reason"
	ApplicationsRetrieveRequestExpandRejectReason                                         ApplicationsRetrieveRequestExpand = "reject_reason"
)

func NewApplicationsRetrieveRequestExpandFromString(s string) (ApplicationsRetrieveRequestExpand, error) {
	switch s {
	case "candidate":
		return ApplicationsRetrieveRequestExpandCandidate, nil
	case "candidate,credited_to":
		return ApplicationsRetrieveRequestExpandCandidateCreditedTo, nil
	case "candidate,credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandCandidateCreditedToCurrentStage, nil
	case "candidate,credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateCreditedToCurrentStageRejectReason, nil
	case "candidate,credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateCreditedToRejectReason, nil
	case "candidate,current_stage":
		return ApplicationsRetrieveRequestExpandCandidateCurrentStage, nil
	case "candidate,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateCurrentStageRejectReason, nil
	case "candidate,job":
		return ApplicationsRetrieveRequestExpandCandidateJob, nil
	case "candidate,job,credited_to":
		return ApplicationsRetrieveRequestExpandCandidateJobCreditedTo, nil
	case "candidate,job,credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandCandidateJobCreditedToCurrentStage, nil
	case "candidate,job,credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateJobCreditedToCurrentStageRejectReason, nil
	case "candidate,job,credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateJobCreditedToRejectReason, nil
	case "candidate,job,current_stage":
		return ApplicationsRetrieveRequestExpandCandidateJobCurrentStage, nil
	case "candidate,job,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateJobCurrentStageRejectReason, nil
	case "candidate,job,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateJobRejectReason, nil
	case "candidate,reject_reason":
		return ApplicationsRetrieveRequestExpandCandidateRejectReason, nil
	case "credited_to":
		return ApplicationsRetrieveRequestExpandCreditedTo, nil
	case "credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandCreditedToCurrentStage, nil
	case "credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandCreditedToCurrentStageRejectReason, nil
	case "credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandCreditedToRejectReason, nil
	case "current_stage":
		return ApplicationsRetrieveRequestExpandCurrentStage, nil
	case "current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandCurrentStageRejectReason, nil
	case "job":
		return ApplicationsRetrieveRequestExpandJob, nil
	case "job,credited_to":
		return ApplicationsRetrieveRequestExpandJobCreditedTo, nil
	case "job,credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandJobCreditedToCurrentStage, nil
	case "job,credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandJobCreditedToCurrentStageRejectReason, nil
	case "job,credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandJobCreditedToRejectReason, nil
	case "job,current_stage":
		return ApplicationsRetrieveRequestExpandJobCurrentStage, nil
	case "job,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandJobCurrentStageRejectReason, nil
	case "job,reject_reason":
		return ApplicationsRetrieveRequestExpandJobRejectReason, nil
	case "offers":
		return ApplicationsRetrieveRequestExpandOffers, nil
	case "offers,candidate":
		return ApplicationsRetrieveRequestExpandOffersCandidate, nil
	case "offers,candidate,credited_to":
		return ApplicationsRetrieveRequestExpandOffersCandidateCreditedTo, nil
	case "offers,candidate,credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandOffersCandidateCreditedToCurrentStage, nil
	case "offers,candidate,credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateCreditedToCurrentStageRejectReason, nil
	case "offers,candidate,credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateCreditedToRejectReason, nil
	case "offers,candidate,current_stage":
		return ApplicationsRetrieveRequestExpandOffersCandidateCurrentStage, nil
	case "offers,candidate,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateCurrentStageRejectReason, nil
	case "offers,candidate,job":
		return ApplicationsRetrieveRequestExpandOffersCandidateJob, nil
	case "offers,candidate,job,credited_to":
		return ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedTo, nil
	case "offers,candidate,job,credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedToCurrentStage, nil
	case "offers,candidate,job,credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedToCurrentStageRejectReason, nil
	case "offers,candidate,job,credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateJobCreditedToRejectReason, nil
	case "offers,candidate,job,current_stage":
		return ApplicationsRetrieveRequestExpandOffersCandidateJobCurrentStage, nil
	case "offers,candidate,job,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateJobCurrentStageRejectReason, nil
	case "offers,candidate,job,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateJobRejectReason, nil
	case "offers,candidate,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCandidateRejectReason, nil
	case "offers,credited_to":
		return ApplicationsRetrieveRequestExpandOffersCreditedTo, nil
	case "offers,credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandOffersCreditedToCurrentStage, nil
	case "offers,credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCreditedToCurrentStageRejectReason, nil
	case "offers,credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCreditedToRejectReason, nil
	case "offers,current_stage":
		return ApplicationsRetrieveRequestExpandOffersCurrentStage, nil
	case "offers,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersCurrentStageRejectReason, nil
	case "offers,job":
		return ApplicationsRetrieveRequestExpandOffersJob, nil
	case "offers,job,credited_to":
		return ApplicationsRetrieveRequestExpandOffersJobCreditedTo, nil
	case "offers,job,credited_to,current_stage":
		return ApplicationsRetrieveRequestExpandOffersJobCreditedToCurrentStage, nil
	case "offers,job,credited_to,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersJobCreditedToCurrentStageRejectReason, nil
	case "offers,job,credited_to,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersJobCreditedToRejectReason, nil
	case "offers,job,current_stage":
		return ApplicationsRetrieveRequestExpandOffersJobCurrentStage, nil
	case "offers,job,current_stage,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersJobCurrentStageRejectReason, nil
	case "offers,job,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersJobRejectReason, nil
	case "offers,reject_reason":
		return ApplicationsRetrieveRequestExpandOffersRejectReason, nil
	case "reject_reason":
		return ApplicationsRetrieveRequestExpandRejectReason, nil
	}
	var t ApplicationsRetrieveRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (a ApplicationsRetrieveRequestExpand) Ptr() *ApplicationsRetrieveRequestExpand {
	return &a
}
