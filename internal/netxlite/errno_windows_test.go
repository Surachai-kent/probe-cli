// Code generated by go generate; DO NOT EDIT.
// Generated: 2021-11-15 13:49:27.135917 +0100 CET m=+0.452914042

package netxlite

import (
	"io"
	"syscall"
	"testing"

	"golang.org/x/sys/windows"
)

func TestClassifySyscallError(t *testing.T) {
	t.Run("for a non-syscall error", func(t *testing.T) {
		if v := classifySyscallError(io.EOF); v != "" {
			t.Fatalf("expected empty string, got '%s'", v)
		}
	})

	t.Run("for WSAECONNREFUSED", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAECONNREFUSED); v != FailureConnectionRefused {
			t.Fatalf("expected '%s', got '%s'", FailureConnectionRefused, v)
		}
	})

	t.Run("for WSAECONNRESET", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAECONNRESET); v != FailureConnectionReset {
			t.Fatalf("expected '%s', got '%s'", FailureConnectionReset, v)
		}
	})

	t.Run("for WSAEHOSTUNREACH", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEHOSTUNREACH); v != FailureHostUnreachable {
			t.Fatalf("expected '%s', got '%s'", FailureHostUnreachable, v)
		}
	})

	t.Run("for WSAETIMEDOUT", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAETIMEDOUT); v != FailureTimedOut {
			t.Fatalf("expected '%s', got '%s'", FailureTimedOut, v)
		}
	})

	t.Run("for WSAEAFNOSUPPORT", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEAFNOSUPPORT); v != FailureAddressFamilyNotSupported {
			t.Fatalf("expected '%s', got '%s'", FailureAddressFamilyNotSupported, v)
		}
	})

	t.Run("for WSAEADDRINUSE", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEADDRINUSE); v != FailureAddressInUse {
			t.Fatalf("expected '%s', got '%s'", FailureAddressInUse, v)
		}
	})

	t.Run("for WSAEADDRNOTAVAIL", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEADDRNOTAVAIL); v != FailureAddressNotAvailable {
			t.Fatalf("expected '%s', got '%s'", FailureAddressNotAvailable, v)
		}
	})

	t.Run("for WSAEISCONN", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEISCONN); v != FailureAlreadyConnected {
			t.Fatalf("expected '%s', got '%s'", FailureAlreadyConnected, v)
		}
	})

	t.Run("for WSAEFAULT", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEFAULT); v != FailureBadAddress {
			t.Fatalf("expected '%s', got '%s'", FailureBadAddress, v)
		}
	})

	t.Run("for WSAEBADF", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEBADF); v != FailureBadFileDescriptor {
			t.Fatalf("expected '%s', got '%s'", FailureBadFileDescriptor, v)
		}
	})

	t.Run("for WSAECONNABORTED", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAECONNABORTED); v != FailureConnectionAborted {
			t.Fatalf("expected '%s', got '%s'", FailureConnectionAborted, v)
		}
	})

	t.Run("for WSAEALREADY", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEALREADY); v != FailureConnectionAlreadyInProgress {
			t.Fatalf("expected '%s', got '%s'", FailureConnectionAlreadyInProgress, v)
		}
	})

	t.Run("for WSAEDESTADDRREQ", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEDESTADDRREQ); v != FailureDestinationAddressRequired {
			t.Fatalf("expected '%s', got '%s'", FailureDestinationAddressRequired, v)
		}
	})

	t.Run("for WSAEINTR", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEINTR); v != FailureInterrupted {
			t.Fatalf("expected '%s', got '%s'", FailureInterrupted, v)
		}
	})

	t.Run("for WSAEINVAL", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEINVAL); v != FailureInvalidArgument {
			t.Fatalf("expected '%s', got '%s'", FailureInvalidArgument, v)
		}
	})

	t.Run("for WSAEMSGSIZE", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEMSGSIZE); v != FailureMessageSize {
			t.Fatalf("expected '%s', got '%s'", FailureMessageSize, v)
		}
	})

	t.Run("for WSAENETDOWN", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAENETDOWN); v != FailureNetworkDown {
			t.Fatalf("expected '%s', got '%s'", FailureNetworkDown, v)
		}
	})

	t.Run("for WSAENETRESET", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAENETRESET); v != FailureNetworkReset {
			t.Fatalf("expected '%s', got '%s'", FailureNetworkReset, v)
		}
	})

	t.Run("for WSAENETUNREACH", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAENETUNREACH); v != FailureNetworkUnreachable {
			t.Fatalf("expected '%s', got '%s'", FailureNetworkUnreachable, v)
		}
	})

	t.Run("for WSAENOBUFS", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAENOBUFS); v != FailureNoBufferSpace {
			t.Fatalf("expected '%s', got '%s'", FailureNoBufferSpace, v)
		}
	})

	t.Run("for WSAENOPROTOOPT", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAENOPROTOOPT); v != FailureNoProtocolOption {
			t.Fatalf("expected '%s', got '%s'", FailureNoProtocolOption, v)
		}
	})

	t.Run("for WSAENOTSOCK", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAENOTSOCK); v != FailureNotASocket {
			t.Fatalf("expected '%s', got '%s'", FailureNotASocket, v)
		}
	})

	t.Run("for WSAENOTCONN", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAENOTCONN); v != FailureNotConnected {
			t.Fatalf("expected '%s', got '%s'", FailureNotConnected, v)
		}
	})

	t.Run("for WSAEWOULDBLOCK", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEWOULDBLOCK); v != FailureOperationWouldBlock {
			t.Fatalf("expected '%s', got '%s'", FailureOperationWouldBlock, v)
		}
	})

	t.Run("for WSAEACCES", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEACCES); v != FailurePermissionDenied {
			t.Fatalf("expected '%s', got '%s'", FailurePermissionDenied, v)
		}
	})

	t.Run("for WSAEPROTONOSUPPORT", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEPROTONOSUPPORT); v != FailureProtocolNotSupported {
			t.Fatalf("expected '%s', got '%s'", FailureProtocolNotSupported, v)
		}
	})

	t.Run("for WSAEPROTOTYPE", func(t *testing.T) {
		if v := classifySyscallError(windows.WSAEPROTOTYPE); v != FailureWrongProtocolType {
			t.Fatalf("expected '%s', got '%s'", FailureWrongProtocolType, v)
		}
	})

	t.Run("for WSANO_DATA", func(t *testing.T) {
		if v := classifySyscallError(windows.WSANO_DATA); v != FailureDNSNoAnswer {
			t.Fatalf("expected '%s', got '%s'", FailureDNSNoAnswer, v)
		}
	})

	t.Run("for WSANO_RECOVERY", func(t *testing.T) {
		if v := classifySyscallError(windows.WSANO_RECOVERY); v != FailureDNSNonRecoverableFailure {
			t.Fatalf("expected '%s', got '%s'", FailureDNSNonRecoverableFailure, v)
		}
	})

	t.Run("for WSATRY_AGAIN", func(t *testing.T) {
		if v := classifySyscallError(windows.WSATRY_AGAIN); v != FailureDNSTemporaryFailure {
			t.Fatalf("expected '%s', got '%s'", FailureDNSTemporaryFailure, v)
		}
	})

	t.Run("for the zero errno value", func(t *testing.T) {
		if v := classifySyscallError(syscall.Errno(0)); v != "" {
			t.Fatalf("expected empty string, got '%s'", v)
		}
	})
}
