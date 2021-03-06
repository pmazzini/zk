// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type GetSASLRequest struct {
	Token []byte // token
}

func (r *GetSASLRequest) GetToken() []byte {
	if r != nil && r.Token != nil {
		return r.Token
	}
	return nil
}

func (r *GetSASLRequest) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Token, err = dec.ReadBuffer()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetSASLRequest) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteBuffer(r.Token); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *GetSASLRequest) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetSASLRequest(%+v)", *r)
}
