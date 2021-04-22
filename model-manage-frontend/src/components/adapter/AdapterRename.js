import { Edit, SimpleForm, TextInput, TextField } from "react-admin";

const AdapterRename = (props) => {
  return (
    <Edit title="Rename Adapter" {...props}>
      <SimpleForm>
        <TextInput source="name" />
        <TextField source="dev_id" />
        <TextField source="model_name" />
      </SimpleForm>
    </Edit>
  );
};

export default AdapterRename;
