package fn

// rpcf_get_stats
func x0047(c conn, p Dict) Dict {
	c.Call(0x0047)
	putI(c, p, "type")
	c.Send()

	return Dict{
		"status":      c.GetI(),
		"uptime":      c.GetI(),
		"uptime_last": c.GetI(),
		"events":      c.GetI(),
		"events_last": c.GetI(),
		"errors":      c.GetI(),
		"errors_last": c.GetI(),
	}
}
