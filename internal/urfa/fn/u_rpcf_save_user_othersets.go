package fn

// rpcf_save_user_othersets
func x9022(c conn, p Dict) Dict {
	c.Call(0x9022)

	// fixme: input has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x9022" name="rpcf_save_user_othersets">
    <input>
      <integer name="user_id" />
      <integer name="count" />
      <for count="count" from="0" name="i">
        <integer array_index="i" name="type" />
        <if condition="eq" value="1" variable="type">
          <integer name="id" />
          <integer name="port" />
        </if>
        <if condition="eq" value="3" variable="type">
          <integer name="currency_id" />
        </if>
      </for>
    </input>
    <output />
  </function>


*/
