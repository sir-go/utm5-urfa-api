package fn

// rpcf_dealer_get_tech_param_by_uid
func x70000075(c conn, p Dict) Dict {
	c.Call(0x70000075)
	putI(c, p, "user_id")
	putI(c, p, "account_id")
	c.Send()

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x70000075" name="rpcf_dealer_get_tech_param_by_uid">
        <input>
            <integer name="user_id" />
            <integer name="account_id" />
        </input>
        <output>
            <integer name="slinks_count" />
            <for count="slinks_count" from="0" name="s">
                <integer array_index="s" name="params_count" />
                <for count="params_count" from="0" name="p">
                    <integer array_index="s,p" name="id" />
                    <integer array_index="s,p" name="type_id" />
                    <string array_index="s,p" name="type_name" />
                    <string array_index="s,p" name="param" />
                    <integer array_index="s,p" name="reg_date" />
                    <integer array_index="s,p" name="slink_id" />
                    <string array_index="s,p" name="service_name" />
                    <string array_index="s,p" name="passwd" />
                </for>
            </for>
        </output>
    </function>


*/
