package fn

// rpcf_whoami
func x4417(c conn, _ Dict) Dict {
	c.Call(0x4417)

	return Dict{
		"my_uid":     c.GetI(),
		"login":      c.GetS(),
		"user_ip4":   c.GetA(),
		"user_mask4": c.GetI(),
		"user_ip6":   c.GetA(),
		"user_mask6": c.GetI(),
		"system_groups": getMapOf(c, func() Dict {
			return Dict{
				"id":   c.GetI(),
				"name": c.GetS(),
				"info": c.GetS(),
			}
		}),
		"allowed_fids": getMapOf(c, func() Dict {
			return Dict{
				"id":     c.GetI(),
				"name":   c.GetS(),
				"module": c.GetS(),
			}
		}),
		"not_allowed": getMapOf(c, func() Dict {
			return Dict{
				"id":     c.GetI(),
				"name":   c.GetS(),
				"module": c.GetS(),
			}
		}),
	}
}

/* xml:
<function id="0x4417" name="rpcf_whoami">
    <input />
    <output>
      <integer name="my_uid" />
      <string name="login" />
      <ip_address name="user_ip4" />
      <integer name="user_mask4" />
      <ip_address name="user_ip6" />
      <integer name="user_mask6" />
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
      <integer name="not_allowed_size" />
      <for count="not_allowed_size" from="0" name="i">
        <integer array_index="i" name="id_not_allowed" />
        <string array_index="i" name="name_not_allowed" />
        <string array_index="i" name="module_not_allowed" />
      </for>
    </output>
  </function>


*/
