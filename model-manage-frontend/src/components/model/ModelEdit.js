import { Edit, SimpleForm, TextInput } from "react-admin";

const ModelEdit = (props) => {
  return (
    <Edit title="Edit Model" {...props}>
      <SimpleForm>
        <TextInput disabled source="id" />
        <TextInput multiline fullWidth source="def" />
      </SimpleForm>
    </Edit>
  );
};

export default ModelEdit;
