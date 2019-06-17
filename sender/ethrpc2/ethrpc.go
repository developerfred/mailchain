// Copyright 2019 Finobo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ethrpc2

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Options struct {
	Tx      *types.Transaction
	ChainID *big.Int
}

func New(address string) (*EthRPC2, error) {
	client, err := ethclient.Dial(address)
	if err != nil {
		return nil, err
	}
	return &EthRPC2{client: client}, nil
}

type EthRPC2 struct {
	client Client
}
