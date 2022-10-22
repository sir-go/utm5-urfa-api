package fn

// rpcf_core_build
func x0046(c conn, _ Dict) Dict {
	c.Call(0x0046)

	return Dict{
		"build": c.GetS(),
	}
}
