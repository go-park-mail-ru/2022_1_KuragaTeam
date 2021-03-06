// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/microservices/profile/proto/profile_grpc.pb.go

// Package usecase is a generated GoMock package.
package usecase

import (
	context "context"
	proto "myapp/internal/microservices/profile/proto"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockProfileClient is a mock of ProfileClient interface.
type MockProfileClient struct {
	ctrl     *gomock.Controller
	recorder *MockProfileClientMockRecorder
}

// MockProfileClientMockRecorder is the mock recorder for MockProfileClient.
type MockProfileClientMockRecorder struct {
	mock *MockProfileClient
}

// NewMockProfileClient creates a new mock instance.
func NewMockProfileClient(ctrl *gomock.Controller) *MockProfileClient {
	mock := &MockProfileClient{ctrl: ctrl}
	mock.recorder = &MockProfileClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileClient) EXPECT() *MockProfileClientMockRecorder {
	return m.recorder
}

// AddLike mocks base method.
func (m *MockProfileClient) AddLike(ctx context.Context, in *proto.LikeData, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddLike", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLike indicates an expected call of AddLike.
func (mr *MockProfileClientMockRecorder) AddLike(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLike", reflect.TypeOf((*MockProfileClient)(nil).AddLike), varargs...)
}

// CheckPaymentsToken mocks base method.
func (m *MockProfileClient) CheckPaymentsToken(ctx context.Context, in *proto.CheckTokenData, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckPaymentsToken", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPaymentsToken indicates an expected call of CheckPaymentsToken.
func (mr *MockProfileClientMockRecorder) CheckPaymentsToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPaymentsToken", reflect.TypeOf((*MockProfileClient)(nil).CheckPaymentsToken), varargs...)
}

// CheckToken mocks base method.
func (m *MockProfileClient) CheckToken(ctx context.Context, in *proto.Token, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CheckToken", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken.
func (mr *MockProfileClientMockRecorder) CheckToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockProfileClient)(nil).CheckToken), varargs...)
}

// CreatePayment mocks base method.
func (m *MockProfileClient) CreatePayment(ctx context.Context, in *proto.CheckTokenData, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreatePayment", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockProfileClientMockRecorder) CreatePayment(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockProfileClient)(nil).CreatePayment), varargs...)
}

// CreateSubscribe mocks base method.
func (m *MockProfileClient) CreateSubscribe(ctx context.Context, in *proto.SubscribeData, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSubscribe", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscribe indicates an expected call of CreateSubscribe.
func (mr *MockProfileClientMockRecorder) CreateSubscribe(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscribe", reflect.TypeOf((*MockProfileClient)(nil).CreateSubscribe), varargs...)
}

// EditAvatar mocks base method.
func (m *MockProfileClient) EditAvatar(ctx context.Context, in *proto.EditAvatarData, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditAvatar", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditAvatar indicates an expected call of EditAvatar.
func (mr *MockProfileClientMockRecorder) EditAvatar(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAvatar", reflect.TypeOf((*MockProfileClient)(nil).EditAvatar), varargs...)
}

// EditProfile mocks base method.
func (m *MockProfileClient) EditProfile(ctx context.Context, in *proto.EditProfileData, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditProfile", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockProfileClientMockRecorder) EditProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockProfileClient)(nil).EditProfile), varargs...)
}

// GetAvatar mocks base method.
func (m *MockProfileClient) GetAvatar(ctx context.Context, in *proto.UserID, opts ...grpc.CallOption) (*proto.FileName, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetAvatar", varargs...)
	ret0, _ := ret[0].(*proto.FileName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatar indicates an expected call of GetAvatar.
func (mr *MockProfileClientMockRecorder) GetAvatar(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockProfileClient)(nil).GetAvatar), varargs...)
}

// GetFavorites mocks base method.
func (m *MockProfileClient) GetFavorites(ctx context.Context, in *proto.UserID, opts ...grpc.CallOption) (*proto.Favorites, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetFavorites", varargs...)
	ret0, _ := ret[0].(*proto.Favorites)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockProfileClientMockRecorder) GetFavorites(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockProfileClient)(nil).GetFavorites), varargs...)
}

// GetMovieRating mocks base method.
func (m *MockProfileClient) GetMovieRating(ctx context.Context, in *proto.MovieRating, opts ...grpc.CallOption) (*proto.Rating, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMovieRating", varargs...)
	ret0, _ := ret[0].(*proto.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieRating indicates an expected call of GetMovieRating.
func (mr *MockProfileClientMockRecorder) GetMovieRating(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieRating", reflect.TypeOf((*MockProfileClient)(nil).GetMovieRating), varargs...)
}

// GetPaymentsToken mocks base method.
func (m *MockProfileClient) GetPaymentsToken(ctx context.Context, in *proto.UserID, opts ...grpc.CallOption) (*proto.Token, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPaymentsToken", varargs...)
	ret0, _ := ret[0].(*proto.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentsToken indicates an expected call of GetPaymentsToken.
func (mr *MockProfileClientMockRecorder) GetPaymentsToken(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentsToken", reflect.TypeOf((*MockProfileClient)(nil).GetPaymentsToken), varargs...)
}

// GetUserProfile mocks base method.
func (m *MockProfileClient) GetUserProfile(ctx context.Context, in *proto.UserID, opts ...grpc.CallOption) (*proto.ProfileData, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserProfile", varargs...)
	ret0, _ := ret[0].(*proto.ProfileData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockProfileClientMockRecorder) GetUserProfile(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockProfileClient)(nil).GetUserProfile), varargs...)
}

// IsSubscription mocks base method.
func (m *MockProfileClient) IsSubscription(ctx context.Context, in *proto.UserID, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IsSubscription", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSubscription indicates an expected call of IsSubscription.
func (mr *MockProfileClientMockRecorder) IsSubscription(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubscription", reflect.TypeOf((*MockProfileClient)(nil).IsSubscription), varargs...)
}

// RemoveLike mocks base method.
func (m *MockProfileClient) RemoveLike(ctx context.Context, in *proto.LikeData, opts ...grpc.CallOption) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveLike", varargs...)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLike indicates an expected call of RemoveLike.
func (mr *MockProfileClientMockRecorder) RemoveLike(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLike", reflect.TypeOf((*MockProfileClient)(nil).RemoveLike), varargs...)
}

// UploadAvatar mocks base method.
func (m *MockProfileClient) UploadAvatar(ctx context.Context, in *proto.UploadInputFile, opts ...grpc.CallOption) (*proto.FileName, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadAvatar", varargs...)
	ret0, _ := ret[0].(*proto.FileName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockProfileClientMockRecorder) UploadAvatar(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockProfileClient)(nil).UploadAvatar), varargs...)
}

// MockProfileServer is a mock of ProfileServer interface.
type MockProfileServer struct {
	ctrl     *gomock.Controller
	recorder *MockProfileServerMockRecorder
}

// MockProfileServerMockRecorder is the mock recorder for MockProfileServer.
type MockProfileServerMockRecorder struct {
	mock *MockProfileServer
}

// NewMockProfileServer creates a new mock instance.
func NewMockProfileServer(ctrl *gomock.Controller) *MockProfileServer {
	mock := &MockProfileServer{ctrl: ctrl}
	mock.recorder = &MockProfileServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileServer) EXPECT() *MockProfileServerMockRecorder {
	return m.recorder
}

// AddLike mocks base method.
func (m *MockProfileServer) AddLike(arg0 context.Context, arg1 *proto.LikeData) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLike", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLike indicates an expected call of AddLike.
func (mr *MockProfileServerMockRecorder) AddLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLike", reflect.TypeOf((*MockProfileServer)(nil).AddLike), arg0, arg1)
}

// CheckPaymentsToken mocks base method.
func (m *MockProfileServer) CheckPaymentsToken(arg0 context.Context, arg1 *proto.CheckTokenData) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPaymentsToken", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPaymentsToken indicates an expected call of CheckPaymentsToken.
func (mr *MockProfileServerMockRecorder) CheckPaymentsToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPaymentsToken", reflect.TypeOf((*MockProfileServer)(nil).CheckPaymentsToken), arg0, arg1)
}

// CheckToken mocks base method.
func (m *MockProfileServer) CheckToken(arg0 context.Context, arg1 *proto.Token) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckToken", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckToken indicates an expected call of CheckToken.
func (mr *MockProfileServerMockRecorder) CheckToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckToken", reflect.TypeOf((*MockProfileServer)(nil).CheckToken), arg0, arg1)
}

// CreatePayment mocks base method.
func (m *MockProfileServer) CreatePayment(arg0 context.Context, arg1 *proto.CheckTokenData) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockProfileServerMockRecorder) CreatePayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockProfileServer)(nil).CreatePayment), arg0, arg1)
}

// CreateSubscribe mocks base method.
func (m *MockProfileServer) CreateSubscribe(arg0 context.Context, arg1 *proto.SubscribeData) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscribe", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSubscribe indicates an expected call of CreateSubscribe.
func (mr *MockProfileServerMockRecorder) CreateSubscribe(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscribe", reflect.TypeOf((*MockProfileServer)(nil).CreateSubscribe), arg0, arg1)
}

// EditAvatar mocks base method.
func (m *MockProfileServer) EditAvatar(arg0 context.Context, arg1 *proto.EditAvatarData) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditAvatar", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditAvatar indicates an expected call of EditAvatar.
func (mr *MockProfileServerMockRecorder) EditAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditAvatar", reflect.TypeOf((*MockProfileServer)(nil).EditAvatar), arg0, arg1)
}

// EditProfile mocks base method.
func (m *MockProfileServer) EditProfile(arg0 context.Context, arg1 *proto.EditProfileData) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditProfile", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditProfile indicates an expected call of EditProfile.
func (mr *MockProfileServerMockRecorder) EditProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditProfile", reflect.TypeOf((*MockProfileServer)(nil).EditProfile), arg0, arg1)
}

// GetAvatar mocks base method.
func (m *MockProfileServer) GetAvatar(arg0 context.Context, arg1 *proto.UserID) (*proto.FileName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvatar", arg0, arg1)
	ret0, _ := ret[0].(*proto.FileName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvatar indicates an expected call of GetAvatar.
func (mr *MockProfileServerMockRecorder) GetAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvatar", reflect.TypeOf((*MockProfileServer)(nil).GetAvatar), arg0, arg1)
}

// GetFavorites mocks base method.
func (m *MockProfileServer) GetFavorites(arg0 context.Context, arg1 *proto.UserID) (*proto.Favorites, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFavorites", arg0, arg1)
	ret0, _ := ret[0].(*proto.Favorites)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockProfileServerMockRecorder) GetFavorites(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockProfileServer)(nil).GetFavorites), arg0, arg1)
}

// GetMovieRating mocks base method.
func (m *MockProfileServer) GetMovieRating(arg0 context.Context, arg1 *proto.MovieRating) (*proto.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMovieRating", arg0, arg1)
	ret0, _ := ret[0].(*proto.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMovieRating indicates an expected call of GetMovieRating.
func (mr *MockProfileServerMockRecorder) GetMovieRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMovieRating", reflect.TypeOf((*MockProfileServer)(nil).GetMovieRating), arg0, arg1)
}

// GetPaymentsToken mocks base method.
func (m *MockProfileServer) GetPaymentsToken(arg0 context.Context, arg1 *proto.UserID) (*proto.Token, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentsToken", arg0, arg1)
	ret0, _ := ret[0].(*proto.Token)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentsToken indicates an expected call of GetPaymentsToken.
func (mr *MockProfileServerMockRecorder) GetPaymentsToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentsToken", reflect.TypeOf((*MockProfileServer)(nil).GetPaymentsToken), arg0, arg1)
}

// GetUserProfile mocks base method.
func (m *MockProfileServer) GetUserProfile(arg0 context.Context, arg1 *proto.UserID) (*proto.ProfileData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", arg0, arg1)
	ret0, _ := ret[0].(*proto.ProfileData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockProfileServerMockRecorder) GetUserProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockProfileServer)(nil).GetUserProfile), arg0, arg1)
}

// IsSubscription mocks base method.
func (m *MockProfileServer) IsSubscription(arg0 context.Context, arg1 *proto.UserID) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsSubscription", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsSubscription indicates an expected call of IsSubscription.
func (mr *MockProfileServerMockRecorder) IsSubscription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsSubscription", reflect.TypeOf((*MockProfileServer)(nil).IsSubscription), arg0, arg1)
}

// RemoveLike mocks base method.
func (m *MockProfileServer) RemoveLike(arg0 context.Context, arg1 *proto.LikeData) (*proto.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveLike", arg0, arg1)
	ret0, _ := ret[0].(*proto.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveLike indicates an expected call of RemoveLike.
func (mr *MockProfileServerMockRecorder) RemoveLike(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveLike", reflect.TypeOf((*MockProfileServer)(nil).RemoveLike), arg0, arg1)
}

// UploadAvatar mocks base method.
func (m *MockProfileServer) UploadAvatar(arg0 context.Context, arg1 *proto.UploadInputFile) (*proto.FileName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", arg0, arg1)
	ret0, _ := ret[0].(*proto.FileName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockProfileServerMockRecorder) UploadAvatar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockProfileServer)(nil).UploadAvatar), arg0, arg1)
}

// MockUnsafeProfileServer is a mock of UnsafeProfileServer interface.
type MockUnsafeProfileServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeProfileServerMockRecorder
}

// MockUnsafeProfileServerMockRecorder is the mock recorder for MockUnsafeProfileServer.
type MockUnsafeProfileServerMockRecorder struct {
	mock *MockUnsafeProfileServer
}

// NewMockUnsafeProfileServer creates a new mock instance.
func NewMockUnsafeProfileServer(ctrl *gomock.Controller) *MockUnsafeProfileServer {
	mock := &MockUnsafeProfileServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeProfileServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeProfileServer) EXPECT() *MockUnsafeProfileServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedProfileServer mocks base method.
func (m *MockUnsafeProfileServer) mustEmbedUnimplementedProfileServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedProfileServer")
}

// mustEmbedUnimplementedProfileServer indicates an expected call of mustEmbedUnimplementedProfileServer.
func (mr *MockUnsafeProfileServerMockRecorder) mustEmbedUnimplementedProfileServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedProfileServer", reflect.TypeOf((*MockUnsafeProfileServer)(nil).mustEmbedUnimplementedProfileServer))
}
