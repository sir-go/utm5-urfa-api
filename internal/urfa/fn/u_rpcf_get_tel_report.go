package fn

// rpcf_get_tel_report
func x5037(c conn, p Dict) Dict {
	c.Call(0x5037)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "apid")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5037" name="rpcf_get_tel_report">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
      <integer name="apid" />
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="dhs_log_size" />
      <for count="dhs_log_size" from="0" name="i">
        <integer name="count" />
        <set dst="count_array" dst_index="i" src="count" />
        <for count="count" from="0" name="j">
          <integer array_index="i,j" name="id" />
          <integer array_index="i,j" name="account_id" />
          <integer array_index="i,j" name="slink_id" />
          <integer array_index="i,j" name="recv_date" />
          <integer array_index="i,j" name="start_time" />
          <integer array_index="i,j" name="end_time" />
          <string array_index="i,j" name="Called_Station_Id" />
          <string array_index="i,j" name="Calling_Station_Id" />
          <integer array_index="i,j" name="nas_port" />
          <string array_index="i,j" name="acct_session_id" />
          <integer array_index="i,j" name="nas_port_type" />
          <string array_index="i,j" name="uname" />
          <integer array_index="i,j" name="service_type" />
          <integer array_index="i,j" name="framed_protocol" />
          <ip_address array_index="i,j" name="nas_ip" />
          <string array_index="i,j" name="nas_id" />
          <integer array_index="i,j" name="acct_status_type" />
          <long array_index="i,j" name="acct_inp_pack" />
          <long array_index="i,j" name="acct_inp_oct" />
          <long array_index="i,j" name="acct_out_pack" />
          <long array_index="i,j" name="acct_out_oct" />
          <integer array_index="i,j" name="zone_id" />
          <integer array_index="i,j" name="did" />
          <long array_index="i,j" name="acct_sess_time" />
          <string array_index="i,j" name="incoming_trunk" />
          <string array_index="i,j" name="outgoing_trunk" />
          <string array_index="i,j" name="pbx_id" />
          <integer array_index="i,j" name="flags" />
          <string array_index="i,j" name="dcause" />
          <long array_index="i,j" name="duration" />
          <double array_index="i,j" name="base_cost" />
          <double array_index="i,j" name="cost_mult" />
          <double array_index="i,j" name="sum_cost" />
        </for>
      </for>
    </output>
  </function>


*/
