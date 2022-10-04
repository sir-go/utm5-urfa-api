package fn

// rpcf_get_contract_type_list
func x4801(c conn, _ Dict) Dict {
	c.Call(0x4801)

	// fixme: output has a complex param
	panic(NotImplemented)
	return nil
}

/* xml:
<function id="0x4801" name="rpcf_get_contract_type_list">
		<input>
		</input>
		<output>
            <integer name="size" />
          <for count="size" from="0" name="i">
                <integer array_index="i" name="id" />
                <string array_index="i" name="name" />
                <integer array_index="i" name="create_date" />
                <integer array_index="i" name="change_date" />
          </for>
		</output>
	</function>


*/
