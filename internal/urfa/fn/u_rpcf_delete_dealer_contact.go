package fn

// rpcf_delete_dealer_contact
func x13013(c conn, p Dict) Dict {
	c.Call(0x13013)
	putI(c, p, "contact_id")
	c.Send()

	return nil
}
