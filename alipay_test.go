package alipay

import (
	"testing"
	"fmt"
)

var (
	appID     = "2016073100129537"
	partnerID = "2088102169227503"

	publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv8dXxi8wNAOqBNOh8Dv5
rh0BTb5KNgk62jDaS536Z1cDqq2JmpBYkBnzJXHAXEgBwPXgX8bGruMMjZKW8P4u
v3Rvj8Am9ewWwUK2U7m2ZB3Oo9MWtyYoiLGX1IA4FFenXzpPgm0WyzaeLX4yJ8j+
hVrRbgwbZzb9Aq0MyepnK5PVoSPLAPXxvWrIBTok1+liughxwD/7R+ldaQQCtWC7
nHBwOOChLkX6jenCOqi6LrTxJ4ycGTWTctngFMJO4YtMmq/2zrw+ovNqmxHJQAZw
uRFnKlZuFoEKPWyMGYtbvK9AWIfC8ubn30O5F9kfLMIHwAHCh0UipPSbKDwQ2BnW
swIDAQAB
-----END PUBLIC KEY-----`)

	privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAv8dXxi8wNAOqBNOh8Dv5rh0BTb5KNgk62jDaS536Z1cDqq2J
mpBYkBnzJXHAXEgBwPXgX8bGruMMjZKW8P4uv3Rvj8Am9ewWwUK2U7m2ZB3Oo9MW
tyYoiLGX1IA4FFenXzpPgm0WyzaeLX4yJ8j+hVrRbgwbZzb9Aq0MyepnK5PVoSPL
APXxvWrIBTok1+liughxwD/7R+ldaQQCtWC7nHBwOOChLkX6jenCOqi6LrTxJ4yc
GTWTctngFMJO4YtMmq/2zrw+ovNqmxHJQAZwuRFnKlZuFoEKPWyMGYtbvK9AWIfC
8ubn30O5F9kfLMIHwAHCh0UipPSbKDwQ2BnWswIDAQABAoIBAH7QyfkSsTRkC+Sf
MaGTd1qscXVAVQCAf/tSfLeuIqx9PL57fNfJhdbcYg2rt8EOGKLJtHKBFlcFawKf
IdMAslcGHtOXA+xxDucDP2AEGVkA4OkyJ/46bGlfzn/Fvc+t2s6812Du1DjSyCxb
G711SuFSGdVEikZpdUt0tVU7/LcyKAEZd45Ct+F9MvrPECbSsfODvTOVDHO2k42f
iwSzLPVmM4wVUc2xA15O87jtDhRiAK/RveQ7J2TWcarkyCR8J+bf5GGA79LdE3vR
Kr/HAk7INVX4T6U9QuDF30mqNRsloQbNGdvqO65nafNHvuVzUiqPdSX7XQwg/cOO
mhSsUbkCgYEA8BQXaHn3psHUZx8zEwQFVyd6rzxb+7jmVlUT+jG1pSiZ4WAWxxqx
YVXhn2dbfatDxWoGOMsrDM/Qp8g81nMG01jtmJr2RKFhAbQl93ipGvvaCNoJ8Lx7
HpFSq7dETcCCAE7tYMk0LlcVwxeaIUWakDyBHgEy4Zp6lLwdwsh115UCgYEAzH8/
E5dTOcYdcxk7HLupEC9MCb+FshZT5UIN9I7zLNljQX2O/8m2THb+oZUoy30RVot+
kYjh5H8M5CYiP0Kkm0Ovq5KC0loyt5SfzWbgwHEldQUVp8woE0YdaJzGB/UnmI9m
dJBON1t3qbMWjlguXOD8bfriDRuefaZd9oVSQycCgYBcz+ecxEoxdY2fsDgWid9m
qiSLylHlJr4lcg6fEsieaOvUbUlg/7jDYGgxL8v28Vbp4us02ZZzBYQs2QRsA1wI
KMDx1jaOobTW68YhvcviWqsX8PMW1kbislu7dsY5KMsZQ2oRmLdLku8e1OkJI9d1
G27vIpeBEC+DgJYgz05/YQKBgQCStWNiQbkihKBSF7LR3Uvf4Z6yi6V16xDLM8Vh
Q0DwVxEfRd3WYjcXynLJJ4J54kMTDMaD0GkHDaMI9taw/bWr8jZQZ67VDILAM68l
o/3v8fyGZFxx4kSJ905X48kqolWC3LYLQA/tJQDHTUUMX/T7CynuGQQdlUfyKu3U
Uzd+FwKBgHW9Nur4eTxK1nIOZyGgCqL1duYsJQcPWyIcRMTSjOoQZ5ZUhQZTw1Hd
2CW0Iu2fXExESTIjwXJ0ZJXnCgFU8acQX5vtItC1BlMaucw9XTx1RBCVQdTZ7DSX
vTlWbWwZHVDP85dioLE9mfo5+Hh3SmHDi3TaVXjxeJsUgHkRgOX7
-----END RSA PRIVATE KEY-----
`)
)

var partnerID = "2088102169227503"

func TestSign(t *testing.T) {

	var client = New(appID, partnerID, publicKey, privateKey, false)

	var p = AliPayTradeWapPay{}
	p.NotifyURL = "http://203.86.24.181:3000/alipay"
	p.ReturnURL = "http://203.86.24.181:3000"
	p.Subject = "修正了中文的 Bug"
	p.OutTradeNo = "trade_no_1234"
	p.TotalAmount = "10.00"
	p.ProductCode = "eeeeee"

	var url, _ = client.TradeWapPay(p)
	fmt.Println(url)
}

func TestAliPayTradeQuery(t *testing.T) {
	client := New(appID, partnerID, publicKey, privateKey, false)

	type arg struct {
		outTradeNo string
		wanted     error
		name       string
	}

	testCaes := []arg{
		{"1111111", nil, "query success"},
		//TODO:add more test case
	}

	for _, tc := range testCaes {
		req := AliPayTradeQuery{
			OutTradeNo: tc.outTradeNo,
		}
		resp, err := client.TradeQuery(req)
		if err != tc.wanted {
			t.Errorf("%s input:%s wanted:%v get:%v", tc.name, tc.outTradeNo, tc.wanted, err)
		} else {
			t.Log(resp)
		}
	}

}
