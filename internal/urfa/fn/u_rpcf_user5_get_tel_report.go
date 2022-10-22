package fn

// rpcf_user5_get_tel_report
func xU4099(c conn, p Dict) Dict {
	c.Call(-0x4099)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4099" name="rpcf_user5_get_tel_report">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="accounts_size" />
      <for count="accounts_size" from="0" name="i">
        <integer name="dhs_log_size" />
        <set dst="dhs_log_size_array" dst_index="i" src="dhs_log_size" />
        <for count="dhs_log_size" from="0" name="j">
          <integer array_index="i,j" name="recv_date" />
          <integer array_index="i,j" name="recv_date_plus_acct_sess_time" />
          <integer array_index="i,j" name="acct_sess_time" />
          <integer array_index="i,j" name="setup_time" />
          <string array_index="i,j" name="Calling_Station_Id" />
          <string array_index="i,j" name="Called_Station_Id" />
          <string array_index="i,j" name="dname" />
          <double array_index="i,j" name="total_cost" />
        </for>
      </for>
    </output>
  </function>


*/
