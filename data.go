/*

maryo/data.go

written by superwhiskers, licensed under gnu gplv3.
if you want a copy, go to http://www.gnu.org/licenses/

*/

package main

// struct for the isitworking endpoint
type isitworkingStruct struct {
	Server string `json:"server"`
}

// TODO: somewhere in here include a list containing a list of templates to add

// pretendo stock config
var pretendoConf = map[string]map[string]string{"config": map[string]string{"decryptOutgoing": "false"}, "endpoints": map[string]string{"account.nintendo.net": "account.pretendo.cc"}}

// local stock config
var localConf = map[string]map[string]string{"config": map[string]string{"decryptOutgoing": "true"}, "endpoints": map[string]string{"account.nintendo.net": "127.0.0.1:8080"}}

// test endpoints
var testEndpoints = map[string]map[string]string{"official": map[string]string{"account": "account.pretendo.cc"}, "local": map[string]string{"account": "127.0.0.1:8080"}, "ninty": map[string]string{"account": "account.nintendo.net"}}

// supposed return value for custom servers
var resMap = map[string]string{"account": "account.nintendo.net"}

// icons used for displaying results
var utilIcons = map[string]string{"success": "√", "failiure": "×", "uncertain": "-"}

// the nintendo CA cert
var nintyCert = []byte(`-----BEGIN CERTIFICATE-----
MIIEwzCCA6ugAwIBAgIBBjANBgkqhkiG9w0BAQsFADBtMQswCQYDVQQGEwJVUzET
MBEGA1UECBMKV2FzaGluZ3RvbjEhMB8GA1UEChMYTmludGVuZG8gb2YgQW1lcmlj
YSBJbmMuMQswCQYDVQQLEwJJUzEZMBcGA1UEAxMQTmludGVuZG8gQ0EgLSBHMzAe
Fw0xMDA1MTMxOTE5NDZaFw0zNzEyMjIxOTE5NDZaMIGlMQswCQYDVQQGEwJVUzET
MBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHUmVkbW9uZDEiMCAGA1UEChMZ
TmludGVuZG8gb2YgQW1lcmljYSwgSW5jLjELMAkGA1UECxMCSVMxGjAYBgNVBAMT
EUNUUiBDb21tb24gUHJvZCAxMSIwIAYJKoZIhvcNAQkBFhNjYUBub2EubmludGVu
ZG8uY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA81Vzs324jZwc
NpbFESgDNooVTRP1TlxvYwz8bbHnJHhImjEJNO29YSTpjmF7wonczooeKXfE/Ry2
+ey9mk92UhzSnvuSHQ6P2zFBbcPnE8eBi73oDnErgixiWe1TKP1G5LvwOqrEkVmX
LN/qnLrsfFp4QNyFc+PLvJ9IAfRSBwdRJHAiSgE9nB9eI7AGcM6DCw7+p9zEz6rN
RHUVRc5I132wJpQa8aoWaqPW7LE8exEC3VSfDHRVPjZUMRhfoBVSi2NfiA3xYsqk
v+Ct3E+bzW8y1aAQ7wIshQ/RGcLtVZE+tkoAznXewVLdKtcC67Vy4awhJ/BqK1tv
c26qV3zIJwIDAQABo4IBMzCCAS8wCQYDVR0TBAIwADAsBglghkgBhvhCAQ0EHxYd
T3BlblNTTCBHZW5lcmF0ZWQgQ2VydGlmaWNhdGUwHQYDVR0OBBYEFIzG7XO5Ojx2
G45r5dTszWF1rcFtMIGXBgNVHSMEgY8wgYyAFATT3tP98MjrwlmSh/sf1z5y+O35
oXGkbzBtMQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEhMB8GA1UE
ChMYTmludGVuZG8gb2YgQW1lcmljYSBJbmMuMQswCQYDVQQLEwJJUzEZMBcGA1UE
AxMQTmludGVuZG8gQ0EgLSBHM4IBATA7BgNVHR8ENDAyMDCgLqAshipodHRwOi8v
Y3JsLm5pbnRlbmRvLmNvbS9uaW50ZW5kby1jYS1nMy5jcmwwDQYJKoZIhvcNAQEL
BQADggEBAEOXZ/3IkNuFUfdxHpP0vrcSCTnDqMk8gsLVbN39BJT8Wqm8e3MFNhS/
Y1YOWgoIPtJp4cd2tXM3cXWzUZgm3SKd1XX/B81PFLEYlk+metUqB4jpF0ApCZs6
RNoXDBTx6XzsC07CA3uaxEdeWjC5Nl29AHuZ1YC/Z+7Da57TwBaa+/APj4y5mGUa
ahbvwpe1t3GSNOS5nBDSeCHAKLmzfnXpliA5qQZxo94RSXIVWK8hilXoFDQCL904
OGpgZnAhz4p3rcJYTq9ub8n6NYr9OJKKbWXfJY1QK4pXFVcIuAph0o/EyzDIEXuT
J4Q4b2km8uI0H4yxsQwUX9Epw6Vbujc=
-----END CERTIFICATE-----`)

// the nintendo CA key
var nintyKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDzVXOzfbiNnBw2
lsURKAM2ihVNE/VOXG9jDPxtseckeEiaMQk07b1hJOmOYXvCidzOih4pd8T9HLb5
7L2aT3ZSHNKe+5IdDo/bMUFtw+cTx4GLvegOcSuCLGJZ7VMo/Ubku/A6qsSRWZcs
3+qcuux8WnhA3IVz48u8n0gB9FIHB1EkcCJKAT2cH14jsAZwzoMLDv6n3MTPqs1E
dRVFzkjXfbAmlBrxqhZqo9bssTx7EQLdVJ8MdFU+NlQxGF+gFVKLY1+IDfFiyqS/
4K3cT5vNbzLVoBDvAiyFD9EZwu1VkT62SgDOdd7BUt0q1wLrtXLhrCEn8GorW29z
bqpXfMgnAgMBAAECggEBAMFOTib2JgmhTax0I8OYVM0b7wYXZ9XDit1WMKZ4INaR
E6QidlzszHiC2WO5v5Zw7M/LW2C3++7Tw+xRjOIsZCOhMBUKZy3cJp4LyB2J9mV5
JUm9KL9oWhcEaXFlHp4+bvZA8vu4M4YAdR86FuhBeqLjQArO5NmGypBivNKIpC1d
rwzSMyPddZvur7AsTIK0Ym9SwWN9eK7F2uBkzAneOugOTEAhq2ZnPMByNtjpTvw/
nvgNAB4Ukz/oomtleCaw92SoSlQlYGzVmuhvt4QqOQzS1V+ToauUhmAPmWEHF3IJ
yL1SiY7UmMlqQoGFV6IH4cwLiwm1wk/IF1tkfvR0gmECgYEA/vhTJYn/Gd3XY2tm
vJs5pvmkJJgNxqgunOCFZxrGxf4BLvQGkRnBlT4MzFa1J8ZNtU6VOVbQsS7cD0ke
2Kyv5rdt/7K2/bNdG4Ocnd+GcWpEtF+ik2wEQB48FQvYSQn/w5NKyJlyYOhXKGe3
XhTiyV5rCwiLfRP8pxAO0xLzVhcCgYEA9FEX+hlxYdg2bVfx+geI9uW0COLi7kzc
xxlPjBvdFltYsltSNLP+0BNqHfp3G5PW/XMTAdZnS9033gy4jLU6fAge8FcnLybh
GrpVNCnMY+DLRyVtlu1fcKKH59BK/ElR7aamxqJyw1JO2Z35u251nIbmrMgs5Txp
8aZ4HXCveHECgYEA1mRybddKhTKPwU53FdKkOK4jgo3Ez61tfIYiRl8ykxuRXSze
NLZmm5qQYmXqb+aEQxcvzQYd907CxaujX2hdhG/q854P1uYyPUd+sxVYVBeaa90a
tEGYlV2XAc9y73+T65z3vhOhJLFZUGVdv6NqSw60jZOCzwq2YLfU71E5AcMCgYEA
jQ6c90rlSYaZtfvGu4LKMzJgBZlpSAicl18nrE8SEKxgw2kyRzd88QmkhPZs+kEb
KW3dFXyCWyy36r4RdzvTLnVJ152aBAFAijv2oY1YcnoBI2yanz8hkVhlexOpl4uF
f95t/9UeyWKmH8KzwuF9igfg+vT/5sJAsMJaKzU6OiECgYEAlKK+RKTYwquD+3gK
o1gsGI+nR96Cb1kvfXzsj+V5UkZchew2pOqhrqpPknGIlFCeTDYjN8jqJyX4EljJ
1FTegfhfSe+XR7KOIh8b+d+fgftyRIp3M//BUF1FtwL789f/VakaIkz6Ret/+8tA
3UHqKKGtgSRL9kiTMVYy64pBG9A=
-----END RSA PRIVATE KEY-----`)
