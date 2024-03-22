<!-- PROJECT LOGO -->
<br />
<p align="center">
  <a href="https://github.com/6boris/web3-go" target="_blank">
    <img src="./assets/img/Ethereum-icon-purple.svg" alt="Logo" width="680" height="256">
  </a>

  <h3 align="center">Web3 Go</h3>

  <p align="center">
    Ethereum Dapp Go API, inspired by 
    <a href="https://github.com/ChainSafe/web3.js" target="_blank">web3.js</a>.
    <br />
    <a href="https://github.com/6boris/web3-go/issues" target="_blank">Report Bug</a>
    ·
    <a href="https://github.com/6boris/web3-go/pulls" target="_blank">Pull Request</a>
  </p>
</p>

[![WEBSITE](https://img.shields.io/badge/Web3-Go-brightgreen)](https://github.com/kylesliu/web3-go)
[![LISTENSE](https://img.shields.io/github/license/6boris/web3-go)](https://github.com/kylesliu/web3-go/blob/main/LICENSE)

## Introduction

This is the Ethereum [Golang API](https://github.com/kylesliu/web3-go) which connects to the Generic [JSON-RPC](https://github.com/ethereum/wiki/wiki/JSON-RPC) spec.

You need to run a local or remote Ethereum node to use this library.

Here is an open source case [Web3 Studio](https://web3-studio.leek.dev/d/demo/web3-studio) reference under.

<a href="https://web3-studio.leek.dev/d/demo/web3-studio" target="_blank">
  <img src="https://s.gin.sh/develop/web3/web3-studio-demo.png" alt="Logo">
</a>


### Client

```bash
go get github.com/6boris/web3-go
```
```go
package main

import (
  "context"
  "fmt"
  "github.com/6boris/web3-go/client"
  clientModel "github.com/6boris/web3-go/model/client"
)

func main() {
  ec := client.NewPool(clientModel.GetDefaultConfPool())
  chainID := int64(1)
  callResp, err := ec.GetEvmClient(chainID).SuggestGasTipCap(context.Background())
  if err != nil {
    panic(err)
  }
  fmt.Println(fmt.Sprintf("ChainID:%d GasPrice:%s", chainID, callResp.String()))
}
/*
Output:
    ChainID:1 GasPrice:10670000
*/
```

## Development Trips
- [X] Client
  - [ ] Base Method
    - [X] eth_chainId
    - [X] web3_clientVersion
    - [X] eth_gasPrice
    - [X] eth_blockNumber
    - [X] eth_getBalance
    - [ ] ...
  - [ ] Middleware
    - [X] LoadBalance
    - [X] Metrics
    - [ ] Grafana
    - [ ] CircuitBreaker
  - [ ] Business Cases
    - [ ] Web3 Studio
- [ ] Other ...



## Community

- [web3.js](https://github.com/ChainSafe/web3.js) Ethereum JavaScript API.
- [Web3j](https://github.com/web3j/web3j) Web3 Java Ethereum Ðapp API.
- [Web3.py](https://github.com/ethereum/web3.py) A Python library for interacting with Ethereum.

## Provider
- https://public.blockpi.io/


## Dev tool

- [JSON RPC](https://www.jsonrpc.org/specification) Defining the JSON RPC specification.
- [Go Ethereum](https://github.com/ethereum/go-ethereum) Official Golang implementation of the Ethereum protocol.
- [Ethereum 1.0 API](https://github.com/ethereum/eth1.0-apis) Ethereum JSON-RPC Specification.
