package fn

// rpcf_del_nf_provider
func x5063(c conn, p Dict) Dict {
	c.Call(0x5063)
	putI(c, p, "id")
	c.Send()

	return Dict{
		"error_message": c.GetS(),
	}
}
