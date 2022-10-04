package fn

// rpcf_get_dhs_report
func x5019(c conn, p Dict) Dict {
	c.Call(0x5019)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "apid")
	putI(c, p, "t_start")
	putI(c, p, "t_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5019" name="rpcf_get_dhs_report">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
      <integer name="apid" />
      <integer name="t_start" />
      <integer name="t_end" />
    </input>
    <output>
      <integer name="dhs_log_size" />
      <for count="dhs_log_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="account_id" />
        <integer array_index="i" name="slink_id" />
        <integer array_index="i" name="recv_date" />
        <integer array_index="i" name="last_update_date" />
        <string array_index="i" name="Called_Station_Id" />
        <string array_index="i" name="Calling_Station_Id" />
        <ip_address array_index="i" name="framed_ip4" />
        <ip_address array_index="i" name="framed_ip6" />
        <integer array_index="i" name="nas_port" />
        <string array_index="i" name="acct_session_id" />
        <integer array_index="i" name="nas_port_type" />
        <string array_index="i" name="uname" />
        <integer array_index="i" name="service_type" />
        <integer array_index="i" name="framed_protocol" />
        <integer array_index="i" name="nas_ip" />
        <string array_index="i" name="nas_id" />
        <integer array_index="i" name="acct_status_type" />
        <long array_index="i" name="acct_inp_pack" />
        <long array_index="i" name="acct_inp_oct" />
        <long array_index="i" name="acct_inp_giga" />
        <long array_index="i" name="acct_out_pack" />
        <long array_index="i" name="acct_out_oct" />
        <long array_index="i" name="acct_out_giga" />
        <long array_index="i" name="acct_sess_time" />
        <integer array_index="i" name="acct_term_cause" />
        <double array_index="i" name="total_cost" />
      </for>
    </output>
  </function>


*/
