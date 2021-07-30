package smartclient

import (
	"context"

	v0 "github.com/authzed/authzed-go/proto/authzed/api/v0"
	v1alpha1 "github.com/authzed/authzed-go/proto/authzed/api/v1alpha1"
	"google.golang.org/grpc"
)

func (sc *SmartClient) Check(ctx context.Context, req *v0.CheckRequest, opts ...grpc.CallOption) (*v0.CheckResponse, error) {
	requestKey := marshalCheckRequest(req)
	backend, err := sc.getConsistentBackend(requestKey)
	if err != nil {
		// TODO rewrite this error?
		return nil, err
	}

	return backend.client_v0.Check(ctx, req, opts...)
}

func (sc *SmartClient) ContentChangeCheck(ctx context.Context, req *v0.ContentChangeCheckRequest, opts ...grpc.CallOption) (*v0.CheckResponse, error) {
	requestKey := marshalContentChangeCheckRequest(req)
	backend, err := sc.getConsistentBackend(requestKey)
	if err != nil {
		// TODO rewrite this error?
		return nil, err
	}

	return backend.client_v0.ContentChangeCheck(ctx, req, opts...)
}

func (sc *SmartClient) Read(ctx context.Context, req *v0.ReadRequest, opts ...grpc.CallOption) (*v0.ReadResponse, error) {
	requestKey := marshalReadRequest(req)
	backend, err := sc.getConsistentBackend(requestKey)
	if err != nil {
		// TODO rewrite this error?
		return nil, err
	}

	return backend.client_v0.Read(ctx, req, opts...)
}

func (sc *SmartClient) Write(ctx context.Context, req *v0.WriteRequest, opts ...grpc.CallOption) (*v0.WriteResponse, error) {
	backend := sc.getRandomBackend()
	return backend.client_v0.Write(ctx, req, opts...)
}

func (sc *SmartClient) Expand(ctx context.Context, req *v0.ExpandRequest, opts ...grpc.CallOption) (*v0.ExpandResponse, error) {
	requestKey := marshalExpandRequest(req)
	backend, err := sc.getConsistentBackend(requestKey)
	if err != nil {
		// TODO rewrite this error?
		return nil, err
	}

	return backend.client_v0.Expand(ctx, req, opts...)
}

func (sc *SmartClient) Lookup(ctx context.Context, req *v0.LookupRequest, opts ...grpc.CallOption) (*v0.LookupResponse, error) {
	requestKey := marshalLookupRequest(req)
	backend, err := sc.getConsistentBackend(requestKey)
	if err != nil {
		// TODO rewrite this error?
		return nil, err
	}

	return backend.client_v0.Lookup(ctx, req, opts...)
}

func (sc *SmartClient) ReadConfig(ctx context.Context, req *v0.ReadConfigRequest, opts ...grpc.CallOption) (*v0.ReadConfigResponse, error) {
	requestKey := marshalReadConfigRequest(req)
	backend, err := sc.getConsistentBackend(requestKey)
	if err != nil {
		// TODO rewrite this error?
		return nil, err
	}

	return backend.client_v0.ReadConfig(ctx, req, opts...)
}

func (sc *SmartClient) WriteConfig(ctx context.Context, req *v0.WriteConfigRequest, opts ...grpc.CallOption) (*v0.WriteConfigResponse, error) {
	backend := sc.getRandomBackend()
	return backend.client_v0.WriteConfig(ctx, req, opts...)
}

func (sc *SmartClient) ReadSchema(ctx context.Context, req *v1alpha1.ReadSchemaRequest, opts ...grpc.CallOption) (*v1alpha1.ReadSchemaResponse, error) {
	requestKey := marshalReadSchemaRequest(req)
	backend, err := sc.getConsistentBackend(requestKey)
	if err != nil {
		// TODO rewrite this error?
		return nil, err
	}

	return backend.client_v1alpha1.ReadSchema(ctx, req, opts...)
}

func (sc *SmartClient) WriteSchema(ctx context.Context, req *v1alpha1.WriteSchemaRequest, opts ...grpc.CallOption) (*v1alpha1.WriteSchemaResponse, error) {
	backend := sc.getRandomBackend()
	return backend.client_v1alpha1.WriteSchema(ctx, req, opts...)
}
