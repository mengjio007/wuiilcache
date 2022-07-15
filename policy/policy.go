package policy

type Policy interface {
	Get(string) (Value, bool)
	Set(string, Value)
}

// NewCache 创建缓存
func NewCache(policy Policy) *Policy {
	return &policy
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

// 标准数据
type entry struct {
	key   string
	value Value
}
