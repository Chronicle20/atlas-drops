package drop

import "time"

const (
	EnvEventTopicDropStatus           = "EVENT_TOPIC_DROP_STATUS"
	StatusEventTypeCreated            = "CREATED"
	StatusEventTypeExpired            = "EXPIRED"
	StatusEventTypePickedUp           = "PICKED_UP"
	StatusEventTypeReserved           = "RESERVED"
	StatusEventTypeReservationFailure = "RESERVATION_FAILURE"
)

type statusEvent[E any] struct {
	WorldId   byte   `json:"worldId"`
	ChannelId byte   `json:"channelId"`
	MapId     uint32 `json:"mapId"`
	DropId    uint32 `json:"dropId"`
	Type      string `json:"type"`
	Body      E      `json:"body"`
}

type createdStatusEventBody struct {
	ItemId          uint32    `json:"itemId"`
	Quantity        uint32    `json:"quantity"`
	Meso            uint32    `json:"meso"`
	Type            byte      `json:"type"`
	X               int16     `json:"x"`
	Y               int16     `json:"y"`
	OwnerId         uint32    `json:"ownerId"`
	OwnerPartyId    uint32    `json:"ownerPartyId"`
	DropTime        time.Time `json:"dropTime"`
	DropperUniqueId uint32    `json:"dropperUniqueId"`
	DropperX        int16     `json:"dropperX"`
	DropperY        int16     `json:"dropperY"`
	PlayerDrop      bool      `json:"playerDrop"`
}

type expiredStatusEventBody struct {
}

type pickedUpStatusEventBody struct {
	CharacterId uint32 `json:"characterId"`
	ItemId      uint32 `json:"itemId"`
	EquipmentId uint32 `json:"equipmentId"`
	Quantity    uint32 `json:"quantity"`
	Meso        uint32 `json:"meso"`
	PetSlot     int8   `json:"petSlot"`
}

type reservedStatusEventBody struct {
	CharacterId uint32 `json:"characterId"`
	ItemId      uint32 `json:"itemId"`
	EquipmentId uint32 `json:"equipmentId"`
	Quantity    uint32 `json:"quantity"`
	Meso        uint32 `json:"meso"`
}

type reservationFailureStatusEventBody struct {
	CharacterId uint32 `json:"characterId"`
}
