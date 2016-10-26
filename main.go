package main

import (
	"flag"
	"fmt"
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
	config.APIPath = "/apis"
	config.GroupVersion = &unversioned.GroupVersion{Group: "shopify.io", Version: "v1"}
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: api.Codecs}
	client, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}
	// topics, err := clientset.Extensions().ThirdPartyResources().Get("kafka-topic.shopify.io")

	r := client.Get().Namespace("default").Resource(resource)
	fmt.Println(r.URL())
	topics, err := r.Do().Get()
	if err != nil {
		panic(err.Error())
	}
	spew.Dump(topics.(*types.KafkaTopicList))
}
