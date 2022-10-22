package fn

// rpcf_user5_traffic_report_group
func xU4010(c conn, p Dict) Dict {
	c.Call(-0x4010)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	putI(c, p, "unused", 2)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4010" name="rpcf_user5_traffic_report_group">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
      <integer default="2" name="unused" />
    </input>
    <output>
      <integer name="unused" />
      <integer array_index="i" name="unused" />
      <double array_index="i" name="bytes_in_kbyte" />
      <integer name="date_count" />
      <for count="date_count" from="0" name="j">
        <integer array_index="j" name="date" />
        <integer name="rows_count" />
        <set dst="rows_count_array" dst_index="j" src="date_count" />
        <for count="rows_count" from="0" name="x">
          <integer array_index="j,x" name="tclass" />
          <string array_index="j,x" name="class_name" />
          <long array_index="j,x" name="bytes" />
          <double array_index="j,x" name="base_cost" />
          <double array_index="j,x" name="discount" />
        </for>
      </for>
    </output>
  </function>


*/
