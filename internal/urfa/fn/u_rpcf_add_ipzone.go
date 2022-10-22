package fn

// rpcf_add_ipzone
func x2801(c conn, p Dict) Dict {
	c.Call(0x2801)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"id": c.GetI(),
	}
}

/* xml:
<function id="0x2801" name="rpcf_add_ipzone">
    <input>
      <integer name="id" />
      <string name="name" />
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="net" />
        <integer array_index="i" name="mask" />
        <integer array_index="i" name="gateaway" />
      </for>
    </input>
    <output>
      <integer name="id" />
    </output>
  </function>


*/
