package fn

// rpcf_user5_edit_user
func xU4040(c conn, p Dict) Dict {
	c.Call(-0x4040)
	putS(c, p, "full_name")
	putS(c, p, "actual_address")
	putS(c, p, "juridical_address")
	putS(c, p, "work_telephone")
	putS(c, p, "home_telephone")
	putS(c, p, "mobile_telephone")
	putS(c, p, "web_page")
	putS(c, p, "icq_number")
	putS(c, p, "pasport")
	putI(c, p, "bank_id")
	putS(c, p, "bank_account")
	putS(c, p, "email")
	c.Send()

	return nil
}
