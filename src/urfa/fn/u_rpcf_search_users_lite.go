package fn

// rpcf_search_users_lite
func x1202(c conn, p Dict) Dict {
	c.Call(0x1202)
	putS(c, p, "login", "")
	putS(c, p, "email", "")
	putS(c, p, "fname", "")
	c.Send()

	return Dict{
		"success": c.GetI(),
		"total":   c.GetI(),
		"users": getMapOf(c, func() Dict {
			return Dict{
				"id":    c.GetI(),
				"login": c.GetS(),
				"email": c.GetS(),
				"name":  c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x1202" name="rpcf_search_users_lite">
    <input>
      <string name="login" />
      <string name="email" />
      <string name="fname" />
    </input>
    <output>
      <integer name="success" />
      <integer name="total" />
      <integer name="show_count" />
      <if condition="ne" value="0" variable="show_count">
        <for count="show_count" from="0" name="i">
          <integer array_index="i" name="id" />
          <string array_index="i" name="login" />
          <string array_index="i" name="email" />
          <string array_index="i" name="name" />
        </for>
      </if>
    </output>
  </function>



*/
