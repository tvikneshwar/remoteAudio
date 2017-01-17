package events

import (
	"os"
	"os/signal"

	"github.com/cskr/pubsub"
)

const (
	RxAudioOn      = "audioOn"
	TxUser         = "txUser"
	MqttConnStatus = "mqttConnStatus"
	Shutdown       = "shutdown"
)

func WatchSystemEvents(evPS *pubsub.PubSub) {

	// Channel to handle OS signals
	osSignals := make(chan os.Signal, 1)

	//subscribe to os.Interrupt (CTRL-C signal)
	signal.Notify(osSignals, os.Interrupt)

	select {
	case osSignal := <-osSignals:
		if osSignal == os.Interrupt {
			evPS.Pub(true, Shutdown)
		}
	}
}