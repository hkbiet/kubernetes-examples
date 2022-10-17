package Pods

import (
	"context"
	"flag"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func GetPods() {
	var kubeconfig *string
	home := homedir.HomeDir()

	if home != "" {

		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")

	} else {

		kubeconfig = flag.String("kubeconfig", "", "absolute path to kubeconfig file")

	}

	flag.Parse()

	// build config to be used in clientSet
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("Unable to get config..")
		panic(err.Error())
	}

	// create clientSet

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Unable to create clientSet...")
		panic(err.Error())
	}

	// get current Pods count from all namespaces
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), v1.ListOptions{})
	if err != nil {
		fmt.Println("Unable to get POds..")
		panic(err.Error())
	}

	fmt.Printf("There are %d Pods in the cluster.", len(pods.Items))

}
