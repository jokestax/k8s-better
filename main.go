package main

import (
	"context"
	"os"

	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
)

type CRD struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   CRDSpec
	Status CRDStatus
}

type CRDSpec struct {
	Group string
	// Scope   ResourceScope
	Names   CRDNames
	Version CRDVersion
}

type CRDNames struct {
	Plural     string
	Singular   string
	Kind       string
	ShortNames []string
}

type CRDVersion struct {
	Name               string
	Served             bool
	Storage            bool
	Deprecated         bool
	DeprecationWarning *string
	// Schema                   *CRDSchema
	// Subresources             *CRDSubresources
	// AdditionalPrinterColumns []CRDPrinterColumn
}

type CRDStatus struct {
}

func main() {

	config, err := clientcmd.BuildConfigFromFlags("", os.Getenv("KUBECONFIG"))
	if err != nil {
		panic(err)
	}

	clientset, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	list, err := clientset.ApiextensionsV1().CustomResourceDefinitions().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, crd := range list.Items {
		println(crd.Name)
	}
}
