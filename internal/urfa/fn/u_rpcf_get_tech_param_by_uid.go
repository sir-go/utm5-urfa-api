package fn

// rpcf_get_tech_param_by_uid
func x9000(c conn, p Dict) Dict {
	c.Call(0x9000)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x9000" name="rpcf_get_tech_param_by_uid">
    <input>
      <integer name="user_id" />
      <integer name="account_id" />
    </input>
    <output>
      <integer name="size_tp" />
      <for count="size_tp" from="0" name="i">
        <integer name="size_vec_ltp" />
        <set dst="size_vec_ltp_array" dst_index="i" src="size_vec_ltp" />
        <for count="size_vec_ltp" from="0" name="j">
          <integer array_index="i,j" name="id" />
          <integer array_index="i,j" name="type_id" />
          <string array_index="i,j" name="type_name" />
          <string array_index="i,j" name="param" />
          <integer array_index="i,j" name="reg_date" />
          <integer array_index="i,j" name="slink_id" />
          <string array_index="i,j" name="service_name" />
          <string array_index="i,j" name="password" />
        </for>
      </for>
    </output>
  </function>


*/
