package fn

// rpcf_get_supplier_id_for_service
func x8021(c conn, p Dict) Dict {
	c.Call(0x8021)
	putI(c, p, "service_id")
	c.Send()

	return Dict{
		"invoice_sup_id": c.GetI(),
	}
}
