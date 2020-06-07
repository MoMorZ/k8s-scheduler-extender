package controller

import (
	"log"
	"math/rand"

	v1 "k8s.io/api/core/v1"
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
)

const fitFormat = "[Info]Pod %v/%v filter:%v"
const luckyOne = "Hey man you are lucky!"
const unluckyOne = "Oh no you are unlucky!"

func filter(args schedulerapi.ExtenderArgs) *schedulerapi.ExtenderFilterResult {
	const Mod = 2

	pod := args.Pod
	nodes := args.Nodes.Items
	filteredNodes := make([]v1.Node, 0, len(args.Nodes.Items))
	failedNodes := make(schedulerapi.FailedNodesMap)

	for _, node := range nodes {
		if rand.Intn(Mod) == 0 {
			filteredNodes = append(filteredNodes, node)
			log.Printf(fitFormat, node.Name, pod.Namespace, luckyOne)
			continue
		}
		failedNodes[node.Name] = ""
		log.Printf(fitFormat, node.Name, pod.Namespace, unluckyOne)
	}
	return &schedulerapi.ExtenderFilterResult{
		Nodes: &v1.NodeList{
			Items: filteredNodes,
		},
		FailedNodes: failedNodes,
		Error:       "",
	}
}

// const (
// 	LuckyPred        = "Lucky"
// 	LuckyPredFailMsg = "Sorry, you're not lucky"
// )

// var predicatesFuncs = map[string]FitPredicate{
// 	LuckyPred: LuckyPredicate,
// }

// type FitPredicate func(pod *v1.Pod, node v1.Node) (bool, []string, error)

// var predicatesSorted = []string{LuckyPred}

// // filter 根据扩展程序定义的预选规则来过滤节点
// // it's webhooked to pkg/scheduler/core/generic_scheduler.go#findNodesThatFit()
// func filter(args schedulerapi.ExtenderArgs) *schedulerapi.ExtenderFilterResult {
// 	var filteredNodes []v1.Node
// 	failedNodes := make(schedulerapi.FailedNodesMap)
// 	pod := args.Pod
// 	for _, node := range args.Nodes.Items {
// 		fits, failReasons, _ := podFitsOnNode(pod, node)
// 		if fits {
// 			filteredNodes = append(filteredNodes, node)
// 		} else {
// 			failedNodes[node.Name] = strings.Join(failReasons, ",")
// 		}
// 	}

// 	result := schedulerapi.ExtenderFilterResult{
// 		Nodes: &v1.NodeList{
// 			Items: filteredNodes,
// 		},
// 		FailedNodes: failedNodes,
// 		Error:       "",
// 	}

// 	return &result
// }

// func podFitsOnNode(pod *v1.Pod, node v1.Node) (bool, []string, error) {
// 	fits := true
// 	var failReasons []string
// 	for _, predicateKey := range predicatesSorted {
// 		fit, failures, err := predicatesFuncs[predicateKey](pod, node)
// 		if err != nil {
// 			return false, nil, err
// 		}
// 		fits = fits && fit
// 		failReasons = append(failReasons, failures...)
// 	}
// 	return fits, failReasons, nil
// }

// func LuckyPredicate(pod *v1.Pod, node v1.Node) (bool, []string, error) {
// 	const Mod = 2
// 	lucky := rand.Intn(Mod) == 0
// 	if lucky {
// 		log.Printf("pod %v/%v is lucky to fit on node %v\n", pod.Name, pod.Namespace, node.Name)
// 		return true, nil, nil
// 	}
// 	log.Printf("pod %v/%v is unlucky to fit on node %v\n", pod.Name, pod.Namespace, node.Name)
// 	return false, []string{LuckyPredFailMsg}, nil
// }
