package factory

// 工厂
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

// FactoryMethods 工厂方法接口
type FactoryMethods interface {
	CreateParser() Factory
}

// FactoryMethodsImpl1  的工厂类
type FactoryMethodsImpl1 struct {
}

// CreateParser 工厂方法实现1
func (y FactoryMethodsImpl1) CreateParser() Factory {
	return implement2{}
}

// FactoryMethodsImpl2 的工厂类
type FactoryMethodsImpl2 struct {
}

// CreateParser 工厂方法实现2
func (j FactoryMethodsImpl2) CreateParser() Factory {
	return implement1{}
}

// NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) FactoryMethods {
	switch t {
	case "1":
		return FactoryMethodsImpl1{}
	case "2":
		return FactoryMethodsImpl2{}
	}
	return nil
}
