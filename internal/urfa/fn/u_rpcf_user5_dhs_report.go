package fn

// rpcf_user5_dhs_report
func xU4015(c conn, p Dict) Dict {
	c.Call(-0x4015)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4015" name="rpcf_user5_dhs_report">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="dhs_log_size" />
      <for count="dhs_log_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <integer array_index="i" name="account_id" />
        <integer array_index="i" name="slink_id" />
        <integer array_index="i" name="recv_date" />
        <integer array_index="i" name="last_update_date" />
        <integer array_index="i" name="framed_ip4" />
        <integer array_index="i" name="framed_ip6" />
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
        <long array_index="i" name="acct_out_pack" />
        <long array_index="i" name="acct_out_oct" />
        <long array_index="i" name="acct_sess_time" />
        <string array_index="i" name="calling_station_id" />
        <string array_index="i" name="called_station_id" />
        <integer name="dhs_sessions_detail_size" />
        <set dst="dhs_sessions_detail_size_array" dst_index="i" src="dhs_sessions_detail_size" />
        <for count="dhs_sessions_detail_size" from="0" name="j">
          <integer array_index="i,j" name="trange_id" />
          <integer array_index="i,j" name="account_id" />
          <long array_index="i,j" name="duration" />
          <double array_index="i,j" name="base_cost" />
          <double array_index="i,j" name="sum_cost" />
        </for>
      </for>
    </output>
  </function>




*/
