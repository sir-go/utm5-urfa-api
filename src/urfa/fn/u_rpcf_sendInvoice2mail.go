package fn

// rpcf_sendInvoice2mail
func x8004(c conn, p Dict) Dict {
	c.Call(0x8004)
	putI(c, p, "aid")
	putI(c, p, "invc_id")
	c.Send()

	return nil
}
