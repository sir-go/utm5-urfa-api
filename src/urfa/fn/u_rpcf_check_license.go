package fn

// rpcf_check_license
func x102(c conn, _ Dict) Dict {
	c.Call(0x102)

	return Dict{
		"has_no_license": c.GetI(),
	}
}
