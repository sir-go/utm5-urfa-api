package fn

// rpcf_save_account
func x1510b(c conn, p Dict) Dict {
	c.Call(0x1510b)
	putI(c, p, "account_id")
	putD(c, p, "credit", 0.0)
	if putI(c, p, "is_blocked", 0) != 0 {
		putI(c, p, "block_start_date", TimeNow())
		putI(c, p, "block_end_date", TimeMax)
	}
	putD(c, p, "vat_rate", 0.0)
	putD(c, p, "sale_tax_rate", 0.0)
	putI(c, p, "int_status", 1)
	putI(c, p, "unlimited", 0)
	putI(c, p, "auto_enable_inet", 1)
	putS(c, p, "external_id", "")
	c.Send()

	return nil
}

/* xml:
<function id="0x1510b" name="rpcf_save_account">
    <input>
      <integer name="account_id" />
      <double name="credit" />
      <integer name="is_blocked" />
      <if condition="ne" value="0" variable="is_blocked">
        <integer default="now()" name="block_start_date" />
        <integer default="max_time()" name="block_end_date" />
      </if>
      <double name="vat_rate" />
      <double name="sale_tax_rate" />
      <integer name="int_status" />
      <integer name="unlimited" />
      <integer name="auto_enable_inet" />
      <string name="external_id" />
    </input>
    <output />
  </function>


*/
