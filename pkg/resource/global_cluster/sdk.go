// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package global_cluster

import (
	"context"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/rds"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/rds-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.RDS{}
	_ = &svcapitypes.GlobalCluster{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadManyInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeGlobalClustersOutput
	resp, err = rm.sdkapi.DescribeGlobalClustersWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeGlobalClusters", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "GlobalClusterNotFoundFault" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.GlobalClusters {
		if elem.DatabaseName != nil {
			ko.Spec.DatabaseName = elem.DatabaseName
		} else {
			ko.Spec.DatabaseName = nil
		}
		if elem.DeletionProtection != nil {
			ko.Spec.DeletionProtection = elem.DeletionProtection
		} else {
			ko.Spec.DeletionProtection = nil
		}
		if elem.Engine != nil {
			ko.Spec.Engine = elem.Engine
		} else {
			ko.Spec.Engine = nil
		}
		if elem.EngineVersion != nil {
			ko.Spec.EngineVersion = elem.EngineVersion
		} else {
			ko.Spec.EngineVersion = nil
		}
		if elem.FailoverState != nil {
			f4 := &svcapitypes.FailoverState{}
			if elem.FailoverState.FromDbClusterArn != nil {
				f4.FromDBClusterARN = elem.FailoverState.FromDbClusterArn
			}
			if elem.FailoverState.Status != nil {
				f4.Status = elem.FailoverState.Status
			}
			if elem.FailoverState.ToDbClusterArn != nil {
				f4.ToDBClusterARN = elem.FailoverState.ToDbClusterArn
			}
			ko.Status.FailoverState = f4
		} else {
			ko.Status.FailoverState = nil
		}
		if elem.GlobalClusterArn != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.GlobalClusterArn)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.GlobalClusterIdentifier != nil {
			ko.Spec.GlobalClusterIdentifier = elem.GlobalClusterIdentifier
		} else {
			ko.Spec.GlobalClusterIdentifier = nil
		}
		if elem.GlobalClusterMembers != nil {
			f7 := []*svcapitypes.GlobalClusterMember{}
			for _, f7iter := range elem.GlobalClusterMembers {
				f7elem := &svcapitypes.GlobalClusterMember{}
				if f7iter.DBClusterArn != nil {
					f7elem.DBClusterARN = f7iter.DBClusterArn
				}
				if f7iter.GlobalWriteForwardingStatus != nil {
					f7elem.GlobalWriteForwardingStatus = f7iter.GlobalWriteForwardingStatus
				}
				if f7iter.IsWriter != nil {
					f7elem.IsWriter = f7iter.IsWriter
				}
				if f7iter.Readers != nil {
					f7elemf3 := []*string{}
					for _, f7elemf3iter := range f7iter.Readers {
						var f7elemf3elem string
						f7elemf3elem = *f7elemf3iter
						f7elemf3 = append(f7elemf3, &f7elemf3elem)
					}
					f7elem.Readers = f7elemf3
				}
				f7 = append(f7, f7elem)
			}
			ko.Status.GlobalClusterMembers = f7
		} else {
			ko.Status.GlobalClusterMembers = nil
		}
		if elem.GlobalClusterResourceId != nil {
			ko.Status.GlobalClusterResourceID = elem.GlobalClusterResourceId
		} else {
			ko.Status.GlobalClusterResourceID = nil
		}
		if elem.Status != nil {
			ko.Status.Status = elem.Status
		} else {
			ko.Status.Status = nil
		}
		if elem.StorageEncrypted != nil {
			ko.Spec.StorageEncrypted = elem.StorageEncrypted
		} else {
			ko.Spec.StorageEncrypted = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadManyInput returns true if there are any fields
// for the ReadMany Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadManyInput(
	r *resource,
) bool {
	return false
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeGlobalClustersInput, error) {
	res := &svcsdk.DescribeGlobalClustersInput{}

	if r.ko.Spec.GlobalClusterIdentifier != nil {
		res.SetGlobalClusterIdentifier(*r.ko.Spec.GlobalClusterIdentifier)
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer exit(err)
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreateGlobalClusterOutput
	_ = resp
	resp, err = rm.sdkapi.CreateGlobalClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateGlobalCluster", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.GlobalCluster.DatabaseName != nil {
		ko.Spec.DatabaseName = resp.GlobalCluster.DatabaseName
	} else {
		ko.Spec.DatabaseName = nil
	}
	if resp.GlobalCluster.DeletionProtection != nil {
		ko.Spec.DeletionProtection = resp.GlobalCluster.DeletionProtection
	} else {
		ko.Spec.DeletionProtection = nil
	}
	if resp.GlobalCluster.Engine != nil {
		ko.Spec.Engine = resp.GlobalCluster.Engine
	} else {
		ko.Spec.Engine = nil
	}
	if resp.GlobalCluster.EngineVersion != nil {
		ko.Spec.EngineVersion = resp.GlobalCluster.EngineVersion
	} else {
		ko.Spec.EngineVersion = nil
	}
	if resp.GlobalCluster.FailoverState != nil {
		f4 := &svcapitypes.FailoverState{}
		if resp.GlobalCluster.FailoverState.FromDbClusterArn != nil {
			f4.FromDBClusterARN = resp.GlobalCluster.FailoverState.FromDbClusterArn
		}
		if resp.GlobalCluster.FailoverState.Status != nil {
			f4.Status = resp.GlobalCluster.FailoverState.Status
		}
		if resp.GlobalCluster.FailoverState.ToDbClusterArn != nil {
			f4.ToDBClusterARN = resp.GlobalCluster.FailoverState.ToDbClusterArn
		}
		ko.Status.FailoverState = f4
	} else {
		ko.Status.FailoverState = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.GlobalCluster.GlobalClusterArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.GlobalCluster.GlobalClusterArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.GlobalCluster.GlobalClusterIdentifier != nil {
		ko.Spec.GlobalClusterIdentifier = resp.GlobalCluster.GlobalClusterIdentifier
	} else {
		ko.Spec.GlobalClusterIdentifier = nil
	}
	if resp.GlobalCluster.GlobalClusterMembers != nil {
		f7 := []*svcapitypes.GlobalClusterMember{}
		for _, f7iter := range resp.GlobalCluster.GlobalClusterMembers {
			f7elem := &svcapitypes.GlobalClusterMember{}
			if f7iter.DBClusterArn != nil {
				f7elem.DBClusterARN = f7iter.DBClusterArn
			}
			if f7iter.GlobalWriteForwardingStatus != nil {
				f7elem.GlobalWriteForwardingStatus = f7iter.GlobalWriteForwardingStatus
			}
			if f7iter.IsWriter != nil {
				f7elem.IsWriter = f7iter.IsWriter
			}
			if f7iter.Readers != nil {
				f7elemf3 := []*string{}
				for _, f7elemf3iter := range f7iter.Readers {
					var f7elemf3elem string
					f7elemf3elem = *f7elemf3iter
					f7elemf3 = append(f7elemf3, &f7elemf3elem)
				}
				f7elem.Readers = f7elemf3
			}
			f7 = append(f7, f7elem)
		}
		ko.Status.GlobalClusterMembers = f7
	} else {
		ko.Status.GlobalClusterMembers = nil
	}
	if resp.GlobalCluster.GlobalClusterResourceId != nil {
		ko.Status.GlobalClusterResourceID = resp.GlobalCluster.GlobalClusterResourceId
	} else {
		ko.Status.GlobalClusterResourceID = nil
	}
	if resp.GlobalCluster.Status != nil {
		ko.Status.Status = resp.GlobalCluster.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.GlobalCluster.StorageEncrypted != nil {
		ko.Spec.StorageEncrypted = resp.GlobalCluster.StorageEncrypted
	} else {
		ko.Spec.StorageEncrypted = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateGlobalClusterInput, error) {
	res := &svcsdk.CreateGlobalClusterInput{}

	if r.ko.Spec.DatabaseName != nil {
		res.SetDatabaseName(*r.ko.Spec.DatabaseName)
	}
	if r.ko.Spec.DeletionProtection != nil {
		res.SetDeletionProtection(*r.ko.Spec.DeletionProtection)
	}
	if r.ko.Spec.Engine != nil {
		res.SetEngine(*r.ko.Spec.Engine)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.GlobalClusterIdentifier != nil {
		res.SetGlobalClusterIdentifier(*r.ko.Spec.GlobalClusterIdentifier)
	}
	if r.ko.Spec.SourceDBClusterIdentifier != nil {
		res.SetSourceDBClusterIdentifier(*r.ko.Spec.SourceDBClusterIdentifier)
	}
	if r.ko.Spec.StorageEncrypted != nil {
		res.SetStorageEncrypted(*r.ko.Spec.StorageEncrypted)
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer exit(err)
	input, err := rm.newUpdateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.ModifyGlobalClusterOutput
	_ = resp
	resp, err = rm.sdkapi.ModifyGlobalClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "ModifyGlobalCluster", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if resp.GlobalCluster.DatabaseName != nil {
		ko.Spec.DatabaseName = resp.GlobalCluster.DatabaseName
	} else {
		ko.Spec.DatabaseName = nil
	}
	if resp.GlobalCluster.DeletionProtection != nil {
		ko.Spec.DeletionProtection = resp.GlobalCluster.DeletionProtection
	} else {
		ko.Spec.DeletionProtection = nil
	}
	if resp.GlobalCluster.Engine != nil {
		ko.Spec.Engine = resp.GlobalCluster.Engine
	} else {
		ko.Spec.Engine = nil
	}
	if resp.GlobalCluster.EngineVersion != nil {
		ko.Spec.EngineVersion = resp.GlobalCluster.EngineVersion
	} else {
		ko.Spec.EngineVersion = nil
	}
	if resp.GlobalCluster.FailoverState != nil {
		f4 := &svcapitypes.FailoverState{}
		if resp.GlobalCluster.FailoverState.FromDbClusterArn != nil {
			f4.FromDBClusterARN = resp.GlobalCluster.FailoverState.FromDbClusterArn
		}
		if resp.GlobalCluster.FailoverState.Status != nil {
			f4.Status = resp.GlobalCluster.FailoverState.Status
		}
		if resp.GlobalCluster.FailoverState.ToDbClusterArn != nil {
			f4.ToDBClusterARN = resp.GlobalCluster.FailoverState.ToDbClusterArn
		}
		ko.Status.FailoverState = f4
	} else {
		ko.Status.FailoverState = nil
	}
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.GlobalCluster.GlobalClusterArn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.GlobalCluster.GlobalClusterArn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.GlobalCluster.GlobalClusterIdentifier != nil {
		ko.Spec.GlobalClusterIdentifier = resp.GlobalCluster.GlobalClusterIdentifier
	} else {
		ko.Spec.GlobalClusterIdentifier = nil
	}
	if resp.GlobalCluster.GlobalClusterMembers != nil {
		f7 := []*svcapitypes.GlobalClusterMember{}
		for _, f7iter := range resp.GlobalCluster.GlobalClusterMembers {
			f7elem := &svcapitypes.GlobalClusterMember{}
			if f7iter.DBClusterArn != nil {
				f7elem.DBClusterARN = f7iter.DBClusterArn
			}
			if f7iter.GlobalWriteForwardingStatus != nil {
				f7elem.GlobalWriteForwardingStatus = f7iter.GlobalWriteForwardingStatus
			}
			if f7iter.IsWriter != nil {
				f7elem.IsWriter = f7iter.IsWriter
			}
			if f7iter.Readers != nil {
				f7elemf3 := []*string{}
				for _, f7elemf3iter := range f7iter.Readers {
					var f7elemf3elem string
					f7elemf3elem = *f7elemf3iter
					f7elemf3 = append(f7elemf3, &f7elemf3elem)
				}
				f7elem.Readers = f7elemf3
			}
			f7 = append(f7, f7elem)
		}
		ko.Status.GlobalClusterMembers = f7
	} else {
		ko.Status.GlobalClusterMembers = nil
	}
	if resp.GlobalCluster.GlobalClusterResourceId != nil {
		ko.Status.GlobalClusterResourceID = resp.GlobalCluster.GlobalClusterResourceId
	} else {
		ko.Status.GlobalClusterResourceID = nil
	}
	if resp.GlobalCluster.Status != nil {
		ko.Status.Status = resp.GlobalCluster.Status
	} else {
		ko.Status.Status = nil
	}
	if resp.GlobalCluster.StorageEncrypted != nil {
		ko.Spec.StorageEncrypted = resp.GlobalCluster.StorageEncrypted
	} else {
		ko.Spec.StorageEncrypted = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.ModifyGlobalClusterInput, error) {
	res := &svcsdk.ModifyGlobalClusterInput{}

	if r.ko.Spec.DeletionProtection != nil {
		res.SetDeletionProtection(*r.ko.Spec.DeletionProtection)
	}
	if r.ko.Spec.EngineVersion != nil {
		res.SetEngineVersion(*r.ko.Spec.EngineVersion)
	}
	if r.ko.Spec.GlobalClusterIdentifier != nil {
		res.SetGlobalClusterIdentifier(*r.ko.Spec.GlobalClusterIdentifier)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeleteGlobalClusterOutput
	_ = resp
	resp, err = rm.sdkapi.DeleteGlobalClusterWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteGlobalCluster", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteGlobalClusterInput, error) {
	res := &svcsdk.DeleteGlobalClusterInput{}

	if r.ko.Spec.GlobalClusterIdentifier != nil {
		res.SetGlobalClusterIdentifier(*r.ko.Spec.GlobalClusterIdentifier)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.GlobalCluster,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}

	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	if err == nil {
		return false
	}
	awsErr, ok := ackerr.AWSError(err)
	if !ok {
		return false
	}
	switch awsErr.Code() {
	case "GlobalClusterAlreadyExistsFault",
		"GlobalClusterQuotaExceededFault",
		"InvalidDBClusterStateFault":
		return true
	default:
		return false
	}
}
