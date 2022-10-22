package fn

// rpcf_payments_report_new
func x3030(c conn, p Dict) Dict {
	c.Call(0x3030)
	putI(c, p, "user_id", 0)
	putI(c, p, "account_id", 0)
	putI(c, p, "group_id", 0)
	putI(c, p, "apid", 0)
	putI(c, p, "time_start", TimeThisMonthBegin())
	putI(c, p, "time_end", TimeNow())
	c.Send()

	rows := make([]Dict, 0)
	usersCount := c.GetI()
	if usersCount < 1 {
		return Dict{"rows": rows}
	}

	for i := 0; i < usersCount; i++ {
		atrCount := c.GetI()
		for j := 0; j < atrCount; j++ {
			rows = append(rows, Dict{
				"id":                 c.GetI(),
				"account_id":         c.GetI(),
				"login":              c.GetS(),
				"actual_date":        c.GetI(),
				"payment_enter_date": c.GetI(),
				"payment":            c.GetD(),
				"payment_incurrency": c.GetD(),
				"currency_id":        c.GetI(),
				"method":             c.GetI(),
				"who_receved":        c.GetI(),
				"admin_comment":      c.GetS(),
				"payment_ext_number": c.GetS(),
				"full_name":          c.GetS(),
				"acc_external_id":    c.GetS(),
				"burnt_date":         c.GetI(),
			})
		}
	}

	return Dict{"rows": rows}
}

/* xml:
<function id="0x3030" name="rpcf_payments_report_new">
    <input>
      <integer default="0" name="user_id" />
      <integer name="account_id" />
      <integer default="0" name="group_id" />
      <integer default="0" name="apid" />
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="users_count" />
      <for count="users_count" from="0" name="i">
        <integer name="atr_size" />
        <set dst="atr_size_array" dst_index="i" src="atr_size" />
        <for count="atr_size" from="0" name="j">
          <integer array_index="i,j" name="id" />
          <integer array_index="i,j" name="account_id" />
          <string array_index="i,j" name="login" />
          <integer array_index="i,j" name="actual_date" />
          <integer array_index="i,j" name="payment_enter_date" />
          <double array_index="i,j" name="payment" />
          <double array_index="i,j" name="payment_incurrency" />
          <integer array_index="i,j" name="currency_id" />
          <integer array_index="i,j" name="method" />
          <integer array_index="i,j" name="who_receved" />
          <string array_index="i,j" name="admin_comment" />
          <string array_index="i,j" name="payment_ext_number" />
          <string array_index="i,j" name="full_name" />
          <string array_index="i,j" name="acc_external_id" />
          <integer array_index="i,j" name="burnt_date" />
        </for>
      </for>
    </output>
  </function>



*/
