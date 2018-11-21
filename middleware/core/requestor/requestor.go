package requestor

import (
	"encoding/json"
	"github.com/jhgv/gocodes/middleware/client/handler"
	"github.com/jhgv/gocodes/middleware/core/message"
	"github.com/jhgv/gocodes/middleware/core/security"
)

type Requestor struct {
}

func (r *Requestor) Invoke(inv *Invocation) (*Termination, error) {

	var termination = &Termination{}
	var encrypter = &security.Encrypter{}

	// Mouting message protocol and marshalling message
	var messageToBeUnmarshalled []byte
	messageToBeMarshalled := &message.Message{}
	// Setting message header
	messageToBeMarshalled.Header = message.MessageHeader{Magic: "MIOP", Version: 0, ByteOrder: false, MessageType: 0, MessageSize: 0.0}

	// TODO: Organize message object construction
	// Setting message body
	requestHeader := message.RequestHeader{Context: "", ObjectKey: inv.GetObjectID(), ResponseExpected: false, RequestID: 0, Operation: inv.GetOperationName()}
	requestBody := message.RequestBody{File: inv.GetFile()}
	messageToBeMarshalled.Body = message.MessageBody{RequestHeader: requestHeader, RequestBody: requestBody}
	messageMarshalled, _ := json.Marshal(messageToBeMarshalled)

	// Communication process through request handler
	client := new(handler.TCPClientHandler)
	client.SetupSocket(inv.GetHost(), inv.GetPort())
	client.Send(encrypter.Encrypt([]byte("1234567890123456"), messageMarshalled))
	messageToBeUnmarshalled, _ = client.Recieve()
	messageToBeUnmarshalled  = encrypter.Decrypt([]byte("1234567890123456"), messageToBeUnmarshalled)
	// Unmarshalling message
	var messageUnmarshalled message.Message
	json.Unmarshal(messageToBeUnmarshalled, &messageUnmarshalled)
	termination.SetResult(messageUnmarshalled.Body.ReplyBody.ConvertedFile)
	return termination, nil
}

func (r *Requestor) InvokeN(inv *Invocation) (*Termination, error) {

	var termination = &Termination{}

	// Mouting message protocol and marshalling message
	var messageToBeUnmarshalled []byte
	messageToBeMarshalled := message.Message{}
	// Setting message header
	messageToBeMarshalled.Header = message.MessageHeader{Magic: "MIOP", Version: 0, ByteOrder: false, MessageType: 0, MessageSize: 0.0}

	// TODO: Organize message object construction
	// Setting message body
	requestHeader := message.RequestHeader{Context: "", ObjectKey: inv.GetObjectID(), ResponseExpected: false, RequestID: 0, Operation: inv.GetOperationName()}
	requestBody := message.RequestBody{Params: inv.GetParams()}
	messageToBeMarshalled.Body = message.MessageBody{RequestHeader: requestHeader, RequestBody: requestBody}
	messageMarshalled, _ := json.Marshal(messageToBeMarshalled)

	// Communication process through request handler
	client := new(handler.TCPClientHandler)
	client.SetupSocket(inv.GetHost(), inv.GetPort())
	client.Send(messageMarshalled)
	messageToBeUnmarshalled, _ = client.Recieve()
	// Unmarshalling message
	var messageUnmarshalled message.Message
	json.Unmarshal(messageToBeUnmarshalled, &messageUnmarshalled)
	termination.SetResult(messageUnmarshalled.Body.ReplyBody.ClientProxy)
	return termination, nil
}
