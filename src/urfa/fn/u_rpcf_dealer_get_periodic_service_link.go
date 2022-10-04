package fn

// rpcf_dealer_get_periodic_service_link
func x70000064(c conn, p Dict) Dict {
	c.Call(0x70000064)
	putI(c, p, "slink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000064" name="rpcf_dealer_get_periodic_service_link">
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
	 <integer name="is_unabon_period" />
	 <integer name="is_unprepay_period" />
	 <integer name="tariff_id" />
	 <integer name="parent_id" />
      </output>
    </function>


*/
