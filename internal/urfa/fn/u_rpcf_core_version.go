package fn

// rpcf_core_version
func x0045(c conn, _ Dict) Dict {
	c.Call(0x0045)

	return Dict{
		"core_version": c.GetS(),
	}
}
