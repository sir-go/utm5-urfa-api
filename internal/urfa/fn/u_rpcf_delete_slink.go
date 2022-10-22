package fn

// rpcf_delete_slink
func x5100(c conn, p Dict) Dict {
	c.Call(0x5100)
	putI(c, p, "slink_id")
	c.Send()

	if c.GetI() != 0 {
		return Dict{"error": Dict{"code": 13, "comment": "unable to delete service link"}}
	}
	return nil
}

/* xml:
<function id="0x5100" name="rpcf_delete_slink">
    <input>
      <integer name="slink_id" />
    </input>
    <output>
      <integer name="error_code" />
      <if condition="ne" value="0" variable="error_code">
        <error code="13" comment="unable to delete service link" />
      </if>
    </output>
  </function>


*/
