// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestartDBInstanceResponseBody
	GetRequestId() *string
}

type RestartDBInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 149C517D-B586-47BE-A107-8673E0ED77C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestartDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBInstanceResponseBody) SetRequestId(v string) *RestartDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
