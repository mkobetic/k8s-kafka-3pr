package types

import (
	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/apimachinery/announced"
	"k8s.io/client-go/pkg/runtime"
)

func init() {
	if err := announced.NewGroupMetaFactory(
		&announced.GroupMetaFactoryArgs{
			GroupName:              GroupName,
			VersionPreferenceOrder: []string{GroupVersion.Version},
			ImportPrefix:           "github.com/mkobetic/k8s-kafka-3pr",
			// RootScopedKinds:        sets.NewString("KafkaTopic", "KafkaTopicList"),
			// AddInternalObjectsToScheme: AddToScheme,
		},
		announced.VersionToSchemeFunc{
			GroupVersion.Version: AddToScheme,
		},
	).Announce().RegisterAndEnable(); err != nil {
		panic(err)
	}
}

// GroupName is the group name use in this package
const GroupName = "shopify.io"

// GroupVersion is group version used to register these objects
var GroupVersion = unversioned.GroupVersion{Group: GroupName, Version: "v1"}

// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return GroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) unversioned.GroupResource {
	return GroupVersion.WithResource(resource).GroupResource()
}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
	AddToScheme   = SchemeBuilder.AddToScheme
)

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	// TODO this gets cleaned up when the types are fixed
	scheme.AddKnownTypes(GroupVersion,
		&KafkaTopicList{},
		&KafkaTopic{},
	)
	return nil
}

type KafkaTopicList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`
	Items                []KafkaTopic `json:"items"`
}

type KafkaTopic struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
	api.ObjectMeta `json:"metadata,omitempty"`
}
