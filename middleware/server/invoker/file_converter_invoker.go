package invoker

import (
	"bytes"
	"encoding/json"
	"github.com/jhgv/gocodes/middleware/core/compress"
	"github.com/jhgv/gocodes/middleware/core/security"
	"github.com/jhgv/gocodes/middleware/utils/constants"
	"io/ioutil"
	"log"
	"os"

	"github.com/jhgv/gocodes/middleware/core/message"
	"github.com/jhgv/gocodes/middleware/core/proxy"
	"github.com/jhgv/gocodes/middleware/core/requestor"
	"github.com/jhgv/gocodes/middleware/server/handler"
	"github.com/jhgv/gocodes/middleware/server/objects"
)

type FileConverterInvoker struct {
}

func (fci *FileConverterInvoker) Invoke(cp proxy.FileConverterProxy) {
	srh := new(handler.TCPServerRequestHanlder)
	srh.SetupSocket(cp.GetHost(), cp.GetPort())

	var encrypter = security.Encrypter{}
	var z = compress.Zipper{}

	var msgToBeUnmarshalled []byte
	var msgUnmarshalled = message.Message{}

	var termination requestor.Termination
	var fileConverter = &objects.FileConverter{}

	var replyHandler = func(convertedFile os.File) {
		//fileInfo, _ := convertedFile.Stat()
		// make file bytes to send back
		//fileInfo, err := convertedFile.Stat()
		//if err != nil {
		//	log.Fatal(err)
		//}

		b, _ := ioutil.ReadFile("file.csv")
		//convertedFile.Read(b)

		// Unmarshalling result
		termination.SetResult(b)

		// Creating message
		messageToBeMarshalled := &message.Message{
			Body: message.MessageBody{
				ReplyHeader: message.ReplyHeader{ReplyStatus: 0, RequestID: 0, ServiceContext: ""},
				ReplyBody:   message.ReplyBody{ConvertedFile: termination.GetResult().([]byte)},
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
		var buff bytes.Buffer
		z.Compress(&buff, messageMarshalled)
		compressedMessageToBeMarsahlled := buff.Bytes()
		messageMarshalled = encrypter.Encrypt([]byte("1234567890123456"), compressedMessageToBeMarsahlled)
		// Send message back through server request handler
		srh.Send(messageMarshalled)
		convertedFile.Close()
	}

	for {
		encryptedMsgToBeUnmarshalled, _ := srh.Recieve()
		// Decrypt
		compressedMessageToBeUnmarshalled := encrypter.Decrypt([]byte("1234567890123456"), encryptedMsgToBeUnmarshalled)
		// Decompress
		var b bytes.Buffer
		z.Decompress(&b, compressedMessageToBeUnmarshalled)
		msgToBeUnmarshalled = b.Bytes()

		json.Unmarshal(msgToBeUnmarshalled, &msgUnmarshalled)
		objectKey := msgUnmarshalled.Body.RequestHeader.ObjectKey
		file := msgUnmarshalled.Body.RequestBody.File

		if objectKey == cp.GetID() {
			switch msgUnmarshalled.Body.RequestHeader.Operation {
			case constants.ConvertFileOperation:
				replyHandler(fileConverter.ConvertFile(file))
			}
		} else {
			log.Printf("Couldn't find remote object")
		}

	}

}
