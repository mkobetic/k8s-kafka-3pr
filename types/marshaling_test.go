package types

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"k8s.io/client-go/pkg/api"
)

func Test_Unmarshaling(t *testing.T) {
	var topics KafkaTopicList
	err := json.Unmarshal([]byte(
		`{"kind":"KafkaTopicList", "items":[
            {"apiVersion":"shopify.io/v1","description":"Test topic 1","kind":"KafkaTopic",
                "metadata":{"name":"test1","namespace":"default","selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics/test1","uid":"463a29e0-9ab9-11e6-9a7b-42010af000bd","resourceVersion":"2618096","creationTimestamp":"2016-10-25T13:45:21Z"}},
            {"apiVersion":"shopify.io/v1","description":"Test topic 2","kind":"KafkaTopic",
                "metadata":{"name":"test2","namespace":"default","selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics/test2","uid":"4640bc21-9ab9-11e6-9a7b-42010af000bd","resourceVersion":"2618097","creationTimestamp":"2016-10-25T13:45:21Z"}},
            {"apiVersion":"shopify.io/v1","description":"Test topic 3","kind":"KafkaTopic",
                "metadata":{"name":"test3","namespace":"default","selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics/test3","uid":"46473d13-9ab9-11e6-9a7b-42010af000bd","resourceVersion":"2618098","creationTimestamp":"2016-10-25T13:45:22Z"}}],
            "metadata":{"selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics","resourceVersion":"2813870"},"apiVersion":"shopify.io/v1"}`),
		&topics)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%#v\n", topics)
	if len(topics.Items) != 3 {
		t.Fatal(topics.Items)
	}
	for i := 0; i < 3; i++ {
		if topic := topics.Items[i]; topic.GetName() != fmt.Sprintf("test%d", i+1) {
			t.Fatal(i, topic)
		}
	}
}

func Test_Decoding(t *testing.T) {
	s, ok := api.Codecs.SerializerForMediaType("application/json", nil)
	if !ok {
		t.Fail()
	}
	ro, gvk, err := s.Decode(
		[]byte(`{"kind":"KafkaTopicList", "items":[
            {"apiVersion":"shopify.io/v1","description":"Test topic 1","kind":"KafkaTopic",
                "metadata":{"name":"test1","namespace":"default","selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics/test1","uid":"463a29e0-9ab9-11e6-9a7b-42010af000bd","resourceVersion":"2618096","creationTimestamp":"2016-10-25T13:45:21Z"}},
            {"apiVersion":"shopify.io/v1","description":"Test topic 2","kind":"KafkaTopic",
                "metadata":{"name":"test2","namespace":"default","selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics/test2","uid":"4640bc21-9ab9-11e6-9a7b-42010af000bd","resourceVersion":"2618097","creationTimestamp":"2016-10-25T13:45:21Z"}},
            {"apiVersion":"shopify.io/v1","description":"Test topic 3","kind":"KafkaTopic",
                "metadata":{"name":"test3","namespace":"default","selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics/test3","uid":"46473d13-9ab9-11e6-9a7b-42010af000bd","resourceVersion":"2618098","creationTimestamp":"2016-10-25T13:45:22Z"}}],
            "metadata":{"selfLink":"/apis/shopify.io/v1/namespaces/default/kafkatopics","resourceVersion":"2813870"},"apiVersion":"shopify.io/v1"}`), nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(gvk)
	spew.Dump(ro)
	topics := ro.(*KafkaTopicList)
	fmt.Printf("%#v\n", topics)
	if len(topics.Items) != 3 {
		t.Fatal(topics.Items)
	}
	for i := 0; i < 3; i++ {
		if topic := topics.Items[i]; topic.GetName() != fmt.Sprintf("test%d", i+1) {
			t.Fatal(i, topic)
		}
	}
}
