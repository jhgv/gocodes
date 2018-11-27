package compress

import (
	"bytes"
	"fmt"
	"log"
	"testing"
)

func TestZipper_Compress(t *testing.T) {
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

	var buf bytes.Buffer
	z := new(Zipper)
	err := z.Compress(&buf, []byte(s))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("compressed size: %d \n\n", buf.Len())

	var buf2 bytes.Buffer
	err = z.Decompress(&buf2, buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("decompressed size: %d", buf2.Len())

	if buf2.String() != s {
		t.Error("Compression/Decompression failed")
	}
}