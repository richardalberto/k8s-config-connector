// Copyright 2022 Google LLC
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

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/core/v1alpha1"
	iamapi "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/iam/v1beta1"
	kcciamclient "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/iam/iamclient"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/reconciliationinterval"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/crdloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/fielddesc"
	crdtemplate "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/crd/template"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension"
	dclcontainer "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/extension/container"
	dclmetadata "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/metadata"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/dcl/schema/dclschemaloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/gvks/supportedgvks"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/k8s"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/krmtotf"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/servicemapping/servicemappingloader"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/text"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/fileutil"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/util/repo"

	"github.com/Masterminds/sprig"
	"github.com/golang-collections/go-datastructures/set"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// convenience struct for converting to the desired human readable output on the docs page from the fielddesc.FieldDescription struct
type HumanReadableFieldDescription struct {
	FullName         string
	RequirementLevel string
	Type             string
	Description      string
}

type IAM struct {
	IsDCLBased               bool
	SupportsConditions       bool
	SupportsAuditConfigs     bool
	ExternalReferenceFormats []string
}

type iamPolicyReference struct {
	Kind                     string
	IsDCLBased               bool
	SupportsConditions       bool
	SupportsAuditConfigs     bool
	ExternalReferenceFormats []string
}

type iamPolicyPartialReference struct {
	Kind                     string
	SupportsConditions       bool
	ExternalReferenceFormats []string
}

type iamPolicyMemberReference struct {
	Kind                     string
	SupportsConditions       bool
	ExternalReferenceFormats []string
}

type iamAuditConfigReference struct {
	Kind                     string
	ExternalReferenceFormats []string
}

type iamPolicyResource struct {
	resource
	SupportedReferences []iamPolicyReference
}

type iamPartialPolicyResource struct {
	resource
	SupportedReferences []iamPolicyPartialReference
}

type iamPolicyMemberResource struct {
	resource
	SupportedReferences []iamPolicyMemberReference
}

type iamAuditConfigResource struct {
	resource
	SupportedReferences []iamAuditConfigReference
}

type resource struct {
	FullyQualifiedName                             string
	Kind                                           string
	Spec                                           string
	Status                                         string
	ShortNames                                     string
	Annotations                                    []string
	IAM                                            *IAM
	SampleYamls                                    map[string]string
	SpecDescriptionContainsRequiredIfParentPresent bool
	SpecDescriptions                               []HumanReadableFieldDescription
	StatusDescriptions                             []HumanReadableFieldDescription
	IsAlphaResource                                bool
	DefaultReconcileInterval                       uint32
}

// some approved spellings in Google public doc
var (
	allowedSpellings = []string{
		"BigQuery",
		"GitHub",
		"gRPC",
	}
)

var (
	dclSchemaLoader       dclschemaloader.DCLSchemaLoader
	serviceMetadataLoader dclmetadata.ServiceMetadataLoader
)

func main() {
	if err := clearGeneratedDocsDir(); err != nil {
		log.Fatalf("error clearing generated docs dir: %v", err)
	}

	smLoader, err := servicemappingloader.New()
	if err != nil {
		log.Fatalf("error creating service mapping loader: %v", err)
	}
	dclSchemaLoader, err = dclschemaloader.New()
	if err != nil {
		log.Fatalf("error creating a DCL schema loader: %v", err)
	}
	serviceMetadataLoader = dclmetadata.New()
	for _, gvk := range supportedgvks.ManualResources(smLoader, serviceMetadataLoader) {
		if err := generateDocForGVK(gvk, smLoader); err != nil {
			log.Fatalf("error generating doc for GVK %v: %v", gvk, err)
		}
	}
}

func generateDocForGVK(gvk schema.GroupVersionKind, smLoader *servicemappingloader.ServiceMappingLoader) error {
	template, err := templateForGVK(gvk)
	if err != nil {
		return fmt.Errorf("error creating template: %v", err)
	}
	templateData, err := templateDataForGVK(gvk, smLoader)
	if err != nil {
		return fmt.Errorf("error preparing template data: %v", err)
	}
	outputFile, err := fileutil.NewEmptyFile(generatedDocPathForGVK(gvk))
	if err != nil {
		return fmt.Errorf("error creating empty file for output: %v", err)
	}
	if err := template.Execute(outputFile, templateData); err != nil {
		return fmt.Errorf("error while executing template: %v", err)
	}
	return nil
}

func templateForGVK(gvk schema.GroupVersionKind) (*template.Template, error) {
	templatesPath := repo.GetG3ResourceReferenceTemplatesPath()
	templateFileName := templateFileNameForGVK(gvk)
	templateFiles := []string{
		filepath.Join(templatesPath, templateFileName),
		filepath.Join(templatesPath, "shared/headercomment.tmpl"),
		filepath.Join(templatesPath, "shared/alphadisclaimer.tmpl"),
		filepath.Join(templatesPath, "shared/bigquerydatasetiamnote.tmpl"),
		filepath.Join(templatesPath, "shared/iamsupport.tmpl"),
		filepath.Join(templatesPath, "shared/resource.tmpl"),
	}
	template, err := template.New(templateFileName).Funcs(sprig.TxtFuncMap()).ParseFiles(templateFiles...)
	if err != nil {
		return nil, fmt.Errorf("error parsing template files: %v", err)
	}
	return template, nil
}

func templateDataForGVK(gvk schema.GroupVersionKind, smLoader *servicemappingloader.ServiceMappingLoader) (interface{}, error) {
	resource, err := constructResourceForGVK(gvk, smLoader)
	if err != nil {
		return nil, fmt.Errorf("error constructing resource data: %v", err)
	}

	switch gvk.Kind {
	case "IAMPolicy":
		supportedReferences, err := referencesSupportedByIAMPolicy(smLoader)
		if err != nil {
			return nil, fmt.Errorf("error determining references supported by IAMPolicy: %v", err)
		}
		return &iamPolicyResource{*resource, supportedReferences}, nil
	case "IAMPartialPolicy":
		// IAMPartialPolicy has the same resource supports as IAMPolicy
		supportedReferences, err := referencesSupportedByIAMPolicy(smLoader)
		if err != nil {
			return nil, fmt.Errorf("error determining references supported by IAMPolicy: %v", err)
		}
		references := make([]iamPolicyPartialReference, 0)
		for _, sr := range supportedReferences {
			r := iamPolicyPartialReference{
				Kind:                     sr.Kind,
				SupportsConditions:       sr.SupportsConditions,
				ExternalReferenceFormats: sr.ExternalReferenceFormats,
			}
			references = append(references, r)
		}
		return &iamPartialPolicyResource{*resource, references}, nil
	case "IAMPolicyMember":
		supportedReferences, err := referencesSupportedByIAMPolicyMember(smLoader)
		if err != nil {
			return nil, fmt.Errorf("error determining references supported by IAMPolicy: %v", err)
		}
		return &iamPolicyMemberResource{*resource, supportedReferences}, nil
	case "IAMAuditConfig":
		supportedReferences, err := referencesSupportedByIAMAuditConfig(smLoader)
		if err != nil {
			return nil, fmt.Errorf("error determining references supported by IAMPolicy: %v", err)
		}
		return &iamAuditConfigResource{*resource, supportedReferences}, nil
	default:
		return resource, nil
	}
}

func constructResourceForGVK(gvk schema.GroupVersionKind, smLoader *servicemappingloader.ServiceMappingLoader) (*resource, error) {
	r := &resource{}

	// crd properties
	crd, err := crdloader.FileToCRD(crdFilePathForGVK(gvk))
	if err != nil {
		return nil, fmt.Errorf("error loading CRD: %v", err)
	}
	r.FullyQualifiedName = crd.Name
	r.Kind = crd.Spec.Names.Kind
	crd.Spec.Names.ShortNames = append(crd.Spec.Names.ShortNames, strings.ToLower(r.Kind))
	r.ShortNames = strings.Join(crd.Spec.Names.ShortNames, "<br>")
	specYaml, err := crdtemplate.SpecToYAML(crd)
	if err != nil {
		return nil, fmt.Errorf("error converting spec to YAML: %v", err)
	}
	if len(specYaml) == 0 {
		r.Spec = r.Kind + " has an empty Spec\n"
	} else {
		r.Spec = string(specYaml)
	}
	statusYaml, err := crdtemplate.StatusToYAML(crd)
	if err != nil {
		return nil, fmt.Errorf("error converting status to YAML: %v", err)
	}
	r.Status = string(statusYaml)
	buildFieldDescriptions(r, crd)
	r.DefaultReconcileInterval = uint32(reconciliationinterval.MeanReconcileReenqueuePeriod(gvk, smLoader, serviceMetadataLoader).Seconds())
	if dclmetadata.IsDCLBasedResourceKind(gvk, serviceMetadataLoader) {
		resourceMetadata, found := serviceMetadataLoader.GetResourceWithGVK(gvk)
		if !found {
			return nil, fmt.Errorf("error finding the DCL based resource %v", gvk)
		}
		r.IsAlphaResource = resourceMetadata.DCLVersion == "alpha"
		if err := handleAnnotationsAndIAMSettingsForDCLBasedResource(r, gvk); err != nil {
			return nil, fmt.Errorf("error processing the DCL based resource %v: %v", gvk, err)
		}
	} else {
		if err := handleAnnotationsAndIAMSettingsForTFBasedResource(r, gvk, smLoader); err != nil {
			return nil, fmt.Errorf("error processing the TF based resource %v: %v", gvk, err)
		}
	}

	r.SampleYamls, err = buildSamples(r.Kind, sampleDirPathForGVK(gvk), smLoader)
	if err != nil {
		return nil, fmt.Errorf("error building samples: %v", err)
	}
	return r, nil
}

func handleAnnotationsAndIAMSettingsForDCLBasedResource(r *resource, gvk schema.GroupVersionKind) error {
	annotationSet := set.New()
	if k8s.ResourceSupportsStateAbsentInSpec(gvk.Kind) {
		annotationSet.Add(k8s.StateIntoSpecAnnotation)
	}
	resourceMetadata, found := serviceMetadataLoader.GetResourceWithGVK(gvk)
	if !found {
		return fmt.Errorf("ServiceMetadata for resource with GroupVersionKind %v not found", gvk)
	}
	containers, err := dclcontainer.GetContainersForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return err
	}
	// TODO(b/186159460): Delete this if-block once all resources support
	// hierarchical references.
	if !resourceMetadata.SupportsHierarchicalReferences {
		for _, c := range containers {
			annotationSet.Add(k8s.GetAnnotationForContainerType(c.Type))
		}
	}
	setAnnotations(r, annotationSet)
	externalReferenceFormat, err := getDCLExternalReferenceFormatIfSupportsIAM(gvk)
	if err != nil {
		return err
	}
	if externalReferenceFormat == "" { // Resource does not support IAM.
		return nil
	}
	r.IAM = &IAM{
		IsDCLBased: true,
		// DCL-based resources support conditions on IAMPolicy but do not support it
		// on IAMPolicyMember. The IAM template is updated to clarify that Conditions
		// are supported for IAMPolicy and IAMPartialPolicy but not IAMPolicyMember.
		SupportsConditions:       true,
		SupportsAuditConfigs:     false, // No DCL-based resources support AuditConfigs.
		ExternalReferenceFormats: []string{externalReferenceFormat},
	}
	return nil
}

func handleAnnotationsAndIAMSettingsForTFBasedResource(r *resource, gvk schema.GroupVersionKind, smLoader *servicemappingloader.ServiceMappingLoader) error {
	annotationSet := set.New()
	if k8s.ResourceSupportsStateAbsentInSpec(gvk.Kind) {
		annotationSet.Add(k8s.StateIntoSpecAnnotation)
	}
	rcs, err := smLoader.GetResourceConfigs(gvk)
	if err != nil {
		return fmt.Errorf("error getting resource configs: %v", err)
	}
	for _, rc := range rcs {
		if rc.Directives != nil {
			for _, d := range rc.Directives {
				annotationSet.Add(k8s.FormatAnnotation(text.SnakeCaseToKebabCase(d)))
			}
		}
		if !krmtotf.SupportsHierarchicalReferences(rc) {
			// TODO(b/193177782): Delete this if-block once all resources
			// support hierarchical references.
			for _, c := range rc.Containers {
				annotationSet.Add(k8s.GetAnnotationForContainerType(c.Type))
			}
		}
	}
	setAnnotations(r, annotationSet)
	if !iamapi.IsHandwrittenIAM(gvk) && resourceSupportsIAMPolicyAndPolicyMember(rcs[0]) {
		externalReferenceFormats, err := iamExternalReferenceFormatsFor(gvk, rcs)
		if err != nil {
			return fmt.Errorf("error determining IAM external reference formats for GVK %v: %v", gvk, err)
		}
		r.IAM = &IAM{
			ExternalReferenceFormats: externalReferenceFormats,

			// All ResourceConfigs of a given kind are assumed to have the same value for SupportsConditions
			SupportsConditions: rcs[0].IAMConfig.SupportsConditions,

			SupportsAuditConfigs: resourceSupportsIAMAuditConfig(rcs[0]),
		}
	}
	return nil
}

func setAnnotations(r *resource, annotationSet *set.Set) {
	if annotationSet.Len() > 0 {
		// arrange annotations alphabetically
		i := annotationSet.Flatten()
		strArr := make([]string, len(i))
		for k, v := range i {
			strArr[k] = v.(string)
		}
		sort.Strings(strArr)
		r.Annotations = strArr
	}
}

func iamExternalReferenceFormatsFor(gvk schema.GroupVersionKind, rcs []*v1alpha1.ResourceConfig) ([]string, error) {
	switch gvk.Kind {
	default:
		formatSet := make(map[string]struct{})
		for _, rc := range rcs {
			if rc.IDTemplate != "" {
				formatSet[rc.IDTemplate] = struct{}{}
			} else if krmtotf.SupportsServerGeneratedIDField(rc) {
				formatSet[krmtotf.ServerGeneratedIdToTemplate(rc)] = struct{}{}
			}
		}
		formatList := make([]string, 0)
		for f, _ := range formatSet {
			formatList = append(formatList, f)
		}
		if len(formatList) == 0 {
			return nil, fmt.Errorf("no IAM external reference formats found for GVK %v", gvk)
		}
		sort.Strings(formatList)
		return formatList, nil
	case "Folder":
		// TODO(kcc-eng): Add an IDTemplate for Folder of the form "folders/{{folder_id}}".
		// This would require adding a "folder_id" field to the google_folder TF resource.
		return []string{"folders/{{folder_id}}"}, nil
	case "KMSCryptoKey":
		// The IDTemplate for KMSCryptoKey is "{{key_ring}}/cryptoKeys/{{name}}" where {{key_ring}} refers to the relative resource name of the key ring.
		// The ID value will be expanded to "projects/{{project}}/locations/{{location}}/keyRings/{{key_ring_id}}/cryptoKeys/{{name}}".
		// To reduce confusion, we want to state an explicit format in resource reference docs.
		return []string{"projects/{{project}}/locations/{{location}}/keyRings/{{key_ring_id}}/cryptoKeys/{{name}}"}, nil
	case "IAMServiceAccount":
		if len(rcs) != 1 || rcs[0].IDTemplate != "projects/{{project}}/serviceAccounts/[{{account_id}}@{{project}}.iam.gserviceaccount.com|{{unique_id}}]" {
			return nil, fmt.Errorf("Assumptions about the format of the IAMSeviceAccount IDTemplate no longer hold true. Modify the above if statement and carefully review this code block for other necessary changes.")
		}
		return []string{"projects/{{project}}/serviceAccounts/{{account_id}}@{{project}}.iam.gserviceaccount.com"}, nil
	}
}

func getDCLExternalReferenceFormatIfSupportsIAM(gvk schema.GroupVersionKind) (string, error) {
	dclSchema, err := dclschemaloader.GetDCLSchemaForGVK(gvk, serviceMetadataLoader, dclSchemaLoader)
	if err != nil {
		return "", err
	}
	supportsIAM, err := extension.HasIam(dclSchema)
	if err != nil {
		return "", err
	}
	if !supportsIAM {
		return "", nil
	}
	// DCL-based resources all have only one x-dcl-id
	externalReferenceFormat, err := extension.GetNameValueTemplate(dclSchema)
	if err != nil {
		return "", fmt.Errorf("error determining IAM external reference formats for GVK %v: %v", gvk, err)
	}
	return externalReferenceFormat, nil
}

func buildSamples(kind, sampleDirPath string, smLoader *servicemappingloader.ServiceMappingLoader) (map[string]string, error) {
	fileInfos, err := ioutil.ReadDir(sampleDirPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory %v: %v", sampleDirPath, err)
	}
	sampleYAMLs := make(map[string]string, 0)
	if containsMultipleSamples(fileInfos) {
		for _, subDir := range fileInfos {
			subDirPath := filepath.Join(sampleDirPath, subDir.Name())
			sample, err := buildSampleYAML(kind, subDirPath)
			if err != nil {
				return nil, fmt.Errorf("error building sample at %v: %v", subDirPath, err)
			}
			name := formatDirectoryName(subDir.Name(), smLoader.GetServiceMappings())
			sampleYAMLs[name] = sample
		}
	} else {
		sample, err := buildSampleYAML(kind, sampleDirPath)
		if err != nil {
			return nil, fmt.Errorf("error building sample at %v: %v", sampleDirPath, err)
		}
		sampleYAMLs["Typical Use Case"] = sample
	}
	return sampleYAMLs, nil
}

func containsMultipleSamples(sampleDirFiles []os.FileInfo) bool {
	// If any of the files is a directory, this is a nested sample directory
	return sampleDirFiles[0].IsDir()
}

func buildSampleYAML(kind, sampleDir string) (string, error) {
	fileInfos, err := ioutil.ReadDir(sampleDir)
	if err != nil {
		return "", fmt.Errorf("error reading directory %v: %v", sampleDir, err)
	}
	objectYAMLs := make([]string, 0)
	for _, f := range fileInfos {
		bytes, err := ioutil.ReadFile(filepath.Join(sampleDir, f.Name()))
		if err != nil {
			return "", fmt.Errorf("error reading file '%v': %v\n", filepath.Join(sampleDir, f.Name()), err)
		}
		objectYAML := strings.Trim(string(bytes), "\n")
		// We want the object for this kind to be at the top. The file will
		// contain the lowercase name of the kind, so use this to find the
		// right one and put it at the front. Strip the Apache2 header from
		// all other objects.
		if strings.Contains(f.Name(), strings.ToLower(kind)+".yaml") {
			objectYAMLs = append([]string{objectYAML}, objectYAMLs...)
		} else {
			objectYAMLs = append(objectYAMLs, stripHeader(objectYAML))
		}
	}
	return strings.Join(objectYAMLs, "\n---\n"), nil
}

func stripHeader(sample string) string {
	sampleSplit := strings.Split(sample, "\n")
	// Find the first line that is not a comment (i.e., does not start with '#')
	startIdx := 0
	for i, line := range sampleSplit {
		if !strings.HasPrefix(line, "#") {
			startIdx = i
			break
		}
	}
	res := strings.Join(sampleSplit[startIdx:], "\n")
	// Remove the newline between the header and the resource definition.
	return strings.Trim(res, "\n")
}

func buildFieldDescriptions(r *resource, crd *apiextensions.CustomResourceDefinition) error {
	specDesc := fielddesc.GetSpecDescription(crd)
	specDescriptions := dropRootAndFlattenChildrenDescriptions(specDesc)
	r.SpecDescriptions = fieldDescriptionsToHumanReadable(specDescriptions)
	r.SpecDescriptionContainsRequiredIfParentPresent = atLeastOneFieldHasRequiredWhenParentPresentRequirementLevel(specDesc)
	statusDesc, err := fielddesc.GetStatusDescription(crd)
	if err != nil {
		return fmt.Errorf("error getting status descriptions: %v", err)
	}
	statusDescriptions := dropRootAndFlattenChildrenDescriptions(statusDesc)
	r.StatusDescriptions = fieldDescriptionsToHumanReadable(statusDescriptions)
	return nil
}

func fieldDescriptionsToHumanReadable(descriptions []fielddesc.FieldDescription) []HumanReadableFieldDescription {
	result := make([]HumanReadableFieldDescription, 0)
	for _, d := range descriptions {
		result = append(result, fieldDescriptionToHumanReadable(d))
	}
	return result
}

func fieldDescriptionToHumanReadable(desc fielddesc.FieldDescription) HumanReadableFieldDescription {
	return HumanReadableFieldDescription{
		FullName:         formatName(desc),
		RequirementLevel: formatRequirementLevel(desc),
		Type:             desc.Type,
		Description:      desc.Description,
	}
}

func formatName(desc fielddesc.FieldDescription) string {
	name := strings.Join(desc.FullName, ".")

	// Remove extraneous period for list field names.
	// (i.e. listField.[].subField -> listField[].subField)
	name = strings.ReplaceAll(name, ".[]", "[]")

	name = strings.TrimPrefix(name, "spec.")
	name = strings.TrimPrefix(name, "status.")

	return name
}

func formatRequirementLevel(desc fielddesc.FieldDescription) string {
	switch desc.RequirementLevel {
	case fielddesc.RequiredRequirementLevel:
		return "Required"
	case fielddesc.RequiredWhenParentPresentRequirementLevel:
		return "Required*"
	case fielddesc.OptionalRequirementLevel:
		return "Optional"
	default:
		panic(fmt.Errorf("unhandled requirement level: %v", desc.RequirementLevel))
	}
}

func dropRootAndFlattenChildrenDescriptions(rootDesc fielddesc.FieldDescription) []fielddesc.FieldDescription {
	result := flattenChildrenDescription(nil, rootDesc)
	return result[1:]
}

func flattenChildrenDescription(result []fielddesc.FieldDescription, fd fielddesc.FieldDescription) []fielddesc.FieldDescription {
	if result == nil {
		result = make([]fielddesc.FieldDescription, 0, 1)
	}
	result = append(result, fd)
	for _, child := range fd.Children {
		result = flattenChildrenDescription(result, child)
	}
	return result
}

func atLeastOneFieldHasRequiredWhenParentPresentRequirementLevel(desc fielddesc.FieldDescription) bool {
	if desc.RequirementLevel == fielddesc.RequiredWhenParentPresentRequirementLevel {
		return true
	}
	for _, c := range desc.Children {
		if atLeastOneFieldHasRequiredWhenParentPresentRequirementLevel(c) {
			return true
		}
	}
	return false
}

func referencesSupportedByIAMPolicy(smLoader *servicemappingloader.ServiceMappingLoader) ([]iamPolicyReference, error) {
	refs := make([]iamPolicyReference, 0)
	for _, gvk := range supportedgvks.BasedOnManualServiceMappings(smLoader) {
		rcs, err := smLoader.GetResourceConfigs(gvk)
		if err != nil {
			return nil, fmt.Errorf("error getting resource configs for GVK %v: %v", gvk, err)
		}
		if !resourceSupportsIAMPolicyAndPolicyMember(rcs[0]) {
			continue
		}
		externalReferenceFormats, err := iamExternalReferenceFormatsFor(gvk, rcs)
		if err != nil {
			return nil, fmt.Errorf("error determining IAM external reference formats for GVK %v: %v", gvk, err)
		}
		refs = append(refs, iamPolicyReference{
			Kind:                     gvk.Kind,
			SupportsConditions:       rcs[0].IAMConfig.SupportsConditions,
			SupportsAuditConfigs:     resourceSupportsIAMAuditConfig(rcs[0]),
			ExternalReferenceFormats: externalReferenceFormats,
		})
	}
	for _, gvk := range supportedgvks.BasedOnDCL(serviceMetadataLoader) {
		externalReferenceFormat, err := getDCLExternalReferenceFormatIfSupportsIAM(gvk)
		if err != nil {
			return nil, err
		}
		if externalReferenceFormat == "" { // Resource does not support IAM.
			continue
		}
		refs = append(refs, iamPolicyReference{
			Kind:       gvk.Kind,
			IsDCLBased: true,
			// DCL-based resources support conditions on IAMPolicy but do not support it
			// on IAMPolicyMember. The IAM template is updated to clarify that Conditions
			// are supported for IAMPolicy and IAMPartialPolicy but not IAMPolicyMember.
			SupportsConditions:       true,
			SupportsAuditConfigs:     false, // No DCL-based resources support AuditConfigs.
			ExternalReferenceFormats: []string{externalReferenceFormat},
		})
	}
	for gvk, extOnlyType := range kcciamclient.ExternalOnlyTypes {
		refs = append(refs, iamPolicyReference{
			Kind:                     gvk.Kind,
			SupportsConditions:       extOnlyType.ResourceConfig.IAMConfig.SupportsConditions,
			SupportsAuditConfigs:     resourceSupportsIAMAuditConfig(extOnlyType.ResourceConfig),
			ExternalReferenceFormats: []string{extOnlyType.ExternalFormat},
		})
	}
	sort.Slice(refs, func(i, j int) bool {
		return refs[i].Kind < refs[j].Kind
	})
	return refs, nil
}

func referencesSupportedByIAMPolicyMember(smLoader *servicemappingloader.ServiceMappingLoader) ([]iamPolicyMemberReference, error) {
	refs := make([]iamPolicyMemberReference, 0)
	for _, gvk := range supportedgvks.BasedOnManualServiceMappings(smLoader) {
		rcs, err := smLoader.GetResourceConfigs(gvk)
		if err != nil {
			return nil, fmt.Errorf("error getting resource configs for GVK %v: %v", gvk, err)
		}
		if !resourceSupportsIAMPolicyAndPolicyMember(rcs[0]) {
			continue
		}
		externalReferenceFormats, err := iamExternalReferenceFormatsFor(gvk, rcs)
		if err != nil {
			return nil, fmt.Errorf("error determining IAM external reference formats for GVK %v: %v", gvk, err)
		}
		refs = append(refs, iamPolicyMemberReference{
			Kind:                     gvk.Kind,
			SupportsConditions:       rcs[0].IAMConfig.SupportsConditions,
			ExternalReferenceFormats: externalReferenceFormats,
		})
	}
	for _, gvk := range supportedgvks.BasedOnDCL(serviceMetadataLoader) {
		externalReferenceFormat, err := getDCLExternalReferenceFormatIfSupportsIAM(gvk)
		if err != nil {
			return nil, err
		}
		if externalReferenceFormat == "" { // Resource does not support IAM.
			continue
		}
		refs = append(refs, iamPolicyMemberReference{
			Kind: gvk.Kind,
			// DCL-based resources do not support conditions on IAMPolicyMember.
			SupportsConditions:       false,
			ExternalReferenceFormats: []string{externalReferenceFormat},
		})
	}
	for gvk, extOnlyType := range kcciamclient.ExternalOnlyTypes {
		refs = append(refs, iamPolicyMemberReference{
			Kind:                     gvk.Kind,
			SupportsConditions:       extOnlyType.ResourceConfig.IAMConfig.SupportsConditions,
			ExternalReferenceFormats: []string{extOnlyType.ExternalFormat},
		})
	}
	sort.Slice(refs, func(i, j int) bool {
		return refs[i].Kind < refs[j].Kind
	})
	return refs, nil
}

func referencesSupportedByIAMAuditConfig(smLoader *servicemappingloader.ServiceMappingLoader) ([]iamAuditConfigReference, error) {
	refs := make([]iamAuditConfigReference, 0)
	for _, gvk := range supportedgvks.BasedOnManualServiceMappings(smLoader) {
		rcs, err := smLoader.GetResourceConfigs(gvk)
		if err != nil {
			return nil, fmt.Errorf("error getting resource configs for GVK %v: %v", gvk, err)
		}
		if !resourceSupportsIAMAuditConfig(rcs[0]) {
			continue
		}
		externalReferenceFormats, err := iamExternalReferenceFormatsFor(gvk, rcs)
		if err != nil {
			return nil, fmt.Errorf("error determining IAM external reference formats for GVK %v: %v", gvk, err)
		}
		refs = append(refs, iamAuditConfigReference{
			Kind:                     gvk.Kind,
			ExternalReferenceFormats: externalReferenceFormats,
		})
	}
	for gvk, extOnlyType := range kcciamclient.ExternalOnlyTypes {
		if !resourceSupportsIAMAuditConfig(extOnlyType.ResourceConfig) {
			continue
		}
		refs = append(refs, iamAuditConfigReference{
			Kind:                     gvk.Kind,
			ExternalReferenceFormats: []string{extOnlyType.ExternalFormat},
		})
	}
	sort.Slice(refs, func(i, j int) bool {
		return refs[i].Kind < refs[j].Kind
	})
	return refs, nil
}

func formatDirectoryName(s string, serviceMappings []v1alpha1.ServiceMapping) string {
	nameToSm := make(map[string]v1alpha1.ServiceMapping)
	for _, sm := range serviceMappings {
		nameToSm[sm.Spec.Name] = sm
	}
	split := strings.Split(s, "-")
	ret := ""
	for i, v := range split {
		title := strings.Title(v)
		if _, ok := nameToSm[strings.ToUpper(title)]; ok {
			// value is IAM, KMS, etc, don't use strings.Title(v) as it will result in Kms or Iam
			title = strings.ToUpper(title)
		}
		for _, s := range allowedSpellings {
			if strings.ToLower(s) == strings.ToLower(title) {
				title = s
			}
		}
		if i != len(split)-1 {
			ret += title + " "
		} else {
			ret += title
		}
	}
	return ret
}

func resourceSupportsIAMPolicyAndPolicyMember(rc *v1alpha1.ResourceConfig) bool {
	return rc.IAMConfig.PolicyName != "" && rc.IAMConfig.PolicyMemberName != ""
}

func resourceSupportsIAMAuditConfig(rc *v1alpha1.ResourceConfig) bool {
	return rc.IAMConfig.AuditConfigName != ""
}

func clearGeneratedDocsDir() error {
	docsDir := repo.GetG3ResourceReferenceGeneratedPath()
	if err := os.RemoveAll(docsDir); err != nil {
		return fmt.Errorf("error deleting generated docs dir at %v: %v", docsDir, err)
	}
	if err := os.Mkdir(docsDir, 0700); err != nil {
		return fmt.Errorf("error recreating generated docs dir at %v: %v", docsDir, err)
	}
	return nil
}

func templateFileNameForGVK(gvk schema.GroupVersionKind) string {
	return strings.ToLower(strings.Join([]string{groupName(gvk), gvk.Kind}, "_")) + ".tmpl"
}

func crdFilePathForGVK(gvk schema.GroupVersionKind) string {
	return filepath.Join(repo.GetCRDsPath(), crdFileNameForGVK(gvk))
}

func crdFileNameForGVK(gvk schema.GroupVersionKind) string {
	crdName := fmt.Sprintf("%s.%s", text.Pluralize(gvk.Kind), gvk.Group)
	return strings.ToLower(strings.Join([]string{"apiextensions.k8s.io_v1_customresourcedefinition", crdName}, "_")) + ".yaml"
}

func sampleDirPathForGVK(gvk schema.GroupVersionKind) string {
	sampleDirName := strings.ToLower(gvk.Kind)
	return filepath.Join(repo.GetResourcesSamplesPath(), sampleDirName)
}

func generatedDocPathForGVK(gvk schema.GroupVersionKind) string {
	generatedDocsPath := repo.GetG3ResourceReferenceGeneratedPath()
	generatedDocDirName := groupName(gvk)
	generatedDocFileName := strings.ToLower(gvk.Kind) + ".md"
	return filepath.Join(generatedDocsPath, generatedDocDirName, generatedDocFileName)
}

func groupName(gvk schema.GroupVersionKind) string {
	return strings.SplitN(gvk.Group, ".", 2)[0]
}
