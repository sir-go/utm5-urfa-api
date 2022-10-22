package fn

// rpcf_count_switches
func x1157(c conn, p Dict) Dict {
	c.Call(0x1157)
	putI(c, p, "field")
	putS(c, p, "value")
	c.Send()

	return Dict{
		"count": c.GetI(),
	}
}
