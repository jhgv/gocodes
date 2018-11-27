package requestor

import (
	"bytes"
	"encoding/json"
	"github.com/jhgv/gocodes/middleware/core/compress"
	"github.com/jhgv/gocodes/middleware/core/security"
	"reflect"
	"testing"
)

type Message struct {
	D []byte
}

func TestRequestor_Invoke(t *testing.T) {

	s := `1 Obadiah	Hatt	ohatt0@jimdo.com	Male	165.44.236.156
	2	Rosemaria	Cardinal	rcardinal1@cdc.gov	Female	125.158.20.77
	3	Johanna	Eggleson	jeggleson2@t-online.de	Female	103.15.198.95
	4	Ettie	Piggott	epiggott3@washingtonpost.com	Female	91.210.151.239
	5	Rafaello	Mart	rmart4@sitemeter.com	Male	2.52.149.58
	6	Caspar	Heading	cheading5@hubpages.com	Male	168.244.82.167
	7	Vin	Mahedy	vmahedy6@seattletimes.com	Male	25.104.66.54
	8	Rosmunda	Rosengren	rrosengren7@opensource.org	Female	113.190.164.167
	9	Franni	Gooderidge	fgooderidge8@columbia.edu	Female	16.43.130.184
	10	Clea	Grunguer	cgrunguer9@miibeian.gov.cn	Female	163.215.197.222
	11	Cass	McIver	cmcivera@java.com	Male	89.103.34.155
	12	Pippa	Reading	preadingb@fastcompany.com	Female	73.18.173.87
	13	Paige	Perchard	pperchardc@marketwatch.com	Male	112.196.222.40
	14	Edythe	Soars	esoarsd@networksolutions.com	Female	189.222.49.80
	15	Pernell	Musgrove	pmusgrovee@tripadvisor.com	Male	29.188.228.63
	16	Carla	Singyard	csingyardf@reference.com	Female	147.180.255.209
	17	Grover	Warboy	gwarboyg@fotki.com	Male	23.65.175.130
	18	Rebecca	Plenderleith	rplenderleithh@clickbank.net	Female	195.38.232.240
	19	Gerrilee	Overell	goverelli@51.la	Female	170.24.33.57
	20	Guillema	Kemwall	gkemwallj@cornell.edu	Female	240.215.12.195`

	var encrypter = &security.Encrypter{}
	var key = []byte("1234567890123456")

	var buff bytes.Buffer
	var buff2 bytes.Buffer
	var messageUnmarshalled Message
	messageToBeMarshalled := &Message{D: []byte(s)}
	z := new(compress.Zipper)

	var messageToBeUnmarshalled []byte

	messageToBeUnmarshalled, _ = json.Marshal(messageToBeMarshalled)

	z.Compress(&buff, messageToBeUnmarshalled)

	enc := encrypter.Encrypt(key, buff.Bytes())

	dec := encrypter.Decrypt(key, enc)

	z.Decompress(&buff2, dec)

	json.Unmarshal(buff2.Bytes(), &messageUnmarshalled)

	if !reflect.DeepEqual(messageUnmarshalled.D, messageToBeMarshalled.D) {
		t.Error("Fail!")
	}

}
