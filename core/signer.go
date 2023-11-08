package core

type Signer interface {
	Sign(msg []byte) ([]byte, error)
}
