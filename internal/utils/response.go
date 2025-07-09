package utils

import "github.com/xryar/golang-grpc-ecommerce/pb/common"

func SuccessResponse() *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    "Success",
	}
}
