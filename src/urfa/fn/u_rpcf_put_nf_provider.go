package fn

// rpcf_put_nf_provider
func x5062(c conn, p Dict) Dict {
	c.Call(0x5062)
	putI(c, p, "id")
	putI(c, p, "collector_id")
	putS(c, p, "name")
	putS(c, p, "comments")
	putA(c, p, "address")
	c.Send()

	return Dict{
		"error_message": c.GetS(),
	}
}
