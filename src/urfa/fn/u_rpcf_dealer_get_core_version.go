package fn

// rpcf_dealer_get_core_version
func x70000047(c conn, _ Dict) Dict {
	c.Call(0x70000047)

	return Dict{
		"version_string": c.GetS(),
	}
}
