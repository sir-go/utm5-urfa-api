package fn

// rpcf_user5_traffic_report
func xU4009(c conn, p Dict) Dict {
	c.Call(-0x4009)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4009" name="rpcf_user5_traffic_report">
    <input>
      <integer name="time_start" />
      <integer default="now()" name="time_end" />
    </input>
    <output>
      <integer name="unused" />
      <double name="bytes_in_kbyte" />
      <integer name="rows_count" />
      <for count="rows_count" from="0" name="i">
        <integer array_index="i" name="tclass" />
        <string array_index="i" name="tclass_name" />
        <long array_index="i" name="bytes" />
        <double array_index="i" name="base_cost" />
        <double array_index="i" name="discount" />
        <double array_index="i" name="discount_with_tax" />
      </for>
    </output>
  </function>


*/
