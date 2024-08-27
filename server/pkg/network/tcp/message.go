package tcp

type Message struct {
	Id		uint64  // 表示Message的操作类型，会在protobuf中定义
	Data	[]byte	// 使用[]byte的原因是可以将json/protobuf等序列化为byte流，通用性好
}