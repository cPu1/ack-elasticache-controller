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

package cache_subnet_group

import (
	"context"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/elasticache"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/elasticache-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &aws.JSONValue{}
	_ = &svcsdk.ElastiCache{}
	_ = &svcapitypes.CacheSubnetGroup{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer exit(err)
	input, err := rm.newListRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DescribeCacheSubnetGroupsOutput
	resp, err = rm.sdkapi.DescribeCacheSubnetGroupsWithContext(ctx, input)
	rm.metrics.RecordAPICall("READ_MANY", "DescribeCacheSubnetGroups", err)
	if err != nil {
		if awsErr, ok := ackerr.AWSError(err); ok && awsErr.Code() == "CacheSubnetGroupNotFoundFault" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	found := false
	for _, elem := range resp.CacheSubnetGroups {
		if elem.ARN != nil {
			if ko.Status.ACKResourceMetadata == nil {
				ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
			}
			tmpARN := ackv1alpha1.AWSResourceName(*elem.ARN)
			ko.Status.ACKResourceMetadata.ARN = &tmpARN
		}
		if elem.CacheSubnetGroupDescription != nil {
			ko.Spec.CacheSubnetGroupDescription = elem.CacheSubnetGroupDescription
		} else {
			ko.Spec.CacheSubnetGroupDescription = nil
		}
		if elem.CacheSubnetGroupName != nil {
			ko.Spec.CacheSubnetGroupName = elem.CacheSubnetGroupName
		} else {
			ko.Spec.CacheSubnetGroupName = nil
		}
		if elem.Subnets != nil {
			f3 := []*svcapitypes.Subnet{}
			for _, f3iter := range elem.Subnets {
				f3elem := &svcapitypes.Subnet{}
				if f3iter.SubnetAvailabilityZone != nil {
					f3elemf0 := &svcapitypes.AvailabilityZone{}
					if f3iter.SubnetAvailabilityZone.Name != nil {
						f3elemf0.Name = f3iter.SubnetAvailabilityZone.Name
					}
					f3elem.SubnetAvailabilityZone = f3elemf0
				}
				if f3iter.SubnetIdentifier != nil {
					f3elem.SubnetIdentifier = f3iter.SubnetIdentifier
				}
				if f3iter.SubnetOutpost != nil {
					f3elemf2 := &svcapitypes.SubnetOutpost{}
					if f3iter.SubnetOutpost.SubnetOutpostArn != nil {
						f3elemf2.SubnetOutpostARN = f3iter.SubnetOutpost.SubnetOutpostArn
					}
					f3elem.SubnetOutpost = f3elemf2
				}
				f3 = append(f3, f3elem)
			}
			ko.Status.Subnets = f3
		} else {
			ko.Status.Subnets = nil
		}
		if elem.VpcId != nil {
			ko.Status.VPCID = elem.VpcId
		} else {
			ko.Status.VPCID = nil
		}
		found = true
		break
	}
	if !found {
		return nil, ackerr.NotFound
	}

	rm.setStatusDefaults(ko)
	// custom set output from response
	ko, err = rm.CustomDescribeCacheSubnetGroupsSetOutput(ctx, r, resp, ko)
	if err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// newListRequestPayload returns SDK-specific struct for the HTTP request
// payload of the List API call for the resource
func (rm *resourceManager) newListRequestPayload(
	r *resource,
) (*svcsdk.DescribeCacheSubnetGroupsInput, error) {
	res := &svcsdk.DescribeCacheSubnetGroupsInput{}

	if r.ko.Spec.CacheSubnetGroupName != nil {
		res.SetCacheSubnetGroupName(*r.ko.Spec.CacheSubnetGroupName)
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

	var resp *svcsdk.CreateCacheSubnetGroupOutput
	_ = resp
	resp, err = rm.sdkapi.CreateCacheSubnetGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreateCacheSubnetGroup", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.CacheSubnetGroup.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.CacheSubnetGroup.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.CacheSubnetGroup.Subnets != nil {
		f3 := []*svcapitypes.Subnet{}
		for _, f3iter := range resp.CacheSubnetGroup.Subnets {
			f3elem := &svcapitypes.Subnet{}
			if f3iter.SubnetAvailabilityZone != nil {
				f3elemf0 := &svcapitypes.AvailabilityZone{}
				if f3iter.SubnetAvailabilityZone.Name != nil {
					f3elemf0.Name = f3iter.SubnetAvailabilityZone.Name
				}
				f3elem.SubnetAvailabilityZone = f3elemf0
			}
			if f3iter.SubnetIdentifier != nil {
				f3elem.SubnetIdentifier = f3iter.SubnetIdentifier
			}
			if f3iter.SubnetOutpost != nil {
				f3elemf2 := &svcapitypes.SubnetOutpost{}
				if f3iter.SubnetOutpost.SubnetOutpostArn != nil {
					f3elemf2.SubnetOutpostARN = f3iter.SubnetOutpost.SubnetOutpostArn
				}
				f3elem.SubnetOutpost = f3elemf2
			}
			f3 = append(f3, f3elem)
		}
		ko.Status.Subnets = f3
	} else {
		ko.Status.Subnets = nil
	}
	if resp.CacheSubnetGroup.VpcId != nil {
		ko.Status.VPCID = resp.CacheSubnetGroup.VpcId
	} else {
		ko.Status.VPCID = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreateCacheSubnetGroupInput, error) {
	res := &svcsdk.CreateCacheSubnetGroupInput{}

	if r.ko.Spec.CacheSubnetGroupDescription != nil {
		res.SetCacheSubnetGroupDescription(*r.ko.Spec.CacheSubnetGroupDescription)
	}
	if r.ko.Spec.CacheSubnetGroupName != nil {
		res.SetCacheSubnetGroupName(*r.ko.Spec.CacheSubnetGroupName)
	}
	if r.ko.Spec.SubnetIDs != nil {
		f2 := []*string{}
		for _, f2iter := range r.ko.Spec.SubnetIDs {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		res.SetSubnetIds(f2)
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

	var resp *svcsdk.ModifyCacheSubnetGroupOutput
	_ = resp
	resp, err = rm.sdkapi.ModifyCacheSubnetGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "ModifyCacheSubnetGroup", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.CacheSubnetGroup.ARN != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.CacheSubnetGroup.ARN)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.CacheSubnetGroup.Subnets != nil {
		f3 := []*svcapitypes.Subnet{}
		for _, f3iter := range resp.CacheSubnetGroup.Subnets {
			f3elem := &svcapitypes.Subnet{}
			if f3iter.SubnetAvailabilityZone != nil {
				f3elemf0 := &svcapitypes.AvailabilityZone{}
				if f3iter.SubnetAvailabilityZone.Name != nil {
					f3elemf0.Name = f3iter.SubnetAvailabilityZone.Name
				}
				f3elem.SubnetAvailabilityZone = f3elemf0
			}
			if f3iter.SubnetIdentifier != nil {
				f3elem.SubnetIdentifier = f3iter.SubnetIdentifier
			}
			if f3iter.SubnetOutpost != nil {
				f3elemf2 := &svcapitypes.SubnetOutpost{}
				if f3iter.SubnetOutpost.SubnetOutpostArn != nil {
					f3elemf2.SubnetOutpostARN = f3iter.SubnetOutpost.SubnetOutpostArn
				}
				f3elem.SubnetOutpost = f3elemf2
			}
			f3 = append(f3, f3elem)
		}
		ko.Status.Subnets = f3
	} else {
		ko.Status.Subnets = nil
	}
	if resp.CacheSubnetGroup.VpcId != nil {
		ko.Status.VPCID = resp.CacheSubnetGroup.VpcId
	} else {
		ko.Status.VPCID = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.ModifyCacheSubnetGroupInput, error) {
	res := &svcsdk.ModifyCacheSubnetGroupInput{}

	if r.ko.Spec.CacheSubnetGroupDescription != nil {
		res.SetCacheSubnetGroupDescription(*r.ko.Spec.CacheSubnetGroupDescription)
	}
	if r.ko.Spec.CacheSubnetGroupName != nil {
		res.SetCacheSubnetGroupName(*r.ko.Spec.CacheSubnetGroupName)
	}
	if r.ko.Spec.SubnetIDs != nil {
		f2 := []*string{}
		for _, f2iter := range r.ko.Spec.SubnetIDs {
			var f2elem string
			f2elem = *f2iter
			f2 = append(f2, &f2elem)
		}
		res.SetSubnetIds(f2)
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer exit(err)
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return err
	}
	_, err = rm.sdkapi.DeleteCacheSubnetGroupWithContext(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeleteCacheSubnetGroup", err)
	return err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeleteCacheSubnetGroupInput, error) {
	res := &svcsdk.DeleteCacheSubnetGroupInput{}

	if r.ko.Spec.CacheSubnetGroupName != nil {
		res.SetCacheSubnetGroupName(*r.ko.Spec.CacheSubnetGroupName)
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.CacheSubnetGroup,
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
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
	}

	if rm.terminalAWSError(err) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		terminalCondition.Status = corev1.ConditionTrue
		awsErr, _ := ackerr.AWSError(err)
		errorMessage := awsErr.Message()
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
				errorMessage = awsErr.Message()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	if terminalCondition != nil || recoverableCondition != nil {
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
	case "CacheSubnetGroupQuotaExceeded",
		"CacheSubnetQuotaExceededFault",
		"SubnetInUse",
		"InvalidSubnet",
		"InvalidParameter",
		"InvalidParameterValue",
		"InvalidParameterCombination":
		return true
	default:
		return false
	}
}
