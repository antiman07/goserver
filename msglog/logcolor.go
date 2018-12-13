package msglog

// 使用github.com/golog的cellnet配色方案
const LogColorDefine = `
{
	"Rule":[
		{"Text":"panic:","Color":"Red"},
		{"Text":"[DB]","Color":"Green"},
		{"Text":"#http.listen","Color":"Blue"},
		{"Text":"#http.recv","Color":"Blue"},
		{"Text":"#http.send","Color":"Purple"},

		{"Text":"#tcp.listen","Color":"Blue"},
		{"Text":"#tcp.accepted","Color":"Blue"},
		{"Text":"#tcp.closed","Color":"Blue"},
		{"Text":"#tcp.recv","Color":"Blue"},
		{"Text":"#tcp.send","Color":"Purple"},
		{"Text":"#tcp.connected","Color":"Blue"},

		{"Text":"#udp.listen","Color":"Blue"},
		{"Text":"#udp.recv","Color":"Blue"},
		{"Text":"#udp.send","Color":"Purple"},

		{"Text":"#rpc.recv","Color":"Blue"},
		{"Text":"#rpc.send","Color":"Purple"},

		{"Text":"#relay.recv","Color":"Blue"},
		{"Text":"#relay.send","Color":"Purple"}
	]
}
`
