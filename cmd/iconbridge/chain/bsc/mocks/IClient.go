// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	bmcperiphery "github.com/icon-project/icon-bridge/cmd/e2etest/chain/bsc/abi/bmcperiphery"

	bsctypes "github.com/icon-project/icon-bridge/cmd/iconbridge/chain/bsc/types"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	ethereum "github.com/ethereum/go-ethereum"

	log "github.com/icon-project/icon-bridge/common/log"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// IClient is an autogenerated mock type for the IClient type
type IClient struct {
	mock.Mock
}

// CallContract provides a mock function with given fields: ctx, msg, blockNumber
func (_m *IClient) CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, msg, blockNumber)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, msg, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, msg, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: ctx, q
func (_m *IClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(ctx, q)

	var r0 []types.Log
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBalance provides a mock function with given fields: ctx, hexAddr
func (_m *IClient) GetBalance(ctx context.Context, hexAddr string) (*big.Int, error) {
	ret := _m.Called(ctx, hexAddr)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context, string) *big.Int); ok {
		r0 = rf(ctx, hexAddr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, hexAddr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockByHash provides a mock function with given fields: hash
func (_m *IClient) GetBlockByHash(hash common.Hash) (*bsctypes.Block, error) {
	ret := _m.Called(hash)

	var r0 *bsctypes.Block
	if rf, ok := ret.Get(0).(func(common.Hash) *bsctypes.Block); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bsctypes.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockNumber provides a mock function with given fields:
func (_m *IClient) GetBlockNumber() (uint64, error) {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockReceipts provides a mock function with given fields: hash
func (_m *IClient) GetBlockReceipts(hash common.Hash) (types.Receipts, error) {
	ret := _m.Called(hash)

	var r0 types.Receipts
	if rf, ok := ret.Get(0).(func(common.Hash) types.Receipts); ok {
		r0 = rf(hash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Receipts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(common.Hash) error); ok {
		r1 = rf(hash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChainID provides a mock function with given fields:
func (_m *IClient) GetChainID() *big.Int {
	ret := _m.Called()

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func() *big.Int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetHeaderByHeight provides a mock function with given fields: ctx, height
func (_m *IClient) GetHeaderByHeight(ctx context.Context, height *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, height)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMedianGasPriceForBlock provides a mock function with given fields: ctx
func (_m *IClient) GetMedianGasPriceForBlock(ctx context.Context) (*big.Int, *big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 *big.Int
	if rf, ok := ret.Get(1).(func(context.Context) *big.Int); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*big.Int)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetStatus provides a mock function with given fields: opts, _link
func (_m *IClient) GetStatus(opts *bind.CallOpts, _link string) (bmcperiphery.TypesLinkStats, error) {
	ret := _m.Called(opts, _link)

	var r0 bmcperiphery.TypesLinkStats
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, string) bmcperiphery.TypesLinkStats); ok {
		r0 = rf(opts, _link)
	} else {
		r0 = ret.Get(0).(bmcperiphery.TypesLinkStats)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.CallOpts, string) error); ok {
		r1 = rf(opts, _link)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HandleRelayMessage provides a mock function with given fields: opts, _prev, _msg
func (_m *IClient) HandleRelayMessage(opts *bind.TransactOpts, _prev string, _msg []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, _prev, _msg)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, string, []byte) *types.Transaction); ok {
		r0 = rf(opts, _prev, _msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, string, []byte) error); ok {
		r1 = rf(opts, _prev, _msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Log provides a mock function with given fields:
func (_m *IClient) Log() log.Logger {
	ret := _m.Called()

	var r0 log.Logger
	if rf, ok := ret.Get(0).(func() log.Logger); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log.Logger)
		}
	}

	return r0
}

// NonceAt provides a mock function with given fields: ctx, account, blockNumber
func (_m *IClient) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	ret := _m.Called(ctx, account, blockNumber)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) uint64); ok {
		r0 = rf(ctx, account, blockNumber)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseMessage provides a mock function with given fields: _a0
func (_m *IClient) ParseMessage(_a0 types.Log) (*bmcperiphery.BmcperipheryMessage, error) {
	ret := _m.Called(_a0)

	var r0 *bmcperiphery.BmcperipheryMessage
	if rf, ok := ret.Get(0).(func(types.Log) *bmcperiphery.BmcperipheryMessage); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bmcperiphery.BmcperipheryMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *IClient) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionByHash provides a mock function with given fields: ctx, blockHash
func (_m *IClient) TransactionByHash(ctx context.Context, blockHash common.Hash) (*types.Transaction, bool, error) {
	ret := _m.Called(ctx, blockHash)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Transaction); ok {
		r0 = rf(ctx, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) bool); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, common.Hash) error); ok {
		r2 = rf(ctx, blockHash)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TransactionCount provides a mock function with given fields: ctx, blockHash
func (_m *IClient) TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error) {
	ret := _m.Called(ctx, blockHash)

	var r0 uint
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) uint); ok {
		r0 = rf(ctx, blockHash)
	} else {
		r0 = ret.Get(0).(uint)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionInBlock provides a mock function with given fields: ctx, blockHash, index
func (_m *IClient) TransactionInBlock(ctx context.Context, blockHash common.Hash, index uint) (*types.Transaction, error) {
	ret := _m.Called(ctx, blockHash, index)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash, uint) *types.Transaction); ok {
		r0 = rf(ctx, blockHash, index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash, uint) error); ok {
		r1 = rf(ctx, blockHash, index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionReceipt provides a mock function with given fields: ctx, txHash
func (_m *IClient) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(ctx, txHash)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, common.Hash) error); ok {
		r1 = rf(ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}