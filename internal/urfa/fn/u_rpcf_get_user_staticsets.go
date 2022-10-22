package fn

// rpcf_get_user_staticsets
func x9020(c conn, p Dict) Dict {
	c.Call(0x9020)
	putI(c, p, "user_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x9020" name="rpcf_get_user_staticsets">
    <input>
      <integer name="user_id" />
    </input>
    <output>
      <integer name="count_of_sets" />
      <if condition="ne" value="0" variable="count_of_sets">
        <integer name="num_of_type_of_sets" />
        <integer name="count" />
        <for count="count" from="0" name="i">
          <string array_index="i" name="router_info" />
          <integer array_index="i" name="currency_id" />
        </for>
      </if>
    </output>
  </function>


*/
