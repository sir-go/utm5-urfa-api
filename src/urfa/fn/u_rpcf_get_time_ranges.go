package fn

// rpcf_get_time_ranges
func x10350(c conn, _ Dict) Dict {
	c.Call(0x10350)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x10350" name="rpcf_get_time_ranges">
    <input />
    <output>
      <integer name="size_tr" />
      <for count="size_tr" from="0" name="i">
        <integer array_index="i" name="range_id" />
        <string array_index="i" name="tr_name" />
	<integer array_index="i" name="priority" />
        <integer name="size_trd" />
        <set dst="size_trd_array" dst_index="i" src="size_trd" />
        <for count="size_trd" from="0" name="j">
          <integer array_index="i,j" name="id" />
          <integer array_index="i,j" name="sec_start" />
          <integer array_index="i,j" name="sec_stop" />
          <integer array_index="i,j" name="min_start" />
          <integer array_index="i,j" name="min_stop" />
          <integer array_index="i,j" name="hour_start" />
          <integer array_index="i,j" name="hour_stop" />
          <integer array_index="i,j" name="wday_start" />
          <integer array_index="i,j" name="wday_stop" />
        </for>
        <integer name="days_size" />
        <set dst="days_size_array" dst_index="i" src="days_size" />
        <for count="days_size" from="0" name="j">
          <integer array_index="i,j" name="internal_id" />
          <integer array_index="i,j" name="mday" />
          <integer array_index="i,j" name="month" />
        </for>
      </for>
    </output>
  </function>


*/
