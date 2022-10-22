package fn

// rpcf_add_card_owner
func x4291(c conn, p Dict) Dict {
	c.Call(0x4291)
	putI(c, p, "pool_id")
	putI(c, p, "owned_id")
	c.Send()

	return nil
}
