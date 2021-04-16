package define

//连接中保存的用户数据
type UserItem struct {
	UserId   int
	UserKey  string
	Nickname string
	Avatar   string
}

//客户端消息通用结构
type ClientMessage struct {
	MessageType int    `json:"message_type"`
	UserId      int    `json:"user_id"`  //用户ID
	UserKey     string `json:"user_key"` //用户key
	Data        struct {
		TypeId int `json:"type_id"` //房间类型
		RoomId int `json:"room_id"` //-1表示单人房
		SeatId int `json:"seat_id"`
	} `json:"data"` //附加参数
}

//错误消息
type ResultError struct {
	MessageId int    `json:"messageId"` //消息ID
	ErrorCode int    `json:"errorCode"` //错误码
	ErrorMsg  string `json:"errorMsg"`  //错误信息
}

type ResultBase struct {
	Result bool `json:"result"`
}

type ResultData struct {
	Result bool        `json:"result"`
	Data   interface{} `json:"data"`
}

//心跳包服务端回包
type TagServer_HeartBeat struct {
	MessageType int `json:"message_type"`
}

const (
	ERROR_TYPE_AUTH     = 501 //身份校验错误
	HEARTBEAT_MSG_ID    = 100 //心跳包
	SERVER_MSG_ID_ERROR = 200 //错误消息
)
