package fn

// rpcf_user5_service_report
func xU4011(c conn, p Dict) Dict {
	c.Call(-0x4011)
	putI(c, p, "time_start")
	putI(c, p, "time_end")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="-0x4011" name="rpcf_user5_service_report">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="aids_size" />
      <for count="aids_size" from="0" name="i">
        <integer name="asr_size" />
        <set dst="asr_size_array" dst_index="i" src="asr_size" />
        <for count="asr_size" from="0" name="j">
          <integer array_index="i,j" name="account_id" />
          <integer array_index="i,j" name="discount_date" />
          <double array_index="i,j" name="discount" />
          <double array_index="i,j" name="discount_with_tax" />
          <string array_index="i,j" name="service_name" />
          <integer array_index="i,j" name="service_type" />
          <string array_index="i,j" name="comment" />
        </for>
      </for>
    </output>
  </function>


*/
