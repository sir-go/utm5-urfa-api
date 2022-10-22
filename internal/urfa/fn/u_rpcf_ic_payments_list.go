package fn

// rpcf_ic_payments_list
func x14005(c conn, p Dict) Dict {
	c.Call(0x14005)
	putI(c, p, "start_date")
	putI(c, p, "end_date")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x14005" name="rpcf_ic_payments_list">
    <input>
      <integer name="start_date" />
      <integer name="end_date" />
    </input>
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="id" />
	<string array_index="i" name="login" />
	<string array_index="i" name="fullname" />
	<integer array_index="i" name="actual_date" />
	<integer array_index="i" name="account_id" />
	<double array_index="i" name="sum" />
	<integer array_index="i" name="currency_id" />
	<integer array_index="i" name="ic_status" />
	<integer array_index="i" name="last_sync_date" />
	<string array_index="i" name="ic_id" />
      </for>
    </output>
  </function>


*/
