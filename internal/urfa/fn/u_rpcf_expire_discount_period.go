package fn

// rpcf_expire_discount_period
func x2606(c conn, p Dict) Dict {
	c.Call(0x2606)
	putI(c, p, "dpid")
	c.Send()

	return nil
}
