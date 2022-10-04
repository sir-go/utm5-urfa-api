package fn

// rpcf_dealer_get_switch
func x70000071(c conn, p Dict) Dict {
	c.Call(0x70000071)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"id":          c.GetI(),
		"name":        c.GetS(),
		"location":    c.GetS(),
		"type":        c.GetI(),
		"ports_count": c.GetI(),
		"remote_id":   c.GetS(),
		"address":     c.GetA(),
		"login":       c.GetS(),
		"password":    c.GetS(),
	}
}
