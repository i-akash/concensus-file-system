package entity

import (
	"rfs/bclib"
	"rfs/config"
	"rfs/secsuit"
	"strconv"
	"time"
)

type Block struct {
	PrevHash   string
	Operations []*Operation
	MinerID    int
	Nonce      int
	TimeStamp  time.Time
	SerialNo   int
}

func (block *Block) String() string {
	//Todo: Need to implement it appropriately like markle root
	str := ""
	str += " " + block.PrevHash
	str += " " + strconv.Itoa(block.MinerID)
	str += " " + strconv.Itoa(block.Nonce)
	str += " " + block.TimeStamp.String()
	str += " " + strconv.Itoa(block.SerialNo)

	for _, operation := range block.Operations {
		str += " " + operation.String()
	}

	return str
}

func (block *Block) Hash() string {
	return secsuit.ComputeHash(block.String())
}

func NewOpBlock(prevblock *Block, operations []*Operation) *Block {

	time.Sleep(time.Duration(bclib.Random(40, 60)) * time.Second)

	config := config.GetSingletonConfigHandler()

	return &Block{
		PrevHash:   prevblock.Hash(),
		Operations: operations,
		MinerID:    config.MinerConfig.MinerId,
		TimeStamp:  time.Now(),
		SerialNo:   prevblock.SerialNo + 1,
	}
}

func NewNoOpBlock(prevblock *Block) *Block {

	time.Sleep(time.Duration(bclib.Random(20, 40)) * time.Second)

	config := config.GetSingletonConfigHandler()

	return &Block{
		PrevHash:  prevblock.Hash(),
		MinerID:   config.MinerConfig.MinerId,
		TimeStamp: time.Now(),
		SerialNo:  prevblock.SerialNo + 1,
	}
}

func CreateGenesisBlock() *Block {
	//Todo: Proper implement - like config
	return &Block{
		SerialNo: 1,
	}
}
