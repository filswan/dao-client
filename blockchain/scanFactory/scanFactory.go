package scanFactory

import (
	"dao-client/blockchain/browsersync/scanlockpayment/bsc"
	"dao-client/blockchain/browsersync/scanlockpayment/goerli"
	"dao-client/blockchain/browsersync/scanlockpayment/nbai"
	"dao-client/blockchain/browsersync/scanlockpayment/polygon"
	"dao-client/common/constants"
)

type IEventScanFactory struct {
}

type BlockChainNetwork interface {
	ScanEventFromChainAndSaveDataToDb()
}

type BscBlockChainNetwork struct {
}

func (b BscBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	bsc.ScanEventFromChainAndSaveDataToDbForBsc()
}

type GoerliBlockChainNetwork struct {
}

func (g GoerliBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	goerli.ScanEventFromChainAndSaveDataToDbForGoerli()
}

type NBAIBlockChainNetwork struct {
}

func (n NBAIBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	nbai.ScanEventFromChainAndSaveDataToDbForNBAI()
}

type PolygonBlockChainNetwork struct {
}

func (p PolygonBlockChainNetwork) ScanEventFromChainAndSaveDataToDb() {
	polygon.ScanEventFromChainAndSaveDataToDbForPolygon()
}

func (*IEventScanFactory) GenerateBlockChainNetwork(network string) BlockChainNetwork {
	switch network {
	case constants.NETWORK_TYPE_BSC:
		return BscBlockChainNetwork{}
	case constants.NETWORK_TYPE_GOERLI:
		return GoerliBlockChainNetwork{}
	case constants.NETWORK_TYPE_NBAI:
		return NBAIBlockChainNetwork{}
	case constants.NETWORK_TYPE_POLYGON:
		return PolygonBlockChainNetwork{}
	}
	return nil
}
