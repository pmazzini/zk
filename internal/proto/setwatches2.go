// Autogenerated jute compiler
// @generated from 'zookeeper.jute'

package proto // github.com/facebookincubator/zk/internal/proto

import (
	"fmt"

	"github.com/go-zookeeper/jute/lib/go/jute"
)

type SetWatches2 struct {
	RelativeZxid               int64    // relativeZxid
	DataWatches                []string // dataWatches
	ExistWatches               []string // existWatches
	ChildWatches               []string // childWatches
	PersistentWatches          []string // persistentWatches
	PersistentRecursiveWatches []string // persistentRecursiveWatches
}

func (r *SetWatches2) GetRelativeZxid() int64 {
	if r != nil {
		return r.RelativeZxid
	}
	return 0
}

func (r *SetWatches2) GetDataWatches() []string {
	if r != nil && r.DataWatches != nil {
		return r.DataWatches
	}
	return nil
}

func (r *SetWatches2) GetExistWatches() []string {
	if r != nil && r.ExistWatches != nil {
		return r.ExistWatches
	}
	return nil
}

func (r *SetWatches2) GetChildWatches() []string {
	if r != nil && r.ChildWatches != nil {
		return r.ChildWatches
	}
	return nil
}

func (r *SetWatches2) GetPersistentWatches() []string {
	if r != nil && r.PersistentWatches != nil {
		return r.PersistentWatches
	}
	return nil
}

func (r *SetWatches2) GetPersistentRecursiveWatches() []string {
	if r != nil && r.PersistentRecursiveWatches != nil {
		return r.PersistentRecursiveWatches
	}
	return nil
}

func (r *SetWatches2) Read(dec jute.Decoder) (err error) {
	var size int
	if err = dec.ReadStart(); err != nil {
		return err
	}
	r.RelativeZxid, err = dec.ReadLong()
	if err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.DataWatches = nil
	} else {
		r.DataWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.DataWatches[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.ExistWatches = nil
	} else {
		r.ExistWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.ExistWatches[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.ChildWatches = nil
	} else {
		r.ChildWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.ChildWatches[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.PersistentWatches = nil
	} else {
		r.PersistentWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.PersistentWatches[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	size, err = dec.ReadVectorStart()
	if err != nil {
		return err
	}
	if size < 0 {
		r.PersistentRecursiveWatches = nil
	} else {
		r.PersistentRecursiveWatches = make([]string, size)
		for i := 0; i < size; i++ {
			r.PersistentRecursiveWatches[i], err = dec.ReadString()
			if err != nil {
				return err
			}
		}
	}
	if err = dec.ReadVectorEnd(); err != nil {
		return err
	}
	if err = dec.ReadEnd(); err != nil {
		return err
	}
	return nil
}

func (r *SetWatches2) Write(enc jute.Encoder) error {
	if err := enc.WriteStart(); err != nil {
		return err
	}
	if err := enc.WriteLong(r.RelativeZxid); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.DataWatches), r.DataWatches == nil); err != nil {
		return err
	}
	for _, v := range r.DataWatches {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.ExistWatches), r.ExistWatches == nil); err != nil {
		return err
	}
	for _, v := range r.ExistWatches {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.ChildWatches), r.ChildWatches == nil); err != nil {
		return err
	}
	for _, v := range r.ChildWatches {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.PersistentWatches), r.PersistentWatches == nil); err != nil {
		return err
	}
	for _, v := range r.PersistentWatches {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteVectorStart(len(r.PersistentRecursiveWatches), r.PersistentRecursiveWatches == nil); err != nil {
		return err
	}
	for _, v := range r.PersistentRecursiveWatches {
		if err := enc.WriteString(v); err != nil {
			return err
		}
	}
	if err := enc.WriteVectorEnd(); err != nil {
		return err
	}
	if err := enc.WriteEnd(); err != nil {
		return err
	}
	return nil
}

func (r *SetWatches2) String() string {
	if r == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SetWatches2(%+v)", *r)
}
