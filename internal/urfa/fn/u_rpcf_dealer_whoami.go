package fn

// rpcf_dealer_whoami
func x7000004a(c conn, _ Dict) Dict {
	c.Call(0x7000004a)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x7000004a" name="rpcf_dealer_whoami">
      <input>
      </input>
      <output>
          <integer name="my_uid" />
	  <string name="login" />
	  <ip_address name="user_ip" />
	  <ip_address name="user_mask" />
	  <integer name="system_group_size" />
	  <for count="system_group_size" from="0" name="i">
	    <integer array_index="i" name="system_group_id" />
	    <string array_index="i" name="system_group_name" />
	    <string array_index="i" name="system_group_info" />
	  </for>
	  <integer name="allowed_fids_size" />
	  <for count="allowed_fids_size" from="0" name="i">
	    <integer array_index="i" name="id" />
	    <string array_index="i" name="name" />
	    <string array_index="i" name="module" />
	  </for>
	  <integer name="not_allowed_fids_size" />
	  <for count="not_allowed_fids_size" from="0" name="i">
	    <integer array_index="i" name="id_not_allowed" />
	    <string array_index="i" name="name_not_allowed" />
	    <string array_index="i" name="module_not_allowed" />
          </for>
	  <string name="fullname" />
	  <string name="org_name" />
      </output>
    </function>


*/
