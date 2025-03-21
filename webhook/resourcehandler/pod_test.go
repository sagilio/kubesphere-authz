// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package resourcehandler

import (
	"encoding/json"
	"testing"

	v1 "k8s.io/api/admission/v1"
	core "k8s.io/api/core/v1"
)
func TestPodHandler(t *testing.T){
	//var review v1.AdmissionReview
	var podObject core.Pod
	var container core.Container
	container.Image="nginx:1.14.2"
	podObject.Spec.Containers=append(podObject.Spec.Containers, container)
	var review v1.AdmissionReview
	review.Request=&v1.AdmissionRequest{}
	review.Request.Namespace="default"
	review.Request.Resource.Resource="pods"
	data,err:=json.Marshal(podObject)
	if err!=nil{
		t.Error(err)
	}
	review.Request.Object.Raw=data
	res:=CheckTrustedImageOfPod(review,"../casbinconfig/image_model.conf","../casbinconfig/image_policy.csv")
	if res==nil{
		t.Error("should be rejected")
		return
	}
	if res.Error()!="casbin rejects the untrusted image nginx:1.14.2"{
		t.Error("should be rejected")

	}
}
func TestPodHandler2(t *testing.T){
	//var review v1.AdmissionReview
	var podObject core.Pod
	var container core.Container
	container.Image="nginx:1.14.3"
	podObject.Spec.Containers=append(podObject.Spec.Containers, container)
	var review v1.AdmissionReview
	review.Request=&v1.AdmissionRequest{}
	review.Request.Resource.Resource="pods"
	review.Request.Namespace="default"
	data,err:=json.Marshal(podObject)
	if err!=nil{
		t.Error(err)
	}
	review.Request.Object.Raw=data
	res:=CheckTrustedImageOfPod(review,"../casbinconfig/image_model.conf","../casbinconfig/image_policy.csv")
	if res!=nil{
		t.Error("should be rejected")
	}

}

func TestPodHandler3(t *testing.T){
	
	var review v1.AdmissionReview
	review.Request=&v1.AdmissionRequest{}
	review.Request.Namespace="default"
	review.Request.Resource.Resource="pods"
	
	review.Request.Operation="DELETE"
	res:=CheckTrustedImageOfPod(review,"../casbinconfig/image_model.conf","../casbinconfig/image_policy.csv")
	if res!=nil{
		t.Error(res.Error())
	}

}
