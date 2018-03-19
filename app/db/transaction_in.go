package db

import (
	"encoding/hex"
	"git.jasonc.me/main/bitcoin/wallet"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"strings"
	"time"
)

type TransactionIn struct {
	Id                    uint   `gorm:"primary_key"`
	TransactionId         uint
	PreviousOutPointHash  []byte
	PreviousOutPointIndex uint32
	SignatureScript       []byte `gorm:"unique;"`
	UnlockString          string
	Witnesses             []*Witness
	Sequence              uint32
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

func (t TransactionIn) GetOutPoint() *wire.OutPoint {
	hash, _ := chainhash.NewHash(t.PreviousOutPointHash)
	return wire.NewOutPoint(hash, t.PreviousOutPointIndex)
}

func (t TransactionIn) GetPrevOutPointString() string {
	return t.GetOutPoint().String()
}

func (t TransactionIn) GetAddress() string {
	split := strings.Split(t.UnlockString, " ")
	if len(split) != 2 {
		return ""
	}
	pubKey, err := hex.DecodeString(split[1])
	if err != nil {
		return ""
	}
	return wallet.GetAddress(pubKey).GetEncoded()
}
