package fn

// rpcf_put_unif_iptr
func x5511(c conn, p Dict) Dict {
	c.Call(0x5511)

	// fixme: function in the default value
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x5511" name="rpcf_put_unif_iptr">
    <input>
      <integer default="size(login)" name="count" />
      <for count="size(login)" from="0" name="i">
        <string array_index="i" name="login" />
        <integer array_index="i" name="ipid" />
        <integer array_index="i" name="tclass" />
        <integer array_index="i" name="d_oct" />
      </for>
    </input>
    <output />
  </function>


*/
