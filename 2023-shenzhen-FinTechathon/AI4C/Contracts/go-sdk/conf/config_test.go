package conf

import (
	"encoding/json"
	"testing"
)

const (
	standardJSON = `[{"IsHTTP":false,"ChainID":1,"CAFile":"ca.crt","TLSCAContext":"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCglNSUlCWkRDQ0FRb0NGQjkvUzViTDZySkY2cUJpSEdmSU9wejJJZDVYTUFvR0NDcUdTTTQ5QkFNQ01EVXhEakFNCglCZ05WQkFNTUJXTm9ZV2x1TVJNd0VRWURWUVFLREFwbWFYTmpieTFpWTI5ek1RNHdEQVlEVlFRTERBVmphR0ZwCgliakFnRncweU1qQTBNVFV3TnpNM05ESmFHQTh5TVRJeU1ETXlNakEzTXpjME1sb3dOVEVPTUF3R0ExVUVBd3dGCglZMmhoYVc0eEV6QVJCZ05WQkFvTUNtWnBjMk52TFdKamIzTXhEakFNQmdOVkJBc01CV05vWVdsdU1GWXdFQVlICglLb1pJemowQ0FRWUZLNEVFQUFvRFFnQUVmdWZLdUJpdE1uMmdpVW5qUVpEUldaRkwwSDdxdUg5d25IVkY2cG9NCglDUEFCbDVRV3lpTEpsM0FTZ3N2YktIMDFZSmY3VG5OZ3FXVmMybVltVXQyNjVEQUtCZ2dxaGtqT1BRUURBZ05JCglBREJGQWlFQTlxdWl0aDg0ZFdyTWtkMlhJVUluVVZRMzZtc1h4bDZjeUVVRkw0c1JrNjRDSUN4MnRWK3JNMHFkCglsNEdEUG1RSmxsQWtrVm55cmU4TFFBYms5dnpuYkZrZAoJLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQoJ","Key":"sdk.key","TLSKeyContext":"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCglNSUdFQWdFQU1CQUdCeXFHU000OUFnRUdCU3VCQkFBS0JHMHdhd0lCQVFRZ0FRcUJheWF1TUlIRDM2NEVyMlZuCglvblFGUGFXdG44T1p5dThxMjZGRWNHV2hSQU5DQUFTMDVuaE0vdzU2YTg2SVg5TVl4SWZTd2lQR1BoZUdEeXQ3CgluelFNMFZNWWd2dHZ2RUpPQzduNC91RDYreWlUSGI2R1JFb0R5aDRRekdVTEdRaEE0NnJmCgktLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCgk=","Cert":"sdk.crt","TLSCertContext":"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCglNSUlCZ2pDQ0FTbWdBd0lCQWdJVVp2ZGVlMERISjVGRXA1VG8vTkprbmorNTRpWXdDZ1lJS29aSXpqMEVBd0l3CglOekVQTUEwR0ExVUVBd3dHWVdkbGJtTjVNUk13RVFZRFZRUUtEQXBtYVhOamJ5MWlZMjl6TVE4d0RRWURWUVFMCglEQVpoWjJWdVkza3dJQmNOTWpJd05ERTFNRGN6TnpReVdoZ1BNakV5TWpBek1qSXdOek0zTkRKYU1ERXhEREFLCglCZ05WQkFNTUEzTmthekVUTUJFR0ExVUVDZ3dLWm1selkyOHRZbU52Y3pFTU1Bb0dBMVVFQ3d3RGMyUnJNRll3CglFQVlIS29aSXpqMENBUVlGSzRFRUFBb0RRZ0FFdE9aNFRQOE9lbXZPaUYvVEdNU0gwc0lqeGo0WGhnOHJlNTgwCglETkZUR0lMN2I3eENUZ3U1K1A3Zyt2c29reDIraGtSS0E4b2VFTXhsQ3hrSVFPT3EzNk1hTUJnd0NRWURWUjBUCglCQUl3QURBTEJnTlZIUThFQkFNQ0JlQXdDZ1lJS29aSXpqMEVBd0lEUndBd1JBSWdBNm42UHNJdFJPWkxmYzF6CglCNS9uWVIraXlyVWdycTg4YzdHVnhZVmliUmdDSUg2QStuT1Fna1Q3Z0dSQlZ0WVVoaFpObEhNem51ZnNhck9PCglzMUhHVURESgoJLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQoJLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCglNSUlCZXpDQ0FTR2dBd0lCQWdJVWR4UGNWTDJ2T05HV25HcFRPZk9td1ZwVUx2Z3dDZ1lJS29aSXpqMEVBd0l3CglOVEVPTUF3R0ExVUVBd3dGWTJoaGFXNHhFekFSQmdOVkJBb01DbVpwYzJOdkxXSmpiM014RGpBTUJnTlZCQXNNCglCV05vWVdsdU1CNFhEVEl5TURReE5UQTNNemMwTWxvWERUTXlNRFF4TWpBM016YzBNbG93TnpFUE1BMEdBMVVFCglBd3dHWVdkbGJtTjVNUk13RVFZRFZRUUtEQXBtYVhOamJ5MWlZMjl6TVE4d0RRWURWUVFMREFaaFoyVnVZM2t3CglWakFRQmdjcWhrak9QUUlCQmdVcmdRUUFDZ05DQUFRU0FReEhHWm14c2tyL2VDbUNyMmtGL1Y0VkZhaG1GZWpqCgk2bURuK3BHYVE1MHFGT3NienRnamtzdkFCOUNDeGplZzFyNUpVV2NKZ25vREhQV1pwSHFMb3hBd0RqQU1CZ05WCglIUk1FQlRBREFRSC9NQW9HQ0NxR1NNNDlCQU1DQTBnQU1FVUNJUUNNOUZhZ0ViMTU3eGhrU3IvVzIvYVM4b1M3CglicnN6L3NsRzBndFZGTW1SYUFJZ0dhWHMwQk56Wm1ZUDZibm5TUGVzbkJCUWtlS25yL1pSY0F4Z2JVTDdOR1k9CgktLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCgktLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS0KCU1JSUJaRENDQVFvQ0ZCOS9TNWJMNnJKRjZxQmlIR2ZJT3B6MklkNVhNQW9HQ0NxR1NNNDlCQU1DTURVeERqQU0KCUJnTlZCQU1NQldOb1lXbHVNUk13RVFZRFZRUUtEQXBtYVhOamJ5MWlZMjl6TVE0d0RBWURWUVFMREFWamFHRnAKCWJqQWdGdzB5TWpBME1UVXdOek0zTkRKYUdBOHlNVEl5TURNeU1qQTNNemMwTWxvd05URU9NQXdHQTFVRUF3d0YKCVkyaGhhVzR4RXpBUkJnTlZCQW9NQ21acGMyTnZMV0pqYjNNeERqQU1CZ05WQkFzTUJXTm9ZV2x1TUZZd0VBWUgKCUtvWkl6ajBDQVFZRks0RUVBQW9EUWdBRWZ1Zkt1Qml0TW4yZ2lVbmpRWkRSV1pGTDBIN3F1SDl3bkhWRjZwb00KCUNQQUJsNVFXeWlMSmwzQVNnc3ZiS0gwMVlKZjdUbk5ncVdWYzJtWW1VdDI2NURBS0JnZ3Foa2pPUFFRREFnTkkKCUFEQkZBaUVBOXF1aXRoODRkV3JNa2QyWElVSW5VVlEzNm1zWHhsNmN5RVVGTDRzUms2NENJQ3gydFYrck0wcWQKCWw0R0RQbVFKbGxBa2tWbnlyZThMUUFiazl2em5iRmtkCgktLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCgk=","IsSMCrypto":false,"DynamicKey":false,"PrivateKey":"uJ1C8SKQBw8jX7j7Ydz5bjsRUWxdT2Mz8m5Ju5Vfi2I=","GroupID":1,"NodeURL":"127.0.0.1:20200"}]`
	fileContent  = `
	[Network]
	#type rpc or channel
	Type="channel"
	CAFile="ca.crt"
	Cert="sdk.crt"
	Key="sdk.key"
	# if the certificate context is not empty, use it, otherwise read from the certificate file
	CAContext='''-----BEGIN CERTIFICATE-----
	MIIBZDCCAQoCFB9/S5bL6rJF6qBiHGfIOpz2Id5XMAoGCCqGSM49BAMCMDUxDjAM
	BgNVBAMMBWNoYWluMRMwEQYDVQQKDApmaXNjby1iY29zMQ4wDAYDVQQLDAVjaGFp
	bjAgFw0yMjA0MTUwNzM3NDJaGA8yMTIyMDMyMjA3Mzc0MlowNTEOMAwGA1UEAwwF
	Y2hhaW4xEzARBgNVBAoMCmZpc2NvLWJjb3MxDjAMBgNVBAsMBWNoYWluMFYwEAYH
	KoZIzj0CAQYFK4EEAAoDQgAEfufKuBitMn2giUnjQZDRWZFL0H7quH9wnHVF6poM
	CPABl5QWyiLJl3ASgsvbKH01YJf7TnNgqWVc2mYmUt265DAKBggqhkjOPQQDAgNI
	ADBFAiEA9quith84dWrMkd2XIUInUVQ36msXxl6cyEUFL4sRk64CICx2tV+rM0qd
	l4GDPmQJllAkkVnyre8LQAbk9vznbFkd
	-----END CERTIFICATE-----
	'''
	KeyContext='''-----BEGIN PRIVATE KEY-----
	MIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQgAQqBayauMIHD364Er2Vn
	onQFPaWtn8OZyu8q26FEcGWhRANCAAS05nhM/w56a86IX9MYxIfSwiPGPheGDyt7
	nzQM0VMYgvtvvEJOC7n4/uD6+yiTHb6GREoDyh4QzGULGQhA46rf
	-----END PRIVATE KEY-----
	'''
	CertContext='''-----BEGIN CERTIFICATE-----
	MIIBgjCCASmgAwIBAgIUZvdee0DHJ5FEp5To/NJknj+54iYwCgYIKoZIzj0EAwIw
	NzEPMA0GA1UEAwwGYWdlbmN5MRMwEQYDVQQKDApmaXNjby1iY29zMQ8wDQYDVQQL
	DAZhZ2VuY3kwIBcNMjIwNDE1MDczNzQyWhgPMjEyMjAzMjIwNzM3NDJaMDExDDAK
	BgNVBAMMA3NkazETMBEGA1UECgwKZmlzY28tYmNvczEMMAoGA1UECwwDc2RrMFYw
	EAYHKoZIzj0CAQYFK4EEAAoDQgAEtOZ4TP8OemvOiF/TGMSH0sIjxj4Xhg8re580
	DNFTGIL7b7xCTgu5+P7g+vsokx2+hkRKA8oeEMxlCxkIQOOq36MaMBgwCQYDVR0T
	BAIwADALBgNVHQ8EBAMCBeAwCgYIKoZIzj0EAwIDRwAwRAIgA6n6PsItROZLfc1z
	B5/nYR+iyrUgrq88c7GVxYVibRgCIH6A+nOQgkT7gGRBVtYUhhZNlHMznufsarOO
	s1HGUDDJ
	-----END CERTIFICATE-----
	-----BEGIN CERTIFICATE-----
	MIIBezCCASGgAwIBAgIUdxPcVL2vONGWnGpTOfOmwVpULvgwCgYIKoZIzj0EAwIw
	NTEOMAwGA1UEAwwFY2hhaW4xEzARBgNVBAoMCmZpc2NvLWJjb3MxDjAMBgNVBAsM
	BWNoYWluMB4XDTIyMDQxNTA3Mzc0MloXDTMyMDQxMjA3Mzc0MlowNzEPMA0GA1UE
	AwwGYWdlbmN5MRMwEQYDVQQKDApmaXNjby1iY29zMQ8wDQYDVQQLDAZhZ2VuY3kw
	VjAQBgcqhkjOPQIBBgUrgQQACgNCAAQSAQxHGZmxskr/eCmCr2kF/V4VFahmFejj
	6mDn+pGaQ50qFOsbztgjksvAB9CCxjeg1r5JUWcJgnoDHPWZpHqLoxAwDjAMBgNV
	HRMEBTADAQH/MAoGCCqGSM49BAMCA0gAMEUCIQCM9FagEb157xhkSr/W2/aS8oS7
	brsz/slG0gtVFMmRaAIgGaXs0BNzZmYP6bnnSPesnBBQkeKnr/ZRcAxgbUL7NGY=
	-----END CERTIFICATE-----
	-----BEGIN CERTIFICATE-----
	MIIBZDCCAQoCFB9/S5bL6rJF6qBiHGfIOpz2Id5XMAoGCCqGSM49BAMCMDUxDjAM
	BgNVBAMMBWNoYWluMRMwEQYDVQQKDApmaXNjby1iY29zMQ4wDAYDVQQLDAVjaGFp
	bjAgFw0yMjA0MTUwNzM3NDJaGA8yMTIyMDMyMjA3Mzc0MlowNTEOMAwGA1UEAwwF
	Y2hhaW4xEzARBgNVBAoMCmZpc2NvLWJjb3MxDjAMBgNVBAsMBWNoYWluMFYwEAYH
	KoZIzj0CAQYFK4EEAAoDQgAEfufKuBitMn2giUnjQZDRWZFL0H7quH9wnHVF6poM
	CPABl5QWyiLJl3ASgsvbKH01YJf7TnNgqWVc2mYmUt265DAKBggqhkjOPQQDAgNI
	ADBFAiEA9quith84dWrMkd2XIUInUVQ36msXxl6cyEUFL4sRk64CICx2tV+rM0qd
	l4GDPmQJllAkkVnyre8LQAbk9vznbFkd
	-----END CERTIFICATE-----
	'''

	[[Network.Connection]]
	NodeURL="127.0.0.1:20200"
	GroupID=1
	# [[Network.Connection]]
	# NodeURL="127.0.0.1:20200"
	# GroupID=2

	[Account]
	# only support PEM format for now
	KeyFile="../.ci/0x83309d045a19c44dc3722d15a6abd472f95866ac.pem"

	[Chain]
	ChainID=1
	SMCrypto=false

	[log]
	Path="./"
	`
)

func TestConfig(t *testing.T) {
	// test parseConfig
	configs, err := ParseConfig([]byte(fileContent))
	if err != nil {
		t.Fatalf("TestConfig failed, err: %v", err)
	}
	jsons, err := json.Marshal(configs)
	if err != nil {
		t.Fatalf("failed when struct transfers to json, err: %v", err)
	}
	if standardJSON != string(jsons) {
		t.Fatalf("parsing the output of test.toml is inconsistent with the standardJson\n the output of test.toml: %v\n standardJson: %v", string(jsons), standardJSON)
	}
	t.Logf("the output of test.toml: %v", string(jsons))
}
