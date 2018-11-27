package message

type RequestBody struct {
	File   []byte
	Params []interface{}
}

type RequestHeader struct {
	Context          string
	RequestID        int
	ResponseExpected bool
	ObjectKey        int
	Compressed       bool
	Operation        string
}

type ReplyBody struct {
	ConvertedFile []byte
	ClientProxy   interface{}
	ServiceName   string
}

type ReplyHeader struct {
	ServiceContext string
	RequestID      int
	ReplyStatus    int
}

type MessageBody struct {
	RequestBody   RequestBody
	RequestHeader RequestHeader
	ReplyHeader   ReplyHeader
	ReplyBody     ReplyBody
}
