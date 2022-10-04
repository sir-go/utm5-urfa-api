package fn

// rpcf_user5_blocks_report
func xU4013(c conn, p Dict) Dict {
	c.Call(-0x4013)
	putI(c, p, "time_start", TimeThisMonthBegin())
	putI(c, p, "time_end", TimeNow())
	c.Send()

	return Dict{
		"blocks": getMapOf(c, func() Dict {
			return Dict{
				"account_id":   c.GetI(),
				"start_date":   c.GetI(),
				"expire_date":  c.GetI(),
				"what_blocked": c.GetI(),
				"block_type":   c.GetI(),
				"comment":      c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="-0x4013" name="rpcf_user5_blocks_report">
    <input>
      <integer name="time_start" />
      <integer name="time_end" />
    </input>
    <output>
      <integer name="blocks_size" />
      <for count="blocks_size" from="0" name="i">
        <integer array_index="i" name="account_id" />
        <integer array_index="i" name="start_date" />
        <integer array_index="i" name="expire_date" />
        <integer array_index="i" name="what_blocked" />
        <integer array_index="i" name="block_type" />
        <string array_index="i" name="comment" />
      </for>
    </output>
  </function>


*/
