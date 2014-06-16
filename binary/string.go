package binary

import "io"

type String string

// String

func (self String) Equals(other Binary) bool {
    return self == other
}

func (self String) Less(other Binary) bool {
    if o, ok := other.(String); ok {
        return self < o
    } else {
        panic("Cannot compare unequal types")
    }
}

func (self String) ByteSize() int {
    return len(self)+4
}

func (self String) WriteTo(w io.Writer) (n int64, err error) {
    var n_ int
    _, err = UInt32(len(self)).WriteTo(w)
    if err != nil { return n, err }
    n_, err = w.Write([]byte(self))
    return int64(n_+4), err
}

func ReadString(r io.Reader) String {
    length := int(ReadUInt32(r))
    bytes := make([]byte, length)
    _, err := io.ReadFull(r, bytes)
    if err != nil { panic(err) }
    return String(bytes)
}
