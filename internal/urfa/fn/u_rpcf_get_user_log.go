package fn

// rpcf_get_user_log
func x2136(c conn, p Dict) Dict {
	c.Call(0x2136)
	putI(c, p, "user_id")
	putI(c, p, "tstart", TimeMin)
	putI(c, p, "tstop", TimeMax)
	putI(c, p, "group_id", 0)
	putI(c, p, "action", ActUnknown)
	c.Send()

	return Dict{
		"actions": getMapOf(c, func() Dict {
			return Dict{
				"user_id":   c.GetI(),
				"ud_login":  c.GetS(),
				"who":       c.GetI(),
				"usd_login": c.GetS(),
				"date":      c.GetI(),
				"action":    c.GetI(),
				"want":      c.GetS(),
				"comment":   c.GetI(),
			}
		}),
	}
}

/* xml:
<function id="0x2136" name="rpcf_get_user_log">
    <input>
      <integer name="user_id" />
      <integer name="tstart" />
      <integer name="tstop" />
      <integer name="group_id" />
      <integer name="action" />
    </input>
    <output>
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="user_id" />
        <string array_index="i" name="ud_login" />
        <integer array_index="i" name="who" />
        <string array_index="i" name="usd_login" />
        <integer array_index="i" name="date" />
        <integer array_index="i" name="action" />
        <string array_index="i" name="want" />
        <string array_index="i" name="comment" />
      </for>
    </output>
  </function>


*/
