package policy

type FIFOCache struct {
	maxBytes int64 // 允许的最大内存
	nBytes   int64 // 已使用内存

}
