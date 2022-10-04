package fn

// rpcf_dealer_get_periodic_service
func x70000027(c conn, p Dict) Dict {
	c.Call(0x70000027)
	putI(c, p, "service_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000027" name="rpcf_dealer_get_periodic_service">
      <input>
         <integer name="service_id" />
      </input>
      <output>
          <integer name="sid" />
          <if condition="eq" value="0" variable="sid">
             <error code="11" comment="service not found or you dont have enough privileges" />
          </if>
          <string name="service_name" />
          <string name="comment" />
	  <integer name="link_by_default" />
	  <double name="cost" />
	  <integer name="deprecated" />
	  <integer name="discount_method" />
	  <integer name="start_date" />
	  <integer name="expire_date" />
          <integer name="param" />
	  <integer name="tariff_id" />
	  <integer name="parent_id" />
      </output>
    </function>


*/
