package fn

// rpcf_get_core_time
func x11112(c conn, _ Dict) Dict {
	c.Call(0x11112)

	return Dict{
		"time":   c.GetI(),
		"tzname": c.GetS(),
	}
}
