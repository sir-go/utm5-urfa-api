package fn

// rpcf_liburfa_list
func x0040(c conn, _ Dict) Dict {
	c.Call(0x0040)
	return Dict{"libs": getMapOf(c, func() Dict {
		return Dict{
			"module":  c.GetS(),
			"version": c.GetS(),
			"path":    c.GetS(),
		}
	})}
}
