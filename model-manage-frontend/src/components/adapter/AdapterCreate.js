import {
  Create,
  SimpleForm,
  TextInput,
  ReferenceInput,
  SelectInput,
  required,
} from "react-admin";

const AdapterCreate = (props) => {
  return (
    <Create title="Create Adapter" {...props}>
      <SimpleForm>
        <TextInput source="id" validate={[required()]} />
        <TextInput source="dev_id" />
        <ReferenceInput source="model_id" reference="models">
          <SelectInput optionText="id" />
        </ReferenceInput>
      </SimpleForm>
    </Create>
  );
};

export default AdapterCreate;
