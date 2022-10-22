package fn

// rpcf_dealer_get_telephony_service_link
func x70000065(c conn, p Dict) Dict {
	c.Call(0x70000065)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000065" name="rpcf_dealer_get_telephony_service_link">
      <input>
        <integer name="slink_id" />
      </input>
      <output>
         <integer name="sid" />
          <if condition="eq" value="0" variable="sid">
             <error code="11" comment="service link not found or you dont have enough privileges" />
          </if>
        <integer name="tariff_link_id" />
        <integer name="is_blocked" />
        <integer name="discount_period_id" />
        <integer name="start_date" />
        <integer name="expire_date" />
        <integer name="policy_id" />
         <double name="cost_coef" />
        <integer name="numbers_size" />
        <for count="numbers_size" from="0" name="numbers">
          <integer array_index="numbers" name="item_id" />
          <string array_index="numbers" name="number" />
          <string array_index="numbers" name="login" />
          <string array_index="numbers" name="password" />
          <string array_index="numbers" name="allowed_cid" />
        </for>
      </output>
    </function>


*/
