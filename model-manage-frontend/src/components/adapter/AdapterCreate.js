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
        <TextInput source="name" />
        <TextInput source="dev_id" />
        <ReferenceInput source="model_id" reference="models">
          <SelectInput />
        </ReferenceInput>
      </SimpleForm>
    </Create>
  );
};

export default AdapterCreate;
