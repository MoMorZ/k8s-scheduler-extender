package controller

import (
	"log"

	v1 "k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

const prioFormat = "[Info]Pod %v/%v score: %v"

func prioritize(args schedulerapi.ExtenderArgs) *schedulerapi.HostPriorityList {
	pod := args.Pod
	nodes := args.Nodes.Items

	hostPriorityList := make(schedulerapi.HostPriorityList, len(nodes))
	for i, node := range nodes {
		tNode := node
		hostPriorityList[i] = func(node *v1.Node) schedulerapi.HostPriority {
			cpu := node.Status.Allocatable.Cpu().MilliValue()
			mem := node.Status.Allocatable.Memory().MilliValue()
			return schedulerapi.HostPriority{
				Host:  node.Name,
				Score: int(cpu * mem),
			}
		}(&tNode)
		log.Printf(prioFormat, hostPriorityList[i].Host, pod.Namespace, hostPriorityList[i].Score)
	}

	return &hostPriorityList
}
