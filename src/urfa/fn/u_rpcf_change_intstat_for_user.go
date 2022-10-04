package fn

// rpcf_change_intstat_for_user
func x2003(c conn, p Dict) Dict {
	c.Call(0x2003)
	putI(c, p, "user_id")
	putI(c, p, "need_block")
	c.Send()

	return nil
}
