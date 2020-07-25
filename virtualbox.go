package virtualbox

import(
	"net/http"
	"bytes"
	"io"
	"fmt"
	"strconv"
	"encoding/json"
)

var daemon = "daemonhost to be written"
var d = "http://" + daemon

type Wallet struct {
	Mnemonic string `json:"mnemonic"`
	PrivateKey string `json:"privateKey"`
	PublicKey string `json:"publicKey"`
}

type NewTx struct {
	PrivateKey string
	Amount float64
	Receiver string
}
type Transaction struct {
	Txid int `json:"txid"`
	Amount float64 `json:"amount"`
	Sender string `json:"sender"`
	Receiver string `json:"receiver"`
	Timestamp float64 `json:"timestamp"`
	Txtype string `json:"txtype"`
}

func bufToString(buffer io.ReadCloser) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(buffer)
	return buf.String()
}

func CreateWallet() Wallet {
	res,err := http.Get(d + "/createwallet")
	if err != nil {
		panic(err)
	}
	var resstring Wallet
	body := bufToString(res.Body)
	err2 := json.Unmarshal([]byte(body),&resstring)
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()
	return resstring
}

func GetTxById(id int) Transaction {
	id2 :=  strconv.Itoa(id)
	res,err := http.Get(d + "/gettx/" + id2)
	body := bufToString(res.Body)
	if err != nil {
		panic(err)
	}
	var resstring Transaction
	err2 := json.Unmarshal([]byte(body),&resstring)
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()
	return resstring
}

func Balance(publicKey string) float64 {
	res,err := http.Get(d + "/balance/" + publicKey)
	if err != nil {
		panic(err)
	}
	var balance float64
	body := bufToString(res.Body)
	err2 := json.Unmarshal([]byte(body),&balance)
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()
	return balance
}

func ReceivedTx(publicKey string) []Transaction {
	res,err := http.Get(d + "/receivedtx/" + publicKey)
	if err != nil {
		panic(err)
	}
	var txs []Transaction
	body := []byte(bufToString(res.Body))
	err2 := json.Unmarshal(body,&txs)
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()
	return txs
}

func SentTx(publicKey string) []Transaction {
	res,err := http.Get(d + "/senttx/" + publicKey)
	if err != nil {
		panic(err)
	}
	var txs []Transaction
	body := []byte(bufToString(res.Body))
	err2 := json.Unmarshal(body,&txs)
	if err2 != nil {
		panic(err2)
	}
	defer res.Body.Close()
	return txs
}

func SendTx(privateKey string,amount float64, receiver string) Transaction {
	thistx := NewTx{PrivateKey : privateKey,Amount : amount,Receiver : receiver}
	jsontx,err0 := json.Marshal(thistx)
	if err0 != nil {
		panic(err0)
	}
	res,err := http.Post(d + "/sendtx","application/json",bytes.NewBuffer(jsontx))
	if err != nil {
		panic(err)
	}
	var tx Transaction
	body := []byte(bufToString(res.Body))
	err2 := json.Unmarshal(body,&tx)
	if err2 != nil {
		panic(err)
	}
	defer res.Body.Close()
	return tx
}
