package fn

// rpcf_liburfa_symtab
func x0044(c conn, _ Dict) Dict {
	c.Call(0x0044)
	return Dict{"libs": getMapOf(c, func() Dict {
		return Dict{
			"id":     c.GetI(),
			"name":   c.GetS(),
			"module": c.GetS(),
		}
	})}
}
