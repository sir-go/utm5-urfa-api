package fn

// rpcf_get_tariffs_list
func x3024(c conn, _ Dict) Dict {
	c.Call(0x3024)

	return Dict{
		"tariffs": getMapOf(c, func() Dict {
			return Dict{
				"id":               c.GetI(),
				"name":             c.GetS(),
				"create_date":      c.GetI(),
				"who_create":       c.GetI(),
				"login":            c.GetS(),
				"change_create":    c.GetI(),
				"who_change":       c.GetI(),
				"login_change":     c.GetS(),
				"expire_date":      c.GetI(),
				"is_blocked":       c.GetI(),
				"balance_rollover": c.GetI(),
				"comment":          c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x3024" name="rpcf_get_tariffs_list">
    <input />
    <output>
      <integer name="tariffs_count" />
      <for count="tariffs_count" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="name" />
        <integer array_index="i" name="create_date" />
        <integer array_index="i" name="who_create" />
        <string array_index="i" name="login" />
        <integer array_index="i" name="change_create" />
        <integer array_index="i" name="who_change" />
        <string array_index="i" name="login_change" />
        <integer array_index="i" name="expire_date" />
        <integer array_index="i" name="is_blocked" />
        <integer array_index="i" name="balance_rollover" />
        <string array_index="i" name="comment" />
      </for>
    </output>
  </function>


*/
