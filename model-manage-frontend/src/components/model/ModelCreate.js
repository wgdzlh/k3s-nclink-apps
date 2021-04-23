import { Create, SimpleForm, TextInput } from "react-admin";

const ModelCreate = (props) => {
  return (
    <Create title="Create Model" {...props}>
      <SimpleForm>
        <TextInput source="id" />
        <TextInput multiline fullWidth source="def" />
      </SimpleForm>
    </Create>
  );
};

export default ModelCreate;
