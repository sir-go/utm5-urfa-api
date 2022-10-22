package fn

// rpcf_add_time_range
func x10351(c conn, p Dict) Dict {
	c.Call(0x10351)

	// fixme: function in the default value
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x10351" name="rpcf_add_time_range">
    <input>
      <string name="tr_name" />
      <integer name="priority" />
      <integer default="size(sec_start)" name="size_trd" />
      <for count="size(sec_start)" from="0" name="i">
        <integer array_index="i" name="sec_start" />
        <integer array_index="i" name="sec_stop" />
        <integer array_index="i" name="min_start" />
        <integer array_index="i" name="min_stop" />
        <integer array_index="i" name="hour_start" />
        <integer array_index="i" name="hour_stop" />
        <integer array_index="i" name="wday_start" />
        <integer array_index="i" name="wday_stop" />
      </for>
      <integer default="size(internal_id)" name="days_size" />
      <for count="size(internal_id)" from="0" name="i">
        <integer array_index="i" name="internal_id" />
        <integer array_index="i" name="mday" />
        <integer array_index="i" name="month" />
      </for>
    </input>
    <output />
  </function>


*/
