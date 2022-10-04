package fn

// rpcf_add_charge_policy
func x15103(c conn, p Dict) Dict {
	c.Call(0x15103)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x15103" name="rpcf_add_charge_policy">
    <input>
      <integer name="flags" />
      <string name="name" />
      <integer name="tm_count" />
      <for count="tm_count" from="0" name="i">
        <integer array_index="i" name="timemark" />
      </for>
    </input>
    <output>
      <integer name="result" />
    </output>
  </function>


*/
