package fn

// rpcf_user5_traffic_report_group_by_ip
func xU404c(c conn, p Dict) Dict {
	c.Call(-0x404c)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x404c" name="rpcf_user5_traffic_report_group_by_ip">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="unused" />
      <integer array_index="i" name="unused" />
      <double array_index="i" name="bytes_in_kbyte" />
      <integer name="ipid_count" />
      <for count="ipid_count" from="0" name="j">
        <ip_address array_index="j" name="ipid" />
        <integer name="traffic_report_entry_size" />
        <set dst="traffic_report_entry_size_array" dst_index="j" src="traffic_report_entry_size" />
        <for count="traffic_report_entry_size" from="0" name="x">
          <integer array_index="j,x" name="tclass" />
          <string array_index="j,x" name="tclass_name" />
          <long array_index="j,x" name="bytes" />
        </for>
      </for>
    </output>
  </function>


*/
