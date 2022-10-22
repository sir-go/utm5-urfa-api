package fn

// rpcf_dealer_get_tps_for_user
func x70000010(c conn, p Dict) Dict {
	c.Call(0x70000010)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "tariff_id")
	putI(c, p, "tplink")
	putI(c, p, "curr")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000010" name="rpcf_dealer_get_tps_for_user">
      <input>
	     <integer name="user_id" />
	     <integer name="account_id" />
	     <integer name="tariff_id" />
	     <integer name="tplink" />
	     <integer name="curr" />
      </input>
      <output>
	     <integer name="service_size" />
	     <if condition="eq" value="0" variable="service_size">
                 <error code="13" comment="tariff not found or you dont have enough privileges" />
             </if>
	     <for count="service_size" from="0" name="i">
	          <integer array_index="i" name="sid" />
		  <string array_index="i" name="service_name" />
		  <integer array_index="i" name="service_type" />
		  <string array_index="i" name="comment" />
		  <integer array_index="i" name="slink" />
		  <integer array_index="i" name="value" />
	     </for>
      </output>
    </function>




*/
