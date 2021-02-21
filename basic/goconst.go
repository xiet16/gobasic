package basic

//关键字 const iota

const (
	Title = "基础学习"
)

const (
	Monday = iota
)

const (
	NetCore = iota
	Go
	Vue
)

func ConstTest() {
	const iotaVar1, iotaVar2, iotaVar3 = iota, iota, iota // 都是0

	// 0 xiet 2
	const (
		iotaVar4 = iota
		iotaVar5 = "xiet"
		iotaVar6 = iota
	)
}
