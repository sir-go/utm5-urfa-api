package fn

// rpcf_unlink_user_tariff
func x3019(c conn, p Dict) Dict {
	c.Call(0x3019)
	putI(c, p, "user_id")
	putI(c, p, "account_id", 0)
	putI(c, p, "tariff_link_id")
	c.Send()

	return nil
}
