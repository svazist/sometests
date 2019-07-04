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

var (
	namespace string = ""
)

func init() {

	kubeServicesCmd.Flags().StringVar(&namespace, "namespace", "", "limit serices by namespace")

	RootCmd.AddCommand(kubeServicesCmd)

}

var kubeServicesCmd = &cobra.Command{
	Use:   "services",
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
		services, err := clientset.CoreV1().Services(namespace).List(metav1.ListOptions{})

		for i := 0; i < len(services.Items); i++ {

			name := services.Items[i].ObjectMeta.Name
			namespace = services.Items[i].ObjectMeta.Namespace
			spec := services.Items[i].Spec
			ip := "undefined"
			if spec.Type == "ClusterIP" {
				ip = spec.ClusterIP
			}

			fmt.Printf("Service Name: %s  Namespace: %s IP: %s\n", name, namespace, ip)
			//fmt.Printf("Service status %s\n", spec)

		}

		fmt.Printf("Kube test config file  %s\n", cubeConfig)

	},
}
