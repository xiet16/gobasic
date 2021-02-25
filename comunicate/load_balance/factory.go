package load_balance

type LbType int

const (
	LbRandom LbType = iota
	LbRoundRobin
	LbWeightRoundRobin
	LbConsistentHash
)

func LoadBalanceFactory(lbType LbType) LoadBalance {
	switch lbType {
	case LbWeightRoundRobin:
		return &WeightRoundRobinBalance{}
	default:
		return &WeightRoundRobinBalance{}
	}
}
