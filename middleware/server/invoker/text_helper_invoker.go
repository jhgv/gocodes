package invoker

import (
	"encoding/json"
	"log"

	"github.com/jhgv/gocodes/middleware/utils/constants"

	"github.com/jhgv/gocodes/middleware/core/message"
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/core/requestor"
	"github.com/jhgv/gocodes/middleware/server/handler"
	"github.com/jhgv/gocodes/middleware/server/objects"
)

type TextHelperInvoker struct {
}

func (thi *TextHelperInvoker) Invoke(cp *proxy.TextHelperProxy) {
	srh := new(handler.TCPServerRequestHanlder)
	srh.SetupSocket(cp.GetHost(), cp.GetPort())

	var msgToBeUnmarshalled []byte
	var msgUnmarshalled = message.Message{}

	var termination requestor.Termination
	var textHelper = &objects.TextHelper{}

	var replyHandler = func(resultText string) {
		// Unmarshalling result
		termination.SetResult(resultText)
		// Creating message
		messageToBeMarshalled := &message.Message{
			Test: "hello",
			Body: message.MessageBody{
				ReplyHeader: message.ReplyHeader{ReplyStatus: 0, RequestID: 0, ServiceContext: ""},
				ReplyBody:   message.ReplyBody{TextResult: termination.GetResult()},
			},
			Header: message.MessageHeader{
				Magic:       "MIOP",
				Version:     0,
				ByteOrder:   false,
				MessageType: 0,
				MessageSize: 0.0,
			},
		}
		// log.Printf("Sending { %s } to client", messageToBeMarshalled.Body.ReplyBody.TextResult)
		messageMarshalled, _ := json.Marshal(messageToBeMarshalled)
		// Send message back through server request handler
		srh.Send(messageMarshalled)
	}

	for {
		msgToBeUnmarshalled, _ = srh.Recieve()
		json.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)
		objectKey := msgUnmarshalled.Body.RequestHeader.ObjectKey
		if objectKey == cp.GetID() {
			switch msgUnmarshalled.Body.RequestHeader.Operation {
			case constants.UpperTextOperation:
				replyHandler(textHelper.UpperText(msgUnmarshalled.Body.RequestBody.Parameters[0]))
			case constants.LowerTextOperation:
				replyHandler(textHelper.LowerText(msgUnmarshalled.Body.RequestBody.Parameters[0]))
			case constants.InvertTextOperation:
				replyHandler(textHelper.InvertText(msgUnmarshalled.Body.RequestBody.Parameters[0]))
			}
		} else {
			log.Printf("Couldn't find remote object")
		}

	}

}
