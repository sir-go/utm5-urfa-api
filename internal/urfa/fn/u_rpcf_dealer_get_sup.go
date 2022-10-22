package fn

// rpcf_dealer_get_sup
func x7000000f(c conn, _ Dict) Dict {
	c.Call(0x7000000f)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000000f" name="rpcf_dealer_get_sup">
      <input>
      </input>
      <output>
          <integer name="count" />
	  <for count="count" from="0" name="i">
	    <integer array_index="i" name="id" />
	    <string array_index="i" name="name" />
	    <string array_index="i" name="ur_adress" />
	    <string array_index="i" name="act_adress" />
	    <string array_index="i" name="inn" />
	    <string array_index="i" name="kpp" />
	    <integer array_index="i" name="bank_id" />
	    <string array_index="i" name="account" />
	    <string array_index="i" name="fio_headman" />
	    <string array_index="i" name="fio_bookeeper" />
	    <string array_index="i" name="fio_headman_sh" />
	    <string array_index="i" name="fio_bookeeper_sh" />
	    <string array_index="i" name="name_sh" />
	    <string array_index="i" name="bank_bic" />
	    <string array_index="i" name="bank_name" />
	    <string array_index="i" name="bank_city" />
	    <string array_index="i" name="bank_kschet" />
	  </for>
      </output>
    </function>


*/
