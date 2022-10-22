package fn

// rpcf_get_discount_periods
func x2600(c conn, _ Dict) Dict {
	c.Call(0x2600)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2600" name="rpcf_get_discount_periods">
    <input />
    <output>
      <integer name="discount_periods_count" />
      <for count="discount_periods_count" from="0" name="i">
        <integer array_index="i" name="static_id" />
        <integer array_index="i" name="discount_period_id" />
        <integer array_index="i" name="start_date" />
        <integer array_index="i" name="end_date" />
        <integer array_index="i" name="periodic_type" />
        <integer array_index="i" name="custom_duration" />
        <integer array_index="i" name="next_discount_period_id" />
        <integer array_index="i" name="canonical_length" />
      </for>
    </output>
  </function>


*/
