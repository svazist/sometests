package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	//"k8s.io/apimachinery/pkg/api/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//"time"
)

func init() {

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
		nodes, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})

		nodeip := []corev1.NodeAddress{}
		conditions := []corev1.NodeCondition{}
		for i := 0; i < len(nodes.Items); i++ {
			nodeip = nodes.Items[i].Status.Addresses
			conditions = nodes.Items[i].Status.Conditions

			fmt.Printf("NodeIp: %s NodeName: %s\n", nodeip[0].Address, nodeip[1].Address)
			fmt.Printf("Conditions: %s \n", conditions)
		}

		fmt.Printf("Kube test config file  %s\n", cubeConfig)

	},
}
