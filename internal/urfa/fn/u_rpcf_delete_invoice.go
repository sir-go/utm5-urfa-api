package fn

// rpcf_delete_invoice
func x7042(c conn, p Dict) Dict {
	c.Call(0x7042)
	putI(c, p, "invoice_id")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
