package fn

// rpcf_get_groups_list
func x2400(c conn, p Dict) Dict {
	c.Call(0x2400)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"groups": getMapOf(c, func() Dict {
			return Dict{
				"id":   c.GetI(),
				"name": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x2400" name="rpcf_get_groups_list">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <integer name="groups_count" />
      <for count="groups_count" from="0" name="i">
        <integer array_index="i" name="group_id" />
        <string array_index="i" name="group_name" />
      </for>
    </output>
  </function>


*/
