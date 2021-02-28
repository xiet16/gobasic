package load_balance

import (
	"errors"
	"strconv"
)

type WeightRoundRobinBalance struct {
	curIndex int
	rss      []*WeightNode
	rsw      []int
}

type WeightNode struct {
	addr            string
	weight          int
	currentWeight   int
	effectiveWeight int
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("param len need 2")
	}

	parInt, err := strconv.ParseInt(params[1], 10, 64)
	if err != nil {
		return err
	}
	node := &WeightNode{
		addr:   params[0],
		weight: int(parInt),
	}
	node.effectiveWeight = node.weight
	r.rss = append(r.rss, node)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	total := 0
	var best *WeightNode
	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]

		//step1 统计所有权重的和
		total += w.effectiveWeight
		//step2 变更节点临时权重的临时权重+节点有效权重
		w.currentWeight += w.effectiveWeight
		//step3 有效权重默认与权重仙童，通讯异常时-1，通讯成功时+1， 知道恢复weigt大小
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}

		//step4 选择最大临时权重节点
		if best == nil || w.currentWeight > best.currentWeight {
			best = w
		}
	}

	if best == nil {
		return ""
	}

	//step5 变临时权重为临时权重-有效权重之和
	best.currentWeight -= total
	return best.addr
}

func (r *WeightRoundRobinBalance) Get(key string) (string, error) {
	return r.Next(), nil
}
