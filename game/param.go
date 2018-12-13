
package game

var (
	Runmode int //进程启动模式 前台运行 后台运行(1:前台运行模式 2:后台运行模式 默认2方式运行)

	Goover string //进程退出标识符

	GameName string //选择的游戏

	GameAddr string //游戏服务器地址

	LogserverAddr string //日志服务器地址

	BeginAccid int //机器人账号(++操作 账号ID 便于批量启动机器人,服务器识别区分)

	Robots int 			//启动的机器人数量
	RealWorkRobots int //退出进程之前 一直在线的玩家数量

	SpaceTime int //批量启动机器人时 间隔时长(见启动机器人for循环处)

	RoomID int //机器人进入的房间号

	Betmulti int //下注倍数

	BetmultiTime int //请求下注倍数阶段等待时长

	SpreadcardTime int //请求翻牌阶段等待时长
	Suspend		   int //轮到自己出牌时暂停的时长 单位是秒
)


