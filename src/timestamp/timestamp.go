package timestamp

import (
    "fmt"
    "strings"
    "time"
)

type Format int
const (
    Unknown Format = iota
    Verbose
    Seconds
)

type options struct {
    UseUTC bool
    Representation Format
}

var Options = options{true,Verbose}

func Now() string {
    now := time.Now()
    if Options.UseUTC { now = now.UTC() }
    switch Options.Representation {
    case Verbose:
        return now.String()
    case Seconds:
        fallthrough
    default:
        return fmt.Sprintf("%d", now.Unix())
    }
}

func AsFormat(s string) Format {
    switch strings.ToLower(s) {
    case "verbose":
        return Verbose
    case "seconds":
        return Seconds
    default:
        return Unknown
    }
}