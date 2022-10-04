package fn

// rpcf_dealer_get_tariff_services_for_user
func x70000049(c conn, p Dict) Dict {
	c.Call(0x70000049)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	putI(c, p, "tplink_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000049" name="rpcf_dealer_get_tariff_services_for_user">
      <input>
          <integer name="user_id" />
	  <integer name="account_id" />
	  <integer name="tplink_id" />
      </input>
      <output>
          <integer name="services_array_size" />
	  <for count="services_array_size" from="0" name="i">
	      <integer array_index="i" name="service_id_array" />
	      <string array_index="i" name="service_name_array" />
	      <integer array_index="i" name="service_type_array" />
	      <string array_index="i" name="service_comment_array" />
	      <integer array_index="i" name="slink_id_array" />
	      <integer array_index="i" name="is_linked_array" />
	  </for>
      </output>
    </function>


*/
