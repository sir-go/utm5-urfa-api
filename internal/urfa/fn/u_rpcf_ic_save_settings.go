package fn

// rpcf_ic_save_settings
func x14002(c conn, p Dict) Dict {
	c.Call(0x14002)
	putI(c, p, "send_users")
	putI(c, p, "send_inv")
	putI(c, p, "send_payments")
	putI(c, p, "recv_payments")
	putI(c, p, "sync_card")
	putI(c, p, "sync_empty_name")
	putI(c, p, "sync_empty_inn")
	putI(c, p, "sync_empty_kpp")
	putI(c, p, "sync_users_since")
	putI(c, p, "sync_users_till")
	putI(c, p, "sync_inv_since")
	putI(c, p, "sync_inv_till")
	putI(c, p, "sync_payments_since")
	putI(c, p, "sync_payments_till")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
