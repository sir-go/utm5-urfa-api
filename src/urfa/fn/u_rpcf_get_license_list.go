package fn

// rpcf_get_license_list
func x2122(c conn, _ Dict) Dict {
	c.Call(0x2122)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x2122" name="rpcf_get_license_list">
      <input />
      <output>
         <integer name="lic_count" />
	 <for count="lic_count" from="0" name="i">
	     <integer array_index="i" name="license_id" />
	     <integer array_index="i" name="license_since" />
	     <integer array_index="i" name="license_till" />
	     <integer name="features_count" />
	     <for count="features_count" from="0" name="j">
	         <integer array_index="i,j" name="feature_id" />
		 <string array_index="i,j" name="feature_value" />
	     </for>
	 </for>
      </output>
  </function>


*/
