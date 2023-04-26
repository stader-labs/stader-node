package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	main_1Raw()
}
func main_1() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64()) // 5671744
	// fmt.Println(block.Time().Uint64())       // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // 144
}

func main_1Raw() {
	x := "http://127.0.0.1:8545/eth_chainId"
	response, err := http.Get(x)
	if err != nil {
		panic(fmt.Sprintf("GET %+v", err))
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Get response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Sprintf("READALL %+v", err))
	}

	var m map[string]interface{}
	json.Unmarshal(body, &m)
	fmt.Printf("%+v", body)
	fmt.Printf("%+v", m)

}

func main_2() {
	x := "http://127.0.0.1:5052/eth/v1/beacon/genesis"
	response, err := http.Get(x)
	if err != nil {
		panic(fmt.Sprintf("GET %+v", err))
	}
	defer func() {
		_ = response.Body.Close()
	}()

	// Get response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Sprintf("READALL %+v", err))
	}

	var m map[string]interface{}
	json.Unmarshal(body, &m)
	fmt.Printf("%+v", body)
	fmt.Printf("%+v", m)

}
