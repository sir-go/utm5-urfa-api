package fn

// rpcf_get_switch
func x1150(c conn, p Dict) Dict {
	c.Call(0x1150)
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
