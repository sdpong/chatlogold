package wechat

import (
	"context"
	"runtime"

	"github.com/sdpong/chatlogold/internal/errors"
	"github.com/sdpong/chatlogold/internal/wechat/model"
	"github.com/sdpong/chatlogold/internal/wechat/process"
)

var DefaultManager *Manager

func init() {
	DefaultManager = NewManager()
	DefaultManager.Load()
}

func Load() error {
	return DefaultManager.Load()
}

func GetAccount(name string) (*Account, error) {
	return DefaultManager.GetAccount(name)
}

func GetProcess(name string) (*model.Process, error) {
	return DefaultManager.GetProcess(name)
}

func GetAccounts() []*Account {
	return DefaultManager.GetAccounts()
}

// Manager 微信管理器
type Manager struct {
	detector   process.Detector
	accounts   []*Account
	processMap map[string]*model.Process
}

// NewManager 创建新的微信管理器
func NewManager() *Manager {
	return &Manager{
		detector:   process.NewDetector(runtime.GOOS),
		accounts:   make([]*Account, 0),
		processMap: make(map[string]*model.Process),
	}
}

// Load 加载微信进程信息
func (m *Manager) Load() error {
	// 查找微信进程
	processes, err := m.detector.FindProcesses()
	if err != nil {
		return err
	}

	// 转换为账号信息
	accounts := make([]*Account, 0, len(processes))
	processMap := make(map[string]*model.Process, len(processes))

	for _, p := range processes {
		account := NewAccount(p)

		accounts = append(accounts, account)
		if account.Name != "" {
			processMap[account.Name] = p
		}
	}

	m.accounts = accounts
	m.processMap = processMap

	return nil
}

// GetAccount 获取指定名称的账号
func (m *Manager) GetAccount(name string) (*Account, error) {
	p, err := m.GetProcess(name)
	if err != nil {
		return nil, err
	}
	return NewAccount(p), nil
}

func (m *Manager) GetProcess(name string) (*model.Process, error) {
	p, ok := m.processMap[name]
	if !ok {
		return nil, errors.WeChatAccountNotFound(name)
	}
	return p, nil
}

// GetAccounts 获取所有账号
func (m *Manager) GetAccounts() []*Account {
	return m.accounts
}

// DecryptDatabase 便捷方法：通过账号名解密数据库
func (m *Manager) DecryptDatabase(ctx context.Context, accountName, dbPath, outputPath string) error {
	// 获取账号
	account, err := m.GetAccount(accountName)
	if err != nil {
		return err
	}

	// 使用账号解密数据库
	return account.DecryptDatabase(ctx, dbPath, outputPath)
}
