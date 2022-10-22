package fn

// rpcf_get_user_contacts_new
func x2040(c conn, p Dict) Dict {
	c.Call(0x2040)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2040" name="rpcf_get_user_contacts_new">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <integer name="size" />
      <for count="size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="descr" />
        <string array_index="i" name="reason" />
        <string array_index="i" name="person" />
        <string array_index="i" name="short_name" />
        <string array_index="i" name="contact" />
        <string array_index="i" name="email" />
        <integer array_index="i" name="id_exec_man" />
      </for>
    </output>
  </function>


*/
