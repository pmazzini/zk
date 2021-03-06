// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type CheckWatchesRequest struct {
	Path string // path
	Type int32  // type
}

func (r *CheckWatchesRequest) GetPath() string {
	if r != nil {
		return r.Path
	}
	return ""
}

func (r *CheckWatchesRequest) GetType() int32 {
	if r != nil {
		return r.Type
	}
	return 0
}

func (r *CheckWatchesRequest) Read(dec jute.Decoder) (err error) {
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.Path, err = dec.ReadString()
	if err != nil {
		return err
	}
	r.Type, err = dec.ReadInt()
	if err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CheckWatchesRequest) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteString(r.Path); err != nil {
		return err
	}
	if err := enc.WriteInt(r.Type); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *CheckWatchesRequest) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CheckWatchesRequest(%+v)", *r)
}
