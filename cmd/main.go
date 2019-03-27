package main

import "../core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("Send 1BTC to Jacky")
	bc.SendData("Send 1EOS to Jack")
	bc.SendData("Send 1EOS to Tony")
	bc.Print()
}
