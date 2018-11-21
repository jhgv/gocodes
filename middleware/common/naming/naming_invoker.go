package naming

import (
	"encoding/json"
	"github.com/jhgv/gocodes/middleware/core/message"
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/server/handler"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"log"
)

type namingInvoker struct {
	naming *Naming
}

func NewNamingInvoker() *namingInvoker {
	n := new(namingInvoker)
	n.naming = NewNaming()
	return n
}

func (ni *namingInvoker) Invoke(port int) {
	srh := new(handler.TCPServerRequestHanlder)
	srh.SetupSocket("", port)

	var msgToBeUnmarshalled []byte
	var msgUnmarshalled = message.Message{}

	for {
		log.Println("Naming server wating for biding o lookup operations...")
		msgToBeUnmarshalled, _ = srh.Recieve()
		json.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)

		switch msgUnmarshalled.Body.RequestHeader.Operation {
		case constants.BindOperation:
			serviceName := msgUnmarshalled.Body.RequestBody.Params[0].(string)
			clientProxyParams := msgUnmarshalled.Body.RequestBody.Params[1].(map[string]interface{})
			clientProxy := proxy.FileConverterProxy{}
			clientProxy.SetHost(clientProxyParams["Host"].(string))
			clientProxy.SetPort(int(clientProxyParams["Port"].(float64)))
			clientProxy.SetID(int(clientProxyParams["ID"].(float64)))
			ni.naming.Bind(serviceName, clientProxy)
			// Creating message
			messageToBeMarshalled := message.Message{
				Body: message.MessageBody{
					ReplyHeader: message.ReplyHeader{ReplyStatus: 0, RequestID: 0, ServiceContext: ""},
					ReplyBody:   message.ReplyBody{ServiceName: serviceName},
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

		case constants.LookupOperation:
			serviceName := msgUnmarshalled.Body.RequestBody.Params[0].(string)

			clientProxy, err := ni.naming.Lookup(serviceName)
			if err != nil {
				log.Fatal(err)
			}

			// Creating message
			messageToBeMarshalled := message.Message{
				Body: message.MessageBody{
					ReplyHeader: message.ReplyHeader{ReplyStatus: 0, RequestID: 0, ServiceContext: ""},
					ReplyBody:   message.ReplyBody{ClientProxy: clientProxy},
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
	}

}
