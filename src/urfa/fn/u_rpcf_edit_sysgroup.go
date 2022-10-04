package fn

// rpcf_edit_sysgroup
func x2405(c conn, p Dict) Dict {
	c.Call(0x2405)

	// fixme: function in the default value
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2405" name="rpcf_edit_sysgroup">
    <input>
      <integer name="group_id" />
      <string name="group_name" />
      <string name="group_info" />
      <integer default="size(fid)" name="num_of_fids" />
      <for count="size(fid)" from="0" name="i">
        <integer array_index="i" name="fid" />
        <integer array_index="i" name="allowed" />
      </for>
    </input>
    <output />
  </function>


*/
