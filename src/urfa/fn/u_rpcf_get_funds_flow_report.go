package fn

// rpcf_get_funds_flow_report
func x501c(c conn, p Dict) Dict {
	c.Call(0x501c)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	putI(c, p, "uid")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x501c" name="rpcf_get_funds_flow_report">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
      <integer name="uid" />
    </input>
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
	  <integer array_index="i" name="account_id_from" />
	  <integer array_index="i" name="account_id_to" />
	  <double array_index="i" name="amount" />
	  <integer array_index="i" name="date" />
	  <integer array_index="i" name="uid" />
	  <string array_index="i" name="login" />
	  <string array_index="i" name="full_name" />
     </for>
    </output>
  </function>


*/
