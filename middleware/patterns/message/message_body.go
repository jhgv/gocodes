package message

type RequestBody struct {
	Parameters []string
}

type RequestHeader struct {
	Context          string
	RequestID        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

type ReplyBody struct {
	OperationResult string
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
