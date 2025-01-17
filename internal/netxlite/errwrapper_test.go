package netxlite

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/ooni/probe-cli/v3/internal/atomicx"
)

func TestErrWrapper(t *testing.T) {
	t.Run("Error", func(t *testing.T) {
		err := &ErrWrapper{Failure: FailureDNSNXDOMAINError}
		if err.Error() != FailureDNSNXDOMAINError {
			t.Fatal("invalid return value")
		}
	})

	t.Run("Unwrap", func(t *testing.T) {
		err := &ErrWrapper{
			Failure:    FailureEOFError,
			WrappedErr: io.EOF,
		}
		if !errors.Is(err, io.EOF) {
			t.Fatal("cannot unwrap error")
		}
	})

	t.Run("MarshalJSON", func(t *testing.T) {
		wrappedErr := &ErrWrapper{
			Failure:    FailureEOFError,
			WrappedErr: io.EOF,
		}
		data, err := json.Marshal(wrappedErr)
		if err != nil {
			t.Fatal(err)
		}
		s := string(data)
		if s != "\""+FailureEOFError+"\"" {
			t.Fatal("invalid serialization", s)
		}
	})
}

func TestNewErrWrapper(t *testing.T) {
	t.Run("panics if the classifier is nil", func(t *testing.T) {
		recovered := &atomicx.Int64{}
		func() {
			defer func() {
				if recover() != nil {
					recovered.Add(1)
				}
			}()
			newErrWrapper(nil, CloseOperation, io.EOF)
		}()
		if recovered.Load() != 1 {
			t.Fatal("did not panic")
		}
	})

	t.Run("panics if the operation is empty", func(t *testing.T) {
		recovered := &atomicx.Int64{}
		func() {
			defer func() {
				if recover() != nil {
					recovered.Add(1)
				}
			}()
			newErrWrapper(classifyGenericError, "", io.EOF)
		}()
		if recovered.Load() != 1 {
			t.Fatal("did not panic")
		}
	})

	t.Run("panics if the error is nil", func(t *testing.T) {
		recovered := &atomicx.Int64{}
		func() {
			defer func() {
				if recover() != nil {
					recovered.Add(1)
				}
			}()
			newErrWrapper(classifyGenericError, CloseOperation, nil)
		}()
		if recovered.Load() != 1 {
			t.Fatal("did not panic")
		}
	})

	t.Run("otherwise, works as intended", func(t *testing.T) {
		ew := newErrWrapper(classifyGenericError, CloseOperation, io.EOF)
		if ew.Failure != FailureEOFError {
			t.Fatal("unexpected failure")
		}
		if ew.Operation != CloseOperation {
			t.Fatal("unexpected operation")
		}
		if ew.WrappedErr != io.EOF {
			t.Fatal("unexpected WrappedErr")
		}
	})

	t.Run("when the underlying error is already a wrapped error", func(t *testing.T) {
		ew := newErrWrapper(classifySyscallError, ReadOperation, ECONNRESET)
		var err1 error = ew
		err2 := fmt.Errorf("cannot read: %w", err1)
		ew2 := newErrWrapper(classifyGenericError, HTTPRoundTripOperation, err2)
		if ew2.Failure != ew.Failure {
			t.Fatal("not the same failure")
		}
		if ew2.Operation != HTTPRoundTripOperation {
			t.Fatal("not the same operation")
		}
		if ew2.WrappedErr != err2 {
			t.Fatal("invalid underlying error")
		}
	})
}

func TestNewTopLevelGenericErrWrapper(t *testing.T) {
	out := NewTopLevelGenericErrWrapper(io.EOF)
	if out.Failure != FailureEOFError {
		t.Fatal("invalid failure")
	}
	if out.Operation != TopLevelOperation {
		t.Fatal("invalid operation")
	}
	if !errors.Is(out, io.EOF) {
		t.Fatal("invalid underlying error using errors.Is")
	}
	if !errors.Is(out.WrappedErr, io.EOF) {
		t.Fatal("invalid WrappedErr")
	}
}

func TestClassifyOperation(t *testing.T) {
	t.Run("for connect", func(t *testing.T) {
		// You're doing HTTP and connect fails. You want to know
		// that connect failed not that HTTP failed.
		err := &ErrWrapper{Operation: ConnectOperation}
		if classifyOperation(err, HTTPRoundTripOperation) != ConnectOperation {
			t.Fatal("unexpected result")
		}
	})

	t.Run("for http_round_trip", func(t *testing.T) {
		// You're doing DoH and something fails inside HTTP. You want
		// to know about the internal HTTP error, not resolve.
		err := &ErrWrapper{Operation: HTTPRoundTripOperation}
		if classifyOperation(err, ResolveOperation) != HTTPRoundTripOperation {
			t.Fatal("unexpected result")
		}
	})

	t.Run("for resolve", func(t *testing.T) {
		// You're doing HTTP and the DNS fails. You want to
		// know that resolve failed.
		err := &ErrWrapper{Operation: ResolveOperation}
		if classifyOperation(err, HTTPRoundTripOperation) != ResolveOperation {
			t.Fatal("unexpected result")
		}
	})

	t.Run("for tls_handshake", func(t *testing.T) {
		// You're doing HTTP and the TLS handshake fails. You want
		// to know about a TLS handshake error.
		err := &ErrWrapper{Operation: TLSHandshakeOperation}
		if classifyOperation(err, HTTPRoundTripOperation) != TLSHandshakeOperation {
			t.Fatal("unexpected result")
		}
	})

	t.Run("for minor operation", func(t *testing.T) {
		// You just noticed that TLS handshake failed and you
		// have a child error telling you that read failed. Here
		// you want to know about a TLS handshake error.
		err := &ErrWrapper{Operation: ReadOperation}
		if classifyOperation(err, TLSHandshakeOperation) != TLSHandshakeOperation {
			t.Fatal("unexpected result")
		}
	})

	t.Run("for quic_handshake", func(t *testing.T) {
		// You're doing HTTP and the QUIC handshake fails. You want
		// to know about a QUIC handshake error.
		err := &ErrWrapper{Operation: QUICHandshakeOperation}
		if classifyOperation(err, HTTPRoundTripOperation) != QUICHandshakeOperation {
			t.Fatal("unexpected result")
		}
	})

	t.Run("for quic_handshake_start", func(t *testing.T) {
		// You're doing HTTP and the QUIC handshake fails. You want
		// to know about a QUIC handshake error.
		err := &ErrWrapper{Operation: "quic_handshake_start"}
		if classifyOperation(err, HTTPRoundTripOperation) != QUICHandshakeOperation {
			t.Fatal("unexpected result")
		}
	})

	t.Run("for quic_handshake_done", func(t *testing.T) {
		// You're doing HTTP and the QUIC handshake fails. You want
		// to know about a QUIC handshake error.
		err := &ErrWrapper{Operation: "quic_handshake_done"}
		if classifyOperation(err, HTTPRoundTripOperation) != QUICHandshakeOperation {
			t.Fatal("unexpected result")
		}
	})
}
