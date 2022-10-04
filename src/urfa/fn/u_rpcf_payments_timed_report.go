package fn

// rpcf_payments_timed_report
func x3006(c conn, p Dict) Dict {
	c.Call(0x3006)
	putI(c, p, "user_id", 0)
	putI(c, p, "account_id")
	putI(c, p, "group_id", 0)
	putI(c, p, "apid")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x3006" name="rpcf_payments_timed_report">
    <input>
      <integer default="0" name="user_id" />
      <integer name="account_id" />
      <integer default="0" name="group_id" />
      <integer name="apid" />
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="accounts_size" />
      <for count="accounts_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <if condition="ne" value="-1" variable="id">
          <string array_index="i" name="login" />
          <integer array_index="i" name="first_date" />
          <integer array_index="i" name="last_date" />
          <integer array_index="i" name="burn_date" />
          <double array_index="i" name="amount" />
          <double array_index="i" name="already_discounted" />
        </if>
      </for>
    </output>
  </function>


*/
