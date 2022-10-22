package fn

// rpcf_tel_report_for_dealer
func x13028(c conn, p Dict) Dict {
	c.Call(0x13028)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x13028" name="rpcf_tel_report_for_dealer">
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
   <for count="accounts_count" from="0" name="k">
       <integer name="dhs_log_size" />
       <set dst="dhs_log_size_array" dst_index="k" src="dhs_log_size" />
       <for count="dhs_log_size" from="0" name="i">
         <integer name="count" />
           <set dst="count_array" dst_index="k,i" src="count" />
         <for count="count" from="0" name="j">
          <integer array_index="k,i,j" name="id" />
          <integer array_index="k,i,j" name="account_id" />
          <integer array_index="k,i,j" name="slink_id" />
          <integer array_index="k,i,j" name="recv_date" />
          <integer array_index="k,i,j" name="acct_sess_time_plus_recv_date" />
          <string array_index="k,i,j" name="Called_Station_Id" />
          <string array_index="k,i,j" name="Calling_Station_Id" />
          <integer array_index="k,i,j" name="nas_port" />
          <string array_index="k,i,j" name="acct_session_id" />
          <integer array_index="k,i,j" name="nas_port_type" />
          <string array_index="k,i,j" name="uname" />
          <integer array_index="k,i,j" name="service_type" />
          <integer array_index="k,i,j" name="framed_protocol" />
          <ip_address array_index="k,i,j" name="nas_ip" />
          <string array_index="k,i,j" name="nas_id" />
          <integer array_index="k,i,j" name="acct_status_type" />
          <long array_index="k,i,j" name="acct_inp_pack" />
          <long array_index="k,i,j" name="acct_inp_oct" />
          <long array_index="k,i,j" name="acct_out_pack" />
          <long array_index="k,i,j" name="acct_out_oct" />
          <integer array_index="k,i,j" name="zone_id" />
          <integer array_index="k,i,j" name="did" />
          <long array_index="k,i,j" name="acct_sess_time" />
          <string array_index="k,i,j" name="dcause" />
          <long array_index="k,i,j" name="duration" />
          <double array_index="k,i,j" name="base_cost" />
          <double array_index="k,i,j" name="sum_cost" />
         </for>
       </for>
   </for>
      </output>
    </function>




*/
