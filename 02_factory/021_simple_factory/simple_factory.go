package factory

// Factory 工厂
type Factory interface {
	Parse(data []byte)
}

type implement1 struct {
}

// Parse 实现1
func (J implement1) Parse(data []byte) {
	panic("implement me")
}

type implement2 struct {
}

// Parse 实现2
func (Y implement2) Parse(data []byte) {
	panic("implement me")
}

// 返回实现1和实现2的共同父类
func NewIRuleConfigParser(t string) Factory {
	switch t {
	case "1":
		return implement1{}
	case "2":
		return implement2{}
	}
	return nil
}
