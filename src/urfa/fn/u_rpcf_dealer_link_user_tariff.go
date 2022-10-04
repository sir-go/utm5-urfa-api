package fn

// rpcf_dealer_link_user_tariff
func x70000069(c conn, p Dict) Dict {
	c.Call(0x70000069)

	// fixme: default value parsing error
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000069" name="rpcf_dealer_link_user_tariff">
      <input>
        <integer name="user_id" />
        <integer default="0" name="account_id" />
        <integer name="tariff_current" />
        <integer default="tariff_current" name="tariff_next" />
        <integer name="discount_period_id" />
        <integer default="0" name="tariff_link_id" />
        <integer name="change_now" />
      </input>
      <output>
        <integer name="tariff_link_id" />
        <if condition="eq" value="0" variable="tariff_link_id">
          <error code="13" comment="unable to link user tariff may by you dont have enough privileges" />
        </if>
      </output>
    </function>


*/
