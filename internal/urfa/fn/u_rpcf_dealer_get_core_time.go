package fn

// rpcf_dealer_get_core_time
func x70000046(c conn, _ Dict) Dict {
	c.Call(0x70000046)

	return Dict{
		"time":   c.GetI(),
		"tzname": c.GetS(),
	}
}
