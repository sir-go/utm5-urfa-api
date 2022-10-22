package fn

// rpcf_other_charges_report
func x3023(c conn, p Dict) Dict {
	c.Call(0x3023)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x3023" name="rpcf_other_charges_report">
      <input>
       <integer default="0" name="user_id" />
       <integer default="0" name="account_id" />
       <integer default="0" name="group_id" />
       <integer default="0" name="apid" />
       <integer name="time_start" />
       <integer default="now()" name="time_end" />
      </input>
      <output>
          <integer name="accounts_count" />
      <for count="accounts_count" from="0" name="j">
         <integer name="atr_size" />
          <set dst="atr_size_array" dst_index="j" src="atr_size" />
         <for count="atr_size" from="0" name="i">
            <integer array_index="j,i" name="account_id" />
            <string array_index="j,i" name="login" />
            <string array_index="j,i" name="full_name" />
            <integer array_index="j,i" name="discount_date" />
            <double array_index="j,i" name="discount" />
            <integer array_index="j,i" name="charge_type" />
         </for>
      </for>
      </output>
  </function>


*/
