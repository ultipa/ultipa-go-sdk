package utils

import (
	"crypto/tls"
	"crypto/x509"
)

func GetCertificate(host string) *x509.Certificate {

	// Create a connection to obtain the certificate of the remote server
	conn, err := tls.Dial("tcp", host, nil)
	if err != nil {
		return nil
	}
	defer conn.Close()

	// Gets the certificate chain returned by the server
	certs := conn.ConnectionState().PeerCertificates
	if len(certs) == 0 {
		return nil
	}

	// Returns the first certificate
	return certs[0]
}
