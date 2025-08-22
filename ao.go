package aogo

import (
	"github.com/liteseed/goar/signer"
	"github.com/liteseed/goar/tag"
	"github.com/liteseed/goar/transaction/data_item"
)

const (
	MuUrl     = "https://mu.ao-testnet.xyz"
	CuUrl     = "https://cu.ao-testnet.xyz"
	SCHEDULER = "_GQ33BkPtZrqxA84vM8Zk-N2aO0toNNu_C-l-rawrBA"
	GATEWAY   = "https://arweave.net"

	SDK = "aogo"
)

type AO struct {
	mu MU
	cu CU
}

type Message struct {
	ID     string     `json:"Id"`
	Target string     `json:"Target"`
	Owner  string     `json:"Owner"`
	Data   any        `json:"Data"`
	Tags   *[]tag.Tag `json:"Tags"`
}

func New(options ...func(*AO)) (*AO, error) {
	ao := &AO{cu: newCU(CuUrl), mu: newMU(MuUrl)}
	for _, o := range options {
		o(ao)
	}
	return ao, nil
}

func WthMU(url string) func(*AO) {
	return func(ao *AO) {
		ao.mu = newMU(url)
	}
}

func WthCU(url string) func(*AO) {
	return func(ao *AO) {
		ao.cu = newCU(url)
	}
}

// MU Functions

func (ao *AO) GenerateMessage(process string, data string, tags *[]tag.Tag, anchor string, s *signer.Signer) (*data_item.DataItem, error) {
	return ao.mu.GenerateMessage(process, data, tags, anchor, s)
}

func (ao *AO) SendMessageDataItem(dataItem *data_item.DataItem) (string, error) {
	return ao.mu.SendMessageDataItem(dataItem)
}

func (ao *AO) GenerateProcess(module string, data []byte, tags []tag.Tag, s *signer.Signer) (*data_item.DataItem, error) {
	return ao.mu.GenerateProcess(module, data, tags, s)
}

func (ao *AO) SendProcessDataItem(dataItem *data_item.DataItem) (string, error) {
	return ao.mu.SendProcessDataItem(dataItem)
}

func (ao *AO) SpawnProcess(module string, data []byte, tags []tag.Tag, s *signer.Signer) (string, error) {
	return ao.mu.SpawnProcess(module, data, tags, s)
}

func (ao *AO) SendMessage(process string, data string, tags *[]tag.Tag, anchor string, s *signer.Signer) (string, error) {
	return ao.mu.SendMessage(process, data, tags, anchor, s)
}

// CU Functions

func (ao *AO) LoadResult(process string, message string) (*Response, error) {
	return ao.cu.LoadResult(process, message)
}

func (ao *AO) DryRun(message Message) (*Response, error) {
	return ao.cu.DryRun(message)
}
