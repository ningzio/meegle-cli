package meegle

import (
	"errors"
	"os"
	"sync"
	"time"
)

type Client interface {
	FetchTasks(projectKey string) ([]Task, error)
	CreateTask(projectKey, name string) (Task, error)
	CreateSubTask(projectKey, taskID, name string) (SubTask, error)
	ToggleSubTaskDone(projectKey, taskID, subTaskID string, done bool) (SubTask, error)
}

func NewClientFromEnv() Client {
	if envConfigured() {
		return NewRealClient(os.Getenv("MEEGLE_BASE_URL"), os.Getenv("MEEGLE_PLUGIN_ID"), os.Getenv("MEEGLE_PLUGIN_SECRET"))
	}
	return NewMockClient()
}

func envConfigured() bool {
	required := []string{
		"MEEGLE_BASE_URL",
		"MEEGLE_PLUGIN_ID",
		"MEEGLE_PLUGIN_SECRET",
		"MEEGLE_PROJECT_KEY",
		"MEEGLE_USER_KEY",
	}
	for _, key := range required {
		if os.Getenv(key) == "" {
			return false
		}
	}
	return true
}

type MockClient struct {
	mu       sync.Mutex
	tasks    []Task
	subTasks map[string][]SubTask
	nextID   int
}

func NewMockClient() *MockClient {
	return &MockClient{
		tasks: []Task{
			{ID: "t-1", Name: "Welcome to Meegle"},
			{ID: "t-2", Name: "Build your first TUI"},
		},
		subTasks: map[string][]SubTask{
			"t-1": {
				{ID: "s-1", TaskID: "t-1", Name: "Say hi to your team", Done: true},
				{ID: "s-2", TaskID: "t-1", Name: "Add one more task", Done: false},
			},
			"t-2": {
				{ID: "s-3", TaskID: "t-2", Name: "Explore Bubble Tea", Done: false},
			},
		},
		nextID: 3,
	}
}

func (m *MockClient) FetchTasks(projectKey string) ([]Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]Task{}, m.tasks...), nil
}

func (m *MockClient) CreateTask(projectKey, name string) (Task, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nextID++
	task := Task{ID: "t-" + itoa(m.nextID), Name: name}
	m.tasks = append([]Task{task}, m.tasks...)
	return task, nil
}

func (m *MockClient) CreateSubTask(projectKey, taskID, name string) (SubTask, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nextID++
	sub := SubTask{ID: "s-" + itoa(m.nextID), TaskID: taskID, Name: name, Done: false}
	m.subTasks[taskID] = append([]SubTask{sub}, m.subTasks[taskID]...)
	return sub, nil
}

func (m *MockClient) ToggleSubTaskDone(projectKey, taskID, subTaskID string, done bool) (SubTask, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	list := m.subTasks[taskID]
	for i, item := range list {
		if item.ID == subTaskID {
			item.Done = done
			list[i] = item
			m.subTasks[taskID] = list
			return item, nil
		}
	}
	return SubTask{}, errors.New("subtask not found")
}

type RealClient struct {
	baseURL      string
	pluginID     string
	pluginSecret string
	mu           sync.Mutex
	auth         *authCache
}

func NewRealClient(baseURL, pluginID, pluginSecret string) *RealClient {
	return &RealClient{
		baseURL:      baseURL,
		pluginID:     pluginID,
		pluginSecret: pluginSecret,
		auth:         &authCache{expiresAt: time.Now().Add(-time.Hour)},
	}
}

func (r *RealClient) FetchTasks(projectKey string) ([]Task, error) {
	return nil, errors.New("real client not implemented; set env vars for mock or implement API calls")
}

func (r *RealClient) CreateTask(projectKey, name string) (Task, error) {
	return Task{}, errors.New("real client not implemented; set env vars for mock or implement API calls")
}

func (r *RealClient) CreateSubTask(projectKey, taskID, name string) (SubTask, error) {
	return SubTask{}, errors.New("real client not implemented; set env vars for mock or implement API calls")
}

func (r *RealClient) ToggleSubTaskDone(projectKey, taskID, subTaskID string, done bool) (SubTask, error) {
	return SubTask{}, errors.New("real client not implemented; set env vars for mock or implement API calls")
}

func itoa(num int) string {
	if num == 0 {
		return "0"
	}
	buf := [20]byte{}
	pos := len(buf)
	for num > 0 {
		pos--
		buf[pos] = byte('0' + num%10)
		num /= 10
	}
	return string(buf[pos:])
}
