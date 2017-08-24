package utils

import "github.com/asaskevich/EventBus"

var RouterBus EventBus.Bus

func init() {
	RouterBus = EventBus.New()
}
