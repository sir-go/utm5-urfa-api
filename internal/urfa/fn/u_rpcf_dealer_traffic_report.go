package fn

// rpcf_dealer_traffic_report
func x7000003a(c conn, p Dict) Dict {
	c.Call(0x7000003a)

	// fixme: function in the default value
	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000003a" name="rpcf_dealer_traffic_report">
          <input>
               <integer name="type" />
           <integer default="0" name="user_id" />
           <integer default="0" name="account_id" />
           <integer default="0" name="group_id" />
           <integer default="0" name="apid" />
           <integer name="time_start" />
           <integer default="now()" name="time_end" />
          </input>
          <output>
             <double name="bytes_in_kbyte" />
             <integer name="users_count" />
             <for count="users_count" from="0" name="i">
               <integer name="atr_size" />
               <set dst="atr_size_array" dst_index="i" src="atr_size" />
               <for count="atr_size" from="0" name="j">
                  <integer array_index="i,j" name="account_id" />
                  <string array_index="i,j" name="login" />
              <if condition="eq" value="1" variable="type">
                  <integer array_index="i,j" name="discount_date" />
              </if>
              <if condition="eq" value="2" variable="type">
                  <integer array_index="i,j" name="discount_date" />
              </if>
              <if condition="eq" value="3" variable="type">
                  <integer array_index="i,j" name="discount_date" />
              </if>
              <if condition="eq" value="4" variable="type">
                  <integer array_index="i,j" name="ip_address" />
              </if>
                  <integer array_index="i,j" name="tclass" />
                  <double array_index="i,j" name="base_cost" />
                  <long array_index="i,j" name="bytes" />
                  <double array_index="i,j" name="discount" />
               </for>
             </for>
          </output>
        </function>


*/