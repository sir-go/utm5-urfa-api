package fn

// rpcf_get_radius_accounts_list
func x10154(c conn, _ Dict) Dict {
	c.Call(0x10154)

	return Dict{
		"accounts": getMapOf(c, func() Dict {
			return Dict{
				"id":            c.GetI(),
				"name":          c.GetS(),
				"password":      c.GetS(),
				"add_nas_attrs": c.GetI(),
			}
		}),
	}
}

/* xml:
<function id="0x10154" name="rpcf_get_radius_accounts_list">
      <input />
      <output>
        <integer name="size" />
        <for count="size" from="0" name="i">
          <integer array_index="i" name="id" />
          <string array_index="i" name="name" />
          <string array_index="i" name="password" />
          <integer array_index="i" name="add_nas_attrs" />
        </for>
      </output>
    </function>


*/
