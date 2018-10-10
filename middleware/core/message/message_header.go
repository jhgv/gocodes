package message

type MessageHeader struct {
	Magic       string
	Version     int
	ByteOrder   bool
	MessageType int
	MessageSize float64
}
