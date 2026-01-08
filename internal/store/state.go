package store

import "meegle-cli/internal/meegle"

type State struct {
	Tasks          []meegle.Task
	SubTasks       map[string][]meegle.SubTask
	SelectedTaskID string
	ReqSeq         int
	LatestReq      map[string]int
}

func NewState() State {
	return State{
		Tasks:    []meegle.Task{},
		SubTasks: map[string][]meegle.SubTask{},
		LatestReq: map[string]int{
			ReqFetchTasks:    0,
			ReqCreateTask:    0,
			ReqCreateSubTask: 0,
			ReqToggleSubTask: 0,
		},
	}
}

const (
	ReqFetchTasks    = "fetchTasks"
	ReqCreateTask    = "createTask"
	ReqCreateSubTask = "createSubTask"
	ReqToggleSubTask = "toggleSubTask"
)

func NextReqID(state *State, op string) int {
	state.ReqSeq++
	state.LatestReq[op] = state.ReqSeq
	return state.ReqSeq
}

func IsLatest(state State, op string, reqID int) bool {
	return state.LatestReq[op] == reqID
}
