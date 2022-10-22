package fn

// rpcf_get_all_discount_periods
func x2607(c conn, _ Dict) Dict {
	c.Call(0x2607)

	return Dict{
		"discount_periods": getMapOf(c, func() Dict {
			return Dict{
				"static_id":               c.GetI(),
				"discount_period_id":      c.GetI(),
				"start_date":              c.GetI(),
				"end_date":                c.GetI(),
				"periodic_type":           c.GetI(),
				"custom_duration":         c.GetI(),
				"next_discount_period_id": c.GetI(),
				"canonical_length":        c.GetI(),
			}
		}),
	}
}

/* xml:
<function id="0x2607" name="rpcf_get_all_discount_periods">
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
