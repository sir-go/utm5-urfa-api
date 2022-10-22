package fn

// rpcf_user5_get_funds_flow_report
func xU15009(c conn, p Dict) Dict {
	c.Call(-0x15009)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x15009" name="rpcf_user5_get_funds_flow_report">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="count" />
      <for count="count" from="0" name="i">
	  <integer array_index="i" name="account_id_from" />
	  <integer array_index="i" name="account_id_to" />
	  <double array_index="i" name="amount" />
	  <integer array_index="i" name="date" />
     </for>
    </output>
  </function>


*/
