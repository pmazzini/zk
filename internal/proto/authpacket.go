// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type AuthPacket struct {
	Type   int32  // type
	Scheme string // scheme
	Auth   []byte // auth
}

func (r *AuthPacket) GetType() int32 {
	if r != nil {
		return r.Type
	}
	return 0
}

func (r *AuthPacket) GetScheme() string {
	if r != nil {
		return r.Scheme
	}
	return ""
}

func (r *AuthPacket) GetAuth() []byte {
	if r != nil && r.Auth != nil {
		return r.Auth
	}
	return nil
}

func (r *AuthPacket) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Type, err = dec.ReadInt()
	if err != nil {
		return err
	}
	r.Scheme, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Auth, err = dec.ReadBuffer()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *AuthPacket) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Type); err != nil {
		return err
	}
	if err := enc.WriteString(r.Scheme); err != nil {
		return err
	}
	if err := enc.WriteBuffer(r.Auth); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *AuthPacket) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AuthPacket(%+v)", *r)
}
