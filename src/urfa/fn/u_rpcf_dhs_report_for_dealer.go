package fn

// rpcf_dhs_report_for_dealer
func x13025(c conn, p Dict) Dict {
	c.Call(0x13025)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x13025" name="rpcf_dhs_report_for_dealer">
      <input>
	   <integer name="dealer_id" />
	   <integer default="0" name="user_id" />
	   <integer default="0" name="account_id" />
	   <integer default="0" name="gid" />
	   <integer default="0" name="apid" />
	   <integer name="time_start" />
	   <integer default="now()" name="time_end" />
      </input>
      <output>
          <integer name="accounts_count" />
	  <for count="accounts_count" from="0" name="j">
	     <integer name="dhs_log_size" />
          <set dst="dhs_log_size_array" dst_index="j" src="dhs_log_size" />
	     <for count="dhs_log_size" from="0" name="i">
	          <integer array_index="j,i" name="id" />
		  <integer array_index="j,i" name="account_id" />
		  <integer array_index="j,i" name="slink_id" />
		  <integer array_index="j,i" name="recv_date" />
		  <integer array_index="j,i" name="last_update_date" />
		  <string array_index="j,i" name="Called_Station_Id" />
		  <string array_index="j,i" name="Calling_Station_Id" />
		  <integer array_index="j,i" name="framed_ip" />
		  <integer array_index="j,i" name="nas_port" />
		  <string array_index="j,i" name="acct_session_id" />
		  <integer array_index="j,i" name="nas_port_type" />
		  <string array_index="j,i" name="uname" />
		  <integer array_index="j,i" name="service_type" />
		  <integer array_index="j,i" name="framed_protocol" />
		  <integer array_index="j,i" name="nas_ip" />
		  <string array_index="j,i" name="nas_id" />
		  <integer array_index="j,i" name="acct_status_type" />
		  <long array_index="j,i" name="acct_inp_pack" />
		  <long array_index="j,i" name="acct_inp_oct" />
		  <long array_index="j,i" name="acct_inp_giga" />
		  <long array_index="j,i" name="acct_out_pack" />
		  <long array_index="j,i" name="acct_out_oct" />
		  <long array_index="j,i" name="acct_out_giga" />
		  <long array_index="j,i" name="acct_sess_time" />
		  <integer array_index="j,i" name="acct_term_cause" />
		  <double array_index="j,i" name="total_cost" />
	     </for>
	  </for>
      </output>
    </function>


*/
