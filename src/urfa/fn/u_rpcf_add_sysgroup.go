package fn

// rpcf_add_sysgroup
func x2404(c conn, p Dict) Dict {
	c.Call(0x2404)

	// fixme: input has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2404" name="rpcf_add_sysgroup">
    <input>
      <integer name="group_id" />
      <string name="group_name" />
      <string name="group_info" />
      <integer name="num_of_fids" />
      <for count="num_of_fids" from="0" name="i">
        <integer array_index="i" name="fid" />
        <integer array_index="i" name="allowed" />
      </for>
    </input>
    <output />
  </function>


*/
