//go:build !go1.16
// +build !go1.16

package plugin

import _ "unsafe"

// NetErrClosed aliases the internal error type poll.ErrNetClosing.
// FUTURE: When Go 1.16 is the minimum supported version for go-plugin, switch to net.ErrClosed.
//go:linkname NetErrClosed internal/poll.ErrNetClosing
var NetErrClosed error
