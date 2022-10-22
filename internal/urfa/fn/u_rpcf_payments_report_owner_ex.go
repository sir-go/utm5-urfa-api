package fn

// rpcf_payments_report_owner_ex
func x300a(c conn, p Dict) Dict {
	c.Call(0x300a)
	putI(c, p, "time_start", TimeTodayMidnight())
	putI(c, p, "time_start", TimeNow())
	c.Send()

	return Dict{
		"rows": getMapOf(c, func() Dict {
			return Dict{
				"id":                 c.GetI(),
				"account_id":         c.GetI(),
				"login":              c.GetS(),
				"user_id":            c.GetI(),
				"full_name":          c.GetS(),
				"actual_date":        c.GetI(),
				"payment_enter_date": c.GetI(),
				"payment":            c.GetD(),
				"payment_incurrency": c.GetD(),
				"currency_id":        c.GetI(),
				"method":             c.GetI(),
				"who_receved":        c.GetI(),
				"admin_comment":      c.GetS(),
				"payment_ext_number": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x300a" name="rpcf_payments_report_owner_ex">
    <input>
      <integer default="0" name="time_start" />
      <integer default="now()" name="time_end" />
    </input>
    <output>
      <integer name="rows_count" />
      <for count="rows_count" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="account_id" />
        <string array_index="i" name="login" />
        <integer array_index="i" name="user_id" />
        <string array_index="i" name="full_name" />
        <integer array_index="i" name="actual_date" />
        <integer array_index="i" name="payment_enter_date" />
        <double array_index="i" name="payment" />
        <double array_index="i" name="payment_incurrency" />
        <integer array_index="i" name="currency_id" />
        <integer array_index="i" name="method" />
        <integer array_index="i" name="who_receved" />
        <string array_index="i" name="admin_comment" />
        <string array_index="i" name="payment_ext_number" />
      </for>
    </output>
  </function>


*/
