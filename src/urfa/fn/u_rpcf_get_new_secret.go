package fn

// rpcf_get_new_secret
func x0060(c conn, p Dict) Dict {
	c.Call(0x0060)
	putI(c, p, "secret_size", 8)
	c.Send()

	return Dict{
		"error":  c.GetS(),
		"secret": c.GetS(),
	}
}
