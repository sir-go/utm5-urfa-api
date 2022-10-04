package fn

// rpcf_cancel_payment_for_account
func x3111(c conn, p Dict) Dict {
	c.Call(0x3111)
	putI(c, p, "pay_t_id")
	putS(c, p, "com_for_user")
	putS(c, p, "com_for_admin")
	c.Send()

	return nil
}
