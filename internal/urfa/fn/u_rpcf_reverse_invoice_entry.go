package fn

// rpcf_reverse_invoice_entry
func x7041(c conn, p Dict) Dict {
	c.Call(0x7041)
	putI(c, p, "invoice_entry_id")
	putD(c, p, "new_summ")
	putS(c, p, "comment")
	c.Send()

	return Dict{
		"result": c.GetI(),
	}
}
