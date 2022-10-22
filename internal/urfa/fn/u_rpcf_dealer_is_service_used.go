package fn

// rpcf_dealer_is_service_used
func x7000002c(c conn, p Dict) Dict {
	c.Call(0x7000002c)
	putI(c, p, "service_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000002c" name="rpcf_dealer_is_service_used">
      <input>
          <integer name="service_id" />
      </input>
      <output>
          <integer name="links_count" />
          <if condition="eq" value="0" variable="links_count">
             <error code="11" comment="service not found or not used or you dont have enough privileges" />
          </if>
      </output>
    </function>


*/
