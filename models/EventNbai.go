package models

import (
	"dao-client/common/constants"
	"dao-client/database"
)

type EventNbai struct {
	ID              int64  `json:"id"`
	TxHash          string `json:"tx_hash"`
	EventName       string `json:"event_name"`
	IndexedData     string `json:"indexed_data"`
	ContractName    string `json:"contract_name"`
	ContractAddress string `json:"contract_address"`
	BytesData       []byte `json:"bytes_data""`
	BlockNo         uint64 `json:"block_no"`
	AddressFrom     string `json:"address_from"`
	AddressTo       string `json:"address_to"`
	CreateAt        string `json:"create_at"`
}

// FindEvents (&Event{Id: "0xadeaCC802D0f2DFd31bE4Fa7434F15782Fd720ac"},"id desc","10","0")
func FindEventNbai(whereCondition interface{}, orderCondition, limit, offset string) ([]*EventNbai, error) {
	db := database.GetDB()
	if offset == "" {
		offset = "0"
	}
	if limit == "" {
		limit = constants.DEFAULT_SELECT_LIMIT
	}
	var models []*EventNbai
	err := db.Where(whereCondition).Offset(offset).Limit(limit).Order(orderCondition).Find(&models).Error
	return models, err
}
