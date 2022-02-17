package demo

type ControllerI interface {
}

type GrpcClient struct {
}

//@autodig outgroup:restControllers
type ControllerDemo struct {
	DigReturn ControllerI
	Service   *Service
	config    string
}

//标记了autodig的类的Init() error会自动在初始化结束后执行
func (c *ControllerDemo) Init() error {
	c.config = "123"
	return nil
}

//@autodig
func NewGrpcClient() *GrpcClient {
	return &GrpcClient{}
}

//@autodig name:abGrpcClient
func NewAbGrpcClient() *GrpcClient {
	return &GrpcClient{}
}

//@autodig
type Service struct {
	Logger       []Logger `autodig:"ingroup:loggers"` //public字段自动注入
	config       string   //private字段会被忽略
	Config       string   `autodig:"-"` //public字段标记-会被忽略
	GrpcClient   *GrpcClient
	AbGrpcClient *GrpcClient `autodig:"name:abGrpcClient"`
}

type Logger struct {
}

//@autodig outgroup:loggers
func NewLogger() Logger {
	return Logger{}
}
