package rpc

type tlsoption struct {
	certFile string
	keyFile string
	openssl bool
}

type changeTlsOption func(*tlsoption)

type ChangeTlsOptionImpl interface {
	apply(*tlsoption)
}

type changeTlsOptionFunc struct {
	f changeTlsOption
}
func(c *changeTlsOptionFunc) apply(cnf *tlsoption){
	c.f(cnf)
}
func newChangeTlsOptionFunc(f changeTlsOption) *changeTlsOptionFunc{
	return &changeTlsOptionFunc{f: f}
}

func AddChangeTlsOptionFunc(certFile,keyFile string) ChangeTlsOptionImpl {
	return newChangeTlsOptionFunc(func(t *tlsoption) {
		if certFile != "" && keyFile != "" {
			t.certFile = certFile
			t.keyFile  = keyFile
			t.openssl = true
		}
	})
}

