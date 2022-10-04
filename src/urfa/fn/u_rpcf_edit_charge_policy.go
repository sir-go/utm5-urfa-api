package fn

// rpcf_edit_charge_policy
func x15104(c conn, p Dict) Dict {
	c.Call(0x15104)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"result": c.GetI(),
	}
}

/* xml:
<function id="0x15104" name="rpcf_edit_charge_policy">
    <input>
      <integer name="policy_id" />
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
