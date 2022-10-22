package fn

// rpcf_get_sysgroup
func x2406(c conn, p Dict) Dict {
	c.Call(0x2406)
	putI(c, p, "group_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2406" name="rpcf_get_sysgroup">
    <input>
      <integer name="group_id" />
    </input>
    <output>
      <string name="group_name" />
      <string name="group_info" />
      <integer name="info_size" />
      <for count="info_size" from="0" name="i">
        <integer array_index="i" name="id" />
        <string array_index="i" name="name" />
        <string array_index="i" name="module" />
        <integer array_index="i" name="fids" />
      </for>
    </output>
  </function>


*/
