package fasthttp

type MutualTLS struct {
	certFile []byte
	keyFile  []byte
}

func CertFile(certFile []byte) *MutualTLS {
	mtls := &MutualTLS{
		certFile: certFile,
	}
	return mtls
}

func (mtls *MutualTLS) CertFile(certFile []byte) *MutualTLS {
	mtls.certFile = certFile
	return mtls
}

func KeyFile(keyFile []byte) *MutualTLS {
	mtls := &MutualTLS{
		keyFile: keyFile,
	}
	return mtls
}

func (mtls *MutualTLS) KeyFile(keyFile []byte) *MutualTLS {
	mtls.keyFile = keyFile
	return mtls
}
