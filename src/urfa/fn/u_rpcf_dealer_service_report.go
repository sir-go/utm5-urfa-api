package fn

// rpcf_dealer_service_report
func x70000039(c conn, p Dict) Dict {
	c.Call(0x70000039)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000039" name="rpcf_dealer_service_report">
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
                <integer array_index="j,i" name="discount_date" />
            <integer array_index="j,i" name="discount_period_id" />
            <double array_index="j,i" name="discount" />
            <string array_index="j,i" name="service_name" />
            <integer array_index="j,i" name="service_type" />
            <string array_index="j,i" name="comment" />
             </for>
          </for>
          </output>
        </function>


*/
