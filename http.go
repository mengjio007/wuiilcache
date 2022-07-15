package wuiilcache

import "github.com/gin-gonic/gin"

func REngine(r *gin.Engine) {
	c := r.Group("/cache/")
	c.POST("/:Group/:Key", Deal)
}

func Deal(c *gin.Context) {
	var (
		g string
		k string
	)
	for _, v := range c.Params {
		if v.Key == "Group" {
			g = v.Value
		}
		if v.Key == "Key" {
			k = v.Value
		}
	}
	var d = struct {
		Op   string      `json:"op"`
		Data interface{} `json:"data"`
	}{}
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(200, gin.H{"error": "don't kn error"})
		return
	}
	if len(d.Op) > 0 {
		// todo Set
	} else {
		group := GetGroup(g)
		if nil == group {
			c.JSON(200, gin.H{"error": "no this group"})
			return
		}
		if view, err := group.Get(k); nil != err {
			c.JSON(200, gin.H{"error": err.Error()})

		} else {
			c.Header("Content-Type", "application/octet-stream")
			c.JSON(200, view.ByteSlice())
		}
	}
}
