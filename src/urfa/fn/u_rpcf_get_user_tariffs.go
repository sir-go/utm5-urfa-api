package fn

// rpcf_get_user_tariffs
func x3017(c conn, p Dict) Dict {
	c.Call(0x3017)
	putI(c, p, "user_id")
	putI(c, p, "account_id", 0)
	c.Send()

	return Dict{
		"tariffs": getMapOf(c, func() Dict {
			return Dict{
				"tariff_current":     c.GetI(),
				"tariff_next":        c.GetI(),
				"discount_period_id": c.GetI(),
				"tariff_link_id":     c.GetI(),
			}
		}),
	}
}

/* xml:
<function id="0x3017" name="rpcf_get_user_tariffs">
    <input>
      <integer name="user_id" />
      <integer default="0" name="account_id" />
    </input>
    <output>
      <integer name="user_tariffs_size" />
      <for count="user_tariffs_size" from="0" name="i">
        <integer array_index="i" name="tariff_current_array" />
        <integer array_index="i" name="tariff_next_array" />
        <integer array_index="i" name="discount_period_id_array" />
        <integer array_index="i" name="tariff_link_id_array" />
      </for>
    </output>
  </function>


*/
