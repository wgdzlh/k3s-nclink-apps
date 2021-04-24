import { Create, SimpleForm, TextInput, required } from "react-admin";

const ModelCreate = (props) => {
  return (
    <Create title="Create Model" {...props}>
      <SimpleForm>
        <TextInput source="id" validate={[required()]} />
        <TextInput multiline fullWidth source="def" />
      </SimpleForm>
    </Create>
  );
};

export default ModelCreate;
