package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	v1 "bytedeath.com/zorder/pkg/apis/order/v1"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "kubeconfig file")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("failed to load kubeconfig file: %s", err.Error())
		os.Exit(1)
	}

	config.APIPath = "/apis/"
	config.NegotiatedSerializer = v1.Codecs.WithoutConversion()
	config.GroupVersion = &v1.GroupVersion

	client, err := rest.RESTClientFor(config)
	if err != nil {
		log.Fatalln(err)
	}

	foo := v1.Order{}
	err = client.Get().Namespace("default").Resource("orders").Name("order-test").Do(context.TODO()).Into(&foo)
	if err != nil {
		log.Fatalln(err)
	}

	newObj := foo.DeepCopy()
	newObj.Spec.Name = "expert"

	fmt.Println(foo.Spec.Name)
	fmt.Println(foo.Spec.Fee)

	fmt.Println(newObj.Spec.Name)
}
