// Code generated by astiav. DO NOT EDIT.
package astiav

import (
	"github.com/asticode/go-astikit"
)

type BuffersinkFlags astikit.BitFlags

func NewBuffersinkFlags(fs ...BuffersinkFlag) BuffersinkFlags {
	o := BuffersinkFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs BuffersinkFlags) Add(f BuffersinkFlag) BuffersinkFlags {
	return BuffersinkFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs BuffersinkFlags) Del(f BuffersinkFlag) BuffersinkFlags {
	return BuffersinkFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs BuffersinkFlags) Has(f BuffersinkFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type BuffersrcFlags astikit.BitFlags

func NewBuffersrcFlags(fs ...BuffersrcFlag) BuffersrcFlags {
	o := BuffersrcFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs BuffersrcFlags) Add(f BuffersrcFlag) BuffersrcFlags {
	return BuffersrcFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs BuffersrcFlags) Del(f BuffersrcFlag) BuffersrcFlags {
	return BuffersrcFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs BuffersrcFlags) Has(f BuffersrcFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type CodecContextFlags astikit.BitFlags

func NewCodecContextFlags(fs ...CodecContextFlag) CodecContextFlags {
	o := CodecContextFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs CodecContextFlags) Add(f CodecContextFlag) CodecContextFlags {
	return CodecContextFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs CodecContextFlags) Del(f CodecContextFlag) CodecContextFlags {
	return CodecContextFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs CodecContextFlags) Has(f CodecContextFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type CodecContextFlags2 astikit.BitFlags

func NewCodecContextFlags2(fs ...CodecContextFlag2) CodecContextFlags2 {
	o := CodecContextFlags2(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs CodecContextFlags2) Add(f CodecContextFlag2) CodecContextFlags2 {
	return CodecContextFlags2(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs CodecContextFlags2) Del(f CodecContextFlag2) CodecContextFlags2 {
	return CodecContextFlags2(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs CodecContextFlags2) Has(f CodecContextFlag2) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type CodecHardwareConfigMethodFlags astikit.BitFlags

func NewCodecHardwareConfigMethodFlags(fs ...CodecHardwareConfigMethodFlag) CodecHardwareConfigMethodFlags {
	o := CodecHardwareConfigMethodFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs CodecHardwareConfigMethodFlags) Add(f CodecHardwareConfigMethodFlag) CodecHardwareConfigMethodFlags {
	return CodecHardwareConfigMethodFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs CodecHardwareConfigMethodFlags) Del(f CodecHardwareConfigMethodFlag) CodecHardwareConfigMethodFlags {
	return CodecHardwareConfigMethodFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs CodecHardwareConfigMethodFlags) Has(f CodecHardwareConfigMethodFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type DictionaryFlags astikit.BitFlags

func NewDictionaryFlags(fs ...DictionaryFlag) DictionaryFlags {
	o := DictionaryFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs DictionaryFlags) Add(f DictionaryFlag) DictionaryFlags {
	return DictionaryFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs DictionaryFlags) Del(f DictionaryFlag) DictionaryFlags {
	return DictionaryFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs DictionaryFlags) Has(f DictionaryFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type FilterCommandFlags astikit.BitFlags

func NewFilterCommandFlags(fs ...FilterCommandFlag) FilterCommandFlags {
	o := FilterCommandFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FilterCommandFlags) Add(f FilterCommandFlag) FilterCommandFlags {
	return FilterCommandFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs FilterCommandFlags) Del(f FilterCommandFlag) FilterCommandFlags {
	return FilterCommandFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs FilterCommandFlags) Has(f FilterCommandFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type FormatContextFlags astikit.BitFlags

func NewFormatContextFlags(fs ...FormatContextFlag) FormatContextFlags {
	o := FormatContextFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FormatContextFlags) Add(f FormatContextFlag) FormatContextFlags {
	return FormatContextFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs FormatContextFlags) Del(f FormatContextFlag) FormatContextFlags {
	return FormatContextFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs FormatContextFlags) Has(f FormatContextFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type FormatContextCtxFlags astikit.BitFlags

func NewFormatContextCtxFlags(fs ...FormatContextCtxFlag) FormatContextCtxFlags {
	o := FormatContextCtxFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FormatContextCtxFlags) Add(f FormatContextCtxFlag) FormatContextCtxFlags {
	return FormatContextCtxFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs FormatContextCtxFlags) Del(f FormatContextCtxFlag) FormatContextCtxFlags {
	return FormatContextCtxFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs FormatContextCtxFlags) Has(f FormatContextCtxFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type FormatEventFlags astikit.BitFlags

func NewFormatEventFlags(fs ...FormatEventFlag) FormatEventFlags {
	o := FormatEventFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs FormatEventFlags) Add(f FormatEventFlag) FormatEventFlags {
	return FormatEventFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs FormatEventFlags) Del(f FormatEventFlag) FormatEventFlags {
	return FormatEventFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs FormatEventFlags) Has(f FormatEventFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type IOContextFlags astikit.BitFlags

func NewIOContextFlags(fs ...IOContextFlag) IOContextFlags {
	o := IOContextFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs IOContextFlags) Add(f IOContextFlag) IOContextFlags {
	return IOContextFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs IOContextFlags) Del(f IOContextFlag) IOContextFlags {
	return IOContextFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs IOContextFlags) Has(f IOContextFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type IOFormatFlags astikit.BitFlags

func NewIOFormatFlags(fs ...IOFormatFlag) IOFormatFlags {
	o := IOFormatFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs IOFormatFlags) Add(f IOFormatFlag) IOFormatFlags {
	return IOFormatFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs IOFormatFlags) Del(f IOFormatFlag) IOFormatFlags {
	return IOFormatFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs IOFormatFlags) Has(f IOFormatFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type PacketFlags astikit.BitFlags

func NewPacketFlags(fs ...PacketFlag) PacketFlags {
	o := PacketFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs PacketFlags) Add(f PacketFlag) PacketFlags {
	return PacketFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs PacketFlags) Del(f PacketFlag) PacketFlags {
	return PacketFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs PacketFlags) Has(f PacketFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type SeekFlags astikit.BitFlags

func NewSeekFlags(fs ...SeekFlag) SeekFlags {
	o := SeekFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs SeekFlags) Add(f SeekFlag) SeekFlags {
	return SeekFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs SeekFlags) Del(f SeekFlag) SeekFlags {
	return SeekFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs SeekFlags) Has(f SeekFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type SoftwareScaleContextFlags astikit.BitFlags

func NewSoftwareScaleContextFlags(fs ...SoftwareScaleContextFlag) SoftwareScaleContextFlags {
	o := SoftwareScaleContextFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs SoftwareScaleContextFlags) Add(f SoftwareScaleContextFlag) SoftwareScaleContextFlags {
	return SoftwareScaleContextFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs SoftwareScaleContextFlags) Del(f SoftwareScaleContextFlag) SoftwareScaleContextFlags {
	return SoftwareScaleContextFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs SoftwareScaleContextFlags) Has(f SoftwareScaleContextFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }

type StreamEventFlags astikit.BitFlags

func NewStreamEventFlags(fs ...StreamEventFlag) StreamEventFlags {
	o := StreamEventFlags(0)
	for _, f := range fs {
		o = o.Add(f)
	}
	return o
}

func (fs StreamEventFlags) Add(f StreamEventFlag) StreamEventFlags {
	return StreamEventFlags(astikit.BitFlags(fs).Add(uint64(f)))
}

func (fs StreamEventFlags) Del(f StreamEventFlag) StreamEventFlags {
	return StreamEventFlags(astikit.BitFlags(fs).Del(uint64(f)))
}

func (fs StreamEventFlags) Has(f StreamEventFlag) bool { return astikit.BitFlags(fs).Has(uint64(f)) }
