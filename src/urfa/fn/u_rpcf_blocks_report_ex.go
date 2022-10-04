package fn

// rpcf_blocks_report_ex
func x300b(c conn, p Dict) Dict {
	c.Call(0x300b)

	putI(c, p, "user_id", 0)
	putI(c, p, "account_id", 0)
	putI(c, p, "group_id", 0)
	putI(c, p, "apid", 0)
	putI(c, p, "time_start", TimeThisMonthBegin())
	putI(c, p, "time_end", TimeNow())
	putI(c, p, "show_all", 1)
	c.Send()

	rows := make([]Dict, 0)
	accountsCount := c.GetI()
	if accountsCount < 1 {
		return Dict{"rows": rows}
	}

	for i := 0; i < accountsCount; i++ {
		atrCount := c.GetI()
		for j := 0; j < atrCount; j++ {
			rows = append(rows, Dict{
				"account_id":  c.GetI(),
				"login":       c.GetS(),
				"start_date":  c.GetI(),
				"expire_date": c.GetI(),
				"block_type":  c.GetI(),
				"id":          c.GetI(),
				"unabon":      c.GetI(),
				"unprepay":    c.GetI(),
				"is_deleted":  c.GetI(),
			})
		}
	}

	return Dict{"rows": rows}
}

/* xml:
<function id="0x300b" name="rpcf_blocks_report_ex">
      <input>
       <integer default="0" name="user_id" />
       <integer default="0" name="account_id" />
       <integer default="0" name="group_id" />
       <integer default="0" name="apid" />
       <integer name="time_start" />
       <integer default="now()" name="time_end" />
       <integer default="1" name="show_all" />
      </input>
      <output>
       <integer name="accounts_count" />
       <for count="accounts_count" from="0" name="i">
        <integer name="atr_size" />
           <set dst="atr_size_array" dst_index="i" src="atr_size" />
        <for count="atr_size" from="0" name="j">
          <integer array_index="i,j" name="account_id" />
          <string array_index="i,j" name="login" />
          <integer array_index="i,j" name="start_date" />
          <integer array_index="i,j" name="expire_date" />
          <integer array_index="i,j" name="block_type" />
          <integer array_index="i,j" name="id" />
          <integer array_index="i,j" name="unabon" />
          <integer array_index="i,j" name="unprepay" />
          <integer array_index="i,j" name="is_deleted" />
        </for>
       </for>
      </output>
  </function>


*/
