package fn

// rpcf_set_supplier_id_for_service
func x8022(c conn, p Dict) Dict {
	c.Call(0x8022)
	putI(c, p, "service_id")
	putI(c, p, "invoice_sup_id")
	c.Send()

	return nil
}
