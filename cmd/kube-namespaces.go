package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	//"k8s.io/apimachinery/pkg/api/errors"
	//corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"time"
)

func init() {

	RootCmd.AddCommand(kubeNamespacesCmd)

}

var kubeNamespacesCmd = &cobra.Command{
	Use:   "namespaces",
	Short: "test kubernates client",
	Long:  `test kubernates client`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(cubeConfig) < 1 {
			fmt.Println("Need set config file ")
			return
		}

		config, err := clientcmd.BuildConfigFromFlags("", cubeConfig)
		if err != nil {
			panic(err.Error())
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}
		namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})

		for i := 0; i < len(namespaces.Items); i++ {

			namespace := namespaces.Items[i]
			fmt.Printf("Namespace Name: %s \n", namespace.ObjectMeta.Name)
		}

		fmt.Printf("Kube test config file  %s\n", cubeConfig)

	},
}
