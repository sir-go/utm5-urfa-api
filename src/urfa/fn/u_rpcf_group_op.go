package fn

// rpcf_group_op
func x240d(c conn, p Dict) Dict {
	c.Call(0x240d)

	// fixme: input has a complex param
	panic(NotImplemented)
	return Dict{
		"error_code": c.GetI(),
		"error_msg":  c.GetS(),
	}
}

/* xml:
<function id="0x240d" name="rpcf_group_op">
    <input>
      <integer name="group_id" />
      <integer name="group_op" />
      <if condition="eq" value="4" variable="group_op">
        <integer name="tpid_from" />
        <integer name="tpid_to" />
      </if>
      <if condition="eq" value="5" variable="group_op">
        <integer name="old_policy" />
        <integer name="new_policy" />
      </if>
    </input>
    <output>
      <integer name="error_code" />
      <string name="error_msg" />
    </output>
  </function>


*/
