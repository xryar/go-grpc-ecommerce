// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: product/product.proto

package product

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProductService_CreateProduct_FullMethodName     = "/product.ProductService/CreateProduct"
	ProductService_DetailProduct_FullMethodName     = "/product.ProductService/DetailProduct"
	ProductService_EditProduct_FullMethodName       = "/product.ProductService/EditProduct"
	ProductService_DeleteProduct_FullMethodName     = "/product.ProductService/DeleteProduct"
	ProductService_ListProduct_FullMethodName       = "/product.ProductService/ListProduct"
	ProductService_ListProductAdmin_FullMethodName  = "/product.ProductService/ListProductAdmin"
	ProductService_HighlightProducts_FullMethodName = "/product.ProductService/HighlightProducts"
)

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	DetailProduct(ctx context.Context, in *DetailProductRequest, opts ...grpc.CallOption) (*DetailProductResponse, error)
	EditProduct(ctx context.Context, in *EditProductRequest, opts ...grpc.CallOption) (*EditProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error)
	ListProductAdmin(ctx context.Context, in *ListProductAdminRequest, opts ...grpc.CallOption) (*ListProductAdminResponse, error)
	HighlightProducts(ctx context.Context, in *HighlightProductRequest, opts ...grpc.CallOption) (*HighlightProductResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, ProductService_CreateProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DetailProduct(ctx context.Context, in *DetailProductRequest, opts ...grpc.CallOption) (*DetailProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DetailProductResponse)
	err := c.cc.Invoke(ctx, ProductService_DetailProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) EditProduct(ctx context.Context, in *EditProductRequest, opts ...grpc.CallOption) (*EditProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditProductResponse)
	err := c.cc.Invoke(ctx, ProductService_EditProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, ProductService_DeleteProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProduct(ctx context.Context, in *ListProductRequest, opts ...grpc.CallOption) (*ListProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductResponse)
	err := c.cc.Invoke(ctx, ProductService_ListProduct_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) ListProductAdmin(ctx context.Context, in *ListProductAdminRequest, opts ...grpc.CallOption) (*ListProductAdminResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProductAdminResponse)
	err := c.cc.Invoke(ctx, ProductService_ListProductAdmin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) HighlightProducts(ctx context.Context, in *HighlightProductRequest, opts ...grpc.CallOption) (*HighlightProductResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HighlightProductResponse)
	err := c.cc.Invoke(ctx, ProductService_HighlightProducts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility.
type ProductServiceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	DetailProduct(context.Context, *DetailProductRequest) (*DetailProductResponse, error)
	EditProduct(context.Context, *EditProductRequest) (*EditProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error)
	ListProductAdmin(context.Context, *ListProductAdminRequest) (*ListProductAdminResponse, error)
	HighlightProducts(context.Context, *HighlightProductRequest) (*HighlightProductResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductServiceServer struct{}

func (UnimplementedProductServiceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductServiceServer) DetailProduct(context.Context, *DetailProductRequest) (*DetailProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetailProduct not implemented")
}
func (UnimplementedProductServiceServer) EditProduct(context.Context, *EditProductRequest) (*EditProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditProduct not implemented")
}
func (UnimplementedProductServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedProductServiceServer) ListProduct(context.Context, *ListProductRequest) (*ListProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProduct not implemented")
}
func (UnimplementedProductServiceServer) ListProductAdmin(context.Context, *ListProductAdminRequest) (*ListProductAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProductAdmin not implemented")
}
func (UnimplementedProductServiceServer) HighlightProducts(context.Context, *HighlightProductRequest) (*HighlightProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HighlightProducts not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}
func (UnimplementedProductServiceServer) testEmbeddedByValue()                        {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DetailProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DetailProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DetailProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DetailProduct(ctx, req.(*DetailProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_EditProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).EditProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_EditProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).EditProduct(ctx, req.(*EditProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ListProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProduct(ctx, req.(*ListProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_ListProductAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProductAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ListProductAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_ListProductAdmin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ListProductAdmin(ctx, req.(*ListProductAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_HighlightProducts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HighlightProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).HighlightProducts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductService_HighlightProducts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).HighlightProducts(ctx, req.(*HighlightProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "product.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductService_CreateProduct_Handler,
		},
		{
			MethodName: "DetailProduct",
			Handler:    _ProductService_DetailProduct_Handler,
		},
		{
			MethodName: "EditProduct",
			Handler:    _ProductService_EditProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _ProductService_DeleteProduct_Handler,
		},
		{
			MethodName: "ListProduct",
			Handler:    _ProductService_ListProduct_Handler,
		},
		{
			MethodName: "ListProductAdmin",
			Handler:    _ProductService_ListProductAdmin_Handler,
		},
		{
			MethodName: "HighlightProducts",
			Handler:    _ProductService_HighlightProducts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product/product.proto",
}
