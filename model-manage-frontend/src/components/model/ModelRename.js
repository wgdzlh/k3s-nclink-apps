import { Edit, SimpleForm, TextInput, TextField } from "react-admin";

const ModelRename = (props) => {
  return (
    <Edit title="Rename Model" {...props}>
      <SimpleForm>
        <TextInput source="name" />
        <TextField source="def" />
      </SimpleForm>
    </Edit>
  );
};

export default ModelRename;
