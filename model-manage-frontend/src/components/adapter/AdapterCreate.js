import {
  Create,
  SimpleForm,
  TextInput,
  ReferenceInput,
  SelectInput,
} from "react-admin";

const AdapterCreate = (props) => {
  return (
    <Create title="Create Adapter" {...props}>
      <SimpleForm>
        <TextInput source="id" />
        <TextInput source="dev_id" />
        <ReferenceInput source="model_id" reference="models">
          <SelectInput optionText="id" />
        </ReferenceInput>
      </SimpleForm>
    </Create>
  );
};

export default AdapterCreate;
