package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/mkobetic/k8s-kafka-3pr/types"

	"k8s.io/client-go/pkg/api"
	"k8s.io/client-go/pkg/api/unversioned"
	"k8s.io/client-go/pkg/runtime/serializer"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig = flag.String("kubeconfig", "$HOME/.kube/config", "absolute path to the kubeconfig file")
)

func main() {
	flag.Parse()
	resource := flag.Arg(0)
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", os.ExpandEnv(*kubeconfig))
	if err != nil {
		panic(err.Error())
	}
	config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
		return &loggingRoundTripper{rt}
	}
	config.APIPath = "/apis"
	config.GroupVersion = &unversioned.GroupVersion{Group: "shopify.io", Version: "v1"}

	// Clientset => can't access custom api groups (have to generate client first)
	// topics, err := clientset.Extensions().ThirdPartyResources().Get("kafka-topic.shopify.io")

	// Dynamic Client => generic UnstructeredList/Unstructured result
	// client, err := dynamic.NewClient(config)
	// if err != nil {
	// 	panic(err)
	// }
	// r := &unversioned.APIResource{Name: resource, Namespaced: false}
	// topics, err := client.Resource(r, "").List(nil)

	// Bare RESTClient => needs JSON marshaling code generated for custom types?
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}
	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	r := client.Get().Namespace("default").Resource(resource)
	fmt.Println(r.URL())
	res := r.Do()
	obj, err := res.Get()
	if err != nil {
		panic(err.Error())
	}
	topics := obj.(*types.KafkaTopicList)
	spew.Dump(topics)
	fmt.Printf("%#v\n", topics)
}
