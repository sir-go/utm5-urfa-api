package fn

// rpcf_dealer_delete_service_link
func x70000020(c conn, p Dict) Dict {
	c.Call(0x70000020)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000020" name="rpcf_dealer_delete_service_link">
      <input>
        <integer name="slink_id" />
      </input>
      <output>
        <integer name="error_code" />
        <if condition="eq" value="0" variable="error_code">
          <error code="13" comment="service link not found or you dont have enough privileges" />
        </if>
      </output>
    </function>




*/
