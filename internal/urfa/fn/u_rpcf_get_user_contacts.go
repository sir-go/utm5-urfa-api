package fn

// rpcf_get_user_contacts
func x2021(c conn, p Dict) Dict {
	c.Call(0x2021)
	putI(c, p, "user_id")
	c.Send()

	return Dict{
		"contacts": getMapOf(c, func() Dict {
			return Dict{
				"id":           c.GetI(),
				"person":       c.GetS(),
				"descr":        c.GetS(),
				"contact":      c.GetS(),
				"email":        c.GetS(),
				"email_notify": c.GetI(),
				"short_name":   c.GetS(),
				"birthday":     c.GetS(),
				"id_exec_man":  c.GetI(),
			}
		}),
	}
}

/* xml:
<function id="0x2021" name="rpcf_get_user_contacts">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="person" />
        <string array_index="i" name="descr" />
        <string array_index="i" name="contact" />
        <string array_index="i" name="email" />
        <integer array_index="i" name="email_notify" />
        <string array_index="i" name="short_name" />
        <string array_index="i" name="birthday" />
        <integer array_index="i" name="id_exec_man" />
      </for>
    </output>
  </function>


*/
