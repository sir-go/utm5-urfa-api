package fn

// rpcf_dealer_get_user_tariffs
func x70000034(c conn, p Dict) Dict {
	c.Call(0x70000034)
	putI(c, p, "user_id")
	putI(c, p, "account_id", 0)
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000034" name="rpcf_dealer_get_user_tariffs">
      <input>
        <integer name="user_id" />
        <integer default="0" name="account_id" />
      </input>
      <output>
        <integer name="user_tariffs_size" />
        <if condition="eq" value="0" variable="user_tariffs_size">
            <error code="13" comment="user has no tariffs or you dont have enough privileges" />
        </if>
        <for count="user_tariffs_size" from="0" name="i">
          <integer array_index="i" name="tariff_current_array" />
          <integer array_index="i" name="tariff_next_array" />
          <integer array_index="i" name="discount_period_id_array" />
          <integer array_index="i" name="tariff_link_id_array" />
        </for>
      </output>
    </function>




*/
