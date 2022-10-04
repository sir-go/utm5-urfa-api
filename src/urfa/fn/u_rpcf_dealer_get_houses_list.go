package fn

// rpcf_dealer_get_houses_list
func x7000000d(c conn, _ Dict) Dict {
	c.Call(0x7000000d)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000000d" name="rpcf_dealer_get_houses_list">
      <input>
      </input>
      <output>
          <integer name="houses_size" />
          <for count="houses_size" from="0" name="i">
	     <integer array_index="i" name="house_id" />
	     <integer array_index="i" name="ip_zone_id" />
	     <integer array_index="i" name="connect_date" />
	     <string array_index="i" name="post_code" />
	     <string array_index="i" name="country" />
	     <string array_index="i" name="region" />
	     <string array_index="i" name="city" />
	     <string array_index="i" name="street" />
	     <string array_index="i" name="number" />
	     <string array_index="i" name="building" />
          </for>
      </output>
    </function>


*/
