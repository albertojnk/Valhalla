package npc

import (
	"github.com/albertojnk/Valhalla/constant/opcode"
	"github.com/albertojnk/Valhalla/mpacket"
)

func packetNpcSetController(npcID int32, isLocal bool) mpacket.Packet {
	p := mpacket.CreateWithOpcode(opcode.SendChannelNpcControl)
	p.WriteBool(isLocal)
	p.WriteInt32(npcID)

	return p
}

func packetNpcMovement(bytes []byte) mpacket.Packet {
	p := mpacket.CreateWithOpcode(opcode.SendChannelNpcMovement)
	p.WriteBytes(bytes)

	return p
}
