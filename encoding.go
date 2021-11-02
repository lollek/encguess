package main

type Encoding int

const (
    UNKNOWN Encoding = iota
    CHARSET_UTF8 Encoding = iota
)

func (e Encoding) String() string {
    switch e {
    default: return "UNKNOWN"
    case CHARSET_UTF8: return "UTF-8"
    }
}
