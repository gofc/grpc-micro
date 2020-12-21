package code

import (
	"fmt"
	"io"

	"github.com/pkg/errors"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewGRPC サーバ間伝搬用のgrpc error
func NewGRPC(code codes.Code, err error) error {
	if err == nil {
		return newGrpcError(code, nil, "", nil)
	}
	if g, ok := err.(GrpcMessage); ok {
		return newWithGrpcMessage(code, err, g)
	}
	return newGrpcError(code, nil, "", err)
}

// NewGRPCWithStatus 外部(client)に返す用のgrpc error、基本localize_interceptorのみで使用
func NewGRPCWithStatus(stat *status.Status, err error) error {
	if grpcErr, ok := err.(*GrpcError); ok {
		return &GrpcError{
			status: stat,
			err:    grpcErr.err,
		}
	}

	if stat == nil {
		return newGrpcError(codes.Internal, nil, "", err)
	}
	return newGrpcError(stat.Code(), stat, stat.Message(), err)
}

//GrpcError Grpcエラー
type GrpcError struct {
	status *status.Status
	err    error
}

const (
	grpcErrorFormatCode = "rpc error: code = %s"
	grpcErrorFormatFull = "rpc error: code = %s desc = %s"
)

// nolint : gocyclo
func newGrpcError(code codes.Code, s *status.Status, m string, e error) error {
	stat := s
	msg := m
	var err error
	if e != nil {
		if ge, ok := e.(grpcStatus); ok && ge.GRPCStatus().Code() == code {
			// gRPCからのerror && エラーコードが同じ
			if msg == "" {
				msg = ge.GRPCStatus().Message()
			}
			err = e // gRPCのエラーをそのまま利用
		} else {
			// gRPC以外からのerror
			if msg == "" {
				msg = grpcStatusMessage(e)
				err = errors.Wrapf(e, grpcErrorFormatCode, code)
			} else {
				err = errors.Wrapf(e, grpcErrorFormatFull, code, msg)
			}
		}
	} else {
		// errorなし
		err = errors.New(fmt.Sprintf(grpcErrorFormatFull, code, msg))
	}
	if stat == nil {
		stat = status.New(code, msg)
	}
	if stat.Err() == nil {
		return nil // OKのためnil
	}
	return &GrpcError{
		status: stat,
		err:    err,
	}
}

// Error message AppError
func (ge *GrpcError) Error() string {
	if ge == nil {
		return ""
	}
	return ge.err.Error()
}

// GRPCStatus gRPCステータス
func (ge *GrpcError) GRPCStatus() *status.Status {
	if ge == nil {
		return nil
	}
	return ge.status
}

// Format errorのフォーマット
func (ge *GrpcError) Format(f fmt.State, c rune) {
	if ge == nil {
		return
	}
	if errFmt, ok := ge.err.(fmt.Formatter); ok {
		errFmt.Format(f, c)
		return
	}
	io.WriteString(f, ge.err.Error()) // nolint : errcheck
}

func newWithGrpcMessage(code codes.Code, orgErr error, m GrpcMessage) error {
	detail, err := m.ToGrpcMessage()
	if err != nil {
		return grpcConvertError(orgErr, err)
	}

	st := status.New(code, grpcStatusMessage(orgErr))
	st, err = st.WithDetails(detail)
	if err != nil {
		return grpcConvertError(orgErr, err)
	}

	return newGrpcError(code, st, "", orgErr)
}

func grpcStatusMessage(err error) string {
	return err.Error()
}

func grpcConvertError(orgErr error, cause error) error {
	return newGrpcError(codes.Internal, nil,
		fmt.Sprintf("fail to convert error: %+v", cause), orgErr)
}

// GrpcMessage gRPCメッセージに変換可能なもの
type GrpcMessage interface {
	ToGrpcMessage() (proto.Message, error)
}

// grpcStatus gRPCステータスに変換可能なもの
type grpcStatus interface {
	GRPCStatus() *status.Status
}
