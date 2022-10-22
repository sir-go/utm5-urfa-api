package fn

// rpcf_get_user_contact
func x2041(c conn, p Dict) Dict {
	c.Call(0x2041)
	putI(c, p, "cid")
	c.Send()

	return Dict{
		"descr":       c.GetS(),
		"reason":      c.GetS(),
		"person":      c.GetS(),
		"short_name":  c.GetS(),
		"contact":     c.GetS(),
		"email":       c.GetS(),
		"id_exec_man": c.GetI(),
	}
}
