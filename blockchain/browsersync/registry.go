// Filesystem registry and backend options

package browsersync

import (
	"dao-client/blockchain/browsersync/scanlockpayment/polygon"
)

func Init() {
	//bsc.Init()
	//goerli.Init()
	//nbai.Init()
	polygon.Init()

}
