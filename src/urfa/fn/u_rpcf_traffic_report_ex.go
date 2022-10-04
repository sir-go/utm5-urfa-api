package fn

// rpcf_traffic_report_ex
func x3009(c conn, p Dict) Dict {
	c.Call(0x3009)
	putI(c, p, "type")
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "group_id")
	putI(c, p, "apid")
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x3009" name="rpcf_traffic_report_ex">
    <input>
      <integer name="type" />
      <integer name="user_id" />
      <integer name="account_id" />
      <integer name="group_id" />
      <integer name="apid" />
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <double name="bytes_in_kbyte" />
      <integer name="users_count" />
      <for count="users_count" from="0" name="i">
        <integer name="atr_size" />
        <set dst="atr_size_array" dst_index="i" src="atr_size" />
        <for count="atr_size" from="0" name="j">
          <integer array_index="i,j" name="account_id" />
          <string array_index="i,j" name="login" />
          <if condition="ne" value="0" variable="type">
            <integer array_index="i,j" name="add_param" />
	  </if>

          <integer array_index="i,j" name="tclass" />
          <double array_index="i,j" name="base_cost" />
          <long array_index="i,j" name="bytes" />
          <double array_index="i,j" name="discount" />
        </for>
      </for>
    </output>
  </function>


*/
