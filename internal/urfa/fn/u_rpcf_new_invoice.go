package fn

// rpcf_new_invoice
func x10205(c conn, p Dict) Dict {
	c.Call(0x10205)

	// fixme: function in the default value
	panic(NotImplemented)
	return Dict{
		"id": c.GetI(),
	}
}

/* xml:
<function id="0x10205" name="rpcf_new_invoice">
    <input>
      <integer name="account_id" />
      <integer name="gen_date" />
      <integer name="ap_id" />
      <integer default="size(name)" name="count" />
      <for count="size(name)" from="0" name="i">
        <string array_index="i" name="name" />
        <double array_index="i" name="qnt" />
        <double array_index="i" name="base_cost" />
        <double array_index="i" name="sum_cost" />
      </for>
    </input>
    <output>
      <integer name="id" />
    </output>
  </function>


*/
