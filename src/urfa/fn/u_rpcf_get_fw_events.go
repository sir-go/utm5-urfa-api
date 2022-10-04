package fn

// rpcf_get_fw_events
func x0048(c conn, _ Dict) Dict {
	c.Call(0x0048)

	return Dict{
		"fw_events": c.GetL(),
	}
}
